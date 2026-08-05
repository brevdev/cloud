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
