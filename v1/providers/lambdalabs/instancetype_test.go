package v1

import (
	"context"
	"testing"
	"time"

	"github.com/alecthomas/units"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/lambdalabs/gen/lambdalabs"
)

func TestLambdaLabsClient_GetInstanceTypes_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockResponse := createMockInstanceTypeResponse()
	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, mockResponse))

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	locations, err := getLambdaLabsLocations()
	require.NoError(t, err)
	assert.Len(t, instanceTypes, len(locations)*2)

	a10Type := findInstanceTypeByName(instanceTypes, "gpu_1x_a10")
	require.NotNil(t, a10Type)
	assert.Equal(t, "gpu_1x_a10", a10Type.Type)
	assert.True(t, a10Type.IsAvailable)
	assert.Len(t, a10Type.SupportedGPUs, 1)
	assert.Equal(t, int32(1), a10Type.SupportedGPUs[0].Count)
	assert.Equal(t, "NVIDIA", a10Type.SupportedGPUs[0].Manufacturer)
	assert.Equal(t, "A10", a10Type.SupportedGPUs[0].Name)
}

func TestLambdaLabsClient_GetInstanceTypes_FilterByLocation(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockResponse := createMockInstanceTypeResponse()
	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, mockResponse))

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		Locations: v1.LocationsFilter{"us-west-1"},
	})
	require.NoError(t, err)
	assert.Len(t, instanceTypes, 2)

	for _, instanceType := range instanceTypes {
		assert.Equal(t, "us-west-1", instanceType.Location)
	}
}

func TestLambdaLabsClient_GetInstanceTypes_FilterByInstanceType(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockResponse := createMockInstanceTypeResponse()
	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, mockResponse))

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		InstanceTypes: []string{"gpu_1x_a10"},
	})
	require.NoError(t, err)
	locations, err := getLambdaLabsLocations()
	require.NoError(t, err)
	assert.Len(t, instanceTypes, len(locations))

	for _, instanceType := range instanceTypes {
		assert.Equal(t, "gpu_1x_a10", instanceType.Type)
	}
}

func TestLambdaLabsClient_GetInstanceTypes_FilterBoth(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockResponse := createMockInstanceTypeResponse()
	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, mockResponse))

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		Locations:     v1.LocationsFilter{"us-east-1"},
		InstanceTypes: []string{"gpu_8x_h100"},
	})
	require.NoError(t, err)
	assert.Len(t, instanceTypes, 1)
	assert.Equal(t, "gpu_8x_h100", instanceTypes[0].Type)
	assert.Equal(t, "us-east-1", instanceTypes[0].Location)
}

func TestLambdaLabsClient_GetInstanceTypes_Empty(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	emptyResponse := openapi.InstanceTypes200Response{
		Data: map[string]openapi.InstanceTypes200ResponseDataValue{},
	}
	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, emptyResponse))

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	assert.Len(t, instanceTypes, 0)
}

func TestLambdaLabsClient_GetInstanceTypes_Error(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewStringResponder(500, `{"error": {"code": "INTERNAL_ERROR", "message": "Internal server error"}}`))

	_, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	assert.Error(t, err)
}

func TestLambdaLabsClient_GetInstanceTypePollTime(t *testing.T) {
	client := &LambdaLabsClient{}
	pollTime := client.GetInstanceTypePollTime()
	assert.Equal(t, 5*time.Minute, pollTime)
}

func TestConvertLambdaLabsInstanceTypeToV1InstanceType(t *testing.T) {
	llInstanceType := createMockLambdaLabsInstanceType("gpu_1x_a10", "1x NVIDIA A10 (24 GB)", "NVIDIA A10", 100)

	v1InstanceType, err := convertLambdaLabsInstanceTypeToV1InstanceType("us-west-1", llInstanceType, true)
	require.NoError(t, err)

	assert.Equal(t, "gpu_1x_a10", v1InstanceType.Type)
	assert.Equal(t, "us-west-1", v1InstanceType.Location)
	assert.True(t, v1InstanceType.IsAvailable)
	assert.Len(t, v1InstanceType.SupportedGPUs, 1)

	gpu := v1InstanceType.SupportedGPUs[0]
	assert.Equal(t, int32(1), gpu.Count)
	assert.Equal(t, "NVIDIA", gpu.Manufacturer)
	assert.Equal(t, "NVIDIA A10", gpu.Name)
	assert.Equal(t, "NVIDIA A10", gpu.Type)
	assert.Equal(t, units.Base2Bytes(24*1024*1024*1024), gpu.Memory)

	assert.NotNil(t, v1InstanceType.BasePrice)
	assert.Equal(t, "USD", v1InstanceType.BasePrice.CurrencyCode())
	assert.Equal(t, "1.00", v1InstanceType.BasePrice.Number())
}

func TestConvertLambdaLabsInstanceTypeToV1InstanceType_CPUOnly(t *testing.T) {
	llInstanceType := createMockLambdaLabsInstanceType("cpu_4x", "4x CPU cores", "", 50)

	v1InstanceType, err := convertLambdaLabsInstanceTypeToV1InstanceType("us-west-1", llInstanceType, true)
	require.NoError(t, err)

	assert.Equal(t, "cpu_4x", v1InstanceType.Type)
	assert.Equal(t, "us-west-1", v1InstanceType.Location)
	assert.True(t, v1InstanceType.IsAvailable)
	assert.Len(t, v1InstanceType.SupportedGPUs, 0)
}

func TestParseGPUFromDescription(t *testing.T) {
	tests := []struct {
		description string
		expected    v1.GPU
	}{
		{
			description: "1x H100 (80 GB SXM5)",
			expected: v1.GPU{
				Count:          1,
				Manufacturer:   "NVIDIA",
				Name:           "H100",
				Type:           "H100.SXM5",
				Memory:         80 * 1024 * 1024 * 1024,
				NetworkDetails: "80 GB SXM5",
				MemoryDetails:  "80 GB",
			},
		},
		{
			description: "8x Tesla V100 (16 GB)",
			expected: v1.GPU{
				Count:         8,
				Manufacturer:  "NVIDIA",
				Name:          "V100",
				Type:          "V100",
				Memory:        16 * 1024 * 1024 * 1024,
				MemoryDetails: "16 GB",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			gpu, err := parseGPUFromDescription(tt.description)
			require.NoError(t, err)
			assert.Equal(t, tt.expected.Count, gpu.Count)
			assert.Equal(t, tt.expected.Manufacturer, gpu.Manufacturer)
			assert.Equal(t, tt.expected.Name, gpu.Name)
			assert.Equal(t, tt.expected.Type, gpu.Type)
			assert.Equal(t, tt.expected.Memory, gpu.Memory)
		})
	}
}

func createMockInstanceTypeResponse() openapi.InstanceTypes200Response {
	return openapi.InstanceTypes200Response{
		Data: map[string]openapi.InstanceTypes200ResponseDataValue{
			"gpu_1x_a10": {
				InstanceType: createMockLambdaLabsInstanceType("gpu_1x_a10", "1x A10 (24 GB)", "A10", 100),
				RegionsWithCapacityAvailable: []openapi.Region{
					createMockRegion("us-west-1", "US West 1"),
					createMockRegion("us-east-1", "US East 1"),
				},
			},
			"gpu_8x_h100": {
				InstanceType: createMockLambdaLabsInstanceType("gpu_8x_h100", "8x H100 (80 GB SXM5)", "H100", 3200),
				RegionsWithCapacityAvailable: []openapi.Region{
					createMockRegion("us-east-1", "US East 1"),
				},
			},
		},
	}
}

func createMockLambdaLabsInstanceType(name, description, gpuDescription string, priceCents int32) openapi.InstanceType {
	gpuCount := int32(0)
	if gpuDescription != "" {
		gpuCount = 1
		if name == "gpu_8x_h100" {
			gpuCount = 8
		}
	}

	return openapi.InstanceType{
		Name:              name,
		Description:       description,
		GpuDescription:    gpuDescription,
		PriceCentsPerHour: priceCents,
		Specs: openapi.InstanceTypeSpecs{
			Vcpus:      8,
			MemoryGib:  32,
			StorageGib: 512,
			Gpus:       gpuCount,
		},
	}
}

func createMockRegion(name, description string) openapi.Region {
	return openapi.Region{
		Name:        name,
		Description: description,
	}
}

func findInstanceTypeByName(instanceTypes []v1.InstanceType, name string) *v1.InstanceType {
	for _, instanceType := range instanceTypes {
		if instanceType.Type == name {
			return &instanceType
		}
	}
	return nil
}
