# Hyperstack Provider

This package implements the minimal Brev Cloud v1 compute surface for Hyperstack using the official Go SDK.

Supported capabilities are instance creation and termination. Locations map to Hyperstack regions, instance types map to flavors, and creation resolves the region's `default-<REGION>` environment through the API. `CreateInstanceAttrs.ImageID` is treated as a Hyperstack image name; when omitted, the provider uses Ubuntu Server 22.04 LTS (Jammy Jellyfish).

VMs are created with a direct inbound TCP/22 security rule for bootstrap SSH access. Caller-provided ingress ports are included as additional direct per-VM rules. Outbound behavior is derived by Hyperstack.

## Read-only validation

Set `HYPERSTACK_API_KEY` and optionally `HYPERSTACK_API_URL`, then run:

```sh
go test -run TestReadOnlyValidation ./v1/providers/hyperstack
```

This validation only lists regions, flavors, prices, and existing virtual machines. It never creates, updates, or deletes a Hyperstack resource.
