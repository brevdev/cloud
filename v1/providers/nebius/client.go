package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
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

	// Determine projectID: use provided ID, or find first available project, or use tenant ID
	if projectID == "" {
		// Try to find an existing project in the tenant for this region
		foundProjectID, err := findProjectForRegion(ctx, sdk, tenantID, location)
		if err == nil && foundProjectID != "" {
			projectID = foundProjectID
		} else {
			// Fallback: try default-project-{region} naming pattern
			projectID = fmt.Sprintf("default-project-%s", location)
		}
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

	return client, nil
}

// findProjectForRegion attempts to find an existing project for the given region
// Priority:
// 1. Project named "default-project-{region}" or "default-{region}"
// 2. First project with region in the name
// 3. First available project
func findProjectForRegion(ctx context.Context, sdk *gosdk.SDK, tenantID, region string) (string, error) {
	pageSize := int64(1000)
	projectsResp, err := sdk.Services().IAM().V1().Project().List(ctx, &iam.ListProjectsRequest{
		ParentId: tenantID,
		PageSize: &pageSize,
	})
	if err != nil {
		return "", fmt.Errorf("failed to list projects: %w", err)
	}

	projects := projectsResp.GetItems()
	if len(projects) == 0 {
		return "", fmt.Errorf("no projects found in tenant %s", tenantID)
	}

	// Priority 1: Look for default-project-{region} or default-{region}
	preferredNames := []string{
		fmt.Sprintf("default-project-%s", region),
		fmt.Sprintf("default-%s", region),
		"default",
	}

	for _, preferredName := range preferredNames {
		for _, project := range projects {
			if project.Metadata != nil && strings.EqualFold(project.Metadata.Name, preferredName) {
				return project.Metadata.Id, nil
			}
		}
	}

	// Priority 2: Look for any project with region in the name
	regionLower := strings.ToLower(region)
	for _, project := range projects {
		if project.Metadata != nil && strings.Contains(strings.ToLower(project.Metadata.Name), regionLower) {
			return project.Metadata.Id, nil
		}
	}

	// Priority 3: Return first available project
	if projects[0].Metadata != nil {
		return projects[0].Metadata.Id, nil
	}

	return "", fmt.Errorf("no suitable project found")
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

// FIXME for b64 decode on cred JSON
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
