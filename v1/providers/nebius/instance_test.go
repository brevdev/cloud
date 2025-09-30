package v1

import (
	"context"
	"strings"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestClient() *NebiusClient {
	return &NebiusClient{
		refID: "test-ref",
		serviceAccountKey: `{
			"subject-credentials": {
				"type": "JWT",
				"alg": "RS256",
				"private-key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n",
				"kid": "publickey-test123",
				"iss": "serviceaccount-test456",
				"sub": "serviceaccount-test456"
			}
		}`,
		tenantID:  "test-tenant",
		projectID: "test-project",
		location:  "eu-north1",
	}
}

func TestNebiusClient_CreateInstance(t *testing.T) {
	client := createTestClient()
	ctx := context.Background()

	attrs := v1.CreateInstanceAttrs{
		RefID:        "test-instance-ref",
		Name:         "test-instance",
		InstanceType: "standard-2",
		ImageID:      "ubuntu-20.04",
		DiskSize:     50,
		Tags: map[string]string{
			"environment": "test",
			"team":        "dev",
		},
	}

	instance, err := client.CreateInstance(ctx, attrs)
	require.NoError(t, err)
	require.NotNil(t, instance)

	// Verify instance attributes
	assert.Equal(t, attrs.RefID, instance.RefID)
	assert.Equal(t, client.refID, instance.CloudCredRefID)
	assert.Equal(t, attrs.Name, instance.Name)
	assert.Equal(t, client.location, instance.Location)
	assert.Equal(t, attrs.InstanceType, instance.InstanceType)
	assert.Equal(t, attrs.ImageID, instance.ImageID)
	assert.Equal(t, attrs.DiskSize, instance.DiskSize)
	assert.Equal(t, attrs.Tags, instance.Tags)

	// Verify generated fields
	assert.Equal(t, v1.CloudProviderInstanceID("nebius-"+attrs.RefID), instance.CloudID)
	assert.Equal(t, v1.LifecycleStatusPending, instance.Status.LifecycleStatus)
	assert.WithinDuration(t, time.Now(), instance.CreatedAt, time.Second)
}

func TestNebiusClient_GetInstance(t *testing.T) {
	client := createTestClient()
	ctx := context.Background()

	instanceID := v1.CloudProviderInstanceID("test-instance-id")

	instance, err := client.GetInstance(ctx, instanceID)
	require.NoError(t, err)
	require.NotNil(t, instance)

	// Verify instance attributes from mock implementation
	assert.Equal(t, "sample-ref", instance.RefID)
	assert.Equal(t, client.refID, instance.CloudCredRefID)
	assert.Equal(t, "sample-instance", instance.Name)
	assert.Equal(t, instanceID, instance.CloudID)
	assert.Equal(t, client.location, instance.Location)
	assert.Equal(t, "sample-type", instance.InstanceType)
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.WithinDuration(t, time.Now(), instance.CreatedAt, time.Second)
}

func TestNebiusClient_NotImplementedMethods(t *testing.T) {
	client := createTestClient()
	ctx := context.Background()
	instanceID := v1.CloudProviderInstanceID("test-instance")

	tests := []struct {
		name string
		fn   func() error
	}{
		{
			name: "TerminateInstance",
			fn: func() error {
				return client.TerminateInstance(ctx, instanceID)
			},
		},
		{
			name: "ListInstances",
			fn: func() error {
				_, err := client.ListInstances(ctx, v1.ListInstancesArgs{})
				return err
			},
		},
		{
			name: "StopInstance",
			fn: func() error {
				return client.StopInstance(ctx, instanceID)
			},
		},
		{
			name: "StartInstance",
			fn: func() error {
				return client.StartInstance(ctx, instanceID)
			},
		},
		{
			name: "RebootInstance",
			fn: func() error {
				return client.RebootInstance(ctx, instanceID)
			},
		},
		{
			name: "ChangeInstanceType",
			fn: func() error {
				return client.ChangeInstanceType(ctx, instanceID, "new-type")
			},
		},
		{
			name: "UpdateInstanceTags",
			fn: func() error {
				return client.UpdateInstanceTags(ctx, v1.UpdateInstanceTagsArgs{
					InstanceID: instanceID,
					Tags: map[string]string{
						"new-tag": "value",
					},
				})
			},
		},
		{
			name: "ResizeInstanceVolume",
			fn: func() error {
				return client.ResizeInstanceVolume(ctx, v1.ResizeInstanceVolumeArgs{
					InstanceID: instanceID,
					Size:       100,
				})
			},
		},
		{
			name: "AddFirewallRulesToInstance",
			fn: func() error {
				return client.AddFirewallRulesToInstance(ctx, v1.AddFirewallRulesToInstanceArgs{
					InstanceID: instanceID,
				})
			},
		},
		{
			name: "RevokeSecurityGroupRules",
			fn: func() error {
				return client.RevokeSecurityGroupRules(ctx, v1.RevokeSecurityGroupRuleArgs{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			assert.Error(t, err)
			// Check for either "implementation pending" or "not yet implemented"
			errorMsg := err.Error()
			hasExpectedMsg := strings.Contains(errorMsg, "implementation pending") ||
							  strings.Contains(errorMsg, "not yet implemented")
			assert.True(t, hasExpectedMsg, "Expected error to contain 'implementation pending' or 'not yet implemented', got: %s", errorMsg)
		})
	}
}

func TestNebiusClient_GetLocations(t *testing.T) {
	client := createTestClient()
	ctx := context.Background()

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.Len(t, locations, 1)

	location := locations[0]
	assert.Equal(t, client.location, location.Name)
	assert.True(t, location.Available)
}

func TestNebiusClient_MergeInstanceForUpdate(t *testing.T) {
	client := createTestClient()

	currInstance := v1.Instance{
		RefID:          "current-ref",
		CloudCredRefID: "current-cred",
		Name:           "current-name",
		Location:       "current-location",
		CreatedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		CloudID:        "current-cloud-id",
		InstanceType:   "current-type",
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusRunning},
	}

	newInstance := v1.Instance{
		RefID:          "new-ref",
		CloudCredRefID: "new-cred",
		Name:           "new-name",
		Location:       "new-location",
		CreatedAt:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		CloudID:        "new-cloud-id",
		InstanceType:   "new-type",
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusStopped},
	}

	merged := client.MergeInstanceForUpdate(currInstance, newInstance)

	// These fields should be preserved from current instance
	assert.Equal(t, currInstance.RefID, merged.RefID)
	assert.Equal(t, currInstance.CloudCredRefID, merged.CloudCredRefID)
	assert.Equal(t, currInstance.Name, merged.Name)
	assert.Equal(t, currInstance.Location, merged.Location)
	assert.Equal(t, currInstance.CreatedAt, merged.CreatedAt)
	assert.Equal(t, currInstance.CloudID, merged.CloudID)

	// These fields should come from new instance
	assert.Equal(t, newInstance.InstanceType, merged.InstanceType)
	assert.Equal(t, newInstance.Status, merged.Status)
}

// BenchmarkCreateInstance benchmarks the CreateInstance method
func BenchmarkCreateInstance(b *testing.B) {
	client := createTestClient()
	ctx := context.Background()

	attrs := v1.CreateInstanceAttrs{
		RefID:        "bench-instance",
		Name:         "bench-test",
		InstanceType: "standard-2",
		ImageID:      "ubuntu-20.04",
		DiskSize:     50,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		attrs.RefID = "bench-instance-" + string(rune(i))
		_, err := client.CreateInstance(ctx, attrs)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGetInstance benchmarks the GetInstance method
func BenchmarkGetInstance(b *testing.B) {
	client := createTestClient()
	ctx := context.Background()
	instanceID := v1.CloudProviderInstanceID("bench-instance")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetInstance(ctx, instanceID)
		if err != nil {
			b.Fatal(err)
		}
	}
}