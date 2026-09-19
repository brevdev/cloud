package hyperstack

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/flavor"
	"github.com/alecthomas/units"
	"github.com/bojanz/currency"

	v1 "github.com/brevdev/cloud/v1"
)

var gpuMemoryPattern = regexp.MustCompile(`(?i)-(\d+)(?:G|GB)(?:-|$)`)

func (c *HyperstackClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	flavorGroups, err := c.listFlavors(ctx)
	if err != nil {
		return nil, err
	}
	rates, err := c.getPricebook(ctx)
	if err != nil {
		return nil, err
	}

	instanceTypes := make([]v1.InstanceType, 0)
	for _, group := range flavorGroups {
		if group.Flavors == nil {
			continue
		}
		for _, providerType := range *group.Flavors {
			instanceType, err := hyperstackInstanceType(providerType, stringValue(group.RegionName), rates)
			if err != nil {
				return nil, err
			}
			if instanceType.Type != "" && v1.IsSelectedByArgs(instanceType, args) {
				instanceTypes = append(instanceTypes, instanceType)
			}
		}
	}

	sort.Slice(instanceTypes, func(i, j int) bool {
		return instanceTypes[i].ID < instanceTypes[j].ID
	})
	return instanceTypes, nil
}

func (c *HyperstackClient) listFlavors(ctx context.Context) ([]flavor.FlavorItemGetResponse, error) {
	response, err := c.flavors.ListFlavorsWithResponse(ctx, nil)
	if err != nil {
		return nil, wrapTransportError("list flavors", err)
	}
	if response.StatusCode() != 200 {
		return nil, responseError("list flavors", response.StatusCode(), response.Body, nil)
	}
	if response.JSON200 == nil || response.JSON200.Data == nil {
		return nil, fmt.Errorf("hyperstack list flavors response did not contain data")
	}
	return *response.JSON200.Data, nil
}

type pricebookEntry struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

func (c *HyperstackClient) getPricebook(ctx context.Context) (map[string]string, error) {
	response, err := c.pricebook.GetPricebook(ctx)
	if err != nil {
		return nil, wrapTransportError("get pricebook", err)
	}
	defer func() { _ = response.Body.Close() }()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, wrapTransportError("read pricebook", err)
	}
	if response.StatusCode != 200 {
		return nil, responseError("get pricebook", response.StatusCode, body, nil)
	}

	var entries []pricebookEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, fmt.Errorf("decode hyperstack pricebook: %w", err)
	}
	rates := make(map[string]string, len(entries))
	for _, entry := range entries {
		name := strings.ToLower(strings.TrimSpace(entry.Name))
		value := strings.Trim(strings.TrimSpace(string(entry.Value)), `"`)
		if name != "" && value != "" && value != "null" {
			rates[name] = value
		}
	}
	return rates, nil
}

func hyperstackInstanceType(providerType flavor.FlavorFields, fallbackLocation string, rates map[string]string) (v1.InstanceType, error) {
	typeName := strings.TrimSpace(stringValue(providerType.Name))
	location := strings.TrimSpace(stringValue(providerType.RegionName))
	if location == "" {
		location = fallbackLocation
	}
	if typeName == "" || location == "" {
		return v1.InstanceType{}, nil
	}

	memoryGB := int64(math.Round(float64(float32Value(providerType.Ram))))
	memory, memoryBytes := byteSizes(memoryGB, v1.Gigabyte)

	storageGB := int64(intValue(providerType.Disk))
	storage, storageBytes := byteSizes(storageGB, v1.Gigabyte)

	gpuType := strings.TrimSpace(stringValue(providerType.Gpu))
	gpuCount := intValue(providerType.GpuCount)

	basePrice, err := flavorPrice(providerType, rates)
	if err != nil {
		return v1.InstanceType{}, fmt.Errorf("price hyperstack instance type %s in %s: %w", typeName, location, err)
	}

	usageClass := "on-demand"
	preemptible := isSpotFlavor(typeName, gpuType)
	if preemptible {
		usageClass = "spot"
	}
	instanceType := v1.InstanceType{
		Type:                   typeName,
		Location:               location,
		Memory:                 memory,
		MemoryBytes:            memoryBytes,
		VCPU:                   int32(intValue(providerType.Cpu)),
		SupportedArchitectures: []v1.Architecture{v1.ArchitectureX86_64},
		SupportedUsageClasses:  []string{usageClass},
		Preemptible:            preemptible,
		Stoppable:              true,
		IsAvailable:            providerType.StockAvailable == nil || *providerType.StockAvailable,
		BasePrice:              basePrice,
		Provider:               CloudProviderID,
	}
	if storageGB > 0 {
		instanceType.SupportedStorage = []v1.Storage{{
			Type:      "ssd",
			Count:     1,
			Size:      storage,
			SizeBytes: storageBytes,
		}}
	}
	if ephemeralGB := int64(intValue(providerType.Ephemeral)); ephemeralGB > 0 {
		ephemeral, ephemeralBytes := byteSizes(ephemeralGB, v1.Gigabyte)
		instanceType.SupportedStorage = append(instanceType.SupportedStorage, v1.Storage{
			Type:                    "ephemeral",
			Count:                   1,
			Size:                    ephemeral,
			SizeBytes:               ephemeralBytes,
			IsEphemeral:             true,
			IsAdditionalDisk:        true,
			RequiresVolumeMountPath: true,
		})
	}
	if gpuCount > 0 && gpuType != "" {
		instanceType.SupportedGPUs = []v1.GPU{hyperstackGPU(gpuType, gpuCount)}
	}
	instanceType.ID = v1.MakeGenericInstanceTypeID(instanceType)
	return instanceType, nil
}

func flavorPrice(providerType flavor.FlavorFields, rates map[string]string) (*currency.Amount, error) {
	gpuCount := intValue(providerType.GpuCount)
	gpuType := strings.TrimSpace(stringValue(providerType.Gpu))
	if gpuCount > 0 && gpuType != "" {
		price, err := rateCost(rates, gpuType, gpuCount)
		return &price, err
	}

	zero, err := currency.NewAmount("0", "USD")
	if err != nil {
		return nil, err
	}
	total := zero
	resources := []struct {
		name  string
		count int
	}{
		{name: "vCPU (cpu-only-flavors)", count: intValue(providerType.Cpu)},
		{name: "RAM (cpu-only-flavors)", count: int(math.Round(float64(float32Value(providerType.Ram))))},
		{name: "hypervisor-local-storage (cpu-only-flavors)", count: intValue(providerType.Disk) + intValue(providerType.Ephemeral)},
	}
	for _, resource := range resources {
		cost, err := rateCost(rates, resource.name, resource.count)
		if err != nil {
			return nil, err
		}
		total, err = total.Add(cost)
		if err != nil {
			return nil, err
		}
	}
	return &total, nil
}

func rateCost(rates map[string]string, resource string, count int) (currency.Amount, error) {
	rate, ok := rates[strings.ToLower(resource)]
	if !ok {
		return currency.Amount{}, fmt.Errorf("pricebook has no rate for %q", resource)
	}
	amount, err := currency.NewAmount(rate, "USD")
	if err != nil {
		return currency.Amount{}, fmt.Errorf("parse %q rate %q: %w", resource, rate, err)
	}
	return amount.Mul(strconv.Itoa(count))
}

func hyperstackGPU(providerGPU string, count int) v1.GPU {
	gpuType := strings.TrimSuffix(providerGPU, "-spot")
	memoryGB := gpuMemoryGB(gpuType)
	memory, memoryBytes := byteSizes(memoryGB, v1.Gigabyte)
	return v1.GPU{
		Count:          int32(count),
		Memory:         memory,
		MemoryBytes:    memoryBytes,
		NetworkDetails: gpuNetworkDetails(gpuType),
		Manufacturer:   v1.ManufacturerNVIDIA,
		Name:           gpuName(gpuType),
		Type:           gpuType,
	}
}

func gpuMemoryGB(gpuType string) int64 {
	if matches := gpuMemoryPattern.FindStringSubmatch(gpuType); len(matches) == 2 {
		memoryGB, _ := strconv.ParseInt(matches[1], 10, 64)
		return memoryGB
	}
	knownMemory := map[string]int64{
		"B200":           180,
		"B300":           288,
		"L40":            48,
		"L40S":           48,
		"RTX-A4000":      16,
		"RTX-A6000":      48,
		"RTX-PRO6000-SE": 96,
	}
	upperType := strings.ToUpper(gpuType)
	for model, memoryGB := range knownMemory {
		if strings.Contains(upperType, model) {
			return memoryGB
		}
	}
	return 0
}

func gpuName(gpuType string) string {
	knownNames := []string{"RTX-PRO6000-SE", "RTX-A6000", "RTX-A4000", "B300", "B200", "H200", "H100", "A100", "L40S", "L40"}
	upperType := strings.ToUpper(gpuType)
	for _, name := range knownNames {
		if strings.Contains(upperType, name) {
			return name
		}
	}
	return gpuType
}

func gpuNetworkDetails(gpuType string) string {
	upperType := strings.ToUpper(gpuType)
	switch {
	case strings.Contains(upperType, "NVLINK"):
		return "NVLink"
	case strings.Contains(upperType, "PCIE"):
		return "PCIe"
	case strings.Contains(upperType, "SXM6"):
		return "SXM6"
	case strings.Contains(upperType, "SXM5"):
		return "SXM5"
	case strings.Contains(upperType, "SXM4"):
		return "SXM4"
	case strings.Contains(upperType, "SXM"):
		return "SXM"
	default:
		return ""
	}
}

func isSpotFlavor(typeName, gpuType string) bool {
	return strings.HasSuffix(strings.ToLower(typeName), "-spot") || strings.HasSuffix(strings.ToLower(gpuType), "-spot")
}

func byteSizes(value int64, unit v1.BytesUnit) (units.Base2Bytes, v1.Bytes) {
	size := v1.NewBytes(v1.BytesValue(value), unit)
	return units.Base2Bytes(size.ByteCount().Int64()), size
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func intValue(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

func float32Value(value *float32) float32 {
	if value == nil {
		return 0
	}
	return *value
}

func (c *HyperstackClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
