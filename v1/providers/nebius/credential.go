package v1

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "nebius"

// NebiusCredential implements the CloudCredential interface for Nebius AI Cloud
type NebiusCredential struct {
	RefID             string
	ServiceAccountKey string `json:"sa_json"`   // JSON service account key
	TenantID          string `json:"tenant_id"` // Nebius tenant ID (top-level organization)
}

var _ v1.CloudCredential = &NebiusCredential{}

// NewNebiusCredential creates a new Nebius credential
func NewNebiusCredential(refID, serviceAccountKey, tenantID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:             refID,
		ServiceAccountKey: serviceAccountKey,
		TenantID:          tenantID,
	}
}

// NewNebiusCredentialWithOrg creates a new Nebius credential with organization ID
func NewNebiusCredentialWithOrg(refID, serviceAccountKey, tenantID, organizationID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:             refID,
		ServiceAccountKey: serviceAccountKey,
		TenantID:          tenantID,
	}
}

// GetReferenceID returns the reference ID for this credential
func (c *NebiusCredential) GetReferenceID() string {
	return c.RefID
}

// GetAPIType returns the API type for Nebius
func (c *NebiusCredential) GetAPIType() v1.APIType {
	return v1.APITypeLocational // Nebius uses location-specific endpoints
}

// GetCloudProviderID returns the cloud provider ID for Nebius
func (c *NebiusCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

// GetTenantID returns a unique project ID for this Brev user within the tenant
// This groups all instances from the same user into a single Nebius project
func (c *NebiusCredential) GetTenantID() (string, error) {
	if c.TenantID == "" {
		return "", fmt.Errorf("tenant ID is required for Nebius project creation")
	}
	// Create a deterministic project ID based on user ID
	// Format: project-{userID} to match Nebius expected project ID format
	// We'll truncate and sanitize the user ID to meet Nebius naming requirements
	sanitizedUserID := sanitizeForNebiusID(c.TenantID)
	return fmt.Sprintf("project-%s", sanitizedUserID), nil
}

// MakeClient creates a new Nebius client from this credential
func (c *NebiusCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	projectID, err := c.GetTenantID()
	if err != nil {
		return nil, fmt.Errorf("failed to get project ID: %w", err)
	}
	return NewNebiusClientWithOrg(ctx, c.RefID, c.ServiceAccountKey, c.TenantID, projectID, "", location)
}

// sanitizeForNebiusID sanitizes a user ID to meet Nebius project ID naming requirements
func sanitizeForNebiusID(userID string) string {
	// Nebius project IDs should be lowercase and contain only alphanumeric characters and hyphens
	// Based on the error pattern: ^([a-z][a-z0-9]{2,49})-([a-z][a-z0-9]{2})(.+?)(?:--([a-z-][a-z0-9-]{0,9}))?$
	// Let's simplify to just use alphanumeric characters

	// Convert to lowercase
	sanitized := strings.ToLower(userID)

	// Replace any non-alphanumeric characters with hyphens
	re := regexp.MustCompile(`[^a-z0-9]`)
	sanitized = re.ReplaceAllString(sanitized, "-")

	// Remove multiple consecutive hyphens
	re = regexp.MustCompile(`-+`)
	sanitized = re.ReplaceAllString(sanitized, "-")

	// Remove leading/trailing hyphens
	sanitized = strings.Trim(sanitized, "-")

	// Limit length to ensure we don't exceed Nebius limits
	if len(sanitized) > 20 {
		sanitized = sanitized[:20]
	}

	// Ensure it starts with a letter
	if len(sanitized) > 0 && !regexp.MustCompile(`^[a-z]`).MatchString(sanitized) {
		sanitized = "u" + sanitized
	}

	return sanitized
}
