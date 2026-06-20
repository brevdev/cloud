package v1

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	cloudv1 "github.com/brevdev/cloud/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	CloudProviderID = "test-kubernetes"

	DefaultNamespace = "default"
	DefaultLocation  = "test-local"

	servicePortName  = "ssh"
	servicePort      = int32(22)
	containerSSHPort = int32(22)
)

type TestKubeAuthMode string

const (
	TestKubeAuthModeKubeconfig TestKubeAuthMode = "kubeconfig"
	TestKubeAuthModeInCluster  TestKubeAuthMode = "in-cluster"
)

// TestKubeCredential authenticates a developer test provider backed by Kubernetes.
type TestKubeCredential struct {
	RefID            string
	AuthMode         TestKubeAuthMode
	KubeconfigBase64 string
	Namespace        string
}

var _ cloudv1.CloudCredential = &TestKubeCredential{}

var restInClusterConfig = rest.InClusterConfig

func NewTestKubeCredential(refID, kubeconfigBase64, namespace string) *TestKubeCredential {
	return &TestKubeCredential{
		RefID:            refID,
		AuthMode:         TestKubeAuthModeKubeconfig,
		KubeconfigBase64: kubeconfigBase64,
		Namespace:        namespace,
	}
}

func NewInClusterTestKubeCredential(refID, namespace string) *TestKubeCredential {
	return &TestKubeCredential{
		RefID:     refID,
		AuthMode:  TestKubeAuthModeInCluster,
		Namespace: namespace,
	}
}

func (c *TestKubeCredential) GetReferenceID() string {
	return c.RefID
}

func (c *TestKubeCredential) GetAPIType() cloudv1.APIType {
	return cloudv1.APITypeGlobal
}

func (c *TestKubeCredential) GetCloudProviderID() cloudv1.CloudProviderID {
	return CloudProviderID
}

func (c *TestKubeCredential) GetTenantID() (string, error) {
	authMode, err := c.validateAuthMode()
	if err != nil {
		return "", err
	}
	var fingerprint string
	switch authMode {
	case TestKubeAuthModeKubeconfig:
		fingerprint = "kubeconfig:" + c.KubeconfigBase64 + ":" + c.Namespace
	case TestKubeAuthModeInCluster:
		fingerprint = "in-cluster:" + c.Namespace
	default:
		return "", fmt.Errorf("unknown testkube auth mode: %s", authMode)
	}
	return fmt.Sprintf("%s-%x", CloudProviderID, sha256.Sum256([]byte(fingerprint))), nil
}

func (c *TestKubeCredential) MakeClient(_ context.Context, location string) (cloudv1.CloudClient, error) {
	restConfig, err := c.restConfig()
	if err != nil {
		return nil, err
	}

	opts := []TestKubeClientOption{
		WithNamespace(c.Namespace),
		WithLocation(firstNonEmpty(location, DefaultLocation)),
	}
	return NewTestKubeClient(c.RefID, restConfig, opts...)
}

func (c *TestKubeCredential) restConfig() (*rest.Config, error) {
	authMode, err := c.validateAuthMode()
	if err != nil {
		return nil, err
	}
	switch authMode {
	case TestKubeAuthModeKubeconfig:
		if c.KubeconfigBase64 == "" {
			return nil, fmt.Errorf("kubeconfigBase64 is required")
		}
		kubeconfig, err := base64.StdEncoding.DecodeString(c.KubeconfigBase64)
		if err != nil {
			return nil, fmt.Errorf("decode kubeconfig: %w", err)
		}
		return clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	case TestKubeAuthModeInCluster:
		return restInClusterConfig()
	default:
		return nil, fmt.Errorf("unknown testkube auth mode: %s", authMode)
	}
}

func (c *TestKubeCredential) authMode() TestKubeAuthMode {
	authMode := TestKubeAuthMode(strings.TrimSpace(string(c.AuthMode)))
	if authMode == "" {
		return TestKubeAuthModeKubeconfig
	}
	return authMode
}

func (c *TestKubeCredential) validateAuthMode() (TestKubeAuthMode, error) {
	authMode := c.authMode()
	switch authMode {
	case TestKubeAuthModeKubeconfig:
		return authMode, nil
	case TestKubeAuthModeInCluster:
		if c.KubeconfigBase64 != "" {
			return "", fmt.Errorf("kubeconfigBase64 must be empty when authMode is %q", authMode)
		}
		return authMode, nil
	default:
		return "", fmt.Errorf("unknown testkube auth mode: %s", authMode)
	}
}

// TestKubeClient implements the CloudClient interface with Kubernetes primitives.
type TestKubeClient struct {
	cloudv1.NotImplCloudClient

	refID     string
	namespace string
	location  string
	k8sClient kubernetes.Interface
}

var _ cloudv1.CloudClient = &TestKubeClient{}

type testKubeClientOptions struct {
	namespace string
	location  string
	k8sClient kubernetes.Interface
}

type TestKubeClientOption func(*testKubeClientOptions) error

func WithNamespace(namespace string) TestKubeClientOption {
	return func(options *testKubeClientOptions) error {
		options.namespace = namespace
		return nil
	}
}

func WithLocation(location string) TestKubeClientOption {
	return func(options *testKubeClientOptions) error {
		options.location = location
		return nil
	}
}

func WithKubernetesClient(k8sClient kubernetes.Interface) TestKubeClientOption {
	return func(options *testKubeClientOptions) error {
		options.k8sClient = k8sClient
		return nil
	}
}

func NewTestKubeClient(refID string, config *rest.Config, opts ...TestKubeClientOption) (*TestKubeClient, error) {
	options := testKubeClientOptions{
		namespace: DefaultNamespace,
		location:  DefaultLocation,
	}
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}
	if strings.TrimSpace(refID) == "" {
		return nil, fmt.Errorf("refID is required")
	}
	if strings.TrimSpace(options.namespace) == "" {
		options.namespace = DefaultNamespace
	}
	if strings.TrimSpace(options.location) == "" {
		options.location = DefaultLocation
	}
	if options.k8sClient == nil {
		if config == nil {
			return nil, fmt.Errorf("kubernetes rest config is required")
		}
		k8sClient, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, fmt.Errorf("create kubernetes client: %w", err)
		}
		options.k8sClient = k8sClient
	}

	return &TestKubeClient{
		refID:     refID,
		namespace: options.namespace,
		location:  options.location,
		k8sClient: options.k8sClient,
	}, nil
}

func (c *TestKubeClient) GetAPIType() cloudv1.APIType {
	return cloudv1.APITypeGlobal
}

func (c *TestKubeClient) GetCloudProviderID() cloudv1.CloudProviderID {
	return CloudProviderID
}

func (c *TestKubeClient) GetReferenceID() string {
	return c.refID
}

func (c *TestKubeClient) GetTenantID() (string, error) {
	return fmt.Sprintf("%s-%x", CloudProviderID, sha256.Sum256([]byte(c.refID+c.namespace))), nil
}

func (c *TestKubeClient) MakeClient(_ context.Context, location string) (cloudv1.CloudClient, error) {
	if location != "" {
		c.location = location
	}
	return c, nil
}

func (c *TestKubeClient) GetInstancePollTime() time.Duration {
	return time.Second
}

func (c *TestKubeClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
