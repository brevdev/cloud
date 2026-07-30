package verda

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"slices"
	"strings"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	verdago "github.com/verda-cloud/verdacloud-sdk-go/pkg/verda"
	"golang.org/x/crypto/ssh"
)

const (
	maxInstanceDescriptionLength = 100
	instanceIdentitySeparator    = "_"
	firewallResourceNamePrefix   = "brev-firewall"
	sshKeyResourceNamePrefix     = "brev-key"
	defaultSSHUser               = "root"
	defaultSSHPort               = 22
)

var resourceNameInvalidCharacters = regexp.MustCompile(`[^a-z0-9-]+`)

func (c *VerdaClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	location := attrs.Location
	if location == "" {
		location = c.location
	}
	if attrs.RefID == "" {
		return nil, errors.New("verda instance RefID is required")
	}
	if attrs.InstanceType == "" {
		return nil, errors.New("verda instance type is required")
	}
	if location == "" {
		return nil, errors.New("verda instance location is required")
	}
	if attrs.PublicKey == "" {
		return nil, errors.New("verda instance public key is required")
	}
	if attrs.UserDataBase64 != "" {
		return nil, errors.New("verda provider does not support instance user data")
	}

	// Verda lacks tags, so we will use the description field to house a small amount of Brev metadata
	description, err := makeInstanceDescription(attrs.RefID, c.refID)
	if err != nil {
		return nil, err
	}

	// "hostname" is essentially the Brev environment name
	hostname := strings.TrimSpace(attrs.Name)
	if hostname == "" {
		hostname = attrs.RefID
	}

	// Images are necessary at creation time, but can be derived by the instance type
	image, err := c.selectImage(ctx, attrs.InstanceType, attrs.ImageID)
	if err != nil {
		return nil, err
	}

	// SSH keys are independent resources, so must be created and cleaned up separately
	sshKeyID, err := c.ensureSSHKey(ctx, attrs.PublicKey, attrs.RefID)
	if err != nil {
		return nil, err
	}

	// Startup scripts are independent resources, so must be created and cleaned up separately
	startupScript, err := buildStartupScript(attrs.FirewallRules)
	if err != nil {
		return nil, errors.Join(err, c.cleanupManagedResources(ctx, attrs.RefID))
	}
	managedScript, err := c.client.StartupScripts.AddStartupScript(ctx, &verdago.CreateStartupScriptRequest{
		Name:   managedResourceName(firewallResourceNamePrefix, attrs.RefID),
		Script: startupScript,
	})
	if err != nil {
		return nil, errors.Join(wrapVerdaError(err), c.cleanupManagedResources(ctx, attrs.RefID))
	}

	request := verdago.CreateInstanceRequest{
		InstanceType:    attrs.InstanceType,
		Image:           image,
		Hostname:        hostname,
		Description:     description,
		SSHKeyIDs:       []string{sshKeyID},
		LocationCode:    location,
		StartupScriptID: &managedScript.ID,
		IsSpot:          attrs.UseSpot,
	}
	if diskSizeGiB := requestedDiskSizeGiB(attrs); diskSizeGiB > 0 {
		request.OSVolume = &verdago.OSVolumeCreateRequest{
			Name: managedResourceName(hostname+"-os", attrs.RefID),
			Size: diskSizeGiB,
		}
		if attrs.UseSpot {
			request.OSVolume.OnSpotDiscontinue = verdago.SpotDiscontinueDeletePermanent
		}
	}

	verdaInstance, err := c.client.Instances.Create(ctx, request)
	if err != nil {
		return nil, errors.Join(wrapVerdaError(err), c.cleanupManagedResources(ctx, attrs.RefID))
	}
	return c.verdaInstanceToInstance(verdaInstance), nil
}

func (c *VerdaClient) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	verdaInstance, err := c.client.Instances.GetByID(ctx, string(id))
	if err != nil {
		return nil, wrapVerdaError(err)
	}
	return c.verdaInstanceToInstance(verdaInstance), nil
}

func (c *VerdaClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	verdaInstances, err := c.client.Instances.Get(ctx, "")
	if err != nil {
		return nil, wrapVerdaError(err)
	}

	instances := make([]v1.Instance, 0, len(verdaInstances))
	for i := range verdaInstances {
		instance := c.verdaInstanceToInstance(&verdaInstances[i])
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

func (c *VerdaClient) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	verdaInstance, err := c.client.Instances.GetByID(ctx, string(id))
	if err != nil {
		wrappedErr := wrapVerdaError(err)
		if errors.Is(wrappedErr, v1.ErrInstanceNotFound) {
			return nil
		}
		return wrappedErr
	}

	volumeIDs := verdaInstance.VolumeIDs
	if verdaInstance.OSVolumeID != nil && !slices.Contains(volumeIDs, *verdaInstance.OSVolumeID) {
		volumeIDs = append(volumeIDs, *verdaInstance.OSVolumeID)
	}
	err = c.performInstanceAction(ctx, verdago.InstanceActionRequest{
		Action:            verdago.ActionDelete,
		ID:                []string{string(id)},
		VolumeIDs:         volumeIDs,
		DeletePermanently: true,
	})
	if err != nil {
		return wrapVerdaError(err)
	}

	refID, _ := parseInstanceDescription(verdaInstance.Description)
	return c.cleanupManagedResources(ctx, refID)
}

func (c *VerdaClient) StopInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return c.performInstanceAction(ctx, verdago.InstanceActionRequest{
		Action: verdago.ActionShutdown,
		ID:     []string{string(id)},
	})
}

func (c *VerdaClient) StartInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return c.performInstanceAction(ctx, verdago.InstanceActionRequest{
		Action: verdago.ActionStart,
		ID:     []string{string(id)},
	})
}

func (c *VerdaClient) performInstanceAction(ctx context.Context, request verdago.InstanceActionRequest) error {
	results, err := c.client.Instances.Action(ctx, request)
	if err != nil {
		return wrapVerdaError(err)
	}
	for _, result := range results {
		if strings.EqualFold(result.Status, "error") {
			return fmt.Errorf(
				"verda instance %s action %s failed: %s",
				result.InstanceID,
				result.Action,
				result.Error,
			)
		}
	}
	return nil
}

func (c *VerdaClient) selectImage(ctx context.Context, instanceType string, requestedImage string) (string, error) {
	if requestedImage != "" {
		return requestedImage, nil
	}
	images, err := c.client.Images.GetImagesByInstanceType(ctx, instanceType)
	if err != nil {
		return "", wrapVerdaError(err)
	}
	for _, image := range images {
		if image.IsDefault && !image.IsCluster {
			return image.ImageType, nil
		}
	}
	for _, image := range images {
		if !image.IsCluster {
			return image.ImageType, nil
		}
	}
	return "", fmt.Errorf("verda has no compatible image for instance type %s", instanceType)
}

func (c *VerdaClient) ensureSSHKey(ctx context.Context, publicKey string, refID string) (string, error) {
	publicKey, err := normalizeSSHPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	keys, err := c.client.SSHKeys.GetAllSSHKeys(ctx)
	if err != nil {
		return "", wrapVerdaError(err)
	}
	for _, key := range keys {
		if strings.TrimSpace(key.PublicKey) == strings.TrimSpace(publicKey) {
			return key.ID, nil
		}
	}

	key, err := c.client.SSHKeys.AddSSHKey(ctx, &verdago.CreateSSHKeyRequest{
		Name:      managedResourceName(sshKeyResourceNamePrefix, refID),
		PublicKey: publicKey,
	})
	if err != nil {
		return "", wrapVerdaError(err)
	}
	return key.ID, nil
}

func normalizeSSHPublicKey(publicKey string) (string, error) {
	publicKey = strings.TrimSpace(publicKey)
	if key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey)); err == nil {
		return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("verda public key must be OpenSSH or PEM encoded")
	}
	parsedKey, pkixErr := x509.ParsePKIXPublicKey(block.Bytes)
	if pkixErr != nil {
		rsaKey, pkcs1Err := x509.ParsePKCS1PublicKey(block.Bytes)
		if pkcs1Err != nil {
			return "", fmt.Errorf("parse verda PEM public key: %w", errors.Join(pkixErr, pkcs1Err))
		}
		parsedKey = rsaKey
	}
	key, err := ssh.NewPublicKey(parsedKey)
	if err != nil {
		return "", fmt.Errorf("convert verda public key to OpenSSH: %w", err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
}

func (c *VerdaClient) cleanupManagedResources(ctx context.Context, refID string) error {
	if refID == "" {
		return nil
	}

	var cleanupErrors []error
	scriptName := managedResourceName(firewallResourceNamePrefix, refID)
	scripts, err := c.client.StartupScripts.GetAllStartupScripts(ctx)
	if err != nil {
		cleanupErrors = append(cleanupErrors, wrapVerdaError(err))
	} else {
		for _, script := range scripts {
			if script.Name != scriptName {
				continue
			}
			if err := c.client.StartupScripts.DeleteStartupScript(ctx, script.ID); err != nil {
				cleanupErrors = append(cleanupErrors, wrapVerdaError(err))
			}
		}
	}

	keyName := managedResourceName(sshKeyResourceNamePrefix, refID)
	keys, err := c.client.SSHKeys.GetAllSSHKeys(ctx)
	if err != nil {
		cleanupErrors = append(cleanupErrors, wrapVerdaError(err))
	} else {
		for _, key := range keys {
			if key.Name != keyName {
				continue
			}
			if err := c.client.SSHKeys.DeleteSSHKey(ctx, key.ID); err != nil {
				cleanupErrors = append(cleanupErrors, wrapVerdaError(err))
			}
		}
	}
	return errors.Join(cleanupErrors...)
}

func (c *VerdaClient) verdaInstanceToInstance(verdaInstance *verdago.Instance) *v1.Instance {
	refID, cloudCredRefID := parseInstanceDescription(verdaInstance.Description)
	if cloudCredRefID == "" {
		cloudCredRefID = c.refID
	}

	publicIP := ""
	if verdaInstance.IP != nil {
		publicIP = *verdaInstance.IP
	}

	instance := &v1.Instance{
		Name:           verdaInstance.Hostname,
		RefID:          refID,
		CloudCredRefID: cloudCredRefID,
		CreatedAt:      verdaInstance.CreatedAt,
		CloudID:        v1.CloudProviderInstanceID(verdaInstance.ID),
		PublicIP:       publicIP,
		PublicDNS:      publicIP,
		Hostname:       verdaInstance.Hostname,
		ImageID:        verdaInstance.Image,
		InstanceType:   verdaInstance.InstanceType,
		SSHUser:        defaultSSHUser,
		SSHPort:        defaultSSHPort,
		Status: v1.Status{
			LifecycleStatus: verdaStatusToLifecycleStatus(verdaInstance.Status),
		},
		Location:   verdaInstance.Location,
		Spot:       verdaInstance.IsSpot,
		Stoppable:  true,
		Rebootable: false,
	}
	instance.InstanceTypeID = v1.MakeGenericInstanceTypeIDFromInstance(*instance)

	if storage := storageDescriptionToStorage(verdaInstance.Storage.Description); len(storage) > 0 {
		instance.DiskSize = storage[0].Size
		instance.DiskSizeBytes = storage[0].SizeBytes
		instance.VolumeType = storage[0].Type
	}
	return instance
}

func verdaStatusToLifecycleStatus(status string) v1.LifecycleStatus {
	switch status {
	case verdago.StatusRunning:
		return v1.LifecycleStatusRunning
	case verdago.StatusOffline:
		return v1.LifecycleStatusStopped
	case verdago.StatusDeleting:
		return v1.LifecycleStatusTerminating
	case verdago.StatusDiscontinued, verdago.StatusNotFound:
		return v1.LifecycleStatusTerminated
	case verdago.StatusError, verdago.StatusNoCapacity:
		return v1.LifecycleStatusFailed
	default:
		return v1.LifecycleStatusPending
	}
}

func makeInstanceDescription(refID string, cloudCredRefID string) (string, error) {
	description := refID + instanceIdentitySeparator + cloudCredRefID
	if len(description) > maxInstanceDescriptionLength {
		return "", fmt.Errorf(
			"verda instance identity is %d characters; maximum description length is %d",
			len(description),
			maxInstanceDescriptionLength,
		)
	}
	return description, nil
}

func parseInstanceDescription(description string) (refID, cloudCredRefID string) {
	refID, cloudCredRefID, found := strings.Cut(description, instanceIdentitySeparator)
	if !found {
		return description, ""
	}
	return refID, cloudCredRefID
}

func managedResourceName(prefix, refID string) string {
	suffix := strings.ToLower(refID)
	suffix = resourceNameInvalidCharacters.ReplaceAllString(suffix, "-")
	suffix = strings.Trim(suffix, "-")
	name := strings.Trim(strings.ToLower(prefix), "-")
	if suffix != "" {
		name += "-" + suffix
	}
	if len(name) > 63 {
		name = strings.TrimRight(name[:63], "-")
	}
	return name
}

func requestedDiskSizeGiB(attrs v1.CreateInstanceAttrs) int {
	var byteCount *big.Int
	switch {
	case attrs.DiskSizeBytes.Value() > 0:
		byteCount = attrs.DiskSizeBytes.ByteCount()
	case attrs.DiskSize > 0:
		byteCount = big.NewInt(int64(attrs.DiskSize))
	default:
		return 0
	}

	bytesPerGiB := big.NewInt(1 << 30)
	sizeGiB, remainder := new(big.Int).QuoRem(byteCount, bytesPerGiB, new(big.Int))
	if remainder.Sign() > 0 {
		sizeGiB.Add(sizeGiB, big.NewInt(1))
	}
	if !sizeGiB.IsInt64() || sizeGiB.Int64() > int64(^uint(0)>>1) {
		return 0
	}
	return int(sizeGiB.Int64())
}

func (c *VerdaClient) GetInstancePollTime() time.Duration {
	return 10 * time.Second
}

func (c *VerdaClient) MergeInstanceForUpdate(_ v1.Instance, newInstance v1.Instance) v1.Instance {
	return newInstance
}

func (c *VerdaClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newInstanceType v1.InstanceType) v1.InstanceType {
	return newInstanceType
}
