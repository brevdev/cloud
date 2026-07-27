# Cloud ARM Instance Catalog Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make Cloud return ARM-capable instance and image catalog entries with truthful canonical architecture metadata while preserving dev-plane's existing ARM exclusion.

**Architecture:** Keep architecture discovery inside each provider adapter, normalize shared NVIDIA Grace model detection in `v1`, and apply the existing caller-owned architecture filter after metadata is populated. Providers with authoritative architecture fields retain those fields as the source of truth; inferred or mapped providers return `unknown` when the architecture cannot be established.

**Tech Stack:** Go 1.25, generated provider SDK clients, `testify`, package-level Go tests, `gofmt`, and `golangci-lint`.

## Global Constraints

- The approved design is
  [`docs/superpowers/specs/2026-07-27-cloud-arm-instance-catalog-design.md`](../specs/2026-07-27-cloud-arm-instance-catalog-design.md).
- Do not modify dev-plane. Its existing `ArchitectureFilter` continues to exclude
  ARM from the synchronized catalog.
- Do not change or remove the public `ArchitectureFilter` contract.
- Do not enable ARM provisioning, image selection, TestKube ARM scheduling, or
  Verb mode.
- An unfiltered Cloud call returns all known entries. A filtered call honors the
  caller's filter.
- Emit only canonical values: `x86_64`, `arm64`, or `unknown`.
- Do not classify B200 or B300 as ARM merely because they are Blackwell GPUs.
- Nebius's current documented platform IDs are mapped explicitly. The current
  platform table documents Intel or AMD host CPUs for all of them:
  <https://docs.nebius.com/compute/virtual-machines/types>.
- Follow red-green-refactor for every behavior change. Characterization tests
  for already-correct providers are expected to pass immediately.
- Run `gofmt` on every touched Go file. Do not edit generated provider clients.

## File and Interface Map

| Area | Existing contract | Planned touchpoint |
| --- | --- | --- |
| Shared types | `v1.Architecture`, `ArchitectureFilter.IsAllowed` | Add shared Grace model classifier |
| Launchpad | Authoritative `SystemArch` | Add mapping and filter characterization tests |
| Shadeform | GPU-name inference | Use shared classifier |
| LambdaLabs | Parsed GPU description, currently hardcoded x86 | Classify parsed Grace GPU models |
| FluidStack | `GpuModel`, currently no architecture metadata/filter | Populate metadata and honor only architecture filtering |
| Nebius instances | Platform IDs, name-based fallback to x86 | Explicit platform map, metadata-driven include/exclude filtering |
| Nebius images | `ImageSpec.CpuArchitecture`, currently ignored | Prefer enum, remove unfiltered x86 default |
| SFCompute v1/v2 | Static H100/H200 x86 metadata | Add regression coverage only |

---

### Task 1: Add the shared NVIDIA Grace classifier

**Files:**

- Modify: `v1/instancetype.go`
- Modify: `v1/instancetype_test.go`

- [ ] **Step 1: Write the failing classifier test**

Append this table test to `v1/instancetype_test.go`:

```go
func TestIsNVIDIAGraceGPU(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		gpuModel string
		want     bool
	}{
		{name: "GH200", gpuModel: "GH200", want: true},
		{name: "GB200", gpuModel: "GB200", want: true},
		{name: "lowercase GH200", gpuModel: "gh200", want: true},
		{name: "trimmed GB200", gpuModel: "  gb200  ", want: true},
		{name: "vendor-prefixed GH200", gpuModel: "NVIDIA GH200", want: true},
		{name: "H100", gpuModel: "H100", want: false},
		{name: "B200", gpuModel: "B200", want: false},
		{name: "empty", gpuModel: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsNVIDIAGraceGPU(tt.gpuModel); got != tt.want {
				t.Fatalf("IsNVIDIAGraceGPU(%q) = %v, want %v", tt.gpuModel, got, tt.want)
			}
		})
	}
}
```

- [ ] **Step 2: Run the focused test and confirm the red state**

Run:

```bash
go test ./v1 -run '^TestIsNVIDIAGraceGPU$' -count=1
```

Expected: compilation fails because `IsNVIDIAGraceGPU` is undefined.

- [ ] **Step 3: Implement the smallest shared classifier**

Add this function beside the existing architecture normalization helpers in
`v1/instancetype.go`:

```go
func IsNVIDIAGraceGPU(gpuModel string) bool {
	normalizedModel := strings.ToUpper(strings.TrimSpace(gpuModel))
	normalizedModel = strings.TrimPrefix(normalizedModel, "NVIDIA ")
	return strings.HasPrefix(normalizedModel, "GH") ||
		strings.HasPrefix(normalizedModel, "GB")
}
```

The function deliberately answers only whether the model is a Grace product.
It does not choose a fallback architecture.

- [ ] **Step 4: Format and prove the green state**

Run:

```bash
gofmt -w v1/instancetype.go v1/instancetype_test.go
go test ./v1 -run '^(TestGetArchitectureAliases|TestIsNVIDIAGraceGPU)$' -count=1
```

Expected: both tests pass.

- [ ] **Step 5: Commit the shared behavior**

```bash
git add v1/instancetype.go v1/instancetype_test.go
git commit -m "feat: classify NVIDIA Grace GPU models"
```

---

### Task 2: Lock Launchpad's authoritative architecture behavior

**Files:**

- Modify: `v1/providers/launchpad/instancetype_test.go`
- Verify only: `v1/providers/launchpad/instancetype.go`

- [ ] **Step 1: Add mapping and caller-filter characterization tests**

Add the generated Launchpad package import:

```go
openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
```

Then append:

```go
func TestLaunchpadArchitectureToArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  openapi.SystemArchEnum
		want v1.Architecture
	}{
		{name: "AMD64", raw: openapi.SystemArchAMD64, want: v1.ArchitectureX86_64},
		{name: "ARM64", raw: openapi.SystemArchARM64, want: v1.ArchitectureARM64},
		{name: "unknown", raw: openapi.SystemArchEnum("riscv64"), want: v1.ArchitectureUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, launchpadArchitectureToArchitecture(tt.raw))
		})
	}
}

func TestLaunchpadArchitectureFilterExcludesARM(t *testing.T) {
	t.Parallel()

	instanceType := v1.InstanceType{
		SupportedArchitectures: []v1.Architecture{
			launchpadArchitectureToArchitecture(openapi.SystemArchARM64),
		},
	}
	args := v1.GetInstanceTypeArgs{
		ArchitectureFilter: &v1.ArchitectureFilter{
			ExcludeArchitectures: []v1.Architecture{v1.ArchitectureARM64},
		},
	}

	require.False(t, v1.IsSelectedByArgs(instanceType, args))
}
```

- [ ] **Step 2: Run the characterization tests**

Run:

```bash
go test ./v1/providers/launchpad -run '^TestLaunchpadArchitecture' -count=1
```

Expected: tests pass without production changes because Launchpad already uses
its authoritative `SystemArch` value and the common filter.

- [ ] **Step 3: Commit the contract tests**

```bash
git add v1/providers/launchpad/instancetype_test.go
git commit -m "test: lock Launchpad architecture mapping"
```

---

### Task 3: Move Shadeform Grace inference to the shared classifier

**Files:**

- Modify: `v1/providers/shadeform/instancetype.go`
- Modify: `v1/providers/shadeform/instancetype_test.go`

- [ ] **Step 1: Write the failing case-normalization test**

Append:

```go
func TestShadeformArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gpuModel string
		want     v1.Architecture
	}{
		{gpuModel: "GH200", want: v1.ArchitectureARM64},
		{gpuModel: "GB200", want: v1.ArchitectureARM64},
		{gpuModel: "gh200", want: v1.ArchitectureARM64},
		{gpuModel: "gb200", want: v1.ArchitectureARM64},
		{gpuModel: "H100", want: v1.ArchitectureX86_64},
		{gpuModel: "B200", want: v1.ArchitectureX86_64},
	}

	for _, tt := range tests {
		t.Run(tt.gpuModel, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, shadeformArchitecture(tt.gpuModel))
		})
	}
}
```

- [ ] **Step 2: Run the focused test and confirm the red state**

Run:

```bash
go test ./v1/providers/shadeform -run '^TestShadeformArchitecture$' -count=1
```

Expected: lowercase GH200 and GB200 cases fail because the current check is
case-sensitive.

- [ ] **Step 3: Delegate classification to the shared helper**

Replace `shadeformArchitecture` with:

```go
func shadeformArchitecture(gpuName string) v1.Architecture {
	if v1.IsNVIDIAGraceGPU(gpuName) {
		return v1.ArchitectureARM64
	}
	return v1.ArchitectureX86_64
}
```

Keep the provider-owned x86 fallback; only recognition is shared.

- [ ] **Step 4: Format and run the package tests**

Run:

```bash
gofmt -w v1/providers/shadeform/instancetype.go v1/providers/shadeform/instancetype_test.go
go test ./v1/providers/shadeform -count=1
```

Expected: all Shadeform tests pass.

- [ ] **Step 5: Commit the provider change**

```bash
git add v1/providers/shadeform/instancetype.go v1/providers/shadeform/instancetype_test.go
git commit -m "feat: normalize Shadeform Grace architectures"
```

---

### Task 4: Classify LambdaLabs Grace offerings and preserve filtering

**Files:**

- Modify: `v1/providers/lambdalabs/instancetype.go`
- Modify: `v1/providers/lambdalabs/instancetype_test.go`

- [ ] **Step 1: Write failing conversion and end-to-end filter tests**

Append these tests:

```go
func TestConvertLambdaLabsInstanceTypeArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		description string
		want        v1.Architecture
	}{
		{name: "GH200", description: "1x GH200 (96 GB)", want: v1.ArchitectureARM64},
		{name: "NVIDIA GH200", description: "1x NVIDIA GH200 (96 GB)", want: v1.ArchitectureARM64},
		{name: "H100", description: "1x H100 (80 GB SXM5)", want: v1.ArchitectureX86_64},
		{name: "B200", description: "1x B200 (192 GB SXM6)", want: v1.ArchitectureX86_64},
		{name: "CPU", description: "4x CPU cores", want: v1.ArchitectureX86_64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			instanceType := createMockLambdaLabsInstanceType(
				"test-"+tt.name,
				tt.description,
				tt.name,
				100,
			)

			got, err := convertLambdaLabsInstanceTypeToV1InstanceType(
				"us-west-1",
				instanceType,
				true,
			)
			require.NoError(t, err)
			require.Equal(t, []v1.Architecture{tt.want}, got.SupportedArchitectures)
		})
	}
}

func TestLambdaLabsClient_GetInstanceTypes_ExcludeARM(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockResponse := createMockInstanceTypeResponse()
	mockResponse.Data["gpu_1x_gh200"] = openapi.InstanceTypes200ResponseDataValue{
		InstanceType: createMockLambdaLabsInstanceType(
			"gpu_1x_gh200",
			"1x GH200 (96 GB)",
			"GH200",
			229,
		),
		RegionsWithCapacityAvailable: []openapi.Region{
			createMockRegion("us-west-1", "US West 1"),
		},
	}
	httpmock.RegisterResponder(
		"GET",
		"https://cloud.lambda.ai/api/v1/instance-types",
		httpmock.NewJsonResponderOrPanic(200, mockResponse),
	)

	instanceTypes, err := client.GetInstanceTypes(
		context.Background(),
		v1.GetInstanceTypeArgs{
			ArchitectureFilter: &v1.ArchitectureFilter{
				ExcludeArchitectures: []v1.Architecture{v1.ArchitectureARM64},
			},
		},
	)
	require.NoError(t, err)
	require.Nil(t, findInstanceTypeByName(instanceTypes, "gpu_1x_gh200"))
	require.NotNil(t, findInstanceTypeByName(instanceTypes, "gpu_1x_a10"))
}
```

- [ ] **Step 2: Run the focused tests and confirm the red state**

Run:

```bash
go test ./v1/providers/lambdalabs -run 'Architecture|ExcludeARM' -count=1
```

Expected: GH200 is reported as x86 and survives the ARM exclusion.

- [ ] **Step 3: Derive architecture from the parsed GPU model**

After parsing `gpus` and before constructing the returned instance type, add:

```go
architecture := v1.ArchitectureX86_64
if len(gpus) > 0 && v1.IsNVIDIAGraceGPU(gpus[0].Name) {
	architecture = v1.ArchitectureARM64
}
```

Then replace the hardcoded field with:

```go
SupportedArchitectures: []v1.Architecture{architecture},
```

CPU-only and non-Grace offerings deliberately retain the provider-owned x86
fallback.

- [ ] **Step 4: Format and run the complete LambdaLabs package**

Run:

```bash
gofmt -w v1/providers/lambdalabs/instancetype.go v1/providers/lambdalabs/instancetype_test.go
go test ./v1/providers/lambdalabs -count=1
```

Expected: all LambdaLabs tests pass.

- [ ] **Step 5: Commit the provider change**

```bash
git add v1/providers/lambdalabs/instancetype.go v1/providers/lambdalabs/instancetype_test.go
git commit -m "feat: report LambdaLabs Grace architectures"
```

---

### Task 5: Populate and filter FluidStack architecture metadata

**Files:**

- Modify: `v1/providers/fluidstack/instancetype.go`
- Create: `v1/providers/fluidstack/instancetype_test.go`

- [ ] **Step 1: Write failing converter and filter tests**

Create `v1/providers/fluidstack/instancetype_test.go`:

```go
package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/fluidstack/gen/fluidstack"
)

func TestConvertFluidStackInstanceTypeArchitecture(t *testing.T) {
	t.Parallel()

	gh200 := "gh200"
	b200 := "B200"
	tests := []struct {
		name     string
		gpuModel *string
		want     cloudv1.Architecture
	}{
		{name: "GH200", gpuModel: &gh200, want: cloudv1.ArchitectureARM64},
		{name: "B200", gpuModel: &b200, want: cloudv1.ArchitectureX86_64},
		{name: "CPU", gpuModel: nil, want: cloudv1.ArchitectureX86_64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gpuCount := int32(1)
			if tt.gpuModel == nil {
				gpuCount = 0
			}
			got := convertFluidStackInstanceTypeToV1InstanceType(
				"",
				openapi.InstanceType{
					Name:     "test-" + tt.name,
					Cpu:      64,
					Memory:   "432GB",
					GpuModel: tt.gpuModel,
					GpuCount: &gpuCount,
				},
				true,
			)
			require.Equal(t, []cloudv1.Architecture{tt.want}, got.SupportedArchitectures)
		})
	}
}

func TestFilterFluidStackInstanceTypesByArchitecture(t *testing.T) {
	t.Parallel()

	x86Type := cloudv1.InstanceType{
		Type:                   "x86",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureX86_64},
	}
	armType := cloudv1.InstanceType{
		Type:                   "arm",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
	}
	instanceTypes := []cloudv1.InstanceType{x86Type, armType}

	require.Equal(
		t,
		instanceTypes,
		filterInstanceTypesByArchitecture(instanceTypes, nil),
	)
	require.Equal(
		t,
		[]cloudv1.InstanceType{x86Type},
		filterInstanceTypesByArchitecture(
			instanceTypes,
			&cloudv1.ArchitectureFilter{
				ExcludeArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
			},
		),
	)
}
```

- [ ] **Step 2: Run the focused tests and confirm the red state**

Run:

```bash
go test ./v1/providers/fluidstack -run 'Architecture' -count=1
```

Expected: the filter helper is undefined and the converter emits no
`SupportedArchitectures`.

- [ ] **Step 3: Populate architecture metadata during conversion**

Before the returned `v1.InstanceType` literal, add:

```go
architecture := v1.ArchitectureX86_64
if fsInstanceType.GpuModel != nil &&
	v1.IsNVIDIAGraceGPU(*fsInstanceType.GpuModel) {
	architecture = v1.ArchitectureARM64
}
```

Add this field to the returned literal:

```go
SupportedArchitectures: []v1.Architecture{architecture},
```

- [ ] **Step 4: Add architecture-only filtering**

Change the `GetInstanceTypes` signature to retain `args`:

```go
func (c *FluidStackClient) GetInstanceTypes(
	ctx context.Context,
	args v1.GetInstanceTypeArgs,
) ([]v1.InstanceType, error) {
```

Return the architecture-filtered values:

```go
return filterInstanceTypesByArchitecture(
	instanceTypes,
	args.ArchitectureFilter,
), nil
```

Add this provider-local helper:

```go
func filterInstanceTypesByArchitecture(
	instanceTypes []v1.InstanceType,
	filter *v1.ArchitectureFilter,
) []v1.InstanceType {
	if filter == nil {
		return instanceTypes
	}

	filtered := make([]v1.InstanceType, 0, len(instanceTypes))
	for _, instanceType := range instanceTypes {
		for _, architecture := range instanceType.SupportedArchitectures {
			if filter.IsAllowed(architecture) {
				filtered = append(filtered, instanceType)
				break
			}
		}
	}
	return filtered
}
```

Do not add location or instance-name filtering in this task.

- [ ] **Step 5: Format and run the complete FluidStack package**

Run:

```bash
gofmt -w v1/providers/fluidstack/instancetype.go v1/providers/fluidstack/instancetype_test.go
go test ./v1/providers/fluidstack -count=1
```

Expected: all FluidStack tests pass.

- [ ] **Step 6: Commit the provider change**

```bash
git add v1/providers/fluidstack/instancetype.go v1/providers/fluidstack/instancetype_test.go
git commit -m "feat: expose FluidStack instance architectures"
```

---

### Task 6: Replace Nebius name inference with explicit platform metadata

**Files:**

- Modify: `v1/providers/nebius/instancetype.go`
- Create: `v1/providers/nebius/instancetype_test.go`

- [ ] **Step 1: Write failing platform-map and filter tests**

Create `v1/providers/nebius/instancetype_test.go`:

```go
package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestNebiusPlatformArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		platform string
		want     cloudv1.Architecture
	}{
		{platform: "gpu-b300-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-b200-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-b200-sxm-a", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-h200-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-h100-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-rtx6000", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-l40s-a", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-l40s-d", want: cloudv1.ArchitectureX86_64},
		{platform: "cpu-d3", want: cloudv1.ArchitectureX86_64},
		{platform: "cpu-e2", want: cloudv1.ArchitectureX86_64},
		{platform: "future-platform", want: cloudv1.ArchitectureUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.platform, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, nebiusPlatformArchitecture(tt.platform))
		})
	}
}

func TestApplyInstanceTypeFiltersUsesArchitectureMetadata(t *testing.T) {
	t.Parallel()

	x86Type := cloudv1.InstanceType{
		Type:                   "x86-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureX86_64},
	}
	armType := cloudv1.InstanceType{
		Type:                   "arm-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
	}
	unknownType := cloudv1.InstanceType{
		Type:                   "unknown-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureUnknown},
	}
	all := []cloudv1.InstanceType{x86Type, armType, unknownType}

	tests := []struct {
		name string
		args cloudv1.GetInstanceTypeArgs
		want []cloudv1.InstanceType
	}{
		{name: "unfiltered", want: all},
		{
			name: "include x86",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureX86_64,
					},
				},
			},
			want: []cloudv1.InstanceType{x86Type},
		},
		{
			name: "include ARM",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
			want: []cloudv1.InstanceType{armType},
		},
		{
			name: "exclude ARM",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					ExcludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
			want: []cloudv1.InstanceType{x86Type, unknownType},
		},
		{
			name: "exclude x86",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					ExcludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureX86_64,
					},
				},
			},
			want: []cloudv1.InstanceType{armType, unknownType},
		},
	}

	client := &NebiusClient{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, client.applyInstanceTypeFilters(all, tt.args))
		})
	}
}
```

- [ ] **Step 2: Run the focused tests and confirm the red state**

Run:

```bash
go test ./v1/providers/nebius -run 'NebiusPlatformArchitecture|ApplyInstanceTypeFiltersUsesArchitectureMetadata' -count=1
```

Expected: `nebiusPlatformArchitecture` is undefined, and exclusion cases expose
the current include-only, type-name-based behavior.

- [ ] **Step 3: Add the explicit provider-local platform map**

Add near the existing platform helpers:

```go
var nebiusPlatformArchitectures = map[string]v1.Architecture{
	"gpu-b300-sxm":  v1.ArchitectureX86_64,
	"gpu-b200-sxm":  v1.ArchitectureX86_64,
	"gpu-b200-sxm-a": v1.ArchitectureX86_64,
	"gpu-h200-sxm":  v1.ArchitectureX86_64,
	"gpu-h100-sxm":  v1.ArchitectureX86_64,
	"gpu-rtx6000":   v1.ArchitectureX86_64,
	"gpu-l40s-a":    v1.ArchitectureX86_64,
	"gpu-l40s-d":    v1.ArchitectureX86_64,
	"cpu-d3":         v1.ArchitectureX86_64,
	"cpu-e2":         v1.ArchitectureX86_64,
}

func nebiusPlatformArchitecture(platformName string) v1.Architecture {
	architecture, ok := nebiusPlatformArchitectures[
		strings.ToLower(strings.TrimSpace(platformName))
	]
	if !ok {
		return v1.ArchitectureUnknown
	}
	return architecture
}
```

Use exact platform IDs rather than GPU-generation inference. This is why B200
and B300 map to x86.

- [ ] **Step 4: Populate `SupportedArchitectures` at construction**

In the `v1.InstanceType` literal inside `getInstanceTypesForLocation`, add:

```go
SupportedArchitectures: []v1.Architecture{
	nebiusPlatformArchitecture(platform.Metadata.Name),
},
```

- [ ] **Step 5: Replace custom architecture filtering**

Add:

```go
func supportsAllowedArchitecture(
	instanceType v1.InstanceType,
	filter *v1.ArchitectureFilter,
) bool {
	for _, architecture := range instanceType.SupportedArchitectures {
		if filter.IsAllowed(architecture) {
			return true
		}
	}
	return false
}
```

Replace the current architecture block in `applyInstanceTypeFilters` with:

```go
if args.ArchitectureFilter != nil &&
	!supportsAllowedArchitecture(instanceType, args.ArchitectureFilter) {
	continue
}
```

Delete `determineInstanceTypeArchitecture`. Do not add a synthetic ARM platform.

- [ ] **Step 6: Format and run the complete Nebius package**

Run:

```bash
gofmt -w v1/providers/nebius/instancetype.go v1/providers/nebius/instancetype_test.go
go test ./v1/providers/nebius -count=1
```

Expected: all Nebius tests pass. Unknown platforms remain present in an
unfiltered result and are handled according to the caller's filter.

- [ ] **Step 7: Commit the provider change**

```bash
git add v1/providers/nebius/instancetype.go v1/providers/nebius/instancetype_test.go
git commit -m "feat: report Nebius platform architectures"
```

---

### Task 7: Make the Nebius image catalog architecture-complete

**Files:**

- Modify: `v1/providers/nebius/image.go`
- Create: `v1/providers/nebius/image_test.go`

- [ ] **Step 1: Write failing extraction and unfiltered-list tests**

Create `v1/providers/nebius/image_test.go`:

```go
package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestExtractArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		image *compute.Image
		want  string
	}{
		{
			name: "authoritative AMD64 enum",
			image: &compute.Image{
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_AMD64,
				},
			},
			want: string(cloudv1.ArchitectureX86_64),
		},
		{
			name: "authoritative ARM64 enum",
			image: &compute.Image{
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_ARM64,
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "enum wins over legacy label",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{
					Labels: map[string]string{"architecture": "amd64"},
				},
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_ARM64,
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "legacy label is canonicalized",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{
					Labels: map[string]string{"arch": "aarch64"},
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "legacy name fallback",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{Name: "ubuntu-24.04-amd64"},
			},
			want: string(cloudv1.ArchitectureX86_64),
		},
		{
			name:  "unknown stays unknown",
			image: &compute.Image{},
			want:  string(cloudv1.ArchitectureUnknown),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, extractArchitecture(tt.image))
		})
	}
}

func TestApplyImageFiltersWithoutArchitectureReturnsAll(t *testing.T) {
	t.Parallel()

	images := []cloudv1.Image{
		{ID: "amd64", Architecture: string(cloudv1.ArchitectureX86_64)},
		{ID: "arm64", Architecture: string(cloudv1.ArchitectureARM64)},
		{ID: "unknown", Architecture: string(cloudv1.ArchitectureUnknown)},
	}

	require.Equal(
		t,
		images,
		applyImageFilters(images, cloudv1.GetImageArgs{}),
	)
}
```

- [ ] **Step 2: Run the focused tests and confirm the red state**

Run:

```bash
go test ./v1/providers/nebius -run 'ExtractArchitecture|ApplyImageFilters' -count=1
```

Expected: the authoritative enum and unknown cases fail, and
`applyImageFilters` is undefined.

- [ ] **Step 3: Make the provider enum authoritative**

Replace `extractArchitecture` with:

```go
func extractArchitecture(image *compute.Image) string {
	if image != nil && image.Spec != nil {
		switch image.Spec.CpuArchitecture {
		case compute.ImageSpec_AMD64:
			return string(v1.ArchitectureX86_64)
		case compute.ImageSpec_ARM64:
			return string(v1.ArchitectureARM64)
		}
	}

	if image != nil && image.Metadata != nil {
		for _, label := range []string{"architecture", "arch"} {
			rawArchitecture, ok := image.Metadata.Labels[label]
			if !ok {
				continue
			}
			architecture := v1.GetArchitecture(rawArchitecture)
			if architecture != v1.ArchitectureUnknown {
				return string(architecture)
			}
		}

		name := strings.ToLower(image.Metadata.Name)
		if strings.Contains(name, ArchitectureArm64) ||
			strings.Contains(name, ArchitectureAArch64) {
			return string(v1.ArchitectureARM64)
		}
		if strings.Contains(name, ArchitectureX86_64) ||
			strings.Contains(name, ArchitectureAMD64) {
			return string(v1.ArchitectureX86_64)
		}
	}

	return string(v1.ArchitectureUnknown)
}
```

- [ ] **Step 4: Centralize and remove the default x86 filter**

Add:

```go
func applyImageFilters(
	images []v1.Image,
	args v1.GetImageArgs,
) []v1.Image {
	images = filterImagesByArchitectures(images, args.Architectures)
	return filterImagesByNameFilters(images, args.NameFilters)
}
```

In `GetImages`, delete the block that substitutes `[]string{"x86_64"}` when
`args.Architectures` is empty. Replace both filter blocks with:

```go
return applyImageFilters(images, args), nil
```

In `getDefaultImages`, replace the hardcoded architecture with:

```go
Architecture: extractArchitecture(image),
```

- [ ] **Step 5: Format and run the complete Nebius package**

Run:

```bash
gofmt -w v1/providers/nebius/image.go v1/providers/nebius/image_test.go
go test ./v1/providers/nebius -count=1
```

Expected: all Nebius tests pass, and an empty architecture filter preserves
AMD64, ARM64, and unknown images.

- [ ] **Step 6: Commit the image-catalog change**

```bash
git add v1/providers/nebius/image.go v1/providers/nebius/image_test.go
git commit -m "feat: expose Nebius image architectures"
```

---

### Task 8: Record SFCompute's current x86-only provider contract

**Files:**

- Create: `v1/providers/sfcompute/instancetype_test.go`
- Create: `v1/providers/sfcomputev2/instancetype_test.go`
- Verify only: `v1/providers/sfcompute/instancetype.go`
- Verify only: `v1/providers/sfcomputev2/instancetype.go`

- [ ] **Step 1: Add SFCompute v1 characterization coverage**

Create `v1/providers/sfcompute/instancetype_test.go`:

```go
package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestSFComputeInstanceTypeArchitectures(t *testing.T) {
	t.Parallel()

	for _, gpuType := range []string{gpuTypeH100, gpuTypeH200} {
		t.Run(gpuType, func(t *testing.T) {
			t.Parallel()
			metadata, err := getInstanceTypeMetadata(gpuType)
			require.NoError(t, err)
			require.Equal(t, cloudv1.ArchitectureX86_64, metadata.architecture)
		})
	}
}
```

- [ ] **Step 2: Add SFCompute v2 characterization and filter coverage**

Create `v1/providers/sfcomputev2/instancetype_test.go`:

```go
package v2

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestSFComputeV2InstanceTypeArchitecture(t *testing.T) {
	t.Parallel()

	instanceType := buildInstanceType(h100InstanceTypeMetadata, true)
	require.Equal(
		t,
		[]cloudv1.Architecture{cloudv1.ArchitectureX86_64},
		instanceType.SupportedArchitectures,
	)
	require.False(
		t,
		cloudv1.IsSelectedByArgs(
			instanceType,
			cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
		),
	)
}
```

- [ ] **Step 3: Run both characterization packages**

Run:

```bash
gofmt -w v1/providers/sfcompute/instancetype_test.go v1/providers/sfcomputev2/instancetype_test.go
go test ./v1/providers/sfcompute ./v1/providers/sfcomputev2 -count=1
```

Expected: tests pass without production changes. H100 and H200 remain truthfully
x86 until SFCompute exposes an authoritative Grace SKU and compatible image.

- [ ] **Step 4: Commit the contract tests**

```bash
git add v1/providers/sfcompute/instancetype_test.go v1/providers/sfcomputev2/instancetype_test.go
git commit -m "test: record SFCompute architecture contracts"
```

---

### Task 9: Run cross-provider verification and inspect the final diff

**Files:**

- Verify all files changed by Tasks 1 through 8

- [ ] **Step 1: Format every touched Go file**

Run:

```bash
gofmt -w \
  v1/instancetype.go \
  v1/instancetype_test.go \
  v1/providers/launchpad/instancetype_test.go \
  v1/providers/shadeform/instancetype.go \
  v1/providers/shadeform/instancetype_test.go \
  v1/providers/lambdalabs/instancetype.go \
  v1/providers/lambdalabs/instancetype_test.go \
  v1/providers/fluidstack/instancetype.go \
  v1/providers/fluidstack/instancetype_test.go \
  v1/providers/nebius/instancetype.go \
  v1/providers/nebius/instancetype_test.go \
  v1/providers/nebius/image.go \
  v1/providers/nebius/image_test.go \
  v1/providers/sfcompute/instancetype_test.go \
  v1/providers/sfcomputev2/instancetype_test.go
```

- [ ] **Step 2: Run the focused cross-provider suite**

Run:

```bash
go test \
  ./v1 \
  ./v1/providers/launchpad \
  ./v1/providers/shadeform \
  ./v1/providers/lambdalabs \
  ./v1/providers/fluidstack \
  ./v1/providers/nebius \
  ./v1/providers/sfcompute \
  ./v1/providers/sfcomputev2 \
  -count=1
```

Expected: all focused packages pass.

- [ ] **Step 3: Run repository-wide tests**

Run:

```bash
go test ./v1/... -count=1
go test ./... -count=1
```

Expected: both commands pass. Investigate any failure before claiming completion;
do not dismiss provider failures without reproducing the exact package.

- [ ] **Step 4: Run formatting and lint checks**

Run:

```bash
make fmt-check
golangci-lint run ./...
```

Expected: both commands pass. If `golangci-lint` is not installed, use the
repository-pinned version from `Makefile` rather than changing tool versions.

- [ ] **Step 5: Verify scope and semantics in the diff**

Run:

```bash
git diff --check
git status --short
git diff --stat HEAD~8..HEAD
git diff HEAD~8..HEAD -- \
  v1/instancetype.go \
  v1/providers/launchpad \
  v1/providers/shadeform \
  v1/providers/lambdalabs \
  v1/providers/fluidstack \
  v1/providers/nebius \
  v1/providers/sfcompute \
  v1/providers/sfcomputev2
```

Confirm:

- no dev-plane or Verb files changed;
- no generated clients changed;
- Launchpad still uses `SystemArch`;
- Shadeform, LambdaLabs, and FluidStack share only Grace recognition;
- FluidStack added only architecture filtering;
- Nebius unknown platforms and images remain `unknown`;
- no synthetic Nebius or TestKube ARM offering was introduced;
- ARM provisioning and boot-image selection remain unchanged.

- [ ] **Step 6: Commit any verification-only formatting adjustment**

Only if Step 1 changed a file after its task commit:

```bash
git add v1
git commit -m "style: format ARM catalog changes"
```

If `git status --short` is empty, do not create an empty commit.
