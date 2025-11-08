# \InventoryLocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryLocationsCreate**](InventoryLocationsAPI.md#V1InventoryLocationsCreate) | **Post** /v1/inventory/locations/ | 
[**V1InventoryLocationsDestroy**](InventoryLocationsAPI.md#V1InventoryLocationsDestroy) | **Delete** /v1/inventory/locations/{id}/ | 
[**V1InventoryLocationsHistoryList**](InventoryLocationsAPI.md#V1InventoryLocationsHistoryList) | **Get** /v1/inventory/locations/{id}/history/ | 
[**V1InventoryLocationsList**](InventoryLocationsAPI.md#V1InventoryLocationsList) | **Get** /v1/inventory/locations/ | 
[**V1InventoryLocationsPartialUpdate**](InventoryLocationsAPI.md#V1InventoryLocationsPartialUpdate) | **Patch** /v1/inventory/locations/{id}/ | 
[**V1InventoryLocationsRetrieve**](InventoryLocationsAPI.md#V1InventoryLocationsRetrieve) | **Get** /v1/inventory/locations/{id}/ | 
[**V1InventoryLocationsUpdate**](InventoryLocationsAPI.md#V1InventoryLocationsUpdate) | **Put** /v1/inventory/locations/{id}/ | 



## V1InventoryLocationsCreate

> Location V1InventoryLocationsCreate(ctx).Location(location).Execute()



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
	location := *openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")}) // Location | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsCreate(context.Background()).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsCreate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | [**Location**](Location.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsDestroy

> V1InventoryLocationsDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this location.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsHistoryList

> PaginatedModelChangeList V1InventoryLocationsHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsHistoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedModelChangeList**](PaginatedModelChangeList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsList

> PaginatedLocationList V1InventoryLocationsList(ctx).Expand(expand).Fields(fields).Id(id).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Region(region).Search(search).Execute()



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
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"provider\"). (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	provider := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	region := "region_example" // string |  (optional)
	search := "search_example" // string | Search for locations by id, name, provider name, region (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsList(context.Background()).Expand(expand).Fields(fields).Id(id).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Region(region).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsList`: PaginatedLocationList
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;provider\&quot;). | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **name** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **string** |  | 
 **region** | **string** |  | 
 **search** | **string** | Search for locations by id, name, provider name, region | 

### Return type

[**PaginatedLocationList**](PaginatedLocationList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsPartialUpdate

> Location V1InventoryLocationsPartialUpdate(ctx, id).Location(location).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this location.
	location := *openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")}) // Location | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsPartialUpdate(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsPartialUpdate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | [**Location**](Location.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsRetrieve

> Location V1InventoryLocationsRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this location.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"provider\"). (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsRetrieve`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;provider\&quot;). | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Location**](Location.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryLocationsUpdate

> Location V1InventoryLocationsUpdate(ctx, id).Location(location).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this location.
	location := *openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")}) // Location | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLocationsAPI.V1InventoryLocationsUpdate(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.V1InventoryLocationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryLocationsUpdate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.V1InventoryLocationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryLocationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **location** | [**Location**](Location.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

