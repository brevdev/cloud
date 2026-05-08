package v1

import (
	"errors"
	"testing"

	cloudv1 "github.com/brevdev/cloud/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	"github.com/nebius/gosdk/serviceerror"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHandleErrToCloudErrMapsNotEnoughResourcesToInsufficientResources(t *testing.T) {
	t.Parallel()

	err := &serviceerror.Error{
		Wrapped: status.Error(codes.ResourceExhausted, "operation failed"),
		Details: []serviceerror.Detail{
			serviceerror.NewDetail(&common.ServiceError{
				Service: "compute",
				Code:    "NotEnoughResources",
				Details: &common.ServiceError_NotEnoughResources{
					NotEnoughResources: &common.NotEnoughResources{
						Violations: []*common.NotEnoughResources_Violation{
							{
								ResourceType: "virtualMachine",
								Requested:    "1gpu-16vcpu-64gb",
								Message:      "VM schedule timeout, most likely due to insufficient hardware resources",
							},
						},
					},
				},
			}),
		},
	}

	require.True(t, errors.Is(handleErrToCloudErr(err), cloudv1.ErrInsufficientResources))
}

func TestHandleErrToCloudErrMapsQuotaFailureToOutOfQuota(t *testing.T) {
	t.Parallel()

	err := &serviceerror.Error{
		Wrapped: status.Error(codes.ResourceExhausted, "operation failed"),
		Details: []serviceerror.Detail{
			serviceerror.NewDetail(&common.ServiceError{
				Service: "compute",
				Code:    "QuotaFailure",
				Details: &common.ServiceError_QuotaFailure{
					QuotaFailure: &common.QuotaFailure{
						Violations: []*common.QuotaFailure_Violation{
							{
								Quota:     "compute.instance.gpu.h100",
								Limit:     "0",
								Requested: "1",
								Message:   "quota exceeded",
							},
						},
					},
				},
			}),
		},
	}

	require.True(t, errors.Is(handleErrToCloudErr(err), cloudv1.ErrOutOfQuota))
}
