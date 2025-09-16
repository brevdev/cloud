# \InventoryLocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryLocationsCreate**](InventoryLocationsAPI.md#InventoryLocationsCreate) | **Post** /v1/inventory/locations/ | 
[**InventoryLocationsDestroy**](InventoryLocationsAPI.md#InventoryLocationsDestroy) | **Delete** /v1/inventory/locations/{id}/ | 
[**InventoryLocationsHistoryList**](InventoryLocationsAPI.md#InventoryLocationsHistoryList) | **Get** /v1/inventory/locations/{id}/history/ | 
[**InventoryLocationsList**](InventoryLocationsAPI.md#InventoryLocationsList) | **Get** /v1/inventory/locations/ | 
[**InventoryLocationsPartialUpdate**](InventoryLocationsAPI.md#InventoryLocationsPartialUpdate) | **Patch** /v1/inventory/locations/{id}/ | 
[**InventoryLocationsRetrieve**](InventoryLocationsAPI.md#InventoryLocationsRetrieve) | **Get** /v1/inventory/locations/{id}/ | 
[**InventoryLocationsUpdate**](InventoryLocationsAPI.md#InventoryLocationsUpdate) | **Put** /v1/inventory/locations/{id}/ | 



## InventoryLocationsCreate

> Location InventoryLocationsCreate(ctx).Location(location).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsCreate(context.Background()).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsCreate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsCreateRequest struct via the builder pattern


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


## InventoryLocationsDestroy

> InventoryLocationsDestroy(ctx, id).Execute()



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
	r, err := apiClient.InventoryLocationsAPI.InventoryLocationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventoryLocationsDestroyRequest struct via the builder pattern


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


## InventoryLocationsHistoryList

> PaginatedModelChangeList InventoryLocationsHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsHistoryListRequest struct via the builder pattern


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


## InventoryLocationsList

> PaginatedLocationList InventoryLocationsList(ctx).Expand(expand).Fields(fields).Id(id).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Region(region).Search(search).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsList(context.Background()).Expand(expand).Fields(fields).Id(id).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Region(region).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsList`: PaginatedLocationList
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsListRequest struct via the builder pattern


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


## InventoryLocationsPartialUpdate

> Location InventoryLocationsPartialUpdate(ctx, id).Location(location).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsPartialUpdate(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsPartialUpdate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsPartialUpdateRequest struct via the builder pattern


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


## InventoryLocationsRetrieve

> Location InventoryLocationsRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsRetrieve`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsRetrieveRequest struct via the builder pattern


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


## InventoryLocationsUpdate

> Location InventoryLocationsUpdate(ctx, id).Location(location).Execute()



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
	resp, r, err := apiClient.InventoryLocationsAPI.InventoryLocationsUpdate(context.Background(), id).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLocationsAPI.InventoryLocationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryLocationsUpdate`: Location
	fmt.Fprintf(os.Stdout, "Response from `InventoryLocationsAPI.InventoryLocationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryLocationsUpdateRequest struct via the builder pattern


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

