package v1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/lambdalabs/gen/lambdalabs"
)

func TestHandleAPIError_InstanceNotFound(t *testing.T) {
	body := `{"error": {"message": "instance does not exist"}}`
	resp := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "404 Not Found",
	}

	err := handleAPIError(context.Background(), resp, errors.New("not found"))

	var permanentErr *backoff.PermanentError
	require.True(t, errors.As(err, &permanentErr))
	assert.Equal(t, v1.ErrInstanceNotFound, permanentErr.Err)
}

func TestHandleAPIError_BannedTemporarily(t *testing.T) {
	body := `{"error": {"message": "banned you temporarily"}}`
	resp := &http.Response{
		StatusCode: 429,
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "429 Too Many Requests",
	}

	err := handleAPIError(context.Background(), resp, errors.New("rate limited"))

	var permanentErr *backoff.PermanentError
	assert.False(t, errors.As(err, &permanentErr))
	assert.Contains(t, err.Error(), "LambdaLabs API error")
	assert.Contains(t, err.Error(), "banned you temporarily")
}

func TestHandleAPIError_ClientError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		status     string
	}{
		{"bad request", 400, "400 Bad Request"},
		{"unauthorized", 401, "401 Unauthorized"},
		{"forbidden", 403, "403 Forbidden"},
		{"not found", 404, "404 Not Found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := `{"error": {"message": "client error"}}`
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(strings.NewReader(body)),
				Request:    &http.Request{URL: &url.URL{Path: "/test"}},
				Status:     tt.status,
			}

			err := handleAPIError(context.Background(), resp, errors.New("client error"))

			var permanentErr *backoff.PermanentError
			require.True(t, errors.As(err, &permanentErr))
			assert.Contains(t, permanentErr.Err.Error(), "LambdaLabs API error")
		})
	}
}

func TestHandleAPIError_TooManyRequests(t *testing.T) {
	body := `{"error": {"message": "too many requests"}}`
	resp := &http.Response{
		StatusCode: 429,
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "429 Too Many Requests",
	}

	err := handleAPIError(context.Background(), resp, errors.New("rate limited"))

	var permanentErr *backoff.PermanentError
	assert.False(t, errors.As(err, &permanentErr))
	assert.Contains(t, err.Error(), "LambdaLabs API error")
}

func TestHandleAPIError_ServerError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		status     string
	}{
		{"internal server error", 500, "500 Internal Server Error"},
		{"bad gateway", 502, "502 Bad Gateway"},
		{"service unavailable", 503, "503 Service Unavailable"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := `{"error": {"message": "server error"}}`
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(strings.NewReader(body)),
				Request:    &http.Request{URL: &url.URL{Path: "/test"}},
				Status:     tt.status,
			}

			err := handleAPIError(context.Background(), resp, errors.New("server error"))

			var permanentErr *backoff.PermanentError
			assert.False(t, errors.As(err, &permanentErr))
			assert.Contains(t, err.Error(), "LambdaLabs API error")
		})
	}
}

func TestHandleAPIError_OpenAPIError(t *testing.T) {
	body := `{"error": {"message": "test error"}}`

	openAPIErr := openapi.GenericOpenAPIError{}

	resp := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "400 Bad Request",
	}

	err := handleAPIError(context.Background(), resp, openAPIErr)

	var permanentErr *backoff.PermanentError
	require.True(t, errors.As(err, &permanentErr))
	assert.Contains(t, permanentErr.Err.Error(), "LambdaLabs API error")
	assert.Contains(t, permanentErr.Err.Error(), "/test")
	assert.Contains(t, permanentErr.Err.Error(), "400 Bad Request")
}

func TestHandleAPIError_EmptyBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("")),
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "400 Bad Request",
	}

	err := handleAPIError(context.Background(), resp, errors.New("test error"))

	var permanentErr *backoff.PermanentError
	require.True(t, errors.As(err, &permanentErr))
	assert.Contains(t, permanentErr.Err.Error(), "LambdaLabs API error")
	assert.Contains(t, permanentErr.Err.Error(), "test error")
}

func TestHandleAPIError_BodyReadError(t *testing.T) {
	resp := &http.Response{
		StatusCode: 400,
		Body:       &errorReader{},
		Request:    &http.Request{URL: &url.URL{Path: "/test"}},
		Status:     "400 Bad Request",
	}

	err := handleAPIError(context.Background(), resp, errors.New("test error"))

	var permanentErr *backoff.PermanentError
	require.True(t, errors.As(err, &permanentErr))
	assert.Contains(t, permanentErr.Err.Error(), "LambdaLabs API error")
}

type errorReader struct{}

func (e *errorReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (e *errorReader) Close() error {
	return nil
}

func TestHandleErrToCloudErr_NilError(t *testing.T) {
	result := handleErrToCloudErr(nil)
	assert.Nil(t, result)
}

func TestHandleErrToCloudErr_CapacityErrors(t *testing.T) {
	tests := []struct {
		name     string
		errMsg   string
		expected error
	}{
		{
			name:     "not enough capacity",
			errMsg:   "Not enough capacity in region",
			expected: v1.ErrInsufficientResources,
		},
		{
			name:     "insufficient capacity",
			errMsg:   "insufficient-capacity error occurred",
			expected: v1.ErrInsufficientResources,
		},
		{
			name:     "capacity with mixed case",
			errMsg:   "Error: Not enough capacity available",
			expected: v1.ErrInsufficientResources,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputErr := errors.New(tt.errMsg)
			result := handleErrToCloudErr(inputErr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHandleErrToCloudErr_RegionErrors(t *testing.T) {
	tests := []struct {
		name     string
		errMsg   string
		expected error
	}{
		{
			name:     "region does not exist",
			errMsg:   "global/invalid-parameters: Region us-invalid-1 does not exist",
			expected: v1.ErrInsufficientResources,
		},
		{
			name:     "region error with different format",
			errMsg:   "global/invalid-parameters error: Region eu-central-99 does not exist in this zone",
			expected: v1.ErrInsufficientResources,
		},
		{
			name:     "invalid parameters without region",
			errMsg:   "global/invalid-parameters: Invalid instance type",
			expected: errors.New("global/invalid-parameters: Invalid instance type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputErr := errors.New(tt.errMsg)
			result := handleErrToCloudErr(inputErr)
			if tt.expected == v1.ErrInsufficientResources {
				assert.Equal(t, tt.expected, result)
			} else {
				assert.Equal(t, tt.expected.Error(), result.Error())
			}
		})
	}
}

func TestHandleErrToCloudErr_OtherErrors(t *testing.T) {
	tests := []struct {
		name   string
		errMsg string
	}{
		{
			name:   "authentication error",
			errMsg: "invalid API key provided",
		},
		{
			name:   "network error",
			errMsg: "connection timeout",
		},
		{
			name:   "generic error",
			errMsg: "something went wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputErr := errors.New(tt.errMsg)
			result := handleErrToCloudErr(inputErr)
			assert.Equal(t, inputErr, result)
		})
	}
}

func TestHandleErrToCloudErr_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		errMsg   string
		expected error
	}{
		{
			name:     "empty error message",
			errMsg:   "",
			expected: errors.New(""),
		},
		{
			name:     "capacity substring in larger message",
			errMsg:   "The request failed because Not enough capacity is available in the selected region",
			expected: v1.ErrInsufficientResources,
		},
		{
			name:     "insufficient capacity with prefix",
			errMsg:   "API Error: insufficient-capacity - please try again later",
			expected: v1.ErrInsufficientResources,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputErr := errors.New(tt.errMsg)
			result := handleErrToCloudErr(inputErr)
			if tt.expected == v1.ErrInsufficientResources {
				assert.Equal(t, tt.expected, result)
			} else {
				assert.Equal(t, tt.expected.Error(), result.Error())
			}
		})
	}
}
