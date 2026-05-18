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
)

// Brev environment config for SFCompute V2.
const (
	brevDefaultImageID = "sfc:image:sfcompute:public:ubuntu-24.04.4-cuda-12.8"
)

func GetDefaultCapacityID(workspace string) string {
	return fmt.Sprintf("sfc:capacity:%s:default:brev-default-capacity", workspace)
}

func GetDefaultImageID() string {
	return brevDefaultImageID
}
