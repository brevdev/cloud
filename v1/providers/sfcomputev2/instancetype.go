package v2

import (
	"context"
	"fmt"
	"sort"
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

// skuFreeCapacity returns, per instance SKU in the configured capacity, how many more
// instances can be created on that SKU right now: the SKU's current node allocation minus the
// number of non-terminated instances already on it. Counts are clamped at zero. Reads the
// per-SKU allocation from the capacity schedule and the per-SKU consumption from the instance
// list, issuing exactly two API calls.
func (c *SFCClientV2) skuFreeCapacity(ctx context.Context) (map[string]int, error) {
	capacityID := c.GetDefaultCapacityResourcePath()

	capResp, err := c.client.Pools.Fetch(ctx, capacityID, nil)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if capResp.PoolResponse == nil {
		return map[string]int{}, nil
	}

	now := time.Now().Unix()
	free := make(map[string]int)
	for skuID, schedule := range capResp.PoolResponse.AllocationSchedule.ByInstanceSku {
		free[skuID] = currentScheduleAllocation(schedule, now)
	}

	resp, err := c.client.Instances.List(ctx, operations.ListInstancesRequest{
		Workspace: c.GetWorkspaceResourcePath(),
		Pool:      []string{capacityID},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp.ListInstancesResponse != nil {
		for _, inst := range resp.ListInstancesResponse.Data {
			// Every non-terminated instance occupies a slot on its SKU, including failed ones.
			if inst.Status == components.InstanceStatusTerminated {
				continue
			}
			sku, ok := inst.GetInstanceSku().Get()
			if !ok || sku == nil {
				continue
			}
			free[sku.ID]--
		}
	}

	for skuID, n := range free {
		free[skuID] = max(n, 0)
	}
	return free, nil
}

// currentScheduleAllocation returns the NodeCount from the schedule entry whose
// [StartAt, EndAt) range is currently in effect. EndAt is null only on the final, unbounded
// entry. Returns 0 if no entry is in effect.
func currentScheduleAllocation(schedule []components.ScheduleEntry, now int64) int {
	for _, entry := range schedule {
		if entry.StartAt > now {
			continue
		}
		// A set, non-null EndAt bounds the range; the final entry's null EndAt is unbounded.
		if endAt, ok := entry.EndAt.Get(); ok && endAt != nil && now >= *endAt {
			continue
		}
		return entry.NodeCount
	}
	return 0
}

// availableSlots returns how many more instances can be created in the configured capacity,
// summed across every SKU.
func (c *SFCClientV2) availableSlots(ctx context.Context) (int, error) {
	free, err := c.skuFreeCapacity(ctx)
	if err != nil {
		return 0, errors.WrapAndTrace(err)
	}
	total := 0
	for _, n := range free {
		total += n
	}
	return total, nil
}

// selectAvailableSku returns an instance SKU in the configured capacity that still has a free
// node. SKUs are considered in sorted order so selection is deterministic; since the goal is to
// fully consume every SKU and the order doesn't matter, this drains one SKU before moving to the
// next. Returns an error if no SKU has free capacity.
func (c *SFCClientV2) selectAvailableSku(ctx context.Context) (string, error) {
	free, err := c.skuFreeCapacity(ctx)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	skuIDs := make([]string, 0, len(free))
	for skuID := range free {
		skuIDs = append(skuIDs, skuID)
	}
	sort.Strings(skuIDs)

	for _, skuID := range skuIDs {
		if free[skuID] > 0 {
			return skuID, nil
		}
	}
	return "", fmt.Errorf("no instance SKU with available capacity in %s", c.GetDefaultCapacityResourcePath())
}

func (c *SFCClientV2) GetLocations(_ context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	return []v1.Location{{
		Name:        sfcLocation,
		Description: fmt.Sprintf("sfc_%s_h100", sfcLocation),
		Available:   true,
	}}, nil
}
