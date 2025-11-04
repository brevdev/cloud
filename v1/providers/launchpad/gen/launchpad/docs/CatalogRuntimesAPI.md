# \CatalogRuntimesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogRuntimesCreate**](CatalogRuntimesAPI.md#V1CatalogRuntimesCreate) | **Post** /v1/catalog/runtimes/ | 
[**V1CatalogRuntimesDestroy**](CatalogRuntimesAPI.md#V1CatalogRuntimesDestroy) | **Delete** /v1/catalog/runtimes/{id}/ | 
[**V1CatalogRuntimesHistoryList**](CatalogRuntimesAPI.md#V1CatalogRuntimesHistoryList) | **Get** /v1/catalog/runtimes/{id}/history/ | 
[**V1CatalogRuntimesList**](CatalogRuntimesAPI.md#V1CatalogRuntimesList) | **Get** /v1/catalog/runtimes/ | 
[**V1CatalogRuntimesPartialUpdate**](CatalogRuntimesAPI.md#V1CatalogRuntimesPartialUpdate) | **Patch** /v1/catalog/runtimes/{id}/ | 
[**V1CatalogRuntimesRetrieve**](CatalogRuntimesAPI.md#V1CatalogRuntimesRetrieve) | **Get** /v1/catalog/runtimes/{id}/ | 
[**V1CatalogRuntimesUpdate**](CatalogRuntimesAPI.md#V1CatalogRuntimesUpdate) | **Put** /v1/catalog/runtimes/{id}/ | 



## V1CatalogRuntimesCreate

> Runtime V1CatalogRuntimesCreate(ctx).Runtime(runtime).Execute()



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
	runtime := *openapiclient.NewRuntime(time.Now(), "Id_example", time.Now(), "Name_example", "Url_example") // Runtime | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesCreate(context.Background()).Runtime(runtime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesCreate`: Runtime
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runtime** | [**Runtime**](Runtime.md) |  | 

### Return type

[**Runtime**](Runtime.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogRuntimesDestroy

> V1CatalogRuntimesDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this runtime.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this runtime. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesDestroyRequest struct via the builder pattern


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


## V1CatalogRuntimesHistoryList

> PaginatedModelChangeList V1CatalogRuntimesHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesHistoryListRequest struct via the builder pattern


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


## V1CatalogRuntimesList

> PaginatedRuntimeList V1CatalogRuntimesList(ctx).Branch(branch).CnsAddonPack(cnsAddonPack).CnsDocker(cnsDocker).CnsDriverVersion(cnsDriverVersion).CnsK8s(cnsK8s).CnsNvidiaDriver(cnsNvidiaDriver).CnsVersion(cnsVersion).Fields(fields).Id(id).Mig(mig).MigProfile(migProfile).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	branch := "branch_example" // string |  (optional)
	cnsAddonPack := true // bool |  (optional)
	cnsDocker := true // bool |  (optional)
	cnsDriverVersion := "cnsDriverVersion_example" // string |  (optional)
	cnsK8s := true // bool |  (optional)
	cnsNvidiaDriver := true // bool |  (optional)
	cnsVersion := "cnsVersion_example" // string |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	mig := true // bool |  (optional)
	migProfile := "migProfile_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for runtimes by branch, cns_driver_version, cns_version, mig_profile, id, name, url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesList(context.Background()).Branch(branch).CnsAddonPack(cnsAddonPack).CnsDocker(cnsDocker).CnsDriverVersion(cnsDriverVersion).CnsK8s(cnsK8s).CnsNvidiaDriver(cnsNvidiaDriver).CnsVersion(cnsVersion).Fields(fields).Id(id).Mig(mig).MigProfile(migProfile).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesList`: PaginatedRuntimeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **branch** | **string** |  | 
 **cnsAddonPack** | **bool** |  | 
 **cnsDocker** | **bool** |  | 
 **cnsDriverVersion** | **string** |  | 
 **cnsK8s** | **bool** |  | 
 **cnsNvidiaDriver** | **bool** |  | 
 **cnsVersion** | **string** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **mig** | **bool** |  | 
 **migProfile** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for runtimes by branch, cns_driver_version, cns_version, mig_profile, id, name, url | 

### Return type

[**PaginatedRuntimeList**](PaginatedRuntimeList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogRuntimesPartialUpdate

> Runtime V1CatalogRuntimesPartialUpdate(ctx, id).Runtime(runtime).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this runtime.
	runtime := *openapiclient.NewRuntime(time.Now(), "Id_example", time.Now(), "Name_example", "Url_example") // Runtime | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesPartialUpdate(context.Background(), id).Runtime(runtime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesPartialUpdate`: Runtime
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this runtime. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runtime** | [**Runtime**](Runtime.md) |  | 

### Return type

[**Runtime**](Runtime.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogRuntimesRetrieve

> Runtime V1CatalogRuntimesRetrieve(ctx, id).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this runtime.
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesRetrieve(context.Background(), id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesRetrieve`: Runtime
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this runtime. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Runtime**](Runtime.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogRuntimesUpdate

> Runtime V1CatalogRuntimesUpdate(ctx, id).Runtime(runtime).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this runtime.
	runtime := *openapiclient.NewRuntime(time.Now(), "Id_example", time.Now(), "Name_example", "Url_example") // Runtime | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogRuntimesAPI.V1CatalogRuntimesUpdate(context.Background(), id).Runtime(runtime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogRuntimesAPI.V1CatalogRuntimesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogRuntimesUpdate`: Runtime
	fmt.Fprintf(os.Stdout, "Response from `CatalogRuntimesAPI.V1CatalogRuntimesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this runtime. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogRuntimesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runtime** | [**Runtime**](Runtime.md) |  | 

### Return type

[**Runtime**](Runtime.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

