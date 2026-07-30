package verda

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	verdago "github.com/verda-cloud/verdacloud-sdk-go/pkg/verda"
)

func wrapVerdaError(err error) error {
	if err == nil {
		return nil
	}

	var apiError *verdago.APIError
	if !errors.As(err, &apiError) {
		return fmt.Errorf("verda API request failed: %w", err)
	}

	message := strings.ToLower(apiError.Message + " " + apiError.Details)
	switch {
	case apiError.StatusCode == http.StatusNotFound:
		return fmt.Errorf("verda API request failed: %w", errors.Join(v1.ErrInstanceNotFound, err))
	case strings.Contains(message, "not enough resources"):
		return fmt.Errorf("verda API request failed: %w", errors.Join(v1.ErrInsufficientResources, err))
	case apiError.StatusCode == http.StatusTooManyRequests || apiError.StatusCode >= http.StatusInternalServerError:
		return fmt.Errorf("verda API request failed: %w", errors.Join(v1.ErrServiceUnavailable, err))
	case strings.Contains(message, "quota") || strings.Contains(message, "limit exceeded"):
		return fmt.Errorf("verda API request failed: %w", errors.Join(v1.ErrOutOfQuota, err))
	default:
		return fmt.Errorf("verda API request failed: %w", err)
	}
}
