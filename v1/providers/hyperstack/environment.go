package hyperstack

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/environment"
)

const (
	defaultEnvironmentTag    = "default-"
	listEnvironmentsPageSize = 100
)

func (c *HyperstackClient) getDefaultEnvironment(ctx context.Context, location string) (environment.EnvironmentFields, error) {
	environmentName := defaultEnvironmentTag + location

	for page := 1; ; page++ {
		providerEnvironments, err := c.listEnvironmentPage(ctx, listEnvironmentPageArgs{
			Page:     page,
			PageSize: listEnvironmentsPageSize,
			Search:   environmentName,
		})
		if err != nil {
			return environment.EnvironmentFields{}, err
		}

		for _, providerEnvironment := range providerEnvironments {
			if stringValue(providerEnvironment.Name) != environmentName || stringValue(providerEnvironment.Region) != location {
				continue
			}
			if providerEnvironment.Id == nil || *providerEnvironment.Id <= 0 {
				return environment.EnvironmentFields{}, fmt.Errorf("hyperstack environment %q did not contain an ID", environmentName)
			}
			return providerEnvironment, nil
		}

		if len(providerEnvironments) < listEnvironmentsPageSize {
			return environment.EnvironmentFields{}, fmt.Errorf("hyperstack environment %q was not found in location %q", environmentName, location)
		}
	}
}

type listEnvironmentPageArgs struct {
	Page     int
	PageSize int
	Search   string
}

func (c *HyperstackClient) listEnvironmentPage(ctx context.Context, args listEnvironmentPageArgs) ([]environment.EnvironmentFields, error) {
	pageNumber := strconv.Itoa(args.Page)
	pageSizeStr := strconv.Itoa(args.PageSize)
	response, err := c.environments.ListEnvironmentsWithResponse(ctx, &environment.ListEnvironmentsParams{
		Page:     &pageNumber,
		PageSize: &pageSizeStr,
		Search:   &args.Search,
	})
	if err != nil {
		return nil, wrapTransportError("list environments", err)
	}
	if response.StatusCode() != http.StatusOK {
		return nil, responseError("list environments", response.StatusCode(), response.Body, nil)
	}
	if response.JSON200 == nil || response.JSON200.Environments == nil {
		return nil, errors.New("hyperstack list environments response did not contain data")
	}
	return *response.JSON200.Environments, nil
}
