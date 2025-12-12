package v1

import (
	"context"

	"github.com/brevdev/cloud/internal/collections"
	"github.com/brevdev/cloud/internal/errors"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
)

func (c *LaunchpadClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	deployments, err := c.paginateListDeployments(ctx, 1000)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	insts, err := collections.MapE(deployments, func(deployment openapi.Deployment) (v1.Instance, error) {
		inst, err := launchpadDeploymentToInstance(&deployment)
		if err != nil {
			return v1.Instance{}, errors.WrapAndTrace(err)
		}
		return inst, nil
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	if len(args.InstanceIDs) > 0 {
		insts = collections.Filter(insts, func(inst v1.Instance) bool {
			return collections.ListContains(args.InstanceIDs, inst.CloudID)
		})
	}
	return insts, nil
}

func (c *LaunchpadClient) paginateListDeployments(ctx context.Context, pageSize int32) ([]openapi.Deployment, error) {
	deployments := make([]openapi.Deployment, 0, pageSize)
	var page int32 = 1
	for {
		listRes, resp, err := c.client.CatalogDeploymentsAPI.V1CatalogDeploymentsList(c.makeAuthContext(ctx)).
			Expand(clusterExpandParameter).
			PageSize(pageSize).
			Page(page).
			Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // handled in err check
		}
		if err != nil && resp == nil {
			return nil, errors.WrapAndTrace(err)
		}
		if err != nil {
			return nil, errors.WrapAndTrace(c.handleLaunchpadAPIErr(ctx, resp, err))
		}
		deployments = append(deployments, listRes.Results...)
		if len(listRes.Results) < int(pageSize) {
			break
		}
		page++
	}
	return deployments, nil
}
