# \InventoryInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryInstancesBulkPartialUpdate**](InventoryInstancesAPI.md#V1InventoryInstancesBulkPartialUpdate) | **Patch** /v1/inventory/instances/bulk/ | 🚧 [Beta Feature]
[**V1InventoryInstancesCreate**](InventoryInstancesAPI.md#V1InventoryInstancesCreate) | **Post** /v1/inventory/instances/ | 🚧 [Beta Feature]
[**V1InventoryInstancesDestroy**](InventoryInstancesAPI.md#V1InventoryInstancesDestroy) | **Delete** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**V1InventoryInstancesHistoryList**](InventoryInstancesAPI.md#V1InventoryInstancesHistoryList) | **Get** /v1/inventory/instances/{id}/history/ | 🚧 [Beta Feature]
[**V1InventoryInstancesList**](InventoryInstancesAPI.md#V1InventoryInstancesList) | **Get** /v1/inventory/instances/ | 🚧 [Beta Feature]
[**V1InventoryInstancesPartialUpdate**](InventoryInstancesAPI.md#V1InventoryInstancesPartialUpdate) | **Patch** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**V1InventoryInstancesRetrieve**](InventoryInstancesAPI.md#V1InventoryInstancesRetrieve) | **Get** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]
[**V1InventoryInstancesUpdate**](InventoryInstancesAPI.md#V1InventoryInstancesUpdate) | **Put** /v1/inventory/instances/{id}/ | 🚧 [Beta Feature]



## V1InventoryInstancesBulkPartialUpdate

> InstanceBulkUpdate V1InventoryInstancesBulkPartialUpdate(ctx).InstanceBulkUpdate(instanceBulkUpdate).Execute()

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
	instanceBulkUpdate := *openapiclient.NewInstanceBulkUpdate(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now(), int32(123), []string{"Ids_example"}, "Result_example") // InstanceBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesBulkPartialUpdate(context.Background()).InstanceBulkUpdate(instanceBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesBulkPartialUpdate`: InstanceBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesBulkPartialUpdateRequest struct via the builder pattern


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


## V1InventoryInstancesCreate

> Instance V1InventoryInstancesCreate(ctx).Instance(instance).Execute()

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
	instance := *openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesCreate(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesCreate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesCreateRequest struct via the builder pattern


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


## V1InventoryInstancesDestroy

> V1InventoryInstancesDestroy(ctx, id).Execute()

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
	r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1InventoryInstancesDestroyRequest struct via the builder pattern


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


## V1InventoryInstancesHistoryList

> PaginatedModelChangeList V1InventoryInstancesHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()

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
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesHistoryListRequest struct via the builder pattern


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


## V1InventoryInstancesList

> PaginatedInstanceList V1InventoryInstancesList(ctx).Cluster(cluster).Expand(expand).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()

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
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesList(context.Background()).Cluster(cluster).Expand(expand).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesList`: PaginatedInstanceList
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesListRequest struct via the builder pattern


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


## V1InventoryInstancesPartialUpdate

> Instance V1InventoryInstancesPartialUpdate(ctx, id).Instance(instance).Execute()

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
	instance := *openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesPartialUpdate(context.Background(), id).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesPartialUpdate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesPartialUpdateRequest struct via the builder pattern


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


## V1InventoryInstancesRetrieve

> Instance V1InventoryInstancesRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()

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
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesRetrieve`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesRetrieveRequest struct via the builder pattern


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


## V1InventoryInstancesUpdate

> Instance V1InventoryInstancesUpdate(ctx, id).Instance(instance).Execute()

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
	instance := *openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.ClusterPipeline_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now()) // Instance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInstancesAPI.V1InventoryInstancesUpdate(context.Background(), id).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInstancesAPI.V1InventoryInstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInstancesUpdate`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InventoryInstancesAPI.V1InventoryInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInstancesUpdateRequest struct via the builder pattern


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

