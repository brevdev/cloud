# Massed Compute Provider

This package implements the minimal Brev Cloud v1 surface for Massed Compute.

Supported operations:

- Discover on-demand instance types and their available regions.
- Create, get, list, and terminate instances.
- Create or reuse SSH keys required at launch.
- Resolve image names through the API, defaulting to `Ubuntu Server 22.04 w/ drivers`.

The provider intentionally does not advertise stop/start, reboot, spot/preemptible instances, firewall mutation, tags, or storage resizing. Spot inventory entries are omitted and spot creation requests are rejected. GPU count, model, and network details are parsed from the product description; bracketed annotations are discarded after spot detection. Marketed VRAM is parsed when present and otherwise filled from a small model lookup table. System RAM, storage, and vCPU count come directly from the inventory specs.

## Generated API client

The generated client is committed under `gen/massedcompute`. Regenerate it from the version-pinned Massed Compute OpenAPI specification with:

```sh
make -C v1/providers/massedcompute generate-massedcompute-client
```

`openapi-v1.0.0.yaml` is the unmodified vendor specification. At generation time, the version-pinned `openapi-v1.0.0.patch` produces `openapi-v1.0.0.final.yaml` with the few repairs needed for strict validation: matching the single-instance operation to its actual `runningInstances` response envelope, declaring the omitted instance `{uuid}` parameter, correcting the launch request's required fields, relocating two request examples, and repairing two misplaced terminate-response fields. Patch application fails if a future vendor document no longer matches these exact locations.

## Live validation

Set `MASSED_COMPUTE_API_TOKEN` and optionally `MASSED_COMPUTE_API_URL`, then run:

```sh
go test -run TestValidationFunctions ./v1/providers/massedcompute
```
