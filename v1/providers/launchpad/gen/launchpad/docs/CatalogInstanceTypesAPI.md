# \CatalogInstanceTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogInstanceTypesList**](CatalogInstanceTypesAPI.md#V1CatalogInstanceTypesList) | **Get** /v1/catalog/instance-types/ | 



## V1CatalogInstanceTypesList

> PaginatedInstanceTypeList V1CatalogInstanceTypesList(ctx).Delivery(delivery).Page(page).PageSize(pageSize).Execute()



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
	delivery := "delivery_example" // string | Limit results to instance types that can meet the given delivery timelines. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogInstanceTypesAPI.V1CatalogInstanceTypesList(context.Background()).Delivery(delivery).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogInstanceTypesAPI.V1CatalogInstanceTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogInstanceTypesList`: PaginatedInstanceTypeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogInstanceTypesAPI.V1CatalogInstanceTypesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogInstanceTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delivery** | **string** | Limit results to instance types that can meet the given delivery timelines. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedInstanceTypeList**](PaginatedInstanceTypeList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

