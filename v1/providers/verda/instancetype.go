package verda

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	v1 "github.com/brevdev/cloud/v1"
	verdago "github.com/verda-cloud/verdacloud-sdk-go/pkg/verda"
)

const defaultCurrency = "usd"

var (
	storageSizePattern     = regexp.MustCompile(`(?i)(\d+)\s*(TiB|TB|GiB|GB|MiB|MB)`)
	gpuMemorySuffixPattern = regexp.MustCompile(`(?i)\s+[0-9]+GB$`)
)

func (c *VerdaClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	verdaTypes, err := c.client.InstanceTypes.Get(ctx, defaultCurrency)
	if err != nil {
		return nil, wrapVerdaError(err)
	}
	availabilities, err := c.client.InstanceAvailability.GetAllAvailabilities(ctx, false, "")
	if err != nil {
		return nil, wrapVerdaError(err)
	}

	typeByName := make(map[string]verdago.InstanceTypeInfo, len(verdaTypes))
	for _, verdaType := range verdaTypes {
		typeByName[verdaType.InstanceType] = verdaType
	}

	instanceTypes := make([]v1.InstanceType, 0)
	for _, availability := range availabilities {
		for _, typeName := range availability.Availabilities {
			if len(args.InstanceTypes) > 0 && !slices.Contains(args.InstanceTypes, typeName) {
				continue
			}
			verdaType, ok := typeByName[typeName]
			if !ok {
				continue
			}
			instanceType, err := verdaInstanceTypeToInstanceType(verdaType, availability.LocationCode)
			if err != nil {
				return nil, err
			}
			if v1.IsSelectedByArgs(instanceType, args) {
				instanceTypes = append(instanceTypes, instanceType)
			}
		}
	}

	sort.Slice(instanceTypes, func(i, j int) bool {
		return instanceTypes[i].ID < instanceTypes[j].ID
	})
	return instanceTypes, nil
}

func verdaInstanceTypeToInstanceType(verdaType verdago.InstanceTypeInfo, location string) (v1.InstanceType, error) {
	basePrice, err := currency.NewAmount(
		strconv.FormatFloat(verdaType.PricePerHour.Float64(), 'f', -1, 64),
		currencyCode(verdaType.Currency),
	)
	if err != nil {
		return v1.InstanceType{}, fmt.Errorf("failed to parse price for verda instance type %s: %w", verdaType.InstanceType, err)
	}

	usageClasses := []string{"on-demand"}
	if verdaType.SpotPrice > 0 {
		usageClasses = append(usageClasses, "spot")
	}

	memory, memoryBytes := byteSizes(int64(verdaType.Memory.SizeInGigabytes), v1.Gigabyte)
	instanceType := v1.InstanceType{
		Type:                   verdaType.InstanceType,
		Location:               location,
		Memory:                 memory,
		MemoryBytes:            memoryBytes,
		VCPU:                   int32(verdaType.CPU.NumberOfCores), //nolint:gosec // ok
		SupportedArchitectures: []v1.Architecture{verdaArchitecture(verdaType.Model)},
		SupportedStorage:       storageDescriptionToStorage(verdaType.Storage.Description),
		ElasticRootVolume:      true,
		SupportedUsageClasses:  usageClasses,
		Stoppable:              true,
		Rebootable:             false,
		IsAvailable:            true,
		BasePrice:              &basePrice,
		Provider:               CloudProviderID,
		Cloud:                  CloudProviderID,
	}

	if verdaType.GPU.NumberOfGPUs > 0 {
		gpuMemory, gpuMemoryBytes := byteSizes(int64(verdaType.GPUMemory.SizeInGigabytes), v1.Gigabyte)
		gpuModel := strings.ToUpper(strings.TrimSpace(verdaType.Model))
		instanceType.SupportedGPUs = []v1.GPU{{
			Count:          int32(verdaType.GPU.NumberOfGPUs), //nolint:gosec // ok
			Memory:         gpuMemory,
			MemoryBytes:    gpuMemoryBytes,
			NetworkDetails: verdaType.P2P,
			Manufacturer:   v1.GetManufacturer(verdaType.Manufacturer),
			Name:           verdaGPUName(gpuModel),
			Type:           gpuModel,
		}}
	}

	instanceType.ID = v1.MakeGenericInstanceTypeID(instanceType)
	return instanceType, nil
}

func verdaGPUName(model string) string {
	return gpuMemorySuffixPattern.ReplaceAllString(model, "")
}

func verdaArchitecture(model string) v1.Architecture {
	// The Verda API does not expose architecture
	if strings.EqualFold(strings.TrimSpace(model), "GB300") {
		return v1.ArchitectureARM64
	}
	return v1.ArchitectureX86_64
}

func currencyCode(code string) string {
	code = strings.ToUpper(code)
	if len(code) == 3 {
		return code
	}
	return "USD"
}

func storageDescriptionToStorage(description string) []v1.Storage {
	match := storageSizePattern.FindStringSubmatch(description)
	if len(match) != 3 {
		return nil
	}

	size, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return nil
	}
	byteUnit, ok := storageByteUnit(match[2])
	if !ok {
		return nil
	}
	legacySize, sizeBytes := byteSizes(size, byteUnit)

	storageType := strings.TrimSpace(description)
	upperDescription := strings.ToUpper(description)
	switch {
	case strings.Contains(upperDescription, "NVME"):
		storageType = "NVMe"
	case strings.Contains(upperDescription, "SSD"):
		storageType = "SSD"
	case strings.Contains(upperDescription, "HDD"):
		storageType = "HDD"
	}

	return []v1.Storage{{
		Count:     1,
		Size:      legacySize,
		SizeBytes: sizeBytes,
		Type:      storageType,
	}}
}

func storageByteUnit(unit string) (v1.BytesUnit, bool) {
	switch strings.ToUpper(unit) {
	case "MB":
		return v1.Megabyte, true
	case "MIB":
		return v1.Mebibyte, true
	case "GB":
		return v1.Gigabyte, true
	case "GIB":
		return v1.Gibibyte, true
	case "TB":
		return v1.Terabyte, true
	case "TIB":
		return v1.Tebibyte, true
	default:
		return v1.BytesUnit{}, false
	}
}

func byteSizes(value int64, unit v1.BytesUnit) (units.Base2Bytes, v1.Bytes) {
	size := v1.NewBytes(v1.BytesValue(value), unit)
	return units.Base2Bytes(size.ByteCount().Int64()), size
}

func (c *VerdaClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
