package hyperstack

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	defaultImageName        = "Ubuntu Server 22.04 LTS (Jammy Jellyfish)"
	defaultSSHPort          = 22
	defaultSSHUser          = "ubuntu"
	defaultPageSize         = 100
	refIDLabelPrefix        = "brev-ref-"
	tagLabelPrefix          = "brev-tag-"
	tagLabelSeparator       = "_"
	managedKeyIDLabelPrefix = "brev-managed-key-id-"
)

var resourceNameInvalidCharacters = regexp.MustCompile(`[^a-zA-Z0-9-]+`)

// Hyperstack permits at most 10 labels. These seven caller tags plus the
// canonical ref ID and optional managed-key ID use at most nine labels.
var instanceTagLabelKeys = []string{
	"dev-plane-managedBy",
	"dev-plane-x-instanceId",
	"dev-plane-x-environmentId",
	"dev-plane-x-userId",
	"dev-plane-x-launchableId",
	"dev-plane-x-cloudCredId",
	"dev-plane-stage",
}

func (c *HyperstackClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	location := strings.TrimSpace(attrs.Location)
	if location == "" {
		location = strings.TrimSpace(c.location)
	}
	if err := validateCreateInstanceAttrs(attrs, location); err != nil {
		return nil, err
	}

	providerEnvironment, err := c.getDefaultEnvironment(ctx, location)
	if err != nil {
		return nil, err
	}
	environmentName := stringValue(providerEnvironment.Name)
	keyPair, err := c.resolveKeyPair(ctx, attrs, environmentName)
	if err != nil {
		return nil, err
	}
	imageName := strings.TrimSpace(attrs.ImageID)
	if imageName == "" {
		imageName = defaultImageName
	}
	securityRules, err := makeSecurityRules(attrs.FirewallRules)
	if err != nil {
		return nil, err
	}
	labels := makeLabels(attrs.RefID, attrs.Tags)
	if keyPair.managedID != 0 {
		labels = append(labels, managedKeyIDLabelPrefix+strconv.Itoa(keyPair.managedID))
	}
	assignFloatingIP := true
	enablePortRandomization := false
	enhancedMonitoringEnabled := false
	userData := readinessCloudConfig

	// Hyperstack root disk sizes are fixed by flavor. Intentionally do not map
	// attrs.DiskSize or attrs.DiskSizeBytes into the provider request.
	response, err := c.virtualMachines.CreateVMsWithResponse(ctx, virtualmachine.CreateInstancesPayload{
		Name:                      managedResourceName(attrs.RefID),
		EnvironmentName:           environmentName,
		KeyName:                   keyPair.name,
		ImageName:                 &imageName,
		FlavorName:                attrs.InstanceType,
		Count:                     1,
		AssignFloatingIp:          &assignFloatingIP,
		EnablePortRandomization:   &enablePortRandomization,
		EnhancedMonitoringEnabled: &enhancedMonitoringEnabled,
		Labels:                    &labels,
		SecurityRules:             &securityRules,
		UserData:                  &userData,
	})
	if err != nil {
		return nil, wrapTransportError("create virtual machine", err)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, responseError("create virtual machine", response.StatusCode(), response.Body, nil)
	}
	if response.JSON200 == nil || response.JSON200.Instances == nil || len(*response.JSON200.Instances) != 1 {
		return nil, errors.New("hyperstack create virtual machine response did not contain exactly one instance")
	}
	providerInstance := (*response.JSON200.Instances)[0]
	if providerInstance.Id == nil || *providerInstance.Id <= 0 {
		return nil, errors.New("hyperstack create virtual machine response did not contain an instance ID")
	}

	instanceID := v1.CloudProviderInstanceID(strconv.Itoa(*providerInstance.Id))
	instance, err := c.GetInstance(ctx, instanceID)
	if err != nil {
		return nil, errors.Join(err, c.TerminateInstance(ctx, instanceID))
	}
	return instance, nil
}

func validateCreateInstanceAttrs(attrs v1.CreateInstanceAttrs, location string) error {
	switch {
	case strings.TrimSpace(attrs.RefID) == "":
		return errors.New("hyperstack instance RefID is required")
	case strings.TrimSpace(attrs.InstanceType) == "":
		return errors.New("hyperstack instance type is required")
	case location == "":
		return errors.New("hyperstack instance location is required")
	case strings.TrimSpace(attrs.PublicKey) == "" && (attrs.KeyPairName == nil || strings.TrimSpace(*attrs.KeyPairName) == ""):
		return errors.New("hyperstack instance public key or key pair name is required")
	case attrs.UserDataBase64 != "":
		return errors.New("hyperstack provider does not support instance user data")
	case len(attrs.AdditionalDisks) > 0:
		return errors.New("hyperstack provider does not support additional disks")
	case attrs.UseSpot != isSpotFlavor(attrs.InstanceType, attrs.InstanceType):
		return errors.New("hyperstack spot selection must match a -spot instance type")
	default:
		return nil
	}
}

func (c *HyperstackClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	numericID, err := parseInstanceID(instanceID)
	if err != nil {
		return nil, err
	}
	providerInstance, err := c.getProviderInstance(ctx, numericID)
	if err != nil {
		return nil, err
	}
	instance, err := c.convertProviderInstance(ctx, providerInstance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func (c *HyperstackClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	instances := make([]v1.Instance, 0)
	for page := 1; ; page++ {
		pageSize := defaultPageSize
		response, err := c.virtualMachines.ListVMsWithResponse(ctx, &virtualmachine.ListVMsParams{
			Page:     &page,
			PageSize: &pageSize,
		})
		if err != nil {
			return nil, wrapTransportError("list virtual machines", err)
		}
		if response.StatusCode() != http.StatusOK {
			return nil, responseError("list virtual machines", response.StatusCode(), response.Body, nil)
		}
		if response.JSON200 == nil || response.JSON200.Instances == nil {
			return nil, errors.New("hyperstack list virtual machines response did not contain data")
		}

		providerInstances := *response.JSON200.Instances
		for _, providerInstance := range providerInstances {
			instance, err := c.convertProviderInstance(ctx, providerInstance)
			if err != nil {
				return nil, err
			}
			if matchesListArgs(instance, args) {
				instances = append(instances, instance)
			}
		}
		if len(providerInstances) < pageSize {
			break
		}
	}
	return instances, nil
}

func (c *HyperstackClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	numericID, err := parseInstanceID(instanceID)
	if err != nil {
		return err
	}
	providerInstance, err := c.getProviderInstance(ctx, numericID)
	if errors.Is(err, v1.ErrInstanceNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	managedKeyPairID, err := managedKeyPairID(providerInstance.Labels)
	if err != nil {
		return err
	}
	response, err := c.virtualMachines.DeleteVMWithResponse(ctx, numericID)
	if err != nil {
		return wrapTransportError("delete virtual machine", err)
	}
	if response.StatusCode() != http.StatusOK && response.StatusCode() != http.StatusNotFound {
		return responseError("delete virtual machine", response.StatusCode(), response.Body, nil)
	}
	if managedKeyPairID == 0 {
		return nil
	}
	return c.deleteManagedKeyPair(ctx, managedKeyPairID)
}

func (c *HyperstackClient) StopInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	numericID, err := parseInstanceID(instanceID)
	if err != nil {
		return err
	}
	response, err := c.virtualMachines.StopVMWithResponse(ctx, numericID)
	if err != nil {
		return wrapTransportError("stop virtual machine", err)
	}
	if response.StatusCode() != http.StatusOK {
		return responseError("stop virtual machine", response.StatusCode(), response.Body, v1.ErrInstanceNotFound)
	}
	return nil
}

func (c *HyperstackClient) StartInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	numericID, err := parseInstanceID(instanceID)
	if err != nil {
		return err
	}
	providerInstance, err := c.getProviderInstance(ctx, numericID)
	if err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(stringValue(providerInstance.Status))) {
	case "active", "running", "starting", "powering-on":
		return nil
	}
	response, err := c.virtualMachines.StartVMWithResponse(ctx, numericID)
	if err != nil {
		return wrapTransportError("start virtual machine", err)
	}
	if response.StatusCode() != http.StatusOK {
		return responseError("start virtual machine", response.StatusCode(), response.Body, v1.ErrInstanceNotFound)
	}
	return nil
}

func (c *HyperstackClient) getProviderInstance(ctx context.Context, instanceID int) (virtualmachine.InstanceFields, error) {
	response, err := c.virtualMachines.GetVMWithResponse(ctx, instanceID)
	if err != nil {
		return virtualmachine.InstanceFields{}, wrapTransportError("get virtual machine", err)
	}
	if response.StatusCode() != http.StatusOK {
		return virtualmachine.InstanceFields{}, responseError("get virtual machine", response.StatusCode(), response.Body, v1.ErrInstanceNotFound)
	}
	if response.JSON200 == nil || response.JSON200.Instance == nil {
		return virtualmachine.InstanceFields{}, errors.New("hyperstack get virtual machine response did not contain data")
	}
	return *response.JSON200.Instance, nil
}

func parseInstanceID(instanceID v1.CloudProviderInstanceID) (int, error) {
	numericID, err := strconv.Atoi(string(instanceID))
	if err != nil || numericID <= 0 {
		return 0, fmt.Errorf("invalid hyperstack instance ID %q", instanceID)
	}
	return numericID, nil
}

func (c *HyperstackClient) convertProviderInstance(ctx context.Context, providerInstance virtualmachine.InstanceFields) (v1.Instance, error) {
	consoleReady := false
	if hyperstackLifecycleStatus(stringValue(providerInstance.Status)) == v1.LifecycleStatusRunning &&
		hyperstackAPIReady(providerInstance, strings.TrimSpace(stringValue(providerInstance.FloatingIp))) {
		var err error
		consoleReady, err = c.vmOperatingSystemReportsReady(ctx, providerInstance)
		if err != nil {
			return v1.Instance{}, err
		}
	}
	return c.convertInstance(providerInstance, consoleReady), nil
}

func (c *HyperstackClient) convertInstance(
	providerInstance virtualmachine.InstanceFields,
	consoleReady bool,
) v1.Instance {
	cloudID := strconv.Itoa(intValue(providerInstance.Id))
	name := strings.TrimSpace(stringValue(providerInstance.Name))
	refID, tags := parseLabels(providerInstance.Labels)
	if refID == "" {
		refID = name
	}

	location := ""
	if providerInstance.Environment != nil {
		location = stringValue(providerInstance.Environment.Region)
	}
	instanceType := ""
	diskSize, diskSizeBytes := byteSizes(0, v1.Gigabyte)
	if providerInstance.Flavor != nil {
		instanceType = stringValue(providerInstance.Flavor.Name)
		diskSize, diskSizeBytes = byteSizes(int64(intValue(providerInstance.Flavor.Disk)), v1.Gigabyte)
	}
	imageName := ""
	if providerInstance.Image != nil {
		imageName = stringValue(providerInstance.Image.Name)
	}
	publicIP := strings.TrimSpace(stringValue(providerInstance.FloatingIp))
	lifecycleStatus := hyperstackLifecycleStatus(stringValue(providerInstance.Status))
	if lifecycleStatus == v1.LifecycleStatusRunning && (!hyperstackAPIReady(providerInstance, publicIP) || !consoleReady) {
		lifecycleStatus = v1.LifecycleStatusPending
	}

	instance := v1.Instance{
		Name:           name,
		RefID:          refID,
		CloudCredRefID: c.refID,
		CloudID:        v1.CloudProviderInstanceID(cloudID),
		PublicIP:       publicIP,
		PublicDNS:      publicIP,
		PrivateIP:      stringValue(providerInstance.FixedIp),
		Hostname:       name,
		ImageID:        imageName,
		InstanceType:   instanceType,
		DiskSize:       diskSize,
		DiskSizeBytes:  diskSizeBytes,
		VolumeType:     "ssd",
		SSHUser:        sshUser(imageName),
		SSHPort:        defaultSSHPort,
		Status: v1.Status{
			LifecycleStatus: lifecycleStatus,
		},
		FirewallRules: providerFirewallRules(providerInstance.SecurityRules),
		Location:      location,
		Tags:          tags,
		Spot:          isSpotFlavor(instanceType, instanceType),
		Stoppable:     true,
	}
	if providerInstance.CreatedAt != nil {
		instance.CreatedAt = providerInstance.CreatedAt.Time
	}
	instance.InstanceTypeID = v1.MakeGenericInstanceTypeID(v1.InstanceType{
		Type:     instance.InstanceType,
		Location: instance.Location,
	})
	return instance
}

func hyperstackAPIReady(providerInstance virtualmachine.InstanceFields, publicIP string) bool {
	if publicIP == "" {
		return false
	}
	return readinessFieldComplete(providerInstance.FloatingIpStatus, "active", "attached") &&
		readinessFieldComplete(providerInstance.VmState, "active", "running") &&
		readinessFieldComplete(providerInstance.PowerState, "active", "running", "on")
}

func readinessFieldComplete(value *string, readyValues ...string) bool {
	status := strings.ToLower(strings.TrimSpace(stringValue(value)))
	return status == "" || slices.Contains(readyValues, status)
}

func hyperstackLifecycleStatus(status string) v1.LifecycleStatus {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "creating", "build", "rebuilding", "initializing":
		return v1.LifecycleStatusPending
	case "active", "running":
		return v1.LifecycleStatusRunning
	case "stopping", "powering-off":
		return v1.LifecycleStatusStopping
	case "stopped", "shutoff", "powered-off":
		return v1.LifecycleStatusStopped
	case "hibernating", "suspending":
		return v1.LifecycleStatusSuspending
	case "hibernated", "suspended":
		return v1.LifecycleStatusSuspended
	case "deleting", "terminating":
		return v1.LifecycleStatusTerminating
	case "deleted", "terminated":
		return v1.LifecycleStatusTerminated
	case "error", "failed":
		return v1.LifecycleStatusFailed
	default:
		return v1.LifecycleStatusPending
	}
}

func providerFirewallRules(providerRules *[]virtualmachine.SecurityRulesFieldsForInstance) v1.FirewallRules {
	if providerRules == nil {
		return v1.FirewallRules{}
	}
	ingressRules := make([]v1.FirewallRule, 0)
	for _, providerRule := range *providerRules {
		if !strings.EqualFold(stringValue(providerRule.Direction), "ingress") {
			continue
		}
		ingressRules = append(ingressRules, v1.FirewallRule{
			ID:       strconv.Itoa(intValue(providerRule.Id)),
			FromPort: int32(intValue(providerRule.PortRangeMin)),
			ToPort:   int32(intValue(providerRule.PortRangeMax)),
			IPRanges: []string{stringValue(providerRule.RemoteIpPrefix)},
		})
	}
	return v1.FirewallRules{IngressRules: ingressRules}
}

func matchesListArgs(instance v1.Instance, args v1.ListInstancesArgs) bool {
	if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, instance.CloudID) {
		return false
	}
	if len(args.Locations) > 0 && !args.Locations.IsAllowed(instance.Location) {
		return false
	}
	for key, values := range args.TagFilters {
		value, found := instance.Tags[key]
		if !found || len(values) > 0 && !slices.Contains(values, value) {
			return false
		}
	}
	return true
}

func makeLabels(refID string, tags v1.Tags) []string {
	labels := []string{refIDLabelPrefix + refID}
	for _, key := range instanceTagLabelKeys {
		if value, ok := tags[key]; ok {
			labels = append(labels, providerTagLabelPrefix(key)+value)
		}
	}
	return labels
}

func parseLabels(providerLabels *[]string) (string, v1.Tags) {
	tags := make(v1.Tags)
	if providerLabels == nil {
		return "", tags
	}
	var refID string
	for _, label := range *providerLabels {
		if key, value, ok := parseTagLabel(label); ok {
			tags[key] = value
			continue
		}
		switch {
		case strings.HasPrefix(label, refIDLabelPrefix):
			refID = strings.TrimPrefix(label, refIDLabelPrefix)
		case strings.HasPrefix(label, managedKeyIDLabelPrefix):
			continue
		default:
			tags[label] = ""
		}
	}
	return refID, tags
}

func providerTagLabelPrefix(key string) string {
	return tagLabelPrefix + strings.ToLower(key) + tagLabelSeparator
}

func parseTagLabel(label string) (string, string, bool) {
	for _, key := range instanceTagLabelKeys {
		prefix := providerTagLabelPrefix(key)
		if strings.HasPrefix(label, prefix) {
			return key, strings.TrimPrefix(label, prefix), true
		}
	}
	return "", "", false
}

func managedKeyPairID(providerLabels *[]string) (int, error) {
	if providerLabels == nil {
		return 0, nil
	}
	for _, label := range *providerLabels {
		if !strings.HasPrefix(label, managedKeyIDLabelPrefix) {
			continue
		}
		keyPairID, err := strconv.Atoi(strings.TrimPrefix(label, managedKeyIDLabelPrefix))
		if err != nil || keyPairID <= 0 {
			return 0, fmt.Errorf("invalid hyperstack managed keypair label %q", label)
		}
		return keyPairID, nil
	}
	return 0, nil
}

func managedResourceName(name string) string {
	name = strings.TrimSpace(name)
	name = resourceNameInvalidCharacters.ReplaceAllString(name, "-")
	name = strings.Trim(name, "-")
	if len(name) > 63 {
		name = strings.TrimRight(name[:63], "-")
	}
	if name == "" {
		return "brev-instance"
	}
	return name
}

func sshUser(imageName string) string {
	lowerName := strings.ToLower(imageName)
	switch {
	case strings.Contains(lowerName, "debian"):
		return "debian"
	case strings.Contains(lowerName, "alma"):
		return "almalinux"
	default:
		return defaultSSHUser
	}
}

func (c *HyperstackClient) GetInstancePollTime() time.Duration {
	return 10 * time.Second
}
