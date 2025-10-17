package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/brevdev/cloud/internal/errors"
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
	logger            v1.Logger
}

var _ v1.CloudClient = &NebiusClient{}

type NebiusClientOption func(c *NebiusClient)

func WithLogger(logger v1.Logger) NebiusClientOption {
	return func(c *NebiusClient) {
		c.logger = logger
	}
}

func NewNebiusClient(ctx context.Context, refID, serviceAccountKey, tenantID, projectID, location string) (*NebiusClient, error) {
	return NewNebiusClientWithOrg(ctx, refID, serviceAccountKey, tenantID, projectID, "", location)
}

func NewNebiusClientWithOrg(ctx context.Context, refID, serviceAccountKey, tenantID, projectID, organizationID, location string, opts ...NebiusClientOption) (*NebiusClient, error) {
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
		return nil, errors.WrapAndTrace(err)
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

	// DEBUG: Log projectID to diagnose corruption
	fmt.Printf("[NEBIUS_DEBUG] NewNebiusClient: refID=%s, location=%s, tenantID=%q (len=%d), projectID=%q (len=%d)\n",
		refID, location, tenantID, len(tenantID), projectID, len(projectID))

	client := &NebiusClient{
		refID:             refID,
		serviceAccountKey: serviceAccountKey,
		tenantID:          tenantID,
		projectID:         projectID,
		organizationID:    organizationID,
		location:          location,
		sdk:               sdk,
		logger:            &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(client)
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
		return "", errors.WrapAndTrace(err)
	}

	projects := projectsResp.GetItems()
	if len(projects) == 0 {
		return "", fmt.Errorf("no projects found in tenant %s", tenantID)
	}

	// Sort projects by ID for deterministic selection
	// This ensures CreateInstance and ListInstances always use the same project!
	sort.Slice(projects, func(i, j int) bool {
		if projects[i].Metadata == nil || projects[j].Metadata == nil {
			return false
		}
		return projects[i].Metadata.Id < projects[j].Metadata.Id
	})

	// Priority 1: Look for default-project-{region} or default-{region}
	preferredNames := []string{
		fmt.Sprintf("default-project-%s", region),
		fmt.Sprintf("default-%s", region),
		"default",
	}

	for _, preferredName := range preferredNames {
		for _, project := range projects {
			if project.Metadata != nil && strings.EqualFold(project.Metadata.Name, preferredName) {
				fmt.Printf("[NEBIUS_DEBUG] findProjectForRegion: Selected project by name match: %s (ID: %s)\n",
					project.Metadata.Name, project.Metadata.Id)
				return project.Metadata.Id, nil
			}
		}
	}

	// Priority 2: Look for any project with region in the name
	regionLower := strings.ToLower(region)
	for _, project := range projects {
		if project.Metadata != nil && strings.Contains(strings.ToLower(project.Metadata.Name), regionLower) {
			fmt.Printf("[NEBIUS_DEBUG] findProjectForRegion: Selected project by region in name: %s (ID: %s)\n",
				project.Metadata.Name, project.Metadata.Id)
			return project.Metadata.Id, nil
		}
	}

	// Priority 3: Return first available project (now deterministic due to sorting)
	if projects[0].Metadata != nil {
		fmt.Printf("[NEBIUS_DEBUG] findProjectForRegion: Selected first available project (sorted): %s (ID: %s)\n",
			projects[0].Metadata.Name, projects[0].Metadata.Id)
		fmt.Printf("[NEBIUS_DEBUG] findProjectForRegion: Total projects: %d, All IDs: %v\n",
			len(projects), func() []string {
				ids := make([]string, 0, len(projects))
				for _, p := range projects {
					if p.Metadata != nil {
						ids = append(ids, p.Metadata.Id)
					}
				}
				return ids
			}())
		return projects[0].Metadata.Id, nil
	}

	return "", fmt.Errorf("no suitable project found")
}

// discoverAllProjects returns all project IDs in the tenant
// This is used by ListInstances to query across all projects
func (c *NebiusClient) discoverAllProjects(ctx context.Context) ([]string, error) {
	pageSize := int64(1000)
	projectsResp, err := c.sdk.Services().IAM().V1().Project().List(ctx, &iam.ListProjectsRequest{
		ParentId: c.tenantID,
		PageSize: &pageSize,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}

	projects := projectsResp.GetItems()
	projectIDs := make([]string, 0, len(projects))
	for _, project := range projects {
		if project.Metadata != nil && project.Metadata.Id != "" {
			projectIDs = append(projectIDs, project.Metadata.Id)
		}
	}

	// Sort for consistency
	sort.Strings(projectIDs)

	return projectIDs, nil
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
	return c.MakeClientWithOptions(ctx, location)
}

func (c *NebiusClient) MakeClientWithOptions(ctx context.Context, location string, opts ...NebiusClientOption) (v1.CloudClient, error) {
	return NewNebiusClientWithOrg(ctx, c.refID, c.serviceAccountKey, c.tenantID, c.projectID, c.organizationID, location, opts...)
}

// GetTenantID returns the project ID (tenant ID) for this Brev user
func (c *NebiusClient) GetTenantID() (string, error) {
	return c.projectID, nil
}

// GetReferenceID returns the reference ID for this client
func (c *NebiusClient) GetReferenceID() string {
	return c.refID
}
