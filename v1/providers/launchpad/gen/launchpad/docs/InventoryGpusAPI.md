# \InventoryGpusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryGpusBulkPartialUpdate**](InventoryGpusAPI.md#V1InventoryGpusBulkPartialUpdate) | **Patch** /v1/inventory/gpus/bulk/ | 
[**V1InventoryGpusCreate**](InventoryGpusAPI.md#V1InventoryGpusCreate) | **Post** /v1/inventory/gpus/ | 
[**V1InventoryGpusDestroy**](InventoryGpusAPI.md#V1InventoryGpusDestroy) | **Delete** /v1/inventory/gpus/{id}/ | 
[**V1InventoryGpusHistoryList**](InventoryGpusAPI.md#V1InventoryGpusHistoryList) | **Get** /v1/inventory/gpus/{id}/history/ | 
[**V1InventoryGpusList**](InventoryGpusAPI.md#V1InventoryGpusList) | **Get** /v1/inventory/gpus/ | 
[**V1InventoryGpusPartialUpdate**](InventoryGpusAPI.md#V1InventoryGpusPartialUpdate) | **Patch** /v1/inventory/gpus/{id}/ | 
[**V1InventoryGpusRetrieve**](InventoryGpusAPI.md#V1InventoryGpusRetrieve) | **Get** /v1/inventory/gpus/{id}/ | 
[**V1InventoryGpusStatsRetrieve**](InventoryGpusAPI.md#V1InventoryGpusStatsRetrieve) | **Get** /v1/inventory/gpus/stats/ | 🚧 [Beta Feature]
[**V1InventoryGpusUpdate**](InventoryGpusAPI.md#V1InventoryGpusUpdate) | **Put** /v1/inventory/gpus/{id}/ | 



## V1InventoryGpusBulkPartialUpdate

> GpuBulkUpdate V1InventoryGpusBulkPartialUpdate(ctx).GpuBulkUpdate(gpuBulkUpdate).Execute()



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
	gpuBulkUpdate := *openapiclient.NewGpuBulkUpdate(time.Now(), "Id_example", "Model_example", time.Now(), int32(123), int32(123), []string{"Ids_example"}, "Result_example") // GpuBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusBulkPartialUpdate(context.Background()).GpuBulkUpdate(gpuBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusBulkPartialUpdate`: GpuBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpuBulkUpdate** | [**GpuBulkUpdate**](GpuBulkUpdate.md) |  | 

### Return type

[**GpuBulkUpdate**](GpuBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryGpusCreate

> Gpu V1InventoryGpusCreate(ctx).Gpu(gpu).Execute()



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
	gpu := *openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now()) // Gpu | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusCreate(context.Background()).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusCreate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpu** | [**Gpu**](Gpu.md) |  | 

### Return type

[**Gpu**](Gpu.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryGpusDestroy

> V1InventoryGpusDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this gpu.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryGpusAPI.V1InventoryGpusDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusDestroyRequest struct via the builder pattern


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


## V1InventoryGpusHistoryList

> PaginatedModelChangeList V1InventoryGpusHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusHistoryListRequest struct via the builder pattern


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


## V1InventoryGpusList

> PaginatedGpuList V1InventoryGpusList(ctx).Fields(fields).FormFactor(formFactor).Id(id).Model(model).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()



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
	formFactor := "formFactor_example" // string | GPU form factor  * `pcie` - PCIe * `sxm` - SXM (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	model := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	priority := int32(56) // int32 |  (optional)
	search := "search_example" // string | Search for gpus by form_factor, id, memory, model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusList(context.Background()).Fields(fields).FormFactor(formFactor).Id(id).Model(model).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusList`: PaginatedGpuList
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Include only the specified fields in the response | 
 **formFactor** | **string** | GPU form factor  * &#x60;pcie&#x60; - PCIe * &#x60;sxm&#x60; - SXM | 
 **id** | **string** |  | 
 **model** | **[]string** | Multiple values may be separated by commas. | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **priority** | **int32** |  | 
 **search** | **string** | Search for gpus by form_factor, id, memory, model | 

### Return type

[**PaginatedGpuList**](PaginatedGpuList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryGpusPartialUpdate

> Gpu V1InventoryGpusPartialUpdate(ctx, id).Gpu(gpu).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this gpu.
	gpu := *openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now()) // Gpu | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusPartialUpdate(context.Background(), id).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusPartialUpdate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gpu** | [**Gpu**](Gpu.md) |  | 

### Return type

[**Gpu**](Gpu.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryGpusRetrieve

> Gpu V1InventoryGpusRetrieve(ctx, id).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this gpu.
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusRetrieve(context.Background(), id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusRetrieve`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Gpu**](Gpu.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryGpusStatsRetrieve

> V1InventoryGpusStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.InventoryGpusAPI.V1InventoryGpusStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusStatsRetrieveRequest struct via the builder pattern


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


## V1InventoryGpusUpdate

> Gpu V1InventoryGpusUpdate(ctx, id).Gpu(gpu).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this gpu.
	gpu := *openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now()) // Gpu | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGpusAPI.V1InventoryGpusUpdate(context.Background(), id).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.V1InventoryGpusUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryGpusUpdate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.V1InventoryGpusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryGpusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gpu** | [**Gpu**](Gpu.md) |  | 

### Return type

[**Gpu**](Gpu.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

