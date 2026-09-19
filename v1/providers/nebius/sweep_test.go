package v1

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	vpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
)

// CIRunIDLabel aliases the validation key so the VM sweep and this resource sweep agree.
const CIRunIDLabel = validation.CIRunIDLabel

// sweepPageSize is the max page size ([1...1000] per the Nebius API). Filtering is client-side, so
// every page must be fetched.
const sweepPageSize = 1000

type resourceSweepResult struct {
	disks    int
	subnets  int
	networks int
	failures int
}

// sweepStandaloneResources deletes this run's leftover network/subnet/disk (VM gone), in dependency
// order: disks, subnets, then VPCs. Matches only ci-run-id == runID, so it never touches prod.
func (c *NebiusClient) sweepStandaloneResources(ctx context.Context, runID string, logf func(string, ...any)) (resourceSweepResult, error) {
	var res resourceSweepResult
	if runID == "" {
		return res, fmt.Errorf("sweepStandaloneResources: runID is required")
	}
	projects, discovered := c.sweepProjects(ctx, logf)
	if !discovered {
		res.failures++ // narrowed scan may miss this run's resources in other projects; fail loudly
	}

	var failed int
	res.disks, failed = resourceSweepPass(ctx, "disk", projects, logf, c.listDisks, keepUnattachedDisk(runID), c.deleteBootDiskIfExists)
	res.failures += failed
	res.subnets, failed = resourceSweepPass(ctx, "subnet", projects, logf, c.listSubnets, keepByLabel[*vpc.Subnet](runID), c.deleteSubnetIfExists)
	res.failures += failed
	res.networks, failed = resourceSweepPass(ctx, "network", projects, logf, c.listNetworks, keepByLabel[*vpc.Network](runID), c.deleteNetworkIfExists)
	res.failures += failed
	return res, nil
}

// sweepProjects lists every project to scan. discovered is false when tenant-wide discovery fails
// and it falls back to the primary project alone — a narrowed scan that may miss resources elsewhere,
// which the caller treats as a failure.
func (c *NebiusClient) sweepProjects(ctx context.Context, logf func(string, ...any)) (projects []string, discovered bool) {
	projectToRegion, err := c.discoverAllProjectsWithRegions(ctx)
	if err != nil || len(projectToRegion) == 0 {
		logf("[sweep] WARNING: project discovery failed (err=%v); scanning ONLY primary project %s — resources in other projects will be missed", err, c.projectID)
		return []string{c.projectID}, false
	}
	projects = make([]string, 0, len(projectToRegion))
	for projectID := range projectToRegion {
		projects = append(projects, projectID)
	}
	return projects, true
}

// resourceSweepPass paginates one resource type across every project: keep selects this run's
// deletable items, del removes them. A list error counts as a failure (so a partial scan is never
// silent) and stops that project's pages; a not-found on delete is success.
func resourceSweepPass[T any](
	ctx context.Context,
	kind string,
	projects []string,
	logf func(string, ...any),
	list func(ctx context.Context, projectID, pageToken string) (items []T, nextPageToken string, err error),
	keep func(item T) (id, name string, ok bool),
	del func(ctx context.Context, id string) error,
) (deleted, failures int) {
	for _, projectID := range projects {
		pageToken := ""
		for {
			items, next, err := list(ctx, projectID, pageToken)
			if err != nil {
				logf("[sweep] list %ss in %s failed: %v", kind, projectID, err)
				failures++
				break
			}
			for _, item := range items {
				id, name, ok := keep(item)
				if !ok {
					continue
				}
				logf("[sweep] %s %s name=%q", kind, id, name)
				if derr := del(ctx, id); derr != nil && !isNotFoundError(derr) {
					logf("[sweep] FAILED %s %s: %v", kind, id, derr)
					failures++
					continue
				}
				deleted++
			}
			if pageToken = next; pageToken == "" {
				break
			}
		}
	}
	return deleted, failures
}

// keepByLabel keeps resources whose ci-run-id label == runID, returning their id/name.
func keepByLabel[T interface {
	GetMetadata() *common.ResourceMetadata
}](runID string) func(item T) (id, name string, ok bool) {
	return func(item T) (string, string, bool) {
		md := item.GetMetadata()
		if md == nil || md.GetLabels()[CIRunIDLabel] != runID {
			return "", "", false
		}
		return md.GetId(), md.GetName(), true
	}
}

// keepUnattachedDisk keeps this run's disks that aren't attached to a VM (attached ones belong to a
// live VM the VM pass handles).
func keepUnattachedDisk(runID string) func(disk *compute.Disk) (id, name string, ok bool) {
	base := keepByLabel[*compute.Disk](runID)
	return func(disk *compute.Disk) (string, string, bool) {
		id, name, ok := base(disk)
		if !ok {
			return "", "", false
		}
		if st := disk.GetStatus(); st.GetReadWriteAttachment() != "" || len(st.GetReadOnlyAttachments()) > 0 {
			return "", "", false
		}
		return id, name, true
	}
}

// listDisks, listSubnets, and listNetworks each fetch one page of a resource type within a project.
func (c *NebiusClient) listDisks(ctx context.Context, projectID, pageToken string) ([]*compute.Disk, string, error) {
	resp, err := c.sdk.Services().Compute().V1().Disk().List(ctx, &compute.ListDisksRequest{
		ParentId: projectID, PageSize: sweepPageSize, PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	return resp.GetItems(), resp.GetNextPageToken(), nil
}

func (c *NebiusClient) listSubnets(ctx context.Context, projectID, pageToken string) ([]*vpc.Subnet, string, error) {
	resp, err := c.sdk.Services().VPC().V1().Subnet().List(ctx, &vpc.ListSubnetsRequest{
		ParentId: projectID, PageSize: sweepPageSize, PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	return resp.GetItems(), resp.GetNextPageToken(), nil
}

func (c *NebiusClient) listNetworks(ctx context.Context, projectID, pageToken string) ([]*vpc.Network, string, error) {
	resp, err := c.sdk.Services().VPC().V1().Network().List(ctx, &vpc.ListNetworksRequest{
		ParentId: projectID, PageSize: sweepPageSize, PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	return resp.GetItems(), resp.GetNextPageToken(), nil
}

// TestSweepOrphans deletes this run's leftover VMs then its standalone network/subnet/disk. Gated on
// SWEEP_ORPHANS=true and CI_RUN_ID; deletes only resources whose ci-run-id == CI_RUN_ID.
func TestSweepOrphans(t *testing.T) {
	if os.Getenv("SWEEP_ORPHANS") != "true" {
		t.Skip("SWEEP_ORPHANS is not set to true, skipping orphan sweep")
	}
	runID := os.Getenv("CI_RUN_ID")
	if runID == "" {
		t.Skip("CI_RUN_ID is not set; the run-scoped sweep has nothing to match")
	}

	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")
	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("Skipping sweep: NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}
	if _, statErr := os.Stat(serviceAccountJSON); statErr == nil {
		//nolint:gosec // maintenance entrypoint reading the service account from a controlled CI env
		data, err := os.ReadFile(serviceAccountJSON)
		if err != nil {
			t.Fatalf("failed to read service account file: %v", err)
		}
		serviceAccountJSON = string(data)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	// Ref ID is arbitrary; ownership comes from the ci-run-id label, not this client.
	client, err := NewNebiusClient(ctx, "orphan-sweeper", serviceAccountJSON, tenantID, "", defaultNebiusLocation)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	// VMs first — terminate cascades their attached network/disk.
	vmRes, err := validation.SweepOrphanedInstances(ctx, client, validation.SweepOpts{RunID: runID, Logf: t.Logf})
	if err != nil {
		t.Fatalf("VM sweep failed: %v", err)
	}

	// Then standalone network/subnet/disk left behind.
	resRes, err := client.sweepStandaloneResources(ctx, runID, t.Logf)
	if err != nil {
		t.Fatalf("resource sweep failed: %v", err)
	}

	t.Logf("sweep summary: runID=%s vms_deleted=%d vms_failed=%d disks=%d subnets=%d networks=%d res_failures=%d",
		runID, len(vmRes.Deleted), len(vmRes.Failed), resRes.disks, resRes.subnets, resRes.networks, resRes.failures)
	if len(vmRes.Failed) > 0 || resRes.failures > 0 {
		t.Errorf("sweep had failures: vms=%d resources=%d (see logs above)", len(vmRes.Failed), resRes.failures)
	}
}
