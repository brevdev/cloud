# Hyperstack Security Notes

Hyperstack VM creation receives the requested SSH public key through an environment-scoped keypair. Key lookup follows all result pages before importing a new key. Port randomization is disabled so Brev can consistently use SSH port 22.

The provider creates each VM with a direct security rule allowing public IPv4 SSH ingress on port 22. Caller-provided TCP ingress ranges are added as direct per-VM rules. Outbound behavior is derived by Hyperstack and is not represented by explicit rules. All other inbound traffic remains denied.

Enhanced monitoring is explicitly disabled when creating VMs. This prevents opting the VM into the guest-installed Hyperstack VM Agent and its metrics ingestion gateway. Hyperstack's platform-level VM metrics are managed by the provider and do not expose a documented per-VM opt-out.

API authentication uses the `api_key` header over Hyperstack's HTTPS endpoint. API keys are never returned from provider methods; tenant identity is derived from a one-way hash of the key.
