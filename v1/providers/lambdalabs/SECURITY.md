# Lambda Labs SECURITY.md for Brev Cloud SDK

This document outlines how the Lambda Labs integration complies with Brev Cloud SDK's security architecture using Lambda Cloud primitives such as firewall rules and virtual network isolation.

## 🔑 SSH Access Requirements

**Lambda Labs instances must support SSH server functionality and SSH key-based authentication for Brev access.**

### SSH Implementation
- **SSH Server**: All Lambda Labs instances have SSH server (OpenSSH) pre-installed and running
- **SSH Key Authentication**: Lambda Labs supports SSH key injection during instance creation
- **Key Management**: Public SSH keys are automatically configured in `~/.ssh/authorized_keys`
- **Access Method**: SSH provides the primary secure access method for instance managementn.

---

## 🔐 Network Security

### Current Model

Lambda Labs recently introduced support for per-instance and per-cluster firewalls, enabling more granular network security controls. **However, the Brev Cloud SDK integration does _not yet_ utilize these new per-instance/cluster firewall features.**

**Currently, we rely on Lambda Labs' account-level global firewall rules to enforce network security.** This means:

- **Inbound:** By default, all inbound traffic is restricted at the account level except for SSH (TCP/22), which is required for instance access and management. All other inbound ports and protocols are closed unless explicitly allowed at the account/global level.
- **Outbound:** All outbound traffic is **unrestricted by default**. Lambda Cloud does not impose egress restrictions, so instances can freely initiate outbound connections.

> **Note:** The configuration of these global firewall rules is performed _outside_ of the Brev Cloud SDK's cloud integration package. Setting up these account-level firewall rules to restrict all inbound traffic except SSH is a **prerequisite for securely adding a new Lambda Labs account** to Brev.

### Explicit Inbound Access

- Brev restricts inbound access to SSH only (TCP/22) using Lambda Labs' global firewall rules.
- No other inbound ports are open by default. If additional access is required, it must be configured at the account/global firewall level (not per-instance).
- The new per-instance/cluster firewall features are not yet integrated with Brev's `FirewallRule` abstraction.

### Implementation Notes

- **Inbound Deny (except SSH):** Achieved by configuring Lambda Labs' global firewall rules to allow only SSH and deny all other inbound traffic.
- **Outbound Allow:** No changes needed; outbound access is unrestricted by Lambda by default.
- **FirewallRule Mapping:** At this time, Brev's `FirewallRule` resources are not mapped to Lambda Labs' new per-instance/cluster firewall APIs. All restrictions are enforced at the account/global level.

> **Planned Improvement:** We plan to update the integration to leverage Lambda Labs' per-instance/cluster firewalls for finer-grained network control in a future release.

---

## 🔒 Data Protection

### Encryption at Rest

- All persistent storage is encrypted at rest using industry-standard encryption (e.g., AES-256) by default.
- Lambda Cloud does not currently support customer-managed encryption keys.

### Encryption in Transit

- All Brev SDK and control plane communication uses HTTPS with TLS 1.2+.
- Users are encouraged to use secure protocols (e.g., SSH, HTTPS) for application-level traffic to and from their instances.

### Key Management

- Provider-managed keys are used for at-rest encryption.
- No customer-managed key option is currently available.

---

## ✅ Implementation Checklist

- [x] **Inbound Deny by Default** using Lambda Labs firewall configuration
- [x] **Outbound Allow by Default** (no action required)
- [x] **Explicit Inbound Access** mapped to Brev `FirewallRule` ➝ Lambda firewall rule
- [x] **Persistent Disk Encryption** enabled by default (AES-256)
- [x] **TLS 1.2+** enforced for all Brev SDK and API communication
- [x] **Provider-Managed Keys** for encryption at rest

---

## 🛡️ Security Contact

For security issues or vulnerability disclosures related to this integration, please contact:

**Email:** [brev@nvidia.com](mailto:brev@nvidia.com)

Please report any vulnerabilities privately before disclosing them publicly.
