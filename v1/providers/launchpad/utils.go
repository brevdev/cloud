package v1

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
)

func (c *LaunchpadClient) handleLaunchpadAPIErr(ctx context.Context, resp *http.Response, err error) error {
	c.logger.Info(ctx, "Launchpad Error", v1.LogField("status", resp.Status))
	body := ""
	defer errors.HandleErrDefer(resp.Body.Close)
	if apiErr, ok := err.(openapi.GenericOpenAPIError); ok {
		body = string(apiErr.Body())
	}
	if body == "" {
		bodyBytes, errr := io.ReadAll(resp.Body)
		if errr != nil {
			c.logger.Error(ctx, errors.Wrap(errr, "Error reading response body"))
		}
		body = string(bodyBytes)
	}
	outErr := errors.Errorf("Launchpad API error\n%s\n%s:\nErr: %s\n%s", resp.Request.URL, resp.Status, err.Error(), body)
	if errors.ErrorContains(outErr, "no available capacity") { //nolint:gocritic // if else preferred
		return errors.WrapAndTrace(errors.Join(v1.ErrInsufficientResources, outErr))
	} else if errors.ErrorContains(outErr, "No Deployment matches the given query") {
		return errors.WrapAndTrace(errors.Join(v1.ErrInstanceNotFound, outErr))

		// Case where deployment has already received terminate request - can treat this is an idempotent successful terminate request
	} else if strings.Contains(body, "Deployment is already being destroyed") && resp.StatusCode == http.StatusBadRequest {
		return nil
	}
	return errors.WrapAndTrace(outErr)
}
