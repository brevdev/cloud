package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestGetInstanceTypes(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewLaunchpadCredential("brev-cloud-sdk-test", apiKey, "https://stage.launchpad.api.nvidia.com"),
	}

	client, err := config.Credential.MakeClient(context.Background(), config.Location)
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		Locations: []string{"all"},
	})
	if err != nil {
		t.Fatalf("failed to get instance types: %v", err)
	}

	t.Logf("instance types: %v", instanceTypes)
}

func TestInstanceTypeInfo(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name             string
		instanceType     string
		instanceTypeInfo instanceTypeInfo
	}{
		{
			name:         "valid instance type",
			instanceType: "nebius.l4x1.pcie.public",
			instanceTypeInfo: instanceTypeInfo{
				cloud:             "nebius",
				gpuName:           "l4",
				gpuCount:          1,
				gpuNetworkDetails: "pcie",
				workshopID:        "public",
			},
		},
		{
			name:         "valid instance type",
			instanceType: "oci.b200x8.sxm6.dgxc",
			instanceTypeInfo: instanceTypeInfo{
				cloud:             "oci",
				gpuName:           "b200",
				gpuCount:          8,
				gpuNetworkDetails: "sxm6",
				workshopID:        "dgxc",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			instanceTypeInfo, err := getInstanceTypeInfo(tt.instanceType)
			if err != nil {
				t.Fatalf("failed to get instance type info: %v", err)
			}

			// validate instance type info
			require.Equal(t, tt.instanceTypeInfo, instanceTypeInfo)

			// validate round trip
			instanceTypeValue := makeInstanceTypeName(tt.instanceTypeInfo)
			require.Equal(t, tt.instanceType, instanceTypeValue)
		})
	}
}

func TestMakeGenericInstanceTypeID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		instanceType   v1.InstanceType
		instanceTypeID v1.InstanceTypeID
	}{
		{
			name: "existing ID",
			instanceType: v1.InstanceType{
				ID:           "abc123",
				Location:     "foo",
				AvailableAzs: []string{"bar"},
				Type:         "baz",
			},
			instanceTypeID: "abc123",
		},
		{
			name: "no available azs",
			instanceType: v1.InstanceType{
				Location: "us-east-1",
				Type:     "a100",
			},
			instanceTypeID: "us-east-1-noSub-a100",
		},
		{
			name: "available azs",
			instanceType: v1.InstanceType{
				Location:     "us-east-1",
				Type:         "a100",
				AvailableAzs: []string{"us-east-1a"},
			},
			instanceTypeID: "us-east-1-us-east-1a-a100",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			instanceTypeID := v1.MakeGenericInstanceTypeID(tt.instanceType)
			require.Equal(t, tt.instanceTypeID, instanceTypeID)
		})
	}
}
