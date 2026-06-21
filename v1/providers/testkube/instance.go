package v1

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/units"
	cloudv1 "github.com/brevdev/cloud/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	labelName           = "app.kubernetes.io/name"
	labelManagedBy      = "app.kubernetes.io/managed-by"
	labelCloudID        = "testkube.brev.dev/cloud-id"
	labelLocation       = "testkube.brev.dev/location"
	labelNameValue      = "test-kubernetes"
	labelManagedByValue = "brev-cloud-sdk"

	annotationRefID          = "testkube.brev.dev/ref-id"
	annotationCloudCredRefID = "testkube.brev.dev/cloud-cred-ref-id" //nolint:gosec // this is a valid annotation
	annotationName           = "testkube.brev.dev/name"
	annotationLocation       = "testkube.brev.dev/location"
	annotationSubLocation    = "testkube.brev.dev/sub-location"
	annotationInstanceType   = "testkube.brev.dev/instance-type"
	annotationImageID        = "testkube.brev.dev/image-id"
	annotationCreatedAt      = "testkube.brev.dev/created-at"
	annotationScenario       = "testkube.brev.dev/scenario"
	annotationTagsJSON       = "testkube.brev.dev/tags-json"

	envInstanceType = "TESTKUBE_INSTANCE_TYPE"
	envScenario     = "TESTKUBE_SCENARIO"
	envFailBuild    = "TESTKUBE_FAIL_BUILD"
)

func (c *TestKubeClient) CreateInstance(ctx context.Context, attrs cloudv1.CreateInstanceAttrs) (*cloudv1.Instance, error) {
	if attrs.RefID == "" {
		return nil, fmt.Errorf("refID is required")
	}
	if attrs.InstanceType == "" {
		return nil, fmt.Errorf("instance type is required")
	}
	instanceTypeSpec, ok := getInstanceTypeSpec(attrs.InstanceType)
	if !ok {
		return nil, fmt.Errorf("unknown testkube instance type: %s", attrs.InstanceType)
	}

	// Immediate provision failures based on the incoming instance type.
	switch attrs.InstanceType {
	case InstanceTypeFailCapacity:
		return nil, cloudv1.ErrInsufficientResources
	case InstanceTypeFailQuota:
		return nil, cloudv1.ErrOutOfQuota
	}

	instance, err := c.createInstanceAsK8sResources(ctx, attrs, instanceTypeSpec)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *TestKubeClient) createInstanceAsK8sResources(ctx context.Context, attrs cloudv1.CreateInstanceAttrs, instanceTypeSpec instanceTypeSpec) (*cloudv1.Instance, error) {
	// Create a "cloud ID" to emulate a provider-provided instance ID.
	cloudID := makeCloudID(c.refID, attrs.RefID)

	location := c.resourceLocation(attrs)
	annotations := c.resourceAnnotations(cloudID, attrs, instanceTypeSpec)

	// Create the service.
	k8sService, err := c.k8sClient.
		CoreV1().
		Services(c.namespace).
		Create(ctx, &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:        string(cloudID),
				Namespace:   c.namespace,
				Labels:      objectLabels(string(cloudID), location),
				Annotations: annotations,
			},
			Spec: corev1.ServiceSpec{
				Type:     instanceTypeSpec.serviceType,
				Selector: selectorLabels(string(cloudID)),
				Ports: []corev1.ServicePort{
					{
						Name:       servicePortName,
						Protocol:   corev1.ProtocolTCP,
						Port:       servicePort,
						TargetPort: intstr.FromInt32(containerSSHPort),
					},
				},
			},
		}, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("create testkube service: %w", err)
	}

	// Create the pod directly.
	k8sPod, err := c.k8sClient.
		CoreV1().
		Pods(c.namespace).
		Create(ctx, &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        string(cloudID),
				Namespace:   c.namespace,
				Labels:      objectLabels(string(cloudID), location),
				Annotations: annotations,
			},
			Spec: corev1.PodSpec{
				RestartPolicy:                 corev1.RestartPolicyNever,
				TerminationGracePeriodSeconds: int64Ptr(1),
				Containers: []corev1.Container{
					{
						Name:            "vm",
						Image:           instanceTypeSpec.image,
						ImagePullPolicy: corev1.PullIfNotPresent,
						SecurityContext: &corev1.SecurityContext{
							Privileged: boolPtr(true),
						},
						Ports: []corev1.ContainerPort{
							{
								Name:          servicePortName,
								ContainerPort: containerSSHPort,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						ReadinessProbe: &corev1.Probe{
							ProbeHandler: corev1.ProbeHandler{
								TCPSocket: &corev1.TCPSocketAction{
									Port: intstr.FromInt32(containerSSHPort),
								},
							},
							InitialDelaySeconds: 1,
							PeriodSeconds:       2,
							FailureThreshold:    30,
						},
						Env: c.containerEnv(attrs),
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("250m"),
								corev1.ResourceMemory: resource.MustParse("512Mi"),
							},
							Limits: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("2"),
								corev1.ResourceMemory: resource.MustParse("4Gi"),
							},
						},
					},
				},
			},
		}, metav1.CreateOptions{})
	if err != nil {
		_ = c.k8sClient.CoreV1().Services(c.namespace).Delete(ctx, string(cloudID), metav1.DeleteOptions{})
		return nil, fmt.Errorf("create testkube pod: %w", err)
	}

	// Map to the brev instance.
	return c.instanceFromResources(k8sPod, k8sService), nil
}

func (c *TestKubeClient) GetInstance(ctx context.Context, instanceID cloudv1.CloudProviderInstanceID) (*cloudv1.Instance, error) {
	pod, err := c.k8sClient.
		CoreV1().
		Pods(c.namespace).
		Get(ctx, string(instanceID), metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("%w: %s", cloudv1.ErrInstanceNotFound, instanceID)
		}
		return nil, fmt.Errorf("get testkube pod: %w", err)
	}

	service, err := c.k8sClient.
		CoreV1().
		Services(c.namespace).
		Get(ctx, string(instanceID), metav1.GetOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, fmt.Errorf("get testkube service: %w", err)
	}
	if apierrors.IsNotFound(err) {
		service = nil
	}

	return c.instanceFromResources(pod, service), nil
}

func (c *TestKubeClient) ListInstances(ctx context.Context, args cloudv1.ListInstancesArgs) ([]cloudv1.Instance, error) {
	pods, err := c.k8sClient.
		CoreV1().
		Pods(c.namespace).
		List(ctx, metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(labels.Set{
				labelManagedBy: labelManagedByValue,
				labelName:      labelNameValue,
			}).String(),
		})
	if err != nil {
		return nil, fmt.Errorf("list testkube pods: %w", err)
	}

	instances := make([]cloudv1.Instance, 0, len(pods.Items))
	for _, pod := range pods.Items {
		instance, err := c.GetInstance(ctx, cloudv1.CloudProviderInstanceID(pod.Name))
		if err != nil {
			return nil, err
		}
		if !matchesListArgs(*instance, args) {
			continue
		}
		instances = append(instances, *instance)
	}
	return instances, nil
}

func (c *TestKubeClient) TerminateInstance(ctx context.Context, instanceID cloudv1.CloudProviderInstanceID) error {
	found := false
	if err := c.k8sClient.CoreV1().Pods(c.namespace).Delete(ctx, string(instanceID), metav1.DeleteOptions{}); err != nil {
		if !apierrors.IsNotFound(err) {
			return fmt.Errorf("delete testkube pod: %w", err)
		}
	} else {
		found = true
	}
	if err := c.k8sClient.CoreV1().Services(c.namespace).Delete(ctx, string(instanceID), metav1.DeleteOptions{}); err != nil {
		if !apierrors.IsNotFound(err) {
			return fmt.Errorf("delete testkube service: %w", err)
		}
	} else {
		found = true
	}
	if !found {
		return fmt.Errorf("%w: %s", cloudv1.ErrInstanceNotFound, instanceID)
	}
	return nil
}

func (c *TestKubeClient) UpdateInstanceTags(ctx context.Context, args cloudv1.UpdateInstanceTagsArgs) error {
	tagsJSON, err := marshalTags(args.Tags)
	if err != nil {
		return err
	}

	pod, err := c.k8sClient.CoreV1().Pods(c.namespace).Get(ctx, string(args.InstanceID), metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("%w: %s", cloudv1.ErrInstanceNotFound, args.InstanceID)
		}
		return fmt.Errorf("get testkube pod for tag update: %w", err)
	}
	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}
	pod.Annotations[annotationTagsJSON] = tagsJSON
	if _, err := c.k8sClient.CoreV1().Pods(c.namespace).Update(ctx, pod, metav1.UpdateOptions{}); err != nil {
		return fmt.Errorf("update testkube pod tags: %w", err)
	}

	service, err := c.k8sClient.CoreV1().Services(c.namespace).Get(ctx, string(args.InstanceID), metav1.GetOptions{})
	if err == nil {
		if service.Annotations == nil {
			service.Annotations = map[string]string{}
		}
		service.Annotations[annotationTagsJSON] = tagsJSON
		_, err = c.k8sClient.CoreV1().Services(c.namespace).Update(ctx, service, metav1.UpdateOptions{})
	}
	if err != nil && !apierrors.IsNotFound(err) {
		return fmt.Errorf("update testkube service tags: %w", err)
	}
	return nil
}

func (c *TestKubeClient) MergeInstanceForUpdate(_ cloudv1.Instance, newInst cloudv1.Instance) cloudv1.Instance {
	return newInst
}

func (c *TestKubeClient) MergeInstanceTypeForUpdate(_ cloudv1.InstanceType, newIt cloudv1.InstanceType) cloudv1.InstanceType {
	return newIt
}

func (c *TestKubeClient) resourceAnnotations(cloudID cloudv1.CloudProviderInstanceID, attrs cloudv1.CreateInstanceAttrs, spec instanceTypeSpec) map[string]string {
	name := attrs.Name
	if name == "" {
		name = string(cloudID)
	}
	tagsJSON, _ := marshalTags(attrs.Tags)
	return map[string]string{
		annotationRefID:          attrs.RefID,
		annotationCloudCredRefID: c.refID,
		annotationName:           name,
		annotationLocation:       c.resourceLocation(attrs),
		annotationSubLocation:    attrs.SubLocation,
		annotationInstanceType:   attrs.InstanceType,
		annotationImageID:        spec.imageID,
		annotationCreatedAt:      time.Now().UTC().Format(time.RFC3339Nano),
		annotationScenario:       scenarioForInstanceType(attrs.InstanceType),
		annotationTagsJSON:       tagsJSON,
	}
}

func (c *TestKubeClient) resourceLocation(attrs cloudv1.CreateInstanceAttrs) string {
	if attrs.Location != "" {
		return attrs.Location
	}
	return c.location
}

func (c *TestKubeClient) containerEnv(attrs cloudv1.CreateInstanceAttrs) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{Name: "USER_NAME", Value: "ubuntu"},
		{Name: "PUBLIC_KEY", Value: attrs.PublicKey},
		{Name: "PASSWORD_ACCESS", Value: "false"},
		{Name: "SUDO_ACCESS", Value: "true"},
		{Name: envInstanceType, Value: attrs.InstanceType},
		{Name: envScenario, Value: scenarioForInstanceType(attrs.InstanceType)},
	}
	if attrs.InstanceType == InstanceTypeFailBuild {
		env = append(env, corev1.EnvVar{Name: envFailBuild, Value: "true"})
	}
	return env
}

func (c *TestKubeClient) instanceFromResources(pod *corev1.Pod, service *corev1.Service) *cloudv1.Instance {
	annotations := pod.Annotations
	instanceType := annotations[annotationInstanceType]
	location := firstNonEmpty(annotations[annotationLocation], c.location)
	instance := &cloudv1.Instance{
		Name:           firstNonEmpty(annotations[annotationName], pod.Name),
		RefID:          annotations[annotationRefID],
		CloudCredRefID: annotations[annotationCloudCredRefID],
		CreatedAt:      createdAt(pod),
		CloudID:        cloudv1.CloudProviderInstanceID(pod.Name),
		Hostname:       pod.Name,
		ImageID:        annotations[annotationImageID],
		InstanceType:   instanceType,
		DiskSize:       units.GiB * 20,
		DiskSizeBytes:  cloudv1.NewBytes(20, cloudv1.Gibibyte),
		VolumeType:     "ephemeral",
		SSHUser:        "ubuntu",
		SSHPort:        int(servicePort),
		Status:         statusFromResources(pod, service),
		Tags:           tagsFromAnnotations(annotations),
		Stoppable:      false,
		Rebootable:     false,
		IsContainer:    false,
		Location:       location,
		SubLocation:    annotations[annotationSubLocation],
		FirewallRules:  sshFirewallRules(),
	}
	instance.InstanceTypeID = cloudv1.MakeGenericInstanceTypeIDFromInstance(*instance)
	if service != nil {
		populateNetwork(service, instance)
	}
	return instance
}

func populateNetwork(service *corev1.Service, instance *cloudv1.Instance) {
	// Default the private IP to the cluster IP.
	if service.Spec.ClusterIP != "" && service.Spec.ClusterIP != corev1.ClusterIPNone {
		instance.PrivateIP = service.Spec.ClusterIP
	}
	// Default the SSH port to the first open port.
	if len(service.Spec.Ports) > 0 {
		instance.SSHPort = int(service.Spec.Ports[0].Port)
	}

	switch service.Spec.Type {
	case corev1.ServiceTypeLoadBalancer:
		// No ingress means no public IP.
		if len(service.Status.LoadBalancer.Ingress) == 0 {
			return
		}
		// Set the public IP to the first ingress IP.
		ingress := service.Status.LoadBalancer.Ingress[0]
		if ingress.IP != "" {
			instance.PublicIP = ingress.IP
			instance.PublicDNS = ingress.IP
		}
		if ingress.Hostname != "" {
			instance.PublicDNS = ingress.Hostname
			if instance.PublicIP == "" {
				instance.PublicIP = ingress.Hostname
			}
		}
	case corev1.ServiceTypeNodePort:
		// No ports means no node port.
		if len(service.Spec.Ports) == 0 {
			return
		}
		// Set the SSH port to the first node port.
		instance.SSHPort = int(service.Spec.Ports[0].NodePort)
	case corev1.ServiceTypeClusterIP:
		// Keep the cluster IP as the private IP.
	}
}

func statusFromResources(pod *corev1.Pod, service *corev1.Service) cloudv1.Status {
	if pod.DeletionTimestamp != nil {
		return cloudv1.Status{LifecycleStatus: cloudv1.LifecycleStatusTerminating}
	}
	if podFailed(*pod) {
		return cloudv1.Status{
			LifecycleStatus: cloudv1.LifecycleStatusFailed,
			Messages:        podMessages(*pod),
		}
	}
	if podReady(*pod) {
		if service == nil {
			return cloudv1.Status{
				LifecycleStatus: cloudv1.LifecycleStatusPending,
				Messages:        append(podMessages(*pod), "waiting for service"),
			}
		}
		if service.Spec.Type == corev1.ServiceTypeLoadBalancer && !loadBalancerReady(service) {
			return cloudv1.Status{
				LifecycleStatus: cloudv1.LifecycleStatusPending,
				Messages:        append(podMessages(*pod), fmt.Sprintf("service %s waiting for load balancer ingress", service.Name)),
			}
		}
		return cloudv1.Status{
			LifecycleStatus: cloudv1.LifecycleStatusRunning,
			Messages:        podMessages(*pod),
		}
	}
	return cloudv1.Status{
		LifecycleStatus: cloudv1.LifecycleStatusPending,
		Messages:        podMessages(*pod),
	}
}

func loadBalancerReady(service *corev1.Service) bool {
	for _, ingress := range service.Status.LoadBalancer.Ingress {
		if ingress.IP != "" || ingress.Hostname != "" {
			return true
		}
	}
	return false
}

func podReady(pod corev1.Pod) bool {
	if pod.Status.Phase != corev1.PodRunning {
		return false
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

func podFailed(pod corev1.Pod) bool {
	if pod.Status.Phase == corev1.PodFailed {
		return true
	}
	for _, status := range pod.Status.ContainerStatuses {
		if status.State.Terminated != nil && status.State.Terminated.ExitCode != 0 {
			return true
		}
		if status.State.Waiting != nil && isFailureWaitingReason(status.State.Waiting.Reason) {
			return true
		}
	}
	return false
}

func isFailureWaitingReason(reason string) bool {
	switch reason {
	case "CrashLoopBackOff", "CreateContainerConfigError", "ErrImagePull", "ImagePullBackOff", "InvalidImageName":
		return true
	default:
		return false
	}
}

func podMessages(pod corev1.Pod) []string {
	messages := []string{}
	if pod.Status.Phase != "" {
		messages = append(messages, fmt.Sprintf("%s: phase=%s", pod.Name, pod.Status.Phase))
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Message != "" {
			messages = append(messages, fmt.Sprintf("%s: %s", pod.Name, condition.Message))
		}
	}
	for _, status := range pod.Status.ContainerStatuses {
		if status.State.Waiting != nil {
			message := status.State.Waiting.Reason
			if status.State.Waiting.Message != "" {
				message += ": " + status.State.Waiting.Message
			}
			messages = append(messages, fmt.Sprintf("%s/%s waiting: %s", pod.Name, status.Name, message))
		}
		if status.State.Terminated != nil {
			message := status.State.Terminated.Reason
			if status.State.Terminated.Message != "" {
				message += ": " + status.State.Terminated.Message
			}
			messages = append(messages, fmt.Sprintf("%s/%s terminated: %s", pod.Name, status.Name, message))
		}
	}
	return messages
}

func createdAt(pod *corev1.Pod) time.Time {
	if pod.Annotations != nil {
		if createdAtRaw := pod.Annotations[annotationCreatedAt]; createdAtRaw != "" {
			if parsed, err := time.Parse(time.RFC3339Nano, createdAtRaw); err == nil {
				return parsed
			}
		}
	}
	return pod.CreationTimestamp.Time
}

func matchesListArgs(instance cloudv1.Instance, args cloudv1.ListInstancesArgs) bool {
	if len(args.InstanceIDs) > 0 && !containsInstanceID(args.InstanceIDs, instance.CloudID) {
		return false
	}
	if len(args.Locations) > 0 && !args.Locations.IsAll() && !args.Locations.IsAllowed(instance.Location) {
		return false
	}
	for tagKey, allowedValues := range args.TagFilters {
		tagValue, ok := instance.Tags[tagKey]
		if !ok {
			return false
		}
		if len(allowedValues) > 0 && !slices.Contains(allowedValues, tagValue) {
			return false
		}
	}
	return true
}

func containsInstanceID(values []cloudv1.CloudProviderInstanceID, value cloudv1.CloudProviderInstanceID) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func selectorLabels(cloudID string) map[string]string {
	return map[string]string{
		labelCloudID: cloudID,
	}
}

func objectLabels(cloudID string, location string) map[string]string {
	labels := selectorLabels(cloudID)
	labels[labelName] = labelNameValue
	labels[labelManagedBy] = labelManagedByValue
	labels[labelLocation] = sanitizeLabelValue(location)
	return labels
}

func makeCloudID(credentialRefID string, refID string) cloudv1.CloudProviderInstanceID {
	sum := sha256.Sum256([]byte(credentialRefID + ":" + refID))
	return cloudv1.CloudProviderInstanceID("tk-" + hex.EncodeToString(sum[:])[:20])
}

func scenarioForInstanceType(instanceType string) string {
	return strings.TrimPrefix(instanceType, "test.")
}

func marshalTags(tags cloudv1.Tags) (string, error) {
	if tags == nil {
		tags = cloudv1.Tags{}
	}
	tagsBytes, err := json.Marshal(tags)
	if err != nil {
		return "", fmt.Errorf("marshal testkube tags: %w", err)
	}
	return string(tagsBytes), nil
}

func tagsFromAnnotations(annotations map[string]string) cloudv1.Tags {
	tags := cloudv1.Tags{}
	if annotations == nil || annotations[annotationTagsJSON] == "" {
		return tags
	}
	if err := json.Unmarshal([]byte(annotations[annotationTagsJSON]), &tags); err != nil {
		return cloudv1.Tags{}
	}
	return tags
}

var invalidLabelValueCharPattern = regexp.MustCompile(`[^a-z0-9_.-]`)

const maxLabelValueLength = 63

func sanitizeLabelValue(value string) string {
	sanitized := invalidLabelValueCharPattern.ReplaceAllString(strings.ToLower(value), "-")
	sanitized = strings.Trim(sanitized, "-_.")
	if len(sanitized) > maxLabelValueLength {
		sanitized = sanitized[:maxLabelValueLength]
		sanitized = strings.TrimRight(sanitized, "-_.")
	}
	if sanitized == "" {
		return "unknown"
	}
	return sanitized
}

func sshFirewallRules() cloudv1.FirewallRules {
	rule := cloudv1.FirewallRule{
		FromPort: servicePort,
		ToPort:   servicePort,
		IPRanges: []string{"0.0.0.0/0"},
	}
	return cloudv1.FirewallRules{
		IngressRules: []cloudv1.FirewallRule{rule},
		EgressRules:  []cloudv1.FirewallRule{rule},
	}
}

func int64Ptr(value int64) *int64 {
	return &value
}

func boolPtr(value bool) *bool {
	return &value
}
