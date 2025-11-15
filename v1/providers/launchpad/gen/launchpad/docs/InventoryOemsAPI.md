# \InventoryOemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryOemsCreate**](InventoryOemsAPI.md#V1InventoryOemsCreate) | **Post** /v1/inventory/oems/ | 
[**V1InventoryOemsDestroy**](InventoryOemsAPI.md#V1InventoryOemsDestroy) | **Delete** /v1/inventory/oems/{id}/ | 
[**V1InventoryOemsHistoryList**](InventoryOemsAPI.md#V1InventoryOemsHistoryList) | **Get** /v1/inventory/oems/{id}/history/ | 
[**V1InventoryOemsList**](InventoryOemsAPI.md#V1InventoryOemsList) | **Get** /v1/inventory/oems/ | 
[**V1InventoryOemsPartialUpdate**](InventoryOemsAPI.md#V1InventoryOemsPartialUpdate) | **Patch** /v1/inventory/oems/{id}/ | 
[**V1InventoryOemsRetrieve**](InventoryOemsAPI.md#V1InventoryOemsRetrieve) | **Get** /v1/inventory/oems/{id}/ | 
[**V1InventoryOemsUpdate**](InventoryOemsAPI.md#V1InventoryOemsUpdate) | **Put** /v1/inventory/oems/{id}/ | 



## V1InventoryOemsCreate

> OEM V1InventoryOemsCreate(ctx).OEM(oEM).Execute()



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
	oEM := *openapiclient.NewOEM(time.Now(), "Id_example", time.Now(), "Name_example") // OEM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsCreate(context.Background()).OEM(oEM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsCreate`: OEM
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oEM** | [**OEM**](OEM.md) |  | 

### Return type

[**OEM**](OEM.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryOemsDestroy

> V1InventoryOemsDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this oem.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryOemsAPI.V1InventoryOemsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this oem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsDestroyRequest struct via the builder pattern


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


## V1InventoryOemsHistoryList

> PaginatedModelChangeList V1InventoryOemsHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsHistoryListRequest struct via the builder pattern


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


## V1InventoryOemsList

> PaginatedOEMList V1InventoryOemsList(ctx).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for oems by id, name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsList(context.Background()).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsList`: PaginatedOEMList
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for oems by id, name | 

### Return type

[**PaginatedOEMList**](PaginatedOEMList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryOemsPartialUpdate

> OEM V1InventoryOemsPartialUpdate(ctx, id).OEM(oEM).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this oem.
	oEM := *openapiclient.NewOEM(time.Now(), "Id_example", time.Now(), "Name_example") // OEM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsPartialUpdate(context.Background(), id).OEM(oEM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsPartialUpdate`: OEM
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this oem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oEM** | [**OEM**](OEM.md) |  | 

### Return type

[**OEM**](OEM.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryOemsRetrieve

> OEM V1InventoryOemsRetrieve(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this oem.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsRetrieve`: OEM
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this oem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OEM**](OEM.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryOemsUpdate

> OEM V1InventoryOemsUpdate(ctx, id).OEM(oEM).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this oem.
	oEM := *openapiclient.NewOEM(time.Now(), "Id_example", time.Now(), "Name_example") // OEM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryOemsAPI.V1InventoryOemsUpdate(context.Background(), id).OEM(oEM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryOemsAPI.V1InventoryOemsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryOemsUpdate`: OEM
	fmt.Fprintf(os.Stdout, "Response from `InventoryOemsAPI.V1InventoryOemsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this oem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryOemsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oEM** | [**OEM**](OEM.md) |  | 

### Return type

[**OEM**](OEM.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

