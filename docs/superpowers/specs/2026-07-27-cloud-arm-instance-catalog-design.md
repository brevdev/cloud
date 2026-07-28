# Cloud ARM Instance Catalog Design

## Goal

Make the Cloud SDK's instance and image catalogs architecture-complete without
changing dev-plane's current x86-only policy.

Cloud providers must:

- return every known instance type when no architecture filter is supplied;
- populate `SupportedArchitectures` with truthful canonical values;
- honor a caller-provided architecture filter;
- report an unknown architecture as `unknown` instead of silently treating it
  as `x86_64`.

This preserves the intended boundary:

```text
Cloud reports provider capabilities.
Callers decide which capabilities their product currently supports.
```

Dev-plane will remain unchanged. Its current request excludes ARM, so ARM
instance types added by this change will not pass into its synchronized catalog.

## Scope

### Shared architecture classification

Add a small, case-insensitive helper that recognizes NVIDIA Grace products from
their GPU model:

- `GH*` is ARM64.
- `GB*` is ARM64.
- A plain `B*` model does not imply ARM64.

The helper should answer only whether a GPU model identifies a Grace system.
Each provider remains responsible for choosing the fallback architecture based
on its own catalog contract.

Use canonical Cloud values:

- `x86_64`
- `arm64`
- `unknown`

### Launchpad

Launchpad already supplies an authoritative `SystemArch` field and maps it to
the Cloud architecture values. No production mapping change is needed.

Add unit coverage proving:

- `SystemArchAMD64` becomes `x86_64`;
- `SystemArchARM64` becomes `arm64`;
- caller-provided ARM exclusions remove ARM catalog entries.

### Shadeform

Retain the existing Grace inference while moving it to the shared helper.
Normalize model casing before classification.

Add conversion coverage for:

- GH200 as ARM64;
- GB200 as ARM64;
- lowercase Grace model names as ARM64;
- H100 and B200 as x86_64.

### LambdaLabs

Replace the unconditional x86 assignment with architecture derived from the
parsed GPU model:

- recognized Grace products become ARM64;
- other currently supported Lambda instance types remain x86_64.

The existing architecture filter remains authoritative and must exclude a
GH200 type when ARM is excluded.

### FluidStack

FluidStack is defunct and no active credential remains. Leave its provider
adapter unchanged and exclude it from this work.

### Nebius instance types

Nebius does not currently expose an architecture field in its Platform API.
Use a provider-local mapping for the platform IDs that Cloud supports.

The currently documented Nebius platforms use Intel or AMD host CPUs, including
the B200 and B300 platforms, and must therefore report `x86_64`. An unrecognized
platform must report `unknown`.

Populate `SupportedArchitectures` during instance-type construction. Replace
the current type-name inference and include-only filtering with filtering over
the populated architecture metadata using `ArchitectureFilter.IsAllowed`, so
both include and exclude filters work.

Do not invent a synthetic Nebius ARM platform. Add a real ARM mapping only when
Nebius publishes an ARM platform identifier or exposes an authoritative
architecture field.

### Nebius images

Remove the catalog-level x86 restriction:

- an unfiltered image request returns all architectures;
- use the Nebius `ImageSpec.CpuArchitecture` enum as the authoritative source;
- retain labels and image-name parsing only as compatibility fallbacks;
- use `unknown` when no source identifies the architecture.

This changes image discovery only. Architecture-aware boot-image selection is
part of the later provisioning work.

### SFCompute and SFComputeV2

Make no production classification change. The current providers expose H100
and H200 offerings backed by x86-only image requirements, and their metadata
already reports `x86_64`.

Add focused regression coverage where practical so this is recorded as a
provider capability rather than an accidental ARM restriction.

Future SFCompute ARM support requires an authoritative Grace SKU plus
architecture-compatible image and create-path support; it is not part of this
catalog change.

### TestKube

Defer ARM advertising.

Safe TestKube ARM support requires:

- a separately named ARM instance type;
- `SupportedArchitectures: [arm64]`;
- a `kubernetes.io/arch: arm64` pod node selector;
- a published and verified `linux/arm64` image manifest.

The existing `test.ok.cpu` type remains x86-only. A single type advertising
both architectures would be unsafe because it could pass an x86 filter and
then schedule onto an ARM node.

## Data Flow

For Cloud callers:

```text
provider API response
  -> provider-specific architecture extraction
  -> canonical Cloud architecture
  -> optional ArchitectureFilter
  -> returned InstanceType or Image
```

An unfiltered call returns x86, ARM, and unknown entries. A filtered call
returns only entries whose architecture is allowed. dev-plane will continue to
send its existing x86-only filter and therefore will not receive ARM entries.

## Error and Unknown-Data Handling

- Do not fail an entire provider catalog because one item has an unknown
  architecture.
- Preserve that item as `unknown` when the provider otherwise supports it.
- Never default an unrecognized provider platform or image to x86.
- Continue returning provider/API errors using the existing error-wrapping
  conventions.

## Testing

Follow test-first development for every behavior change.

Focused tests will cover:

- shared Grace GPU classification, including casing and B200 non-ARM behavior;
- Launchpad authoritative AMD64 and ARM64 conversion;
- Shadeform GH/GB conversion;
- LambdaLabs GH200 conversion and architecture filtering;
- Nebius known x86 and unknown platform mappings;
- Nebius include and exclude filter behavior;
- Nebius authoritative AMD64/ARM64 image metadata and unfiltered image listing;
- existing SFCompute x86 metadata where a focused unit seam exists.

Run package-level tests while iterating, then:

```bash
go test ./v1/...
go test ./...
```

Run `gofmt` on every touched Go file and run the repository's configured
linting if practical.

## Non-Goals

- No dev-plane changes.
- No removal of the public `ArchitectureFilter` contract.
- No enabling ARM instance creation in dev-plane.
- No architecture-aware provider boot-image selection beyond image catalog
  metadata.
- No TestKube image publishing or ARM scheduling change.
- No Verb-mode changes.
