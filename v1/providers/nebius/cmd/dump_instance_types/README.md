# Instance Types Dump Utility

This utility aggregates Nebius instance types across regions into a single view per preset configuration, matching the LaunchPad API format.

**Features**:
- ✅ Cross-region aggregation with capacity maps
- ✅ **Real pricing from Nebius Billing Calculator API** (optional)
- ✅ LaunchPad-compatible JSON format
- ✅ Elastic storage details (50GB-2560GB)

## Usage

### Quick Mode (No Pricing)

```bash
# Set credentials
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'

# Run the dump (instant, pricing = 0)
cd /home/jmorgan/VS/brev-cloud-sdk/v1/providers/nebius
go run ./cmd/dump_instance_types/main.go > instance_types.json
```

### With Real Pricing (Recommended)

```bash
# Set tenant-level credentials only (no project ID needed!)
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-e00xxx'

# Run with pricing (takes ~60 seconds for 20+ instance types)
FETCH_PRICING=true go run ./cmd/dump_instance_types/main.go > complete_catalog.json
```

**Note**: Pricing is catalog-level and doesn't vary by project. The tool automatically creates/finds a project just for the API request structure.

### Query the Output

```bash
# View all GPU types with pricing
cat complete_catalog.json | jq '.[] | select(.gpu != null)'

# Show L40S options with pricing
cat complete_catalog.json | jq '.[] | select(.gpu.family == "l40s") | {preset, regions, price}'

# Compare pricing across GPU families
cat complete_catalog.json | jq -r '.[] | select(.gpu != null) | "\(.gpu.count)x \(.gpu.family): $\(.price.on_demand_per_hour)/hr"'
```

## Output Format

The output matches the LaunchPad API format with semantic IDs:

```json
{
  "id": "gpu-l40s-d-1gpu-16vcpu-96gb",
  "nebius_platform_id": "computeplatform-e00xxx",
  "cloud": "nebius",
  "platform": "gpu-l40s-d",
  "preset": "1gpu-16vcpu-96gb",
  "capacity": {
    "eu-north1": 1,
    "eu-west1": 0,
    "us-central1": 1
  },
  "regions": ["eu-north1", "us-central1"],
  "cpu": 16,
  "memory_gb": 96,
  "gpu": {
    "count": 1,
    "family": "l40s",
    "model": "NVIDIA L40S",
    "manufacturer": "NVIDIA",
    "memory_gb": 48,
    "interconnection_type": "pcie"
  },
  "storage": [
    {
      "type": "network-ssd",
      "size_min_gb": 50,
      "size_max_gb": 2560,
      "is_elastic": true
    }
  ],
  "system_arch": "amd64",
  "price": {
    "currency": "USD",
    "on_demand_per_hour": 2.45
  }
}
```

## Fields

- **capacity**: Map of region -> availability (1 = available, 0 = no quota)
- **regions**: List of regions where this preset has quota
- **gpu.family**: Lowercase GPU type (l40s, h100, h200)
- **storage.is_elastic**: Nebius supports elastic volumes (50GB-2560GB)
- **price**: From Nebius billing API (currently placeholder)

## Differences from SDK Output

The SDK `GetInstanceTypes()` returns **one entry per region** (like LaunchPad SDK does).
This utility **aggregates them** for easier visualization and comparison.

Both representations are valid - this is just for human readability and testing.

