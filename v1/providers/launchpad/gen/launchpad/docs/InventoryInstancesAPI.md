# \InventoryInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryInstancesBulkPartialUpdate**](InventoryInstancesAPI.md#InventoryInstancesBulkPartialUpdate) | **Patch** /v1/inventory/instances/bulk/ | 🚧 [Beta Feature]
[**InventoryInstancesCreate**](InventoryInstancesAPI.md#InventoryInstancesCreate) | **Post** /v1/inventory/instances/ | 🚧 [Beta Feature]
[**InventoryInstancesDestroy**](InventoryInstancesAPI.md#InventoryInstancesDestroy) | **Delete** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**InventoryInstancesHistoryList**](InventoryInstancesAPI.md#InventoryInstancesHistoryList) | **Get** /v1/inventory/instances/{id}/history/ | 🚧 [Beta Feature]
[**InventoryInstancesList**](InventoryInstancesAPI.md#InventoryInstancesList) | **Get** /v1/inventory/instances/ | 🚧 [Beta Feature]
[**InventoryInstancesPartialUpdate**](InventoryInstancesAPI.md#InventoryInstancesPartialUpdate) | **Patch** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**InventoryInstancesRetrieve**](InventoryInstancesAPI.md#InventoryInstancesRetrieve) | **Get** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**InventoryInstancesUpdate**](InventoryInstancesAPI.md#InventoryInstancesUpdate) | **Put** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]



## InventoryInstancesBulkPartialUpdate

> InstanceBulkUpdate InventoryInstancesBulkPartialUpdate(ctx).InstanceBulkUpdate(instanceBulkUpdate).Execute()

🚧 [Beta Feature]

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
	instanceBulkUpdate := *openapiclient.NewInstanceBulkUpdate(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now(), int32(123), []string{"Ids_example"}, "Result_example") // InstanceBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesBulkPartialUpdate(context.Background()).InstanceBulkUpdate(instanceBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesBulkPartialUpdate`: InstanceBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceBulkUpdate** | [**InstanceBulkUpdate**](InstanceBulkUpdate.md) |  | 

### Return type

[**InstanceBulkUpdate**](InstanceBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInstancesCreate

> Instance InventoryInstancesCreate(ctx).Instance(instance).Execute()

🚧 [Beta Feature]

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
	instance := *openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesCreate(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesCreate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Instance**](Instance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInstancesDestroy

> InventoryInstancesDestroy(ctx, id).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryInstancesAPI.InventoryInstancesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesDestroyRequest struct via the builder pattern


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


## InventoryInstancesHistoryList

> PaginatedModelChangeList InventoryInstancesHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesHistoryListRequest struct via the builder pattern


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


## InventoryInstancesList

> PaginatedInstanceList InventoryInstancesList(ctx).Cluster(cluster).Expand(expand).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()

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
	cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	instanceId := "instanceId_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for instances by cluster, id, instance_id, name, state, tags (optional)
	state := "state_example" // string | Current lifecycle state of this instance  * `running` - Instance is running * `starting` - Instance is starting * `stopped` - Instance is stopped * `stopping` - Instance is stopping * `unknown` - Instance state is currently unknown (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesList(context.Background()).Cluster(cluster).Expand(expand).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesList`: PaginatedInstanceList
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster\&quot;). | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **instanceId** | **string** |  | 
 **name** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for instances by cluster, id, instance_id, name, state, tags | 
 **state** | **string** | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | 

### Return type

[**PaginatedInstanceList**](PaginatedInstanceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInstancesPartialUpdate

> Instance InventoryInstancesPartialUpdate(ctx, id).Instance(instance).Execute()

🚧 [Beta Feature]

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this instance.
	instance := *openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesPartialUpdate(context.Background(), id).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesPartialUpdate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**Instance**](Instance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInstancesRetrieve

> Instance InventoryInstancesRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this instance.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesRetrieve`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster\&quot;). | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Instance**](Instance.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInstancesUpdate

> Instance InventoryInstancesUpdate(ctx, id).Instance(instance).Execute()

🚧 [Beta Feature]

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this instance.
	instance := *openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.InventoryInstancesUpdate(context.Background(), id).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.InventoryInstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryInstancesUpdate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.InventoryInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInstancesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**Instance**](Instance.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

