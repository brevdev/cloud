# \InstancesAPI

All URIs are relative to *https://&lt;region&gt;.atlas.fluidstack.io/api/v1alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstancesAPI.md#CreateInstance) | **Post** /instances | Create a machine instance
[**DeleteInstance**](InstancesAPI.md#DeleteInstance) | **Delete** /instances/{id} | Delete a machine instance
[**GetInstance**](InstancesAPI.md#GetInstance) | **Get** /instances/{id} | Get a machine instance
[**ListInstances**](InstancesAPI.md#ListInstances) | **Get** /instances | List all machine instances
[**StartInstance**](InstancesAPI.md#StartInstance) | **Post** /instances/{id}/actions/start | Start a machine instance
[**StopInstance**](InstancesAPI.md#StopInstance) | **Post** /instances/{id}/actions/stop | Stop a machine instance



## CreateInstance

> Instance CreateInstance(ctx).XPROJECTID(xPROJECTID).InstancesPostRequest(instancesPostRequest).Execute()

Create a machine instance

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	instancesPostRequest := *openapiclient.NewInstancesPostRequest("instance-1", "h100-hgx-80gb.8x") // InstancesPostRequest | Details of the instance to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.CreateInstance(context.Background()).XPROJECTID(xPROJECTID).InstancesPostRequest(instancesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 
 **instancesPostRequest** | [**InstancesPostRequest**](InstancesPostRequest.md) | Details of the instance to create | 

### Return type

[**Instance**](Instance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, id).XPROJECTID(xPROJECTID).Execute()

Delete a machine instance

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstancesAPI.DeleteInstance(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> Instance GetInstance(ctx, id).XPROJECTID(xPROJECTID).Execute()

Get a machine instance

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstance(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**Instance**](Instance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> []Instance ListInstances(ctx).XPROJECTID(xPROJECTID).Execute()

List all machine instances

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.ListInstances(context.Background()).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: []Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

### Return type

[**[]Instance**](Instance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartInstance

> StartInstance(ctx, id).XPROJECTID(xPROJECTID).Execute()

Start a machine instance

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstancesAPI.StartInstance(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopInstance

> StopInstance(ctx, id).XPROJECTID(xPROJECTID).Execute()

Stop a machine instance

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstancesAPI.StopInstance(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

