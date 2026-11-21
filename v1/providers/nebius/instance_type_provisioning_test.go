package v1

import (
	"context"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/conn"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	"github.com/nebius/gosdk/serviceerror"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Intercept SDK calls before transport so these tests cannot provision real resources.
func instanceTypeTestClient(t *testing.T, call func(string, any, any) error) *NebiusClient {
	t.Helper()
	sdk, err := gosdk.New(context.Background(),
		gosdk.WithCredentials(gosdk.NoCredentials()),
		gosdk.WithoutParentID(),
		gosdk.WithResolvers(conn.NewConstantResolver("unused.invalid:443")),
		gosdk.WithDialOptions(
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(func(_ context.Context, method string, req, reply any, _ *grpc.ClientConn, _ grpc.UnaryInvoker, _ ...grpc.CallOption) error {
				return call(method, req, reply)
			}),
		),
	)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, sdk.Close()) })
	return &NebiusClient{sdk: sdk, projectID: "project-test", logger: &v1.NoopLogger{}}
}

func testPlatform(name, preset string) *compute.Platform {
	return &compute.Platform{
		Metadata: &common.ResourceMetadata{Name: name},
		Spec:     &compute.PlatformSpec{Presets: []*compute.Preset{{Name: preset}}},
	}
}

func TestCreateInstanceRejectsUnavailableTypeBeforeCreatingResources(t *testing.T) {
	t.Parallel()
	var calls []string
	client := instanceTypeTestClient(t, func(method string, _, reply any) error {
		calls = append(calls, method)
		require.Equal(t, compute.PlatformService_List_FullMethodName, method)
		response, ok := reply.(*compute.ListPlatformsResponse)
		require.True(t, ok)
		response.Items = []*compute.Platform{testPlatform("gpu-h100-sxm", "8gpu-128vcpu-1600gb")}
		return nil
	})
	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		RefID: "cpu-request", InstanceType: "cpu-e2.2vcpu-8gb",
	})
	require.ErrorIs(t, err, v1.ErrInsufficientResources)
	require.ErrorContains(t, err, "cpu-e2.2vcpu-8gb")
	require.Nil(t, instance)
	require.Equal(t, []string{compute.PlatformService_List_FullMethodName}, calls)
}

func TestCreateInstanceReturnsCapacityErrorWithoutSubstitution(t *testing.T) {
	t.Parallel()
	for _, asynchronous := range []bool{false, true} {
		name := "create RPC reports exhausted capacity"
		if asynchronous {
			name = "operation reports exhausted capacity"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			capacityStatus, err := status.New(codes.ResourceExhausted, "requested CPU capacity is exhausted").WithDetails(&common.ServiceError{
				Service: "compute", Code: "NotEnoughResources",
				Details: &common.ServiceError_NotEnoughResources{NotEnoughResources: &common.NotEnoughResources{}},
			})
			require.NoError(t, err)
			var createCalls int
			client := instanceTypeTestClient(t, func(method string, req, reply any) error {
				switch method {
				case compute.PlatformService_List_FullMethodName:
					response, ok := reply.(*compute.ListPlatformsResponse)
					require.True(t, ok)
					response.Items = []*compute.Platform{
						testPlatform("gpu-h100-sxm", "8gpu-128vcpu-1600gb"),
						testPlatform("cpu-e2", "2vcpu-8gb"),
					}
				case compute.InstanceService_Create_FullMethodName:
					createCalls++
					request, ok := req.(*compute.CreateInstanceRequest)
					require.True(t, ok)
					require.Equal(t, "cpu-e2", request.Spec.Resources.Platform)
					require.Equal(t, "2vcpu-8gb", request.Spec.Resources.GetPreset())
					if asynchronous {
						response, ok := reply.(proto.Message)
						require.True(t, ok)
						proto.Merge(response, &common.Operation{
							Id: "operation-failed", CreatedAt: timestamppb.Now(), Status: capacityStatus.Proto(),
						})
						return nil
					}
					// Match the SDK's service-error interceptor, bypassed by this test transport.
					return serviceerror.FromError(capacityStatus.Err(), capacityStatus.Proto().Details...)
				default:
					// Network, subnet and disk creation and their cleanup complete immediately.
					// Any other API call fails the test instead of reaching a provider.
					resourceID, ok := map[string]string{
						"/nebius.vpc.v1.NetworkService/Create":  "network-test",
						"/nebius.vpc.v1.SubnetService/Create":   "subnet-test",
						"/nebius.compute.v1.DiskService/Create": "disk-test",
						"/nebius.vpc.v1.NetworkService/Delete":  "network-test",
						"/nebius.vpc.v1.SubnetService/Delete":   "subnet-test",
						"/nebius.compute.v1.DiskService/Delete": "disk-test",
					}[method]
					require.True(t, ok, "unexpected SDK call: %s", method)
					response, ok := reply.(proto.Message)
					require.True(t, ok)
					proto.Merge(response, &common.Operation{
						Id: "operation-test", ResourceId: resourceID,
						CreatedAt: timestamppb.Now(), Status: status.New(codes.OK, "").Proto(),
					})
				}
				return nil
			})
			instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
				RefID: "cpu-request", InstanceType: "cpu-e2.2vcpu-8gb",
			})
			require.ErrorIs(t, err, v1.ErrInsufficientResources)
			require.Nil(t, instance)
			require.Equal(t, 1, createCalls, "capacity failure must not try another instance type")
		})
	}
}
