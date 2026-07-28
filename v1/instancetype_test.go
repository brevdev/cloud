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
