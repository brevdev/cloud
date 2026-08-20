package v2

import (
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

func TestH100InstanceTypeUsesNominalGBForVRAM(t *testing.T) {
	t.Parallel()

	instanceType := buildInstanceType(h100InstanceTypeMetadata, true)

	require.Equal(t, v1.NewBytes(80, v1.Gigabyte), instanceType.SupportedGPUs[0].MemoryBytes)
}
