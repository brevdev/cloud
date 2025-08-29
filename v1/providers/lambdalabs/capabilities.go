package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

// getLambdaLabsCapabilities returns the unified capabilities for Lambda Labs
// Based on API documentation at https://cloud.lambda.ai/api/v1/openapi.json
func getLambdaLabsCapabilities() v1.Capabilities {
	return v1.Capabilities{
		// SUPPORTED FEATURES (with API evidence):

		// Instance Management
		v1.CapabilityCreateInstance,          // POST /api/v1/instance-operations/launch
		v1.CapabilityTerminateInstance,       // POST /api/v1/instance-operations/terminate
		v1.CapabilityCreateTerminateInstance, // Combined create/terminate capability
		v1.CapabilityRebootInstance,          // POST /api/v1/instance-operations/restart

		// UNSUPPORTED FEATURES (no API evidence found):
		// - v1.CapabilityModifyFirewall        // Firewall management is project-level, not instance-level
		// - v1.CapabilityStopStartInstance     // No stop/start endpoints
		// - v1.CapabilityResizeInstanceVolume  // No volume resizing endpoints
		// - v1.CapabilityMachineImage          // No image endpoints
		// - v1.CapabilityTags                  // No tagging endpoints
	}
}

// GetCapabilities returns the capabilities of Lambda Labs client
func (c *LambdaLabsClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getLambdaLabsCapabilities(), nil
}

// GetCapabilities returns the capabilities for Lambda Labs credential
func (c *LambdaLabsCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getLambdaLabsCapabilities(), nil
}
