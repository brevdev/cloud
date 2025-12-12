# ExperienceBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **NullableString** | User ID who is responsible for manual provisioning | [optional] 
**Autoapprove** | Pointer to **bool** | Can the experience be provisioned without human approval? | [optional] 
**Autoprovision** | Pointer to **bool** | Can the experience be provisioned without human intervention? | [optional] 
**Bootstrap** | Pointer to **bool** | Can this experience be used to bootstrap another compatible experience? | [optional] 
**CatalogId** | **string** | Unique ID for this experience in the sales catalog. Must be unique. | [readonly] 
**CatalogIdAlias** | **NullableString** | Human-readable identifier for the experience in the sales catalog (ex: LP-15). Must be unique. | [readonly] 
**Category** | [**CategoryEnum**](CategoryEnum.md) | Functional group that this experience is targetting  * &#x60;AI&#x60; - AI * &#x60;Clara&#x60; - Clara * &#x60;Data Science&#x60; - Data Science * &#x60;3D Design Collaboration and Simulation&#x60; - 3D Design Collaboration and Simulation * &#x60;Developer&#x60; - Developer * &#x60;Infrastructure Optimization&#x60; - Infrastructure Optimization | 
**CollectionBranch** | Pointer to **string** | Ansible collection branch initialized within the pipeline | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Experience** | **string** | Experience name slug | 
**ExperienceBranch** | Pointer to **string** | Experience branch name used during deployment (default: origin/main) | [optional] 
**FcPlatform** | Pointer to [**NullableFcPlatformEnum**](FcPlatformEnum.md) |  | [optional] 
**FcSupport** | Pointer to **bool** | Does the experience support Flight Control? | [optional] 
**Featured** | **bool** | Is the experience featured in the catalog? | [readonly] 
**GarageId** | Pointer to **NullableString** | ID of the garage where nodes for this experience should be selected from | [optional] 
**GcBranch** | Pointer to **string** | Ground Control branch name (default: main) | [optional] 
**GpuCount** | Pointer to **int32** | Number of GPUs used | [optional] 
**GpuOs** | [**GpuOs**](GpuOs.md) |  | 
**Id** | **string** |  | [readonly] 
**InactivityMax** | Pointer to **int32** | Number of days without user interaction before the experience is torn down (default: 3) | [optional] 
**Lifetime** | Pointer to **NullableInt32** | Default number of days a provisioned experience should remain active (default: 3). A null lifetime will cause a deployment to remain active indefinitely. | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | Pointer to **int32** | Number of Nodes used | [optional] 
**Oem** | Pointer to **NullableString** | If set, the experience must be provisioned to hardware manufactured by the given OEM | [optional] 
**Persona** | **string** |  | 
**Pipeline** | **int64** | Pipeline ID used for provisioning | 
**Platform** | [**PlatformEnum**](PlatformEnum.md) | Base platform that the experience will be provisioned onto  * &#x60;air&#x60; - NVIDIA Air * &#x60;flight_deck&#x60; - Flight Deck * &#x60;kvm_bastion&#x60; - KVM Bastion * &#x60;lp-vmware-platform&#x60; - lp-vmware-platform * &#x60;minimal&#x60; - minimal * &#x60;openshift&#x60; - OpenShift * &#x60;proxmox&#x60; - Proxmox * &#x60;vsphere&#x60; - vSphere * &#x60;vsphere_horizon&#x60; - VMware Horizon * &#x60;vsphere7&#x60; - vSphere 7 * &#x60;vsphere8&#x60; - vSphere 8 | 
**Provider** | Pointer to **NullableString** | If set, the experience must be provisioned to the given provider | [optional] 
**Published** | Pointer to [**PublishedEnum**](PublishedEnum.md) | Is the experience published for use?  * &#x60;draft&#x60; - draft * &#x60;no&#x60; - no * &#x60;yes&#x60; - yes | [optional] 
**Repo** | Pointer to **NullableString** | URL of the repository for provisioning automation | [optional] 
**RequiredGpus** | Pointer to **[]string** | If set, the experience must be provisioned using one of the given GPU types. GPU requirements are evaluated in the order they are set (ex: [\&quot;{l40s_gpu_uuid}\&quot;, \&quot;{a100_gpu_uuid}\&quot;] will prefer an l40s). | [optional] 
**Runtime** | Pointer to **NullableString** | The default runtime to use when provisioning this experience | [optional] 
**SaLab** | Pointer to **bool** | Is this a persistent experience for SAs? | [optional] 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | Required CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**Title** | **string** |  | 
**VgpuProfile** | Pointer to **NullableString** |  | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewExperienceBulkUpdate

`func NewExperienceBulkUpdate(catalogId string, catalogIdAlias NullableString, category CategoryEnum, created time.Time, experience string, featured bool, gpuOs GpuOs, id string, modified time.Time, persona string, pipeline int64, platform PlatformEnum, systemArch SystemArchEnum, title string, count int32, ids []string, result string, ) *ExperienceBulkUpdate`

NewExperienceBulkUpdate instantiates a new ExperienceBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceBulkUpdateWithDefaults

`func NewExperienceBulkUpdateWithDefaults() *ExperienceBulkUpdate`

NewExperienceBulkUpdateWithDefaults instantiates a new ExperienceBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ExperienceBulkUpdate) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ExperienceBulkUpdate) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ExperienceBulkUpdate) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ExperienceBulkUpdate) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *ExperienceBulkUpdate) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *ExperienceBulkUpdate) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAutoapprove

`func (o *ExperienceBulkUpdate) GetAutoapprove() bool`

GetAutoapprove returns the Autoapprove field if non-nil, zero value otherwise.

### GetAutoapproveOk

`func (o *ExperienceBulkUpdate) GetAutoapproveOk() (*bool, bool)`

GetAutoapproveOk returns a tuple with the Autoapprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoapprove

`func (o *ExperienceBulkUpdate) SetAutoapprove(v bool)`

SetAutoapprove sets Autoapprove field to given value.

### HasAutoapprove

`func (o *ExperienceBulkUpdate) HasAutoapprove() bool`

HasAutoapprove returns a boolean if a field has been set.

### GetAutoprovision

`func (o *ExperienceBulkUpdate) GetAutoprovision() bool`

GetAutoprovision returns the Autoprovision field if non-nil, zero value otherwise.

### GetAutoprovisionOk

`func (o *ExperienceBulkUpdate) GetAutoprovisionOk() (*bool, bool)`

GetAutoprovisionOk returns a tuple with the Autoprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoprovision

`func (o *ExperienceBulkUpdate) SetAutoprovision(v bool)`

SetAutoprovision sets Autoprovision field to given value.

### HasAutoprovision

`func (o *ExperienceBulkUpdate) HasAutoprovision() bool`

HasAutoprovision returns a boolean if a field has been set.

### GetBootstrap

`func (o *ExperienceBulkUpdate) GetBootstrap() bool`

GetBootstrap returns the Bootstrap field if non-nil, zero value otherwise.

### GetBootstrapOk

`func (o *ExperienceBulkUpdate) GetBootstrapOk() (*bool, bool)`

GetBootstrapOk returns a tuple with the Bootstrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrap

`func (o *ExperienceBulkUpdate) SetBootstrap(v bool)`

SetBootstrap sets Bootstrap field to given value.

### HasBootstrap

`func (o *ExperienceBulkUpdate) HasBootstrap() bool`

HasBootstrap returns a boolean if a field has been set.

### GetCatalogId

`func (o *ExperienceBulkUpdate) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ExperienceBulkUpdate) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ExperienceBulkUpdate) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.


### GetCatalogIdAlias

`func (o *ExperienceBulkUpdate) GetCatalogIdAlias() string`

GetCatalogIdAlias returns the CatalogIdAlias field if non-nil, zero value otherwise.

### GetCatalogIdAliasOk

`func (o *ExperienceBulkUpdate) GetCatalogIdAliasOk() (*string, bool)`

GetCatalogIdAliasOk returns a tuple with the CatalogIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdAlias

`func (o *ExperienceBulkUpdate) SetCatalogIdAlias(v string)`

SetCatalogIdAlias sets CatalogIdAlias field to given value.


### SetCatalogIdAliasNil

`func (o *ExperienceBulkUpdate) SetCatalogIdAliasNil(b bool)`

 SetCatalogIdAliasNil sets the value for CatalogIdAlias to be an explicit nil

### UnsetCatalogIdAlias
`func (o *ExperienceBulkUpdate) UnsetCatalogIdAlias()`

UnsetCatalogIdAlias ensures that no value is present for CatalogIdAlias, not even an explicit nil
### GetCategory

`func (o *ExperienceBulkUpdate) GetCategory() CategoryEnum`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExperienceBulkUpdate) GetCategoryOk() (*CategoryEnum, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExperienceBulkUpdate) SetCategory(v CategoryEnum)`

SetCategory sets Category field to given value.


### GetCollectionBranch

`func (o *ExperienceBulkUpdate) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *ExperienceBulkUpdate) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *ExperienceBulkUpdate) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *ExperienceBulkUpdate) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### GetCreated

`func (o *ExperienceBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExperienceBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExperienceBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *ExperienceBulkUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperienceBulkUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperienceBulkUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperienceBulkUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExperienceBulkUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExperienceBulkUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExperience

`func (o *ExperienceBulkUpdate) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ExperienceBulkUpdate) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ExperienceBulkUpdate) SetExperience(v string)`

SetExperience sets Experience field to given value.


### GetExperienceBranch

`func (o *ExperienceBulkUpdate) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *ExperienceBulkUpdate) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *ExperienceBulkUpdate) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *ExperienceBulkUpdate) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### GetFcPlatform

`func (o *ExperienceBulkUpdate) GetFcPlatform() FcPlatformEnum`

GetFcPlatform returns the FcPlatform field if non-nil, zero value otherwise.

### GetFcPlatformOk

`func (o *ExperienceBulkUpdate) GetFcPlatformOk() (*FcPlatformEnum, bool)`

GetFcPlatformOk returns a tuple with the FcPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPlatform

`func (o *ExperienceBulkUpdate) SetFcPlatform(v FcPlatformEnum)`

SetFcPlatform sets FcPlatform field to given value.

### HasFcPlatform

`func (o *ExperienceBulkUpdate) HasFcPlatform() bool`

HasFcPlatform returns a boolean if a field has been set.

### SetFcPlatformNil

`func (o *ExperienceBulkUpdate) SetFcPlatformNil(b bool)`

 SetFcPlatformNil sets the value for FcPlatform to be an explicit nil

### UnsetFcPlatform
`func (o *ExperienceBulkUpdate) UnsetFcPlatform()`

UnsetFcPlatform ensures that no value is present for FcPlatform, not even an explicit nil
### GetFcSupport

`func (o *ExperienceBulkUpdate) GetFcSupport() bool`

GetFcSupport returns the FcSupport field if non-nil, zero value otherwise.

### GetFcSupportOk

`func (o *ExperienceBulkUpdate) GetFcSupportOk() (*bool, bool)`

GetFcSupportOk returns a tuple with the FcSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSupport

`func (o *ExperienceBulkUpdate) SetFcSupport(v bool)`

SetFcSupport sets FcSupport field to given value.

### HasFcSupport

`func (o *ExperienceBulkUpdate) HasFcSupport() bool`

HasFcSupport returns a boolean if a field has been set.

### GetFeatured

`func (o *ExperienceBulkUpdate) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ExperienceBulkUpdate) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ExperienceBulkUpdate) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetGarageId

`func (o *ExperienceBulkUpdate) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *ExperienceBulkUpdate) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *ExperienceBulkUpdate) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *ExperienceBulkUpdate) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *ExperienceBulkUpdate) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *ExperienceBulkUpdate) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGcBranch

`func (o *ExperienceBulkUpdate) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *ExperienceBulkUpdate) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *ExperienceBulkUpdate) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *ExperienceBulkUpdate) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### GetGpuCount

`func (o *ExperienceBulkUpdate) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ExperienceBulkUpdate) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ExperienceBulkUpdate) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ExperienceBulkUpdate) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuOs

`func (o *ExperienceBulkUpdate) GetGpuOs() GpuOs`

GetGpuOs returns the GpuOs field if non-nil, zero value otherwise.

### GetGpuOsOk

`func (o *ExperienceBulkUpdate) GetGpuOsOk() (*GpuOs, bool)`

GetGpuOsOk returns a tuple with the GpuOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOs

`func (o *ExperienceBulkUpdate) SetGpuOs(v GpuOs)`

SetGpuOs sets GpuOs field to given value.


### GetId

`func (o *ExperienceBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperienceBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperienceBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetInactivityMax

`func (o *ExperienceBulkUpdate) GetInactivityMax() int32`

GetInactivityMax returns the InactivityMax field if non-nil, zero value otherwise.

### GetInactivityMaxOk

`func (o *ExperienceBulkUpdate) GetInactivityMaxOk() (*int32, bool)`

GetInactivityMaxOk returns a tuple with the InactivityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityMax

`func (o *ExperienceBulkUpdate) SetInactivityMax(v int32)`

SetInactivityMax sets InactivityMax field to given value.

### HasInactivityMax

`func (o *ExperienceBulkUpdate) HasInactivityMax() bool`

HasInactivityMax returns a boolean if a field has been set.

### GetLifetime

`func (o *ExperienceBulkUpdate) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *ExperienceBulkUpdate) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *ExperienceBulkUpdate) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *ExperienceBulkUpdate) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### SetLifetimeNil

`func (o *ExperienceBulkUpdate) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *ExperienceBulkUpdate) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetModified

`func (o *ExperienceBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ExperienceBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ExperienceBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *ExperienceBulkUpdate) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ExperienceBulkUpdate) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ExperienceBulkUpdate) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ExperienceBulkUpdate) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetOem

`func (o *ExperienceBulkUpdate) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *ExperienceBulkUpdate) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *ExperienceBulkUpdate) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *ExperienceBulkUpdate) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *ExperienceBulkUpdate) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *ExperienceBulkUpdate) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetPersona

`func (o *ExperienceBulkUpdate) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *ExperienceBulkUpdate) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *ExperienceBulkUpdate) SetPersona(v string)`

SetPersona sets Persona field to given value.


### GetPipeline

`func (o *ExperienceBulkUpdate) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *ExperienceBulkUpdate) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *ExperienceBulkUpdate) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.


### GetPlatform

`func (o *ExperienceBulkUpdate) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ExperienceBulkUpdate) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ExperienceBulkUpdate) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.


### GetProvider

`func (o *ExperienceBulkUpdate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ExperienceBulkUpdate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ExperienceBulkUpdate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ExperienceBulkUpdate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ExperienceBulkUpdate) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ExperienceBulkUpdate) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetPublished

`func (o *ExperienceBulkUpdate) GetPublished() PublishedEnum`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ExperienceBulkUpdate) GetPublishedOk() (*PublishedEnum, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ExperienceBulkUpdate) SetPublished(v PublishedEnum)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ExperienceBulkUpdate) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *ExperienceBulkUpdate) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ExperienceBulkUpdate) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ExperienceBulkUpdate) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ExperienceBulkUpdate) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *ExperienceBulkUpdate) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *ExperienceBulkUpdate) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetRequiredGpus

`func (o *ExperienceBulkUpdate) GetRequiredGpus() []string`

GetRequiredGpus returns the RequiredGpus field if non-nil, zero value otherwise.

### GetRequiredGpusOk

`func (o *ExperienceBulkUpdate) GetRequiredGpusOk() (*[]string, bool)`

GetRequiredGpusOk returns a tuple with the RequiredGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGpus

`func (o *ExperienceBulkUpdate) SetRequiredGpus(v []string)`

SetRequiredGpus sets RequiredGpus field to given value.

### HasRequiredGpus

`func (o *ExperienceBulkUpdate) HasRequiredGpus() bool`

HasRequiredGpus returns a boolean if a field has been set.

### GetRuntime

`func (o *ExperienceBulkUpdate) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ExperienceBulkUpdate) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ExperienceBulkUpdate) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ExperienceBulkUpdate) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *ExperienceBulkUpdate) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *ExperienceBulkUpdate) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetSaLab

`func (o *ExperienceBulkUpdate) GetSaLab() bool`

GetSaLab returns the SaLab field if non-nil, zero value otherwise.

### GetSaLabOk

`func (o *ExperienceBulkUpdate) GetSaLabOk() (*bool, bool)`

GetSaLabOk returns a tuple with the SaLab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLab

`func (o *ExperienceBulkUpdate) SetSaLab(v bool)`

SetSaLab sets SaLab field to given value.

### HasSaLab

`func (o *ExperienceBulkUpdate) HasSaLab() bool`

HasSaLab returns a boolean if a field has been set.

### GetSystemArch

`func (o *ExperienceBulkUpdate) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *ExperienceBulkUpdate) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *ExperienceBulkUpdate) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetTitle

`func (o *ExperienceBulkUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExperienceBulkUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExperienceBulkUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVgpuProfile

`func (o *ExperienceBulkUpdate) GetVgpuProfile() string`

GetVgpuProfile returns the VgpuProfile field if non-nil, zero value otherwise.

### GetVgpuProfileOk

`func (o *ExperienceBulkUpdate) GetVgpuProfileOk() (*string, bool)`

GetVgpuProfileOk returns a tuple with the VgpuProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuProfile

`func (o *ExperienceBulkUpdate) SetVgpuProfile(v string)`

SetVgpuProfile sets VgpuProfile field to given value.

### HasVgpuProfile

`func (o *ExperienceBulkUpdate) HasVgpuProfile() bool`

HasVgpuProfile returns a boolean if a field has been set.

### SetVgpuProfileNil

`func (o *ExperienceBulkUpdate) SetVgpuProfileNil(b bool)`

 SetVgpuProfileNil sets the value for VgpuProfile to be an explicit nil

### UnsetVgpuProfile
`func (o *ExperienceBulkUpdate) UnsetVgpuProfile()`

UnsetVgpuProfile ensures that no value is present for VgpuProfile, not even an explicit nil
### GetCount

`func (o *ExperienceBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ExperienceBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ExperienceBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *ExperienceBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ExperienceBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ExperienceBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *ExperienceBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExperienceBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExperienceBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


