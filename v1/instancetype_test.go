package v1

import "testing"

func TestGetArchitectureAliases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
		want Architecture
	}{
		{name: "x86_64", raw: "x86_64", want: ArchitectureX86_64},
		{name: "amd64", raw: "amd64", want: ArchitectureX86_64},
		{name: "arm64", raw: "arm64", want: ArchitectureARM64},
		{name: "aarch64", raw: "aarch64", want: ArchitectureARM64},
		{name: "trim and case", raw: " AARCH64\n", want: ArchitectureARM64},
		{name: "unknown", raw: "riscv64", want: ArchitectureUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetArchitecture(tt.raw); got != tt.want {
				t.Fatalf("GetArchitecture(%q) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}
}

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
