# Nebius SECURITY.md for Brev Cloud SDK

This document explains how Nebius VMs meet Brev Cloud SDK's security requirements using Nebius primitives like Security Groups, VPCs, and projects.

## 🔑 SSH Access Requirements

**Nebius VMs must support SSH server functionality and SSH key-based authentication for Brev access.**

### SSH Implementation
- **SSH Server**: All Nebius VM instances include SSH server (OpenSSH) installed and running by default
- **SSH Key Authentication**: Nebius supports SSH public key injection during VM creation via metadata
- **Key Management**: SSH keys are automatically configured in the VM's `~/.ssh/authorized_keys` file
- **Security Integration**: SSH access works within the Security Group firewall rules defined for the instances.

---

## Network Security

### Default Rules

* **Inbound:** All inbound traffic is **denied by default** using a custom Nebius Security Group with no inbound rules.
* **Outbound:** We explicitly **allow all outbound traffic** by adding a wide egress rule (all ports/protocols to `0.0.0.0/0`).

### Explicit Access

* All inbound access must be added manually via Brev’s `FirewallRule` interface.
* These are mapped to Nebius Security Group rules that allow specific ports and sources.

### Isolation

* Each cluster uses its own Security Group.

---

## Cluster Security

* Instances in the same cluster:

  * Share a Security Group.
  * Can talk to each other using a "self" rule (Nebius allows rules that permit traffic from the same group).
* No traffic is allowed from outside the cluster unless explicitly opened.
* Different clusters use different Security Groups to ensure isolation.

---

## Data Protection

### At Rest

* Nebius encrypts all persistent disks by default using AES-256 or equivalent.

### In Transit

* All Brev SDK API calls use HTTPS (TLS 1.2+).
* Internal instance traffic should use secure protocols (e.g., SSH, HTTPS).

---

## Implementation Checklist

* [ ] Default deny-all inbound using custom Nebius Security Group
* [ ] Allow-all outbound via security group egress rule
* [ ] `FirewallRule` maps to explicit Nebius SG ingress rule
* [ ] Instances in the same cluster can talk via shared SG "self" rule
* [ ] Different clusters are isolated using separate SGs or VPCs
* [x] Disk encryption enabled by default (Nebius default)
* [x] TLS used for all API and external communication (Nebius SDK default)

## Authentication Implementation

### Service Account Setup

Nebius uses JWT-based service account authentication:

1. **Service Account Creation**: Create a service account in Nebius IAM
2. **Key Generation**: Generate a JSON service account key file
3. **JWT Token Exchange**: SDK automatically handles JWT signing and token exchange
4. **API Authentication**: All API calls use Bearer token authentication

### Authentication Flow

```
1. Load service account JSON key
2. Generate JWT with RS256 signing (kid, iss, sub, exp claims)
3. Exchange JWT for IAM token via TokenExchangeService
4. Use IAM token in Authorization header for compute API calls
```

### Implementation Details

The `NebiusClient` uses the official Nebius Go SDK which handles:
- Automatic JWT token generation and refresh
- gRPC connection management with TLS 1.2+
- Service discovery for Nebius API endpoints
- Retry logic and error handling

---

## Security Contact

* Email: [brev@nvidia.com](mailto:brev@nvidia.com)
* Please report vulnerabilities privately before disclosing publicly.
