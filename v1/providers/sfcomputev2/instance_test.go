package v2

import (
	"errors"
	"fmt"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/sfcompute/sfc-go/models/apierrors"
	"github.com/stretchr/testify/require"
)

func TestNormalizeTerminateInstanceErrorTreatsNotFoundAsAlreadyGone(t *testing.T) {
	t.Parallel()

	providerErr := fmt.Errorf("release instance: %w", &apierrors.NotFoundError{})

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, v1.ErrInstanceNotFound)
	var notFoundErr *apierrors.NotFoundError
	require.False(t, errors.As(err, &notFoundErr))
}

func TestNormalizeTerminateInstanceErrorPreservesRetryableFailure(t *testing.T) {
	t.Parallel()

	providerErr := errors.New("provider rate limited")

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, providerErr)
}
