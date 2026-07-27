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
