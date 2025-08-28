package v1

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/brevdev/cloud/pkg/ssh"
)

type CloudMachineImage interface {
	GetImages(ctx context.Context, args GetImageArgs) ([]Image, error)
}

type GetImageArgs struct {
	Owners        []string // self, amazon, aws-marketplace, <account id>, project id for GCP
	Architectures []string // i386, x86_64, arm64
	NameFilters   []string // name of the image (wildcard permitted)
	ImageIDs      []string
}

type Image struct {
	ID           string
	Architecture string
	Description  string
	Name         string
	CreatedAt    time.Time
}

func ValidateInstanceImage(ctx context.Context, instance Instance, privateKey string) error {
	sshClient, err := connectToInstance(ctx, instance, privateKey)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := sshClient.Close(); closeErr != nil {
			fmt.Printf("warning: failed to close SSH connection: %v\n", closeErr)
		}
	}()

	arch, err := validateArchitecture(ctx, sshClient)
	if err != nil {
		return err
	}

	osVersion, err := validateOSVersion(ctx, sshClient)
	if err != nil {
		return err
	}

	homeDir, err := validateHomeDirectory(ctx, sshClient, instance.SSHUser)
	if err != nil {
		return err
	}

	systemdStatus, err := validateSystemd(ctx, sshClient)
	if err != nil {
		return err
	}

	fmt.Printf("Instance image validation passed for %s: architecture=%s, os=%s, home=%s, systemd=%s\n",
		instance.CloudID, arch, osVersion, homeDir, systemdStatus)

	return nil
}

func connectToInstance(ctx context.Context, instance Instance, privateKey string) (*ssh.Client, error) {
	if instance.SSHUser == "" {
		return nil, fmt.Errorf("SSH user is not set for instance %s", instance.CloudID)
	}
	if instance.SSHPort == 0 {
		return nil, fmt.Errorf("SSH port is not set for instance %s", instance.CloudID)
	}
	if instance.PublicIP == "" {
		return nil, fmt.Errorf("public IP is not available for instance %s", instance.CloudID)
	}

	sshClient, err := ssh.ConnectToHost(ctx, ssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: fmt.Sprintf("%s:%d", instance.PublicIP, instance.SSHPort),
		PrivKey:  privateKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to instance via SSH: %w", err)
	}
	return sshClient, nil
}

func validateArchitecture(ctx context.Context, sshClient *ssh.Client) (string, error) {
	stdout, stderr, err := sshClient.RunCommand(ctx, "uname -m")
	if err != nil {
		return "", fmt.Errorf("failed to check architecture: %w, stdout: %s, stderr: %s", err, stdout, stderr)
	}
	arch := strings.TrimSpace(stdout)
	if !strings.Contains(arch, "x86_64") {
		return "", fmt.Errorf("expected x86_64 architecture, got: %s", arch)
	}
	return "x86_64", nil
}

func validateOSVersion(ctx context.Context, sshClient *ssh.Client) (string, error) {
	stdout, stderr, err := sshClient.RunCommand(ctx, "cat /etc/os-release | grep PRETTY_NAME")
	if err != nil {
		return "", fmt.Errorf("failed to check OS version: %w, stdout: %s, stderr: %s", err, stdout, stderr)
	}

	parts := strings.Split(strings.TrimSpace(stdout), "=")
	if len(parts) != 2 {
		return "", fmt.Errorf("error: os pretty name not in format PRETTY_NAME=\"Ubuntu\": %s", stdout)
	}

	osVersion := strings.Trim(parts[1], "\"")
	ubuntuRegex := regexp.MustCompile(`Ubuntu 20\.04|22\.04`)
	if !ubuntuRegex.MatchString(osVersion) {
		return "", fmt.Errorf("expected Ubuntu 20.04 or 22.04, got: %s", osVersion)
	}
	return osVersion, nil
}

func validateHomeDirectory(ctx context.Context, sshClient *ssh.Client, sshUser string) (string, error) {
	stdout, stderr, err := sshClient.RunCommand(ctx, "cd ~ && pwd")
	if err != nil {
		return "", fmt.Errorf("failed to check home directory: %w, stdout: %s, stderr: %s", err, stdout, stderr)
	}

	homeDir := strings.TrimSpace(stdout)
	if sshUser == "ubuntu" {
		if !strings.Contains(homeDir, "/home/ubuntu") {
			return "", fmt.Errorf("expected ubuntu user home directory to contain /home/ubuntu, got: %s", homeDir)
		}
	} else {
		if !strings.Contains(homeDir, "/root") && !strings.Contains(homeDir, fmt.Sprintf("/home/%v", sshUser)) {
			return "", fmt.Errorf("expected non-ubuntu user home directory to contain /root, got: %s", homeDir)
		}
	}
	return homeDir, nil
}

func validateSystemd(ctx context.Context, sshClient *ssh.Client) (string, error) {
	stdout, stderr, err := sshClient.RunCommand(ctx, "systemctl is-system-running")

	systemdStatus := strings.TrimSpace(stdout)
	if systemdStatus == "running" || systemdStatus == "degraded" {
		return systemdStatus, nil
	}

	if err != nil {
		return "", fmt.Errorf("failed to check systemd status: %w, stdout: %s, stderr: %s", err, stdout, stderr)
	}

	return "", fmt.Errorf("expected systemd to be running or degraded, got: %s", systemdStatus)
}
