package v2

import (
	"testing"

	"github.com/alecthomas/units"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

func TestH100InstanceTypeReports80GiBGPUVRAM(t *testing.T) {
	instanceType := buildInstanceType(h100InstanceTypeMetadata, true)
	require.Len(t, instanceType.SupportedGPUs, 1)

	gpu := instanceType.SupportedGPUs[0]
	require.Equal(t, 80*units.Gibibyte, gpu.Memory)
	require.True(t, gpu.MemoryBytes.Equal(v1.NewBytes(80, v1.Gibibyte)))
}
