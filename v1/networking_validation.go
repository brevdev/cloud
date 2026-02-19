package v1

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/brevdev/cloud/internal/ssh"
)

const (
	// DefaultFirewallTestPort is the port used for testing firewall rules
	// This port should NOT be in the allowed ingress rules
	DefaultFirewallTestPort = 9999

	// FirewallTestTimeout is the timeout for testing port accessibility
	FirewallTestTimeout = 10 * time.Second

	// PortConnectionTimeout is the timeout for a single connection attempt
	PortConnectionTimeout = 5 * time.Second
)

// ValidateFirewallBlocksPort validates that a port is NOT accessible from outside the instance.
// This is used to verify that firewall rules (UFW, iptables) are working correctly.
func ValidateFirewallBlocksPort(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string, port int) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	// First, start a test server on the instance
	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	// Start a simple HTTP server on 0.0.0.0:port in the background
	// The server will respond with "OK" to any request
	startServerCmd := fmt.Sprintf(
		"nohup sh -c 'echo -e \"HTTP/1.1 200 OK\\r\\nContent-Length: 2\\r\\n\\r\\nOK\" | nc -l -p %d' > /dev/null 2>&1 &",
		port,
	)
	_, _, err = sshClient.RunCommand(ctx, startServerCmd)
	if err != nil {
		return fmt.Errorf("failed to start test server on instance: %w", err)
	}

	// Give the server a moment to start
	time.Sleep(500 * time.Millisecond)

	// Now try to connect to the port from outside - this should FAIL
	err = checkPortBlocked(ctx, publicIP, port)
	if err != nil {
		return err
	}

	// Clean up: kill any remaining nc processes on that port
	killCmd := fmt.Sprintf("pkill -f 'nc -l -p %d' || true", port)
	_, _, _ = sshClient.RunCommand(ctx, killCmd)

	return nil
}

// ValidateDockerFirewallBlocksPort validates that a Docker container listening on 0.0.0.0
// is NOT accessible from outside the instance due to DOCKER-USER iptables rules.
func ValidateDockerFirewallBlocksPort(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string, port int) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	dockerCmd, err := setupDockerCommand(ctx, sshClient, instance.CloudID)
	if err != nil {
		return err
	}

	// Start a Docker container with a simple HTTP server
	// Using nginx as it's commonly available
	containerName := fmt.Sprintf("firewall-test-%d", port)
	startDockerCmd := fmt.Sprintf(
		"%s run -d --rm --name %s -p %d:%d nginx:alpine",
		dockerCmd, containerName, port, 80,
	)
	_, stderr, err := sshClient.RunCommand(ctx, startDockerCmd)
	if err != nil {
		return fmt.Errorf("failed to start docker container: %w, stderr: %s", err, stderr)
	}

	// Wait for container to be running and service to be ready
	if err := waitForDockerService(ctx, sshClient, dockerCmd, containerName, port); err != nil {
		_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s rm -f %s || true", dockerCmd, containerName))
		return err
	}

	// Debug: show iptables rules for DOCKER-USER chain
	iptablesOut, _, _ := sshClient.RunCommand(ctx, "sudo iptables -L DOCKER-USER -n -v 2>&1 || echo 'DOCKER-USER chain not found'")
	fmt.Printf("Instance %s DOCKER-USER iptables rules:\n%s\n", publicIP, iptablesOut)

	// Now try to connect to the port from outside - this should FAIL due to iptables rules
	fmt.Printf("Testing external connectivity to %s:%d\n", publicIP, port)
	err = checkPortBlocked(ctx, publicIP, port)

	// Clean up: stop and remove the container
	stopDockerCmd := fmt.Sprintf("%s rm -f %s || true", dockerCmd, containerName)
	_, _, _ = sshClient.RunCommand(ctx, stopDockerCmd)

	if err != nil {
		return err
	}

	return nil
}

func ValidateDockerFirewallAllowsEgress(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	dockerCmd, err := setupDockerCommand(ctx, sshClient, instance.CloudID)
	if err != nil {
		return err
	}

	// Pull the alpine image
	cmd := fmt.Sprintf(
		"%s pull alpine",
		dockerCmd,
	)
	_, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to pull alpine image: %w, stderr: %s", err, stderr)
	}

	// Start a Docker container to ping Google's DNS server
	cmd = fmt.Sprintf(
		"%s run --rm alpine ping -c 3 8.8.8.8",
		dockerCmd,
	)
	stdout, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to connect to Google's DNS server: %w, stderr: %s", err, stderr)
	}
	if !strings.Contains(stdout, "3 packets transmitted, 3 packets received") {
		return fmt.Errorf("expected successful ping, got: %s", stdout)
	}

	return nil
}

func ValidateDockerFirewallAllowsContainerToContainerCommunication(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	dockerCmd, err := setupDockerCommand(ctx, sshClient, instance.CloudID)
	if err != nil {
		return err
	}

	// Create a docker network
	networkName := "firewall-test-network"
	cmd := fmt.Sprintf(
		"%s network create %s",
		dockerCmd, networkName,
	)
	_, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to create docker network: %w, stderr: %s", err, stderr)
	}

	// Pull the alpine image
	cmd = fmt.Sprintf(
		"%s pull alpine",
		dockerCmd,
	)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to pull alpine image: %w, stderr: %s", err, stderr)
	}

	// Pull the nginx image
	cmd = fmt.Sprintf(
		"%s pull nginx:alpine",
		dockerCmd,
	)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to pull nginx image: %w, stderr: %s", err, stderr)
	}

	// Start a Docker container in the background
	containerName := "firewall-test-container-to-container"
	cmd = fmt.Sprintf(
		"%s run -d --name %s --network %s nginx:alpine",
		dockerCmd, containerName, networkName,
	)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to start docker container: %w, stderr: %s", err, stderr)
	}

	// Start a second Docker container to connect to the first container
	cmd = fmt.Sprintf(
		"%s run --network %s --rm alpine wget -q -O- http://%s",
		dockerCmd, networkName, containerName,
	)
	stdout, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to connect to nginx container: %w, stderr: %s", err, stderr)
	}

	if !strings.Contains(stdout, "Welcome to nginx") {
		return fmt.Errorf("expected successful wget, got: %s", stdout)
	}
	return nil
}

func ValidateMicroK8sFirewallAllowsEgress(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	microK8sCmd, err := setupMicroK8sCommand(ctx, sshClient, instance.CloudID)
	if err != nil {
		return err
	}

	// Ensure prior run artifacts do not interfere.
	_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete pod mk8s-egress-test --ignore-not-found=true", microK8sCmd))

	cmd := fmt.Sprintf(
		"%s kubectl run mk8s-egress-test --image=alpine:3.20 --restart=Never --command -- sh -c 'ping -c 3 8.8.8.8'",
		microK8sCmd,
	)
	_, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to create microk8s egress test pod: %w, stderr: %s", err, stderr)
	}

	defer func() {
		_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete pod mk8s-egress-test --ignore-not-found=true", microK8sCmd))
	}()

	cmd = fmt.Sprintf("%s kubectl wait --for=jsonpath='{.status.phase}'=Succeeded pod/mk8s-egress-test --timeout=180s", microK8sCmd)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		logsCmd := fmt.Sprintf("%s kubectl logs mk8s-egress-test 2>/dev/null || true", microK8sCmd)
		logs, _, _ := sshClient.RunCommand(ctx, logsCmd)
		return fmt.Errorf("microk8s egress test pod did not succeed: %w, stderr: %s, logs: %s", err, stderr, logs)
	}

	cmd = fmt.Sprintf("%s kubectl logs mk8s-egress-test", microK8sCmd)
	stdout, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to get microk8s egress test pod logs: %w, stderr: %s", err, stderr)
	}
	if !strings.Contains(stdout, "3 packets transmitted, 3 packets received") {
		return fmt.Errorf("expected successful pod egress ping, got logs: %s", stdout)
	}

	return nil
}

func ValidateMicroK8sFirewallAllowsPodToPodCommunication(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	microK8sCmd, err := setupMicroK8sCommand(ctx, sshClient, instance.CloudID)
	if err != nil {
		return err
	}

	// Ensure prior run artifacts do not interfere.
	_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete pod mk8s-nginx mk8s-c2c-test --ignore-not-found=true", microK8sCmd))
	_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete service mk8s-nginx-svc --ignore-not-found=true", microK8sCmd))

	cmd := fmt.Sprintf("%s kubectl run mk8s-nginx --image=nginx:alpine --restart=Never --port=80", microK8sCmd)
	_, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to create microk8s nginx pod: %w, stderr: %s", err, stderr)
	}

	defer func() {
		_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete pod mk8s-nginx mk8s-c2c-test --ignore-not-found=true", microK8sCmd))
		_, _, _ = sshClient.RunCommand(ctx, fmt.Sprintf("%s kubectl delete service mk8s-nginx-svc --ignore-not-found=true", microK8sCmd))
	}()

	cmd = fmt.Sprintf("%s kubectl wait --for=condition=Ready pod/mk8s-nginx --timeout=180s", microK8sCmd)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("microk8s nginx pod did not become ready: %w, stderr: %s", err, stderr)
	}

	cmd = fmt.Sprintf("%s kubectl expose pod mk8s-nginx --name=mk8s-nginx-svc --port=80 --target-port=80", microK8sCmd)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to create microk8s nginx service: %w, stderr: %s", err, stderr)
	}

	cmd = fmt.Sprintf(
		"%s kubectl run mk8s-c2c-test --image=alpine:3.20 --restart=Never --command -- sh -c 'wget -q -O- http://mk8s-nginx-svc'",
		microK8sCmd,
	)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to create microk8s pod-to-pod test pod: %w, stderr: %s", err, stderr)
	}

	cmd = fmt.Sprintf("%s kubectl wait --for=jsonpath='{.status.phase}'=Succeeded pod/mk8s-c2c-test --timeout=180s", microK8sCmd)
	_, stderr, err = sshClient.RunCommand(ctx, cmd)
	if err != nil {
		logsCmd := fmt.Sprintf("%s kubectl logs mk8s-c2c-test 2>/dev/null || true", microK8sCmd)
		logs, _, _ := sshClient.RunCommand(ctx, logsCmd)
		return fmt.Errorf("microk8s pod-to-pod test pod did not succeed: %w, stderr: %s, logs: %s", err, stderr, logs)
	}

	cmd = fmt.Sprintf("%s kubectl logs mk8s-c2c-test", microK8sCmd)
	stdout, stderr, err := sshClient.RunCommand(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to get microk8s pod-to-pod test pod logs: %w, stderr: %s", err, stderr)
	}
	if !strings.Contains(stdout, "Welcome to nginx") {
		return fmt.Errorf("expected successful pod-to-pod communication, got logs: %s", stdout)
	}

	return nil
}

// setupDockerCommand ensures Docker is available and returns the command to use (always with sudo)
func setupDockerCommand(ctx context.Context, sshClient *ssh.Client, instanceID CloudProviderInstanceID) (string, error) {
	// Check if Docker is available
	_, _, err := sshClient.RunCommand(ctx, "sudo docker --version")
	if err != nil {
		// Docker not installed, try installing it
		fmt.Printf("Docker not found, attempting to install on instance %s\n", instanceID)
		_, stderr, installErr := sshClient.RunCommand(ctx, "curl -fsSL https://get.docker.com | sudo sh")
		if installErr != nil {
			return "", fmt.Errorf("docker not available and failed to install: %w, stderr: %s", installErr, stderr)
		}
	}

	// Verify Docker works with sudo
	_, _, err = sshClient.RunCommand(ctx, "sudo docker ps")
	if err != nil {
		return "", fmt.Errorf("docker not accessible with sudo: %w", err)
	}
	return "sudo docker", nil
}

// setupMicroK8sCommand ensures MicroK8s is available and returns the command to use (always with sudo).
func setupMicroK8sCommand(ctx context.Context, sshClient *ssh.Client, instanceID CloudProviderInstanceID) (string, error) {
	checkCmd := "sudo microk8s status --wait-ready --timeout 120"
	_, _, err := sshClient.RunCommand(ctx, checkCmd)
	if err != nil {
		fmt.Printf("MicroK8s not found or not ready, attempting to install on instance %s\n", instanceID)
		_, stderr, installErr := sshClient.RunCommand(ctx, "sudo snap install microk8s --classic")
		if installErr != nil {
			return "", fmt.Errorf("microk8s not available and failed to install: %w, stderr: %s", installErr, stderr)
		}
		_, stderr, readyErr := sshClient.RunCommand(ctx, checkCmd)
		if readyErr != nil {
			return "", fmt.Errorf("microk8s installed but not ready: %w, stderr: %s", readyErr, stderr)
		}
	}

	_, stderr, err := sshClient.RunCommand(ctx, "sudo microk8s enable dns")
	if err != nil && !strings.Contains(stderr, "Nothing to do for dns") && !strings.Contains(stderr, "is already enabled") {
		return "", fmt.Errorf("failed to enable microk8s dns addon: %w, stderr: %s", err, stderr)
	}

	return "sudo microk8s", nil
}

// waitForDockerService waits for a Docker container's service to be ready and responding
func waitForDockerService(ctx context.Context, sshClient *ssh.Client, dockerCmd, containerName string, port int) error {
	for i := 0; i < 30; i++ { // Try for up to 30 seconds
		time.Sleep(1 * time.Second)

		// Check container is running
		checkContainerCmd := fmt.Sprintf("%s ps --filter name=%s --format '{{.Names}}'", dockerCmd, containerName)
		stdout, _, err := sshClient.RunCommand(ctx, checkContainerCmd)
		if err != nil || stdout == "" {
			continue
		}

		// Check if the port is listening inside the container (via localhost)
		checkPortCmd := fmt.Sprintf("curl -s -o /dev/null -w '%%{http_code}' --connect-timeout 2 http://localhost:%d/ || echo 'failed'", port)
		stdout, _, err = sshClient.RunCommand(ctx, checkPortCmd)
		if err == nil && stdout != "" && stdout != "failed" && stdout != "000" {
			fmt.Printf("Docker container ready after %d seconds, curl returned: %s\n", i+1, stdout)
			return nil
		}
	}
	return fmt.Errorf("docker container service did not become ready within 30 seconds")
}

// checkPortBlocked verifies that a port is NOT accessible from outside
// Returns nil if the port is blocked (expected), returns an error if the port is accessible
func checkPortBlocked(ctx context.Context, host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	// Try multiple times to be sure - if ANY connection succeeds, the port is accessible
	for attempt := 1; attempt <= 3; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, PortConnectionTimeout)

		dialer := net.Dialer{Timeout: PortConnectionTimeout}
		conn, err := dialer.DialContext(attemptCtx, "tcp", addr)
		cancel()

		if err != nil {
			fmt.Printf("checkPortBlocked attempt %d: connection to %s failed (expected): %v\n", attempt, addr, err)
			continue
		}

		// Connection succeeded - port is accessible, which is a problem
		_ = conn.Close()
		return fmt.Errorf("port %d is accessible from outside but should be blocked by firewall (attempt %d succeeded)", port, attempt)
	}

	// All attempts failed to connect - port is blocked as expected
	fmt.Printf("checkPortBlocked: confirmed port %d is blocked after 3 attempts\n", port)
	return nil
}

// checkPortAccessible verifies that a port IS accessible from outside
// Returns nil if the port is accessible, returns an error if the port is blocked
func checkPortAccessible(ctx context.Context, host string, port int) error {
	ctx, cancel := context.WithTimeout(ctx, PortConnectionTimeout)
	defer cancel()

	addr := fmt.Sprintf("%s:%d", host, port)

	dialer := net.Dialer{Timeout: PortConnectionTimeout}
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("port %d is not accessible: %w", port, err)
	}
	_ = conn.Close()
	return nil
}

// ValidateFirewallAllowsPort validates that a port IS accessible from outside the instance
// when it's in the allowed ingress rules.
func ValidateFirewallAllowsPort(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string, port int) error {
	var err error
	instance, err = WaitForInstanceLifecycleStatus(ctx, client, instance, LifecycleStatusRunning, PendingToRunningTimeout)
	if err != nil {
		return fmt.Errorf("failed to wait for instance running: %w", err)
	}

	publicIP := instance.PublicIP
	if publicIP == "" {
		return fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", publicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to SSH into instance: %w", err)
	}
	defer func() { _ = sshClient.Close() }()

	// Start a simple HTTP server on 0.0.0.0:port
	startServerCmd := fmt.Sprintf(
		"nohup sh -c 'while true; do echo -e \"HTTP/1.1 200 OK\\r\\nContent-Length: 2\\r\\n\\r\\nOK\" | nc -l -p %d; done' > /dev/null 2>&1 &",
		port,
	)
	_, _, err = sshClient.RunCommand(ctx, startServerCmd)
	if err != nil {
		return fmt.Errorf("failed to start test server on instance: %w", err)
	}

	time.Sleep(500 * time.Millisecond)

	// Try to connect - this should succeed for allowed ports
	err = checkPortAccessible(ctx, publicIP, port)

	// Clean up
	killCmd := fmt.Sprintf("pkill -f 'nc -l -p %d' || true", port)
	_, _, _ = sshClient.RunCommand(ctx, killCmd)

	if err != nil {
		return fmt.Errorf("allowed port %d is not accessible: %w", port, err)
	}

	return nil
}

// ValidateFirewallRules validates that firewall rules are working correctly by:
// 1. Checking that a non-allowed port is blocked
// 2. Checking that an allowed port (SSH) is accessible
func ValidateFirewallRules(ctx context.Context, client CloudInstanceReader, instance *Instance, privateKey string) error {
	var validationErr error

	// Validate that SSH port is accessible (should be allowed)
	err := checkPortAccessible(ctx, instance.PublicIP, instance.SSHPort)
	if err != nil {
		validationErr = errors.Join(validationErr, fmt.Errorf("SSH port should be accessible: %w", err))
	}

	// Validate that a non-standard port is blocked
	err = ValidateFirewallBlocksPort(ctx, client, instance, privateKey, DefaultFirewallTestPort)
	if err != nil {
		validationErr = errors.Join(validationErr, fmt.Errorf("firewall should block port %d: %w", DefaultFirewallTestPort, err))
	}

	return validationErr
}

// ValidateHTTPPortBlocked validates that an HTTP port is not accessible via HTTP request
func ValidateHTTPPortBlocked(ctx context.Context, host string, port int) error {
	ctx, cancel := context.WithTimeout(ctx, PortConnectionTimeout)
	defer cancel()

	url := fmt.Sprintf("http://%s:%d", host, port)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: PortConnectionTimeout}
	resp, err := client.Do(req)
	if err != nil {
		// Connection failed - expected behavior for blocked port
		return nil
	}
	defer func() { _ = resp.Body.Close() }()

	// If we got a response, the port is accessible - this is a problem
	return fmt.Errorf("HTTP port %d is accessible (status: %d) but should be blocked", port, resp.StatusCode)
}
