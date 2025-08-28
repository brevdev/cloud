package v1

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/google/go-cmp/cmp"
)

type InstanceTypeID string

type InstanceType struct {
	ID       InstanceTypeID // this id should be unique across all regions and stable
	Provider string
	Type     string

	Location    string
	SubLocation string

	SupportedGPUs          []GPU
	SupportedStorage       []Storage
	SupportedUsageClasses  []string
	SupportedArchitectures []string
	SupportedNumCores      []int32

	VCPU            int32
	Memory          units.Base2Bytes
	DefaultCores    int32
	ClockSpeedInGhz float64

	MaximumNetworkInterfaces int32
	NetworkPerformance       string

	IsAvailable bool
	Quota       InstanceTypeQuota
	BasePrice   *currency.Amount

	EstimatedDeployTime *time.Duration

	// capabilities
	CanModifyFirewallRules          bool // can we modify the firewall rules?
	IsContainer                     bool // is the instance a container?
	UserPrivilegeEscalationDisabled bool // can the user escalate privileges? (processes can not be more privileged than initial process)
	NotPrivileged                   bool // is the instance not privileged? (i.e. no sudo)
	Stoppable                       bool // can the instance be stopped?
	Rebootable                      bool // can the instance be rebooted?
	VariablePrice                   bool // will the price change over time?
	Preemptible                     bool // can the instance be preempted?
	ElasticRootVolume               bool // can we change the root volume size? (i.e. can we resize the root volume?)
	SubLocationTypeChangeable       bool // can we change the instance type to a different type in the same sublocation?
}

func MakeGenericInstanceTypeID(instanceType InstanceType) InstanceTypeID {
	if instanceType.ID != "" {
		return instanceType.ID
	}
	subLoc := noSubLocation
	return InstanceTypeID(fmt.Sprintf("%s-%s-%s", instanceType.Location, subLoc, instanceType.Type))
}

func MakeGenericInstanceTypeIDFromInstance(instance Instance) InstanceTypeID {
	if instance.InstanceTypeID != "" {
		return instance.InstanceTypeID
	}
	subLoc := noSubLocation
	if instance.SubLocation != "" {
		subLoc = instance.SubLocation
	}
	return InstanceTypeID(fmt.Sprintf("%s-%s-%s", instance.Location, subLoc, instance.InstanceType))
}

type GPU struct {
	Count          int32
	Memory         units.Base2Bytes
	MemoryDetails  string
	NetworkDetails string
	Manufacturer   string
	Name           string
	Type           string
}

type InstanceTypeQuota struct {
	OnDemand Quota
	Spot     Quota
	Reserved Quota
}

type CloudInstanceType interface {
	GetInstanceTypes(ctx context.Context, args GetInstanceTypeArgs) ([]InstanceType, error)
	GetInstanceTypePollTime() time.Duration
	CloudLocation
}

type GetInstanceTypeArgs struct {
	Locations              LocationsFilter
	SupportedArchitectures []string
	InstanceTypes          []string
}

// ValidateGetInstanceTypes validates that the GetInstanceTypes functionality works correctly
// by testing that filtering by specific instance types returns the expected results
func ValidateGetInstanceTypes(ctx context.Context, client CloudInstanceType) error { //nolint:funlen,gocyclo // todo refactor
	// Get all instance types first
	allTypes, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{})
	if err != nil {
		return fmt.Errorf("failed to get all instance types: %w", err)
	}

	if len(allTypes) == 0 {
		return errors.New("no instance types available for validation")
	}

	// Test 1: Deterministic results - multiple calls should return the same results (order-insensitive)
	allTypes2, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{})
	if err != nil {
		return fmt.Errorf("failed to get all instance types on second call: %w", err)
	}

	// Remove volatile fields for comparison
	normalizedTypes1 := normalizeInstanceTypes(allTypes)
	normalizedTypes2 := normalizeInstanceTypes(allTypes2)

	// Build maps keyed by ID for order-insensitive comparison
	map1 := make(map[InstanceTypeID]InstanceType)
	for _, t := range normalizedTypes1 {
		map1[t.ID] = t
	}
	map2 := make(map[InstanceTypeID]InstanceType)
	for _, t := range normalizedTypes2 {
		map2[t.ID] = t
	}

	// Compare keys
	if len(map1) != len(map2) {
		return fmt.Errorf("instance types are not deterministic between calls: different number of types (%d vs %d)", len(map1), len(map2))
	}
	for id, t1 := range map1 {
		t2, ok := map2[id]
		if !ok {
			return fmt.Errorf("instance type ID %s present in first call but missing in second", id)
		}
		if !reflect.DeepEqual(t1, t2) {
			diff := cmp.Diff(t1, t2)
			fmt.Printf("Instance type with ID %s differs between calls. Diff:\n%s\n", id, diff)
			return fmt.Errorf("instance type with ID %s differs between calls", id)
		}
	}

	// Test 2: ID stability and uniqueness
	idMap := make(map[InstanceTypeID]InstanceType)
	for _, instanceType := range allTypes {
		if existing, exists := idMap[instanceType.ID]; exists {
			return fmt.Errorf("duplicate instance type ID found: %s (types: %s, %s)",
				instanceType.ID, existing.Type, instanceType.Type)
		}
		idMap[instanceType.ID] = instanceType
	}

	// Test 3: Filtering by instance type name
	firstType := allTypes[0]
	filteredTypes, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{
		InstanceTypes: []string{firstType.Type},
	})
	if err != nil {
		return fmt.Errorf("failed to get filtered instance types: %w", err)
	}

	if len(filteredTypes) == 0 {
		return fmt.Errorf("no instance types returned when filtering by type: %s", firstType.Type)
	}

	// Compare the first type with the filtered result, ignoring fields that may vary
	// between different calls or implementations
	expectedType := firstType
	expectedType.ID = ""
	expectedType.SubLocation = ""

	// Find the matching type in filteredTypes by ID (since order is not guaranteed)
	var actualType InstanceType
	found := false
	for _, t := range filteredTypes {
		tmp := t
		tmp.ID = ""
		tmp.SubLocation = ""
		if reflect.DeepEqual(expectedType, tmp) {
			actualType = tmp
			found = true
			break
		}
	}
	if !found {
		// If not found by struct equality, just compare the first filtered type for debugging
		actualType = filteredTypes[0]
		actualType.ID = ""
		actualType.SubLocation = ""
		diff := cmp.Diff(expectedType, actualType)
		fmt.Printf("Filtered instance type does not match expected type. Diff:\n%s\n", diff)
		return fmt.Errorf("filtered instance type does not match expected type: expected %+v, got %+v", expectedType, actualType)
	}

	return nil
}

// ValidateLocationalInstanceTypes validates that locational filtering works correctly
// by comparing locational results with all-location results using CloudLocation capabilities
func ValidateLocationalInstanceTypes(ctx context.Context, client CloudInstanceType) error {
	// Get all-location instance types by requesting from all locations
	allLocationTypes, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{
		Locations: All,
	})
	if err != nil {
		// If all-location is not supported, skip this validation
		return fmt.Errorf("all-location instance types not supported: %w", err)
	}

	if len(allLocationTypes) == 0 {
		return errors.New("no all-location instance types available for validation")
	}

	locationToTest := allLocationTypes[0].Location
	// Get locational instance types (default behavior - typically current location)
	locationalTypes, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{
		Locations: LocationsFilter{locationToTest},
	})
	if err != nil {
		return fmt.Errorf("failed to get locational instance types: %w", err)
	}

	if len(locationalTypes) == 0 {
		return errors.New("no locational instance types available for validation")
	}

	// Validate that locational results are a subset of all-location results
	if len(locationalTypes) >= len(allLocationTypes) {
		return fmt.Errorf("locational instance types (%d) should be fewer than all-location types (%d)",
			len(locationalTypes), len(allLocationTypes))
	}

	// Create a map of all-location types for efficient lookup
	allLocationMap := make(map[InstanceTypeID]InstanceType)
	for _, instanceType := range allLocationTypes {
		allLocationMap[instanceType.ID] = instanceType
	}

	// Validate that all locational types exist in all-location results
	for _, locationalType := range locationalTypes {
		if _, exists := allLocationMap[locationalType.ID]; !exists {
			return fmt.Errorf("locational instance type %s not found in all-location results", locationalType.ID)
		}
	}

	// Additional validation: ensure locational types have appropriate location information
	for _, locationalType := range locationalTypes {
		if locationalType.Location == "" {
			return fmt.Errorf("locational instance type %s should have location information", locationalType.ID)
		}
	}

	return nil
}

// normalizeInstanceTypes removes volatile fields that may change between calls
func normalizeInstanceTypes(types []InstanceType) []InstanceType {
	normalized := make([]InstanceType, len(types))
	for i, instanceType := range types {
		normalized[i] = instanceType
		// Remove fields that may vary between calls
		normalized[i].BasePrice = nil
		normalized[i].Quota = InstanceTypeQuota{}
		normalized[i].EstimatedDeployTime = nil
	}
	return normalized
}

// ValidateStableInstanceTypeIDs validates that the provided stable instance type IDs are valid and stable
// This function ensures that stable IDs exist in the current instance types and have required properties
func ValidateStableInstanceTypeIDs(ctx context.Context, client CloudInstanceType, stableIDs []InstanceTypeID) error {
	// Get all instance types
	allTypes, err := client.GetInstanceTypes(ctx, GetInstanceTypeArgs{})
	if err != nil {
		return fmt.Errorf("failed to get instance types: %w", err)
	}

	if len(allTypes) == 0 {
		return errors.New("no instance types available for validation")
	}

	// Group types by ID for efficient lookup
	typesByID := make(map[InstanceTypeID][]InstanceType)
	for _, instanceType := range allTypes {
		typesByID[instanceType.ID] = append(typesByID[instanceType.ID], instanceType)
	}

	// Validate that each ID has exactly one instance type (uniqueness)
	for id, types := range typesByID {
		if len(types) != 1 {
			return fmt.Errorf("instance type id %s should be unique, found %d instances", id, len(types))
		}
	}

	// Validate that stable IDs are not empty
	if len(stableIDs) == 0 {
		return errors.New("stable IDs list cannot be empty")
	}

	// Validate that all stable IDs exist in current instance types, collecting all errors
	var errs []error
	for _, stableID := range stableIDs {
		if _, exists := typesByID[stableID]; !exists {
			errs = append(errs, fmt.Errorf("instance type id %s should be stable but not found", stableID)) // if this fails, we may need to coordinate a migration of the stable ID
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	// Validate that all instance types have required properties
	for _, instanceType := range allTypes {
		// Check that instance type has base price
		if instanceType.BasePrice == nil {
			return fmt.Errorf("instance type %s should have base price", instanceType.ID)
		}

		// Check that supported storage has price information
		for i, storage := range instanceType.SupportedStorage {
			if storage.MinSize != nil {
				if storage.PricePerGBHr == nil {
					return fmt.Errorf("instance type %s should have storage %d price", instanceType.ID, i)
				}
			}
		}
	}

	return nil
}
