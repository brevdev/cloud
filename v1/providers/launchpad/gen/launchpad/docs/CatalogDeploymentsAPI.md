# \CatalogDeploymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogDeploymentsBulkPartialUpdate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsBulkPartialUpdate) | **Patch** /v1/catalog/deployments/bulk/ | 
[**V1CatalogDeploymentsCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsCreate) | **Post** /v1/catalog/deployments/ | 
[**V1CatalogDeploymentsDestroy**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsDestroy) | **Delete** /v1/catalog/deployments/{id}/ | 
[**V1CatalogDeploymentsHistoryList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsHistoryList) | **Get** /v1/catalog/deployments/{id}/history/ | 
[**V1CatalogDeploymentsInstancesList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsInstancesList) | **Get** /v1/catalog/deployments/{deployment_id}/instances/ | 🚧 [Beta Feature]
[**V1CatalogDeploymentsList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsList) | **Get** /v1/catalog/deployments/ | 
[**V1CatalogDeploymentsNotesCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/notes/ | 
[**V1CatalogDeploymentsNotesDestroy**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesDestroy) | **Delete** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**V1CatalogDeploymentsNotesList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesList) | **Get** /v1/catalog/deployments/{deployment_id}/notes/ | 
[**V1CatalogDeploymentsNotesPartialUpdate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesPartialUpdate) | **Patch** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**V1CatalogDeploymentsNotesRetrieve**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesRetrieve) | **Get** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**V1CatalogDeploymentsNotesUpdate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsNotesUpdate) | **Put** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**V1CatalogDeploymentsPartialUpdate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsPartialUpdate) | **Patch** /v1/catalog/deployments/{id}/ | 
[**V1CatalogDeploymentsPipelinesCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsPipelinesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/pipelines/ | 
[**V1CatalogDeploymentsPipelinesList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsPipelinesList) | **Get** /v1/catalog/deployments/{deployment_id}/pipelines/ | 
[**V1CatalogDeploymentsRetrieve**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsRetrieve) | **Get** /v1/catalog/deployments/{id}/ | 
[**V1CatalogDeploymentsServicesCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsServicesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/services/ | 
[**V1CatalogDeploymentsServicesList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsServicesList) | **Get** /v1/catalog/deployments/{deployment_id}/services/ | 
[**V1CatalogDeploymentsSshKeysCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsSshKeysCreate) | **Post** /v1/catalog/deployments/{deployment_id}/ssh-keys/ | 
[**V1CatalogDeploymentsSshKeysDestroy**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsSshKeysDestroy) | **Delete** /v1/catalog/deployments/{deployment_id}/ssh-keys/{id}/ | 
[**V1CatalogDeploymentsSshKeysList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsSshKeysList) | **Get** /v1/catalog/deployments/{deployment_id}/ssh-keys/ | 
[**V1CatalogDeploymentsStatsRetrieve**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsStatsRetrieve) | **Get** /v1/catalog/deployments/stats/ | 🚧 [Beta Feature]
[**V1CatalogDeploymentsTasksCreate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsTasksCreate) | **Post** /v1/catalog/deployments/{deployment_id}/tasks/ | 🚧 [Beta Feature]
[**V1CatalogDeploymentsTasksList**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsTasksList) | **Get** /v1/catalog/deployments/{deployment_id}/tasks/ | 🚧 [Beta Feature]
[**V1CatalogDeploymentsTasksRetrieve**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsTasksRetrieve) | **Get** /v1/catalog/deployments/{deployment_id}/tasks/{id}/ | 🚧 [Beta Feature]
[**V1CatalogDeploymentsUpdate**](CatalogDeploymentsAPI.md#V1CatalogDeploymentsUpdate) | **Put** /v1/catalog/deployments/{id}/ | 
[**V1CatalogExperiencesNotesCreate**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesCreate) | **Post** /v1/catalog/experiences/{experience_id}/notes/ | 
[**V1CatalogExperiencesNotesDestroy**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesDestroy) | **Delete** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesList**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesList) | **Get** /v1/catalog/experiences/{experience_id}/notes/ | 
[**V1CatalogExperiencesNotesPartialUpdate**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesPartialUpdate) | **Patch** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesRetrieve**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesRetrieve) | **Get** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesUpdate**](CatalogDeploymentsAPI.md#V1CatalogExperiencesNotesUpdate) | **Put** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 



## V1CatalogDeploymentsBulkPartialUpdate

> DeploymentBulkUpdate V1CatalogDeploymentsBulkPartialUpdate(ctx).DeploymentBulkUpdate(deploymentBulkUpdate).Execute()



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
	deploymentBulkUpdate := *openapiclient.NewDeploymentBulkUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), "Experience_example", "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example", int32(123), []string{"Ids_example"}, "Result_example") // DeploymentBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsBulkPartialUpdate(context.Background()).DeploymentBulkUpdate(deploymentBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsBulkPartialUpdate`: DeploymentBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentBulkUpdate** | [**DeploymentBulkUpdate**](DeploymentBulkUpdate.md) |  | 

### Return type

[**DeploymentBulkUpdate**](DeploymentBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsCreate

> Deployment V1CatalogDeploymentsCreate(ctx).Deployment(deployment).Expand(expand).Execute()



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
	deployment := *openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed")) // Deployment | 
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsCreate(context.Background()).Deployment(deployment).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsCreate`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment** | [**Deployment**](Deployment.md) |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster\&quot;). | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsDestroy

> DocDeploymentDelete V1CatalogDeploymentsDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this deployment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsDestroy`: DocDeploymentDelete
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocDeploymentDelete**](DocDeploymentDelete.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsHistoryList

> PaginatedModelChangeList V1CatalogDeploymentsHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsHistoryListRequest struct via the builder pattern


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


## V1CatalogDeploymentsInstancesList

> PaginatedDeploymentInstanceList V1CatalogDeploymentsInstancesList(ctx, deploymentId).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()

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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	instanceId := "instanceId_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for deployment-instances by id, instance_id, name, state, tags (optional)
	state := "state_example" // string | Current lifecycle state of this instance  * `running` - Instance is running * `starting` - Instance is starting * `stopped` - Instance is stopped * `stopping` - Instance is stopping * `unknown` - Instance state is currently unknown (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsInstancesList(context.Background(), deploymentId).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsInstancesList`: PaginatedDeploymentInstanceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsInstancesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **instanceId** | **string** |  | 
 **name** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for deployment-instances by id, instance_id, name, state, tags | 
 **state** | **string** | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | 

### Return type

[**PaginatedDeploymentInstanceList**](PaginatedDeploymentInstanceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsList

> PaginatedDeploymentList V1CatalogDeploymentsList(ctx).BastionOperatingSystem(bastionOperatingSystem).Cluster(cluster).ClusterGpusModel(clusterGpusModel).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Expired(expired).ExpiresAt(expiresAt).Expiring(expiring).Fields(fields).FlightcontrolRelease(flightcontrolRelease).GarageId(garageId).GcBranch(gcBranch).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).OemName(oemName).Omit(omit).Ordering(ordering).OrgName(orgName).Page(page).PageSize(pageSize).PersistOnFailure(persistOnFailure).Persona(persona).Pipeline(pipeline).PipelineBranch(pipelineBranch).Platform(platform).Priority(priority).ProviderName(providerName).Region(region).RequestId(requestId).RequesterEmail(requesterEmail).RequesterName(requesterName).RuntimeBranch(runtimeBranch).RuntimeCnsAddonPack(runtimeCnsAddonPack).RuntimeCnsDocker(runtimeCnsDocker).RuntimeCnsDriverVersion(runtimeCnsDriverVersion).RuntimeCnsK8s(runtimeCnsK8s).RuntimeCnsNvidiaDriver(runtimeCnsNvidiaDriver).RuntimeCnsVersion(runtimeCnsVersion).RuntimeMig(runtimeMig).RuntimeMigProfile(runtimeMigProfile).SalesId(salesId).SalesOwnerEmail(salesOwnerEmail).SalesOwnerName(salesOwnerName).Search(search).State(state).StateNot(stateNot).Workshop(workshop).WorkshopId(workshopId).Execute()



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
	bastionOperatingSystem := "bastionOperatingSystem_example" // string |  (optional)
	cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	clusterGpusModel := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	collectionBranch := "collectionBranch_example" // string |  (optional)
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)
	experience := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	experienceBranch := "experienceBranch_example" // string |  (optional)
	expired := true // bool |  (optional)
	expiresAt := time.Now() // string |  (optional)
	expiring := "expiring_example" // string | Include deployments whose expires_at value is within the given range (inclusive), specified as \"today\", \"tomorrow\", or \"{start}[,end]\". Start and end times must be in ISO format. (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	flightcontrolRelease := "flightcontrolRelease_example" // string |  (optional)
	garageId := "garageId_example" // string |  (optional)
	gcBranch := "gcBranch_example" // string |  (optional)
	gpuAlias := "gpuAlias_example" // string |  (optional)
	gpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string |  (optional)
	gpuOsName := "gpuOsName_example" // string |  (optional)
	gpuOsRelease := "gpuOsRelease_example" // string |  (optional)
	gpuOsVersion := "gpuOsVersion_example" // string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	nodeCount := int32(56) // int32 |  (optional)
	oemName := "oemName_example" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	orgName := "orgName_example" // string |  (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	persistOnFailure := true // bool |  (optional)
	persona := "persona_example" // string |  (optional)
	pipeline := int32(56) // int32 |  (optional)
	pipelineBranch := "pipelineBranch_example" // string |  (optional)
	platform := "platform_example" // string |  (optional)
	priority := "priority_example" // string | Priority level for the request  * `p0` - p0 * `p1` - p1 * `p2` - p2 * `p3` - p3 (optional)
	providerName := "providerName_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	requestId := "requestId_example" // string |  (optional)
	requesterEmail := "requesterEmail_example" // string |  (optional)
	requesterName := "requesterName_example" // string |  (optional)
	runtimeBranch := "runtimeBranch_example" // string |  (optional)
	runtimeCnsAddonPack := true // bool |  (optional)
	runtimeCnsDocker := true // bool |  (optional)
	runtimeCnsDriverVersion := "runtimeCnsDriverVersion_example" // string |  (optional)
	runtimeCnsK8s := true // bool |  (optional)
	runtimeCnsNvidiaDriver := true // bool |  (optional)
	runtimeCnsVersion := "runtimeCnsVersion_example" // string |  (optional)
	runtimeMig := true // bool |  (optional)
	runtimeMigProfile := "runtimeMigProfile_example" // string |  (optional)
	salesId := "salesId_example" // string |  (optional)
	salesOwnerEmail := "salesOwnerEmail_example" // string |  (optional)
	salesOwnerName := "salesOwnerName_example" // string |  (optional)
	search := "search_example" // string | Search for deployments by experience catalog_id, experience catalog_id_alias, experience id, experience title, expires_at, id, org_name, provisioning_config bastion_operating_system, provisioning_config collection_branch, provisioning_config experience_branch, provisioning_config flightcontrol_release, provisioning_config garage_id, provisioning_config gc_branch, provisioning_config gpu_alias, provisioning_config gpu_model, provisioning_config gpu_os_name, provisioning_config gpu_os_release, provisioning_config gpu_os_version, provisioning_config oem_name, provisioning_config persona, provisioning_config pipeline_branch, provisioning_config platform, provisioning_config provider_name, provisioning_config region, provisioning_config runtime_branch, provisioning_config runtime_cns_driver_version, provisioning_config runtime_cns_version, provisioning_config runtime_mig_profile, provisioning_config runtime_url, provisioning_config workshop_id, provisioning_request bastion_operating_system, provisioning_request collection_branch, provisioning_request experience_branch, provisioning_request flightcontrol_release, provisioning_request garage_id, provisioning_request gc_branch, provisioning_request gpu_alias, provisioning_request gpu_model, provisioning_request gpu_os_name, provisioning_request gpu_os_release, provisioning_request gpu_os_version, provisioning_request oem_name, provisioning_request persona, provisioning_request pipeline_branch, provisioning_request platform, provisioning_request provider_name, provisioning_request region, provisioning_request runtime_branch, provisioning_request runtime_cns_driver_version, provisioning_request runtime_cns_version, provisioning_request runtime_mig_profile, provisioning_request runtime_url, provisioning_request workshop_id, request_id, requester_email, requester_name, sales_id, sales_owner_email, sales_owner_name, services url, state, tags (optional)
	state := []string{"State_example"} // []string | Multiple values may be separated by commas.  * `destroyed` - Deployment has been fully destroyed * `destroying` - Deployment is being destroyed * `error` - Deployment has encountered a fatal error and will not be retried * `failed` - Deployment has failed but may be retried * `paused` - Deployment is paused but may be retried later * `ready` - Deployment is ready and all instances are running * `retrying` - Deployment is retrying * `starting` - Deployment instances are starting * `stopped` - Deployment instances are stopped * `stopping` - Deployment instances are stopping * `waiting` - Waiting for deployment to be ready (optional)
	stateNot := []string{"StateNot_example"} // []string | Multiple values may be separated by commas.  * `destroyed` - Deployment has been fully destroyed * `destroying` - Deployment is being destroyed * `error` - Deployment has encountered a fatal error and will not be retried * `failed` - Deployment has failed but may be retried * `paused` - Deployment is paused but may be retried later * `ready` - Deployment is ready and all instances are running * `retrying` - Deployment is retrying * `starting` - Deployment instances are starting * `stopped` - Deployment instances are stopped * `stopping` - Deployment instances are stopping * `waiting` - Waiting for deployment to be ready (optional)
	workshop := true // bool |  (optional)
	workshopId := "workshopId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsList(context.Background()).BastionOperatingSystem(bastionOperatingSystem).Cluster(cluster).ClusterGpusModel(clusterGpusModel).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Expired(expired).ExpiresAt(expiresAt).Expiring(expiring).Fields(fields).FlightcontrolRelease(flightcontrolRelease).GarageId(garageId).GcBranch(gcBranch).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).OemName(oemName).Omit(omit).Ordering(ordering).OrgName(orgName).Page(page).PageSize(pageSize).PersistOnFailure(persistOnFailure).Persona(persona).Pipeline(pipeline).PipelineBranch(pipelineBranch).Platform(platform).Priority(priority).ProviderName(providerName).Region(region).RequestId(requestId).RequesterEmail(requesterEmail).RequesterName(requesterName).RuntimeBranch(runtimeBranch).RuntimeCnsAddonPack(runtimeCnsAddonPack).RuntimeCnsDocker(runtimeCnsDocker).RuntimeCnsDriverVersion(runtimeCnsDriverVersion).RuntimeCnsK8s(runtimeCnsK8s).RuntimeCnsNvidiaDriver(runtimeCnsNvidiaDriver).RuntimeCnsVersion(runtimeCnsVersion).RuntimeMig(runtimeMig).RuntimeMigProfile(runtimeMigProfile).SalesId(salesId).SalesOwnerEmail(salesOwnerEmail).SalesOwnerName(salesOwnerName).Search(search).State(state).StateNot(stateNot).Workshop(workshop).WorkshopId(workshopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsList`: PaginatedDeploymentList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bastionOperatingSystem** | **string** |  | 
 **cluster** | **string** |  | 
 **clusterGpusModel** | **[]string** | Multiple values may be separated by commas. | 
 **collectionBranch** | **string** |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster\&quot;). | 
 **experience** | **string** |  | 
 **experienceBranch** | **string** |  | 
 **expired** | **bool** |  | 
 **expiresAt** | **string** |  | 
 **expiring** | **string** | Include deployments whose expires_at value is within the given range (inclusive), specified as \&quot;today\&quot;, \&quot;tomorrow\&quot;, or \&quot;{start}[,end]\&quot;. Start and end times must be in ISO format. | 
 **fields** | **string** | Include only the specified fields in the response | 
 **flightcontrolRelease** | **string** |  | 
 **garageId** | **string** |  | 
 **gcBranch** | **string** |  | 
 **gpuAlias** | **string** |  | 
 **gpuCount** | **int32** |  | 
 **gpuModel** | **string** |  | 
 **gpuOsName** | **string** |  | 
 **gpuOsRelease** | **string** |  | 
 **gpuOsVersion** | **string** |  | 
 **id** | **string** |  | 
 **nodeCount** | **int32** |  | 
 **oemName** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **orgName** | **string** |  | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **persistOnFailure** | **bool** |  | 
 **persona** | **string** |  | 
 **pipeline** | **int32** |  | 
 **pipelineBranch** | **string** |  | 
 **platform** | **string** |  | 
 **priority** | **string** | Priority level for the request  * &#x60;p0&#x60; - p0 * &#x60;p1&#x60; - p1 * &#x60;p2&#x60; - p2 * &#x60;p3&#x60; - p3 | 
 **providerName** | **string** |  | 
 **region** | **string** |  | 
 **requestId** | **string** |  | 
 **requesterEmail** | **string** |  | 
 **requesterName** | **string** |  | 
 **runtimeBranch** | **string** |  | 
 **runtimeCnsAddonPack** | **bool** |  | 
 **runtimeCnsDocker** | **bool** |  | 
 **runtimeCnsDriverVersion** | **string** |  | 
 **runtimeCnsK8s** | **bool** |  | 
 **runtimeCnsNvidiaDriver** | **bool** |  | 
 **runtimeCnsVersion** | **string** |  | 
 **runtimeMig** | **bool** |  | 
 **runtimeMigProfile** | **string** |  | 
 **salesId** | **string** |  | 
 **salesOwnerEmail** | **string** |  | 
 **salesOwnerName** | **string** |  | 
 **search** | **string** | Search for deployments by experience catalog_id, experience catalog_id_alias, experience id, experience title, expires_at, id, org_name, provisioning_config bastion_operating_system, provisioning_config collection_branch, provisioning_config experience_branch, provisioning_config flightcontrol_release, provisioning_config garage_id, provisioning_config gc_branch, provisioning_config gpu_alias, provisioning_config gpu_model, provisioning_config gpu_os_name, provisioning_config gpu_os_release, provisioning_config gpu_os_version, provisioning_config oem_name, provisioning_config persona, provisioning_config pipeline_branch, provisioning_config platform, provisioning_config provider_name, provisioning_config region, provisioning_config runtime_branch, provisioning_config runtime_cns_driver_version, provisioning_config runtime_cns_version, provisioning_config runtime_mig_profile, provisioning_config runtime_url, provisioning_config workshop_id, provisioning_request bastion_operating_system, provisioning_request collection_branch, provisioning_request experience_branch, provisioning_request flightcontrol_release, provisioning_request garage_id, provisioning_request gc_branch, provisioning_request gpu_alias, provisioning_request gpu_model, provisioning_request gpu_os_name, provisioning_request gpu_os_release, provisioning_request gpu_os_version, provisioning_request oem_name, provisioning_request persona, provisioning_request pipeline_branch, provisioning_request platform, provisioning_request provider_name, provisioning_request region, provisioning_request runtime_branch, provisioning_request runtime_cns_driver_version, provisioning_request runtime_cns_version, provisioning_request runtime_mig_profile, provisioning_request runtime_url, provisioning_request workshop_id, request_id, requester_email, requester_name, sales_id, sales_owner_email, sales_owner_name, services url, state, tags | 
 **state** | **[]string** | Multiple values may be separated by commas.  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | 
 **stateNot** | **[]string** | Multiple values may be separated by commas.  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | 
 **workshop** | **bool** |  | 
 **workshopId** | **string** |  | 

### Return type

[**PaginatedDeploymentList**](PaginatedDeploymentList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsNotesCreate

> DeploymentNote V1CatalogDeploymentsNotesCreate(ctx, deploymentId).DeploymentNote(deploymentNote).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesCreate(context.Background(), deploymentId).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsNotesCreate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentNote** | [**DeploymentNote**](DeploymentNote.md) |  | 

### Return type

[**DeploymentNote**](DeploymentNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsNotesDestroy

> V1CatalogDeploymentsNotesDestroy(ctx, deploymentId, id).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesDestroy(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesDestroyRequest struct via the builder pattern


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


## V1CatalogDeploymentsNotesList

> PaginatedDeploymentNoteList V1CatalogDeploymentsNotesList(ctx, deploymentId).CreatedBy(createdBy).Deployment(deployment).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createdBy := int32(56) // int32 |  (optional)
	deployment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	modifiedBy := int32(56) // int32 |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for deployment-notes by content (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesList(context.Background(), deploymentId).CreatedBy(createdBy).Deployment(deployment).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsNotesList`: PaginatedDeploymentNoteList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBy** | **int32** |  | 
 **deployment** | **string** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **modifiedBy** | **int32** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for deployment-notes by content | 

### Return type

[**PaginatedDeploymentNoteList**](PaginatedDeploymentNoteList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsNotesPartialUpdate

> DeploymentNote V1CatalogDeploymentsNotesPartialUpdate(ctx, deploymentId, id).DeploymentNote(deploymentNote).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesPartialUpdate(context.Background(), deploymentId, id).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsNotesPartialUpdate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentNote** | [**DeploymentNote**](DeploymentNote.md) |  | 

### Return type

[**DeploymentNote**](DeploymentNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsNotesRetrieve

> DeploymentNote V1CatalogDeploymentsNotesRetrieve(ctx, deploymentId, id).Fields(fields).Omit(omit).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesRetrieve(context.Background(), deploymentId, id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsNotesRetrieve`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**DeploymentNote**](DeploymentNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsNotesUpdate

> DeploymentNote V1CatalogDeploymentsNotesUpdate(ctx, deploymentId, id).DeploymentNote(deploymentNote).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsNotesUpdate(context.Background(), deploymentId, id).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsNotesUpdate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsNotesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsNotesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentNote** | [**DeploymentNote**](DeploymentNote.md) |  | 

### Return type

[**DeploymentNote**](DeploymentNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsPartialUpdate

> DeploymentUpdate V1CatalogDeploymentsPartialUpdate(ctx, id).DeploymentUpdate(deploymentUpdate).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this deployment.
	deploymentUpdate := *openapiclient.NewDeploymentUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), "Experience_example", "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example") // DeploymentUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsPartialUpdate(context.Background(), id).DeploymentUpdate(deploymentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsPartialUpdate`: DeploymentUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentUpdate** | [**DeploymentUpdate**](DeploymentUpdate.md) |  | 

### Return type

[**DeploymentUpdate**](DeploymentUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsPipelinesCreate

> DeploymentPipeline V1CatalogDeploymentsPipelinesCreate(ctx, deploymentId).DeploymentPipeline(deploymentPipeline).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentPipeline := *openapiclient.NewDeploymentPipeline(openapiclient.PipelineAction("apply"), time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), int64(123), "Url_example") // DeploymentPipeline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesCreate(context.Background(), deploymentId).DeploymentPipeline(deploymentPipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsPipelinesCreate`: DeploymentPipeline
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsPipelinesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentPipeline** | [**DeploymentPipeline**](DeploymentPipeline.md) |  | 

### Return type

[**DeploymentPipeline**](DeploymentPipeline.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsPipelinesList

> PaginatedDeploymentPipelineList V1CatalogDeploymentsPipelinesList(ctx, deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).PipelineId(pipelineId).Search(search).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	action := "action_example" // string | Action for the pipeline to run  * `apply` - apply * `destroy` - destroy * `notify` - notify (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	pipelineId := int32(56) // int32 |  (optional)
	search := "search_example" // string | Search for deployment-pipelines by action, id, pipeline_id, url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesList(context.Background(), deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).PipelineId(pipelineId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsPipelinesList`: PaginatedDeploymentPipelineList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsPipelinesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsPipelinesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy * &#x60;notify&#x60; - notify | 
 **id** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **pipelineId** | **int32** |  | 
 **search** | **string** | Search for deployment-pipelines by action, id, pipeline_id, url | 

### Return type

[**PaginatedDeploymentPipelineList**](PaginatedDeploymentPipelineList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsRetrieve

> Deployment V1CatalogDeploymentsRetrieve(ctx, id).Expand(expand).Expiring(expiring).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this deployment.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)
	expiring := "expiring_example" // string | Include deployments whose expires_at value is within the given range (inclusive), specified as \"today\", \"tomorrow\", or \"{start}[,end]\". Start and end times must be in ISO format. (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsRetrieve(context.Background(), id).Expand(expand).Expiring(expiring).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsRetrieve`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;cluster\&quot;). | 
 **expiring** | **string** | Include deployments whose expires_at value is within the given range (inclusive), specified as \&quot;today\&quot;, \&quot;tomorrow\&quot;, or \&quot;{start}[,end]\&quot;. Start and end times must be in ISO format. | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsServicesCreate

> DeploymentService V1CatalogDeploymentsServicesCreate(ctx, deploymentId).DeploymentService(deploymentService).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentService := *openapiclient.NewDeploymentService(time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "Name_example", "Url_example") // DeploymentService | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsServicesCreate(context.Background(), deploymentId).DeploymentService(deploymentService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsServicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsServicesCreate`: DeploymentService
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsServicesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsServicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentService** | [**DeploymentService**](DeploymentService.md) |  | 

### Return type

[**DeploymentService**](DeploymentService.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsServicesList

> PaginatedDeploymentServiceList V1CatalogDeploymentsServicesList(ctx, deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for deployment-services by id, name, url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsServicesList(context.Background(), deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsServicesList`: PaginatedDeploymentServiceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsServicesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsServicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for deployment-services by id, name, url | 

### Return type

[**PaginatedDeploymentServiceList**](PaginatedDeploymentServiceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsSshKeysCreate

> DeploymentKey V1CatalogDeploymentsSshKeysCreate(ctx, deploymentId).DeploymentKey(deploymentKey).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentKey := *openapiclient.NewDeploymentKey(time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "Name_example", "PublicKey_example") // DeploymentKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysCreate(context.Background(), deploymentId).DeploymentKey(deploymentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsSshKeysCreate`: DeploymentKey
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsSshKeysCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentKey** | [**DeploymentKey**](DeploymentKey.md) |  | 

### Return type

[**DeploymentKey**](DeploymentKey.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsSshKeysDestroy

> V1CatalogDeploymentsSshKeysDestroy(ctx, deploymentId, id).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysDestroy(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsSshKeysDestroyRequest struct via the builder pattern


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


## V1CatalogDeploymentsSshKeysList

> PaginatedDeploymentKeyList V1CatalogDeploymentsSshKeysList(ctx, deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for deployment-keys by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysList(context.Background(), deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsSshKeysList`: PaginatedDeploymentKeyList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsSshKeysList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsSshKeysListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for deployment-keys by name | 

### Return type

[**PaginatedDeploymentKeyList**](PaginatedDeploymentKeyList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsStatsRetrieve

> V1CatalogDeploymentsStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsStatsRetrieveRequest struct via the builder pattern


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


## V1CatalogDeploymentsTasksCreate

> DeploymentTask V1CatalogDeploymentsTasksCreate(ctx, deploymentId).DeploymentTask(deploymentTask).Execute()

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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentTask := *openapiclient.NewDeploymentTask(openapiclient.DeploymentTaskActionEnum("start_instances"), time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), int32(123), openapiclient.StatusEnum("completed"), "StatusText_example") // DeploymentTask | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsTasksCreate(context.Background(), deploymentId).DeploymentTask(deploymentTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsTasksCreate`: DeploymentTask
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsTasksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentTask** | [**DeploymentTask**](DeploymentTask.md) |  | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsTasksList

> PaginatedDeploymentTaskList V1CatalogDeploymentsTasksList(ctx, deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()

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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	action := "action_example" // string | The action the task will perform  * `start_instances` - Start all instances in the deployment * `stop_instances` - Stop all instances in the deployment (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for deployment-tasks by action, id, status, status_text (optional)
	status := "status_example" // string | Current status of the task  * `completed` - completed * `failed` - failed * `pending` - pending * `processing` - processing * `retrying` - retrying (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsTasksList(context.Background(), deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsTasksList`: PaginatedDeploymentTaskList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | The action the task will perform  * &#x60;start_instances&#x60; - Start all instances in the deployment * &#x60;stop_instances&#x60; - Stop all instances in the deployment | 
 **id** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for deployment-tasks by action, id, status, status_text | 
 **status** | **string** | Current status of the task  * &#x60;completed&#x60; - completed * &#x60;failed&#x60; - failed * &#x60;pending&#x60; - pending * &#x60;processing&#x60; - processing * &#x60;retrying&#x60; - retrying | 

### Return type

[**PaginatedDeploymentTaskList**](PaginatedDeploymentTaskList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsTasksRetrieve

> DeploymentTask V1CatalogDeploymentsTasksRetrieve(ctx, deploymentId, id).Execute()

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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsTasksRetrieve(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsTasksRetrieve`: DeploymentTask
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogDeploymentsUpdate

> DeploymentUpdate V1CatalogDeploymentsUpdate(ctx, id).DeploymentUpdate(deploymentUpdate).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this deployment.
	deploymentUpdate := *openapiclient.NewDeploymentUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), "Experience_example", "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", *openapiclient.NewOverrides(), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example") // DeploymentUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogDeploymentsUpdate(context.Background(), id).DeploymentUpdate(deploymentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogDeploymentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogDeploymentsUpdate`: DeploymentUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogDeploymentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogDeploymentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentUpdate** | [**DeploymentUpdate**](DeploymentUpdate.md) |  | 

### Return type

[**DeploymentUpdate**](DeploymentUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesNotesCreate

> ExperienceNote V1CatalogExperiencesNotesCreate(ctx, experienceId).ExperienceNote(experienceNote).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesCreate(context.Background(), experienceId).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesCreate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogExperiencesNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experienceNote** | [**ExperienceNote**](ExperienceNote.md) |  | 

### Return type

[**ExperienceNote**](ExperienceNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesNotesDestroy

> V1CatalogExperiencesNotesDestroy(ctx, experienceId, id).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesDestroy(context.Background(), experienceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesDestroyRequest struct via the builder pattern


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


## V1CatalogExperiencesNotesList

> PaginatedExperienceNoteList V1CatalogExperiencesNotesList(ctx, experienceId).CreatedBy(createdBy).Experience(experience).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createdBy := int32(56) // int32 |  (optional)
	experience := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	modifiedBy := int32(56) // int32 |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for experience-notes by content (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesList(context.Background(), experienceId).CreatedBy(createdBy).Experience(experience).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesList`: PaginatedExperienceNoteList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogExperiencesNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBy** | **int32** |  | 
 **experience** | **string** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **id** | **string** |  | 
 **modifiedBy** | **int32** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for experience-notes by content | 

### Return type

[**PaginatedExperienceNoteList**](PaginatedExperienceNoteList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesNotesPartialUpdate

> ExperienceNote V1CatalogExperiencesNotesPartialUpdate(ctx, experienceId, id).ExperienceNote(experienceNote).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesPartialUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesPartialUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogExperiencesNotesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experienceNote** | [**ExperienceNote**](ExperienceNote.md) |  | 

### Return type

[**ExperienceNote**](ExperienceNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesNotesRetrieve

> ExperienceNote V1CatalogExperiencesNotesRetrieve(ctx, experienceId, id).Fields(fields).Omit(omit).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesRetrieve(context.Background(), experienceId, id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesRetrieve`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogExperiencesNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**ExperienceNote**](ExperienceNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesNotesUpdate

> ExperienceNote V1CatalogExperiencesNotesUpdate(ctx, experienceId, id).ExperienceNote(experienceNote).Execute()



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
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.V1CatalogExperiencesNotesUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.V1CatalogExperiencesNotesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.V1CatalogExperiencesNotesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesNotesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experienceNote** | [**ExperienceNote**](ExperienceNote.md) |  | 

### Return type

[**ExperienceNote**](ExperienceNote.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

