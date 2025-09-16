package v1

import (
	"context"
	"fmt"
	"slices"
)

type CloudLocation interface {
	GetLocations(ctx context.Context, args GetLocationsArgs) ([]Location, error)
}

type GetLocationsArgs struct {
	IncludeUnavailable bool
}

type Location struct {
	Name        string // basically the id
	Description string
	Available   bool
	Endpoint    string
	Priority    int
	Country     string // ISO 3166-1 alpha-3 https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
}

type LocationsFilter []string

var All = []string{"all"}

func (f LocationsFilter) IsAll() bool {
	for _, v := range f {
		if v == "*" || v == "all" {
			return true
		}
	}
	return false
}

func (f LocationsFilter) IsAllowed(location string) bool {
	if f.IsAll() {
		return true
	}
	return slices.Contains(f, location)
}

// ValidateGetLocations validates that the CloudLocation implementation returns at least one available location without error.
func ValidateGetLocations(ctx context.Context, client CloudLocation) error {
	locs, err := client.GetLocations(ctx, GetLocationsArgs{})
	if err != nil {
		return err
	}
	if len(locs) == 0 {
		return fmt.Errorf("no locations returned from GetLocations")
	}
	// Optionally, check that at least one location is available
	hasAvailable := false
	for _, loc := range locs {
		if loc.Available {
			hasAvailable = true
			break
		}
	}
	if !hasAvailable {
		return fmt.Errorf("no available locations found in GetLocations result")
	}
	return nil
}

const noSubLocation = "noSub"
