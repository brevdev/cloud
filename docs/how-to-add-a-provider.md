# How to Add a Cloud Provider

This guide explains how to add a new cloud provider to the Brev Cloud SDK (v1). The Lambda Labs provider is the best working, well-tested example—use it as your canonical reference.

Goals:

- Implement a provider-specific CloudCredential (factory) and CloudClient (implementation) that satisfy pkg/v1 interfaces.
- Accurately declare Capabilities based on the provider’s API surface.
- Implement at least instance lifecycle and instance types, adhering to security requirements.
- Add validation tests and (optionally) a GitHub Actions workflow to run them with real credentials.

Helpful background:

- Architecture overview: ../docs/ARCHITECTURE.md
- Security requirements: ../docs/SECURITY.md
- Validation testing framework: ../docs/VALIDATION_TESTING.md
- v1 design notes: ../pkg/v1/V1_DESIGN_NOTES.md

Provider examples:

- Lambda Labs (canonical): ../internal/lambdalabs/v1/README.md
- Nebius (in progress): ../internal/nebius/v1/README.md
- Fluidstack (in progress): ../internal/fluidstack/v1/README.md

---

## Core v1 Interfaces You Must Target

CloudClient is a composed interface of provider capabilities. You don’t need to implement everything—only what your provider supports—but you must advertise Capabilities correctly.

- CloudClient composition: ../pkg/v1/client.go
  - Key aggregation: CloudBase, CloudQuota, CloudRebootInstance, CloudStopStartInstance, CloudResizeInstanceVolume, CloudMachineImage, CloudChangeInstanceType, CloudModifyFirewall, CloudInstanceTags, UpdateHandler
- Capabilities system: ../pkg/v1/capabilities.go
- Instance lifecycle, validation helpers, and types: ../pkg/v1/instance.go
- Instance types and validation helpers: ../pkg/v1/instancetype.go

Patterns to follow:

- Embed v1.NotImplCloudClient in your client so unsupported methods gracefully return ErrNotImplemented (see ../pkg/v1/notimplemented.go).
- Accurately return capability flags that match your provider’s real API.
- Prefer stable, provider-native identifiers; otherwise use MakeGenericInstanceTypeID/MakeGenericInstanceTypeIDFromInstance.

---

---

## Compute Brokers & Marketplaces (Aggregators)

This SDK supports providers that aggregate compute from multiple upstream sources (multi-cloud brokers, marketplaces, or exchanges). When implementing an aggregator, use these to differentiate where the compute comes from while keeping the interface consistent:

- Provider (CloudProviderID): Identify your aggregator (e.g., "mybroker"). If you expose underlying vendors, include that metadata on returned resources (e.g., tags/labels) or encode it in stable IDs that you control.
- Location and SubLocation: Map upstream regions/zones into `Location` and `SubLocation` so users can choose placement consistently across sources. For example, `Location="us-west"` and `SubLocation="vendorA/zone-2"` or `SubLocation="sv15/DC3"` for finer placement.
- InstanceType IDs: If upstream vendors don’t provide stable, cross-market IDs, generate stable IDs using `MakeGenericInstanceTypeID` and include upstream hints in IDs or metadata. Ensure stability over time to avoid breaking consumers.
- InstanceType attributes (recommended): Use instance type attributes to delineate behavior differences across upstream sources (e.g., performance, network, storage, locality). There is also a `provider` attribute on the instance type you can use to indicate the originating vendor/source.

Notes:

- Capabilities represent what your broker can support. Differences between upstream vendors should be reflected in instance type attributes rather than reducing declared capabilities to the lowest common denominator.
- Keep your `Location`/`SubLocation` stable even if upstream identifiers change; translate upstream → broker-stable naming.
- Conform to the default-deny inbound model; document any upstream limitations under `internal/{provider}/SECURITY.md`.

## Directory Layout

Create a new provider folder:

- internal/{provider}/
  - SECURITY.md (provider-specific notes; link to top-level security expectations)
  - CONTRIBUTE.md (optional provider integration notes)
  - v1/
    - client.go (credentials and client)
    - instance.go (instance lifecycle + helpers)
    - instancetype.go (instance types)
    - capabilities.go (capability declarations)
    - networking.go, image.go, storage.go, tags.go, quota.go, location.go (as applicable)
    - validation_test.go (validation suite entry point)

Use Lambda Labs as the pattern:

- ../internal/lambdalabs/v1/client.go
- ../internal/lambdalabs/v1/instance.go
- ../internal/lambdalabs/v1/capabilities.go

---

## Minimal Scaffold (Copy/Paste Template)

Place in internal/{provider}/v1/client.go. Adjust names, imports, and fields for your provider.

```go
package v1

import (
    "context"

    v1 "github.com/brevdev/cloud/pkg/v1"
)

type {Provider}Credential struct {
    RefID string
    // Add auth fields (e.g., APIKey, ClientID, Secret, Tenant, etc.)
}

var _ v1.CloudCredential = &{Provider}Credential{}

func New{Provider}Credential(refID string /* auth fields */) *{Provider}Credential {
    return &{Provider}Credential{
        RefID: refID,
        // ...
    }
}

func (c *{Provider}Credential) GetReferenceID() string { return c.RefID }
func (c *{Provider}Credential) GetAPIType() v1.APIType  { return v1.APITypeLocational /* or v1.APITypeGlobal */ }
func (c *{Provider}Credential) GetCloudProviderID() v1.CloudProviderID {
    return "{provider-id}" // e.g., "lambdalabs"
}
func (c *{Provider}Credential) GetTenantID() (string, error) {
    // Derive stable tenant ID for quota/account scoping if possible
    return "", nil
}

func (c *{Provider}Credential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
    return get{Provider}Capabilities(), nil
}

func (c *{Provider}Credential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
    // Create a client configured for a given location if locational API
    return New{Provider}Client(c.RefID /* auth fields */).MakeClient(ctx, location)
}

// ---------------- Client ----------------

type {Provider}Client struct {
    v1.NotImplCloudClient
    refID    string
    location string
    // add http/sdk client fields, base URLs, etc.
}

var _ v1.CloudClient = &{Provider}Client{}

func New{Provider}Client(refID string /* auth fields */) *{Provider}Client {
    return &{Provider}Client{
        refID: refID,
        // init http/sdk clients here
    }
}

func (c *{Provider}Client) GetAPIType() v1.APIType { return v1.APITypeLocational /* or Global */ }
func (c *{Provider}Client) GetCloudProviderID() v1.CloudProviderID { return "{provider-id}" }
func (c *{Provider}Client) GetReferenceID() string                 { return c.refID }
func (c *{Provider}Client) GetTenantID() (string, error)           { return "", nil }

func (c *{Provider}Client) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
    c.location = location
    return c, nil
}
```

Declare capabilities in internal/{provider}/v1/capabilities.go:

```go
package v1

import (
    "context"

    v1 "github.com/brevdev/cloud/pkg/v1"
)

func get{Provider}Capabilities() v1.Capabilities {
    return v1.Capabilities{
        v1.CapabilityCreateInstance,
        v1.CapabilityTerminateInstance,
        v1.CapabilityCreateTerminateInstance,
        // add others supported by your provider: reboot, stop/start, machine-image, tags, resize-volume, modify-firewall, etc.
    }
}

func (c *{Provider}Client) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
    return get{Provider}Capabilities(), nil
}

func (c *{Provider}Credential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
    return get{Provider}Capabilities(), nil
}
```

Implement instance lifecycle in internal/{provider}/v1/instance.go (map to provider API):

```go
package v1

import (
    "context"
    "fmt"

    v1 "github.com/brevdev/cloud/pkg/v1"
)

func (c *{Provider}Client) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
    // 1) ensure SSH key present (or inject via API) per ../docs/SECURITY.md
    // 2) map attrs to provider request (location, instance type, image, tags, firewall rules if supported)
    // 3) launch and return instance converted to v1.Instance
    return nil, fmt.Errorf("not implemented")
}

func (c *{Provider}Client) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
    return nil, fmt.Errorf("not implemented")
}

func (c *{Provider}Client) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
    return nil, fmt.Errorf("not implemented")
}

func (c *{Provider}Client) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
    return fmt.Errorf("not implemented")
}

// Optional if supported:
func (c *{Provider}Client) RebootInstance(ctx context.Context, id v1.CloudProviderInstanceID) error { return fmt.Errorf("not implemented") }
func (c *{Provider}Client) StopInstance(ctx context.Context, id v1.CloudProviderInstanceID) error   { return fmt.Errorf("not implemented") }
func (c *{Provider}Client) StartInstance(ctx context.Context, id v1.CloudProviderInstanceID) error  { return fmt.Errorf("not implemented") }

// Merge strategies (pass-through is acceptable baseline).
func (c *{Provider}Client) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance         { return newInst }
func (c *{Provider}Client) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.Type { return newIt }
```

See the canonical mapping and conversion logic in Lambda Labs:

- Create/terminate/list/reboot: ../internal/lambdalabs/v1/instance.go
- Capabilities: ../internal/lambdalabs/v1/capabilities.go
- Client/credential + NotImpl: ../internal/lambdalabs/v1/client.go

Implement instance types in internal/{provider}/v1/instancetype.go:

- Implement:
  - GetInstanceTypes(ctx, args GetInstanceTypeArgs) ([]InstanceType, error)
  - GetInstanceTypePollTime() time.Duration
- Use stable IDs if provider offers them. If not, use MakeGenericInstanceTypeID.
- Validate with helpers:
  - ValidateGetInstanceTypes: ../pkg/v1/instancetype.go
  - ValidateLocationalInstanceTypes: ../pkg/v1/instancetype.go
  - ValidateStableInstanceTypeIDs (if you maintain a stable ID list)

---

## Capabilities: Be Precise

The SDK uses a three-level capability system to accurately represent what operations are supported:

### 1. Provider-Level Capabilities

These are high-level features that your cloud provider's API supports, declared in your `GetCapabilities()` method. Capability flags live in ../pkg/v1/capabilities.go. Only include capabilities your API actually supports. For example, Lambda Labs supports:

- Create/terminate/reboot instance (`CapabilityCreateInstance`, `CapabilityTerminateInstance`, `CapabilityRebootInstance`)
- Does not (currently) support stop/start, resize volume, machine image, tags

### 2. Instance Type Capabilities

These are hardware-specific features that vary by instance configuration, expressed as boolean fields on the `InstanceType` struct:

- `Stoppable`: Whether instances of this type can be stopped/started
- `Rebootable`: Whether instances of this type can be rebooted
- `CanModifyFirewallRules`: Whether firewall rules can be modified for this instance type
- `Preemptible`: Whether this instance type supports spot/preemptible pricing
- `TunneledSSH`: Whether connections must be routed through a client-side tunnel proxy. This is required for instances that do not have public IP addresses.

### 3. Instance Capabilities

These are capability boolean fields replicated on individual `Instance` objects, similar to Instance Type capabilities but applied to the running instance rather than the type template. While these fields could theoretically be derived from the associated `InstanceType`, they are duplicated on the instance for performance and convenience reasons. Examples include:

- `Stoppable`: Whether this specific instance can be stopped/started
- `Rebootable`: Whether this specific instance can be rebooted
- `TunneledSSH`: Whether connections must be routed through a client-side tunnel proxy. This is required for instances that do not have public IP addresses.

These fields must be kept accurate and in sync with the corresponding InstanceType capabilities, even though they appear redundant. They can also reflect runtime state-dependent variations - for example, a running instance might support certain operations that a stopped instance cannot, based on the current `LifecycleStatus`.

Reference:

- Lambda capabilities: ../internal/lambdalabs/v1/capabilities.go

---

## Security Requirements

All providers must conform to ../docs/SECURITY.md:

- Default deny all inbound, allow all outbound
- SSH server must be available with key-based auth
- Firewall rules should be explicitly configured via FirewallRule when supported
- If your provider’s firewall model is global/project-scoped rather than per-instance, document limitations in internal/{provider}/SECURITY.md and reflect that by omitting CapabilityModifyFirewall if applicable.

Provider-specific security doc examples:

- Lambda Labs: ../internal/lambdalabs/SECURITY.md
- Nebius: ../internal/nebius/SECURITY.md
- Fluidstack: ../internal/fluidstack/v1/SECURITY.md

---

## Validation Testing and CI

Use the shared validation suite to test your provider with real credentials.

- Validation framework and instructions: ../docs/VALIDATION_TESTING.md
- Shared package: ../internal/validation/suite.go

Steps:

1. Create internal/{provider}/v1/validation_test.go:

```go
package v1

import (
    "os"
    "testing"

    "github.com/brevdev/cloud/internal/validation"
)

func TestValidationFunctions(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping validation tests in short mode")
    }

    apiKey := os.Getenv("YOUR_PROVIDER_API_KEY")
    if apiKey == "" {
        t.Skip("YOUR_PROVIDER_API_KEY not set, skipping validation tests")
    }

    cfg := validation.ProviderConfig{
        Credential: New{Provider}Credential("validation-test" /* auth fields from env, e.g., apiKey */),
    }
    validation.RunValidationSuite(t, cfg)
}
```

2. Local runs:

- make test # skips validation (short)
- make test-validation # runs validation (long)
- make test-all # runs everything

3. CI workflow (recommended):

- Add .github/workflows/validation-{provider}.yml (copy Lambda Labs workflow if available or follow VALIDATION_TESTING.md).
- Store secrets in GitHub Actions (e.g., YOUR_PROVIDER_API_KEY).

---

## Checklist

- [ ] Add internal/{provider}/v1 with client.go, instance.go, capabilities.go, instancetype.go
- [ ] Embed v1.NotImplCloudClient in client and only implement supported methods
- [ ] Accurately set Capabilities
- [ ] Implement instance types with stable IDs where possible
- [ ] Conform to security model; document provider-specific nuances
- [ ] Add validation_test.go and (optionally) CI workflow
- [ ] Run make lint and make test locally
- [ ] Add provider docs (README.md under provider folder) describing API mapping and feature coverage

---

## References

- Architecture: ../docs/ARCHITECTURE.md
- Security: ../docs/SECURITY.md
- Validation testing: ../docs/VALIDATION_TESTING.md
- CloudClient and composition: ../pkg/v1/client.go
- Capabilities: ../pkg/v1/capabilities.go
- Instance lifecycle and validations: ../pkg/v1/instance.go
- Instance types and validations: ../pkg/v1/instancetype.go
- Lambda Labs example:
  - Client/Credential: ../internal/lambdalabs/v1/client.go
  - Capabilities: ../internal/lambdalabs/v1/capabilities.go
  - Instance operations: ../internal/lambdalabs/v1/instance.go
  - Provider README: ../internal/lambdalabs/v1/README.md
