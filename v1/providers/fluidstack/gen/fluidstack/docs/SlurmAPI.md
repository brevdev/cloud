# \SlurmAPI

All URIs are relative to *https://&lt;region&gt;.atlas.fluidstack.io/api/v1alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlurmCluster**](SlurmAPI.md#CreateSlurmCluster) | **Post** /slurm/clusters | Create a cluster
[**CreateSlurmNodePool**](SlurmAPI.md#CreateSlurmNodePool) | **Post** /slurm/clusters/{cluster_id}/node-pools | Create a node pool
[**DeleteSlurmCluster**](SlurmAPI.md#DeleteSlurmCluster) | **Delete** /slurm/clusters/{id} | Delete a cluster
[**DeleteSlurmNodePool**](SlurmAPI.md#DeleteSlurmNodePool) | **Delete** /slurm/clusters/{cluster_id}/node-pools/{node_pool_id} | Delete a node pool
[**GetSlurmCluster**](SlurmAPI.md#GetSlurmCluster) | **Get** /slurm/clusters/{id} | Get a cluster
[**GetSlurmNodePool**](SlurmAPI.md#GetSlurmNodePool) | **Get** /slurm/clusters/{cluster_id}/node-pools/{node_pool_id} | Get a node pool
[**ListSlurmClusters**](SlurmAPI.md#ListSlurmClusters) | **Get** /slurm/clusters | List all clusters
[**ListSlurmNodePools**](SlurmAPI.md#ListSlurmNodePools) | **Get** /slurm/clusters/{cluster_id}/node-pools | List all node pools for a cluster



## CreateSlurmCluster

> SlurmCluster CreateSlurmCluster(ctx).XPROJECTID(xPROJECTID).SlurmClustersPostRequest(slurmClustersPostRequest).Execute()

Create a cluster

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	slurmClustersPostRequest := *openapiclient.NewSlurmClustersPostRequest("cluster-01") // SlurmClustersPostRequest | Details of the cluster to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.CreateSlurmCluster(context.Background()).XPROJECTID(xPROJECTID).SlurmClustersPostRequest(slurmClustersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.CreateSlurmCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlurmCluster`: SlurmCluster
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.CreateSlurmCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlurmClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 
 **slurmClustersPostRequest** | [**SlurmClustersPostRequest**](SlurmClustersPostRequest.md) | Details of the cluster to create | 

### Return type

[**SlurmCluster**](SlurmCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSlurmNodePool

> SlurmNodePool CreateSlurmNodePool(ctx, clusterId).XPROJECTID(xPROJECTID).SlurmNodePoolPostRequest(slurmNodePoolPostRequest).Execute()

Create a node pool

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster
	slurmNodePoolPostRequest := *openapiclient.NewSlurmNodePoolPostRequest("pool-01", "h100-hgx-80gb.8x", int32(3)) // SlurmNodePoolPostRequest | Details of the node pool to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.CreateSlurmNodePool(context.Background(), clusterId).XPROJECTID(xPROJECTID).SlurmNodePoolPostRequest(slurmNodePoolPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.CreateSlurmNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlurmNodePool`: SlurmNodePool
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.CreateSlurmNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlurmNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

 **slurmNodePoolPostRequest** | [**SlurmNodePoolPostRequest**](SlurmNodePoolPostRequest.md) | Details of the node pool to create | 

### Return type

[**SlurmNodePool**](SlurmNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlurmCluster

> DeleteSlurmCluster(ctx, id).XPROJECTID(xPROJECTID).Execute()

Delete a cluster

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlurmAPI.DeleteSlurmCluster(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.DeleteSlurmCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the cluster to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlurmClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlurmNodePool

> DeleteSlurmNodePool(ctx, clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()

Delete a node pool

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster
	nodePoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the node pool to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlurmAPI.DeleteSlurmNodePool(context.Background(), clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.DeleteSlurmNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 
**nodePoolId** | **string** | Unique identifier of the node pool to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlurmNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlurmCluster

> SlurmCluster GetSlurmCluster(ctx, id).XPROJECTID(xPROJECTID).Execute()

Get a cluster

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.GetSlurmCluster(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.GetSlurmCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlurmCluster`: SlurmCluster
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.GetSlurmCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlurmClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**SlurmCluster**](SlurmCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlurmNodePool

> SlurmNodePool GetSlurmNodePool(ctx, clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()

Get a node pool

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster
	nodePoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the node pool

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.GetSlurmNodePool(context.Background(), clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.GetSlurmNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlurmNodePool`: SlurmNodePool
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.GetSlurmNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 
**nodePoolId** | **string** | Unique identifier of the node pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlurmNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 



### Return type

[**SlurmNodePool**](SlurmNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlurmClusters

> []SlurmCluster ListSlurmClusters(ctx).XPROJECTID(xPROJECTID).Execute()

List all clusters

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.ListSlurmClusters(context.Background()).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.ListSlurmClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlurmClusters`: []SlurmCluster
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.ListSlurmClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSlurmClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

### Return type

[**[]SlurmCluster**](SlurmCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlurmNodePools

> []SlurmNodePool ListSlurmNodePools(ctx, clusterId).XPROJECTID(xPROJECTID).Execute()

List all node pools for a cluster

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.ListSlurmNodePools(context.Background(), clusterId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.ListSlurmNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlurmNodePools`: []SlurmNodePool
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.ListSlurmNodePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlurmNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**[]SlurmNodePool**](SlurmNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

