package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/lambdalabs/gen/lambdalabs"
	"github.com/cenkalti/backoff/v4"
)

func handleAPIError(_ context.Context, resp *http.Response, err error) error {
	body := ""
	e, ok := err.(openapi.GenericOpenAPIError)
	if ok {
		body = string(e.Body())
	}
	if body == "" {
		bodyBytes, errr := io.ReadAll(resp.Body)
		if errr != nil {
			fmt.Printf("Error reading response body: %v\n", errr)
		}
		body = string(bodyBytes)
	}
	outErr := fmt.Errorf("LambdaLabs API error\n%s\n%s:\nErr: %s\n%s", resp.Request.URL, resp.Status, err.Error(), body)
	if strings.Contains(body, "instance does not exist") { //nolint:gocritic // ignore
		return backoff.Permanent(v1.ErrInstanceNotFound)
	} else if strings.Contains(body, "banned you temporarily") {
		return outErr
	} else if resp.StatusCode < 500 && resp.StatusCode != 429 { // 429 Too Many Requests (use back off)
		return backoff.Permanent(outErr)
	} else {
		return outErr
	}
}

func handleErrToCloudErr(e error) error {
	if e == nil {
		return nil
	}
	if strings.Contains(e.Error(), "Not enough capacity") || strings.Contains(e.Error(), "insufficient-capacity") { //nolint:gocritic // ignore
		return v1.ErrInsufficientResources
	} else if strings.Contains(e.Error(), "global/invalid-parameters") && strings.Contains(e.Error(), "Region") && strings.Contains(e.Error(), "does not exist") {
		return v1.ErrInsufficientResources
	} else {
		return e
	}
}
