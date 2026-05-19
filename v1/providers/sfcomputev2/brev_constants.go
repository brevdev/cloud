package v2

import "fmt"

// Package-internal constants — SSH defaults and internal tag keys.
const (
	defaultPort        = 22
	defaultSSHUsername = "ubuntu"

	// Internal tag keys written to every SFCompute V2 instance. These are stripped from
	// v1.Instance.Tags on read so they don't surface as user-facing tags.
	tagKeyCloudCredRefID = "brev-cloud-cred-ref-id" //nolint:gosec // not a secret
	tagKeyRefID          = "brev-ref-id"

	// Brev environment config for SFCompute V2.
	brevDefaultImageResourcePath = "sfc:image:sfcompute:public:ubuntu-24.04.4-cuda-12.8"
)

func (c *SFCClientV2) GetDefaultCapacityResourcePath() string {
	return fmt.Sprintf("sfc:capacity:%s:%s:brev-default-capacity", c.organization, c.workspace)
}

func (c *SFCClientV2) GetWorkspaceResourcePath() string {
	return fmt.Sprintf("sfc:workspace:%s:%s", c.organization, c.workspace)
}

func (c *SFCClientV2) GetDefaultImageResourcePath() string {
	return brevDefaultImageResourcePath
}
