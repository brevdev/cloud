# Nebius Provider

This directory contains the Nebius provider implementation for the compute package.

## Overview

The Nebius provider implements the CloudClient interface defined in `pkg/v1` to provide access to Nebius AI Cloud infrastructure. This implementation is based on the official Nebius API documentation at https://github.com/nebius/api and uses the Nebius Go SDK.

## Supported Features

Based on the Nebius API documentation, the following features are **SUPPORTED**:

### Instance Management
- ✅ **Create Instance**: `InstanceService.Create` in compute/v1/instance_service.proto
- ✅ **Get Instance**: `InstanceService.Get` and `InstanceService.GetByName` 
- ✅ **List Instances**: `InstanceService.List` with pagination support
- ✅ **Terminate Instance**: `InstanceService.Delete`
- ✅ **Stop Instance**: `InstanceService.Stop`
- ✅ **Start Instance**: `InstanceService.Start`

### Instance Updates
- ✅ **Update Instance Tags**: Maps to `UpdateInstanceTags` in CloudClient interface
- ✅ **Change Instance Type**: Maps to `ChangeInstanceType` in CloudClient interface via `ResourcesSpec.preset` field in `InstanceService.Update`

### GPU Cluster Management
- ✅ **Create GPU Cluster**: `GpuClusterService.Create` in compute/v1/gpu_cluster_service.proto
- ✅ **Get GPU Cluster**: `GpuClusterService.Get` and `GpuClusterService.GetByName`
- ✅ **List GPU Clusters**: `GpuClusterService.List` with pagination support
- ✅ **Delete GPU Cluster**: `GpuClusterService.Delete`
- ✅ **Update GPU Cluster**: `GpuClusterService.Update`

### Machine Images
- ✅ **Get Images**: `ImageService.Get`, `ImageService.GetByName`, `ImageService.GetLatestByFamily`
- ✅ **List Images**: `ImageService.List` with filtering support

### Quota Management
- ✅ **Get Quotas**: `QuotaAllowanceService` in quotas/v1/quota_allowance_service.proto

## Unsupported Features

The following features are **NOT SUPPORTED** (no clear API endpoints found):

### Instance Operations
- ❌ **Reboot Instance**: No reboot endpoint found in instance_service.proto
- ❌ **General Instance Updates**: Nebius InstanceService.Update exists but most InstanceSpec fields are immutable; only specific updates like tags and instance type are supported through dedicated CloudClient methods

### Volume Management
- ❌ **Resize Instance Volume**: Volume resizing not clearly documented

### Location Management
- ❌ **Get Locations**: No location listing service found

### Firewall Management
- ❌ **Firewall Rules**: Network security handled through VPC service, not instance-level firewall rules

## Implementation Approach

This implementation uses the `NotImplCloudClient` pattern for unsupported features:
- Supported features have TODO implementations with API service references
- Unsupported features return `ErrNotImplemented` (handled by embedded NotImplCloudClient)
- Full CloudClient interface compliance is maintained

## Nebius API

The provider integrates with the Nebius AI Cloud API:
- Base URL: `{service-name}.api.nebius.cloud:443` (gRPC)
- Authentication: Service account based (JWT tokens)
- SDK: `github.com/nebius/gosdk`
- Documentation: https://github.com/nebius/api
- API Type: Locational (location-specific endpoints)

## Key Features

Nebius AI Cloud is known for:
- GPU instances and GPU clusters for AI/ML workloads
- Comprehensive compute, storage, and networking services
- gRPC-based API with strong typing
- Service account authentication with JWT tokens
- Location-specific API endpoints
- Advanced operations tracking and idempotency
- Integration with VPC, IAM, billing, and quota services
- Container registry and managed services

## Implementation Notes

### Platform Name vs Platform ID
The Nebius API requires **platform NAME** (e.g., `"gpu-h100-sxm"`) in `ResourcesSpec.Platform`, **NOT** platform ID (e.g., `"computeplatform-e00caqbn6nysa972yq"`). The `parseInstanceType` function must always return `platform.Metadata.Name`, not `platform.Metadata.Id`.

### GPU VRAM Mapping
GPU memory (VRAM) is populated via static mapping since the Nebius SDK doesn't natively provide this information:
- L40S: 48 GiB
- H100: 80 GiB
- H200: 141 GiB
- A100: 80 GiB
- V100: 32 GiB
- A10: 24 GiB
- T4: 16 GiB
- L4: 24 GiB
- B200: 192 GiB

See `getGPUMemory()` in `instancetype.go` for the complete mapping.

### Logging Support
The Nebius provider supports structured logging via the `v1.Logger` interface. To enable logging:

```go
import (
    nebiusv1 "github.com/brevdev/cloud/v1/providers/nebius"
    "github.com/brevdev/cloud/v1"
)

// Create a logger (implement v1.Logger interface)
logger := myLogger{}

// Option 1: Via credential
cred := nebiusv1.NewNebiusCredential(refID, serviceKey, tenantID)
client, err := cred.MakeClientWithOptions(ctx, location, nebiusv1.WithLogger(logger))

// Option 2: Via direct client construction
client, err := nebiusv1.NewNebiusClientWithOrg(ctx, refID, serviceKey, tenantID, projectID, orgID, location, nebiusv1.WithLogger(logger))
```

Without a logger, the client defaults to `v1.NoopLogger{}` which discards all log messages.

### Error Tracing
Critical error paths use `errors.WrapAndTrace()` from `github.com/brevdev/cloud/internal/errors` to add stack traces and detailed context to errors. This improves debugging when errors propagate through the system.

### Resource Naming and Correlation
All Nebius resources (instances, VPCs, subnets, boot disks) are named using the `RefID` (environment ID) for easy correlation:
- VPC: `{refID}-vpc`
- Subnet: `{refID}-subnet`
- Boot Disk: `{refID}-boot-disk`
- Instance: User-provided name

All resources include the `environment-id` label for filtering and tracking.

### Automatic Cleanup on Failure
If instance creation fails at any step, all created resources are automatically cleaned up to prevent orphaned resources:
- Boot disks
- Subnets
- VPC networks

This cleanup is handled via a deferred function that tracks all created resource IDs and deletes them if the operation doesn't complete successfully.

## TODO

- [ ] Add comprehensive error handling and retry logic
- [ ] Investigate VPC integration for networking features
- [ ] Verify instance type changes work correctly via ResourcesSpec.preset field
