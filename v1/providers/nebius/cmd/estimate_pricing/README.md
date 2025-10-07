# Nebius Pricing Estimator

This tool queries the **real Nebius Billing Calculator API** to get actual pricing for all instance types.

Based on: https://github.com/nebius/api/blob/main/nebius/billing/v1alpha1/calculator_service.proto

## Usage

```bash
# Set credentials
export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
export NEBIUS_TENANT_ID='tenant-xxx'
export NEBIUS_PROJECT_ID='project-xxx'

# Run the pricing estimator
go run ./cmd/estimate_pricing/main.go > pricing_estimates.json

# View all pricing
cat pricing_estimates.json | jq '.'

# Filter by GPU family
cat pricing_estimates.json | jq '.[] | select(.platform_name | contains("l40s"))'

# Show pricing summary
cat pricing_estimates.json | jq -r '.[] | "\(.platform_name) \(.preset_name): $\(.hourly_rate)/hr"'
```

## Output Format

```json
{
  "platform_id": "computeplatform-e00xxx",
  "platform_name": "gpu-l40s-d",
  "preset_name": "4gpu-128vcpu-768gb",
  "region": "eu-north1",
  "currency": "USD",
  "hourly_rate": 9.1376,
  "daily_rate": 219.3024,
  "monthly_rate": 6670.272,
  "annual_rate": 80043.264
}
```

## Actual Pricing (from Nebius Billing API)

### L40S GPU Pricing
- 1×L40S (16vcpu-96gb): **$1.82/hr** (~$1,326/month)
- 2×L40S (64vcpu-384gb): **$4.57/hr** (~$3,335/month)
- 4×L40S (128vcpu-768gb): **$9.14/hr** (~$6,670/month)

### H100 GPU Pricing  
- 1×H100 (16vcpu-200gb): **$2.95/hr** (~$2,153/month)
- 8×H100 (128vcpu-1600gb): **$23.60/hr** (~$17,228/month)

### H200 GPU Pricing
- 1×H200 (16vcpu-200gb): **$3.50/hr** (~$2,555/month)
- 8×H200 (128vcpu-1600gb): **$28.00/hr** (~$20,440/month)

### CPU Pricing
- 4vcpu-16gb: **$0.10/hr** (~$72/month)
- 8vcpu-32gb: **$0.20/hr** (~$145/month)
- 16vcpu-64gb: **$0.40/hr** (~$290/month)

## Combine with Instance Types

To create a complete view with both availability and pricing:

```bash
# Join instance types with pricing
jq -s '
  [.[0][] as $it | .[1][] as $price |
   if ($it.id | startswith($price.platform_id)) and ($it.preset == $price.preset_name)
   then $it + {
     price: {
       currency: $price.currency,
       on_demand_per_hour: $price.hourly_rate,
       estimated_monthly: $price.monthly_rate
     }
   }
   else empty end]
' instance_types_aggregated.json pricing_estimates.json > complete_catalog.json
```

This creates a complete instance type catalog with:
- ✅ Regional availability (capacity map)
- ✅ Instance specs (CPU, memory, GPU)
- ✅ **Real pricing** from Nebius Billing API
- ✅ Elastic storage details

## Implementation Details

The tool uses:
1. `nebius.compute.v1.Platform.List()` - Get all platforms
2. `nebius.billing.v1alpha1.Calculator.Estimate()` - Get pricing for each platform/preset
3. Minimal `CreateInstanceRequest` spec (only platform + preset required for pricing)

Pricing is calculated based on:
- Platform resources (CPU, memory, GPU)
- Network-SSD boot disk (50GB default)
- On-demand/unspecified offer type (no contract discounts)

**Note**: Pricing shown is for `eu-north1` region. Rates may vary slightly by region.



