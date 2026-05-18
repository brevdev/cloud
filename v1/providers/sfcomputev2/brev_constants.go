package v2

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
// TODO: source these from environment variables rather than hardcoding them here.
const (
	// BrevProductionCapacityID is the SFCompute V2 capacity ID for Brev production instances.
	BrevProductionCapacityID = "brev-production-capacity"

	// BrevProductionImageID is the SFCompute image for Brev production instances
	// (ubuntu-24.04.4-cuda-12.8, vm_images.vm_image_id).
	BrevProductionImageID = "vmi_4GwEvmclFURy7ztFQjOdr"
)
