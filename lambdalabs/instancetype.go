package lambdalabs

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/brevdev/cloud"
	"github.com/brevdev/cloud/internal/collections"
	openapi "github.com/brevdev/cloud/internal/gen/lambdalabs"
)

// GetInstanceTypePollTime returns the polling interval for instance types
func (c *LambdaLabsClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *LambdaLabsClient) GetInstanceTypes(ctx context.Context, args cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	instanceTypesResp, err := c.getInstanceTypes(ctx)
	if err != nil {
		return nil, err
	}

	locations, err := c.GetLocations(ctx, cloud.GetLocationsArgs{})
	if err != nil {
		return nil, err
	}

	instanceTypes, err := collections.MapE(collections.GetMapValues(instanceTypesResp.Data), func(resp openapi.InstanceTypes200ResponseDataValue) ([]cloud.InstanceType, error) {
		currentlyAvailableRegions := collections.GroupBy(resp.RegionsWithCapacityAvailable, func(lambdaRegion openapi.Region) string {
			return lambdaRegion.Name
		})
		its, err1 := collections.MapE(locations, func(region cloud.Location) (cloud.InstanceType, error) {
			isAvailable := false
			if _, ok := currentlyAvailableRegions[region.Name]; ok {
				isAvailable = true
			}
			it, err2 := convertLambdaLabsInstanceTypeToV1InstanceType(region.Name, resp.InstanceType, isAvailable)
			if err2 != nil {
				return cloud.InstanceType{}, err2
			}
			return it, nil
		})
		if err1 != nil {
			return []cloud.InstanceType{}, err1
		}
		return its, nil
	})
	if err != nil {
		return nil, err
	}
	instanceTypesFlattened := collections.Flatten(instanceTypes)

	if len(args.Locations) == 0 {
		if c.location != "" {
			args.Locations = []string{c.location}
		} else {
			args.Locations = cloud.All
		}
	}

	if !args.Locations.IsAll() {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(it cloud.InstanceType) bool {
			return collections.ListContains(args.Locations, it.Location)
		})
	}

	if len(args.SupportedArchitectures) > 0 {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(instanceType cloud.InstanceType) bool {
			for _, arch := range args.SupportedArchitectures {
				if collections.ListContains(instanceType.SupportedArchitectures, arch) {
					return true
				}
			}
			return false
		})
	}
	if len(args.InstanceTypes) > 0 {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(instanceType cloud.InstanceType) bool {
			return collections.ListContains(args.InstanceTypes, instanceType.Type)
		})
	}

	return instanceTypesFlattened, nil
}

func (c *LambdaLabsClient) getInstanceTypes(ctx context.Context) (*openapi.InstanceTypes200Response, error) {
	ilr, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.InstanceTypes200Response, error) {
		res, resp, err := c.client.DefaultAPI.InstanceTypes(c.makeAuthContext(ctx)).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.InstanceTypes200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}

	return ilr, nil
}

func parseGPUFromDescription(input string) (cloud.GPU, error) {
	var gpu cloud.GPU

	// Extract the count
	countRegex := regexp.MustCompile(`(\d+)x`)
	countMatch := countRegex.FindStringSubmatch(input)
	if len(countMatch) == 0 {
		return cloud.GPU{}, fmt.Errorf("could not find count in %s", input)
	}
	count, _ := strconv.ParseInt(countMatch[1], 10, 32)
	gpu.Count = int32(count)

	// Extract the memory
	memoryRegex := regexp.MustCompile(`(\d+) GB`)
	memoryMatch := memoryRegex.FindStringSubmatch(input)
	if len(memoryMatch) == 0 {
		return cloud.GPU{}, fmt.Errorf("could not find memory in %s", input)
	}
	memoryStr := memoryMatch[1]
	memoryGiB, _ := strconv.Atoi(memoryStr)
	gpu.Memory = units.GiB * units.Base2Bytes(memoryGiB)

	// Extract the network details
	networkRegex := regexp.MustCompile(`(\w+\s?)+\)`)
	networkMatch := networkRegex.FindStringSubmatch(input)
	if len(networkMatch) == 0 {
		return cloud.GPU{}, fmt.Errorf("could not find network details in %s", input)
	}
	networkStr := strings.TrimSuffix(networkMatch[0], ")")
	networkDetails := strings.TrimSpace(strings.ReplaceAll(networkStr, memoryStr+" GB", ""))
	gpu.NetworkDetails = networkDetails

	// Extract the name
	nameRegex := regexp.MustCompile(`x (.*?) \(`)
	nameMatch := nameRegex.FindStringSubmatch(input)
	if len(nameMatch) == 0 {
		return cloud.GPU{}, fmt.Errorf("could not find name in %s", input)
	}
	nameStr := strings.TrimRight(strings.TrimLeft(nameMatch[0], "x "), " (")
	nameStr = regexp.MustCompile(`(?i)^Tesla\s+`).ReplaceAllString(nameStr, "")
	gpu.Name = nameStr
	if networkDetails != "" {
		gpu.Type = nameStr + "." + networkDetails
	} else {
		gpu.Type = nameStr
	}

	gpu.Manufacturer = "NVIDIA"

	return gpu, nil
}

func convertLambdaLabsInstanceTypeToV1InstanceType(location string, instType openapi.InstanceType, isAvailable bool) (cloud.InstanceType, error) {
	gpus := []cloud.GPU{}
	if !strings.Contains(instType.Description, "CPU") {
		gpu, err := parseGPUFromDescription(instType.Description)
		if err != nil {
			return cloud.InstanceType{}, err
		}
		gpus = append(gpus, gpu)
	}
	amount, err := currency.NewAmountFromInt64(int64(instType.PriceCentsPerHour), "USD")
	if err != nil {
		return cloud.InstanceType{}, err
	}
	it := cloud.InstanceType{
		Location:      location,
		Type:          instType.Name,
		SupportedGPUs: gpus,
		SupportedStorage: []cloud.Storage{
			{
				Type:  "ssd",
				Count: 1,
				Size:  units.GiB * units.Base2Bytes(instType.Specs.StorageGib),
			},
		},
		SupportedUsageClasses:    []string{"on-demand"},
		Memory:                   units.GiB * units.Base2Bytes(instType.Specs.MemoryGib),
		MaximumNetworkInterfaces: 0,
		NetworkPerformance:       "",
		SupportedNumCores:        []int32{},
		DefaultCores:             0,
		VCPU:                     instType.Specs.Vcpus,
		SupportedArchitectures:   []string{"x86_64"},
		ClockSpeedInGhz:          0,
		Stoppable:                false,
		Rebootable:               true,
		IsAvailable:              isAvailable,
		BasePrice:                &amount,
		Provider:                 string(CloudProviderID),
	}
	it.ID = cloud.MakeGenericInstanceTypeID(it)
	return it, nil
}
