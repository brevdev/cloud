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

## TODO

- [ ] Implement actual API integration for supported features
- [ ] Add proper service account authentication handling
- [ ] Add comprehensive error handling and retry logic
- [ ] Add logging and monitoring
- [ ] Add comprehensive testing
- [ ] Investigate VPC integration for networking features
- [ ] Verify instance type changes work correctly via ResourcesSpec.preset field
