# \CapacityAPI

All URIs are relative to *https://&lt;region&gt;.atlas.fluidstack.io/api/v1alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCapacity**](CapacityAPI.md#ListCapacity) | **Get** /capacity | List available capacity



## ListCapacity

> []ListCapacity200ResponseInner ListCapacity(ctx).XPROJECTID(xPROJECTID).Execute()

List available capacity

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
	resp, r, err := apiClient.CapacityAPI.ListCapacity(context.Background()).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.ListCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCapacity`: []ListCapacity200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.ListCapacity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

### Return type

[**[]ListCapacity200ResponseInner**](ListCapacity200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

