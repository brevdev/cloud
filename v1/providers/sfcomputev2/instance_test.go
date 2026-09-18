package v2

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

func TestNormalizeTerminateInstanceErrorTreatsNotFoundAsAlreadyGone(t *testing.T) {
	t.Parallel()

	providerErr := fmt.Errorf("release instance: %w", &apiError{
		statusCode: http.StatusNotFound,
		body:       `{"error":"not found"}`,
	})

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, v1.ErrInstanceNotFound)
	var responseErr *apiError
	require.False(t, errors.As(err, &responseErr))
}

func TestNormalizeTerminateInstanceErrorPreservesRetryableFailure(t *testing.T) {
	t.Parallel()

	providerErr := errors.New("provider rate limited")

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, providerErr)
}

func TestMakeSFCNameIsDeterministic(t *testing.T) {
	t.Parallel()

	tags := v1.Tags{"dev-plane-x-environmentId": "p82qfn5qs"}

	first := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)
	second := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)

	require.Equal(t, first, second)
}

func TestMakeSFCNameIsUniquePerInstance(t *testing.T) {
	t.Parallel()

	// Same environment, two instances: re-provisioning must not reuse the previous name.
	tags := v1.Tags{"dev-plane-x-environmentId": "p82qfn5qs"}

	first := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)
	second := makeSFCName("inst-2gAgPPd3QC1nmSjaWg8kc7UfyDx", tags)

	require.NotEqual(t, first, second)
}
