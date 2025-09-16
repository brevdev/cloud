# \InventoryClustersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryClustersBulkPartialUpdate**](InventoryClustersAPI.md#InventoryClustersBulkPartialUpdate) | **Patch** /v1/inventory/clusters/bulk/ | 
[**InventoryClustersCreate**](InventoryClustersAPI.md#InventoryClustersCreate) | **Post** /v1/inventory/clusters/ | 
[**InventoryClustersDestroy**](InventoryClustersAPI.md#InventoryClustersDestroy) | **Delete** /v1/inventory/clusters/{id}/ | 
[**InventoryClustersHistoryList**](InventoryClustersAPI.md#InventoryClustersHistoryList) | **Get** /v1/inventory/clusters/{id}/history/ | 
[**InventoryClustersList**](InventoryClustersAPI.md#InventoryClustersList) | **Get** /v1/inventory/clusters/ | 
[**InventoryClustersPartialUpdate**](InventoryClustersAPI.md#InventoryClustersPartialUpdate) | **Patch** /v1/inventory/clusters/{id}/ | 
[**InventoryClustersRetrieve**](InventoryClustersAPI.md#InventoryClustersRetrieve) | **Get** /v1/inventory/clusters/{id}/ | 
[**InventoryClustersStatsRetrieve**](InventoryClustersAPI.md#InventoryClustersStatsRetrieve) | **Get** /v1/inventory/clusters/stats/ | 🚧 [Beta Feature]
[**InventoryClustersTenantsCreate**](InventoryClustersAPI.md#InventoryClustersTenantsCreate) | **Post** /v1/inventory/clusters/{cluster_id}/tenants/ | 
[**InventoryClustersTenantsDestroy**](InventoryClustersAPI.md#InventoryClustersTenantsDestroy) | **Delete** /v1/inventory/clusters/{cluster_id}/tenants/{id}/ | 
[**InventoryClustersUpdate**](InventoryClustersAPI.md#InventoryClustersUpdate) | **Put** /v1/inventory/clusters/{id}/ | 



## InventoryClustersBulkPartialUpdate

> ClusterBulkUpdate InventoryClustersBulkPartialUpdate(ctx).ClusterBulkUpdate(clusterBulkUpdate).Execute()



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
	clusterBulkUpdate := *openapiclient.NewClusterBulkUpdate(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "GpuAlias_example", int32(123), "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: }, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: }}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: }}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), int32(123), false, "PublicAddress_example", "RequestId_example", []string{"TenantIds_example"}, int32(123), []string{"Ids_example"}, "Result_example") // ClusterBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersBulkPartialUpdate(context.Background()).ClusterBulkUpdate(clusterBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersBulkPartialUpdate`: ClusterBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersBulkPartialUpdateRequest struct via the builder pattern


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


## InventoryClustersCreate

> Cluster InventoryClustersCreate(ctx).Cluster(cluster).Execute()



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
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersCreate(context.Background()).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersCreate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersCreateRequest struct via the builder pattern


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


## InventoryClustersDestroy

> InventoryClustersDestroy(ctx, id).Execute()



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
	r, err := apiClient.InventoryClustersAPI.InventoryClustersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventoryClustersDestroyRequest struct via the builder pattern


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


## InventoryClustersHistoryList

> PaginatedModelChangeList InventoryClustersHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersHistoryListRequest struct via the builder pattern


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


## InventoryClustersList

> PaginatedClusterList InventoryClustersList(ctx).Available(available).BastionName(bastionName).Deployment(deployment).Enabled(enabled).Expand(expand).Experience(experience).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).HasDeployment(hasDeployment).HasRequestId(hasRequestId).HasWorkshopId(hasWorkshopId).Id(id).Location(location).LocationName(locationName).LocationRegion(locationRegion).Maintenance(maintenance).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).MinNodeCount(minNodeCount).MinProvisioningAttempts(minProvisioningAttempts).MinTenantCount(minTenantCount).Netmask(netmask).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persist(persist).Provider(provider).ProviderCapacity(providerCapacity).ProviderName(providerName).ProviderNodeId(providerNodeId).ProvisionUser(provisionUser).ProvisioningAttempts(provisioningAttempts).ProvisioningState(provisioningState).PublicAddress(publicAddress).Rack(rack).RequestId(requestId).Reservation(reservation).Search(search).SystemArch(systemArch).VlanId(vlanId).Workshop(workshop).WorkshopId(workshopId).WorkshopIdNot(workshopIdNot).Execute()



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
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersList(context.Background()).Available(available).BastionName(bastionName).Deployment(deployment).Enabled(enabled).Expand(expand).Experience(experience).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).HasDeployment(hasDeployment).HasRequestId(hasRequestId).HasWorkshopId(hasWorkshopId).Id(id).Location(location).LocationName(locationName).LocationRegion(locationRegion).Maintenance(maintenance).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).MinNodeCount(minNodeCount).MinProvisioningAttempts(minProvisioningAttempts).MinTenantCount(minTenantCount).Netmask(netmask).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persist(persist).Provider(provider).ProviderCapacity(providerCapacity).ProviderName(providerName).ProviderNodeId(providerNodeId).ProvisionUser(provisionUser).ProvisioningAttempts(provisioningAttempts).ProvisioningState(provisioningState).PublicAddress(publicAddress).Rack(rack).RequestId(requestId).Reservation(reservation).Search(search).SystemArch(systemArch).VlanId(vlanId).Workshop(workshop).WorkshopId(workshopId).WorkshopIdNot(workshopIdNot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersList`: PaginatedClusterList
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersListRequest struct via the builder pattern


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


## InventoryClustersPartialUpdate

> Cluster InventoryClustersPartialUpdate(ctx, id).Cluster(cluster).Execute()



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
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersPartialUpdate(context.Background(), id).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersPartialUpdate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersPartialUpdateRequest struct via the builder pattern


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


## InventoryClustersRetrieve

> Cluster InventoryClustersRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersRetrieve`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersRetrieveRequest struct via the builder pattern


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


## InventoryClustersStatsRetrieve

> InventoryClustersStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.InventoryClustersAPI.InventoryClustersStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersStatsRetrieveRequest struct via the builder pattern


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


## InventoryClustersTenantsCreate

> Cluster InventoryClustersTenantsCreate(ctx, clusterId).Tenant(tenant).Execute()



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
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersTenantsCreate(context.Background(), clusterId).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersTenantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersTenantsCreate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersTenantsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersTenantsCreateRequest struct via the builder pattern


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


## InventoryClustersTenantsDestroy

> InventoryClustersTenantsDestroy(ctx, clusterId, id).Execute()



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
	r, err := apiClient.InventoryClustersAPI.InventoryClustersTenantsDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersTenantsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInventoryClustersTenantsDestroyRequest struct via the builder pattern


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


## InventoryClustersUpdate

> Cluster InventoryClustersUpdate(ctx, id).Cluster(cluster).Execute()



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
	cluster := *openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: openapiclient.NewCluster(false, time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, []openapiclient.ClusterGpusInner{openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}}, "Id_example", []openapiclient.ClusterInstancesInner{openapiclient.Cluster_instances_inner{Instance: openapiclient.NewInstance(openapiclient.Deployment_cluster{Cluster: }, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"})}, time.Now(), "Id_example", "InstanceId_example", time.Now(), time.Now())}}, time.Now(), time.Now(), false, []string{"TenantIds_example"}) // Cluster |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryClustersAPI.InventoryClustersUpdate(context.Background(), id).Cluster(cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryClustersAPI.InventoryClustersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryClustersUpdate`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `InventoryClustersAPI.InventoryClustersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryClustersUpdateRequest struct via the builder pattern


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

