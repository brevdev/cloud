package massedcompute

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"
)

func wrapMassedComputeError(err error, response *http.Response) error {
	if err == nil {
		return nil
	}

	message := err.Error()
	var apiError *openapi.GenericOpenAPIError
	if errors.As(err, &apiError) {
		message += " " + string(apiError.Body())
	}
	lowerMessage := strings.ToLower(message)

	statusCode := 0
	if response != nil {
		statusCode = response.StatusCode
	}

	switch {
	case statusCode == http.StatusTooManyRequests || statusCode >= http.StatusInternalServerError:
		return fmt.Errorf("massed compute API request failed: %w", errors.Join(v1.ErrServiceUnavailable, err))
	case strings.Contains(lowerMessage, "capacity") || strings.Contains(lowerMessage, "out of stock"):
		return fmt.Errorf("massed compute API request failed: %w", errors.Join(v1.ErrInsufficientResources, err))
	case strings.Contains(lowerMessage, "quota") || strings.Contains(lowerMessage, "limit exceeded"):
		return fmt.Errorf("massed compute API request failed: %w", errors.Join(v1.ErrOutOfQuota, err))
	default:
		return fmt.Errorf("massed compute API request failed: %w", err)
	}
}
