# \CatalogDeploymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogDeploymentsBulkPartialUpdate**](CatalogDeploymentsAPI.md#CatalogDeploymentsBulkPartialUpdate) | **Patch** /v1/catalog/deployments/bulk/ | 
[**CatalogDeploymentsCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsCreate) | **Post** /v1/catalog/deployments/ | 
[**CatalogDeploymentsDestroy**](CatalogDeploymentsAPI.md#CatalogDeploymentsDestroy) | **Delete** /v1/catalog/deployments/{id}/ | 
[**CatalogDeploymentsHistoryList**](CatalogDeploymentsAPI.md#CatalogDeploymentsHistoryList) | **Get** /v1/catalog/deployments/{id}/history/ | 
[**CatalogDeploymentsInstancesList**](CatalogDeploymentsAPI.md#CatalogDeploymentsInstancesList) | **Get** /v1/catalog/deployments/{deployment_id}/instances/ | 🚧 [Beta Feature]
[**CatalogDeploymentsList**](CatalogDeploymentsAPI.md#CatalogDeploymentsList) | **Get** /v1/catalog/deployments/ | 
[**CatalogDeploymentsNotesCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/notes/ | 
[**CatalogDeploymentsNotesDestroy**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesDestroy) | **Delete** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**CatalogDeploymentsNotesList**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesList) | **Get** /v1/catalog/deployments/{deployment_id}/notes/ | 
[**CatalogDeploymentsNotesPartialUpdate**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesPartialUpdate) | **Patch** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**CatalogDeploymentsNotesRetrieve**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesRetrieve) | **Get** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**CatalogDeploymentsNotesUpdate**](CatalogDeploymentsAPI.md#CatalogDeploymentsNotesUpdate) | **Put** /v1/catalog/deployments/{deployment_id}/notes/{id}/ | 
[**CatalogDeploymentsPartialUpdate**](CatalogDeploymentsAPI.md#CatalogDeploymentsPartialUpdate) | **Patch** /v1/catalog/deployments/{id}/ | 
[**CatalogDeploymentsPipelinesCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsPipelinesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/pipelines/ | 
[**CatalogDeploymentsPipelinesList**](CatalogDeploymentsAPI.md#CatalogDeploymentsPipelinesList) | **Get** /v1/catalog/deployments/{deployment_id}/pipelines/ | 
[**CatalogDeploymentsRetrieve**](CatalogDeploymentsAPI.md#CatalogDeploymentsRetrieve) | **Get** /v1/catalog/deployments/{id}/ | 
[**CatalogDeploymentsServicesCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsServicesCreate) | **Post** /v1/catalog/deployments/{deployment_id}/services/ | 
[**CatalogDeploymentsServicesList**](CatalogDeploymentsAPI.md#CatalogDeploymentsServicesList) | **Get** /v1/catalog/deployments/{deployment_id}/services/ | 
[**CatalogDeploymentsSshKeysCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsSshKeysCreate) | **Post** /v1/catalog/deployments/{deployment_id}/ssh-keys/ | 
[**CatalogDeploymentsSshKeysDestroy**](CatalogDeploymentsAPI.md#CatalogDeploymentsSshKeysDestroy) | **Delete** /v1/catalog/deployments/{deployment_id}/ssh-keys/{id}/ | 
[**CatalogDeploymentsSshKeysList**](CatalogDeploymentsAPI.md#CatalogDeploymentsSshKeysList) | **Get** /v1/catalog/deployments/{deployment_id}/ssh-keys/ | 
[**CatalogDeploymentsStatsRetrieve**](CatalogDeploymentsAPI.md#CatalogDeploymentsStatsRetrieve) | **Get** /v1/catalog/deployments/stats/ | 🚧 [Beta Feature]
[**CatalogDeploymentsTasksCreate**](CatalogDeploymentsAPI.md#CatalogDeploymentsTasksCreate) | **Post** /v1/catalog/deployments/{deployment_id}/tasks/ | 🚧 [Beta Feature]
[**CatalogDeploymentsTasksList**](CatalogDeploymentsAPI.md#CatalogDeploymentsTasksList) | **Get** /v1/catalog/deployments/{deployment_id}/tasks/ | 🚧 [Beta Feature]
[**CatalogDeploymentsTasksRetrieve**](CatalogDeploymentsAPI.md#CatalogDeploymentsTasksRetrieve) | **Get** /v1/catalog/deployments/{deployment_id}/tasks/{id}/ | 🚧 [Beta Feature]
[**CatalogDeploymentsUpdate**](CatalogDeploymentsAPI.md#CatalogDeploymentsUpdate) | **Put** /v1/catalog/deployments/{id}/ | 
[**CatalogExperiencesNotesCreate**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesCreate) | **Post** /v1/catalog/experiences/{experience_id}/notes/ | 
[**CatalogExperiencesNotesDestroy**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesDestroy) | **Delete** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**CatalogExperiencesNotesList**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesList) | **Get** /v1/catalog/experiences/{experience_id}/notes/ | 
[**CatalogExperiencesNotesPartialUpdate**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesPartialUpdate) | **Patch** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**CatalogExperiencesNotesRetrieve**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesRetrieve) | **Get** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**CatalogExperiencesNotesUpdate**](CatalogDeploymentsAPI.md#CatalogExperiencesNotesUpdate) | **Put** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 



## CatalogDeploymentsBulkPartialUpdate

> DeploymentBulkUpdate CatalogDeploymentsBulkPartialUpdate(ctx).DeploymentBulkUpdate(deploymentBulkUpdate).Execute()



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
	deploymentBulkUpdate := *openapiclient.NewDeploymentBulkUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", interface{}(123), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example", int32(123), []string{"Ids_example"}, "Result_example") // DeploymentBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsBulkPartialUpdate(context.Background()).DeploymentBulkUpdate(deploymentBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsBulkPartialUpdate`: DeploymentBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsBulkPartialUpdateRequest struct via the builder pattern


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


## CatalogDeploymentsCreate

> Deployment CatalogDeploymentsCreate(ctx).Deployment(deployment).Expand(expand).Execute()



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
	deployment := *openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed")) // Deployment | 
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"cluster\"). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsCreate(context.Background()).Deployment(deployment).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsCreate`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsCreateRequest struct via the builder pattern


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


## CatalogDeploymentsDestroy

> DocDeploymentDelete CatalogDeploymentsDestroy(ctx, id).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsDestroy`: DocDeploymentDelete
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsDestroyRequest struct via the builder pattern


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


## CatalogDeploymentsHistoryList

> PaginatedModelChangeList CatalogDeploymentsHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsHistoryListRequest struct via the builder pattern


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


## CatalogDeploymentsInstancesList

> PaginatedDeploymentInstanceList CatalogDeploymentsInstancesList(ctx, deploymentId).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()

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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsInstancesList(context.Background(), deploymentId).Fields(fields).Id(id).InstanceId(instanceId).Name(name).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsInstancesList`: PaginatedDeploymentInstanceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsInstancesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsInstancesListRequest struct via the builder pattern


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


## CatalogDeploymentsList

> PaginatedDeploymentList CatalogDeploymentsList(ctx).BastionOperatingSystem(bastionOperatingSystem).Cluster(cluster).ClusterGpusModel(clusterGpusModel).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Expired(expired).ExpiresAt(expiresAt).Expiring(expiring).Fields(fields).FlightcontrolRelease(flightcontrolRelease).GarageId(garageId).GcBranch(gcBranch).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).OemName(oemName).Omit(omit).Ordering(ordering).OrgName(orgName).Page(page).PageSize(pageSize).PersistOnFailure(persistOnFailure).Persona(persona).Pipeline(pipeline).PipelineBranch(pipelineBranch).Platform(platform).Priority(priority).ProviderName(providerName).Region(region).RequestId(requestId).RequesterEmail(requesterEmail).RequesterName(requesterName).SalesId(salesId).SalesOwnerEmail(salesOwnerEmail).SalesOwnerName(salesOwnerName).Search(search).State(state).Workshop(workshop).WorkshopId(workshopId).Execute()



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
	platform := "platform_example" // string | Override the default platform selection  * `air` - NVIDIA Air * `flight_deck` - Flight Deck * `kvm_bastion` - KVM Bastion * `lp-vmware-platform` - lp-vmware-platform * `minimal` - minimal * `openshift` - OpenShift * `vsphere` - vSphere * `vsphere_horizon` - VMware Horizon * `vsphere7` - vSphere 7 * `vsphere8` - vSphere 8 (optional)
	priority := "priority_example" // string | Priority level for the request  * `p0` - p0 * `p1` - p1 * `p2` - p2 * `p3` - p3 (optional)
	providerName := "providerName_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	requestId := "requestId_example" // string |  (optional)
	requesterEmail := "requesterEmail_example" // string |  (optional)
	requesterName := "requesterName_example" // string |  (optional)
	salesId := "salesId_example" // string |  (optional)
	salesOwnerEmail := "salesOwnerEmail_example" // string |  (optional)
	salesOwnerName := "salesOwnerName_example" // string |  (optional)
	search := "search_example" // string | Search for deployments by bastion_operating_system, collection_branch, experience_branch, experience catalog_id, experience catalog_id_alias, experience id, experience title, expires_at, flightcontrol_release, garage_id, gc_branch, gpu_alias, gpu_model, gpu_os_name, gpu_os_release, gpu_os_version, id, oem_name, org_name, persona, pipeline_branch, platform, provider_name, region, request_id, requester_email, requester_name, sales_id, sales_owner_email, sales_owner_name, services url, state, tags, workshop_id (optional)
	state := []string{"State_example"} // []string | Multiple values may be separated by commas.  * `destroyed` - Deployment has been fully destroyed * `destroying` - Deployment is being destroyed * `error` - Deployment has encountered a fatal error and will not be retried * `failed` - Deployment has failed but may be retried * `paused` - Deployment is paused but may be retried later * `ready` - Deployment is ready and all instances are running * `retrying` - Deployment is retrying * `starting` - Deployment instances are starting * `stopped` - Deployment instances are stopped * `stopping` - Deployment instances are stopping * `waiting` - Waiting for deployment to be ready (optional)
	workshop := true // bool |  (optional)
	workshopId := "workshopId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsList(context.Background()).BastionOperatingSystem(bastionOperatingSystem).Cluster(cluster).ClusterGpusModel(clusterGpusModel).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Expired(expired).ExpiresAt(expiresAt).Expiring(expiring).Fields(fields).FlightcontrolRelease(flightcontrolRelease).GarageId(garageId).GcBranch(gcBranch).GpuAlias(gpuAlias).GpuCount(gpuCount).GpuModel(gpuModel).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).OemName(oemName).Omit(omit).Ordering(ordering).OrgName(orgName).Page(page).PageSize(pageSize).PersistOnFailure(persistOnFailure).Persona(persona).Pipeline(pipeline).PipelineBranch(pipelineBranch).Platform(platform).Priority(priority).ProviderName(providerName).Region(region).RequestId(requestId).RequesterEmail(requesterEmail).RequesterName(requesterName).SalesId(salesId).SalesOwnerEmail(salesOwnerEmail).SalesOwnerName(salesOwnerName).Search(search).State(state).Workshop(workshop).WorkshopId(workshopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsList`: PaginatedDeploymentList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsListRequest struct via the builder pattern


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
 **platform** | **string** | Override the default platform selection  * &#x60;air&#x60; - NVIDIA Air * &#x60;flight_deck&#x60; - Flight Deck * &#x60;kvm_bastion&#x60; - KVM Bastion * &#x60;lp-vmware-platform&#x60; - lp-vmware-platform * &#x60;minimal&#x60; - minimal * &#x60;openshift&#x60; - OpenShift * &#x60;vsphere&#x60; - vSphere * &#x60;vsphere_horizon&#x60; - VMware Horizon * &#x60;vsphere7&#x60; - vSphere 7 * &#x60;vsphere8&#x60; - vSphere 8 | 
 **priority** | **string** | Priority level for the request  * &#x60;p0&#x60; - p0 * &#x60;p1&#x60; - p1 * &#x60;p2&#x60; - p2 * &#x60;p3&#x60; - p3 | 
 **providerName** | **string** |  | 
 **region** | **string** |  | 
 **requestId** | **string** |  | 
 **requesterEmail** | **string** |  | 
 **requesterName** | **string** |  | 
 **salesId** | **string** |  | 
 **salesOwnerEmail** | **string** |  | 
 **salesOwnerName** | **string** |  | 
 **search** | **string** | Search for deployments by bastion_operating_system, collection_branch, experience_branch, experience catalog_id, experience catalog_id_alias, experience id, experience title, expires_at, flightcontrol_release, garage_id, gc_branch, gpu_alias, gpu_model, gpu_os_name, gpu_os_release, gpu_os_version, id, oem_name, org_name, persona, pipeline_branch, platform, provider_name, region, request_id, requester_email, requester_name, sales_id, sales_owner_email, sales_owner_name, services url, state, tags, workshop_id | 
 **state** | **[]string** | Multiple values may be separated by commas.  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | 
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


## CatalogDeploymentsNotesCreate

> DeploymentNote CatalogDeploymentsNotesCreate(ctx, deploymentId).DeploymentNote(deploymentNote).Execute()



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
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesCreate(context.Background(), deploymentId).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsNotesCreate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesCreateRequest struct via the builder pattern


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


## CatalogDeploymentsNotesDestroy

> CatalogDeploymentsNotesDestroy(ctx, deploymentId, id).Execute()



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
	r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesDestroy(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesDestroyRequest struct via the builder pattern


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


## CatalogDeploymentsNotesList

> PaginatedDeploymentNoteList CatalogDeploymentsNotesList(ctx, deploymentId).CreatedBy(createdBy).Deployment(deployment).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesList(context.Background(), deploymentId).CreatedBy(createdBy).Deployment(deployment).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsNotesList`: PaginatedDeploymentNoteList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesListRequest struct via the builder pattern


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


## CatalogDeploymentsNotesPartialUpdate

> DeploymentNote CatalogDeploymentsNotesPartialUpdate(ctx, deploymentId, id).DeploymentNote(deploymentNote).Execute()



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
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesPartialUpdate(context.Background(), deploymentId, id).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsNotesPartialUpdate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsNotesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesPartialUpdateRequest struct via the builder pattern


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


## CatalogDeploymentsNotesRetrieve

> DeploymentNote CatalogDeploymentsNotesRetrieve(ctx, deploymentId, id).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesRetrieve(context.Background(), deploymentId, id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsNotesRetrieve`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesRetrieveRequest struct via the builder pattern


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


## CatalogDeploymentsNotesUpdate

> DeploymentNote CatalogDeploymentsNotesUpdate(ctx, deploymentId, id).DeploymentNote(deploymentNote).Execute()



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
	deploymentNote := *openapiclient.NewDeploymentNote(time.Now(), "CreatedBy_example", openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "ModifiedBy_example") // DeploymentNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsNotesUpdate(context.Background(), deploymentId, id).DeploymentNote(deploymentNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsNotesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsNotesUpdate`: DeploymentNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsNotesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsNotesUpdateRequest struct via the builder pattern


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


## CatalogDeploymentsPartialUpdate

> DeploymentUpdate CatalogDeploymentsPartialUpdate(ctx, id).DeploymentUpdate(deploymentUpdate).Execute()



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
	deploymentUpdate := *openapiclient.NewDeploymentUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", interface{}(123), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example") // DeploymentUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsPartialUpdate(context.Background(), id).DeploymentUpdate(deploymentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsPartialUpdate`: DeploymentUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsPartialUpdateRequest struct via the builder pattern


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


## CatalogDeploymentsPipelinesCreate

> DeploymentPipeline CatalogDeploymentsPipelinesCreate(ctx, deploymentId).DeploymentPipeline(deploymentPipeline).Execute()



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
	deploymentPipeline := *openapiclient.NewDeploymentPipeline(openapiclient.DeploymentPipelineActionEnum("apply"), time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), int64(123), "Url_example") // DeploymentPipeline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsPipelinesCreate(context.Background(), deploymentId).DeploymentPipeline(deploymentPipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsPipelinesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsPipelinesCreate`: DeploymentPipeline
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsPipelinesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsPipelinesCreateRequest struct via the builder pattern


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


## CatalogDeploymentsPipelinesList

> PaginatedDeploymentPipelineList CatalogDeploymentsPipelinesList(ctx, deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).PipelineId(pipelineId).Search(search).Execute()



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
	action := "action_example" // string | Action for the pipeline to run  * `apply` - apply * `destroy` - destroy (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	pipelineId := int32(56) // int32 |  (optional)
	search := "search_example" // string | Search for deployment-pipelines by action, id, pipeline_id, url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsPipelinesList(context.Background(), deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).PipelineId(pipelineId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsPipelinesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsPipelinesList`: PaginatedDeploymentPipelineList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsPipelinesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsPipelinesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy | 
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


## CatalogDeploymentsRetrieve

> Deployment CatalogDeploymentsRetrieve(ctx, id).Expand(expand).Expiring(expiring).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsRetrieve(context.Background(), id).Expand(expand).Expiring(expiring).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsRetrieve`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsRetrieveRequest struct via the builder pattern


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


## CatalogDeploymentsServicesCreate

> DeploymentService CatalogDeploymentsServicesCreate(ctx, deploymentId).DeploymentService(deploymentService).Execute()



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
	deploymentService := *openapiclient.NewDeploymentService(time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "Name_example", "Url_example") // DeploymentService | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsServicesCreate(context.Background(), deploymentId).DeploymentService(deploymentService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsServicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsServicesCreate`: DeploymentService
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsServicesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsServicesCreateRequest struct via the builder pattern


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


## CatalogDeploymentsServicesList

> PaginatedDeploymentServiceList CatalogDeploymentsServicesList(ctx, deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsServicesList(context.Background(), deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsServicesList`: PaginatedDeploymentServiceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsServicesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsServicesListRequest struct via the builder pattern


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


## CatalogDeploymentsSshKeysCreate

> DeploymentKey CatalogDeploymentsSshKeysCreate(ctx, deploymentId).DeploymentKey(deploymentKey).Execute()



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
	deploymentKey := *openapiclient.NewDeploymentKey(time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), "Name_example", "PublicKey_example") // DeploymentKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsSshKeysCreate(context.Background(), deploymentId).DeploymentKey(deploymentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsSshKeysCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsSshKeysCreate`: DeploymentKey
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsSshKeysCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsSshKeysCreateRequest struct via the builder pattern


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


## CatalogDeploymentsSshKeysDestroy

> CatalogDeploymentsSshKeysDestroy(ctx, deploymentId, id).Execute()



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
	r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsSshKeysDestroy(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsSshKeysDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCatalogDeploymentsSshKeysDestroyRequest struct via the builder pattern


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


## CatalogDeploymentsSshKeysList

> PaginatedDeploymentKeyList CatalogDeploymentsSshKeysList(ctx, deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsSshKeysList(context.Background(), deploymentId).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsSshKeysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsSshKeysList`: PaginatedDeploymentKeyList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsSshKeysList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsSshKeysListRequest struct via the builder pattern


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


## CatalogDeploymentsStatsRetrieve

> CatalogDeploymentsStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsStatsRetrieveRequest struct via the builder pattern


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


## CatalogDeploymentsTasksCreate

> DeploymentTask CatalogDeploymentsTasksCreate(ctx, deploymentId).DeploymentTask(deploymentTask).Execute()

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
	deploymentTask := *openapiclient.NewDeploymentTask(openapiclient.DeploymentTaskActionEnum("start_instances"), time.Now(), openapiclient.Cluster_deployment{Deployment: openapiclient.NewDeployment(time.Now(), "Id_example", time.Now(), "OrgName_example", interface{}(123), []string{"Pipelines_example"}, "RequesterEmail_example", "RequesterName_example", int32(123), []string{"Services_example"}, int32(123), "SshUser_example", openapiclient.DeploymentState("destroyed"))}, "Id_example", time.Now(), int32(123), openapiclient.StatusEnum("completed"), "StatusText_example") // DeploymentTask | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsTasksCreate(context.Background(), deploymentId).DeploymentTask(deploymentTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsTasksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsTasksCreate`: DeploymentTask
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsTasksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsTasksCreateRequest struct via the builder pattern


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


## CatalogDeploymentsTasksList

> PaginatedDeploymentTaskList CatalogDeploymentsTasksList(ctx, deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()

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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsTasksList(context.Background(), deploymentId).Action(action).Id(id).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsTasksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsTasksList`: PaginatedDeploymentTaskList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsTasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsTasksListRequest struct via the builder pattern


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


## CatalogDeploymentsTasksRetrieve

> DeploymentTask CatalogDeploymentsTasksRetrieve(ctx, deploymentId, id).Execute()

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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsTasksRetrieve(context.Background(), deploymentId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsTasksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsTasksRetrieve`: DeploymentTask
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsTasksRetrieveRequest struct via the builder pattern


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


## CatalogDeploymentsUpdate

> DeploymentUpdate CatalogDeploymentsUpdate(ctx, id).DeploymentUpdate(deploymentUpdate).Execute()



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
	deploymentUpdate := *openapiclient.NewDeploymentUpdate("BastionOperatingSystem_example", "Cluster_example", time.Now(), openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "GarageId_example", "GpuAlias_example", NullableInt32(123), "GpuModel_example", "GpuOsName_example", "GpuOsRelease_example", "GpuOsVersion_example", "Id_example", NullableInt32(123), NullableInt32(123), time.Now(), NullableInt32(123), "OemName_example", interface{}(123), []string{"Pipelines_example"}, openapiclient.PriorityEnum("p0"), "ProviderName_example", "PublicKey_example", "Region_example", "RequestId_example", time.Now(), "SalesId_example", []string{"Services_example"}, false, "WorkshopId_example", "WorkshopOverridePassword_example") // DeploymentUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogDeploymentsUpdate(context.Background(), id).DeploymentUpdate(deploymentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogDeploymentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogDeploymentsUpdate`: DeploymentUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogDeploymentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogDeploymentsUpdateRequest struct via the builder pattern


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


## CatalogExperiencesNotesCreate

> ExperienceNote CatalogExperiencesNotesCreate(ctx, experienceId).ExperienceNote(experienceNote).Execute()



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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesCreate(context.Background(), experienceId).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogExperiencesNotesCreate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogExperiencesNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesCreateRequest struct via the builder pattern


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


## CatalogExperiencesNotesDestroy

> CatalogExperiencesNotesDestroy(ctx, experienceId, id).Execute()



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
	r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesDestroy(context.Background(), experienceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesDestroyRequest struct via the builder pattern


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


## CatalogExperiencesNotesList

> PaginatedExperienceNoteList CatalogExperiencesNotesList(ctx, experienceId).CreatedBy(createdBy).Experience(experience).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesList(context.Background(), experienceId).CreatedBy(createdBy).Experience(experience).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogExperiencesNotesList`: PaginatedExperienceNoteList
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogExperiencesNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesListRequest struct via the builder pattern


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


## CatalogExperiencesNotesPartialUpdate

> ExperienceNote CatalogExperiencesNotesPartialUpdate(ctx, experienceId, id).ExperienceNote(experienceNote).Execute()



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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesPartialUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogExperiencesNotesPartialUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogExperiencesNotesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesPartialUpdateRequest struct via the builder pattern


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


## CatalogExperiencesNotesRetrieve

> ExperienceNote CatalogExperiencesNotesRetrieve(ctx, experienceId, id).Fields(fields).Omit(omit).Execute()



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
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesRetrieve(context.Background(), experienceId, id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogExperiencesNotesRetrieve`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogExperiencesNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesRetrieveRequest struct via the builder pattern


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


## CatalogExperiencesNotesUpdate

> ExperienceNote CatalogExperiencesNotesUpdate(ctx, experienceId, id).ExperienceNote(experienceNote).Execute()



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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.Deployment_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogDeploymentsAPI.CatalogExperiencesNotesUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogDeploymentsAPI.CatalogExperiencesNotesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogExperiencesNotesUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogDeploymentsAPI.CatalogExperiencesNotesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogExperiencesNotesUpdateRequest struct via the builder pattern


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

