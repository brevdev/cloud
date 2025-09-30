package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	iam "github.com/nebius/gosdk/proto/nebius/iam/v1"
)


// It embeds NotImplCloudClient to handle unsupported features
type NebiusClient struct {
	v1.NotImplCloudClient
	refID             string
	serviceAccountKey string
	tenantID          string // Nebius tenant (organization)
	projectID         string // Nebius project (per-user)
	organizationID    string // Brev organization ID (maps to tenant_uuid)
	location          string
	sdk               *gosdk.SDK
}

var _ v1.CloudClient = &NebiusClient{}

func NewNebiusClient(ctx context.Context, refID, serviceAccountKey, tenantID, projectID, location string) (*NebiusClient, error) {
	return NewNebiusClientWithOrg(ctx, refID, serviceAccountKey, tenantID, projectID, "", location)
}

func NewNebiusClientWithOrg(ctx context.Context, refID, serviceAccountKey, tenantID, projectID, organizationID, location string) (*NebiusClient, error) {
	// Initialize SDK with proper service account credentials
	var creds gosdk.Credentials

	// Check if serviceAccountKey is a file path or JSON content
	if _, err := os.Stat(serviceAccountKey); err == nil {
		// It's a file path - use ServiceAccountCredentialsFileParser
		parser := auth.NewServiceAccountCredentialsFileParser(nil, serviceAccountKey)
		creds = gosdk.ServiceAccountReader(parser)
	} else {
		// It's JSON content - parse it manually and create ServiceAccount
		var credFile auth.ServiceAccountCredentials
		if err := json.Unmarshal([]byte(serviceAccountKey), &credFile); err != nil {
			return nil, fmt.Errorf("failed to parse service account key JSON: %w", err)
		}

		// Basic validation of the structure
		if credFile.SubjectCredentials.Alg != "RS256" {
			return nil, fmt.Errorf("invalid service account algorithm: %s. Only RS256 is supported", credFile.SubjectCredentials.Alg)
		}
		if credFile.SubjectCredentials.Issuer != credFile.SubjectCredentials.Subject {
			return nil, fmt.Errorf("invalid service account subject must be the same as issuer")
		}

		// Create service account parser from the parsed content
		parser := auth.NewPrivateKeyParser(
			[]byte(credFile.SubjectCredentials.PrivateKey),
			credFile.SubjectCredentials.KeyID,
			credFile.SubjectCredentials.Subject,
		)
		creds = gosdk.ServiceAccountReader(parser)
	}

	sdk, err := gosdk.New(ctx, gosdk.WithCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Nebius SDK: %w", err)
	}

	client := &NebiusClient{
		refID:             refID,
		serviceAccountKey: serviceAccountKey,
		tenantID:          tenantID,
		projectID:         projectID,
		organizationID:    organizationID,
		location:          location,
		sdk:               sdk,
	}

	// Ensure the user's project exists (create if needed)
	if err := client.ensureProjectExists(ctx); err != nil {
		return nil, fmt.Errorf("failed to ensure project exists: %w", err)
	}

	return client, nil
}

// GetAPIType returns the API type for Nebius
func (c *NebiusClient) GetAPIType() v1.APIType {
	return v1.APITypeLocational
}

// GetCloudProviderID returns the cloud provider ID for Nebius
func (c *NebiusClient) GetCloudProviderID() v1.CloudProviderID {
	return "nebius"
}

// MakeClient creates a new client instance for a different location
func (c *NebiusClient) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return NewNebiusClient(ctx, c.refID, c.serviceAccountKey, c.tenantID, c.projectID, location)
}

// GetTenantID returns the project ID (tenant ID) for this Brev user
func (c *NebiusClient) GetTenantID() (string, error) {
	return c.projectID, nil
}

// GetReferenceID returns the reference ID for this client
func (c *NebiusClient) GetReferenceID() string {
	return c.refID
}

// ensureProjectExists creates a Nebius project for this user if it doesn't exist
func (c *NebiusClient) ensureProjectExists(ctx context.Context) error {
	// First, try to find existing project by name pattern
	existingProjectID, err := c.findExistingProject(ctx)
	if err == nil && existingProjectID != "" {
		// Update our project ID to use the existing project
		c.projectID = existingProjectID
		return nil
	}

	// Try to get the project by ID to see if it exists
	_, err = c.sdk.Services().IAM().V1().Project().Get(ctx, &iam.GetProjectRequest{
		Id: c.projectID,
	})
	if err != nil {
		// Check if the error is "not found", then create the project
		if isNotFoundError(err) {
			// Project doesn't exist, create it
			return c.createProject(ctx)
		}
		// Some other error occurred
		return fmt.Errorf("failed to check if project exists: %w", err)
	}

	// Project exists, we're good
	return nil
}

// createProject creates a new project within the tenant
func (c *NebiusClient) createProject(ctx context.Context) error {
	labels := map[string]string{
		"created-by":     "brev-cloud-sdk",
		"brev-user":      c.refID,
		"project-type":   "user-instances",
	}

	// Add organization ID if available (correlates to Brev Organization)
	if c.organizationID != "" {
		labels["tenant-uuid"] = c.organizationID // Maps to tenant_uuid in Terraform
		labels["brev-organization"] = c.organizationID
	}

	createReq := &iam.CreateProjectRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.tenantID,
			Name:     fmt.Sprintf("brev-user-%s", c.refID),
			Labels:   labels,
		},
		// Spec: &iam.ProjectSpec{
		//	// Add any specific project configuration if needed
		// },
	}

	operation, err := c.sdk.Services().IAM().V1().Project().Create(ctx, createReq)
	if err != nil {
		// Check if project already exists (this is OK)
		if isAlreadyExistsError(err) {
			return nil // Project already exists, we're good
		}
		return fmt.Errorf("failed to create project: %w", err)
	}

	// Wait for project creation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for project creation: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("project creation failed: %v", finalOp.Status())
	}

	return nil
}

// findExistingProject finds an existing project by looking for the expected name pattern
func (c *NebiusClient) findExistingProject(ctx context.Context) (string, error) {
	expectedName := fmt.Sprintf("brev-user-%s", c.refID)

	resp, err := c.sdk.Services().IAM().V1().Project().List(ctx, &iam.ListProjectsRequest{
		ParentId: c.tenantID,
	})
	if err != nil {
		return "", err
	}

	// Look for project with matching name
	for _, project := range resp.GetItems() {
		if project.Metadata != nil && project.Metadata.Name == expectedName {
			return project.Metadata.Id, nil
		}
	}

	return "", fmt.Errorf("no existing project found with name: %s", expectedName)
}

