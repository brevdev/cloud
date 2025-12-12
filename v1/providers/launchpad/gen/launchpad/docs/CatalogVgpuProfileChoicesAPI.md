# \CatalogVgpuProfileChoicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogVgpuProfileChoicesCreate**](CatalogVgpuProfileChoicesAPI.md#V1CatalogVgpuProfileChoicesCreate) | **Post** /v1/catalog/vgpu-profile-choices/ | 
[**V1CatalogVgpuProfileChoicesList**](CatalogVgpuProfileChoicesAPI.md#V1CatalogVgpuProfileChoicesList) | **Get** /v1/catalog/vgpu-profile-choices/ | 



## V1CatalogVgpuProfileChoicesCreate

> VGpuProfileChoice V1CatalogVgpuProfileChoicesCreate(ctx).VGpuProfileChoice(vGpuProfileChoice).Execute()



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
	vGpuProfileChoice := *openapiclient.NewVGpuProfileChoice(time.Now(), int32(123), time.Now(), "Name_example") // VGpuProfileChoice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesCreate(context.Background()).VGpuProfileChoice(vGpuProfileChoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogVgpuProfileChoicesCreate`: VGpuProfileChoice
	fmt.Fprintf(os.Stdout, "Response from `CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogVgpuProfileChoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vGpuProfileChoice** | [**VGpuProfileChoice**](VGpuProfileChoice.md) |  | 

### Return type

[**VGpuProfileChoice**](VGpuProfileChoice.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogVgpuProfileChoicesList

> PaginatedVGpuProfileChoiceList V1CatalogVgpuProfileChoicesList(ctx).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	search := "search_example" // string | Search for vgpuprofilechoices by id, name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesList(context.Background()).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogVgpuProfileChoicesList`: PaginatedVGpuProfileChoiceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogVgpuProfileChoicesAPI.V1CatalogVgpuProfileChoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogVgpuProfileChoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for vgpuprofilechoices by id, name | 

### Return type

[**PaginatedVGpuProfileChoiceList**](PaginatedVGpuProfileChoiceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

