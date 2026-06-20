package v1

import (
	"context"
	"encoding/base64"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	internalssh "github.com/brevdev/cloud/internal/ssh"
	"github.com/brevdev/cloud/internal/validation"
	cloudv1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	gossh "golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	testKubeKubeconfigBase64EnvVar = "TESTKUBE_KUBECONFIG_BASE64"
	testKubeAuthModeEnvVar         = "TESTKUBE_AUTH_MODE"
	testKubeNamespaceEnvVar        = "TESTKUBE_NAMESPACE"
	testKubeLocationEnvVar         = "TESTKUBE_LOCATION"
)

var (
	validationSSHKeysOnce sync.Once
	validationSSHKeysErr  error
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()
	checkValidationSkip(t)

	location := testKubeLocation()
	config := validation.ProviderConfig{
		Credential: testKubeCredential(),
		Location:   location,
		StableIDs: []cloudv1.InstanceTypeID{
			cloudv1.InstanceTypeID(location + "-noSub-" + InstanceTypeOKCPU),
		},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	checkValidationSkip(t)

	location := testKubeLocation()
	config := validation.ProviderConfig{
		Credential: testKubeCredential(),
		Location:   location,
		StableIDs: []cloudv1.InstanceTypeID{
			cloudv1.InstanceTypeID(location + "-noSub-" + InstanceTypeOKCPU),
		},
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func TestFailureInstanceTypesValidation(t *testing.T) {
	t.Parallel()
	checkValidationSkip(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)

	client, err := testKubeCredential().MakeClient(ctx, testKubeLocation())
	require.NoError(t, err)

	_, err = client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        "validation-capacity",
		Name:         "validation-capacity",
		Location:     testKubeLocation(),
		InstanceType: InstanceTypeFailCapacity,
	})
	require.ErrorIs(t, err, cloudv1.ErrInsufficientResources)

	_, err = client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        "validation-quota",
		Name:         "validation-quota",
		Location:     testKubeLocation(),
		InstanceType: InstanceTypeFailQuota,
	})
	require.ErrorIs(t, err, cloudv1.ErrOutOfQuota)
}

func TestImageBackedInstanceValidation(t *testing.T) {
	checkValidationSkip(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	t.Cleanup(cancel)

	client := testKubeImageClient(t)
	refID := "validation-image-ok-cpu-" + strconv.FormatInt(time.Now().UnixNano(), 36)

	instance, err := client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        refID,
		Name:         refID,
		Location:     testKubeLocation(),
		InstanceType: InstanceTypeOKCPU,
		PublicKey:    testAuthorizedKey(t),
		Tags: cloudv1.Tags{
			"test": "image-validation",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, instance)
	t.Cleanup(func() {
		_ = client.TerminateInstance(context.Background(), instance.CloudID)
	})

	statefulSet, err := client.k8sClient.AppsV1().StatefulSets(client.namespace).Get(ctx, string(instance.CloudID), metav1.GetOptions{})
	require.NoError(t, err)
	require.Equal(t, DefaultImage, statefulSet.Spec.Template.Spec.Containers[0].Image)

	instance = waitForValidationInstanceStatus(ctx, t, client, instance.CloudID, cloudv1.LifecycleStatusRunning, 4*time.Minute)
	require.NotEmpty(t, instance.Hostname)
	require.NotEmpty(t, instance.PublicIP)
	require.NotZero(t, instance.SSHPort)

	sshCtx, cancelSSH := context.WithTimeout(ctx, 2*time.Minute)
	t.Cleanup(cancelSSH)
	require.NoError(t, internalssh.WaitForSSH(sshCtx, internalssh.ConnectionConfig{
		User:     instance.SSHUser,
		HostPort: net.JoinHostPort(instance.PublicIP, strconv.Itoa(instance.SSHPort)),
		PrivKey:  internalssh.DoNotUseDummyPrivateKey,
	}, internalssh.WaitForSSHOptions{
		Timeout:           90 * time.Second,
		ConnectionTimeout: 10 * time.Second,
		CheckCmd:          "sudo -n true && command -v apt-get && command -v systemctl && systemctl list-units --type=service --no-pager >/dev/null",
		WaitTime:          2 * time.Second,
	}))
}

func checkValidationSkip(t *testing.T) {
	t.Helper()

	kubeconfigBase64 := os.Getenv(testKubeKubeconfigBase64EnvVar)
	authMode := testKubeAuthMode()
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if authMode == TestKubeAuthModeKubeconfig && kubeconfigBase64 == "" && isValidationTest != "" {
		t.Fatalf("%s not set, but VALIDATION_TEST is set", testKubeKubeconfigBase64EnvVar)
	}
	if authMode == TestKubeAuthModeKubeconfig && kubeconfigBase64 == "" {
		t.Skipf("%s not set, skipping testkube validation tests", testKubeKubeconfigBase64EnvVar)
	}
	ensureValidationSSHKeys(t)
}

func testKubeCredential() *TestKubeCredential {
	if testKubeAuthMode() == TestKubeAuthModeInCluster {
		return NewInClusterTestKubeCredential("validation-test", testKubeNamespace())
	}
	return NewTestKubeCredential("validation-test", os.Getenv(testKubeKubeconfigBase64EnvVar), testKubeNamespace())
}

func testKubeImageClient(t *testing.T) *TestKubeClient {
	t.Helper()

	credential := testKubeCredential()
	restConfig, err := credential.restConfig()
	require.NoError(t, err)

	client, err := NewTestKubeClient(credential.RefID, restConfig,
		WithNamespace(testKubeNamespace()),
		WithLocation(testKubeLocation()),
	)
	require.NoError(t, err)
	return client
}

func testKubeNamespace() string {
	if namespace := os.Getenv(testKubeNamespaceEnvVar); namespace != "" {
		return namespace
	}
	return DefaultNamespace
}

func testKubeAuthMode() TestKubeAuthMode {
	if authMode := os.Getenv(testKubeAuthModeEnvVar); authMode != "" {
		return TestKubeAuthMode(authMode)
	}
	return TestKubeAuthModeKubeconfig
}

func testKubeLocation() string {
	if location := os.Getenv(testKubeLocationEnvVar); location != "" {
		return location
	}
	return DefaultLocation
}

func testAuthorizedKey(t *testing.T) string {
	t.Helper()

	authorizedKey, err := defaultAuthorizedKey()
	require.NoError(t, err)
	return authorizedKey
}

func defaultAuthorizedKey() (string, error) {
	signer, err := gossh.ParsePrivateKey([]byte(internalssh.DoNotUseDummyPrivateKey))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(gossh.MarshalAuthorizedKey(signer.PublicKey()))), nil
}

func ensureValidationSSHKeys(t *testing.T) {
	t.Helper()

	validationSSHKeysOnce.Do(func() {
		if os.Getenv("TEST_PRIVATE_KEY_BASE64") == "" {
			validationSSHKeysErr = os.Setenv(
				"TEST_PRIVATE_KEY_BASE64",
				base64.StdEncoding.EncodeToString([]byte(internalssh.DoNotUseDummyPrivateKey)),
			)
			if validationSSHKeysErr != nil {
				return
			}
		}

		if os.Getenv("TEST_PUBLIC_KEY_BASE64") == "" {
			authorizedKey, err := defaultAuthorizedKey()
			if err != nil {
				validationSSHKeysErr = err
				return
			}
			validationSSHKeysErr = os.Setenv(
				"TEST_PUBLIC_KEY_BASE64",
				base64.StdEncoding.EncodeToString([]byte(authorizedKey)),
			)
		}
	})
	require.NoError(t, validationSSHKeysErr)
}

func waitForValidationInstanceStatus(ctx context.Context, t *testing.T, client *TestKubeClient, instanceID cloudv1.CloudProviderInstanceID, status cloudv1.LifecycleStatus, timeout time.Duration) *cloudv1.Instance {
	t.Helper()

	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()

	var lastInstance *cloudv1.Instance
	var lastErr error
	for {
		instance, err := client.GetInstance(ctx, instanceID)
		if err != nil {
			lastErr = err
		} else {
			lastErr = nil
			lastInstance = instance
			if instance.Status.LifecycleStatus == status {
				return instance
			}
			if instance.Status.LifecycleStatus == cloudv1.LifecycleStatusFailed && status != cloudv1.LifecycleStatusFailed {
				t.Fatalf("instance %s failed while waiting for %s: %v", instanceID, status, instance.Status.Messages)
			}
		}

		select {
		case <-ctx.Done():
			require.NoError(t, lastErr)
			t.Fatalf("context ended waiting for instance %s to become %s: %v", instanceID, status, ctx.Err())
		case <-deadline.C:
			require.NoError(t, lastErr)
			if lastInstance != nil {
				t.Fatalf("instance %s status is %s, waiting for %s: %v", instanceID, lastInstance.Status.LifecycleStatus, status, lastInstance.Status.Messages)
			}
			t.Fatalf("timed out waiting for instance %s to become %s", instanceID, status)
		case <-tick.C:
		}
	}
}
