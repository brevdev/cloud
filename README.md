# Cloud SDK (v1000)

An early-stage, vendor-agnostic Go SDK for managing **clusterable, GPU-accelerated compute** across cloud providers.

---

## Project Goals

- Define a clean, minimal interface for cloud compute primitives:
  - `Instance`
  - `Storage`
  - `FirewallRule`
  - `InstanceType`
  - `Location`

- Enable **clusterable GPU workloads** across multiple providers, with shared semantics and L3 network guarantees. (WIP)

---

## Security

All cloud integrations must follow our [Security Requirements](SECURITY.md), which define:

- **Network Security**: Default "deny all inbound, allow all outbound" model
- **Cluster Security**: Internal instance communication with external isolation
- **Data Protection**: Encryption requirements for data at rest and in transit
- **Implementation Guidelines**: Security checklists for cloud provider integrations

See [SECURITY.md](docs/SECURITY.md) for complete security specifications and implementation requirements.

---

## Status

- Version: `v1` — internal interface, open-sourced
- Current scope: core types + interfaces + tests
- Cloud provider implementations are internal-only for now
- `v2` will be shaped by feedback and contributions from the community

## Platform Support

- **Operating System**: Currently supports Ubuntu 22 only
- **Architecture**: Designed for GPU-accelerated compute workloads
- **Access Method**: Requires SSH server and SSH key-based authentication
- **System Requirements**: Requires systemd to be running and accessible

---

## Who This Is For

- **NVIDIA Cloud Partners (NCPs)** looking to offer Brev-compatible GPU compute
- **Infra teams** building cluster-aware systems or abstractions on raw compute
- **Cloud providers** interested in contributing to a shared interface for accelerated compute
- **Compute brokers & marketplaces (aggregators)** offering multi-cloud compute

## Documentation

- **[V1 Design Notes](pkg/v1/V1_DESIGN_NOTES.md)**: Design decisions, known quirks, and AWS-inspired patterns in the v1 API
- **[Architecture Overview](docs/ARCHITECTURE.md)**: How the Cloud SDK fits into Brev's overall architecture
- **[Security Requirements](docs/SECURITY.md)**: Security specifications and implementation requirements
- **[How to Add a Provider](docs/how-to-add-a-provider.md)**: Step-by-step guide to implement a new cloud provider using the Lambda Labs example

---

## Get Involved

This is a foundation — we're opening it early to **learn with the community** and shape a clean, composable `v2`. If you're building GPU compute infrastructure or tooling, we'd love your input.

