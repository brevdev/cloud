package massedcompute

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"
)

var (
	gpuMemoryPattern             = regexp.MustCompile(`\(([0-9]+)GB\)`)
	descriptionAnnotationPattern = regexp.MustCompile(`\[[^]]*\]`)
	gpuMemoryGBByModel           = map[string]v1.BytesValue{
		"H100 NVL":  94,
		"B200 SXM6": 180,
		"B300 SXM6": 288,
	}
)

type massedComputeGPUDescription struct {
	count          int32
	model          string
	networkDetails string
	memoryBytes    v1.Bytes
}

func (c *MassedComputeClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	inventory, err := c.getInventory(ctx)
	if err != nil {
		return nil, err
	}

	instanceTypes := make([]v1.InstanceType, 0)
	for _, item := range inventory {
		if item.InstanceType == nil || item.InstanceType.Name == nil || item.InstanceType.Specs == nil {
			// Cannot parse instance type
			continue
		}

		typeName := strings.TrimSpace(*item.InstanceType.Name)
		if typeName == "" || isSpotDescription(stringValue(item.InstanceType.Description)) {
			continue
		}
		for _, region := range item.RegionsWithCapacityAvailable {
			location := stringValue(region.Name)
			if location == "" {
				continue
			}
			instanceType, err := massedComputeInstanceType(item, location)
			if err != nil {
				return nil, err
			}
			if instanceType.IsAvailable && v1.IsSelectedByArgs(instanceType, args) {
				instanceTypes = append(instanceTypes, instanceType)
			}
		}
	}

	return instanceTypes, nil
}

func (c *MassedComputeClient) getInventory(ctx context.Context) (map[string]openapi.GPUInventoryV1GpuInventoryValue, error) {
	result, httpResp, err := c.client.DefaultAPI.GpuInventoryGet(ctx).Execute()
	defer closeResponseBody(httpResp)
	if err != nil {
		return nil, wrapMassedComputeError(err, httpResp)
	}
	if result == nil || result.GpuInventory == nil {
		return nil, fmt.Errorf("massed compute instance-type response did not contain data")
	}

	return *result.GpuInventory, nil
}

func massedComputeInstanceType(item openapi.GPUInventoryV1GpuInventoryValue, location string) (v1.InstanceType, error) {
	providerType := item.InstanceType
	specs := providerType.Specs
	typeName := stringValue(providerType.Name)
	description := stringValue(providerType.Description)

	price, err := currency.NewAmountFromInt64(int64(int32Value(providerType.PriceCentsPerHour)), "USD")
	if err != nil {
		return v1.InstanceType{}, fmt.Errorf("parse price for massed compute instance type %s: %w", typeName, err)
	}

	memoryBytes := v1.NewBytes(v1.BytesValue(int32Value(specs.MemoryGib)), v1.Gibibyte)
	storageBytes := v1.NewBytes(v1.BytesValue(int32Value(specs.StorageGb)), v1.Gigabyte)
	architecture := massedComputeArchitecture(description)
	gpu := massedComputeGPUs(description)

	instanceType := v1.InstanceType{
		Type:                   typeName,
		Location:               location,
		Memory:                 legacyBytes(memoryBytes),
		MemoryBytes:            memoryBytes,
		VCPU:                   int32Value(specs.VcpuCount),
		SupportedArchitectures: []v1.Architecture{architecture},
		SupportedGPUs:          gpu,
		SupportedUsageClasses:  []string{"on-demand"},
		IsAvailable:            item.CapacityAvailable == nil || *item.CapacityAvailable > 0,
		BasePrice:              &price,
		Provider:               CloudProviderID,
	}
	if storageBytes.Value() > 0 {
		instanceType.SupportedStorage = []v1.Storage{{
			Type:      "ssd",
			Count:     1,
			Size:      legacyBytes(storageBytes),
			SizeBytes: storageBytes,
		}}
	}
	instanceType.ID = v1.MakeGenericInstanceTypeID(instanceType)
	return instanceType, nil
}

func massedComputeGPUs(description string) []v1.GPU {
	parsed, ok := parseMassedComputeGPUDescription(description)
	if !ok {
		return nil
	}

	gpu := v1.GPU{
		Count:          parsed.count,
		Manufacturer:   v1.ManufacturerNVIDIA,
		Name:           parsed.model,
		Type:           parsed.model,
		NetworkDetails: parsed.networkDetails,
		MemoryBytes:    parsed.memoryBytes,
	}
	if gpu.MemoryBytes.Value() > 0 {
		gpu.Memory = legacyBytes(gpu.MemoryBytes)
	}
	return []v1.GPU{gpu}
}

func massedComputeArchitecture(description string) v1.Architecture {
	lowerDescription := strings.ToLower(description)
	if strings.Contains(lowerDescription, "gh200") || strings.Contains(lowerDescription, "gb200") || strings.Contains(lowerDescription, "gb300") {
		return v1.ArchitectureARM64
	}
	return v1.ArchitectureX86_64
}

func isSpotDescription(description string) bool {
	return strings.Contains(strings.ToLower(description), "[spot]")
}

func parseMassedComputeGPUDescription(description string) (massedComputeGPUDescription, bool) {
	if isSpotDescription(description) {
		return massedComputeGPUDescription{}, false
	}

	// Massed Compute GPU descriptions start with "nx <model>".
	parts := strings.SplitN(strings.TrimSpace(description), " ", 2)
	if len(parts) != 2 {
		return massedComputeGPUDescription{}, false
	}
	countText, ok := strings.CutSuffix(strings.ToLower(parts[0]), "x")
	if !ok {
		return massedComputeGPUDescription{}, false
	}
	count, err := strconv.ParseInt(countText, 10, 32)
	if err != nil || count <= 0 {
		return massedComputeGPUDescription{}, false
	}

	// Retain the remaining description for later processing
	description = parts[1]
	memoryBytes := v1.Bytes{}

	// The description usually (but not always!) contains the GB memory size in parenthesis: "nx <model> (<memory>GB)"
	if match := gpuMemoryPattern.FindStringSubmatch(description); len(match) == 2 {
		if memoryGB, err := strconv.ParseInt(match[1], 10, 64); err == nil && memoryGB > 0 {
			memoryBytes = v1.NewBytes(v1.BytesValue(memoryGB), v1.Gigabyte)
		}
	}

	// Remove the memory size from the description
	description = gpuMemoryPattern.ReplaceAllString(description, "")

	// Annotations sometimes (but not always!) exist as additional information in the description: "nx <model> [annotation]"
	description = descriptionAnnotationPattern.ReplaceAllString(description, "")

	// The description is now in the format "<model> <interconnect>", but note that the model name may contain spaces, and the interconnect
	// may not be present.
	fields := strings.Fields(description)
	modelKey := strings.ToUpper(strings.Join(fields, " "))

	model := make([]string, 0, len(fields))
	networkDetails := ""
	for _, field := range fields {
		if interconnect, ok := massedComputeInterconnect(field); ok {
			// We reached the interconnect, so we can stop processing the fields
			networkDetails = interconnect
			continue
		}
		model = append(model, field)
	}
	if len(model) == 0 {
		return massedComputeGPUDescription{}, false
	}

	modelName := strings.Join(model, " ")
	// We have the model, but we failed to parse the memory size from the description. As a backup, we can use the model name to look up the memory size.
	if memoryBytes.Value() == 0 {
		if memoryGB, ok := gpuMemoryGBByModel[modelKey]; ok {
			memoryBytes = v1.NewBytes(memoryGB, v1.Gigabyte)
		}
	}

	return massedComputeGPUDescription{
		count:          int32(count),
		model:          modelName,
		networkDetails: networkDetails,
		memoryBytes:    memoryBytes,
	}, true
}

func massedComputeInterconnect(value string) (string, bool) {
	value = strings.ToUpper(strings.TrimSpace(value))
	if value == "NVL" || value == "NVLINK" {
		return "NVLink", true
	}
	if strings.HasPrefix(value, "SXM") {
		return value, true
	}
	return "", false
}

func legacyBytes(size v1.Bytes) units.Base2Bytes {
	return units.Base2Bytes(size.ByteCount().Int64())
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func int32Value(value *int32) int32 {
	if value == nil {
		return 0
	}
	return *value
}

func (c *MassedComputeClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
