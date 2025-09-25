package v1

import (
	"context"
	"net/http"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
)

func (c *LaunchpadClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	_, err := c.GetInstance(ctx, instanceID) // ensures instance exists and is from credential provider
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	_, resp, err := c.client.CatalogDeploymentsAPI.CatalogDeploymentsDestroy(c.makeAuthContext(ctx), string(instanceID)).
		Execute()
	if resp != nil {
		defer resp.Body.Close() //nolint:errcheck // handled in err check
	}

	// No error returned, so we can consider this a successful terminate request
	if err == nil {
		return nil
	}

	// If there was no response, but there was an error, return the error
	if resp == nil {
		return errors.WrapAndTrace(err)
	}
	// If the response is not 200, return the error
	if resp.StatusCode != http.StatusOK {
		return errors.WrapAndTrace(c.handleLaunchpadAPIErr(ctx, resp, err))
	}

	// We have an error, but the response is 200, so we can consider this a successful terminate request
	return nil
}
