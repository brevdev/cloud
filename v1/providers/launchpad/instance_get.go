package v1

import (
	"context"

	"github.com/alecthomas/units"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
)

const (
	clusterExpandParameter = "cluster.nodes.location.provider,cluster.nodes.gpu"
)

func (c *LaunchpadClient) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	getDeployment, resp, err := c.client.CatalogDeploymentsAPI.V1CatalogDeploymentsRetrieve(c.makeAuthContext(ctx), string(id)).
		Expand(clusterExpandParameter).
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
	inst, err := launchpadDeploymentToInstance(getDeployment)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &inst, nil
}

func launchpadDeploymentToInstance(deployment *openapi.Deployment) (v1.Instance, error) {
	if deployment == nil {
		return v1.Instance{}, errors.WrapAndTrace(errors.New("deployment is nil"))
	}

	tags, err := launchpadTagsToInstanceTags(deployment.Tags)
	if err != nil {
		return v1.Instance{}, errors.WrapAndTrace(err)
	}

	var totalStorageSize int32
	nodes := deployment.GetCluster().Cluster.GetNodes()
	if len(nodes) == 0 {
		totalStorageSize = 0
	} else {
		node := nodes[0]

		// Calculate disk size
		storage := node.Node.GetStorage()
		for _, s := range storage {
			totalStorageSize += s.GetSize()
		}
	}

	inst := v1.Instance{
		Name:           getValueFromTags("Name", tags),
		RefID:          getValueFromTags("RefID", tags),
		CloudCredRefID: getValueFromTags("CloudCredRefID", tags),
		CloudID:        v1.CloudProviderInstanceID(deployment.Id),
		CreatedAt:      deployment.Created,
		SSHUser:        deployment.SshUser,
		SSHPort:        int(deployment.SshPort),
		Status: v1.Status{
			LifecycleStatus: launchpadStateToLifecycleStatus(deployment.State),
		},
		Tags: tags,
		InternalPortMappings: []v1.PortMapping{
			{
				FromPort: int(deployment.SshPort),
				ToPort:   2022,
			},
		},
		DiskSize:          units.Base2Bytes(totalStorageSize) * units.GiB,
		DiskSizeByteValue: v1.NewBytes(v1.BytesValue(totalStorageSize), v1.Gigabyte),
		Location:          deployment.GetRegion(),
		PublicDNS:         deployment.GetCluster().Cluster.GetPublicAddress(),
		PublicIP:          deployment.GetCluster().Cluster.GetPublicAddress(),
	}

	cluster := deployment.GetCluster().Cluster
	if cluster != nil {
		instanceType := launchpadClusterToInstanceType(*deployment.GetCluster().Cluster)
		if instanceType != nil {
			inst.InstanceType = instanceType.Type
			inst.InstanceTypeID = instanceType.ID
		}
	}

	return inst, nil
}

func launchpadTagsToInstanceTags(tags interface{}) (map[string]string, error) {
	tagsMap, ok := tags.(map[string]interface{})
	if !ok {
		return nil, errors.WrapAndTrace(errors.New("tags interface casting error"))
	}
	result := make(map[string]string)
	for key, value := range tagsMap {
		valueString, ok := value.(string)
		if !ok {
			return nil, errors.WrapAndTrace(errors.New("tags interface casting error"))
		}
		result[key] = valueString
	}
	return result, nil
}

func launchpadStateToLifecycleStatus(launchpadState openapi.DeploymentState) v1.LifecycleStatus {
	switch launchpadState {
	case openapi.DeploymentStateWaiting: // waiting for deployment to be ready
		return v1.LifecycleStatusPending
	case openapi.DeploymentStateReady: // deployment is ready and all instances are running
		return v1.LifecycleStatusRunning
	case openapi.DeploymentStateFailed: // deployment has failed but may be retried
		return v1.LifecycleStatusPending
	case openapi.DeploymentStateDestroying: // deployment is being destroyed
		return v1.LifecycleStatusTerminating
	case openapi.DeploymentStateDestroyed: // deployment has been fully destroyed
		return v1.LifecycleStatusTerminated
	case openapi.DeploymentStateError: // deployment has encountered a fatal error and will not be retried
		return v1.LifecycleStatusFailed
	case openapi.DeploymentStateStopping: // deployment instances are stopping
		return v1.LifecycleStatusStopping
	case openapi.DeploymentStateStarting: // deployment instances are starting
		return v1.LifecycleStatusPending
	case openapi.DeploymentStateStopped: // deployment instances are stopped
		return v1.LifecycleStatusStopped
	case openapi.DeploymentStateRetrying: // deployment is retrying
		return v1.LifecycleStatusPending
	default:
		return v1.LifecycleStatusPending
	}
}

func getValueFromTags(tagKey string, tags map[string]string) string {
	value, ok := tags[tagKey]
	if !ok {
		return ""
	}
	return value
}
