# Brev Cloud SDK

Vendor-agnostic Go SDK for integrating cloud providers into Brev's GPU compute platform.

Cloud providers implement a standard set of interfaces — instance lifecycle, instance types, networking, storage — and Brev handles orchestration, inventory sync, and end-user experience.

## Who This Is For

- **Cloud providers** integrating GPU compute into Brev
- **NVIDIA Cloud Partners (NCPs)** offering Brev-compatible infrastructure
- **Compute brokers and marketplaces** aggregating multi-cloud GPU capacity

## Get Started

Three resources to go from zero to a working provider integration:

1. **[Cloud Manual](docs/CloudManual.md)** — Complete reference for the SDK's interfaces, types, capabilities, billing, and provider expectations.
2. **[Architecture](docs/ARCHITECTURE.md)** — How the Cloud SDK connects to Brev's control plane: inventory sync, provisioning flow, and credential management.
3. **[Integration Guide](docs/IntegrationGuide.md)** — Practical walkthrough to implement a new provider, with copy/paste scaffolding and validation tests.

## Additional Documentation

- [Security Requirements](docs/SECURITY.md) — Network model, SSH access, firewall rules, and encryption expectations.
- [Validation Testing](docs/VALIDATION_TESTING.md) — Shared test suite for verifying provider implementations against real APIs.
- [V1 Design Notes](v1/V1_DESIGN_NOTES.md) — Design decisions, known quirks, and AWS-inspired patterns in the v1 API.

## Status

- Version: `v1` — internal interface, open-sourced
- Cloud provider implementations are internal-only for now
- `v2` will be shaped by feedback and real-world integration experience

## Contributing

We welcome and encourage contributions. If you're building GPU compute infrastructure or working on a provider integration, open an issue or submit a PR — we want to build this with the community.
