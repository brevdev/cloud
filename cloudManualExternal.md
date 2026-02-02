# Brev Cloud Provider Integration Guide

**For Cloud Infrastructure Providers Integrating with Brev**

---

## Table of Contents

1. [Integration Overview](#1-integration-overview)
2. [How Brev Discovers Your Inventory](#2-how-brev-discovers-your-inventory)
3. [Instance Types: Your Compute Catalog](#3-instance-types-your-compute-catalog)
4. [Location Model](#4-location-model)
5. [GPU Normalization](#5-gpu-normalization)
6. [Credential and Authentication Model](#6-credential-and-authentication-model)
7. [Instance Lifecycle Operations](#7-instance-lifecycle-operations)
8. [SSH Connectivity](#8-ssh-connectivity)
9. [Firewall and Security Groups](#9-firewall-and-security-groups)
10. [Instance Metadata and Tags](#10-instance-metadata-and-tags)
11. [Error Handling and Status Reporting](#11-error-handling-and-status-reporting)
12. [Pricing and Billing](#12-pricing-and-billing)
13. [Common Questions](#13-common-questions)

---

## 1. Integration Overview

### What Does Integration Mean?

When you integrate with Brev, you're allowing Brev's control plane to:
1. **Sync** your available GPU instance types into Brev's catalog
2. **Provision** instances on your infrastructure via API calls
3. **Manage** instance lifecycle (start, stop, terminate) through your API
4. **Connect** to running instances via SSH to configure them

### What Brev Needs From You (Cloud Provider)

| Requirement | Purpose |
|-------------|---------|
| **Instance Type Listing API** | Discover your available instance types |
| **Instance Lifecycle APIs** | Create, get, start, stop, terminate |
| **API Credentials for Brev** | Authenticate Brev's calls to your API |
| **SSH Key Injection** | Accept SSH public key at VM creation |
| **SSH Access** | Control plane communication to VMs |

### Integration Architecture

### System Architecture Diagram

```
┌────────────────────────────────────────────────────────────────────────────────────┐
│                              Brev Control Plane (dev-plane)                        │
│                                                                                    │
│  ┌──────────────────────────────────┐    ┌──────────────────────────────────────┐  │
│  │         Syncer Layer             │    │     Instance Service Layer           │  │
│  │    (Continuous Reconciliation)   │    │       (User-Triggered Actions)       │  │
│  │                                  │    │                                      │  │
│  │  ┌────────────────────────────┐  │    │  ┌────────────────────────────────┐  │  │
│  │  │  InstanceTypeSyncer        │  │    │  │  Instance Lifecycle            │  │  │
│  │  │  ─────────────────────     │  │    │  │  ─────────────────────         │  │  │
│  │  │  Calls:                    │  │    │  │  Calls:                        │  │  │
│  │  │  • GetInstanceTypes()      │  │    │  │  • CreateInstance()            │  │  │
│  │  │  • GetLocations()          │  │    │  │  • TerminateInstance()         │  │  │
│  │  │  • GetInstanceTypePollTime │  │    │  │  • StopInstance()              │  │  │
│  │  │                            │  │    │  │  • StartInstance()             │  │  │
│  │  │  Interval: 1-5 min         │  │    │  │                                │  │  │
│  │  └────────────┬───────────────┘  │    │  └──────────────┬─────────────────┘  │  │
│  │               │                  │    │                 │                    │  │
│  │  ┌────────────┴───────────────┐  │    │  ┌──────────────┴─────────────────┐  │  │
│  │  │  InstanceSyncer            │  │    │  │  Instance State & Queries      │  │  │
│  │  │  ─────────────────────     │  │    │  │  ─────────────────────         │  │  │
│  │  │  Calls:                    │  │    │  │  Calls:                        │  │  │
│  │  │  • ListInstances()         │  │    │  │  • GetInstance()               │  │  │
│  │  │                            │  │    │  │  • ListInstances()             │  │  │
│  │  │  Interval: 5 sec           │  │    │  │  • AddFirewallRulesToInstance  │  │  │
│  │  └────────────┬───────────────┘  │    │  │  • ResizeInstanceVolume()      │  │  │
│  │               │                  │    │  │  • UpdateInstanceTags()        │  │  │
│  └───────────────┼──────────────────┘    │  └──────────────┬─────────────────┘  │  │
│                  │                       └─────────────────┼────────────────────┘  │
│                  │                                         │                       │
└──────────────────┼─────────────────────────────────────────┼───────────────────────┘
                   │                                         │
                   │       ┌─────────────────────────────────┘
                   │       │
                   ▼       ▼
┌────────────────────────────────────────────────────────────────────────────────────┐
│                           CLOUD SDK (v1) - This Repo                               │
│                                                                                    │
│  ┌──────────────────────────────────────────────────────────────────────────────┐  │
│  │                         CloudClient Interface                                │  │
│  │  (Composed of: CloudCredential, CloudBase, CloudQuota, CloudStopStart,       │  │
│  │   CloudReboot, CloudResizeVolume, CloudModifyFirewall, CloudInstanceTags...) │  │
│  └──────────────────────────────────────────────────────────────────────────────┘  │
│                                                                                    │
│  ┌──────────────────────────────────────────────────────────────────────────────┐  │
│  │                        Provider Implementations                              │  │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐            │  │
│  │  │ Lambda   │ │ Fluidstk │ │ Shadefrm │ │  Nebius  │ │  Your    │            │  │
│  │  │ Labs     │ │          │ │          │ │          │ │ Provider │   • • •    │  │
│  │  └────┬─────┘ └────┬─────┘ └────┬─────┘ └────┬─────┘ └────┬─────┘            │  │
│  └───────┼────────────┼────────────┼────────────┼────────────┼──────────────────┘  │
│          │            │            │            │            │                     │
└──────────┼────────────┼────────────┼────────────┼────────────┼─────────────────────┘
           │            │            │            │            │
           ▼            ▼            ▼            ▼            ▼
┌────────────────────────────────────────────────────────────────────────────────────┐
│                            CLOUD PROVIDER APIs                                     │
│                                                                                    │
│  Each provider's native REST/gRPC API for instance management                      │
└────────────────────────────────────────────────────────────────────────────────────┘
```

## 2. How Brev Discovers Your Inventory

### The Instance Type Syncer

Brev runs a **continuous synchronization process** that periodically queries your API to understand what compute is available.

**Sync Behavior:**
- Polls your instance type listing API at a configurable interval you define via `GetInstanceTypePollTime()` (default: 1 minute; existing implementations use 1-5 minutes depending on provider needs)
- Compares current catalog to previous state
- Updates availability, pricing, and specs as they change
- Marks types as unavailable when removed from your API
- Adds new types when they appear

### What We Query

We need an API endpoint that returns your available instance types. For each type, we map your data to the `v1.InstanceType` struct (defined in `cloud/v1/instancetype.go`):

**Core Instance Type Fields:**

| Struct Field | Type | Description | Example |
|--------------|------|-------------|---------|
| `Type` | `string` | Your internal type name | `"gpu_1x_a100_80gb_sxm4"` |
| `Location` | `string` | Region identifier | `"us-west-1"` |
| `VCPU` | `int32` | vCPU count | `128` |
| `MemoryBytes` | `Bytes` | RAM (use `v1.NewBytes()`) | `v1.NewBytes(1024, v1.Gibibyte)` |
| `BasePrice` | `*currency.Amount` | Hourly price in USD | `currency.NewAmountFromInt64(3200, "USD")` (= $32.00/hr) |
| `IsAvailable` | `bool` | Currently launchable | `true` |

**GPU Details (`SupportedGPUs []GPU`):**

| Struct Field | Type | Description | Example |
|--------------|------|-------------|---------|
| `Count` | `int32` | Number of GPUs | `8` |
| `Name` | `string` | GPU model name | `"A100"` |
| `MemoryBytes` | `Bytes` | VRAM per GPU | `v1.NewBytes(80, v1.Gibibyte)` |
| `NetworkDetails` | `string` | Interconnect type | `"SXM4"`, `"PCIe"` |
| `Manufacturer` | `Manufacturer` | GPU vendor | `v1.ManufacturerNVIDIA` |

**Storage Details (`SupportedStorage []Storage`):**

| Struct Field | Type | Description | Example |
|--------------|------|-------------|---------|
| `SizeBytes` | `Bytes` | Disk size | `v1.NewBytes(2000, v1.Gibibyte)` |
| `Type` | `string` | Storage type | `"ssd"`, `"nvme"` |
| `PricePerGBHr` | `*currency.Amount` | Additional storage cost | `nil` (if included in base price) |

**Example: Converting Provider Data to `v1.InstanceType`**

From Lambda Labs implementation (`cloud/v1/providers/lambdalabs/instancetype.go`):

```go
it := v1.InstanceType{
    Location:      location,
    Type:          instType.Name,                                           // "gpu_1x_a100_80gb_sxm4"
    SupportedGPUs: []v1.GPU{{
        Count:       8,
        Name:        "A100",
        MemoryBytes: v1.NewBytes(80, v1.Gibibyte),
        NetworkDetails: "SXM4",
        Manufacturer: v1.ManufacturerNVIDIA,
    }},
    SupportedStorage: []v1.Storage{{
        Type:      "ssd",
        SizeBytes: v1.NewBytes(instType.Specs.StorageGib, v1.Gibibyte),
    }},
    VCPU:        instType.Specs.Vcpus,
    MemoryBytes: v1.NewBytes(instType.Specs.MemoryGib, v1.Gibibyte),
    BasePrice:   &amount,
    IsAvailable: isAvailable,
    Provider:    CloudProviderID,
    Cloud:       CloudProviderID,
}
it.ID = v1.MakeGenericInstanceTypeID(it)  // Generate ID using helper (or set your own)
```

### API Type Declaration

When implementing the Cloud SDK, you declare how Brev's control plane should query your integration via `GetAPIType()`:

| API Type | Meaning | Control Plane Behavior |
|----------|---------|------------------------|
| `APITypeGlobal` | Your `GetInstanceTypes()` returns all regions in one call | Brev calls once with `locations = ["all"]` |
| `APITypeLocational` | Your `GetInstanceTypes()` is region-scoped | Brev iterates over `GetLocations()` results |

**You handle the mapping internally.** The SDK doesn't call your API directly—your implementation does. Whether your cloud's native API is regional, global, or something else entirely, you write the conversion logic in `GetInstanceTypes()`.

**Example: Global API (Lambda Labs)**
Lambda Labs' API returns all instance types with regional availability embedded. The SDK implementation fetches once and expands to per-region `v1.InstanceType` entries:

```go
// Simplified from cloud/v1/providers/lambdalabs/instancetype.go
func (c *LambdaLabsClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
    resp, _ := c.client.InstanceTypes(ctx)  // Single API call returns all types
    
    // Expand each type to all its available regions
    for _, instType := range resp.Data {
        for _, region := range locations {
            isAvailable := slices.Contains(instType.RegionsWithCapacityAvailable, region.Name)
            instanceTypes = append(instanceTypes, convertToV1(region.Name, instType, isAvailable))
        }
    }
    return instanceTypes, nil
}
```

**Example: Locational API (Nebius)**
Nebius requires per-region quota checks. The SDK implementation iterates regions internally:

```go
// Simplified from cloud/v1/providers/nebius/instancetype.go
func (c *NebiusClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
    platforms, _ := c.sdk.Compute().Platform().List(ctx, c.projectID)
    
    for _, location := range locations {
        // Check quota per-region
        isAvailable := c.checkQuotaAvailability(platform, location.Name, quotaMap)
        instanceTypes = append(instanceTypes, convertToV1(location.Name, platform, isAvailable))
    }
    return instanceTypes, nil
}
```

**Key point:** You decide how to call your cloud's API. Brev only cares that `GetInstanceTypes()` returns properly formatted `v1.InstanceType` entries with accurate `Location` and `IsAvailable` fields.

---

## 3. Instance Types: Your Compute Catalog

### What Is an Instance Type to Brev?

Brev treats compute as **inventory**. Each instance type represents a distinct compute configuration in your catalog. Users browse your instance types filtered by GPU, region, price, and availability.

### The Canonical Instance Type Model

When we ingest your instance types, we normalize them to the `v1.InstanceType` struct. Here are the key fields (see `cloud/v1/instancetype.go` for the complete definition):

| Field | Type | Description |
|-------|------|-------------|
| `ID` | `InstanceTypeID` | Stable, unique identifier (you define the format—see below) |
| `Cloud` | `string` | Your cloud identifier (e.g., `"lambdalabs"`, `"crusoe"`) |
| `Provider` | `string` | Provider identifier (often same as `Cloud`) |
| `Type` | `string` | Your native type name |
| `Location` | `string` | Primary region identifier |
| `SubLocation` | `string` | Availability zone (optional; helper uses `"noSub"` if empty) |
| `AvailableAzs` | `[]string` | All zones where this type is available |
| `SupportedGPUs` | `[]GPU` | GPU details (see `GPU` struct below) |
| `VCPU` | `int32` | vCPU count |
| `MemoryBytes` | `Bytes` | RAM (use `v1.NewBytes()` helper) |
| `SupportedStorage` | `[]Storage` | Storage options (see `Storage` struct) |
| `BasePrice` | `*currency.Amount` | Hourly price in USD |
| `IsAvailable` | `bool` | Currently launchable |
| `Stoppable` | `bool` | Can instances be stopped/resumed |
| `Rebootable` | `bool` | Can instances be rebooted |

**The `GPU` struct** (`cloud/v1/instancetype.go`):

| Field | Type | Description |
|-------|------|-------------|
| `Count` | `int32` | Number of GPUs |
| `Name` | `string` | GPU model name (e.g., `"A100"`, `"H100"`) |
| `Type` | `string` | Full GPU type (e.g., `"A100.SXM4"`) |
| `MemoryBytes` | `Bytes` | VRAM per GPU |
| `MemoryDetails` | `string` | Memory type: `"HBM"`, `"GDDR"`, etc. |
| `NetworkDetails` | `string` | Interconnect: `"PCIe"`, `"SXM4"`, `"SXM5"` |
| `Manufacturer` | `Manufacturer` | `ManufacturerNVIDIA`, `ManufacturerIntel`, etc. |

**The `Storage` struct** (`cloud/v1/storage.go`):

| Field | Type | Description |
|-------|------|-------------|
| `Count` | `int32` | Number of disks |
| `SizeBytes` | `Bytes` | Disk size |
| `Type` | `string` | Storage type (e.g., `"ssd"`, `"nvme"`) |
| `PricePerGBHr` | `*currency.Amount` | Additional storage cost (if applicable) |
| `IsEphemeral` | `bool` | Lost on stop/terminate |

### Instance Type ID

The `ID` field must be a **stable, unique identifier** for each instance type across all regions. You control the format.

**Requirements:**
- **Stable**: The same instance type must return the same ID on every sync
- **Unique**: No two instance types can share an ID
- **Deterministic**: IDs must not change between API calls

**Option 1: Use the Helper Function**

The SDK provides `MakeGenericInstanceTypeID()` which generates IDs using this pattern:

```
{location}-{subLocation}-{type}
```

If your instance type has no sublocation, the helper uses `"noSub"` as a placeholder.

```go
// Set all fields first, then call the helper at the END
it := v1.InstanceType{
    Location: "us-west-1",
    Type:     "gpu_1x_a100",
    // ... other fields
}
it.ID = v1.MakeGenericInstanceTypeID(it)  // Result: "us-west-1-noSub-gpu_1x_a100"
```

**Option 2: Define Your Own Format**

If you prefer a different ID format, set `ID` directly:

```go
// Shadeform uses: {cloud}_{instanceType}_{region}
it := v1.InstanceType{
    ID:       v1.InstanceTypeID("massedcompute_L40_desmoines-usa-1"),
    Location: "desmoines-usa-1",
    Type:     "massedcompute_L40",
    // ... other fields
}
```

**Why Stability Matters:**

Brev uses this ID to track inventory and match provisioning requests. If your IDs change between syncs, Brev loses the ability to correlate instance types correctly.

### CRITICAL: ID Consistency Between InstanceType and Instance

> **Warning**: This is the most common cause of integration failures. Instance types may sync successfully but instances fail to provision or appear "orphaned."

When Brev provisions an instance, it looks up the corresponding instance type using the instance's `InstanceTypeID`. **These IDs must match exactly.**

**A Common Problem:**

The SDK has two helper functions that generate IDs differently:

| Function | Used For | SubLocation Source |
|----------|----------|-------------------|
| `MakeGenericInstanceTypeID()` | InstanceType structs | `AvailableAzs[0]` (first AZ) |
| `MakeGenericInstanceTypeIDFromInstance()` | Instance structs | `SubLocation` field |

If `AvailableAzs[0]` and `SubLocation` don't match, the IDs diverge and lookup fails.

**The Mistakes:**

```go
// WRONG - Manually setting InstanceTypeID
inst := &v1.Instance{
    InstanceType:   "gpu-h100-8x",
    InstanceTypeID: v1.InstanceTypeID("gpu-h100-8x"),  // BUG: Missing location!
}

// WRONG - Inconsistent SubLocation vs AvailableAzs
instanceType := v1.InstanceType{
    Location:     "us-east-1",
    SubLocation:  "us-east-1a",      // Set to "us-east-1a"
    AvailableAzs: []string{"us-east-1b"},  // But AZs has "us-east-1b"!
}
```

**The Fix:**

1. **For InstanceType**: Set all fields first, then call `MakeGenericInstanceTypeID()` at the END
2. **For Instance**: Set all fields first, then call `MakeGenericInstanceTypeIDFromInstance()` at the END
3. **Ensure consistency**: If you set both `SubLocation` and `AvailableAzs`, make sure `SubLocation == AvailableAzs[0]`

```go
// CORRECT - InstanceType
it := v1.InstanceType{
    Location:     "us-east-1",
    AvailableAzs: []string{"us-east-1a"},
    Type:         "gpu-h100-8x",
    // ... other fields
}
it.ID = v1.MakeGenericInstanceTypeID(it)  // LAST

// CORRECT - Instance
inst := &v1.Instance{
    Location:     "us-east-1",
    SubLocation:  "us-east-1a",  // Matches the AZ
    InstanceType: "gpu-h100-8x",
    // ... other fields
}
inst.InstanceTypeID = v1.MakeGenericInstanceTypeIDFromInstance(*inst)  // LAST
```

**Symptoms of ID Mismatch:**
- Instance types sync successfully but don't appear in the Brev catalog
- `CreateInstance` succeeds but subsequent operations fail
- "instance type not found" errors during provisioning
- Instances appear "orphaned" (no associated instance type)

### Validating Your Instance Type IDs

The SDK provides validation functions to catch ID generation issues early. **Run these in your test suite:**

**1. `ValidateStableInstanceTypeIDs`** - Ensures your instance type IDs are stable and unique:
```go
// In your validation tests
err := v1.ValidateStableInstanceTypeIDs(ctx, client, stableIDs)
require.NoError(t, err, "ValidateStableInstanceTypeIDs should pass")
```

This validates:
- Each instance type ID is unique (no duplicates)
- Your designated stable IDs exist in the current instance types
- All instance types have required properties (base price, storage pricing)

**2. `ValidateCreateInstance`** - Validates that instance and instance type IDs match:
```go
// In your validation tests
instance, err := v1.ValidateCreateInstance(ctx, client, attrs, selectedType)
require.NoError(t, err, "ValidateCreateInstance should pass")
```

This validates (among other things):
- `instance.InstanceTypeID == selectedType.ID` — **catches ID generation mismatches**
- `instance.RefID` matches the provided RefID
- Location and instance type fields are consistent

> **Why this matters:** If `MakeGenericInstanceTypeID()` and `MakeGenericInstanceTypeIDFromInstance()` produce different IDs for the same logical type, the control plane cannot correlate instances with their types. `ValidateCreateInstance` catches this.

See [`internal/validation/suite.go`](internal/validation/suite.go) for the full validation test suite you can use as a reference.

## 4. Location Model

### The Location Hierarchy

Brev uses a three-level location model to represent where compute resources exist:

| Level | Field | Description | Example |
|-------|-------|-------------|---------|
| **Region** | `Location` | Primary geographic region | `"us-west-1"`, `"europe-west4"` |
| **Availability Zone** | `SubLocation` | Specific zone within a region | `"us-west-1a"`, `"europe-west4-b"` |
| **Available Zones** | `AvailableAzs` | All zones where this type can launch | `["us-west-1a", "us-west-1b"]` |

> **Note:** The distinction between these fields can be confusing. `Location` is the region, `SubLocation` is a specific zone (used for instances), and `AvailableAzs` lists all zones where an instance type is available (used for instance types).

### The Location Struct

When implementing `GetLocations()`, you return a list of `Location` structs (defined in `cloud/v1/location.go`):

| Field | Type | Description |
|-------|------|-------------|
| `Name` | `string` | Region identifier (acts as the ID) |
| `Description` | `string` | Human-readable name |
| `Available` | `bool` | Whether the region is currently operational |
| `Endpoint` | `string` | API endpoint for this region (if applicable) |
| `Priority` | `int` | Preference order for region selection |
| `Country` | `string` | ISO 3166-1 alpha-3 country code |

### Availability on Instance Types

Availability is tracked **per instance type** using two fields on the `InstanceType` struct:

| Field | Type | Meaning |
|-------|------|---------|
| `IsAvailable` | `bool` | Whether this type can currently be launched |
| `AvailableAzs` | `[]string` | Which availability zones have capacity |

**Interpreting Availability:**
- `IsAvailable: true` + `AvailableAzs: ["us-west-1a", "us-west-1b"]` = Can launch in either AZ
- `IsAvailable: false` = Type exists but is currently out of stock or disabled
- Empty `AvailableAzs` with `IsAvailable: true` = Region-level availability only (no AZ granularity)

---

## 5. GPU Normalization

### The GPU Struct

The Cloud SDK represents GPUs with these fields:

```go
type GPU struct {
    Name           string           // Base model: "H100", "A100", "L40S"
    Count          int32            // Number of GPUs
    Memory         units.Base2Bytes // VRAM per GPU (deprecated, use MemoryBytes)
    MemoryBytes    Bytes            // VRAM per GPU in structured format
    MemoryDetails  string           // Memory type: "HBM2", "HBM3", "HBM2e", "GDDR"
    NetworkDetails string           // Form factor: "PCIe", "SXM", "SXM4", "SXM5"
    Manufacturer   Manufacturer     // "NVIDIA", "AMD", "Intel"
    Type           string           // Optional: original type identifier
}
```

### Implementer Responsibility

**You are responsible for normalizing GPU data.** Brev does not automatically parse GPU descriptions. Your `GetInstanceTypes` must populate the `GPU` struct.

| Field | Example | Notes |
|-------|---------|-------|
| `Name` | `"H100"`, `"A100"` | Base model, uppercase |
| `Count` | `8` | GPUs per instance |
| `MemoryBytes` | `v1.NewBytes(80, v1.Gibibyte)` | VRAM per GPU |
| `NetworkDetails` | `"SXM4"`, `"PCIe"` | Form factor |
| `Manufacturer` | `"NVIDIA"` |

### Provider Examples

**Lambda Labs** (`cloud/v1/providers/lambdalabs/instancetype.go:parseGPUFromDescription`)

Parses `"8x A100 (40 GB SXM4)"` using regex:

```go
gpu.Count = int32(count)           // from (\d+)x
gpu.Name = nameStr                 // from x (.*?) \(
gpu.MemoryBytes = v1.NewBytes(v1.BytesValue(memoryGiB), v1.Gibibyte)
gpu.NetworkDetails = networkDetails // remainder after "GB"
gpu.Manufacturer = "NVIDIA"
```

**Launchpad** (`cloud/v1/providers/launchpad/instancetype.go:launchpadGpusToGpus`)

Maps structured API fields:

```go
gpus[i] = v1.GPU{
    Name:           strings.ToUpper(gp.Family),
    Count:          gp.Count,
    MemoryBytes:    v1.NewBytes(v1.BytesValue(gp.MemoryGb), v1.Gigabyte),
    NetworkDetails: string(gp.InterconnectionType),
    Manufacturer:   v1.GetManufacturer(gp.Manufacturer),
}
```

### Key Points

- `Name`: base model only (`"H100"` not `"NVIDIA H100 80GB"`)
- `NetworkDetails`: `"SXM"`, `"SXM4"`, `"SXM5"`, or `"PCIe"`
- `Manufacturer`: always set to `"NVIDIA"`

---

## 6. Credential and Authentication Model

### How Brev Authenticates to Your API

Brev stores credentials for your cloud provider and uses them to make API calls. This is a direct relationship between **Brev's control plane** and **your cloud API**.

### What You Need to Provide

| Requirement | Details |
|-------------|---------|
| **API Credentials** | A JSON-serializable Go struct containing your authentication fields (API key, token, service account, etc.) |
| **Authentication Endpoint** | How Brev authenticates (API key header, OAuth, etc.) |

### Credential Storage Model

Credentials are stored in Brev's control plane database as **raw JSON** (`json.RawMessage`). This means your credential struct must be JSON-serializable with proper struct tags.

**How it works:**
1. **You define** a credential struct with JSON tags for each field
2. **Brev stores** the struct as raw JSON bytes in the database (encrypted at rest)
3. **Brev deserializes** the JSON back into your struct type when making API calls

**Example credential struct:**

```go
type MyProviderCredential struct {
    RefID  string            // Set by Brev (the cloud_cred ID)
    APIKey string `json:"api_key"`
    Region string `json:"region,omitempty"`  // Optional fields use omitempty
}
```

**Key requirements:**
- All fields you need serialized must have `json:"field_name"` tags
- The `RefID` field is set by Brev after storage (it's the database record ID)
- Use `json:"...,omitempty"` for optional fields
- The struct must implement the `CloudCredential` interface

### Credential Exchange Process

1. **You provide** API credentials to Brev during integration setup
2. **Brev stores** credentials securely (encrypted at rest)
3. **Brev uses** credentials to call your API for sync and provisioning

### Credential Types

Providers define their own credential struct with whatever fields they need. The struct fields use JSON tags that determine the field names in the stored JSON.

| Provider | Struct Fields | JSON Fields |
|----------|---------------|-------------|
| **Lambda Labs** | `APIKey string` | `api_key` |
| **Shadeform** | `APIKey string` | `api_key` |
| **FluidStack** | `APIKey string` | `api_key` |
| **AWS** | `AccessKeyID`, `SecretAccessKey` | `access_key_id`, `secret_access_key` |
| **Nebius** | `ServiceAccountKey`, `TenantID` | `service_account_key`, `tenant_id` |
| **Launchpad** | `APIToken`, `APIURL` | `api_token`, `api_url` |

**Complete credential struct example (from Launchpad):**

```go
type LaunchpadCredential struct {
    RefID    string            // Not serialized - set by Brev after storage
    APIToken string `json:"api_token"`
    APIURL   string `json:"api_url"`
}

var _ v1.CloudCredential = &LaunchpadCredential{}  // Compile-time interface check

func (c *LaunchpadCredential) Validate() error {
    return validation.ValidateStruct(c,
        validation.Field(&c.APIToken, validation.Required),
        validation.Field(&c.APIURL, validation.Required),
    )
}
```

Your credential struct must implement the `CloudCredential` interface, which requires these methods:

```go
type CloudCredential interface {
    MakeClient(ctx context.Context, location string) (CloudClient, error)
    GetTenantID() (string, error)
    GetReferenceID() string
    GetAPIType() APIType
    GetCapabilities(ctx context.Context) (Capabilities, error)
    GetCloudProviderID() CloudProviderID
}
```

### SSH Keys (Separate from API Credentials)

SSH keys are passed at instance creation time via the `PublicKey` field in `CreateInstanceAttrs`.

Your implementation must:
1. Accept this public key in your create instance API
2. Install it in the VM's default user `~/.ssh/authorized_keys` before the instance becomes accessible

Brev manages SSH keys per user. The public key provided in `CreateInstanceAttrs.PublicKey` belongs to the user, and the control plane retains the corresponding private key to connect after creation.

---

## 7. Instance Lifecycle Operations

This section describes each lifecycle operation, its requirements, and expected behavior. Not all operations are required—providers declare their capabilities via `GetCapabilities()`.

### Lifecycle States

The SDK defines these states in `LifecycleStatus` (from `cloud/v1/instance.go`):

| State | Meaning |
|-------|---------|
| `pending` | Create initiated, VM provisioning |
| `running` | Instance is up with a public IP |
| `stopping` | Stop requested, shutting down |
| `stopped` | Powered off, storage preserved |
| `suspending` | Suspend requested |
| `suspended` | Hibernated state |
| `terminating` | Terminate requested |
| `terminated` | Instance destroyed |
| `failed` | Provisioning or operation failed |

### Create Instance (Required)

**Interface:** `CloudCreateTerminateInstance.CreateInstance(ctx, CreateInstanceAttrs) (*Instance, error)`

**Contract:**
- On success: Return an `*Instance` with a valid `CloudID`. The instance must exist in your system.
- On error: Return an error **and ensure no instance was created**. Brev will not attempt cleanup on errors.

**Key input fields from `CreateInstanceAttrs`:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `RefID` | `string` | Yes | Brev's reference ID; use for idempotency |
| `InstanceType` | `string` | Yes | Your instance type name |
| `Location` | `string` | Yes | Region to launch in |
| `SubLocation` | `string` | No | Specific availability zone |
| `PublicKey` | `string` | Yes | SSH public key (OpenSSH format) |
| `Name` | `string` | No | Display name for the instance |
| `ImageID` | `string` | No | OS image; use your default if empty |
| `DiskSize` | `units.Base2Bytes` | No | Boot disk size |
| `FirewallRules` | `FirewallRules` | No | Ports to open (SSH port is always required) |
| `Tags` | `Tags` | No | Key-value metadata |
| `UserDataBase64` | `string` | No | Cloud-init or startup script |

**Key output fields on `Instance`:**

| Field | When Required | Description |
|-------|---------------|-------------|
| `CloudID` | Always | Your unique instance identifier |
| `Status.LifecycleStatus` | Always | Current state (`pending` or `running`) |
| `Location` | Always | Region where launched |
| `InstanceType` | Always | Instance type that was provisioned |
| `PublicIP` | When running | Public IPv4 for SSH access |
| `SSHUser` | Always | Username for SSH (e.g., `ubuntu`, `root`) |
| `SSHPort` | Always | SSH port (typically `22`) |
| `RefID` | Always | Echo back the input `RefID` |

**Example flow (from Lambda Labs implementation):**

```go
// 1. Register the SSH key with your API
keyPairResp, err := c.addSSHKey(ctx, openapi.AddSSHKeyRequest{
    Name:      attrs.RefID,
    PublicKey: &attrs.PublicKey,
})

// 2. Launch the instance with the key
resp, err := c.launchInstance(ctx, openapi.LaunchInstanceRequest{
    RegionName:       attrs.Location,
    InstanceTypeName: attrs.InstanceType,
    SshKeyNames:      []string{keyPairName},
})

// 3. Return instance details
return c.GetInstance(ctx, v1.CloudProviderInstanceID(resp.Data.InstanceIds[0]))
```

### Terminate Instance (Required)

**Interface:** `CloudCreateTerminateInstance.TerminateInstance(ctx, instanceID) error`

**Contract:**
- Initiate instance termination. Storage may or may not be preserved (provider-dependent).
- Return `nil` on success, even if the instance is already terminated.
- The instance should eventually reach `terminated` state.

**Idempotency:** Should succeed if called multiple times on the same instance.

### Stop Instance (Optional)

**Capability:** `CapabilityStopStartInstance`

**Interface:** `CloudStopStartInstance.StopInstance(ctx, instanceID) error`

**Contract:**
- Power off the instance while preserving storage.
- Return `nil` once the stop operation is initiated.
- Instance should transition: `running` → `stopping` → `stopped`

**When to implement:** Only if your platform supports instances that can stop and preserve storage. Lambda Labs does not support this, but Nebius does.

### Start Instance (Optional)

**Capability:** `CapabilityStopStartInstance`

**Interface:** `CloudStopStartInstance.StartInstance(ctx, instanceID) error`

**Contract:**
- Power on a previously stopped instance.
- Return `nil` once the start operation is initiated.
- Instance should transition: `stopped` → `pending` → `running`

**Note:** If you implement `StopInstance`, you must also implement `StartInstance`.

### Stop/Start: Three Levels of Control

Stop/start support is controlled at three levels:

| Level | What to Set | Purpose |
|-------|-------------|---------|
| **Provider Capability** | `CapabilityStopStartInstance` in `GetCapabilities()` | Indicates your API supports stop/start operations |
| **Instance Type** | `InstanceType.Stoppable = true/false` | Indicates whether this instance type can be stopped (e.g., spot instances typically cannot) |
| **Instance** | `Instance.Stoppable = true/false` | Indicates whether this specific instance can be stopped |

**Example - Nebius (supports stop/start):**
```go
// In GetCapabilities()
v1.CapabilityStopStartInstance,  // API supports it

// In GetInstanceTypes() - instance type level
instanceType := v1.InstanceType{
    Stoppable: true,  // This type supports stop/start
    // ...
}

// In GetInstance()/CreateInstance() - instance level
instance := v1.Instance{
    Stoppable: true,  // This instance can be stopped
    // ...
}
```

**Example - Lambda Labs (no stop/start support):**
```go
// In GetCapabilities()
// CapabilityStopStartInstance NOT included

// In GetInstanceTypes()
instanceType := v1.InstanceType{
    Stoppable: false,  // Cannot be stopped
    // ...
}

// In GetInstance()/CreateInstance()
instance := v1.Instance{
    Stoppable: false,  // Cannot be stopped
    // ...
}
```

The control plane checks all three levels before allowing a stop/start operation. If any level indicates `false`, the operation won't be attempted.


### Get Instance (Required)

**Interface:** `CloudInstanceReader.GetInstance(ctx, instanceID) (*Instance, error)`

**Contract:**
- Return current state of the instance.
- Return `ErrResourceNotFound` if the instance doesn't exist.

### List Instances (Required)

**Interface:** `CloudInstanceReader.ListInstances(ctx, ListInstancesArgs) ([]Instance, error)`

**Contract:**
- Return all instances matching the filter criteria.
- Used by the Instance Syncer to reconcile state (called every ~5 seconds).

### Capability Declaration

Your credential's `GetCapabilities()` must return the capabilities you support:

```go
func (c *MyCredential) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
    return v1.Capabilities{
        v1.CapabilityCreateInstance,           // Required
        v1.CapabilityTerminateInstance,        // Required
        v1.CapabilityCreateTerminateInstance,  // Required (composite)
        // Optional:
        v1.CapabilityStopStartInstance,        // If you support stop/start
        v1.CapabilityRebootInstance,           // If you support reboot
        v1.CapabilityTags,                     // If your API supports instance tags/labels (see Section 10)
        v1.CapabilityModifyFirewall,           // If you support dynamic firewall rules
        v1.CapabilityResizeInstanceVolume,     // If you support volume resizing
    }, nil
}
```

Brev checks capabilities before calling optional methods. If you don't declare a capability, Brev won't attempt that operation.

> **Note on `CapabilityTags`:** This capability is optional, but `RefID` and `CloudCredRefID` data is **required** regardless. If your API doesn't support tags, you must use an alternative mechanism to store and retrieve this data. See [Section 10: Instance Metadata and Tags](#10-instance-metadata-and-tags) for details and examples.

---

## 8. SSH Connectivity

### Core Requirement

Brev's control plane must be able to connect to your instances via SSH using the provided keys. This is the **only hard requirement** for network connectivity.

After your VM is running, Brev connects via SSH to:

1. **Configure the environment**: Install Brev agent, set up development tools
2. **Enable connections**: Set up tunnels and connection paths for users
3. **Manage instance**: Execute commands, transfer files, health checks

### What You Provide at Launch

When provisioning, we pass:
- **SSH public key**: Key to install in `authorized_keys` (via `CreateInstanceAttrs.PublicKey`)
- **Firewall rules**: Ports to open (see Section 9)

### Instance Requirements

Your instances must return these fields so Brev can connect:

| Field | Required | Description |
|-------|----------|-------------|
| `SSHUser` | Yes | Username for SSH (e.g., `ubuntu`, `root`, `ec2-user`) |
| `SSHPort` | Yes | SSH port (commonly `22`, but can be any port) |
| `PublicIP` | Yes | Publicly routable address for SSH connection |

**Note:** While `PublicIP` is the required field, public routing via DNS also works in practice. The key requirement is that Brev can reach your instance over SSH.

### SSH User

Brev connects as the default user your image provides:

| Image | Default User |
|-------|--------------|
| Ubuntu | `ubuntu` |
| Debian | `admin` or `debian` |
| Amazon Linux | `ec2-user` |
| Custom | Whatever you configure |

### Runtime Requirements

| Requirement | Details |
|-------------|---------|
| **SSHD running** | On the port specified by `Instance.SSHPort` |
| **Port publicly reachable** | No NAT or firewall blocking inbound SSH |
| **Key installed** | The public key from `CreateInstanceAttrs.PublicKey` in `authorized_keys` |

---

## 9. Firewall and Security Groups

**Can you dynamically expose ports at instance creation?** Yes, if you support user-data or have a native firewall API.

**Can you modify firewall rules after creation without SSH/reboot?** Only if you have a native API. Most GPU clouds don't.


### SDK Structures

```go
type FirewallRules struct {
    IngressRules []FirewallRule
    EgressRules  []FirewallRule
}

type FirewallRule struct {
    FromPort int32
    ToPort   int32
    IPRanges []string // CIDR notation
}
```

Passed via `CreateInstanceAttrs.FirewallRules`.

### If You Have a Native API

Use it. Implement `CloudModifyFirewall` for post-creation changes:

```go
type CloudModifyFirewall interface {
    AddFirewallRulesToInstance(ctx context.Context, args AddFirewallRulesToInstanceArgs) error
    RevokeSecurityGroupRules(ctx context.Context, args RevokeSecurityGroupRuleArgs) error
}
```

Add `CapabilityModifyFirewall` to your capabilities.

### If You Only Have User-Data

Inject UFW commands at boot. See `cloud/v1/providers/shadeform/ufw.go`.

```go
// Core pattern
commands := []string{
    "ufw --force reset",
    "ufw default deny incoming",
    "ufw default allow outgoing",
    "ufw allow 22/tcp",
}
for _, rule := range firewallRules.IngressRules {
    for _, cidr := range rule.IPRanges {
        commands = append(commands, fmt.Sprintf("ufw allow in from %s to any port %d", cidr, rule.FromPort))
    }
}
commands = append(commands, "ufw --force enable")

// Base64 encode and pass as user-data
script := strings.Join(commands, "\n")
encoded := base64.StdEncoding.EncodeToString([]byte(script))
```

**Do not** implement `CloudModifyFirewall`. Return `ErrNotImplemented`.

### If You Only Have IP Allowlists

See `cloud/v1/providers/launchpad/instance_create.go`. You can only restrict by source IP, not port. Extract `/32`s from the rules and pass to your API:

```go
ips := []string{}
for _, rule := range firewallRules.IngressRules {
    for _, cidr := range rule.IPRanges {
        _, ipNet, _ := net.ParseCIDR(cidr)
        ones, bits := ipNet.Mask.Size()
        if ones == bits { // /32 only
            ips = append(ips, ipNet.IP.String())
        }
    }
}
```

---

## 10. Instance Metadata and Tags

Brev uses metadata to track and correlate instances. The control plane requires certain data to be persisted with instances and retrievable later.

### Required Instance Data

These values **MUST** be stored with the instance and returned in `GetInstance`/`ListInstances`:

| Field | Purpose |
|-------|---------|
| `RefID` | Instance correlation and idempotency (passed in `CreateInstanceAttrs.RefID`) |
| `CloudCredRefID` | Identifies which credential created the instance (from `GetReferenceID()`) |

### The `CapabilityTags` Capability

If your cloud provider's API supports instance tagging/labeling, declare `v1.CapabilityTags` in your capabilities:

```go
func (c *MyCredential) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
    return v1.Capabilities{
        v1.CapabilityCreateInstance,
        v1.CapabilityTerminateInstance,
        v1.CapabilityTags,  // Declare this if your API supports tags/labels
    }, nil
}
```

**When `CapabilityTags` is declared:**
- Store `RefID`, `CloudCredRefID`, and any additional tags via `CreateInstanceAttrs.Tags`
- The control plane will call `UpdateInstanceTags()` to add metadata after creation
- `ListInstances()` should support filtering via `TagFilters` for efficient queries

**Example (Shadeform with tags):**
```go
// At creation - store RefID and CloudCredRefID as tags
refIDTag := fmt.Sprintf("refID=%s", attrs.RefID)
cloudCredRefIDTag := fmt.Sprintf("cloudCredRefID=%s", c.GetReferenceID())
tags := []string{refIDTag, cloudCredRefIDTag}

// When reading back - extract from tags
refID := tags["refID"]
cloudCredRefID := tags["cloudCredRefID"]
```

### Alternative: When Tags Are NOT Supported

If your API doesn't support tags, you **still must** persist and return `RefID` and `CloudCredRefID`. Use creative alternatives:

**Example (Lambda Labs without tags):**
```go
// At creation - encode CloudCredRefID in instance name
name := fmt.Sprintf("%s--%s", c.GetReferenceID(), time.Now().UTC().Format(timeFormat))
// Use RefID as the SSH key pair name
keyPairName := attrs.RefID

// When reading back - extract from name and SSH key
nameParts := strings.Split(instance.Name, "--")
cloudCredRefID := nameParts[0]
refID := instance.SshKeyNames[0]
```

### Recommendation: Use Tags If Possible

**Tags are the recommended and easiest integration path.** They provide:
- Clean separation of metadata from instance properties
- Efficient server-side filtering via `TagFilters`
- Full billing/usage tracking capabilities
- Straightforward implementation

If your cloud API supports any form of instance tagging, labels, or metadata—**use it**.


> **Before implementing a custom solution**, please reach out to the Brev team. We can help design an approach that works reliably with the control plane and avoid edge cases that could cause instance correlation issues.

---

## 11. Error Handling and Status Reporting

### Error Categories

Your provider implementation should translate API errors into the standard error constants defined in [`v1/errors.go`](v1/errors.go):

| Category | Examples | Return This Error Constant |
|----------|----------|---------------------------|
| **Out of Stock** | No capacity in region | `v1.ErrInsufficientResources` |
| **Quota Exceeded** | Hit account limit | `v1.ErrOutOfQuota` |
| **Resource Not Found** | Instance/image doesn't exist | `v1.ErrResourceNotFound`, `v1.ErrInstanceNotFound`, `v1.ErrImageNotFound` |
| **Service Unavailable** | API temporarily down | `v1.ErrServiceUnavailable` |
| **Auth Failed** | Bad API key | Return HTTP 401/403 error |
| **Internal Error** | Your system issue | Return error with HTTP 500 details |

**Reference:** See [`v1/errors.go`](v1/errors.go) for the full list of error constants:

```go
var (
	ErrInsufficientResources = errors.New("zone has insufficient resources to fulfill the request, InsufficientCapacity")
	ErrOutOfQuota            = errors.New("out of quota in the region fulfill the request, InsufficientQuota")
	ErrImageNotFound         = errors.New("image not found")
	ErrDuplicateFirewallRule = errors.New("duplicate firewall rule")
	ErrInstanceNotFound      = errors.New("instance not found")
	ErrResourceNotFound      = errors.New("resource not found")
	ErrServiceUnavailable    = errors.New("api is temporarily unavailable")
)
```

### Out of Stock Handling

"Out of stock" is common with GPUs. Your implementation should return `v1.ErrInsufficientResources`:

1. Your API returns your specific "no capacity" error
2. Your provider translates this to `v1.ErrInsufficientResources`
3. Brev marks that type as temporarily unavailable in that region
4. The syncer will re-check availability on the next poll

**Example from Shadeform provider** ([`v1/providers/shadeform/instance.go`](v1/providers/shadeform/instance.go)):

```go
if shadeformErrorResponse.ErrorCode == outOfStockErrorCode {
    return v1.ErrInsufficientResources
}
```

**Example from Lambda Labs provider** ([`v1/providers/lambdalabs/errors.go`](v1/providers/lambdalabs/errors.go)):

```go
if strings.Contains(e.Error(), "Not enough capacity") || strings.Contains(e.Error(), "insufficient-capacity") {
    return v1.ErrInsufficientResources
}
```

---

## 12. Pricing and Billing

### How Pricing Works

Brev displays your prices via `InstanceType.BasePrice` (see [`v1/instancetype.go`](v1/instancetype.go)).

| Field | Type | Notes |
|-------|------|-------|
| **BasePrice** | `*currency.Amount` | From [`github.com/bojanz/currency`](https://pkg.go.dev/github.com/bojanz/currency#Amount) |
| **Currency** | Up to implementer | Most providers use `"USD"` |

### Billing

Billing arrangements are handled separately during the integration partnership setup.


---

## 13. Common Questions

### "Do you need access to our admin console?"

No. We only need programmatic API access. All operations go through your public API—see Section 6 for credential details.

### "What images/OS should our VMs run?"

| Requirement | Details |
|-------------|---------|
| **OS** | Ubuntu 22.04 (preferred) or 24.04 |

Custom images work if they meet these requirements. The SDK validates image compatibility via `ValidateInstanceImage()`.

### "What if we don't have public IPs?"

Public IP with SSH access is required for standard integration. Bastion/jump host routing is supported (see `InternalPortMappings` in the `Instance` struct). Other alternatives (VPN, Cloudflare tunnels) require custom integration work.

### "How do you track GPU interconnect (NVLink, SXM, PCIe)?"

We track interconnect type via the `GPU.NetworkDetails` field. Your implementation should populate this with values like `"PCIe"`, `"SXM"`, `"SXM4"`, or `"SXM5"`. If you have multiple variants (e.g., PCIe vs SXM versions of the same GPU), surface them as separate instance types.


### "What SLA/uptime do you expect from our API?"

| Requirement | Target |
|-------------|--------|
| **Availability** | 99%+ uptime |
| **Response time** | < 5 seconds typical |
| **Idempotency** | Supported where possible |

The Instance Syncer is resilient to brief outages—it retries and recovers automatically.

### "What does Brev do on the VMs after launch?"

After `CreateInstance` returns successfully:

1. **SSH connection**: Brev waits for SSH to become available (up to 10 minutes via `ValidateInstanceSSHAccessible`)
2. **Key bootstrapping**: Brev adds admin keys to `authorized_keys` via SSH
3. **Agent setup**: Brev installs a lightweight agent for tunnel management and environment configuration

You don't need to do anything special—just ensure the SSH public key from `CreateInstanceAttrs.PublicKey` is installed before the instance becomes accessible.

---

## Next Steps

To begin integration:

1. **Share your API documentation**: Instance types, lifecycle, auth
2. **Provide API credentials**: For Brev to access your API
3. **Technical call**: We'll align on specifics
4. **Implementation**: We build the adapter in our Cloud SDK
5. **Testing**: Validate end-to-end flow
6. **Launch**: Enable in Brev's catalog

Contact the Brev team to start the integration process.

---

## Glossary

| Term | Definition |
|------|------------|
| **Cloud Provider (You)** | Your company, providing GPU compute infrastructure |
| **Brev Control Plane** | Brev's system that syncs inventory and provisions instances |
| **Instance Type** | A compute configuration (CPU, GPU, RAM, storage, etc.) |
| **Location** | Primary region identifier (e.g., `us-west-2`) |
| **SubLocation** | Availability zone within a region (e.g., `us-west-2a`) |
| **noSub** | Placeholder used by `MakeGenericInstanceTypeID()` when no availability zone exists |
| **Syncer** | Brev's continuous process that polls your API for inventory |
| **Cloud SDK** | Brev's internal layer that adapts to different cloud provider APIs |
| **InstanceTypeID** | Stable, unique identifier for an instance type (format defined by implementer) |
| **SSH Key Injection** | Your API accepting Brev's SSH public key at VM creation |

---

*Document version: 2.0*
*For Brev integration partners*
