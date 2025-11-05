package v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
)

const (
	brevDGXCWorkshopID = "brev-dgxc"
)

func (c *LaunchpadClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	launchpadInstanceTypes, err := c.paginateInstanceTypes(ctx, 100)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	instanceTypes := []v1.InstanceType{}
	for _, launchpadInstanceType := range launchpadInstanceTypes {
		for region, capacity := range launchpadInstanceType.Capacity {
			// If capacity is 0, the instance type is not available in the region
			if capacity == 0 {
				continue
			}

			// Convert the Launchpad instance type to a v1 instance type
			instanceType, err := launchpadInstanceTypeToInstanceType(launchpadInstanceType, region)
			if err != nil {
				return nil, errors.WrapAndTrace(err)
			}

			// Collect the instance type if it is selected by the args
			if isSelectedByArgs(*instanceType, args) {
				instanceTypes = append(instanceTypes, *instanceType)
			} else {
				continue
			}
		}
	}

	return instanceTypes, nil
}

func isSelectedByArgs(instanceType v1.InstanceType, args v1.GetInstanceTypeArgs) bool {
	if args.Locations != nil {
		for _, location := range instanceType.Location {
			if !args.Locations.IsAllowed(string(location)) {
				return false
			}
		}
	}

	if args.GPUManufactererFilter != nil {
		for _, supportedGPU := range instanceType.SupportedGPUs {
			if !args.GPUManufactererFilter.IsAllowed(supportedGPU.Manufacturer) {
				return false
			}
		}
	}

	if args.CloudFilter != nil {
		if !args.CloudFilter.IsAllowed(instanceType.Cloud) {
			return false
		}
	}

	if args.ArchitectureFilter != nil {
		for _, architecture := range instanceType.SupportedArchitectures {
			if !args.ArchitectureFilter.IsAllowed(architecture) {
				return false
			}
		}
	}

	return true
}

func (c *LaunchpadClient) paginateInstanceTypes(ctx context.Context, pageSize int32) ([]openapi.InstanceType, error) {
	instanceTypes := make([]openapi.InstanceType, 0, pageSize)
	var page int32 = 1
	for {
		// Fetch page
		instanceTypesRes, resp, err := c.client.CatalogInstanceTypesAPI.V1CatalogInstanceTypesList(c.makeAuthContext(ctx)).
			Delivery(string(openapi.OnDemandSpeedFast)).
			PageSize(pageSize).
			Page(page).
			Execute()
		defer func() {
			err = resp.Body.Close()
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "failed to close response body"))
			}
		}()
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}

		// Append results to list
		instanceTypes = append(instanceTypes, instanceTypesRes.Results...)

		// Check if there are more pages
		if len(instanceTypesRes.Results) < int(pageSize) {
			break
		}

		// Increment page
		page++
	}
	return instanceTypes, nil
}

type instanceTypeInfo struct {
	cloud             string
	gpuName           string
	gpuCount          int32
	gpuNetworkDetails string
	workshopID        string
}

func makeInstanceTypeName(info instanceTypeInfo) string {
	var workshopSuffix string
	if info.workshopID == "" {
		workshopSuffix = ""
	} else {
		workshopSuffix = "." + info.workshopID
	}
	return fmt.Sprintf("%s.%sx%d.%s%s", info.cloud, strings.ToLower(info.gpuName), info.gpuCount, strings.ToLower(info.gpuNetworkDetails), workshopSuffix)
}

func getInstanceTypeInfo(name string) (instanceTypeInfo, error) {
	parts := strings.Split(name, ".")
	if len(parts) != 3 && len(parts) != 4 {
		return instanceTypeInfo{}, errors.Errorf("invalid instance type name: %s", name)
	}
	cloud := parts[0]

	// Parse gpuName and gpuCount from "gpuNamexCOUNT" format
	gpuPart := parts[1]
	xIndex := strings.LastIndex(gpuPart, "x")
	if xIndex == -1 {
		return instanceTypeInfo{}, errors.New("invalid GPU part format")
	}
	gpuName := gpuPart[:xIndex]
	gpuCountStr := gpuPart[xIndex+1:]
	gpuCount, err := strconv.ParseInt(gpuCountStr, 10, 32)
	if err != nil {
		return instanceTypeInfo{}, errors.WrapAndTrace(err)
	}

	gpuNetworkDetails := parts[2]

	var workshopID string
	if len(parts) == 4 {
		workshopID = parts[3]
	} else {
		workshopID = ""
	}

	return instanceTypeInfo{
		cloud:             cloud,
		gpuName:           gpuName,
		gpuCount:          int32(gpuCount),
		gpuNetworkDetails: gpuNetworkDetails,
		workshopID:        workshopID,
	}, nil
}

func launchpadInstanceTypeToInstanceType(launchpadInstanceType openapi.InstanceType, region string) (*v1.InstanceType, error) {
	// Launchpad may return multiple storage options
	storage := launchpadStorageToStorages(launchpadInstanceType.Storage)

	// Launchpad returns only a single GPU
	gpus := launchpadGpusToGpus([]openapi.InstanceTypeGpu{launchpadInstanceType.Gpu})
	if len(gpus) != 1 {
		return nil, errors.Errorf("expected 1 GPU, got %d", len(gpus))
	}
	gpu := gpus[0]

	// The fact that we are here means there is capacity in the region, and therefore the instance type is available
	isAvailable := true

	// Get the hourly price
	hourlyPrice := getInstanceTypePrice(&gpu)

	info := instanceTypeInfo{
		cloud:             launchpadInstanceType.Cloud,
		gpuName:           gpu.Name,
		gpuCount:          gpu.Count,
		gpuNetworkDetails: gpu.NetworkDetails,
		workshopID:        launchpadInstanceType.WorkshopId,
	}
	typeName := makeInstanceTypeName(info)

	it := &v1.InstanceType{
		Type:                   typeName,
		VCPU:                   launchpadInstanceType.Cpu,
		Memory:                 gbToBytes(launchpadInstanceType.MemoryGb),
		MemoryByteValue:        v1.NewByteValue(launchpadInstanceType.MemoryGb, v1.Gigabyte),
		SupportedGPUs:          []v1.GPU{gpu},
		SupportedStorage:       storage,
		SupportedArchitectures: []v1.Architecture{launchpadArchitectureToArchitecture(launchpadInstanceType.SystemArch)},
		IsAvailable:            isAvailable,
		Location:               region,
		BasePrice:              hourlyPrice,
		Provider:               CloudProviderID,
		Cloud:                  launchpadInstanceType.Cloud,
		ReservedInstancePoolID: launchpadWorkshopIDToReservedInstancePoolID(info.workshopID),
	}

	// Make the instance type ID
	it.ID = v1.MakeGenericInstanceTypeID(*it)
	return it, nil
}

func launchpadWorkshopIDToReservedInstancePoolID(workshopID string) string {
	// If the workshop is empty, consider this to be a public instance type
	if workshopID == "" {
		return ""
	}

	// If the workshop is the Brev DGXC workshop, consider this to be a public instance type
	if workshopID == brevDGXCWorkshopID {
		return ""
	}

	// Otherwise, use the workshop ID as the reserved instance pool ID
	return workshopID
}

func launchpadStorageToStorages(launchpadStorage []openapi.InstanceTypeStorage) []v1.Storage {
	if len(launchpadStorage) == 0 {
		return nil
	}
	storage := make([]v1.Storage, len(launchpadStorage))
	for i, s := range launchpadStorage {
		storage[i] = v1.Storage{
			Count:         1,
			Size:          gbToBytes(s.SizeGb),
			SizeByteValue: v1.NewByteValue(s.SizeGb, v1.Gigabyte),
			Type:          string(s.Type),
		}
	}
	return storage
}

func launchpadGpusToGpus(lpGpus []openapi.InstanceTypeGpu) []v1.GPU {
	if len(lpGpus) == 0 {
		return nil
	}
	gpus := make([]v1.GPU, len(lpGpus))
	for i, gp := range lpGpus {
		gpus[i] = v1.GPU{
			Name:           strings.ToUpper(gp.Family),
			Manufacturer:   v1.GetManufacturer(gp.Manufacturer),
			Count:          gp.Count,
			Memory:         gbToBytes(gp.MemoryGb),
			NetworkDetails: string(gp.InterconnectionType),
			Type:           strings.ToUpper(gp.Model),
		}
	}
	return gpus
}

func launchpadArchitectureToArchitecture(launchpadArchitecture openapi.SystemArchEnum) v1.Architecture {
	switch launchpadArchitecture {
	case openapi.SystemArchAMD64:
		return v1.ArchitectureX86_64
	case openapi.SystemArchARM64:
		return v1.ArchitectureARM64
	}
	return v1.ArchitectureUnknown
}

// TODO: this will convert 1GB to 1GiB which is incorrect
func gbToBytes(gb int32) units.Base2Bytes {
	gbInt64 := int64(gb)
	mult := 1024 * 1024 * 1024
	return units.Base2Bytes((gbInt64 * int64(mult)))
}

func launchpadClusterToInstanceType(cluster openapi.Cluster) *v1.InstanceType {
	// Technically launchpad clusters can have multiple nodes, but we only support one for now, and should expect only one will be returned.
	if len(cluster.Nodes) == 0 || cluster.Nodes[0].Node == nil || cluster.Nodes[0].Node.GpuCount == nil {
		return nil
	}
	node := *cluster.Nodes[0].Node

	storage := launchpadNodeToSupportedStorage(node)
	if storage == nil {
		return nil
	}

	gpu := launchpadGputoGpu(node)
	if gpu == nil {
		return nil
	}

	var vcpu int32
	if node.Cpu != nil {
		vcpu = *node.Cpu
	}
	var memory units.Base2Bytes
	var memoryValueBytes v1.ByteValue
	if node.Memory != nil {
		memory = gbToBytes(*node.Memory)
		memoryValueBytes = v1.NewByteValue(*node.Memory, v1.Gigabyte)
	}

	isAvailable := (cluster.ProvisioningState != nil && *cluster.ProvisioningState == openapi.ProvisioningStateReady)
	location := node.GetLocation().Location.GetName()
	cloud := node.GetLocation().Location.GetProvider().Provider.GetName()

	typeName := makeInstanceTypeName(instanceTypeInfo{
		cloud:             cloud,
		gpuName:           gpu.Name,
		gpuCount:          *node.GpuCount,
		gpuNetworkDetails: gpu.NetworkDetails,
		workshopID:        cluster.GetWorkshopId(),
	})

	it := &v1.InstanceType{
		Type:             typeName,
		SupportedGPUs:    []v1.GPU{*gpu},
		SupportedStorage: storage,
		Memory:           memory,
		MemoryByteValue:  memoryValueBytes,
		VCPU:             vcpu,
		IsAvailable:      isAvailable,
		Location:         location,
		BasePrice:        getInstanceTypePrice(gpu),
		Cloud:            cloud,
		Provider:         CloudProviderID,
	}
	it.ID = v1.MakeGenericInstanceTypeID(*it)
	return it
}

func launchpadGputoGpu(node openapi.Node) *v1.GPU {
	lpGpu := node.GetGpu().Gpu
	if lpGpu == nil {
		return nil
	}

	lpGpuModel := strings.ToUpper(lpGpu.Model)
	lpGpuCount := node.GetGpuCount()

	var lpGpuFormFactor string
	if lpGpu.FormFactor != nil {
		switch *lpGpu.FormFactor {
		case openapi.InterconnectionTypePCIe:
			lpGpuFormFactor = "PCIE"
		case openapi.InterconnectionTypeSXM:
			lpGpuFormFactor = "SXM"
		}
	}

	var lpGpuMemory units.Base2Bytes
	if lpGpu.Memory != nil {
		lpGpuMemory = gbToBytes(*lpGpu.Memory)
	}

	gpu := &v1.GPU{
		Name:           lpGpuModel,
		Count:          lpGpuCount,
		NetworkDetails: lpGpuFormFactor,
		Memory:         lpGpuMemory,
		Manufacturer:   "NVIDIA", // The only supported manufacturer for Launchpad
	}
	return gpu
}

func launchpadNodeToSupportedStorage(node openapi.Node) []v1.Storage {
	if len(node.Storage) == 0 {
		return nil
	}
	storage := make([]v1.Storage, 0, len(node.Storage))
	for _, s := range node.Storage {
		storage = append(storage, v1.Storage{
			Count: 1,
			Size:  gbToBytes(*s.Size),
			Type:  launchpadStorageTypeToStorageType(s.Type),
		})
	}
	return storage
}

func launchpadStorageTypeToStorageType(storageType openapi.TypeEnum) string {
	switch storageType {
	case openapi.TypeNVMe:
		return "nvme"
	case openapi.TypeSSD:
		return "ssd"
	default:
		return ""
	}
}

var (
	defaultCurrencyCode = "USD"
	defaultPricePerGPU  = newCurrencyAmount("2.50", defaultCurrencyCode)
	gpuNameToPrice      = map[string]*currency.Amount{
		"h100":   newCurrencyAmount("2.37", defaultCurrencyCode),
		"h200":   newCurrencyAmount("2.80", defaultCurrencyCode),
		"a40":    newCurrencyAmount("0.98", defaultCurrencyCode),
		"ggl40s": newCurrencyAmount("1.98", defaultCurrencyCode),
		"gh200":  newCurrencyAmount("2.80", defaultCurrencyCode),
		"l40s":   newCurrencyAmount("1.58", defaultCurrencyCode),
		"a30":    newCurrencyAmount("0.12", defaultCurrencyCode),
		"a100":   newCurrencyAmount("2.01", defaultCurrencyCode),
	}
)

func newCurrencyAmount(amount string, currencyCode string) *currency.Amount {
	currencyAmount, err := currency.NewAmount(amount, currencyCode)
	if err != nil {
		panic(err)
	}
	return &currencyAmount
}

// uses gpuNameToPrice map to get the price per GPU
func getInstanceTypePrice(gpu *v1.GPU) *currency.Amount {
	if gpu == nil {
		return nil
	}

	gpuNameLower := strings.ToLower(gpu.Name)

	// get price per GPU from gpuNameToPrice map, defaulting if not found
	pricePerGPU := defaultPricePerGPU
	for gpu, price := range gpuNameToPrice {
		if gpuNameLower == gpu {
			pricePerGPU = price
			break
		}
	}

	// multiply price per GPU by gpuCount
	price, err := pricePerGPU.Mul(strconv.Itoa(int(gpu.Count)))
	if err != nil {
		panic(err)
	}

	// calculate upcharge
	upcharge := "1.0"
	if gpu.NetworkDetails == "SXM" {
		// Apply 25% upcharge for SXM variants
		upcharge = "1.25"
	}

	// apply upcharge
	price, err = price.Mul(upcharge)
	if err != nil {
		panic(err)
	}

	// round to 2 decimal places
	price = price.RoundTo(2, currency.RoundHalfUp)

	return &price
}
