# SFCompute V2 networking

Configurable ingress is optional. Existing credentials continue using the SSH
proxy and do not create firewalls. To request public IPv4 for new instances:

```go
credential := v2.NewSFCCredentialV2(refID, apiKey, organization, workspace)
credential.EnableConfigurableFirewall = true
```

The equivalent JSON field is `"enable_configurable_firewall": true`. Enable it
only after the SFCompute migration, API, and capacity reconciler have finished
deploying. Roll out this Brev integration afterward. The pool must advertise
`public_ipv4_skus`; the adapter selects only those SKUs. Missing support returns
an error before creating an instance or firewall.

The API key needs firewall read, create, write, and delete permissions in the
pool's workspace, in addition to the existing instance and pool permissions.
Each new instance gets its own firewall. TCP ports 22 and 2222 remain open for
SSH; additional port ranges apply to TCP and UDP. Sources must be public IPv4
CIDRs or `0.0.0.0/0`. Narrowing outbound traffic is unsupported. SFCompute's
firewall quotas apply, including the limit of 99 custom firewalls per workspace.

`GetInstance` and `ListInstances` return the public IP, SSH port 22, and stable
ingress rule IDs for revocation. Firewall edits preserve unrelated rules and
retry concurrent changes through `/integrations/brev/v1/firewalls`. Rule updates
require the version returned by that integration endpoint. Termination deletes
the attached firewall only when its ID matches the instance's Brev ownership tag.
A cleanup failure is returned to the caller; retrying termination retries
firewall deletion. Cleanup uses a separate bounded context
so cancellation of a create request does not cancel its firewall cleanup.

Turning the credential option off affects new instances. Existing public-IP
instances remain accessible and their firewalls can still be managed. Existing
SSH-proxy instances cannot acquire public networking through a rule update.
