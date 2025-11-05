package v1

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/brevdev/cloud/internal/collections"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/lambdalabs/gen/lambdalabs"
)

// GetInstanceTypePollTime returns the polling interval for instance types
func (c *LambdaLabsClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *LambdaLabsClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	instanceTypesResp, err := c.getInstanceTypes(ctx)
	if err != nil {
		return nil, err
	}

	locations, err := c.GetLocations(ctx, v1.GetLocationsArgs{})
	if err != nil {
		return nil, err
	}

	instanceTypes, err := collections.MapE(collections.GetMapValues(instanceTypesResp.Data), func(resp openapi.InstanceTypes200ResponseDataValue) ([]v1.InstanceType, error) {
		currentlyAvailableRegions := collections.GroupBy(resp.RegionsWithCapacityAvailable, func(lambdaRegion openapi.Region) string {
			return lambdaRegion.Name
		})
		its, err1 := collections.MapE(locations, func(region v1.Location) (v1.InstanceType, error) {
			isAvailable := false
			if _, ok := currentlyAvailableRegions[region.Name]; ok {
				isAvailable = true
			}
			it, err2 := convertLambdaLabsInstanceTypeToV1InstanceType(region.Name, resp.InstanceType, isAvailable)
			if err2 != nil {
				return v1.InstanceType{}, err2
			}
			return it, nil
		})
		if err1 != nil {
			return []v1.InstanceType{}, err1
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
			args.Locations = v1.All
		}
	}

	if !args.Locations.IsAll() {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(it v1.InstanceType) bool {
			return collections.ListContains(args.Locations, it.Location)
		})
	}

	if args.ArchitectureFilter != nil {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(instanceType v1.InstanceType) bool {
			for _, arch := range instanceType.SupportedArchitectures {
				if args.ArchitectureFilter.IsAllowed(arch) {
					return true
				}
			}
			return false
		})
	}
	if len(args.InstanceTypes) > 0 {
		instanceTypesFlattened = collections.Filter(instanceTypesFlattened, func(instanceType v1.InstanceType) bool {
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

func parseGPUFromDescription(input string) (v1.GPU, error) {
	var gpu v1.GPU

	// Extract the count
	countRegex := regexp.MustCompile(`(\d+)x`)
	countMatch := countRegex.FindStringSubmatch(input)
	if len(countMatch) == 0 {
		return v1.GPU{}, fmt.Errorf("could not find count in %s", input)
	}
	count, _ := strconv.ParseInt(countMatch[1], 10, 32)
	gpu.Count = int32(count)

	// Extract the memory
	memoryRegex := regexp.MustCompile(`(\d+) GB`)
	memoryMatch := memoryRegex.FindStringSubmatch(input)
	if len(memoryMatch) == 0 {
		return v1.GPU{}, fmt.Errorf("could not find memory in %s", input)
	}
	memoryStr := memoryMatch[1]
	memoryGiB, _ := strconv.Atoi(memoryStr)
	gpu.Memory = units.GiB * units.Base2Bytes(memoryGiB)
	gpu.MemoryByteValue = v1.NewByteValue(int32(memoryGiB), v1.Gibibyte)

	// Extract the network details
	networkRegex := regexp.MustCompile(`(\w+\s?)+\)`)
	networkMatch := networkRegex.FindStringSubmatch(input)
	if len(networkMatch) == 0 {
		return v1.GPU{}, fmt.Errorf("could not find network details in %s", input)
	}
	networkStr := strings.TrimSuffix(networkMatch[0], ")")
	networkDetails := strings.TrimSpace(strings.ReplaceAll(networkStr, memoryStr+" GB", ""))
	gpu.NetworkDetails = networkDetails

	// Extract the name
	nameRegex := regexp.MustCompile(`x (.*?) \(`)
	nameMatch := nameRegex.FindStringSubmatch(input)
	if len(nameMatch) == 0 {
		return v1.GPU{}, fmt.Errorf("could not find name in %s", input)
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

func convertLambdaLabsInstanceTypeToV1InstanceType(location string, instType openapi.InstanceType, isAvailable bool) (v1.InstanceType, error) {
	gpus := []v1.GPU{}
	if !strings.Contains(instType.Description, "CPU") {
		gpu, err := parseGPUFromDescription(instType.Description)
		if err != nil {
			return v1.InstanceType{}, err
		}
		gpus = append(gpus, gpu)
	}
	amount, err := currency.NewAmountFromInt64(int64(instType.PriceCentsPerHour), "USD")
	if err != nil {
		return v1.InstanceType{}, err
	}

	it := v1.InstanceType{
		Location:      location,
		Type:          instType.Name,
		SupportedGPUs: gpus,
		SupportedStorage: []v1.Storage{
			{
				Type:          "ssd",
				Count:         1,
				Size:          units.GiB * units.Base2Bytes(instType.Specs.StorageGib),
				SizeByteValue: v1.NewByteValue(instType.Specs.StorageGib, v1.Gibibyte),
			},
		},
		SupportedUsageClasses:    []string{"on-demand"},
		Memory:                   units.GiB * units.Base2Bytes(instType.Specs.MemoryGib),
		MemoryByteValue:          v1.NewByteValue(instType.Specs.MemoryGib, v1.Gibibyte),
		MaximumNetworkInterfaces: 0,
		NetworkPerformance:       "",
		SupportedNumCores:        []int32{},
		DefaultCores:             0,
		VCPU:                     instType.Specs.Vcpus,
		SupportedArchitectures:   []v1.Architecture{v1.ArchitectureX86_64},
		ClockSpeedInGhz:          0,
		Stoppable:                false,
		Rebootable:               true,
		IsAvailable:              isAvailable,
		BasePrice:                &amount,
		Provider:                 CloudProviderID,
		Cloud:                    CloudProviderID,
	}
	it.ID = v1.MakeGenericInstanceTypeID(it)
	return it, nil
}
