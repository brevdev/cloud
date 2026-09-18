package hyperstack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

type errorResponse struct {
	Message     string `json:"message"`
	ErrorReason string `json:"error_reason"`
}

func wrapTransportError(operation string, err error) error {
	return fmt.Errorf("hyperstack %s failed: %w", operation, err)
}

func responseError(operation string, statusCode int, body []byte, notFound error) error {
	var apiError errorResponse
	_ = json.Unmarshal(body, &apiError)
	detail := strings.TrimSpace(strings.Join([]string{apiError.Message, apiError.ErrorReason}, ": "))
	detail = strings.Trim(detail, ": ")
	if detail == "" {
		detail = http.StatusText(statusCode)
	}

	requestError := fmt.Errorf("HTTP %d: %s", statusCode, detail)
	lowerDetail := strings.ToLower(detail)
	var sentinel error
	switch {
	case statusCode == http.StatusNotFound && notFound != nil:
		sentinel = notFound
	case statusCode == http.StatusTooManyRequests || statusCode >= http.StatusInternalServerError:
		sentinel = v1.ErrServiceUnavailable
	case strings.Contains(lowerDetail, "capacity") || strings.Contains(lowerDetail, "out of stock") || strings.Contains(lowerDetail, "stock unavailable"):
		sentinel = v1.ErrInsufficientResources
	case strings.Contains(lowerDetail, "quota") || strings.Contains(lowerDetail, "limit exceeded"):
		sentinel = v1.ErrOutOfQuota
	}
	if sentinel != nil {
		requestError = errors.Join(sentinel, requestError)
	}
	return fmt.Errorf("hyperstack %s failed: %w", operation, requestError)
}
