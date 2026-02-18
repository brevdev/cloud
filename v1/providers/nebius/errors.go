package v1

import (
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NebiusError represents a Nebius-specific error
type NebiusError struct {
	Code    codes.Code
	Message string
	Details string
}

func (e *NebiusError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("nebius error (code: %s): %s - %s", e.Code.String(), e.Message, e.Details)
	}
	return fmt.Sprintf("nebius error (code: %s): %s", e.Code.String(), e.Message)
}

// isNotFoundError checks if an error is a "not found" error
func isNotFoundError(err error) bool {
	// Check for gRPC NotFound status code
	if status, ok := status.FromError(err); ok {
		return status.Code() == codes.NotFound
	}
	return false
}

func handleErrToCloudErr(e error) error {
    if e == nil {
        return nil
    }

    // Check for gRPC ResourceExhausted status code
    if grpcStatus, ok := status.FromError(e); ok {
        if grpcStatus.Code() == codes.ResourceExhausted {
            return v1.ErrOutOfQuota
        }
    }
    return e
}


// isAlreadyExistsError checks if an error is an "already exists" error
//
//nolint:unused // Reserved for future error handling improvements
func isAlreadyExistsError(err error) bool {
	// Check for gRPC AlreadyExists status code
	if status, ok := status.FromError(err); ok {
		return status.Code() == codes.AlreadyExists
	}
	return false
}

// wrapNebiusError wraps a gRPC error into a NebiusError
//
//nolint:unused // Reserved for future error handling improvements
func wrapNebiusError(err error, context string) error {
	if err == nil {
		return nil
	}

	if grpcStatus, ok := status.FromError(err); ok {
		nebiusErr := &NebiusError{
			Code:    grpcStatus.Code(),
			Message: grpcStatus.Message(),
			Details: context,
		}
		return nebiusErr
	}

	// Return original error if not a gRPC error
	return err
}
