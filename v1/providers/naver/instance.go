package v1

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/units"
	cloud "github.com/brevdev/cloud/v1"
)

type serverInstanceListResponse struct {
	Create    serverInstanceList `json:"createServerInstancesResponse"`
	Get       serverInstanceList `json:"getServerInstanceListResponse"`
	Detail    serverInstanceList `json:"getServerInstanceDetailResponse"`
	Terminate responseMeta       `json:"terminateServerInstancesResponse"`
	Reboot    responseMeta       `json:"rebootServerInstancesResponse"`
	Stop      responseMeta       `json:"stopServerInstancesResponse"`
	Start     responseMeta       `json:"startServerInstancesResponse"`
}

func (r *serverInstanceListResponse) apiError() error {
	for _, meta := range []responseMeta{r.Create.responseMeta, r.Get.responseMeta, r.Detail.responseMeta, r.Terminate, r.Reboot, r.Stop, r.Start} {
		if err := meta.apiError(); err != nil {
			return err
		}
	}
	return nil
}

type serverInstanceList struct {
	responseMeta
	TotalRows          int                   `json:"totalRows"`
	ServerInstanceList []naverServerInstance `json:"serverInstanceList"`
}

type naverServerInstance struct {
	ServerInstanceNo                    string   `json:"serverInstanceNo"`
	ServerName                          string   `json:"serverName"`
	ServerDescription                   string   `json:"serverDescription"`
	CPUCount                            int32    `json:"cpuCount"`
	MemorySize                          int64    `json:"memorySize"`
	PlatformType                        codeName `json:"platformType"`
	LoginKeyName                        string   `json:"loginKeyName"`
	PublicIPInstanceNo                  string   `json:"publicIpInstanceNo"`
	PublicIP                            string   `json:"publicIp"`
	PrivateIP                           string   `json:"privateIp"`
	ServerInstanceStatus                codeName `json:"serverInstanceStatus"`
	ServerInstanceOperation             codeName `json:"serverInstanceOperation"`
	ServerInstanceStatusName            string   `json:"serverInstanceStatusName"`
	CreateDate                          string   `json:"createDate"`
	ServerImageProductCode              string   `json:"serverImageProductCode"`
	ServerImageNo                       string   `json:"serverImageNo"`
	ServerProductCode                   string   `json:"serverProductCode"`
	ServerSpecCode                      string   `json:"serverSpecCode"`
	ZoneCode                            string   `json:"zoneCode"`
	RegionCode                          string   `json:"regionCode"`
	VPCNo                               string   `json:"vpcNo"`
	SubnetNo                            string   `json:"subnetNo"`
	BaseBlockStorageSize                int64    `json:"baseBlockStorageSize"`
	BaseBlockStorageDiskType            codeName `json:"baseBlockStorageDiskType"`
	BaseBlockStorageDiskDetailType      codeName `json:"baseBlockStorageDiskDetailType"`
	ServerInstanceType                  codeName `json:"serverInstanceType"`
	IsProtectServerTermination          bool     `json:"isProtectServerTermination"`
	IsPreInstallGPUDriver               bool     `json:"isPreInstallGpuDriver"`
	NetworkInterfaceNoList              []string `json:"networkInterfaceNoList"`
	FabricClusterPoolNo                 string   `json:"fabricClusterPoolNo"`
	FabricClusterPoolName               string   `json:"fabricClusterPoolName"`
	FabricClusterMode                   string   `json:"fabricClusterMode"`
	HypervisorType                      codeName `json:"hypervisorType"`
	BaseBlockStorageDiskDetailTypeAlias codeName `json:"baseBlockStorageDiskDetailTypeCode"`
}

type importLoginKeyResponse struct {
	Response loginKeyList `json:"importLoginKeyResponse"`
}

func (r *importLoginKeyResponse) apiError() error {
	return r.Response.responseMeta.apiError()
}

type loginKeyList struct {
	responseMeta
	LoginKeyList []naverLoginKey `json:"loginKeyList"`
}

type naverLoginKey struct {
	KeyName     string `json:"keyName"`
	Fingerprint string `json:"fingerprint"`
	CreateDate  string `json:"createDate"`
}

func (c *NaverClient) CreateInstance(ctx context.Context, attrs cloud.CreateInstanceAttrs) (*cloud.Instance, error) {
	if attrs.VPCID == "" {
		return nil, fmt.Errorf("VPCID is required for NAVER VPC server creation")
	}
	if attrs.SubnetID == "" {
		return nil, fmt.Errorf("SubnetID is required for NAVER VPC server creation")
	}
	if attrs.InstanceType == "" {
		return nil, fmt.Errorf("InstanceType is required")
	}
	if attrs.ImageID == "" {
		return nil, fmt.Errorf("ImageID is required")
	}

	location := firstNonEmpty(attrs.Location, c.location, defaultRegionCode)
	keyName := ""
	if attrs.KeyPairName != nil {
		keyName = *attrs.KeyPairName
	}
	if attrs.PublicKey != "" {
		imported, err := c.importLoginKey(ctx, location, naverResourceName("brev-"+attrs.RefID), attrs.PublicKey)
		if err != nil {
			return nil, err
		}
		keyName = imported
	}
	if keyName == "" {
		return nil, fmt.Errorf("KeyPairName or PublicKey is required")
	}

	params := url.Values{}
	params.Set("regionCode", location)
	params.Set("vpcNo", attrs.VPCID)
	params.Set("subnetNo", attrs.SubnetID)
	params.Set("serverName", naverResourceName(firstNonEmpty(attrs.Name, attrs.RefID, c.refID)))
	params.Set("serverCreateCount", "1")
	params.Set("networkInterfaceList.1.networkInterfaceOrder", "0")
	params.Set("networkInterfaceList.1.subnetNo", attrs.SubnetID)
	params.Set("loginKeyName", keyName)
	params.Set("isProtectServerTermination", "false")
	params.Set("serverImageProductCode", attrs.ImageID)
	params.Set("serverProductCode", attrs.InstanceType)
	if attrs.UserDataBase64 != "" {
		params.Set("userData", attrs.UserDataBase64)
	}

	var resp serverInstanceListResponse
	if err := c.do(ctx, "createServerInstances", params, &resp); err != nil {
		return nil, err
	}
	if len(resp.Create.ServerInstanceList) != 1 {
		return nil, fmt.Errorf("expected 1 server instance, got %d", len(resp.Create.ServerInstanceList))
	}
	inst := c.convertServerInstance(resp.Create.ServerInstanceList[0])
	inst.RefID = attrs.RefID
	inst.CloudCredRefID = c.refID
	inst.Tags = attrs.Tags
	return inst, nil
}

func (c *NaverClient) GetInstance(ctx context.Context, id cloud.CloudProviderInstanceID) (*cloud.Instance, error) {
	params := url.Values{}
	if c.location != "" {
		params.Set("regionCode", c.location)
	}
	params.Set("serverInstanceNo", string(id))

	var resp serverInstanceListResponse
	if err := c.do(ctx, "getServerInstanceDetail", params, &resp); err != nil {
		return nil, err
	}
	if len(resp.Detail.ServerInstanceList) == 0 {
		return nil, fmt.Errorf("server instance %s not found", id)
	}
	return c.convertServerInstance(resp.Detail.ServerInstanceList[0]), nil
}

func (c *NaverClient) ListInstances(ctx context.Context, args cloud.ListInstancesArgs) ([]cloud.Instance, error) {
	params := url.Values{}
	if c.location != "" {
		params.Set("regionCode", c.location)
	}
	for i, id := range args.InstanceIDs {
		params.Set(indexedParam("serverInstanceNoList", i+1), string(id))
	}

	var resp serverInstanceListResponse
	if err := c.do(ctx, "getServerInstanceList", params, &resp); err != nil {
		return nil, err
	}
	instances := make([]cloud.Instance, 0, len(resp.Get.ServerInstanceList))
	for _, server := range resp.Get.ServerInstanceList {
		if len(args.Locations) > 0 && !args.Locations.IsAllowed(server.RegionCode) {
			continue
		}
		if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, cloud.CloudProviderInstanceID(server.ServerInstanceNo)) {
			continue
		}
		instances = append(instances, *c.convertServerInstance(server))
	}
	return instances, nil
}

func (c *NaverClient) TerminateInstance(ctx context.Context, id cloud.CloudProviderInstanceID) error {
	return c.instanceAction(ctx, "terminateServerInstances", id)
}

func (c *NaverClient) RebootInstance(ctx context.Context, id cloud.CloudProviderInstanceID) error {
	return c.instanceAction(ctx, "rebootServerInstances", id)
}

func (c *NaverClient) StopInstance(ctx context.Context, id cloud.CloudProviderInstanceID) error {
	return c.instanceAction(ctx, "stopServerInstances", id)
}

func (c *NaverClient) StartInstance(ctx context.Context, id cloud.CloudProviderInstanceID) error {
	return c.instanceAction(ctx, "startServerInstances", id)
}

func (c *NaverClient) instanceAction(ctx context.Context, action string, id cloud.CloudProviderInstanceID) error {
	params := url.Values{}
	if c.location != "" {
		params.Set("regionCode", c.location)
	}
	params.Set("serverInstanceNoList.1", string(id))
	var resp serverInstanceListResponse
	return c.do(ctx, action, params, &resp)
}

func (c *NaverClient) importLoginKey(ctx context.Context, location, keyName, publicKey string) (string, error) {
	params := url.Values{}
	params.Set("regionCode", location)
	params.Set("keyName", keyName)
	params.Set("publicKey", publicKey)

	var resp importLoginKeyResponse
	if err := c.do(ctx, "importLoginKey", params, &resp); err != nil {
		return "", err
	}
	if len(resp.Response.LoginKeyList) == 0 {
		return "", fmt.Errorf("importLoginKey returned no login keys")
	}
	return resp.Response.LoginKeyList[0].KeyName, nil
}

func (c *NaverClient) convertServerInstance(server naverServerInstance) *cloud.Instance {
	createdAt, _ := parseNaverTime(server.CreateDate)
	instanceType := firstNonEmpty(server.ServerProductCode, server.ServerSpecCode)
	imageID := firstNonEmpty(server.ServerImageProductCode, server.ServerImageNo)
	diskBytes := server.BaseBlockStorageSize
	volumeType := firstNonEmpty(server.BaseBlockStorageDiskDetailType.Code, server.BaseBlockStorageDiskType.Code)
	inst := &cloud.Instance{
		Name:              server.ServerName,
		CloudID:           cloud.CloudProviderInstanceID(server.ServerInstanceNo),
		CloudCredRefID:    c.refID,
		CreatedAt:         createdAt,
		PublicIP:          server.PublicIP,
		PublicDNS:         server.PublicIP,
		PrivateIP:         server.PrivateIP,
		ImageID:           imageID,
		InstanceType:      instanceType,
		DiskSize:          units.Base2Bytes(diskBytes),
		DiskSizeBytes:     cloud.NewBytes(cloud.BytesValue(diskBytes), cloud.Byte),
		VolumeType:        volumeType,
		SSHUser:           defaultSSHUser,
		SSHPort:           defaultSSHPort,
		VPCID:             server.VPCNo,
		SubnetID:          server.SubnetNo,
		Location:          server.RegionCode,
		SubLocation:       server.ZoneCode,
		Stoppable:         true,
		Rebootable:        true,
		FirewallRules:     cloud.FirewallRules{},
		InstanceTypeID:    "",
		AdditionalDisks:   nil,
		IPAllocationID:    optionalString(server.PublicIPInstanceNo),
		PubKeyFingerprint: "",
		Status: cloud.Status{
			LifecycleStatus: naverStatusToLifecycle(server.ServerInstanceStatus.Code, server.ServerInstanceStatusName, server.ServerInstanceOperation.Code),
		},
	}
	inst.InstanceTypeID = cloud.MakeGenericInstanceTypeIDFromInstance(*inst)
	return inst
}

func naverStatusToLifecycle(code, name, operation string) cloud.LifecycleStatus {
	normalized := strings.ToLower(strings.Join([]string{code, name, operation}, " "))
	switch {
	case strings.Contains(normalized, "termt") || strings.Contains(normalized, "terminating"):
		return cloud.LifecycleStatusTerminating
	case strings.Contains(normalized, "terminated"):
		return cloud.LifecycleStatusTerminated
	case strings.Contains(normalized, "nstop") || strings.Contains(normalized, "stopped"):
		return cloud.LifecycleStatusStopped
	case strings.Contains(normalized, "stop"):
		return cloud.LifecycleStatusStopping
	case strings.Contains(normalized, "run") || strings.Contains(normalized, "running"):
		return cloud.LifecycleStatusRunning
	case strings.Contains(normalized, "fail") || strings.Contains(normalized, "error"):
		return cloud.LifecycleStatusFailed
	default:
		return cloud.LifecycleStatusPending
	}
}

func parseNaverTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	for _, layout := range []string{"2006-01-02T15:04:05-0700", time.RFC3339} {
		if parsed, err := time.Parse(layout, value); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported NAVER time format %q", value)
}

func naverResourceName(value string) string {
	value = strings.ToLower(value)
	value = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	if value == "" || value[0] < 'a' || value[0] > 'z' {
		value = "n-" + value
	}
	if len(value) > 30 {
		value = value[:30]
	}
	value = strings.Trim(value, "-")
	if len(value) < 3 {
		value += strings.Repeat("0", 3-len(value))
	}
	return value
}

func indexedParam(prefix string, index int) string {
	return fmt.Sprintf("%s.%d", prefix, index)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func (c *NaverClient) MergeInstanceForUpdate(_ cloud.Instance, newInst cloud.Instance) cloud.Instance {
	return newInst
}

func (c *NaverClient) MergeInstanceTypeForUpdate(_ cloud.InstanceType, newIt cloud.InstanceType) cloud.InstanceType {
	return newIt
}
