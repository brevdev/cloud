package v1

import (
	"context"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	cloud "github.com/brevdev/cloud/v1"
)

type serverProductListResponse struct {
	Response productList `json:"getServerProductListResponse"`
}

func (r *serverProductListResponse) apiError() error {
	return r.Response.responseMeta.apiError()
}

type productList struct {
	responseMeta
	TotalRows   int            `json:"totalRows"`
	ProductList []naverProduct `json:"productList"`
}

type naverProduct struct {
	ProductCode          string   `json:"productCode"`
	ProductName          string   `json:"productName"`
	ProductType          codeName `json:"productType"`
	ProductDescription   string   `json:"productDescription"`
	InfraResourceType    codeName `json:"infraResourceType"`
	CPUCount             int32    `json:"cpuCount"`
	MemorySize           int64    `json:"memorySize"`
	BaseBlockStorageSize int64    `json:"baseBlockStorageSize"`
	OSInformation        string   `json:"osInformation"`
	DiskType             codeName `json:"diskType"`
	PlatformType         codeName `json:"platformType"`
	GenerationCode       string   `json:"generationCode"`
}

func (c *NaverClient) GetInstanceTypePollTime() time.Duration {
	return defaultInstanceTypePollMinutes * time.Minute
}

func (c *NaverClient) GetInstanceTypes(ctx context.Context, args cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	locations := args.Locations
	if len(locations) == 0 {
		locations = cloud.LocationsFilter{c.location}
	}
	if locations.IsAll() {
		locs, err := c.GetLocations(ctx, cloud.GetLocationsArgs{})
		if err != nil {
			return nil, err
		}
		locations = make(cloud.LocationsFilter, 0, len(locs))
		for _, loc := range locs {
			locations = append(locations, loc.Name)
		}
	}

	var out []cloud.InstanceType
	for _, location := range locations {
		params := url.Values{}
		params.Set("regionCode", location)
		params.Set("serverImageProductCode", defaultServerImageProductCode)
		var resp serverProductListResponse
		if err := c.do(ctx, "getServerProductList", params, &resp); err != nil {
			return nil, err
		}
		for _, product := range resp.Response.ProductList {
			it := product.toInstanceType(location)
			if len(args.InstanceTypes) > 0 && !slices.Contains(args.InstanceTypes, it.Type) {
				continue
			}
			if args.CloudFilter != nil && !args.CloudFilter.IsAllowed(it.Cloud) {
				continue
			}
			if args.ArchitectureFilter != nil && !args.ArchitectureFilter.IsAllowed(cloud.ArchitectureX86_64) {
				continue
			}
			if args.GPUManufactererFilter != nil {
				allowed := len(it.SupportedGPUs) == 0
				for _, gpu := range it.SupportedGPUs {
					if args.GPUManufactererFilter.IsAllowed(gpu.Manufacturer) {
						allowed = true
					}
				}
				if !allowed {
					continue
				}
			}
			out = append(out, it)
		}
	}
	return out, nil
}

func (p naverProduct) toInstanceType(location string) cloud.InstanceType {
	storage := cloud.Storage{
		Type:      firstNonEmpty(p.DiskType.Code, "NET"),
		Count:     1,
		Size:      units.Base2Bytes(p.BaseBlockStorageSize),
		SizeBytes: cloud.NewBytes(cloud.BytesValue(p.BaseBlockStorageSize), cloud.Byte),
	}
	it := cloud.InstanceType{
		Location:                 location,
		Type:                     p.ProductCode,
		SupportedGPUs:            parseGPU(p),
		SupportedStorage:         []cloud.Storage{storage},
		SupportedUsageClasses:    []string{"on-demand"},
		Memory:                   units.Base2Bytes(p.MemorySize),
		MemoryBytes:              cloud.NewBytes(cloud.BytesValue(p.MemorySize), cloud.Byte),
		MaximumNetworkInterfaces: 3,
		NetworkPerformance:       "",
		VCPU:                     p.CPUCount,
		SupportedArchitectures:   []cloud.Architecture{cloud.ArchitectureX86_64},
		Stoppable:                true,
		Rebootable:               true,
		IsAvailable:              true,
		Provider:                 CloudProviderID,
		Cloud:                    CloudProviderID,
	}
	it.ID = cloud.MakeGenericInstanceTypeID(it)
	return it
}

func parseGPU(product naverProduct) []cloud.GPU {
	text := strings.Join([]string{product.ProductCode, product.ProductName, product.ProductDescription, product.ProductType.Code, product.ProductType.CodeName}, " ")
	if !strings.Contains(strings.ToUpper(text), "GPU") && !containsKnownGPU(text) {
		return nil
	}

	name := "GPU"
	count := int32(1)
	gpuRe := regexp.MustCompile(`(?i)(?:NVIDIA\s+)?(H200|H100|A100|A10|L40S|V100|T4|RTX[0-9A-Z]+)\s*(?:N\d+)?\s*(\d+)?\s*EA?`)
	if match := gpuRe.FindStringSubmatch(text); len(match) > 0 {
		name = strings.ToUpper(match[1])
		if len(match) > 2 && match[2] != "" {
			if parsed, err := strconv.ParseInt(match[2], 10, 32); err == nil && parsed > 0 {
				count = int32(parsed)
			}
		}
	}

	return []cloud.GPU{{
		Count:        count,
		Name:         name,
		Type:         name,
		Manufacturer: cloud.ManufacturerNVIDIA,
	}}
}

func containsKnownGPU(text string) bool {
	upper := strings.ToUpper(text)
	for _, token := range []string{"H200", "H100", "A100", "A10", "L40S", "V100", "T4"} {
		if strings.Contains(upper, token) {
			return true
		}
	}
	return false
}
