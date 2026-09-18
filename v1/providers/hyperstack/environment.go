package hyperstack

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/environment"
)

func (c *HyperstackClient) getDefaultEnvironment(ctx context.Context, location string) (environment.EnvironmentFields, error) {
	environmentName := defaultEnvironmentTag + location
	pageSize := strconv.Itoa(defaultPageSize)
	for page := 1; ; page++ {
		pageNumber := strconv.Itoa(page)
		response, err := c.environments.ListEnvironmentsWithResponse(ctx, &environment.ListEnvironmentsParams{
			Page:     &pageNumber,
			PageSize: &pageSize,
			Search:   &environmentName,
		})
		if err != nil {
			return environment.EnvironmentFields{}, wrapTransportError("list environments", err)
		}
		if response.StatusCode() != http.StatusOK {
			return environment.EnvironmentFields{}, responseError("list environments", response.StatusCode(), response.Body, nil)
		}
		if response.JSON200 == nil || response.JSON200.Environments == nil {
			return environment.EnvironmentFields{}, errors.New("hyperstack list environments response did not contain data")
		}

		providerEnvironments := *response.JSON200.Environments
		for _, providerEnvironment := range providerEnvironments {
			if stringValue(providerEnvironment.Name) != environmentName || stringValue(providerEnvironment.Region) != location {
				continue
			}
			if providerEnvironment.Id == nil || *providerEnvironment.Id <= 0 {
				return environment.EnvironmentFields{}, fmt.Errorf("hyperstack environment %q did not contain an ID", environmentName)
			}
			return providerEnvironment, nil
		}
		if len(providerEnvironments) < defaultPageSize {
			return environment.EnvironmentFields{}, fmt.Errorf("hyperstack environment %q was not found in location %q", environmentName, location)
		}
	}
}
