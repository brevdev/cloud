package massedcompute

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"
	"golang.org/x/crypto/ssh"
)

const (
	defaultSSHPort      = 22
	sshKeyNamePrefix    = "brevkey"
	instanceNameDivider = "_"
	defaultImageName    = "Ubuntu Server 22.04 w/ drivers"
)

var resourceNameInvalidCharacters = regexp.MustCompile(`[^a-zA-Z0-9_-]+`)

func (c *MassedComputeClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	location := attrs.Location
	if location == "" {
		location = c.location
	}
	if err := validateCreateInstanceAttrs(attrs, location); err != nil {
		return nil, err
	}

	startupCommand, err := buildStartupCommand(attrs.FirewallRules)
	if err != nil {
		return nil, err
	}

	imageID, err := c.resolveImageID(ctx, attrs.ImageID)
	if err != nil {
		return nil, err
	}

	publicKey, err := normalizeSSHPublicKey(attrs.PublicKey)
	if err != nil {
		return nil, err
	}

	keyName, err := c.ensureSSHKey(ctx, publicKey, attrs.RefID)
	if err != nil {
		return nil, err
	}

	providerName := makeProviderInstanceName(attrs.RefID, attrs.Name)
	instanceID, err := c.launchInstance(ctx, openapi.InstanceLaunchPostRequest{
		ImageId:      imageID,
		ProductName:  attrs.InstanceType,
		RegionName:   location,
		InstanceName: &providerName,
		Command:      &startupCommand,
		SshKeys:      []string{keyName},
	})
	if err != nil {
		return nil, err
	}

	instance, err := c.GetInstance(ctx, v1.CloudProviderInstanceID(instanceID))
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func validateCreateInstanceAttrs(attrs v1.CreateInstanceAttrs, location string) error {
	switch {
	case attrs.RefID == "":
		return errors.New("massed compute instance RefID is required")
	case attrs.InstanceType == "":
		return errors.New("massed compute instance type is required")
	case location == "":
		return errors.New("massed compute instance location is required")
	case strings.TrimSpace(attrs.PublicKey) == "":
		return errors.New("massed compute instance public key is required")
	case attrs.UserDataBase64 != "":
		return errors.New("massed compute does not support instance user data")
	case attrs.UseSpot || strings.HasSuffix(strings.ToLower(strings.TrimSpace(attrs.InstanceType)), "_spot"):
		return errors.New("massed compute spot instances are not supported")
	default:
		return nil
	}
}

func normalizeSSHPublicKey(publicKey string) (string, error) {
	publicKey = strings.TrimSpace(publicKey)
	if key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey)); err == nil {
		return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("massed compute public key must be OpenSSH or PEM encoded")
	}
	parsedKey, pkixErr := x509.ParsePKIXPublicKey(block.Bytes)
	if pkixErr != nil {
		rsaKey, pkcs1Err := x509.ParsePKCS1PublicKey(block.Bytes)
		if pkcs1Err != nil {
			return "", fmt.Errorf("parse massed compute PEM public key: %w", errors.Join(pkixErr, pkcs1Err))
		}
		parsedKey = rsaKey
	}
	key, err := ssh.NewPublicKey(parsedKey)
	if err != nil {
		return "", fmt.Errorf("convert massed compute public key to OpenSSH: %w", err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
}

func (c *MassedComputeClient) launchInstance(ctx context.Context, req openapi.InstanceLaunchPostRequest) (string, error) {
	resp, httpResp, err := c.client.InstancesAPI.InstanceLaunchPost(ctx).
		InstanceLaunchPostRequest(req).
		Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return "", wrapMassedComputeError(err, httpResp)
	}
	if resp == nil || resp.Response == nil || *resp.Response == "" {
		return "", errors.New("massed compute launch response did not contain an instance UUID")
	}
	return stringValue(resp.Response), nil
}

func (c *MassedComputeClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	resp, err := c.getRunningInstance(ctx, instanceID)
	if err != nil {
		return nil, err
	}
	return c.convertInstanceToV1Instance(ctx, resp[0])
}

func (c *MassedComputeClient) getRunningInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) ([]openapi.RetrieveAllRunningInstancesV1RunningInstancesInner, error) {
	resp, httpResp, err := c.client.InstancesAPI.InstanceUuidGet(ctx, string(instanceID)).Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("massed compute instance %s: %w", instanceID, v1.ErrInstanceNotFound)
		}
		return nil, wrapMassedComputeError(err, httpResp)
	}
	if resp == nil || len(resp.RunningInstances) == 0 {
		return nil, errors.New("massed compute get response did not contain instance data")
	}
	return resp.RunningInstances, nil
}

func (c *MassedComputeClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	resp, err := c.listRunningInstances(ctx)
	if err != nil {
		return nil, err
	}

	instances := make([]v1.Instance, 0, len(resp))
	for _, providerInstance := range resp {
		instance, err := c.convertInstanceToV1Instance(ctx, providerInstance)
		if err != nil {
			return nil, err
		}
		if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, instance.CloudID) {
			continue
		}
		if len(args.Locations) > 0 && !args.Locations.IsAllowed(instance.Location) {
			continue
		}
		instances = append(instances, *instance)
	}
	return instances, nil
}

func (c *MassedComputeClient) listRunningInstances(ctx context.Context) ([]openapi.RetrieveAllRunningInstancesV1RunningInstancesInner, error) {
	resp, httpResp, err := c.client.InstancesAPI.InstanceGet(ctx).Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return nil, wrapMassedComputeError(err, httpResp)
	}
	if resp == nil {
		return nil, errors.New("massed compute list response was empty")
	}
	return resp.RunningInstances, nil
}

func (c *MassedComputeClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	req := openapi.InstanceRestartPostRequest{
		InstanceUuids: []string{string(instanceID)},
	}
	err := c.terminateInstance(ctx, req)
	return err
}

func (c *MassedComputeClient) terminateInstance(ctx context.Context, req openapi.InstanceRestartPostRequest) error {
	_, httpResp, err := c.client.InstancesAPI.InstanceTerminatePost(ctx).
		InstanceRestartPostRequest(req).
		Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil
		}
		return wrapMassedComputeError(err, httpResp)
	}
	return nil
}

func (c *MassedComputeClient) resolveImageID(ctx context.Context, requestedImage string) (int32, error) {
	requestedImage = strings.TrimSpace(requestedImage)
	if requestedImage != "" {
		imageID, err := strconv.ParseInt(requestedImage, 10, 32)
		if err == nil && imageID > 0 {
			return int32(imageID), nil
		}
	}
	imageName := requestedImage
	if imageName == "" {
		imageName = defaultImageName
	}

	images, err := c.getImages(ctx)
	if err != nil {
		return 0, err
	}
	for _, image := range images {
		if strings.EqualFold(strings.TrimSpace(stringValue(image.VmImageName)), imageName) && int32Value(image.VmImageId) > 0 {
			return int32Value(image.VmImageId), nil
		}
	}
	return 0, fmt.Errorf("massed compute image %q was not found", imageName)
}

func (c *MassedComputeClient) resolveImageName(ctx context.Context, imageID int32) (string, error) {
	images, err := c.getImages(ctx)
	if err != nil {
		return "", err
	}
	for _, image := range images {
		if int32Value(image.VmImageId) == imageID && strings.TrimSpace(stringValue(image.VmImageName)) != "" {
			return strings.TrimSpace(stringValue(image.VmImageName)), nil
		}
	}
	return "", fmt.Errorf("massed compute image ID %d was not found", imageID)
}

func (c *MassedComputeClient) getImages(ctx context.Context) ([]openapi.ImagesV1ImagesInner, error) {
	resp, httpResp, err := c.client.DefaultAPI.ImagesGet(ctx).Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return nil, wrapMassedComputeError(err, httpResp)
	}
	if resp == nil {
		return nil, errors.New("massed compute image response was empty")
	}
	return resp.Images, nil
}

func (c *MassedComputeClient) ensureSSHKey(ctx context.Context, publicKey string, refID string) (string, error) {
	resp, err := c.listSSHKeys(ctx)
	if err != nil {
		return "", err
	}
	if resp == nil || resp.SshKeys == nil {
		return "", errors.New("massed compute SSH-key response did not contain data")
	}
	for _, key := range resp.SshKeys {
		if strings.TrimSpace(stringValue(key.PublicKey)) == publicKey {
			return stringValue(key.Name), nil
		}
	}

	// SSH keys cannot contain punctuation other than spaces
	keyRefID := strings.ReplaceAll(refID, "-", "")
	req := openapi.SshKeysPostRequest{
		Name:      managedResourceName(sshKeyNamePrefix, " ", keyRefID),
		PublicKey: publicKey,
	}
	createdKey, err := c.createSSHKey(ctx, req)
	if err != nil {
		return "", err
	}
	if createdKey == nil || createdKey.SshKey == nil || createdKey.SshKey.Name == nil {
		return "", errors.New("massed compute SSH-key response did not contain data")
	}
	return stringValue(createdKey.SshKey.Name), nil
}

func (c *MassedComputeClient) listSSHKeys(ctx context.Context) (*openapi.SSHKey, error) {
	resp, httpResp, err := c.client.SSHKeysAPI.SshKeysGet(ctx).Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return nil, wrapMassedComputeError(err, httpResp)
	}
	return resp, nil
}

func (c *MassedComputeClient) createSSHKey(ctx context.Context, req openapi.SshKeysPostRequest) (*openapi.POSTSSHKey, error) {
	resp, httpResp, err := c.client.SSHKeysAPI.SshKeysPost(ctx).
		SshKeysPostRequest(req).
		Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return nil, wrapMassedComputeError(err, httpResp)
	}
	return resp, nil
}

func (c *MassedComputeClient) convertInstanceToV1Instance(ctx context.Context, providerInstance openapi.RetrieveAllRunningInstancesV1RunningInstancesInner) (*v1.Instance, error) {
	providerName := stringValue(providerInstance.Name)
	refID, name := parseProviderInstanceName(providerName)
	sshUser := stringValue(providerInstance.Username)

	var instanceType string
	var storageGB int32
	if providerInstance.Product != nil {
		instanceType = stringValue(providerInstance.Product.Name)
		storageGB = int32Value(providerInstance.Product.Storage)
	}

	imageName := ""
	if providerInstance.Image != nil {
		imageName = strings.TrimSpace(stringValue(providerInstance.Image.Name))
		if imageName == "" && int32Value(providerInstance.Image.Id) > 0 {
			resolvedImageName, err := c.resolveImageName(ctx, int32Value(providerInstance.Image.Id))
			if err != nil {
				return nil, err
			}
			imageName = resolvedImageName
		}
	}

	ip := stringValue(providerInstance.Ip)
	storageBytes := v1.NewBytes(v1.BytesValue(storageGB), v1.Gigabyte)
	instance := &v1.Instance{
		Name:           name,
		RefID:          refID,
		CloudCredRefID: c.refID,
		CloudID:        v1.CloudProviderInstanceID(stringValue(providerInstance.Uuid)),
		PublicIP:       ip,
		PublicDNS:      ip,
		Hostname:       providerName,
		ImageID:        imageName,
		InstanceType:   instanceType,
		SSHUser:        sshUser,
		SSHPort:        defaultSSHPort,
		Status: v1.Status{
			LifecycleStatus: massedComputeLifecycleStatus(stringValue(providerInstance.Status)),
		},
		Location:      c.instanceLocation(providerInstance.AdditionalProperties),
		DiskSizeBytes: storageBytes,
		DiskSize:      legacyBytes(storageBytes),
	}
	if storageBytes.Value() > 0 {
		instance.VolumeType = "ssd"
	}
	if createdAt, err := time.Parse(time.RFC3339Nano, stringValue(providerInstance.Created)); err == nil {
		instance.CreatedAt = createdAt
	}
	instance.InstanceTypeID = v1.MakeGenericInstanceTypeIDFromInstance(*instance)
	return instance, nil
}

func (c *MassedComputeClient) instanceLocation(additionalProperties map[string]any) string {
	if region, ok := additionalProperties["region"].(map[string]any); ok {
		if location, ok := region["name"].(string); ok && strings.TrimSpace(location) != "" {
			return strings.TrimSpace(location)
		}
	}
	return c.location
}

func massedComputeLifecycleStatus(status string) v1.LifecycleStatus {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "active", "running", "rented":
		return v1.LifecycleStatusRunning
	case "stopping":
		return v1.LifecycleStatusStopping
	case "stopped":
		return v1.LifecycleStatusStopped
	case "terminating", "deleting":
		return v1.LifecycleStatusTerminating
	case "terminated", "deleted":
		return v1.LifecycleStatusTerminated
	case "failed", "error":
		return v1.LifecycleStatusFailed
	default:
		return v1.LifecycleStatusPending
	}
}

func managedResourceName(prefix string, separator string, refID string) string {
	suffix := strings.ToLower(refID)
	suffix = resourceNameInvalidCharacters.ReplaceAllString(suffix, separator)
	suffix = strings.Trim(suffix, separator)
	name := strings.Trim(strings.ToLower(prefix), separator)
	if suffix != "" {
		name += separator + suffix
	}
	if len(name) > 63 {
		name = strings.TrimRight(name[:63], separator)
	}
	return name
}

func makeProviderInstanceName(refID, name string) string {
	if name == "" {
		return refID
	}
	return refID + instanceNameDivider + name
}

func parseProviderInstanceName(providerName string) (string, string) {
	refID, name, found := strings.Cut(providerName, instanceNameDivider)
	if !found {
		return providerName, providerName
	}
	return refID, name
}

func closeResponseBody(response *http.Response) {
	if response != nil && response.Body != nil {
		_ = response.Body.Close()
	}
}

func (c *MassedComputeClient) GetInstancePollTime() time.Duration {
	return 10 * time.Second
}

func (c *MassedComputeClient) MergeInstanceForUpdate(_ v1.Instance, updated v1.Instance) v1.Instance {
	return updated
}

func (c *MassedComputeClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, updated v1.InstanceType) v1.InstanceType {
	return updated
}
