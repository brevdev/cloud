# \KubernetesAPI

All URIs are relative to *https://&lt;region&gt;.atlas.fluidstack.io/api/v1alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKubernetesCluster**](KubernetesAPI.md#CreateKubernetesCluster) | **Post** /kubernetes/clusters | Create a cluster
[**CreateKubernetesNodePool**](KubernetesAPI.md#CreateKubernetesNodePool) | **Post** /kubernetes/clusters/{cluster_id}/node-pools | Create a node pool
[**DeleteKubernetesCluster**](KubernetesAPI.md#DeleteKubernetesCluster) | **Delete** /kubernetes/clusters/{id} | Delete a cluster
[**DeleteKubernetesNodePool**](KubernetesAPI.md#DeleteKubernetesNodePool) | **Delete** /kubernetes/clusters/{cluster_id}/node-pools/{node_pool_id} | Delete a node pool
[**GetKubernetesCluster**](KubernetesAPI.md#GetKubernetesCluster) | **Get** /kubernetes/clusters/{id} | Get a cluster
[**GetKubernetesClusterCredentials**](KubernetesAPI.md#GetKubernetesClusterCredentials) | **Get** /kubernetes/clusters/{cluster_id}/credentials | Get kubeconfig credentials for a cluster
[**GetKubernetesNodePool**](KubernetesAPI.md#GetKubernetesNodePool) | **Get** /kubernetes/clusters/{cluster_id}/node-pools/{node_pool_id} | Get a node pool
[**ListKubernetesClusters**](KubernetesAPI.md#ListKubernetesClusters) | **Get** /kubernetes/clusters | List all clusters
[**ListKubernetesNodePools**](KubernetesAPI.md#ListKubernetesNodePools) | **Get** /kubernetes/clusters/{cluster_id}/node-pools | List all node pools for a cluster



## CreateKubernetesCluster

> KubernetesCluster CreateKubernetesCluster(ctx).XPROJECTID(xPROJECTID).KubernetesClustersPostRequest(kubernetesClustersPostRequest).Execute()

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
	kubernetesClustersPostRequest := *openapiclient.NewKubernetesClustersPostRequest("cluster-01") // KubernetesClustersPostRequest | Details of the cluster to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.CreateKubernetesCluster(context.Background()).XPROJECTID(xPROJECTID).KubernetesClustersPostRequest(kubernetesClustersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.CreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesCluster`: KubernetesCluster
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.CreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 
 **kubernetesClustersPostRequest** | [**KubernetesClustersPostRequest**](KubernetesClustersPostRequest.md) | Details of the cluster to create | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKubernetesNodePool

> KubernetesNodePool CreateKubernetesNodePool(ctx, clusterId).XPROJECTID(xPROJECTID).KubernetesNodePoolPostRequest(kubernetesNodePoolPostRequest).Execute()

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
	kubernetesNodePoolPostRequest := *openapiclient.NewKubernetesNodePoolPostRequest("pool-01", "h100-hgx-80gb.8x", int32(3)) // KubernetesNodePoolPostRequest | Details of the node pool to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.CreateKubernetesNodePool(context.Background(), clusterId).XPROJECTID(xPROJECTID).KubernetesNodePoolPostRequest(kubernetesNodePoolPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.CreateKubernetesNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesNodePool`: KubernetesNodePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.CreateKubernetesNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

 **kubernetesNodePoolPostRequest** | [**KubernetesNodePoolPostRequest**](KubernetesNodePoolPostRequest.md) | Details of the node pool to create | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubernetesCluster

> DeleteKubernetesCluster(ctx, id).XPROJECTID(xPROJECTID).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the kubernetes cluster to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.DeleteKubernetesCluster(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.DeleteKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the kubernetes cluster to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesClusterRequest struct via the builder pattern


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


## DeleteKubernetesNodePool

> DeleteKubernetesNodePool(ctx, clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()

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
	r, err := apiClient.KubernetesAPI.DeleteKubernetesNodePool(context.Background(), clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.DeleteKubernetesNodePool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKubernetesNodePoolRequest struct via the builder pattern


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


## GetKubernetesCluster

> KubernetesCluster GetKubernetesCluster(ctx, id).XPROJECTID(xPROJECTID).Execute()

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the kubernetes cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.GetKubernetesCluster(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesCluster`: KubernetesCluster
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the kubernetes cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusterCredentials

> string GetKubernetesClusterCredentials(ctx, clusterId).XPROJECTID(xPROJECTID).Execute()

Get kubeconfig credentials for a cluster

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the kubernetes cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.GetKubernetesClusterCredentials(context.Background(), clusterId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetKubernetesClusterCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesClusterCredentials`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetKubernetesClusterCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the kubernetes cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesNodePool

> KubernetesNodePool GetKubernetesNodePool(ctx, clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()

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
	resp, r, err := apiClient.KubernetesAPI.GetKubernetesNodePool(context.Background(), clusterId, nodePoolId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetKubernetesNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesNodePool`: KubernetesNodePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetKubernetesNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 
**nodePoolId** | **string** | Unique identifier of the node pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 



### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesClusters

> []KubernetesCluster ListKubernetesClusters(ctx).XPROJECTID(xPROJECTID).Execute()

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
	resp, r, err := apiClient.KubernetesAPI.ListKubernetesClusters(context.Background()).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ListKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesClusters`: []KubernetesCluster
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.ListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

### Return type

[**[]KubernetesCluster**](KubernetesCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesNodePools

> []KubernetesNodePool ListKubernetesNodePools(ctx, clusterId).XPROJECTID(xPROJECTID).Execute()

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
	resp, r, err := apiClient.KubernetesAPI.ListKubernetesNodePools(context.Background(), clusterId).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ListKubernetesNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesNodePools`: []KubernetesNodePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.ListKubernetesNodePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Unique identifier of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**[]KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

