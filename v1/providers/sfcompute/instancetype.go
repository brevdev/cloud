package v1

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	sfcnodes "github.com/sfcompute/nodes-go"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	gpuTypeH100 = "h100"
	gpuTypeH200 = "h200"

	deliveryTypeVM         = "VM"
	interconnectInfiniband = "infiniband"
)

var (
	allowedZones = []string{"hayesvalley", "yerba"}

	gpuToVRAM = map[string]v1.Bytes{
		gpuTypeH100: v1.NewBytes(80, v1.Gigabyte),
		gpuTypeH200: v1.NewBytes(141, v1.Gigabyte),
	}
	gpuToFormFactor = map[string]string{
		gpuTypeH100: "sxm5",
		gpuTypeH200: "sxm5",
	}
	gpuToArchitecture = map[string]v1.Architecture{
		gpuTypeH100: v1.ArchitectureX86_64,
		gpuTypeH200: v1.ArchitectureX86_64,
	}

	defaultGPUCountPerNode  = int32(8)
	defaultGPUManufacturer  = "nvidia"
	defaultRAMPerNode       = v1.NewBytes(960, v1.Gigabyte)
	defaultStoragePerNode   = v1.NewBytes(1500, v1.Gigabyte)
	defaultProvisioningTime = 5 * time.Minute
	defaultPricePerGPU      = makeDefaultInstanceTypePrice("2.00", "USD")
)

func makeDefaultInstanceTypePrice(amount string, currencyCode string) currency.Amount {
	instanceTypePrice, err := currency.NewAmount(amount, currencyCode)
	if err != nil {
		panic(err)
	}
	return instanceTypePrice
}

func (c *SFCClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	// Fetch all available zones
	includeUnavailable := false
	zones, err := c.getZones(ctx, includeUnavailable)
	if err != nil {
		return nil, err
	}

	instanceTypes := make([]v1.InstanceType, 0, len(zones))
	for _, zone := range zones {
		gpuType := strings.ToLower(string(zone.HardwareType))

		if !gpuTypeIsAllowed(gpuType) {
			continue
		}

		instanceType, err := getInstanceTypeForZone(zone)
		if err != nil {
			return nil, err
		}

		if v1.IsSelectedByArgs(*instanceType, args) {
			instanceTypes = append(instanceTypes, *instanceType)
		}
	}

	return instanceTypes, nil
}

func getInstanceTypeForZone(zone sfcnodes.ZoneListResponseData) (*v1.InstanceType, error) {
	gpuType := strings.ToLower(string(zone.HardwareType))

	ramInt64, err := defaultRAMPerNode.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	ram := units.Base2Bytes(ramInt64 * int64(units.Gibibyte))

	memoryInt64, err := gpuToVRAM[gpuType].ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	memory := units.Base2Bytes(memoryInt64 * int64(units.Gibibyte))

	diskSizeInt64, err := defaultStoragePerNode.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	diskSize := units.Base2Bytes(diskSizeInt64 * int64(units.Gibibyte))

	instanceType := v1.InstanceType{
		IsAvailable:         true,
		Type:                makeInstanceTypeName(zone),
		Memory:              ram,
		MemoryBytes:         defaultRAMPerNode,
		Location:            zoneToLocation(zone).Name,
		Stoppable:           false,
		Rebootable:          false,
		IsContainer:         false,
		Provider:            CloudProviderID,
		BasePrice:           &defaultPricePerGPU,
		EstimatedDeployTime: &defaultProvisioningTime,
		SupportedGPUs: []v1.GPU{{
			Count:          defaultGPUCountPerNode,
			Type:           strings.ToUpper(gpuType),
			Manufacturer:   v1.GetManufacturer(defaultGPUManufacturer),
			Name:           strings.ToUpper(gpuType),
			Memory:         memory,
			MemoryBytes:    gpuToVRAM[gpuType],
			NetworkDetails: gpuToFormFactor[gpuType],
		}},
		SupportedStorage: []v1.Storage{{
			Type:      "ssd",
			Count:     1,
			Size:      diskSize,
			SizeBytes: defaultStoragePerNode,
		}},
		SupportedArchitectures: []v1.Architecture{gpuToArchitecture[gpuType]},
	}

	instanceType.ID = v1.MakeGenericInstanceTypeID(instanceType)

	return &instanceType, nil
}

func gpuTypeIsAllowed(gpuType string) bool {
	return gpuType == gpuTypeH100 || gpuType == gpuTypeH200
}

func makeInstanceTypeName(zone sfcnodes.ZoneListResponseData) string {
	interconnect := ""
	if strings.ToLower(zone.InterconnectType) == interconnectInfiniband {
		interconnect = ".ib"
	}
	return fmt.Sprintf("%s%s", strings.ToLower(string(zone.HardwareType)), interconnect)
}

func (c *SFCClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	zones, err := c.getZones(ctx, args.IncludeUnavailable)
	if err != nil {
		return nil, err
	}

	locations := make([]v1.Location, 0, len(zones))
	for _, zone := range zones {
		location := zoneToLocation(zone)
		locations = append(locations, location)
	}

	return locations, nil
}

func (c *SFCClient) getZones(ctx context.Context, includeUnavailable bool) ([]sfcnodes.ZoneListResponseData, error) {
	// Fetch the zones from the API
	resp, err := c.client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}

	// If there are no zones, return an empty list
	if resp == nil || len(resp.Data) == 0 {
		return []sfcnodes.ZoneListResponseData{}, nil
	}

	zones := make([]sfcnodes.ZoneListResponseData, 0, len(resp.Data))
	for _, zone := range resp.Data {
		// If the zone is not allowed, skip it
		if !slices.Contains(allowedZones, strings.ToLower(zone.Name)) {
			continue
		}

		// If the there is no available capacity, and skip it
		if len(zone.AvailableCapacity) == 0 && !includeUnavailable {
			continue
		}

		// If the delivery type is not VM, skip it
		if zone.DeliveryType != deliveryTypeVM {
			continue
		}

		// Add the zone to the list
		zones = append(zones, zone)
	}

	return zones, nil
}

func zoneToLocation(zone sfcnodes.ZoneListResponseData) v1.Location {
	return v1.Location{
		Name:        zone.Name,
		Description: fmt.Sprintf("sfc_%s_%s", zone.Name, string(zone.HardwareType)),
		Available:   true,
	}
}
