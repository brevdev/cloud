package v2

import (
	"context"
	"fmt"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/sfcompute/sfc-go/models/components"
	"github.com/sfcompute/sfc-go/models/operations"
)

const (
	h100InstanceType = "h100.ib"
	sfcVCPU          = 112
	sfcGPUCount      = 8
	sfcLocation      = "sfc"
	diskTypeSSD      = "ssd"
	formFactorSXM5   = "sxm5"
)

type sfcInstanceTypeMetadata struct {
	diskBytes       v1.Bytes
	memoryBytes     v1.Bytes
	gpuVRAM         v1.Bytes
	vcpu            int32
	gpuCount        int32
	gpuManufacturer v1.Manufacturer
	architecture    v1.Architecture
	deployTime      time.Duration
	price           currency.Amount
	instanceTypeID  v1.InstanceTypeID
}

var h100InstanceTypeMetadata = func() sfcInstanceTypeMetadata {
	price, err := currency.NewAmount("16.00", "USD")
	if err != nil {
		panic(err)
	}
	m := sfcInstanceTypeMetadata{
		diskBytes:       v1.NewBytes(1500, v1.Gigabyte),
		memoryBytes:     v1.NewBytes(960, v1.Gigabyte),
		gpuVRAM:         v1.NewBytes(80, v1.Gigabyte),
		vcpu:            sfcVCPU,
		gpuCount:        sfcGPUCount,
		gpuManufacturer: v1.ManufacturerNVIDIA,
		architecture:    v1.ArchitectureX86_64,
		deployTime:      14 * time.Minute,
		price:           price,
	}

	// Compute the instance type ID from a representative InstanceType so it matches
	// what Brev expects when validating or storing the type.
	it := buildInstanceType(m, true)
	m.instanceTypeID = it.ID
	return m
}()

func buildInstanceType(m sfcInstanceTypeMetadata, isAvailable bool) v1.InstanceType {
	ramInt64, _ := m.memoryBytes.ByteCountInUnitInt64(v1.Gibibyte)
	ram := units.Base2Bytes(ramInt64 * int64(units.Gibibyte))

	vramInt64, _ := m.gpuVRAM.ByteCountInUnitInt64(v1.Gibibyte)
	vram := units.Base2Bytes(vramInt64 * int64(units.Gibibyte))

	diskInt64, _ := m.diskBytes.ByteCountInUnitInt64(v1.Gibibyte)
	diskSize := units.Base2Bytes(diskInt64 * int64(units.Gibibyte))

	it := v1.InstanceType{
		IsAvailable:         isAvailable,
		Type:                h100InstanceType,
		Memory:              ram,
		MemoryBytes:         m.memoryBytes,
		VCPU:                m.vcpu,
		Location:            sfcLocation,
		Stoppable:           false,
		Rebootable:          false,
		IsContainer:         false,
		Provider:            CloudProviderID,
		BasePrice:           &m.price,
		EstimatedDeployTime: &m.deployTime,
		SupportedGPUs: []v1.GPU{{
			Count:          m.gpuCount,
			Type:           "H100",
			Manufacturer:   m.gpuManufacturer,
			Name:           "H100",
			Memory:         vram,
			MemoryBytes:    m.gpuVRAM,
			NetworkDetails: formFactorSXM5,
		}},
		SupportedStorage: []v1.Storage{{
			Type:      diskTypeSSD,
			Count:     1,
			Size:      diskSize,
			SizeBytes: m.diskBytes,
		}},
		SupportedArchitectures: []v1.Architecture{m.architecture},
	}
	it.ID = v1.MakeGenericInstanceTypeID(it)
	return it
}

func (c *SFCClientV2) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	c.logger.Debug(ctx, "sfcv2: GetInstanceTypes start",
		v1.LogField("location", c.location),
	)

	available, err := c.availableSlots(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	if available <= 0 {
		c.logger.Debug(ctx, "sfcv2: GetInstanceTypes no available slots")
		return []v1.InstanceType{}, nil
	}

	instanceType := buildInstanceType(h100InstanceTypeMetadata, true)

	if !v1.IsSelectedByArgs(instanceType, args) {
		return []v1.InstanceType{}, nil
	}

	c.logger.Debug(ctx, "sfcv2: GetInstanceTypes end",
		v1.LogField("available slots", available),
	)

	return []v1.InstanceType{instanceType}, nil
}

// availableSlots returns how many more instances can be created in the configured capacity.
// It subtracts the count of non-terminated instances from the current capacity allocation.
func (c *SFCClientV2) availableSlots(ctx context.Context) (int, error) {
	allocated, err := c.currentCapacityAllocation(ctx)
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}

	active, err := c.activeInstanceCount(ctx)
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}

	return max(allocated-active, 0), nil
}

// currentCapacityAllocation returns the NodeAllocation from the most recent schedule entry
// in BrevProductionCapacityID that is currently in effect (EffectiveAt <= now).
func (c *SFCClientV2) currentCapacityAllocation(ctx context.Context) (int, error) {
	resp, err := c.client.Capacities.Fetch(ctx, BrevDefaultCapacityID, nil, nil)
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}
	if resp.CapacityResponse == nil {
		return 0, nil
	}

	now := time.Now().Unix()
	allocation := 0
	latestAt := int64(-1)
	for _, entry := range resp.CapacityResponse.AllocationSchedule.Total {
		if entry.EffectiveAt <= now && entry.EffectiveAt > latestAt {
			latestAt = entry.EffectiveAt
			allocation = entry.NodeAllocation
		}
	}
	return allocation, nil
}

// activeInstanceCount returns the number of non-terminated instances in BrevProductionCapacityID.
// All non-terminated instances occupy a slot in the capacity, including failed ones.
func (c *SFCClientV2) activeInstanceCount(ctx context.Context) (int, error) {
	capacityID := BrevDefaultCapacityID
	resp, err := c.client.Instances.List(ctx, operations.ListInstancesRequest{
		Capacity: &capacityID,
	})
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}
	if resp.ListInstancesResponse == nil {
		return 0, nil
	}

	count := 0
	for _, inst := range resp.ListInstancesResponse.Data {
		if inst.Status != components.InstanceStatusTerminated {
			count++
		}
	}
	return count, nil
}

func (c *SFCClientV2) GetLocations(_ context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	return []v1.Location{{
		Name:        sfcLocation,
		Description: fmt.Sprintf("sfc_%s_h100", sfcLocation),
		Available:   true,
	}}, nil
}
