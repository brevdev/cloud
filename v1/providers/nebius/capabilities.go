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
		v1.CapabilityModifyFirewall,    // Nebius has Security Groups for firewall management
		v1.CapabilityMachineImage,      // Nebius supports custom machine images
		v1.CapabilityTags,              // Nebius supports resource tagging
		v1.CapabilityInstanceUserData,  // Nebius supports user data in instance creation
		v1.CapabilityVPC,               // Nebius supports VPCs
		v1.CapabilityManagedKubernetes, // Nebius supports managed Kubernetes clusters
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
