# \InventoryGpusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryGpusBulkPartialUpdate**](InventoryGpusAPI.md#InventoryGpusBulkPartialUpdate) | **Patch** /v1/inventory/gpus/bulk/ | 
[**InventoryGpusCreate**](InventoryGpusAPI.md#InventoryGpusCreate) | **Post** /v1/inventory/gpus/ | 
[**InventoryGpusDestroy**](InventoryGpusAPI.md#InventoryGpusDestroy) | **Delete** /v1/inventory/gpus/{id}/ | 
[**InventoryGpusHistoryList**](InventoryGpusAPI.md#InventoryGpusHistoryList) | **Get** /v1/inventory/gpus/{id}/history/ | 
[**InventoryGpusList**](InventoryGpusAPI.md#InventoryGpusList) | **Get** /v1/inventory/gpus/ | 
[**InventoryGpusPartialUpdate**](InventoryGpusAPI.md#InventoryGpusPartialUpdate) | **Patch** /v1/inventory/gpus/{id}/ | 
[**InventoryGpusRetrieve**](InventoryGpusAPI.md#InventoryGpusRetrieve) | **Get** /v1/inventory/gpus/{id}/ | 
[**InventoryGpusStatsRetrieve**](InventoryGpusAPI.md#InventoryGpusStatsRetrieve) | **Get** /v1/inventory/gpus/stats/ | 🚧 [Beta Feature]
[**InventoryGpusUpdate**](InventoryGpusAPI.md#InventoryGpusUpdate) | **Put** /v1/inventory/gpus/{id}/ | 



## InventoryGpusBulkPartialUpdate

> GpuBulkUpdate InventoryGpusBulkPartialUpdate(ctx).GpuBulkUpdate(gpuBulkUpdate).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusBulkPartialUpdate(context.Background()).GpuBulkUpdate(gpuBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusBulkPartialUpdate`: GpuBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusBulkPartialUpdateRequest struct via the builder pattern


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


## InventoryGpusCreate

> Gpu InventoryGpusCreate(ctx).Gpu(gpu).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusCreate(context.Background()).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusCreate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusCreateRequest struct via the builder pattern


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


## InventoryGpusDestroy

> InventoryGpusDestroy(ctx, id).Execute()



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
	r, err := apiClient.InventoryGpusAPI.InventoryGpusDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventoryGpusDestroyRequest struct via the builder pattern


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


## InventoryGpusHistoryList

> PaginatedModelChangeList InventoryGpusHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusHistoryListRequest struct via the builder pattern


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


## InventoryGpusList

> PaginatedGpuList InventoryGpusList(ctx).Fields(fields).FormFactor(formFactor).Id(id).Model(model).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusList(context.Background()).Fields(fields).FormFactor(formFactor).Id(id).Model(model).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Priority(priority).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusList`: PaginatedGpuList
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusListRequest struct via the builder pattern


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


## InventoryGpusPartialUpdate

> Gpu InventoryGpusPartialUpdate(ctx, id).Gpu(gpu).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusPartialUpdate(context.Background(), id).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusPartialUpdate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusPartialUpdateRequest struct via the builder pattern


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


## InventoryGpusRetrieve

> Gpu InventoryGpusRetrieve(ctx, id).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusRetrieve(context.Background(), id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusRetrieve`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusRetrieveRequest struct via the builder pattern


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


## InventoryGpusStatsRetrieve

> InventoryGpusStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.InventoryGpusAPI.InventoryGpusStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusStatsRetrieveRequest struct via the builder pattern


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


## InventoryGpusUpdate

> Gpu InventoryGpusUpdate(ctx, id).Gpu(gpu).Execute()



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
	resp, r, err := apiClient.InventoryGpusAPI.InventoryGpusUpdate(context.Background(), id).Gpu(gpu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGpusAPI.InventoryGpusUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryGpusUpdate`: Gpu
	fmt.Fprintf(os.Stdout, "Response from `InventoryGpusAPI.InventoryGpusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this gpu. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGpusUpdateRequest struct via the builder pattern


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

