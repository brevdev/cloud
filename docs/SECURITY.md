# 🔒 Security Requirements

This document outlines the security requirements and best practices for implementing cloud integrations with the Brev Compute SDK. If the cloud provider has a different security model, please indicate in the providers SECURITY.md file.

## 🔑 SSH Access Requirements

**Brev depends on SSH server functionality and SSH key-based authentication for instance access.**

### SSH Server Requirements

1. **SSH Server**: All instances must have an SSH server running and accessible
2. **SSH Key Authentication**: Instances must support SSH key-based authentication
3. **Key Injection**: The SDK must support passing SSH public keys during instance creation
4. **Secure Access**: SSH access provides the primary method for secure instance management and user access

**Implementation Requirements:**

- SSH server (OpenSSH or equivalent) must be installed and running on all instances
- systemd must be running and accessible via systemctl command
- SSH key pairs must be supported for authentication
- Public keys must be injectable during instance provisioning
- SSH access must be available through the configured firewall rules

Providers supporting `CreateInstanceAttrs.SSHCertificateAuthority` install resource-scoped CA trust instead of the user's public key and report `Instance.SSHCertificateBootstrap` on creation. Providers that do not support CA bootstrap retain their existing key setup and leave that result false. CA private keys remain in the caller's secret store; the SDK receives only public trust material.

## 🌐 Network Security Requirements

### Network Requirements

**All instances must implement a "deny all inbound, allow all outbound" security model by default.**

```
┌─────────────────────────────────────────────────────────────┐
│                    Instance Security Model                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                Instance Boundary                    │    │
│  │                                                     │    │
│  │  ┌─────────────────────────────────────────────┐    │    │
│  │  │           Inbound Traffic                   │    │    │
│  │  │                                           │    │    │
│  │  │  ❌ DENIED BY DEFAULT                      │    │    │
│  │  │  • All ports blocked                       │    │    │
│  │  │  • All protocols blocked                   │    │    │
│  │  │  • All source IPs blocked                  │    │    │
│  │  │                                           │    │    │
│  │  └─────────────────────────────────────────────┘    │    │
│  │                                                     │    │
│  │  ┌─────────────────────────────────────────────┐    │    │
│  │  │           Outbound Traffic                  │    │    │
│  │  │                                           │    │    │
│  │  │  ✅ ALLOWED BY DEFAULT                     │    │    │
│  │  │  • All ports allowed                       │    │    │
│  │  │  • All protocols allowed                   │    │    │
│  │  │  • All destination IPs allowed             │    │    │
│  │  │                                           │    │    │
│  │  └─────────────────────────────────────────────┘    │    │
│  │                                                     │    │
│  └─────────────────────────────────────────────────────┘    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Implementation Requirements:**

1. **Default State**: All inbound traffic must be blocked by default (exemption may be made to 22, though the sdk prefers to explicitly set this)
2. **Explicit Allow**: Inbound access must be explicitly granted through `FirewallRule` resources
3. **Outbound Freedom**: Outbound traffic should be unrestricted by default
5. **Security Groups**: Use cloud provider security groups or equivalent (AWS Security Groups, GCP Firewall Rules, Azure NSGs) for network isolation
6. **Default Deny**: Configure security groups with default deny rules for all inbound traffic

### Cluster Security

**Instances within the same cluster can communicate with each other while maintaining external inbound restrictions.**

```
┌─────────────────────────────────────────────────────────────┐
│                    Cluster Security Model                    │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────────────────────────────────────────────┐    │
│  │                Cluster Boundary                    │    │
│  │                                                     │    │
│  │  ┌─────────────────┐    ┌─────────────────┐        │    │
│  │  │   Instance A    │    │   Instance B    │        │    │
│  │  │                 │    │                 │        │    │
│  │  │  ✅ Internal    │◄──►│  ✅ Internal    │        │    │
│  │  │  Communication  │    │  Communication  │        │    │
│  │  │                 │    │                 │        │    │
│  │  └─────────────────┘    └─────────────────┘        │    │
│  │                                                     │    │
│  │  ┌─────────────────────────────────────────────┐    │    │
│  │  │           External Inbound                 │    │    │
│  │  │                                           │    │    │
│  │  │  ❌ STILL DENIED                          │    │    │
│  │  │  • All external ports blocked             │    │    │
│  │  │  • All external protocols blocked         │    │    │
│  │  │  • All external source IPs blocked        │    │    │
│  │  │                                           │    │    │
│  │  └─────────────────────────────────────────────┘    │    │
│  │                                                     │    │
│  └─────────────────────────────────────────────────────┘    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Cluster Communication Requirements:**

1. **Internal Allow**: Instances within the same cluster can communicate on all ports and protocols
2. **External Deny**: External inbound traffic remains blocked by default
3. **Cluster Isolation**: Different clusters are isolated from each other unless explicitly configured
4. **Shared Security Groups**: Instances in the same cluster share security group rules for internal communication
5. **Cross-Cluster Access**: Inter-cluster communication requires explicit firewall rules

## 🛡️ Data Protection

### Encryption Requirements

- **Data at Rest**: All persistent storage must be encrypted
- **Data in Transit**: All network communications must be encrypted (TLS 1.2+)
- **Encryption Algorithms**: Use industry-standard encryption algorithms (AES-256, etc.)

## 📋 Implementation Checklist

### Cloud Provider Integration

- [ ] Implement default "deny all inbound" firewall rules
- [ ] Support explicit firewall rule creation/mutation through SDK
- [ ] Enable encryption for all persistent storage
- [ ] Document security configurations

## 📞 Security Contact

For security issues, vulnerabilities, or questions:

- **Security Email**: brev@nvidia.com
- **Responsible Disclosure**: Please report vulnerabilities through our security email

---

**Note**: This document is a living document and will be updated as security requirements evolve. All cloud integrations must comply with these requirements to ensure the security and integrity of the Brev Compute SDK ecosystem.    
