# \InventoryProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryProvidersBulkPartialUpdate**](InventoryProvidersAPI.md#V1InventoryProvidersBulkPartialUpdate) | **Patch** /v1/inventory/providers/bulk/ | 
[**V1InventoryProvidersCreate**](InventoryProvidersAPI.md#V1InventoryProvidersCreate) | **Post** /v1/inventory/providers/ | 
[**V1InventoryProvidersDestroy**](InventoryProvidersAPI.md#V1InventoryProvidersDestroy) | **Delete** /v1/inventory/providers/{id}/ | 
[**V1InventoryProvidersHistoryList**](InventoryProvidersAPI.md#V1InventoryProvidersHistoryList) | **Get** /v1/inventory/providers/{id}/history/ | 
[**V1InventoryProvidersList**](InventoryProvidersAPI.md#V1InventoryProvidersList) | **Get** /v1/inventory/providers/ | 
[**V1InventoryProvidersPartialUpdate**](InventoryProvidersAPI.md#V1InventoryProvidersPartialUpdate) | **Patch** /v1/inventory/providers/{id}/ | 
[**V1InventoryProvidersRetrieve**](InventoryProvidersAPI.md#V1InventoryProvidersRetrieve) | **Get** /v1/inventory/providers/{id}/ | 
[**V1InventoryProvidersStatsRetrieve**](InventoryProvidersAPI.md#V1InventoryProvidersStatsRetrieve) | **Get** /v1/inventory/providers/stats/ | 🚧 [Beta Feature]
[**V1InventoryProvidersUpdate**](InventoryProvidersAPI.md#V1InventoryProvidersUpdate) | **Put** /v1/inventory/providers/{id}/ | 



## V1InventoryProvidersBulkPartialUpdate

> ProviderBulkUpdate V1InventoryProvidersBulkPartialUpdate(ctx).ProviderBulkUpdate(providerBulkUpdate).Execute()



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
	providerBulkUpdate := *openapiclient.NewProviderBulkUpdate(time.Now(), "Id_example", time.Now(), "Name_example", int32(123), []string{"Ids_example"}, "Result_example") // ProviderBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersBulkPartialUpdate(context.Background()).ProviderBulkUpdate(providerBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersBulkPartialUpdate`: ProviderBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerBulkUpdate** | [**ProviderBulkUpdate**](ProviderBulkUpdate.md) |  | 

### Return type

[**ProviderBulkUpdate**](ProviderBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryProvidersCreate

> Provider V1InventoryProvidersCreate(ctx).Provider(provider).Execute()



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
	provider := *openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example") // Provider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersCreate(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersCreate`: Provider
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | [**Provider**](Provider.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryProvidersDestroy

> V1InventoryProvidersDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersDestroyRequest struct via the builder pattern


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


## V1InventoryProvidersHistoryList

> PaginatedModelChangeList V1InventoryProvidersHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersHistoryListRequest struct via the builder pattern


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


## V1InventoryProvidersList

> PaginatedProviderList V1InventoryProvidersList(ctx).DisplayName(displayName).Fields(fields).Id(id).InstanceLimit(instanceLimit).Name(name).Omit(omit).OnDemandSpeed(onDemandSpeed).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()



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
	displayName := "displayName_example" // string |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	instanceLimit := int32(56) // int32 |  (optional)
	name := "name_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	onDemandSpeed := "onDemandSpeed_example" // string | Speed of on-demand inventory provisioning  * `fast` - fast * `slow` - slow (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	priority := int32(56) // int32 |  (optional)
	search := "search_example" // string | Search for providers by display_name, id, name, on_demand_speed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersList(context.Background()).DisplayName(displayName).Fields(fields).Id(id).InstanceLimit(instanceLimit).Name(name).Omit(omit).OnDemandSpeed(onDemandSpeed).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersList`: PaginatedProviderList
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **string** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **instanceLimit** | **int32** |  | 
 **name** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **onDemandSpeed** | **string** | Speed of on-demand inventory provisioning  * &#x60;fast&#x60; - fast * &#x60;slow&#x60; - slow | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **priority** | **int32** |  | 
 **search** | **string** | Search for providers by display_name, id, name, on_demand_speed | 

### Return type

[**PaginatedProviderList**](PaginatedProviderList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryProvidersPartialUpdate

> Provider V1InventoryProvidersPartialUpdate(ctx, id).Provider(provider).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this provider.
	provider := *openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example") // Provider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersPartialUpdate(context.Background(), id).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersPartialUpdate`: Provider
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | [**Provider**](Provider.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryProvidersRetrieve

> Provider V1InventoryProvidersRetrieve(ctx, id).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this provider.
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersRetrieve(context.Background(), id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersRetrieve`: Provider
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Provider**](Provider.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryProvidersStatsRetrieve

> V1InventoryProvidersStatsRetrieve(ctx).Execute()

🚧 [Beta Feature]

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
	r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersStatsRetrieveRequest struct via the builder pattern


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


## V1InventoryProvidersUpdate

> Provider V1InventoryProvidersUpdate(ctx, id).Provider(provider).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this provider.
	provider := *openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example") // Provider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryProvidersAPI.V1InventoryProvidersUpdate(context.Background(), id).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryProvidersAPI.V1InventoryProvidersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryProvidersUpdate`: Provider
	fmt.Fprintf(os.Stdout, "Response from `InventoryProvidersAPI.V1InventoryProvidersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryProvidersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | [**Provider**](Provider.md) |  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

