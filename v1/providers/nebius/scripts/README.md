# Nebius Provider Scripts

This directory contains utility scripts for testing and enumerating Nebius cloud resources. All scripts are implemented as Go test files with the `scripts` build tag.

## Prerequisites

Export your Nebius credentials as environment variables:

```bash
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'
export NEBIUS_LOCATION='eu-north1'  # Optional, defaults to eu-north1
```

## Instance Type Enumeration

### Enumerate All Regions

Lists all instance types across all Nebius regions with GPU type breakdowns:

```bash
cd v1/providers/nebius
go test -tags scripts -v -run Test_EnumerateInstanceTypes ./scripts/
```

**Output:**
- Console summary with region-by-region GPU counts
- JSON file: `instance_types_all_regions.json`

### Enumerate Single Region

Lists instance types for a specific region with detailed specifications:

```bash
export NEBIUS_LOCATION='eu-north1'
go test -tags scripts -v -run Test_EnumerateInstanceTypesSingleRegion ./scripts/
```

**Output:**
- Console summary categorized by CPU/GPU types
- JSON file: `instance_types_eu-north1.json`

### GPU Types Only

Displays only GPU instance types in a formatted table:

```bash
export NEBIUS_LOCATION='eu-north1'
go test -tags scripts -v -run Test_EnumerateGPUTypes ./scripts/
```

**Example Output:**
```
ID                                                 GPU Type        Count    vCPUs      RAM (GB)   VRAM/GPU (GB)
------------------------------------------------------------------------------------------------------------------------
nebius-eu-north1-l40s-1gpu-16vcpu-96gb            L40S            1        16         96         48
nebius-eu-north1-l40s-4gpu-128vcpu-768gb          L40S            4        128        768        48
nebius-eu-north1-h100-8gpu-128vcpu-1600gb         H100            8        128        1600       80
```

## Image Enumeration

### Enumerate Images (Single Region)

Lists all available images in a specific region:

```bash
export NEBIUS_LOCATION='eu-north1'
go test -tags scripts -v -run Test_EnumerateImages ./scripts/
```

**Output:**
- Console summary organized by OS
- JSON file: `images_eu-north1.json`

### Enumerate Images (All Regions)

Lists images across all Nebius regions:

```bash
go test -tags scripts -v -run Test_EnumerateImagesAllRegions ./scripts/
```

**Output:**
- Console summary with image counts per region
- JSON file: `images_all_regions.json`

### Filter GPU-Optimized Images

Shows only images suitable for GPU instances (CUDA, ML, etc.):

```bash
export NEBIUS_LOCATION='eu-north1'
go test -tags scripts -v -run Test_FilterGPUImages ./scripts/
```

## VPC and Kubernetes Scripts

### Create VPC

Creates a test VPC with public/private subnets:

```bash
go test -tags scripts -v -run TestCreateVPC ./scripts/
```

### Create Kubernetes Cluster

Creates a Kubernetes cluster with VPC:

```bash
go test -tags scripts -v -run Test_CreateVPCAndCluster ./scripts/
```

## Running All Scripts

To run all enumeration scripts at once:

```bash
go test -tags scripts -v ./scripts/
```

## Output Files

Scripts generate JSON files in the current directory:
- `instance_types_all_regions.json` - All instance types across regions
- `instance_types_<region>.json` - Instance types for specific region
- `images_all_regions.json` - All images across regions
- `images_<region>.json` - Images for specific region

## Tips

### Pretty Print JSON Output

```bash
cat instance_types_eu-north1.json | jq '.'
```

### Filter JSON Results

```bash
# Show only L40S instance types
cat instance_types_eu-north1.json | jq '.[] | select(.supported_gpus[0].type == "L40S")'

# Show instance types with pricing
cat instance_types_eu-north1.json | jq '.[] | select(.price != null) | {id, price}'

# Count GPU types
cat instance_types_all_regions.json | jq -r '.[].supported_gpus[0].type' | sort | uniq -c
```

### Redirect Output to File

```bash
go test -tags scripts -v -run Test_EnumerateGPUTypes ./scripts/ > gpu_types_output.txt 2>&1
```

## Integration with Testing Guide

These scripts complement the integration tests documented in [`NEBIUS_TESTING_GUIDE.md`](../NEBIUS_TESTING_GUIDE.md). Use them for:
- Discovery: Finding available instance types and regions
- Validation: Verifying quota and availability
- Development: Testing new features with real Nebius resources
