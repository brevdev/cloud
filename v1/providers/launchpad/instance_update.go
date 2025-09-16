package v1

import (
	v1 "github.com/brevdev/cloud/v1"
)

func (c *LaunchpadClient) MergeInstanceForUpdate(_ v1.Instance, newInstance v1.Instance) v1.Instance {
	return newInstance
}

func (c *LaunchpadClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newInstanceType v1.InstanceType) v1.InstanceType {
	return newInstanceType
}
