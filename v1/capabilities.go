package v1

import "slices"

type Capability string

type Capabilities []Capability

func (c Capabilities) IsCapable(cc Capability) bool {
	return slices.Contains(c, cc)
}

const (
	CapabilityCreateInstance           Capability = "create-instance"
	CapabilityCreateIdempotentInstance Capability = "create-instance-idempotent"
	CapabilityTerminateInstance        Capability = "terminate-instance"
	CapabilityCreateTerminateInstance  Capability = "create-terminate-instance"
	CapabilityInstanceUserData         Capability = "instance-userdata" // specify user data when creating an instance in CreateInstanceAttrs // should be in instance type
	CapabilityTags                     Capability = "tags"
	CapabilityRebootInstance           Capability = "reboot-instance"
	CapabilityResizeInstanceVolume     Capability = "resize-instance-volume"
	CapabilityStopStartInstance        Capability = "stop-start-instance"
	CapabilityMachineImage             Capability = "machine-image"
	CapabilityModifyFirewall           Capability = "modify-firewall"
	CapabilityVPC                      Capability = "vpc"
	CapabilityManagedKubernetes        Capability = "managed-kubernetes"
)
