# \CatalogExperiencesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogExperiencesBulkCreate**](CatalogExperiencesAPI.md#V1CatalogExperiencesBulkCreate) | **Post** /v1/catalog/experiences/bulk/ | 
[**V1CatalogExperiencesBulkPartialUpdate**](CatalogExperiencesAPI.md#V1CatalogExperiencesBulkPartialUpdate) | **Patch** /v1/catalog/experiences/bulk/ | 
[**V1CatalogExperiencesCreate**](CatalogExperiencesAPI.md#V1CatalogExperiencesCreate) | **Post** /v1/catalog/experiences/ | 
[**V1CatalogExperiencesDestroy**](CatalogExperiencesAPI.md#V1CatalogExperiencesDestroy) | **Delete** /v1/catalog/experiences/{id}/ | 
[**V1CatalogExperiencesHistoryList**](CatalogExperiencesAPI.md#V1CatalogExperiencesHistoryList) | **Get** /v1/catalog/experiences/{id}/history/ | 
[**V1CatalogExperiencesList**](CatalogExperiencesAPI.md#V1CatalogExperiencesList) | **Get** /v1/catalog/experiences/ | 
[**V1CatalogExperiencesNotesCreate**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesCreate) | **Post** /v1/catalog/experiences/{experience_id}/notes/ | 
[**V1CatalogExperiencesNotesDestroy**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesDestroy) | **Delete** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesList**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesList) | **Get** /v1/catalog/experiences/{experience_id}/notes/ | 
[**V1CatalogExperiencesNotesPartialUpdate**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesPartialUpdate) | **Patch** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesRetrieve**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesRetrieve) | **Get** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesNotesUpdate**](CatalogExperiencesAPI.md#V1CatalogExperiencesNotesUpdate) | **Put** /v1/catalog/experiences/{experience_id}/notes/{id}/ | 
[**V1CatalogExperiencesPartialUpdate**](CatalogExperiencesAPI.md#V1CatalogExperiencesPartialUpdate) | **Patch** /v1/catalog/experiences/{id}/ | 
[**V1CatalogExperiencesRetrieve**](CatalogExperiencesAPI.md#V1CatalogExperiencesRetrieve) | **Get** /v1/catalog/experiences/{id}/ | 
[**V1CatalogExperiencesStatsRetrieve**](CatalogExperiencesAPI.md#V1CatalogExperiencesStatsRetrieve) | **Get** /v1/catalog/experiences/stats/ | 🚧 [Beta Feature]
[**V1CatalogExperiencesUpdate**](CatalogExperiencesAPI.md#V1CatalogExperiencesUpdate) | **Put** /v1/catalog/experiences/{id}/ | 



## V1CatalogExperiencesBulkCreate

> ExperienceBulk V1CatalogExperiencesBulkCreate(ctx).CsvFile(csvFile).Execute()





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
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesBulkCreate(context.Background()).CsvFile(csvFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesBulkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesBulkCreate`: ExperienceBulk
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesBulkCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csvFile** | ***os.File** |  | 

### Return type

[**ExperienceBulk**](ExperienceBulk.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesBulkPartialUpdate

> ExperienceBulkUpdate V1CatalogExperiencesBulkPartialUpdate(ctx).ExperienceBulkUpdate(experienceBulkUpdate).Execute()



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
	experienceBulkUpdate := *openapiclient.NewExperienceBulkUpdate("CatalogId_example", "CatalogIdAlias_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example", int32(123), []string{"Ids_example"}, "Result_example") // ExperienceBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesBulkPartialUpdate(context.Background()).ExperienceBulkUpdate(experienceBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesBulkPartialUpdate`: ExperienceBulkUpdate
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experienceBulkUpdate** | [**ExperienceBulkUpdate**](ExperienceBulkUpdate.md) |  | 

### Return type

[**ExperienceBulkUpdate**](ExperienceBulkUpdate.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesCreate

> Experience V1CatalogExperiencesCreate(ctx).Experience(experience).Execute()



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
	experience := *openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example") // Experience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesCreate(context.Background()).Experience(experience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesCreate`: Experience
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experience** | [**Experience**](Experience.md) |  | 

### Return type

[**Experience**](Experience.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesDestroy

> V1CatalogExperiencesDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this experience.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this experience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesDestroyRequest struct via the builder pattern


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


## V1CatalogExperiencesHistoryList

> PaginatedModelChangeList V1CatalogExperiencesHistoryList(ctx, id).Page(page).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesHistoryList(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesHistoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesHistoryList`: PaginatedModelChangeList
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesHistoryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesHistoryListRequest struct via the builder pattern


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


## V1CatalogExperiencesList

> PaginatedExperienceList V1CatalogExperiencesList(ctx).Assignee(assignee).Autoapprove(autoapprove).Autoprovision(autoprovision).Bootstrap(bootstrap).CatalogId(catalogId).CatalogIdAlias(catalogIdAlias).Category(category).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Featured(featured).Fields(fields).GarageId(garageId).GcBranch(gcBranch).GpuCount(gpuCount).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persona(persona).Pipeline(pipeline).Platform(platform).Provider(provider).Published(published).RequiresGpu(requiresGpu).Runtime(runtime).SaLab(saLab).Search(search).SystemArch(systemArch).VgpuProfile(vgpuProfile).Execute()



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
	assignee := "assignee_example" // string |  (optional)
	autoapprove := true // bool |  (optional)
	autoprovision := true // bool |  (optional)
	bootstrap := true // bool |  (optional)
	catalogId := "catalogId_example" // string |  (optional)
	catalogIdAlias := "catalogIdAlias_example" // string |  (optional)
	category := []string{"Category_example"} // []string | Multiple values may be separated by commas.  * `AI` - AI * `Clara` - Clara * `Data Science` - Data Science * `3D Design Collaboration and Simulation` - 3D Design Collaboration and Simulation * `Developer` - Developer * `Infrastructure Optimization` - Infrastructure Optimization (optional)
	collectionBranch := "collectionBranch_example" // string |  (optional)
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"required_gpus\"). (optional)
	experience := "experience_example" // string |  (optional)
	experienceBranch := "experienceBranch_example" // string |  (optional)
	featured := true // bool |  (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	garageId := "garageId_example" // string |  (optional)
	gcBranch := "gcBranch_example" // string |  (optional)
	gpuCount := int32(56) // int32 |  (optional)
	gpuOsName := "gpuOsName_example" // string |  (optional)
	gpuOsRelease := "gpuOsRelease_example" // string |  (optional)
	gpuOsVersion := "gpuOsVersion_example" // string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	nodeCount := int32(56) // int32 |  (optional)
	oem := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	persona := "persona_example" // string |  (optional)
	pipeline := int32(56) // int32 |  (optional)
	platform := "platform_example" // string | Base platform that the experience will be provisioned onto  * `air` - NVIDIA Air * `flight_deck` - Flight Deck * `kvm_bastion` - KVM Bastion * `lp-vmware-platform` - lp-vmware-platform * `minimal` - minimal * `openshift` - OpenShift * `proxmox` - Proxmox * `vsphere` - vSphere * `vsphere_horizon` - VMware Horizon * `vsphere7` - vSphere 7 * `vsphere8` - vSphere 8 (optional)
	provider := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	published := []string{"Published_example"} // []string | Multiple values may be separated by commas.  * `draft` - draft * `no` - no * `yes` - yes (optional)
	requiresGpu := "requiresGpu_example" // string | Only include experiences that require a given GPU ID or model (optional)
	runtime := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	saLab := true // bool |  (optional)
	search := "search_example" // string | Search for experiences by assignee, catalog_id, catalog_id_alias, category, collection_branch, description, experience, experience_branch, gc_branch, gpu_os_name, gpu_os_release, gpu_os_version, id, oem name, persona, pipeline, platform, provider name, required_gpus model, runtime name, system_arch, title, vgpu_profile (optional)
	systemArch := "systemArch_example" // string | Required CPU architecture  * `amd64` - amd64 * `arm64` - arm64 (optional)
	vgpuProfile := "vgpuProfile_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesList(context.Background()).Assignee(assignee).Autoapprove(autoapprove).Autoprovision(autoprovision).Bootstrap(bootstrap).CatalogId(catalogId).CatalogIdAlias(catalogIdAlias).Category(category).CollectionBranch(collectionBranch).Expand(expand).Experience(experience).ExperienceBranch(experienceBranch).Featured(featured).Fields(fields).GarageId(garageId).GcBranch(gcBranch).GpuCount(gpuCount).GpuOsName(gpuOsName).GpuOsRelease(gpuOsRelease).GpuOsVersion(gpuOsVersion).Id(id).NodeCount(nodeCount).Oem(oem).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Persona(persona).Pipeline(pipeline).Platform(platform).Provider(provider).Published(published).RequiresGpu(requiresGpu).Runtime(runtime).SaLab(saLab).Search(search).SystemArch(systemArch).VgpuProfile(vgpuProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesList`: PaginatedExperienceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignee** | **string** |  | 
 **autoapprove** | **bool** |  | 
 **autoprovision** | **bool** |  | 
 **bootstrap** | **bool** |  | 
 **catalogId** | **string** |  | 
 **catalogIdAlias** | **string** |  | 
 **category** | **[]string** | Multiple values may be separated by commas.  * &#x60;AI&#x60; - AI * &#x60;Clara&#x60; - Clara * &#x60;Data Science&#x60; - Data Science * &#x60;3D Design Collaboration and Simulation&#x60; - 3D Design Collaboration and Simulation * &#x60;Developer&#x60; - Developer * &#x60;Infrastructure Optimization&#x60; - Infrastructure Optimization | 
 **collectionBranch** | **string** |  | 
 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;required_gpus\&quot;). | 
 **experience** | **string** |  | 
 **experienceBranch** | **string** |  | 
 **featured** | **bool** |  | 
 **fields** | **string** | Include only the specified fields in the response | 
 **garageId** | **string** |  | 
 **gcBranch** | **string** |  | 
 **gpuCount** | **int32** |  | 
 **gpuOsName** | **string** |  | 
 **gpuOsRelease** | **string** |  | 
 **gpuOsVersion** | **string** |  | 
 **id** | **string** |  | 
 **nodeCount** | **int32** |  | 
 **oem** | **string** |  | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **persona** | **string** |  | 
 **pipeline** | **int32** |  | 
 **platform** | **string** | Base platform that the experience will be provisioned onto  * &#x60;air&#x60; - NVIDIA Air * &#x60;flight_deck&#x60; - Flight Deck * &#x60;kvm_bastion&#x60; - KVM Bastion * &#x60;lp-vmware-platform&#x60; - lp-vmware-platform * &#x60;minimal&#x60; - minimal * &#x60;openshift&#x60; - OpenShift * &#x60;proxmox&#x60; - Proxmox * &#x60;vsphere&#x60; - vSphere * &#x60;vsphere_horizon&#x60; - VMware Horizon * &#x60;vsphere7&#x60; - vSphere 7 * &#x60;vsphere8&#x60; - vSphere 8 | 
 **provider** | **string** |  | 
 **published** | **[]string** | Multiple values may be separated by commas.  * &#x60;draft&#x60; - draft * &#x60;no&#x60; - no * &#x60;yes&#x60; - yes | 
 **requiresGpu** | **string** | Only include experiences that require a given GPU ID or model | 
 **runtime** | **string** |  | 
 **saLab** | **bool** |  | 
 **search** | **string** | Search for experiences by assignee, catalog_id, catalog_id_alias, category, collection_branch, description, experience, experience_branch, gc_branch, gpu_os_name, gpu_os_release, gpu_os_version, id, oem name, persona, pipeline, platform, provider name, required_gpus model, runtime name, system_arch, title, vgpu_profile | 
 **systemArch** | **string** | Required CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
 **vgpuProfile** | **string** |  | 

### Return type

[**PaginatedExperienceList**](PaginatedExperienceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesCreate(context.Background(), experienceId).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesCreate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesNotesCreate`: %v\n", resp)
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
	r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesDestroy(context.Background(), experienceId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesDestroy``: %v\n", err)
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
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesList(context.Background(), experienceId).CreatedBy(createdBy).Experience(experience).Fields(fields).Id(id).ModifiedBy(modifiedBy).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesList`: PaginatedExperienceNoteList
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesNotesList`: %v\n", resp)
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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesPartialUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesPartialUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesNotesPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesRetrieve(context.Background(), experienceId, id).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesRetrieve`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesNotesRetrieve`: %v\n", resp)
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
	experienceNote := *openapiclient.NewExperienceNote(time.Now(), "CreatedBy_example", openapiclient.ExperienceNote_experience{Experience: openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example")}, "Id_example", time.Now(), "ModifiedBy_example") // ExperienceNote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesNotesUpdate(context.Background(), experienceId, id).ExperienceNote(experienceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesNotesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesNotesUpdate`: ExperienceNote
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesNotesUpdate`: %v\n", resp)
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


## V1CatalogExperiencesPartialUpdate

> Experience V1CatalogExperiencesPartialUpdate(ctx, id).Experience(experience).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this experience.
	experience := *openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example") // Experience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesPartialUpdate(context.Background(), id).Experience(experience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesPartialUpdate`: Experience
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this experience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experience** | [**Experience**](Experience.md) |  | 

### Return type

[**Experience**](Experience.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesRetrieve

> Experience V1CatalogExperiencesRetrieve(ctx, id).Expand(expand).Fields(fields).Omit(omit).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this experience.
	expand := "expand_example" // string | Expand related field(s) instead of only showing a UUID (ex: \"required_gpus\"). (optional)
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesRetrieve(context.Background(), id).Expand(expand).Fields(fields).Omit(omit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesRetrieve`: Experience
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this experience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand related field(s) instead of only showing a UUID (ex: \&quot;required_gpus\&quot;). | 
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 

### Return type

[**Experience**](Experience.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogExperiencesStatsRetrieve

> V1CatalogExperiencesStatsRetrieve(ctx).Execute()

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
	r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesStatsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesStatsRetrieveRequest struct via the builder pattern


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


## V1CatalogExperiencesUpdate

> Experience V1CatalogExperiencesUpdate(ctx, id).Experience(experience).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this experience.
	experience := *openapiclient.NewExperience("CatalogId_example", openapiclient.CategoryEnum("AI"), time.Now(), "Experience_example", false, *openapiclient.NewGpuOs("Name_example", "Release_example", "Version_example"), "Id_example", time.Now(), "Persona_example", int64(123), openapiclient.PlatformEnum("air"), openapiclient.SystemArchEnum("amd64"), "Title_example") // Experience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogExperiencesAPI.V1CatalogExperiencesUpdate(context.Background(), id).Experience(experience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogExperiencesAPI.V1CatalogExperiencesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogExperiencesUpdate`: Experience
	fmt.Fprintf(os.Stdout, "Response from `CatalogExperiencesAPI.V1CatalogExperiencesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this experience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogExperiencesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experience** | [**Experience**](Experience.md) |  | 

### Return type

[**Experience**](Experience.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

