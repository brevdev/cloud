# Massed Compute Provider

This package implements the minimal Brev Cloud v1 surface for Massed Compute.

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
