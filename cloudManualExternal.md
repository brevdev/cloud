# Brev Cloud Provider Integration Guide

**For Cloud Infrastructure Providers Integrating with Brev**

---

## Table of Contents

1. [Integration Overview](#1-integration-overview)
2. [How Brev Discovers Your Inventory](#2-how-brev-discovers-your-inventory)
3. [Instance Types: Your SKU Catalog](#3-instance-types-your-sku-catalog)
4. [Location and Availability Model](#4-location-and-availability-model)
5. [GPU Normalization](#5-gpu-normalization)
6. [Credential and Authentication Model](#6-credential-and-authentication-model)
7. [Provisioning Lifecycle](#7-provisioning-lifecycle)
8. [Network Requirements](#8-network-requirements)
9. [SSH and Control Plane Access](#9-ssh-and-control-plane-access)
10. [Firewall and Security Groups](#10-firewall-and-security-groups)
11. [Instance Metadata and Tags](#11-instance-metadata-and-tags)
12. [Error Handling and Status Reporting](#12-error-handling-and-status-reporting)
13. [Pricing and Billing](#13-pricing-and-billing)
14. [Common Questions](#14-common-questions)

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
| **Instance Type Listing API** | Discover your available SKUs |
| **Instance Lifecycle APIs** | Create, get, start, stop, terminate |
| **API Credentials for Brev** | Authenticate Brev's calls to your API |
| **SSH Key Injection** | Accept SSH public key at VM creation |
| **SSH Access on Port 22** | Control plane communication to VMs |

### Integration Architecture

### System Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              Brev Control Plane                                 │
│  ┌───────────────────────────────────────────────────────────────────────────┐  │
│  │                        Syncer Layer                                       │  │
│  │  ┌─────────────────────┐    ┌─────────────────────────────┐               │  │
│  │  │   InstanceSyncer    │    │   InstanceTypeSyncer        │               │  │
│  │  │ (Real-time state)   │    │ (Catalog sync every 1-5min) │               │  │
│  │  └──────────┬──────────┘    └──────────────┬──────────────┘               │  │
│  └─────────────┼──────────────────────────────┼──────────────────────────────┘  │
│                │                              │                                 │
└────────────────┼──────────────────────────────┼─────────────────────────────────┘
                 │                              │
                 ▼                              ▼
┌─────────────────────────────────────────────────────────────────────────────────┐
│                           CLOUD SDK (v1) - This Repo                            │   
│  ┌────────────────────────────────────────────────────────────────────────────┐ |
│  │              Provider Implementations                                      │ │
│  │  ┌─────────┐ ┌───────────┐ ┌─────────▼───┐ ┌───────────┐ ┌──────────────┐  │ │
│  │  │   A   │ │ │    B      │ │         C   │ │     D     │ │       E      │  │ │ 
│  │  │ Provider│ │  Provider │ │   Provider  │ │  Provider │ │   Provider   │  │ │
│  │  └────┬────┘ └─────┬─────┘ └──────┬──────┘ └─────┬─────┘ └──────┬───────┘  │ │
│  └───────┼────────────┼──────────────┼──────────────┼───────────────┼─────────┘ │
└──────────┼────────────┼──────────────┼──────────────┼───────────────┼───────────┘
           │            │              │              │               │
           ▼            ▼              ▼              ▼               ▼
┌──────────────────────────────────────────────────────────────────────────────────┐
│                          CLOUD PROVIDER APIs                                     │
│                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────┘

---
```

## 2. How Brev Discovers Your Inventory

### The Instance Type Syncer

Brev runs a **continuous synchronization process** that periodically queries your API to understand what compute is available. This isn't a one-time import—it's an ongoing reconciliation.

**Sync Behavior:**
- Polls your instance type listing API at regular intervals (typically every 1-5 minutes)
- Compares current catalog to previous state
- Updates availability, pricing, and specs as they change
- Marks types as unavailable when removed from your API
- Adds new types when they appear

### What We Query

We need an API endpoint that returns your available instance types. For each type, we extract:

| Field | What We Need | Example |
|-------|--------------|---------|
| **Type identifier** | Your internal name for this SKU | `gpu_1x_a100_sxm4` |
| **GPU model** | What GPU is in this instance | `A100 SXM4 80GB` |
| **GPU count** | How many GPUs | `8` |
| **CPU cores** | vCPU count | `128` |
| **Memory** | RAM in GB | `1024` |
| **Storage** | Disk in GB | `2000` |
| **Regions/Availability** | Where this type can launch | `us-west-1, us-east-2` |
| **Pricing** | Cost per hour (USD cents) | `3200` (= $32.00/hr) |

### API Patterns We Support

**Pattern A: Locational API (like AWS, GCP)**
Your API returns different availability per region. We query each region separately or you provide region-specific results.

```
GET /regions/us-west-1/instance-types → returns types available in us-west-1
GET /regions/us-east-2/instance-types → returns types available in us-east-2
```

**Pattern B: Global API (like Lambda Labs)**
Your API returns all types with their regional availability embedded.

```
GET /instance-types → returns all types with "available_regions": ["us-west-1", "us-east-2"]
```

Both patterns work. We adapt our sync logic to your API design.

---

## 3. Instance Types: Your SKU Catalog

### What Is an Instance Type to Brev?

Brev treats compute as **inventory**. Each instance type is a **SKU** (Stock Keeping Unit) in your catalog. Users browse your SKUs filtered by GPU, region, price, and availability.

### The Canonical Instance Type Model

When we ingest your instance types, we normalize them to this structure:

| Field | Type | Description |
|-------|------|-------------|
| `ID` | string | Brev's composite identifier (see below) |
| `Cloud` | string | Your cloud identifier (e.g., `"lambdalabs"`, `"crusoe"`) |
| `Type` | string | Your native type name |
| `Location` | string | Primary region identifier |
| `SubLocation` | string | Availability zone (or `"noSub"` if N/A) |
| `AvailableAzs` | []string | All zones where this type is available |
| `GPU` | string | Normalized GPU model name |
| `GPUCount` | int | Number of GPUs |
| `CPUCores` | int | vCPU count |
| `MemoryMB` | int | RAM in megabytes |
| `StorageMB` | int | Disk in megabytes |
| `PriceHr` | int | Price in cents per hour |
| `IsAvailable` | bool | Currently launchable |

### The Instance Type ID Format

Brev generates a unique ID for each instance type using this pattern:

```
{location}-{subLocation}-{type}
```

**Examples:**
- `us-west-1-us-west-1a-gpu_1x_a100` (locational cloud with AZs)
- `us-east-noSub-1x_a100_80gb_sxm4` (global cloud, no sublocation concept)
- `eu-central-1-noSub-h100_8x` (locational region, but you don't expose AZs)

**Why This Matters:**
This ID is how Brev tracks inventory. When provisioning, this ID connects the request to the correct SKU in your catalog.

### The "noSub" Convention

If your cloud doesn't have sub-locations (availability zones), we use the literal string `"noSub"` as a placeholder. This keeps the ID format consistent across all providers.

---

## 4. Location and Availability Model

### Location Hierarchy

Brev uses a two-tier location model:

```
Location (Region)
└── SubLocation (Availability Zone)
```

**Examples:**

| Your Term | Brev Location | Brev SubLocation |
|-----------|---------------|------------------|
| AWS `us-west-2a` | `us-west-2` | `us-west-2a` |
| GCP `us-central1-a` | `us-central1` | `us-central1-a` |
| Lambda Labs `us-tx-1` | `us-tx-1` | `noSub` |
| Your DC `phoenix-dc1` | `phoenix-dc1` | `noSub` |

### How Availability Is Tracked

For each instance type, we track:

1. **AvailableAzs**: List of all sub-locations where this type exists
2. **IsAvailable**: Boolean indicating if it's currently launchable

**Availability Meaning:**
- `IsAvailable: true` + `AvailableAzs: ["us-west-1a", "us-west-1b"]` = Can launch in either AZ
- `IsAvailable: false` = Type exists but is currently out of stock or disabled

### Region Normalization

We typically use your region identifiers as-is. If you have unique region names (`phoenix-main`, `denver-gpu-cluster`), those become the Location value.

---

## 5. GPU Normalization

### Why GPU Normalization Matters

Users search for GPUs by model. They want "H100" not "NVIDIA H100 80GB HBM3 SXM5 Accelerator". We normalize your GPU descriptions to standard names.

### The GPU Taxonomy

Brev normalizes GPUs to these canonical identifiers:

| Your Description | Brev GPU |
|------------------|----------|
| `NVIDIA H100 80GB HBM3` | `H100_SXM5` or `H100_PCIE` |
| `NVIDIA A100 SXM4 80GB` | `A100_SXM4_80GB` |
| `NVIDIA A100 PCIe 40GB` | `A100_PCIE_40GB` |
| `NVIDIA A10` | `A10` |
| `NVIDIA L40S` | `L40S` |
| `AMD MI300X` | `MI300X` |

### What We Parse

From your GPU field/description, we extract:
- **Model family**: H100, A100, L40S, etc.
- **Form factor**: SXM vs PCIe (affects interconnect and performance)
- **Memory size**: 40GB vs 80GB variants
- **Generation**: SXM4 vs SXM5

### Providing Clean GPU Data

The cleaner your GPU data, the better the user experience. Ideally provide:
- `gpu_model`: `"H100"` or `"A100"`
- `gpu_memory_gb`: `80`
- `gpu_variant`: `"SXM5"` or `"PCIe"`

If you only provide a description string, we'll parse it, but structured data is preferred.

---

## 6. Credential and Authentication Model

### How Brev Authenticates to Your API

Brev stores credentials for your cloud provider and uses them to make API calls. This is a direct relationship between **Brev's control plane** and **your cloud API**.

### What You Need to Provide

| Requirement | Details |
|-------------|---------|
| **API Credentials** | API key, token, or service account for Brev to use |
| **Authentication Endpoint** | How Brev authenticates (API key header, OAuth, etc.) |
| **Required Permissions** | List instance types, create/get/start/stop/terminate instances |

### Credential Exchange Process

1. **You provide** API credentials to Brev during integration setup
2. **Brev stores** credentials securely (encrypted at rest)
3. **Brev uses** credentials to call your API for sync and provisioning

### Credential Types

Providers define their own credential struct with whatever fields they need. Examples from existing providers:

| Provider | Credential Fields |
|----------|-------------------|
| **Lambda Labs** | `APIKey` |
| **Shadeform** | `APIKey` |
| **FluidStack** | `APIKey` |
| **AWS** | `AccessKeyID`, `SecretAccessKey` |
| **Nebius** | `ServiceAccountKey` (JSON), `TenantID` |
| **Launchpad** | `APIToken`, `APIURL` |

Your credential struct just needs to implement the `CloudCredential` interface.

### SSH Keys (Separate from API Credentials)

For each VM launch, Brev provides an SSH public key in the create request. **You need to:**
1. Accept an SSH public key parameter in your create instance API
2. Install that key in the VM's default user `~/.ssh/authorized_keys`
3. Ensure SSHD is running on port 22

Brev generates and manages these SSH keys—you just need to accept and install them.

---

## 7. Provisioning Lifecycle

### Instance States

Brev tracks instances through these states:

| State | Meaning |
|-------|---------|
| `pending` | Create request sent, waiting for VM |
| `running` | Instance is up and accessible |
| `stopping` | Stop request sent |
| `stopped` | Instance stopped but not terminated |
| `terminating` | Terminate request sent |
| `terminated` | Instance terminated |
| `failed` | Provisioning failed |

### What Your Create API Should Return

| Field | Required | Description |
|-------|----------|-------------|
| `instance_id` | Yes | Your unique identifier |
| `status` | Yes | Current state |
| `public_ip` | When running | IPv4 address for SSH |
| `region` | Yes | Where it launched |
| `instance_type` | Yes | What SKU was provisioned |

### Polling vs Webhooks

Most integrations use **polling**—Brev periodically calls your Get Instance API until status is `running`. If you support webhooks for state changes, that can reduce API load.

---

## 8. Network Requirements

### Critical Requirement: Public IP with SSH Access

Every instance **must** have a publicly routable IP address with port 22 (SSH) accessible. This is how Brev's control plane communicates with the instance.

### Network Configuration at Launch

When provisioning, we pass:
- **SSH public key**: Key to install in `authorized_keys`
- **Firewall rules**: Ports to open (see Section 10)

### IP Assignment

| Scenario | Requirement |
|----------|-------------|
| **Ideal** | Public IPv4 assigned automatically at launch |
| **Acceptable** | Public IP available via API after instance starts |
| **Not Supported** | NAT-only instances with no public ingress |

### IPv6

IPv6-only instances are not currently supported. We require IPv4 for SSH connectivity.

---

## 9. SSH and Control Plane Access

### Why SSH Is Critical

SSH (port 22) is Brev's **control channel**. After your VM is running, Brev connects via SSH to:

1. **Configure the environment**: Install Brev agent, set up development tools
2. **Enable connections**: Set up connection paths for users
3. **Manage instance**: Execute commands, transfer files

### What You Must Support

| Requirement | Details |
|-------------|---------|
| **Accept SSH key in create request** | Your API must accept an SSH public key parameter |
| **Install key in VM** | Key goes in default user's `~/.ssh/authorized_keys` |
| **SSHD running on port 22** | Standard SSH daemon, default config is fine |
| **Port 22 reachable** | Public IP with port 22 open |

### SSH User

We typically connect as:
- `root` (if permitted)
- `ubuntu` (common on Ubuntu images)
- Whatever default user your images provide

Let us know your default SSH user during integration setup.

---

## 10. Firewall and Security Groups

### Brev's Firewall Model

Brev uses a provider-agnostic firewall model that maps to your security group / firewall implementation:

**Ingress Rules** (inbound traffic):
```
Port(s)     Protocol    Source
22          TCP         0.0.0.0/0    # SSH - REQUIRED
443         TCP         0.0.0.0/0    # HTTPS (optional)
8080        TCP         0.0.0.0/0    # User app (optional)
```

**Egress Rules** (outbound traffic):
```
Port(s)     Protocol    Destination
*           *           0.0.0.0/0    # Allow all outbound
```

### Minimum Required Ports

| Port | Protocol | Direction | Purpose |
|------|----------|-----------|---------|
| **22** | TCP | Inbound | SSH (mandatory) |

All other ports are configurable based on workload needs.

### Mapping to Your System

Your firewall / security group implementation should:
1. Accept our firewall rules in the create request (or apply defaults)
2. Ensure port 22 is open for Brev's control plane
3. Allow additional ports to be specified for applications

---

## 11. Instance Metadata and Tags

### Tags We Set

Brev may set tags/labels on instances for identification:

| Tag | Value | Purpose |
|-----|-------|---------|
| `brev-instance-id` | Brev's internal ID | Cross-reference |
| `Name` | User-specified | Display name |

### Tag Requirements

Your API should support:
- Setting tags at instance creation
- Updating tags on running instances
- Querying instances by tag (helpful but not required)

If you don't support tags, we track the mapping on our side.

---

## 12. Error Handling and Status Reporting

### Error Categories

| Category | Examples | How to Report |
|----------|----------|---------------|
| **Out of Stock** | No capacity in region | Return specific error code |
| **Quota Exceeded** | Hit account limit | Return quota error |
| **Invalid Request** | Bad instance type | Return validation error |
| **Auth Failed** | Bad API key | Return 401/403 |
| **Internal Error** | Your system issue | Return 500 with details |

### Preferred Error Format

We prefer errors that include:
- **Error code**: Machine-readable identifier
- **Message**: Human-readable description
- **Region** (if relevant): Where the failure occurred

### Out of Stock Handling

"Out of stock" is common with GPUs. Ideal handling:
1. Your API returns a clear "no capacity" error
2. We mark that type as temporarily unavailable in that region
3. The syncer will re-check availability on the next poll

---

## 13. Pricing and Billing

### How Pricing Works

Brev displays your prices. We need:

| Field | Format | Example |
|-------|--------|---------|
| **Hourly price** | Cents (integer) | `3200` = $32.00/hr |
| **Currency** | USD assumed | - |

### Billing

Billing arrangements are handled separately during the integration partnership setup.

### Price Sync

Prices sync along with instance types. When you update pricing in your system, we pick it up in the next sync cycle.

---

## 14. Common Questions

### "What credentials do you need from us?"

We need API credentials that allow Brev to:
- List your available instance types
- Create, get, start, stop, and terminate instances
- Optionally: update tags, modify firewall rules

This is typically an API key or service account.

### "Do you need access to our admin console?"

No. We only need API access. All operations go through your public API.

### "What images/OS should our VMs run?"

We work best with:
- **Ubuntu 22.04 or 24.04** (preferred)
- **CUDA pre-installed** (for GPU instances)
- **Python 3.10+** available
- **SSHD running** on port 22

Custom images can work, but Ubuntu with CUDA is the smoothest path.

### "How do you handle the SSH keys?"

For each VM:
1. Brev generates an SSH key pair
2. Brev passes the public key in the create request
3. You install it in the VM's `authorized_keys`
4. Brev connects using the private key

You don't manage these keys—just accept them at VM creation.

### "What if we don't have public IPs?"

Public IP with SSH access is required for the standard integration. Alternatives:
- **VPN/Private connectivity**: Custom integration needed
- **Bastion host**: Brev can SSH through a jump box
- **Cloudflare tunnel**: Instance calls out, no inbound needed

These require additional integration work.

### "How do you handle multi-GPU interconnect (NVLink, etc.)?"

We track GPU configuration but don't currently differentiate NVLink vs PCIe interconnect in the UI. If you have multiple variants (NVLink cluster vs standalone), surface them as different instance types.

### "What about bare metal vs VMs?"

Both work. From Brev's perspective, if it has an IP and SSH access, it's an instance. Bare metal instances are provisioned the same way.

### "How do we test the integration?"

Typical integration process:
1. **Staging environment**: Brev tests against your sandbox/dev API
2. **Test credentials**: You provide test account with limited quota
3. **Validation**: We verify create, get, stop, start, terminate
4. **Production**: Enable in Brev's catalog

### "What SLA/uptime do you expect from our API?"

Your API should be:
- **Available**: 99%+ uptime for instance operations
- **Responsive**: <5 second response times typical
- **Consistent**: Idempotent operations where possible

Sync polling is resilient to brief outages—we retry and recover.

### "What does Brev do on the VMs after launch?"

After Brev creates a VM via your API:
1. Brev SSHs into the VM using the key we provided at creation
2. Brev installs a lightweight agent and configures the environment
3. Brev sets up connection paths

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
| **Instance Type** | A SKU representing a compute configuration (CPU, GPU, RAM, etc.) |
| **Location** | Primary region identifier (e.g., `us-west-2`) |
| **SubLocation** | Availability zone within a region (e.g., `us-west-2a`) |
| **noSub** | Placeholder when your cloud doesn't have availability zones |
| **Syncer** | Brev's continuous process that polls your API for inventory |
| **Cloud SDK** | Brev's internal layer that adapts to different cloud provider APIs |
| **InstanceTypeID** | Brev's composite identifier: `{location}-{subLocation}-{type}` |
| **SSH Key Injection** | Your API accepting Brev's SSH public key at VM creation |

---

*Document version: 2.0*
*For Brev integration partners*
