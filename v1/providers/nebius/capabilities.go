package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

// getNebiusCapabilities returns the unified capabilities for Nebius AI Cloud
// Based on Nebius compute API and our implementation
func getNebiusCapabilities() v1.Capabilities {
	return v1.Capabilities{
		// SUPPORTED FEATURES:

		// Instance Management
		v1.CapabilityCreateInstance,          // Nebius compute instance creation
		v1.CapabilityTerminateInstance,       // Nebius compute instance termination
		v1.CapabilityCreateTerminateInstance, // Combined create/terminate capability
		v1.CapabilityRebootInstance,          // Nebius instance restart
		v1.CapabilityStopStartInstance,       // Nebius instance stop/start operations
		v1.CapabilityResizeInstanceVolume,    // Nebius volume resizing

		// Resource Management
		v1.CapabilityMachineImage, // Nebius image management
		v1.CapabilityTags,         // Nebius resource labeling

		// PARTIALLY SUPPORTED (infrastructure implemented):
		// - Network management (VPC, subnets) - implemented
		// - Project management - implemented
		// - Boot disk management - implemented

		// FUTURE ENHANCEMENTS:
		// - v1.CapabilityModifyFirewall  // Network security groups (future)
	}
}

// GetCapabilities returns the capabilities of Nebius client
func (c *NebiusClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getNebiusCapabilities(), nil
}

// GetCapabilities returns the capabilities for Nebius credential
func (c *NebiusCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getNebiusCapabilities(), nil
}