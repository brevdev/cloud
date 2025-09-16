# \InventoryNodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryNodesBulkCreate**](InventoryNodesAPI.md#InventoryNodesBulkCreate) | **Post** /v1/inventory/nodes/bulk/ | 
[**InventoryNodesBulkPartialUpdate**](InventoryNodesAPI.md#InventoryNodesBulkPartialUpdate) | **Patch** /v1/inventory/nodes/bulk/ | 
[**InventoryNodesCreate**](InventoryNodesAPI.md#InventoryNodesCreate) | **Post** /v1/inventory/nodes/ | 
[**InventoryNodesDestroy**](InventoryNodesAPI.md#InventoryNodesDestroy) | **Delete** /v1/inventory/nodes/{id}/ | 
[**InventoryNodesHistoryList**](InventoryNodesAPI.md#InventoryNodesHistoryList) | **Get** /v1/inventory/nodes/{id}/history/ | 
[**InventoryNodesList**](InventoryNodesAPI.md#InventoryNodesList) | **Get** /v1/inventory/nodes/ | 
[**InventoryNodesPartialUpdate**](InventoryNodesAPI.md#InventoryNodesPartialUpdate) | **Patch** /v1/inventory/nodes/{id}/ | 
[**InventoryNodesRetrieve**](InventoryNodesAPI.md#InventoryNodesRetrieve) | **Get** /v1/inventory/nodes/{id}/ | 
[**InventoryNodesUpdate**](InventoryNodesAPI.md#InventoryNodesUpdate) | **Put** /v1/inventory/nodes/{id}/ | 



## InventoryNodesBulkCreate

> NodeBulk InventoryNodesBulkCreate(ctx).CsvFile(csvFile).Execute()





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
	csvFile := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesBulkCreate(context.Background()).CsvFile(csvFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesBulkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesBulkCreate`: NodeBulk
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesBulkCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csvFile** | ***os.File** |  | 

### Return type

[**NodeBulk**](NodeBulk.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesBulkPartialUpdate

> NodeBulkUpdate InventoryNodesBulkPartialUpdate(ctx).NodeBulkUpdate(nodeBulkUpdate).Execute()



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
	nodeBulkUpdate := *openapiclient.NewNodeBulkUpdate(time.Now(), openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}, "GpuAlias_example", "GpuModel_example", "Id_example", openapiclient.Node_location{Location: openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")})}, time.Now(), "SerialNumber_example", []openapiclient.NodeStorage{*openapiclient.NewNodeStorage(openapiclient.TypeEnum("nvme"))}, openapiclient.SystemArchEnum("amd64"), int32(123), []string{"Ids_example"}, "Result_example") // NodeBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesBulkPartialUpdate(context.Background()).NodeBulkUpdate(nodeBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesBulkPartialUpdate`: NodeBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeBulkUpdate** | [**NodeBulkUpdate**](NodeBulkUpdate.md) |  | 

### Return type

[**NodeBulkUpdate**](NodeBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesCreate

> Node InventoryNodesCreate(ctx).Node(node).Execute()



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
	node := *openapiclient.NewNode(time.Now(), openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}, "GpuAlias_example", "GpuModel_example", "Id_example", openapiclient.Node_location{Location: openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")})}, time.Now(), []openapiclient.NodeStorage{*openapiclient.NewNodeStorage(openapiclient.TypeEnum("nvme"))}, openapiclient.SystemArchEnum("amd64")) // Node | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesCreate(context.Background()).Node(node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesCreate`: Node
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **node** | [**Node**](Node.md) |  | 

### Return type

[**Node**](Node.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesDestroy

> InventoryNodesDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryNodesAPI.InventoryNodesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesDestroyRequest struct via the builder pattern


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


## InventoryNodesHistoryList

> PaginatedModelChangeList InventoryNodesHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesHistoryListRequest struct via the builder pattern


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


## InventoryNodesList

> PaginatedNodeList InventoryNodesList(ctx).BmcIp(bmcIp).BmcMac(bmcMac).BmcPassword(bmcPassword).BmcUser(bmcUser).Cluster(cluster).Cpu(cpu).CpuManufacturer(cpuManufacturer).CpuModel(cpuModel).Expand(expand).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuVbios(gpuVbios).Id(id).Location(location).Memory(memory).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).Model(model).NetworkType(networkType).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).ProviderNodeId(providerNodeId).Rack(rack).RackUnit(rackUnit).Search(search).SerialNumber(serialNumber).SystemArch(systemArch).Tee(tee).Execute()



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
	bmcIp := "bmcIp_example" // string |  (optional)
	bmcMac := "bmcMac_example" // string |  (optional)
	bmcPassword := "bmcPassword_example" // string |  (optional)
	bmcUser := "bmcUser_example" // string |  (optional)
	cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	cpu := int32(56) // int32 |  (optional)
	cpuManufacturer := "cpuManufacturer_example" // string | Manufacturer of the node's CPU  * `amd` - AMD * `arm` - ARM * `intel` - Intel (optional)
	cpuModel := "cpuModel_example" // string |  (optional)
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \"location.provider\"). Separate multiple fields with a comma (ex: \"location,oem\") (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	garageId := "garageId_example" // string |  (optional)
	gpu := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	gpuAlias := "gpuAlias_example" // string | Alias for GPU plan (i.e. installed GPU type and count) (optional)
	gpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string | Model of GPU(s) installed (optional)
	gpuVbios := "gpuVbios_example" // string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	location := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	memory := int32(56) // int32 |  (optional)
	mgmtIp := "mgmtIp_example" // string |  (optional)
	mgmtMac := "mgmtMac_example" // string |  (optional)
	minGpuCount := int32(56) // int32 | Only include nodes that have a gpu_count greater than or equal to this value (optional)
	model := "model_example" // string |  (optional)
	networkType := "networkType_example" // string | Type of networking technology used  * `ethernet` - Ethernet * `infiniband` - InfiniBand (optional)
	oem := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	provider := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	providerNodeId := "providerNodeId_example" // string |  (optional)
	rack := "rack_example" // string |  (optional)
	rackUnit := int32(56) // int32 |  (optional)
	search := "search_example" // string | Search for nodes by bmc_ip, bmc_mac, bmc_password, bmc_user, cpu_manufacturer, cpu_model, garage_id, gpu_alias, gpu model, gpu_vbios, id, location name, location provider name, memory, mgmt_ip, mgmt_mac, model, network_type, nic_prefixes, notes, oem name, provider_node_id, rack, rack_unit, serial_number, storage, system_arch (optional)
	serialNumber := "serialNumber_example" // string |  (optional)
	systemArch := "systemArch_example" // string | CPU architecture  * `amd64` - amd64 * `arm64` - arm64 (optional)
	tee := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesList(context.Background()).BmcIp(bmcIp).BmcMac(bmcMac).BmcPassword(bmcPassword).BmcUser(bmcUser).Cluster(cluster).Cpu(cpu).CpuManufacturer(cpuManufacturer).CpuModel(cpuModel).Expand(expand).Fields(fields).GarageId(garageId).Gpu(gpu).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuVbios(gpuVbios).Id(id).Location(location).Memory(memory).MgmtIp(mgmtIp).MgmtMac(mgmtMac).MinGpuCount(minGpuCount).Model(model).NetworkType(networkType).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).ProviderNodeId(providerNodeId).Rack(rack).RackUnit(rackUnit).Search(search).SerialNumber(serialNumber).SystemArch(systemArch).Tee(tee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesList`: PaginatedNodeList
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bmcIp** | **string** |  | 
 **bmcMac** | **string** |  | 
 **bmcPassword** | **string** |  | 
 **bmcUser** | **string** |  | 
 **cluster** | **string** |  | 
 **cpu** | **int32** |  | 
 **cpuManufacturer** | **string** | Manufacturer of the node&#39;s CPU  * &#x60;amd&#x60; - AMD * &#x60;arm&#x60; - ARM * &#x60;intel&#x60; - Intel | 
 **cpuModel** | **string** |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \&quot;location.provider\&quot;). Separate multiple fields with a comma (ex: \&quot;location,oem\&quot;) | 
 **fields** | **string** | Include only the specified fields in the response | 
 **garageId** | **string** |  | 
 **gpu** | **[]string** | Multiple values may be separated by commas. | 
 **gpuAlias** | **string** | Alias for GPU plan (i.e. installed GPU type and count) | 
 **gpuCount** | **int32** |  | 
 **gpuModel** | **string** | Model of GPU(s) installed | 
 **gpuVbios** | **string** |  | 
 **id** | **string** |  | 
 **location** | **string** |  | 
 **memory** | **int32** |  | 
 **mgmtIp** | **string** |  | 
 **mgmtMac** | **string** |  | 
 **minGpuCount** | **int32** | Only include nodes that have a gpu_count greater than or equal to this value | 
 **model** | **string** |  | 
 **networkType** | **string** | Type of networking technology used  * &#x60;ethernet&#x60; - Ethernet * &#x60;infiniband&#x60; - InfiniBand | 
 **oem** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **string** |  | 
 **providerNodeId** | **string** |  | 
 **rack** | **string** |  | 
 **rackUnit** | **int32** |  | 
 **search** | **string** | Search for nodes by bmc_ip, bmc_mac, bmc_password, bmc_user, cpu_manufacturer, cpu_model, garage_id, gpu_alias, gpu model, gpu_vbios, id, location name, location provider name, memory, mgmt_ip, mgmt_mac, model, network_type, nic_prefixes, notes, oem name, provider_node_id, rack, rack_unit, serial_number, storage, system_arch | 
 **serialNumber** | **string** |  | 
 **systemArch** | **string** | CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
 **tee** | **bool** |  | 

### Return type

[**PaginatedNodeList**](PaginatedNodeList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesPartialUpdate

> Node InventoryNodesPartialUpdate(ctx, id).Node(node).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this node.
	node := *openapiclient.NewNode(time.Now(), openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}, "GpuAlias_example", "GpuModel_example", "Id_example", openapiclient.Node_location{Location: openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")})}, time.Now(), []openapiclient.NodeStorage{*openapiclient.NewNodeStorage(openapiclient.TypeEnum("nvme"))}, openapiclient.SystemArchEnum("amd64")) // Node | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesPartialUpdate(context.Background(), id).Node(node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesPartialUpdate`: Node
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | [**Node**](Node.md) |  | 

### Return type

[**Node**](Node.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesRetrieve

> Node InventoryNodesRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this node.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \"location.provider\"). Separate multiple fields with a comma (ex: \"location,oem\") (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesRetrieve`: Node
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID. Separate nested relationships with a period (ex: \&quot;location.provider\&quot;). Separate multiple fields with a comma (ex: \&quot;location,oem\&quot;) | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Node**](Node.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryNodesUpdate

> Node InventoryNodesUpdate(ctx, id).Node(node).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this node.
	node := *openapiclient.NewNode(time.Now(), openapiclient.Cluster_gpus_inner{Gpu: openapiclient.NewGpu(time.Now(), "Id_example", "Model_example", time.Now())}, "GpuAlias_example", "GpuModel_example", "Id_example", openapiclient.Node_location{Location: openapiclient.NewLocation(time.Now(), "Id_example", time.Now(), "Name_example", openapiclient.Location_provider{Provider: openapiclient.NewProvider(time.Now(), "Id_example", time.Now(), "Name_example")})}, time.Now(), []openapiclient.NodeStorage{*openapiclient.NewNodeStorage(openapiclient.TypeEnum("nvme"))}, openapiclient.SystemArchEnum("amd64")) // Node | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryNodesAPI.InventoryNodesUpdate(context.Background(), id).Node(node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryNodesAPI.InventoryNodesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryNodesUpdate`: Node
	fmt.Fprintf(os.Stdout, "Response from `InventoryNodesAPI.InventoryNodesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryNodesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | [**Node**](Node.md) |  | 

### Return type

[**Node**](Node.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

