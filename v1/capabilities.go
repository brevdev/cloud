package v1

type Capability string

type Capabilities []Capability

func (c Capabilities) IsCapable(cc Capability) bool {
	for _, capability := range c {
		if capability == cc {
			return true
		}
	}
	return false
}

const (
	CapabilityCreateInstance           Capability = "create-instance"
	CapabilityCreateIdempotentInstance Capability = "create-instance-idempotent"
	CapabilityTerminateInstance        Capability = "terminate-instance"
)

const (
	CapabilityCreateTerminateInstance Capability = "create-terminate-instance"
	CapabilityInstanceUserData        Capability = "instance-userdata" // specify user data when creating an instance in CreateInstanceAttrs // should be in instance type
)

const CapabilityTags Capability = "tags"

const CapabilityRebootInstance Capability = "reboot-instance"

const CapabilityResizeInstanceVolume Capability = "resize-instance-volume"

const CapabilityStopStartInstance Capability = "stop-start-instance"

const CapabilityMachineImage Capability = "machine-image"

const CapabilityModifyFirewall Capability = "modify-firewall"
