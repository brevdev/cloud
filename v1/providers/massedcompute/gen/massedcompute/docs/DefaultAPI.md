# \DefaultAPI

All URIs are relative to *https://vm.massedcompute.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GpuInventoryGet**](DefaultAPI.md#GpuInventoryGet) | **Get** /gpu-inventory | Retrieve a list of avaialable GPU configurations.
[**ImagesGet**](DefaultAPI.md#ImagesGet) | **Get** /images | Retrieve list of available images.



## GpuInventoryGet

> GPUInventoryV1 GpuInventoryGet(ctx).Execute()

Retrieve a list of avaialable GPU configurations.



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
	resp, r, err := apiClient.DefaultAPI.GpuInventoryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GpuInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GpuInventoryGet`: GPUInventoryV1
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GpuInventoryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGpuInventoryGetRequest struct via the builder pattern


### Return type

[**GPUInventoryV1**](GPUInventoryV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesGet

> ImagesV1 ImagesGet(ctx).Execute()

Retrieve list of available images.



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
	resp, r, err := apiClient.DefaultAPI.ImagesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ImagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesGet`: ImagesV1
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ImagesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImagesGetRequest struct via the builder pattern


### Return type

[**ImagesV1**](ImagesV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

