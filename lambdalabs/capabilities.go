package lambdalabs

import (
	"context"

	"github.com/brevdev/cloud"
)

// getLambdaLabsCapabilities returns the unified capabilities for Lambda Labs
// Based on API documentation at https://cloud.lambda.ai/api/v1/openapi.json
func getLambdaLabsCapabilities() cloud.Capabilities {
	return cloud.Capabilities{
		// SUPPORTED FEATURES (with API evidence):

		// Instance Management
		cloud.CapabilityCreateInstance,          // POST /api/v1/instance-operations/launch
		cloud.CapabilityTerminateInstance,       // POST /api/v1/instance-operations/terminate
		cloud.CapabilityCreateTerminateInstance, // Combined create/terminate capability
		cloud.CapabilityRebootInstance,          // POST /api/v1/instance-operations/restart

		// UNSUPPORTED FEATURES (no API evidence found):
		// - cloud.CapabilityModifyFirewall        // Firewall management is project-level, not instance-level
		// - cloud.CapabilityStopStartInstance     // No stop/start endpoints
		// - cloud.CapabilityResizeInstanceVolume  // No volume resizing endpoints
		// - cloud.CapabilityMachineImage          // No image endpoints
		// - cloud.CapabilityTags                  // No tagging endpoints
	}
}

// GetCapabilities returns the capabilities of Lambda Labs client
func (c *LambdaLabsClient) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getLambdaLabsCapabilities(), nil
}

// GetCapabilities returns the capabilities for Lambda Labs credential
func (c *LambdaLabsCredential) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getLambdaLabsCapabilities(), nil
}
