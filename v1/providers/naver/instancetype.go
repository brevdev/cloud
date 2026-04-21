package v1

import (
	"context"
	"encoding/json"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	cloud "github.com/brevdev/cloud/v1"
)

const (
	defaultPayCurrencyCode     = "USD"
	ncloudServerProductKindVPC = "VSVR"
	ncloudMeterRatePriceType   = "MTRAT"
	ncloudHourlyUsageUnit      = "USAGE_HH"
)

type serverProductListResponse struct {
	Response productList `json:"getServerProductListResponse"`
}

func (r *serverProductListResponse) apiError() error {
	return r.Response.apiError()
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

type productPriceListResponse struct {
	Response naverProductPriceList `json:"getProductPriceListResponse"`
}

func (r *productPriceListResponse) apiError() error {
	return r.Response.apiError()
}

type naverProductPriceList struct {
	responseMeta
	TotalRows        int                 `json:"totalRows"`
	ProductPriceList []naverProductPrice `json:"productPriceList"`
}

type naverProductPrice struct {
	ProductCode string       `json:"productCode"`
	PriceList   []naverPrice `json:"priceList"`
}

type naverPrice struct {
	PriceType   codeName         `json:"priceType"`
	Unit        codeName         `json:"unit"`
	Price       naverPriceAmount `json:"price"`
	PayCurrency naverPayCurrency `json:"payCurrency"`
}

type naverPayCurrency struct {
	Code string `json:"code"`
}

type naverPriceAmount string

func (a *naverPriceAmount) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err == nil {
		*a = naverPriceAmount(value)
		return nil
	}

	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}
	*a = naverPriceAmount(number.String())
	return nil
}

func (c *NaverClient) GetInstanceTypePollTime() time.Duration {
	return defaultInstanceTypePollMinutes * time.Minute
}

func (c *NaverClient) GetInstanceTypes(ctx context.Context, args cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	locations, err := c.instanceTypeLocations(ctx, args.Locations)
	if err != nil {
		return nil, err
	}

	var out []cloud.InstanceType
	for _, location := range locations {
		instanceTypes, err := c.instanceTypesForLocation(ctx, location, args)
		if err != nil {
			return nil, err
		}
		out = append(out, instanceTypes...)
	}
	return out, nil
}

func (c *NaverClient) instanceTypeLocations(ctx context.Context, locations cloud.LocationsFilter) (cloud.LocationsFilter, error) {
	if len(locations) == 0 {
		return cloud.LocationsFilter{c.location}, nil
	}
	if !locations.IsAll() {
		return locations, nil
	}

	locs, err := c.GetLocations(ctx, cloud.GetLocationsArgs{})
	if err != nil {
		return nil, err
	}
	resolved := make(cloud.LocationsFilter, 0, len(locs))
	for _, loc := range locs {
		resolved = append(resolved, loc.Name)
	}
	return resolved, nil
}

func (c *NaverClient) instanceTypesForLocation(ctx context.Context, location string, args cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	params := url.Values{}
	params.Set("regionCode", location)
	params.Set("serverImageProductCode", defaultServerImageProductCode)

	var resp serverProductListResponse
	if err := c.do(ctx, "getServerProductList", params, &resp); err != nil {
		return nil, err
	}

	prices := c.productPrices(ctx, location)
	out := make([]cloud.InstanceType, 0, len(resp.Response.ProductList))
	for _, product := range resp.Response.ProductList {
		it := product.toInstanceType(location, prices[product.ProductCode])
		if includeInstanceType(it, args) {
			out = append(out, it)
		}
	}
	return out, nil
}

func includeInstanceType(it cloud.InstanceType, args cloud.GetInstanceTypeArgs) bool {
	if len(args.InstanceTypes) > 0 && !slices.Contains(args.InstanceTypes, it.Type) {
		return false
	}
	if args.CloudFilter != nil && !args.CloudFilter.IsAllowed(it.Cloud) {
		return false
	}
	if args.ArchitectureFilter != nil && !args.ArchitectureFilter.IsAllowed(cloud.ArchitectureX86_64) {
		return false
	}
	return allowsGPUManufacturer(it.SupportedGPUs, args.GPUManufactererFilter)
}

func allowsGPUManufacturer(gpus []cloud.GPU, filter *cloud.GPUManufacturerFilter) bool {
	if filter == nil || len(gpus) == 0 {
		return true
	}
	for _, gpu := range gpus {
		if filter.IsAllowed(gpu.Manufacturer) {
			return true
		}
	}
	return false
}

func (c *NaverClient) productPrices(ctx context.Context, location string) map[string]*currency.Amount {
	params := url.Values{}
	params.Set("regionCode", location)
	params.Set("productItemKindCode", ncloudServerProductKindVPC)
	params.Set("payCurrencyCode", defaultPayCurrencyCode)
	params.Set("pageSize", "1000")

	var resp productPriceListResponse
	if err := c.doBilling(ctx, "product/getProductPriceList", params, &resp); err != nil {
		return nil
	}

	prices := make(map[string]*currency.Amount, len(resp.Response.ProductPriceList))
	for _, productPrice := range resp.Response.ProductPriceList {
		price := hourlyPrice(productPrice.PriceList)
		if price != nil {
			prices[productPrice.ProductCode] = price
		}
	}
	return prices
}

func hourlyPrice(prices []naverPrice) *currency.Amount {
	for _, price := range prices {
		if price.PriceType.Code != ncloudMeterRatePriceType || price.Unit.Code != ncloudHourlyUsageUnit {
			continue
		}

		amount, err := currency.NewAmount(string(price.Price), price.PayCurrency.Code)
		if err != nil {
			return nil
		}
		return &amount
	}
	return nil
}

func (p naverProduct) toInstanceType(location string, basePrice *currency.Amount) cloud.InstanceType {
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
		BasePrice:                basePrice,
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
