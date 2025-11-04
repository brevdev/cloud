# \InventoryClustersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryClustersBulkPartialUpdate**](InventoryClustersAPI.md#V1InventoryClustersBulkPartialUpdate) | **Patch** /v1/inventory/clusters/bulk/ | 
[**V1InventoryClustersCreate**](InventoryClustersAPI.md#V1InventoryClustersCreate) | **Post** /v1/inventory/clusters/ | 
[**V1InventoryClustersDestroy**](InventoryClustersAPI.md#V1InventoryClustersDestroy) | **Delete** /v1/inventory/clusters/{id}/ | 
[**V1InventoryClustersHistoryList**](InventoryClustersAPI.md#V1InventoryClustersHistoryList) | **Get** /v1/inventory/clusters/{id}/history/ | 
[**V1InventoryClustersList**](InventoryClustersAPI.md#V1InventoryClustersList) | **Get** /v1/inventory/clusters/ | 
[**V1InventoryClustersPartialUpdate**](InventoryClustersAPI.md#V1InventoryClustersPartialUpdate) | **Patch** /v1/inventory/clusters/{id}/ | 
[**V1InventoryClustersPipelinesTriggerCreate**](InventoryClustersAPI.md#V1InventoryClustersPipelinesTriggerCreate) | **Post** /v1/inventory/clusters/{id}/pipelines/trigger/ | 
[**V1InventoryClustersProvisionCreate**](InventoryClustersAPI.md#V1InventoryClustersProvisionCreate) | **Post** /v1/inventory/clusters/provision/ | 
[**V1InventoryClustersProvisionDestroy**](InventoryClustersAPI.md#V1InventoryClustersProvisionDestroy) | **Delete** /v1/inventory/clusters/{id}/provision/ | 
[**V1InventoryClustersProvisionPartialUpdate**](InventoryClustersAPI.md#V1InventoryClustersProvisionPartialUpdate) | **Patch** /v1/inventory/clusters/{id}/provision/ | 
[**V1InventoryClustersRetrieve**](InventoryClustersAPI.md#V1InventoryClustersRetrieve) | **Get** /v1/inventory/clusters/{id}/ | 
[**V1InventoryClustersStatsRetrieve**](InventoryClustersAPI.md#V1InventoryClustersStatsRetrieve) | **Get** /v1/inventory/clusters/stats/ | 🚧 [Beta Feature]
[**V1InventoryClustersTenantsCreate**](InventoryClustersAPI.md#V1InventoryClustersTenantsCreate) | **Post** /v1/inventory/clusters/{cluster_id}/tenants/ | 
[**V1InventoryClustersTenantsDestroy**](InventoryClustersAPI.md#V1InventoryClustersTenantsDestroy) | **Delete** /v1/inventory/clusters/{cluster_id}/tenants/{id}/ | 
[**V1InventoryClustersUpdate**](InventoryClustersAPI.md#V1InventoryClustersUpdate) | **Put** /v1/inventory/clusters/{id}/ | 



## V1InventoryClustersBulkPartialUpdate

> ClusterBulkUpdate V1InventoryClustersBulkPartialUpdate(ctx).ClusterBulkUpdate(clusterBulkUpdate).Execute()



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
	clusterBulkUpdate := *openapiclient.NewClusterBulkUpdate(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "GpuAlias_example", int32(123), "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: }, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: }}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), int32(123), false, "PublicAddress_example", "RequestId_example", []string{"TenantIds_example"}, int32(123), []string{"Ids_example"}, "Result_example") // ClusterBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersBulkPartialUpdate(context.Background()).ClusterBulkUpdate(clusterBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersBulkPartialUpdate`: ClusterBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterBulkUpdate** | [**ClusterBulkUpdate**](ClusterBulkUpdate.md) |  | 

### Return type

[**ClusterBulkUpdate**](ClusterBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersCreate

> Cluster V1InventoryClustersCreate(ctx).Cluster(cluster).Execute()



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
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersCreate(context.Background()).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersCreate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**Cluster**](Cluster.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersDestroy

> V1InventoryClustersDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryClustersAPI.V1InventoryClustersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersDestroyRequest struct via the builder pattern


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


## V1InventoryClustersHistoryList

> PaginatedModelChangeList V1InventoryClustersHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersHistoryListRequest struct via the builder pattern


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


## V1InventoryClustersList

> PaginatedClusterList V1InventoryClustersList(ctx).Available(available).BastionName(bastionName).Deployment(deployment).Enabled(enabled).Expand(expand).Experience(experience).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).HasDeployment(hasDeployment).HasRequestId(hasRequestId).HasWorkshopId(hasWorkshopId).Id(id).Location(location).LocationName(locationName).LocationRegion(locationRegion).Maintenance(maintenance).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).MinNodeCount(minNodeCount).MinProvisioningAttempts(minProvisioningAttempts).MinTenantCount(minTenantCount).Netmask(netmask).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persist(persist).Provider(provider).ProviderCapacity(providerCapacity).ProviderName(providerName).ProviderNodeId(providerNodeId).ProvisionUser(provisionUser).ProvisioningAttempts(provisioningAttempts).ProvisioningState(provisioningState).PublicAddress(publicAddress).Rack(rack).RequestId(requestId).Reservation(reservation).Search(search).SystemArch(systemArch).VlanId(vlanId).Workshop(workshop).WorkshopId(workshopId).WorkshopIdNot(workshopIdNot).Execute()



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
	available := true // bool | Is the cluster currently available for provisioning? (optional)
	bastionName := "bastionName_example" // string |  (optional)
	deployment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	enabled := true // bool |  (optional)
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \"nodes.location\"). Separate multiple fields with a comma (ex: \"gpus,nodes\") (optional)
	experience := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	garageId := "garageId_example" // string | Only include clusters whose nodes have the given garage ID (optional)
	gpu := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	gpuAlias := "gpuAlias_example" // string | Alias for GPU plan (i.e. installed GPU type and count) (optional)
	gpuCount := float32(8.14) // float32 | Only include clusters with a physical GPU count equal to this value (optional)
	gpuModel := "gpuModel_example" // string | Only include clusters with the given GPU model name (optional)
	hasDeployment := true // bool |  (optional)
	hasRequestId := true // bool |  (optional)
	hasWorkshopId := true // bool |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	location := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	locationName := "locationName_example" // string | Only include clusters whose nodes are in the location with the given name (optional)
	locationRegion := "locationRegion_example" // string | Only include clusters whose nodes are in the location in the given region (optional)
	maintenance := true // bool |  (optional)
	mgmtIp := "mgmtIp_example" // string |  (optional)
	mgmtMac := "mgmtMac_example" // string |  (optional)
	minGpuCount := float32(8.14) // float32 | Only include clusters that have a gpu_count greater than or equal to this value (optional)
	minNodeCount := float32(8.14) // float32 | Only include clusters that have a node_count greater than or equal to this value (optional)
	minProvisioningAttempts := int32(56) // int32 | Only include clusters that have a provisioning_attempts value greater than or equal to this value (optional)
	minTenantCount := float32(8.14) // float32 | Only include clusters whose number of tenant_ids is greater than or equal to this value (optional)
	netmask := int32(56) // int32 |  (optional)
	nodeCount := float32(8.14) // float32 | Only include clusters with a node count equal to this value (optional)
	oem := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only include clusters with nodes that have the given OEM ID (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	persist := true // bool |  (optional)
	provider := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	providerCapacity := true // bool |  (optional)
	providerName := "providerName_example" // string | Only include clusters whose nodes are from the provider with the given name (optional)
	providerNodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	provisionUser := "provisionUser_example" // string |  (optional)
	provisioningAttempts := int32(56) // int32 |  (optional)
	provisioningState := []string{"ProvisioningState_example"} // []string | Multiple values may be separated by commas.  * `deployed` - Cluster is in use by a deployment * `deploying` - Provisioning is in progress * `destroying` - Cluster is being destroyed * `pending` - Provisioning will begin soon * `ready` - Provisioning has completed and is ready for a deployment * `reserved` - Cluster is unprovisioned but reserved for later use * `unprovisioned` - Cluster has not yet been provisioned (optional)
	publicAddress := "publicAddress_example" // string |  (optional)
	rack := "rack_example" // string | Only include clusters whose nodes are in the given rack (optional)
	requestId := "requestId_example" // string |  (optional)
	reservation := true // bool |  (optional)
	search := "search_example" // string | Search for clusters by bastion_name, experience id, gpu_alias, gpus model, id, mgmt_ip, mgmt_mac, netmask, nodes garage_id, nodes location name, nodes location provider name, nodes oem name, nodes provider_node_id, nodes rack, notes, provision_user, provisioning_state, public_address, request_id, tenant_ids, workshop_id (optional)
	systemArch := "systemArch_example" // string | Only include clusters whose nodes have the given CPU architecture (optional)
	vlanId := int32(56) // int32 |  (optional)
	workshop := true // bool |  (optional)
	workshopId := "workshopId_example" // string |  (optional)
	workshopIdNot := "workshopIdNot_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersList(context.Background()).Available(available).BastionName(bastionName).Deployment(deployment).Enabled(enabled).Expand(expand).Experience(experience).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).HasDeployment(hasDeployment).HasRequestId(hasRequestId).HasWorkshopId(hasWorkshopId).Id(id).Location(location).LocationName(locationName).LocationRegion(locationRegion).Maintenance(maintenance).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).MinNodeCount(minNodeCount).MinProvisioningAttempts(minProvisioningAttempts).MinTenantCount(minTenantCount).Netmask(netmask).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persist(persist).Provider(provider).ProviderCapacity(providerCapacity).ProviderName(providerName).ProviderNodeId(providerNodeId).ProvisionUser(provisionUser).ProvisioningAttempts(provisioningAttempts).ProvisioningState(provisioningState).PublicAddress(publicAddress).Rack(rack).RequestId(requestId).Reservation(reservation).Search(search).SystemArch(systemArch).VlanId(vlanId).Workshop(workshop).WorkshopId(workshopId).WorkshopIdNot(workshopIdNot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersList`: PaginatedClusterList
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **available** | **bool** | Is the cluster currently available for provisioning? | 
 **bastionName** | **string** |  | 
 **deployment** | **string** |  | 
 **enabled** | **bool** |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \&quot;nodes.location\&quot;). Separate multiple fields with a comma (ex: \&quot;gpus,nodes\&quot;) | 
 **experience** | **string** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **garageId** | **string** | Only include clusters whose nodes have the given garage ID | 
 **gpu** | **[]string** | Multiple values may be separated by commas. | 
 **gpuAlias** | **string** | Alias for GPU plan (i.e. installed GPU type and count) | 
 **gpuCount** | **float32** | Only include clusters with a physical GPU count equal to this value | 
 **gpuModel** | **string** | Only include clusters with the given GPU model name | 
 **hasDeployment** | **bool** |  | 
 **hasRequestId** | **bool** |  | 
 **hasWorkshopId** | **bool** |  | 
 **id** | **string** |  | 
 **location** | **string** |  | 
 **locationName** | **string** | Only include clusters whose nodes are in the location with the given name | 
 **locationRegion** | **string** | Only include clusters whose nodes are in the location in the given region | 
 **maintenance** | **bool** |  | 
 **mgmtIp** | **string** |  | 
 **mgmtMac** | **string** |  | 
 **minGpuCount** | **float32** | Only include clusters that have a gpu_count greater than or equal to this value | 
 **minNodeCount** | **float32** | Only include clusters that have a node_count greater than or equal to this value | 
 **minProvisioningAttempts** | **int32** | Only include clusters that have a provisioning_attempts value greater than or equal to this value | 
 **minTenantCount** | **float32** | Only include clusters whose number of tenant_ids is greater than or equal to this value | 
 **netmask** | **int32** |  | 
 **nodeCount** | **float32** | Only include clusters with a node count equal to this value | 
 **oem** | **string** | Only include clusters with nodes that have the given OEM ID | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **persist** | **bool** |  | 
 **provider** | **string** |  | 
 **providerCapacity** | **bool** |  | 
 **providerName** | **string** | Only include clusters whose nodes are from the provider with the given name | 
 **providerNodeId** | **string** |  | 
 **provisionUser** | **string** |  | 
 **provisioningAttempts** | **int32** |  | 
 **provisioningState** | **[]string** | Multiple values may be separated by commas.  * &#x60;deployed&#x60; - Cluster is in use by a deployment * &#x60;deploying&#x60; - Provisioning is in progress * &#x60;destroying&#x60; - Cluster is being destroyed * &#x60;pending&#x60; - Provisioning will begin soon * &#x60;ready&#x60; - Provisioning has completed and is ready for a deployment * &#x60;reserved&#x60; - Cluster is unprovisioned but reserved for later use * &#x60;unprovisioned&#x60; - Cluster has not yet been provisioned | 
 **publicAddress** | **string** |  | 
 **rack** | **string** | Only include clusters whose nodes are in the given rack | 
 **requestId** | **string** |  | 
 **reservation** | **bool** |  | 
 **search** | **string** | Search for clusters by bastion_name, experience id, gpu_alias, gpus model, id, mgmt_ip, mgmt_mac, netmask, nodes garage_id, nodes location name, nodes location provider name, nodes oem name, nodes provider_node_id, nodes rack, notes, provision_user, provisioning_state, public_address, request_id, tenant_ids, workshop_id | 
 **systemArch** | **string** | Only include clusters whose nodes have the given CPU architecture | 
 **vlanId** | **int32** |  | 
 **workshop** | **bool** |  | 
 **workshopId** | **string** |  | 
 **workshopIdNot** | **string** |  | 

### Return type

[**PaginatedClusterList**](PaginatedClusterList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersPartialUpdate

> Cluster V1InventoryClustersPartialUpdate(ctx, id).Cluster(cluster).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this cluster.
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersPartialUpdate(context.Background(), id).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersPartialUpdate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cluster** | [**Cluster**](Cluster.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersPipelinesTriggerCreate

> ClusterPipelineTrigger V1InventoryClustersPipelinesTriggerCreate(ctx, id).ClusterPipelineTrigger(clusterPipelineTrigger).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	clusterPipelineTrigger := *openapiclient.NewClusterPipelineTrigger(openapiclient.PipelineAction("apply"), time.Now(), "Id_example", time.Now(), int32(123), "RequestId_example", "Url_example") // ClusterPipelineTrigger | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersPipelinesTriggerCreate(context.Background(), id).ClusterPipelineTrigger(clusterPipelineTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersPipelinesTriggerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersPipelinesTriggerCreate`: ClusterPipelineTrigger
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersPipelinesTriggerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersPipelinesTriggerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterPipelineTrigger** | [**ClusterPipelineTrigger**](ClusterPipelineTrigger.md) |  | 

### Return type

[**ClusterPipelineTrigger**](ClusterPipelineTrigger.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersProvisionCreate

> ProvisioningRequest V1InventoryClustersProvisionCreate(ctx).ProvisioningRequest(provisioningRequest).Expand(expand).Execute()



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
	provisioningRequest := *openapiclient.NewProvisioningRequest("Experience_example", "RequestId_example") // ProvisioningRequest | 
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster,experience\"). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersProvisionCreate(context.Background()).ProvisioningRequest(provisioningRequest).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersProvisionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersProvisionCreate`: ProvisioningRequest
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersProvisionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersProvisionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningRequest** | [**ProvisioningRequest**](ProvisioningRequest.md) |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster,experience\&quot;). | 

### Return type

[**ProvisioningRequest**](ProvisioningRequest.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersProvisionDestroy

> Cluster V1InventoryClustersProvisionDestroy(ctx, id).Expand(expand).Execute()



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
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster,experience\"). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersProvisionDestroy(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersProvisionDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersProvisionDestroy`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersProvisionDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersProvisionDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster,experience\&quot;). | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersProvisionPartialUpdate

> ProvisioningRequest V1InventoryClustersProvisionPartialUpdate(ctx, id).ProvisioningRequest(provisioningRequest).Execute()



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
	provisioningRequest := *openapiclient.NewProvisioningRequest("Experience_example", "RequestId_example") // ProvisioningRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersProvisionPartialUpdate(context.Background(), id).ProvisioningRequest(provisioningRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersProvisionPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersProvisionPartialUpdate`: ProvisioningRequest
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersProvisionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersProvisionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningRequest** | [**ProvisioningRequest**](ProvisioningRequest.md) |  | 

### Return type

[**ProvisioningRequest**](ProvisioningRequest.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersRetrieve

> Cluster V1InventoryClustersRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this cluster.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \"nodes.location\"). Separate multiple fields with a comma (ex: \"gpus,nodes\") (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersRetrieve`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \&quot;nodes.location\&quot;). Separate multiple fields with a comma (ex: \&quot;gpus,nodes\&quot;) | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersStatsRetrieve

> V1InventoryClustersStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.InventoryClustersAPI.V1InventoryClustersStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersStatsRetrieveRequest struct via the builder pattern


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


## V1InventoryClustersTenantsCreate

> Cluster V1InventoryClustersTenantsCreate(ctx, clusterId).Tenant(tenant).Execute()



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenant := *openapiclient.NewTenant("Id_example") // Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersTenantsCreate(context.Background(), clusterId).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersTenantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersTenantsCreate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersTenantsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryClustersTenantsDestroy

> V1InventoryClustersTenantsDestroy(ctx, clusterId, id).Execute()



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryClustersAPI.V1InventoryClustersTenantsDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersTenantsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersTenantsDestroyRequest struct via the builder pattern


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


## V1InventoryClustersUpdate

> Cluster V1InventoryClustersUpdate(ctx, id).Cluster(cluster).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this cluster.
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.V1InventoryClustersUpdate(context.Background(), id).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.V1InventoryClustersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryClustersUpdate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.V1InventoryClustersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryClustersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cluster** | [**Cluster**](Cluster.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

