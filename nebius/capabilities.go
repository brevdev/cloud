package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func getNebiusCapabilities() cloud.Capabilities {
	return cloud.Capabilities{
		// SUPPORTED FEATURES (with API evidence):

		// Instance Management
		cloud.CapabilityCreateInstance,          // Nebius compute API supports instance creation
		cloud.CapabilityTerminateInstance,       // Nebius compute API supports instance deletion
		cloud.CapabilityCreateTerminateInstance, // Combined create/terminate capability
		cloud.CapabilityRebootInstance,          // Nebius supports instance restart operations
		cloud.CapabilityStopStartInstance,       // Nebius supports instance stop/start operations

		cloud.CapabilityModifyFirewall,       // Nebius has Security Groups for firewall management
		cloud.CapabilityMachineImage,         // Nebius supports custom machine images
		cloud.CapabilityResizeInstanceVolume, // Nebius supports disk resizing
		cloud.CapabilityTags,                 // Nebius supports resource tagging
		cloud.CapabilityInstanceUserData,     // Nebius supports user data in instance creation

	}
}

// GetCapabilities returns the capabilities of Nebius client
func (c *NebiusClient) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getNebiusCapabilities(), nil
}

// GetCapabilities returns the capabilities for Nebius credential
func (c *NebiusCredential) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getNebiusCapabilities(), nil
}
