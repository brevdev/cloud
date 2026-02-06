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
	formFactorSXM5         = "sxm5"
	diskTypeSSD            = "ssd"
)

var allowedZones = []string{"hayesvalley", "yerba"}

func makeDefaultInstanceTypePrice(amount string, currencyCode string) currency.Amount {
	instanceTypePrice, err := currency.NewAmount(amount, currencyCode)
	if err != nil {
		panic(err)
	}
	return instanceTypePrice
}

func (c *SFCClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	c.logger.Debug(ctx, "sfc: GetInstanceTypes start",
		v1.LogField("location", c.location),
		v1.LogField("args", fmt.Sprintf("%+v", args)),
	)

	// Fetch all available zones
	includeUnavailable := false
	zones, err := c.getZones(ctx, includeUnavailable)
	if err != nil {
		return nil, err
	}

	c.logger.Debug(ctx, "sfc: GetInstanceTypes zones list",
		v1.LogField("zone count", len(zones)),
	)

	instanceTypes := make([]v1.InstanceType, 0, len(zones))
	for _, zone := range zones {
		gpuType := strings.ToLower(string(zone.HardwareType))

		if !gpuTypeIsAllowed(gpuType) {
			c.logger.Debug(ctx, "sfc: GetInstanceTypes gpu type not allowed",
				v1.LogField("gpuType", gpuType),
			)
			continue
		}

		instanceType, err := getInstanceTypeForZone(zone)
		if err != nil {
			return nil, err
		}

		if !v1.IsSelectedByArgs(*instanceType, args) {
			c.logger.Debug(ctx, "sfc: GetInstanceTypes instance type not selected by args",
				v1.LogField("instanceType", instanceType.Type),
			)
			continue
		}

		instanceTypes = append(instanceTypes, *instanceType)
	}

	c.logger.Debug(ctx, "sfc: GetInstanceTypes end",
		v1.LogField("instanceType count", len(instanceTypes)),
	)

	return instanceTypes, nil
}

func getInstanceTypeForZone(zone sfcnodes.ZoneListResponseData) (*v1.InstanceType, error) {
	gpuType := strings.ToLower(string(zone.HardwareType))

	gpuMetadata, err := getInstanceTypeMetadata(gpuType)
	if err != nil {
		return nil, err
	}

	ramInt64, err := gpuMetadata.memoryBytes.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	ram := units.Base2Bytes(ramInt64 * int64(units.Gibibyte))

	memoryInt64, err := gpuMetadata.gpuVRAM.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	memory := units.Base2Bytes(memoryInt64 * int64(units.Gibibyte))

	diskSizeInt64, err := gpuMetadata.diskBytes.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	diskSize := units.Base2Bytes(diskSizeInt64 * int64(units.Gibibyte))

	instanceType := v1.InstanceType{
		IsAvailable:         true,
		Type:                makeInstanceTypeName(zone),
		Memory:              ram,
		MemoryBytes:         gpuMetadata.memoryBytes,
		Location:            zoneToLocation(zone).Name,
		Stoppable:           false,
		Rebootable:          false,
		IsContainer:         false,
		Provider:            CloudProviderID,
		BasePrice:           &gpuMetadata.price,
		EstimatedDeployTime: &gpuMetadata.estimatedDeployTime,
		SupportedGPUs: []v1.GPU{{
			Count:          gpuMetadata.gpuCount,
			Type:           strings.ToUpper(gpuType),
			Manufacturer:   gpuMetadata.gpuManufacturer,
			Name:           strings.ToUpper(gpuType),
			Memory:         memory,
			MemoryBytes:    gpuMetadata.gpuVRAM,
			NetworkDetails: gpuMetadata.formFactor,
		}},
		SupportedStorage: []v1.Storage{{
			Type:      diskTypeSSD,
			Count:     1,
			Size:      diskSize,
			SizeBytes: gpuMetadata.diskBytes,
		}},
		SupportedArchitectures: []v1.Architecture{gpuMetadata.architecture},
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

// sfcInstanceTypeMetadata is a struct that contains the metadata for a given instance type.
// These values are not currently provided by the SFCompute API, so we need to hardcode them.
type sfcInstanceTypeMetadata struct {
	gpuType             string
	formFactor          string
	architecture        v1.Architecture
	memoryBytes         v1.Bytes
	diskBytes           v1.Bytes
	gpuCount            int32
	gpuManufacturer     v1.Manufacturer
	gpuVRAM             v1.Bytes
	estimatedDeployTime time.Duration
	price               currency.Amount
}

func getInstanceTypeMetadata(gpuType string) (*sfcInstanceTypeMetadata, error) {
	switch gpuType {
	case gpuTypeH100:
		return &h100InstanceTypeMetadata, nil
	case gpuTypeH200:
		return &h200InstanceTypeMetadata, nil
	default:
		return nil, fmt.Errorf("invalid GPU type: %s", gpuType)
	}
}

var h100InstanceTypeMetadata = sfcInstanceTypeMetadata{
	gpuType:             gpuTypeH100,
	formFactor:          formFactorSXM5,
	architecture:        v1.ArchitectureX86_64,
	memoryBytes:         v1.NewBytes(960, v1.Gigabyte),
	diskBytes:           v1.NewBytes(1500, v1.Gigabyte),
	gpuCount:            8,
	gpuManufacturer:     v1.ManufacturerNVIDIA,
	gpuVRAM:             v1.NewBytes(80, v1.Gigabyte),
	estimatedDeployTime: 14 * time.Minute,
	price:               makeDefaultInstanceTypePrice("16.00", "USD"),
}

var h200InstanceTypeMetadata = sfcInstanceTypeMetadata{
	gpuType:             gpuTypeH200,
	formFactor:          formFactorSXM5,
	architecture:        v1.ArchitectureX86_64,
	memoryBytes:         v1.NewBytes(960, v1.Gigabyte),
	diskBytes:           v1.NewBytes(1500, v1.Gigabyte),
	gpuCount:            8,
	gpuManufacturer:     v1.ManufacturerNVIDIA,
	gpuVRAM:             v1.NewBytes(141, v1.Gigabyte),
	estimatedDeployTime: 14 * time.Minute,
	price:               makeDefaultInstanceTypePrice("24.00", "USD"),
}
