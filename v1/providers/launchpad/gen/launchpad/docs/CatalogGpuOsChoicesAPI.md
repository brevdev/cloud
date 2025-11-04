# \CatalogGpuOsChoicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogGpuOsChoicesCreate**](CatalogGpuOsChoicesAPI.md#V1CatalogGpuOsChoicesCreate) | **Post** /v1/catalog/gpu-os-choices/ | 
[**V1CatalogGpuOsChoicesList**](CatalogGpuOsChoicesAPI.md#V1CatalogGpuOsChoicesList) | **Get** /v1/catalog/gpu-os-choices/ | 



## V1CatalogGpuOsChoicesCreate

> GpuOsChoice V1CatalogGpuOsChoicesCreate(ctx).GpuOsChoice(gpuOsChoice).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	gpuOsChoice := *openapiclient.NewGpuOsChoice(time.Now(), int32(123), time.Now(), "Name_example", "Release_example", "Version_example") // GpuOsChoice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesCreate(context.Background()).GpuOsChoice(gpuOsChoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogGpuOsChoicesCreate`: GpuOsChoice
	fmt.Fprintf(os.Stdout, "Response from `CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogGpuOsChoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpuOsChoice** | [**GpuOsChoice**](GpuOsChoice.md) |  | 

### Return type

[**GpuOsChoice**](GpuOsChoice.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogGpuOsChoicesList

> PaginatedGpuOsChoiceList V1CatalogGpuOsChoicesList(ctx).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for gpuoschoices by id, name, release, version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesList(context.Background()).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogGpuOsChoicesList`: PaginatedGpuOsChoiceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogGpuOsChoicesAPI.V1CatalogGpuOsChoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogGpuOsChoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for gpuoschoices by id, name, release, version | 

### Return type

[**PaginatedGpuOsChoiceList**](PaginatedGpuOsChoiceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

