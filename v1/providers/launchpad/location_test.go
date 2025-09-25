package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
)

func TestGetLocations(t *testing.T) {
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

	locations, err := client.GetLocations(context.Background(), v1.GetLocationsArgs{})
	if err != nil {
		t.Fatalf("failed to get locations: %v", err)
	}

	t.Logf("locations: %v", locations)
}

func TestInstanceTypesToLocations(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                   string
		launchpadInstanceTypes []openapi.InstanceType
		locations              []v1.Location
		args                   v1.GetLocationsArgs
	}{
		{
			name: "include unavailable",
			launchpadInstanceTypes: []openapi.InstanceType{
				{
					Capacity: map[string]int32{
						"us-east-1": 1,
						"us-west-1": 0,
					},
				},
			},
			args: v1.GetLocationsArgs{
				IncludeUnavailable: true,
			},
			locations: []v1.Location{
				{
					Name:        "us-east-1",
					Description: "us-east-1",
					Available:   true,
				},
				{
					Name:        "us-west-1",
					Description: "us-west-1",
					Available:   false,
				},
			},
		},
		{
			name: "exclude unavailable",
			launchpadInstanceTypes: []openapi.InstanceType{
				{
					Capacity: map[string]int32{
						"us-east-1": 1,
						"us-west-1": 0,
					},
				},
			},
			args: v1.GetLocationsArgs{
				IncludeUnavailable: false,
			},
			locations: []v1.Location{
				{
					Name:        "us-east-1",
					Description: "us-east-1",
					Available:   true,
				},
			},
		},
		{
			name: "mixed availability across instance types",
			launchpadInstanceTypes: []openapi.InstanceType{
				{
					Capacity: map[string]int32{
						"us-east-1": 1,
						"us-west-1": 0, // unavailable
					},
				},
				{
					Capacity: map[string]int32{
						"us-west-1": 1, // available
					},
				},
			},
			args: v1.GetLocationsArgs{
				IncludeUnavailable: false,
			},
			locations: []v1.Location{
				{
					Name:        "us-east-1",
					Description: "us-east-1",
					Available:   true,
				},
				{
					Name:        "us-west-1",
					Description: "us-west-1",
					Available:   true,
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			locations := launchpadInstanceTypesToLocations(tt.launchpadInstanceTypes, tt.args)
			require.ElementsMatch(t, tt.locations, locations)
		})
	}
}
