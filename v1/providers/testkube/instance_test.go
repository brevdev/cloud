package v1

import (
	"context"
	"errors"
	"testing"

	cloudv1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateInstanceProvisionFailures(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		name         string
		instanceType string
		expectedErr  error
	}{
		{
			name:         "capacity",
			instanceType: InstanceTypeFailCapacity,
			expectedErr:  cloudv1.ErrInsufficientResources,
		},
		{
			name:         "quota",
			instanceType: InstanceTypeFailQuota,
			expectedErr:  cloudv1.ErrOutOfQuota,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestClient(t)

			instance, err := client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
				RefID:        tc.name,
				Name:         tc.name,
				InstanceType: tc.instanceType,
			})
			require.Nil(t, instance)
			require.ErrorIs(t, err, tc.expectedErr)

			pods, err := client.k8sClient.CoreV1().Pods(client.namespace).List(ctx, metav1.ListOptions{})
			require.NoError(t, err)
			require.Empty(t, pods.Items)

			services, err := client.k8sClient.CoreV1().Services(client.namespace).List(ctx, metav1.ListOptions{})
			require.NoError(t, err)
			require.Empty(t, services.Items)
		})
	}
}

func TestInstanceLifecycle(t *testing.T) { //nolint:funlen // test ok
	ctx := context.Background()
	client := newTestClient(t)

	instance, err := client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        "env-1",
		Name:         "dev env",
		InstanceType: InstanceTypeOKCPU,
		PublicKey:    "ssh-rsa test",
		Tags: cloudv1.Tags{
			"purpose": "test",
		},
	})
	require.NoError(t, err)
	require.Equal(t, cloudv1.LifecycleStatusPending, instance.Status.LifecycleStatus)
	require.Equal(t, "env-1", instance.RefID)
	require.Equal(t, "test-credential", instance.CloudCredRefID)
	require.Equal(t, InstanceTypeOKCPU, instance.InstanceType)
	spec, ok := getInstanceTypeSpec(InstanceTypeOKCPU)
	require.True(t, ok)
	require.Equal(t, spec.imageID, instance.ImageID)

	listed, err := client.ListInstances(ctx, cloudv1.ListInstancesArgs{
		TagFilters: map[string][]string{
			"purpose": {"test"},
		},
	})
	require.NoError(t, err)
	require.Len(t, listed, 1)

	require.ErrorIs(t, client.StopInstance(ctx, instance.CloudID), cloudv1.ErrNotImplemented)
	require.ErrorIs(t, client.StartInstance(ctx, instance.CloudID), cloudv1.ErrNotImplemented)
	require.ErrorIs(t, client.RebootInstance(ctx, instance.CloudID), cloudv1.ErrNotImplemented)
	setPodReady(t, client, instance.CloudID)

	pendingLB, err := client.GetInstance(ctx, instance.CloudID)
	require.NoError(t, err)
	require.Equal(t, cloudv1.LifecycleStatusPending, pendingLB.Status.LifecycleStatus)
	require.Contains(t, pendingLB.Status.Messages, "service "+string(instance.CloudID)+" waiting for load balancer ingress")

	setServiceLoadBalancerIngress(t, client, instance.CloudID, "203.0.113.10", "")

	running, err := client.GetInstance(ctx, instance.CloudID)
	require.NoError(t, err)
	require.Equal(t, cloudv1.LifecycleStatusRunning, running.Status.LifecycleStatus)
	require.Equal(t, string(instance.CloudID), running.Hostname)
	require.Equal(t, "203.0.113.10", running.PublicIP)
	require.Equal(t, "203.0.113.10", running.PublicDNS)
	require.Equal(t, 22, running.SSHPort)

	require.NoError(t, client.UpdateInstanceTags(ctx, cloudv1.UpdateInstanceTagsArgs{
		InstanceID: instance.CloudID,
		Tags: cloudv1.Tags{
			"purpose": "updated",
		},
	}))
	updated, err := client.GetInstance(ctx, instance.CloudID)
	require.NoError(t, err)
	require.Equal(t, "updated", updated.Tags["purpose"])

	require.NoError(t, client.TerminateInstance(ctx, instance.CloudID))
	_, err = client.GetInstance(ctx, instance.CloudID)
	require.True(t, errors.Is(err, cloudv1.ErrInstanceNotFound))
}

func TestScenarioEnvironment(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	instance, err := client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        "build",
		Name:         "build",
		InstanceType: InstanceTypeFailBuild,
	})
	require.NoError(t, err)

	pod, err := client.k8sClient.CoreV1().Pods(client.namespace).Get(ctx, string(instance.CloudID), metav1.GetOptions{})
	require.NoError(t, err)
	service, err := client.k8sClient.CoreV1().Services(client.namespace).Get(ctx, string(instance.CloudID), metav1.GetOptions{})
	require.NoError(t, err)
	spec, ok := getInstanceTypeSpec(InstanceTypeFailBuild)
	require.True(t, ok)
	require.Equal(t, spec.serviceType, service.Spec.Type)
	require.Zero(t, service.Spec.Ports[0].NodePort)
	container := pod.Spec.Containers[0]
	require.Equal(t, spec.image, container.Image)
	require.Zero(t, container.Ports[0].HostPort)
	envByName := envMap(container.Env)
	require.Equal(t, "fail.build", envByName[envScenario])
	require.Equal(t, "true", envByName[envFailBuild])
}

func TestInstanceUsesBakedImageSpec(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	instance, err := client.CreateInstance(ctx, cloudv1.CreateInstanceAttrs{
		RefID:        "image-spec",
		Name:         "image spec",
		InstanceType: InstanceTypeOKCPU,
	})
	require.NoError(t, err)
	spec, ok := getInstanceTypeSpec(InstanceTypeOKCPU)
	require.True(t, ok)
	require.Equal(t, spec.imageID, instance.ImageID)

	pod, err := client.k8sClient.CoreV1().Pods(client.namespace).Get(ctx, string(instance.CloudID), metav1.GetOptions{})
	require.NoError(t, err)
	container := pod.Spec.Containers[0]
	require.Equal(t, spec.image, container.Image)
	require.NotNil(t, container.ReadinessProbe)
	require.NotNil(t, container.ReadinessProbe.TCPSocket)
	require.Equal(t, containerSSHPort, container.ReadinessProbe.TCPSocket.Port.IntVal)
	for _, mount := range container.VolumeMounts {
		require.NotEqual(t, "/sys/fs/cgroup", mount.MountPath)
	}
}

func TestPopulateNetworkLoadBalancer(t *testing.T) {
	instance := &cloudv1.Instance{}
	populateNetwork(&corev1.Service{
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeLoadBalancer,
			ClusterIP: "10.96.119.41",
			Ports: []corev1.ServicePort{
				{
					Port: 22,
				},
			},
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{Hostname: "testkube.example.com"},
				},
			},
		},
	}, instance)

	require.Equal(t, "10.96.119.41", instance.PrivateIP)
	require.Equal(t, "testkube.example.com", instance.PublicIP)
	require.Equal(t, "testkube.example.com", instance.PublicDNS)
	require.Equal(t, 22, instance.SSHPort)
}

func setPodReady(t *testing.T, client *TestKubeClient, instanceID cloudv1.CloudProviderInstanceID) {
	t.Helper()

	pod, err := client.k8sClient.CoreV1().Pods(client.namespace).Get(context.Background(), string(instanceID), metav1.GetOptions{})
	require.NoError(t, err)
	pod.Status.Phase = corev1.PodRunning
	pod.Status.Conditions = []corev1.PodCondition{
		{
			Type:   corev1.PodReady,
			Status: corev1.ConditionTrue,
		},
	}
	_, err = client.k8sClient.CoreV1().Pods(client.namespace).UpdateStatus(context.Background(), pod, metav1.UpdateOptions{})
	require.NoError(t, err)
}

func setServiceLoadBalancerIngress(t *testing.T, client *TestKubeClient, instanceID cloudv1.CloudProviderInstanceID, ip, hostname string) {
	t.Helper()

	service, err := client.k8sClient.CoreV1().Services(client.namespace).Get(context.Background(), string(instanceID), metav1.GetOptions{})
	require.NoError(t, err)
	service.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{
			IP:       ip,
			Hostname: hostname,
		},
	}
	_, err = client.k8sClient.CoreV1().Services(client.namespace).UpdateStatus(context.Background(), service, metav1.UpdateOptions{})
	require.NoError(t, err)
}

func envMap(envVars []corev1.EnvVar) map[string]string {
	envByName := map[string]string{}
	for _, envVar := range envVars {
		envByName[envVar.Name] = envVar.Value
	}
	return envByName
}
