package v2

const (
	defaultPort        = 22
	defaultSSHUsername = "ubuntu"

	// Tag keys used to persist Brev metadata on SFCompute V2 instances.
	tagKeyCloudCredRefID = "brev-cloud-cred-ref-id"
	tagKeyStage          = "brev-stage"
	tagKeyRefID          = "brev-ref-id"
	tagKeyName           = "brev-name"

	// BrevProductionCapacityID is the SFCompute capacity used for Brev-managed instances.
	// TODO: replace with dynamic lookup from the Brev credential config; this is a stand-in.
	BrevProductionCapacityID = "brev-production-capacity"

	// BrevProductionImageID is the public SFCompute image "ubuntu-24.04.4-cuda-12.8" (vm_images.vm_image_id).
	BrevProductionImageID = "vmi_4GwEvmclFURy7ztFQjOdr"
)
