# \InstancesAPI

All URIs are relative to *https://vm.massedcompute.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceGet**](InstancesAPI.md#InstanceGet) | **Get** /instance | Retrieve list of all running instances.
[**InstanceLaunchPost**](InstancesAPI.md#InstanceLaunchPost) | **Post** /instance/launch | Deploy new instances.
[**InstanceRestartPost**](InstancesAPI.md#InstanceRestartPost) | **Post** /instance/restart | Restart an instances.
[**InstanceTerminatePost**](InstancesAPI.md#InstanceTerminatePost) | **Post** /instance/terminate | Terminate an instances.
[**InstanceUuidGet**](InstancesAPI.md#InstanceUuidGet) | **Get** /instance/{uuid} | Retrieve single running instances.



## InstanceGet

> RetrieveAllRunningInstancesV1 InstanceGet(ctx).Execute()

Retrieve list of all running instances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.InstanceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.InstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceGet`: RetrieveAllRunningInstancesV1
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.InstanceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceGetRequest struct via the builder pattern


### Return type

[**RetrieveAllRunningInstancesV1**](RetrieveAllRunningInstancesV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceLaunchPost

> InstanceLaunchPost202Response InstanceLaunchPost(ctx).InstanceLaunchPostRequest(instanceLaunchPostRequest).Execute()

Deploy new instances.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	instanceLaunchPostRequest := *openapiclient.NewInstanceLaunchPostRequest(int32(123), "ProductName_example", "RegionName_example") // InstanceLaunchPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.InstanceLaunchPost(context.Background()).InstanceLaunchPostRequest(instanceLaunchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.InstanceLaunchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceLaunchPost`: InstanceLaunchPost202Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.InstanceLaunchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceLaunchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceLaunchPostRequest** | [**InstanceLaunchPostRequest**](InstanceLaunchPostRequest.md) |  |

### Return type

[**InstanceLaunchPost202Response**](InstanceLaunchPost202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceRestartPost

> RestartInstanceV1 InstanceRestartPost(ctx).InstanceRestartPostRequest(instanceRestartPostRequest).Execute()

Restart an instances.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	instanceRestartPostRequest := *openapiclient.NewInstanceRestartPostRequest([]string{"InstanceUuids_example"}) // InstanceRestartPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.InstanceRestartPost(context.Background()).InstanceRestartPostRequest(instanceRestartPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.InstanceRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceRestartPost`: RestartInstanceV1
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.InstanceRestartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceRestartPostRequest** | [**InstanceRestartPostRequest**](InstanceRestartPostRequest.md) |  |

### Return type

[**RestartInstanceV1**](RestartInstanceV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceTerminatePost

> TerminateInstanceV1 InstanceTerminatePost(ctx).InstanceRestartPostRequest(instanceRestartPostRequest).Execute()

Terminate an instances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	instanceRestartPostRequest := *openapiclient.NewInstanceRestartPostRequest([]string{"InstanceUuids_example"}) // InstanceRestartPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.InstanceTerminatePost(context.Background()).InstanceRestartPostRequest(instanceRestartPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.InstanceTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceTerminatePost`: TerminateInstanceV1
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.InstanceTerminatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceTerminatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceRestartPostRequest** | [**InstanceRestartPostRequest**](InstanceRestartPostRequest.md) |  |

### Return type

[**TerminateInstanceV1**](TerminateInstanceV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceUuidGet

> RetrieveAllRunningInstancesV1 InstanceUuidGet(ctx, uuid).Execute()

Retrieve single running instances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	uuid := "uuid_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.InstanceUuidGet(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.InstanceUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceUuidGet`: RetrieveAllRunningInstancesV1
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.InstanceUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveAllRunningInstancesV1**](RetrieveAllRunningInstancesV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

