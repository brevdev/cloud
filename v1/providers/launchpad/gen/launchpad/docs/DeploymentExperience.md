# DeploymentExperience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | User ID who is responsible for manual provisioning | [optional] 
**Autoapprove** | Pointer to **bool** | Can the experience be provisioned without human approval? | [optional] 
**Autoprovision** | Pointer to **bool** | Can the experience be provisioned without human intervention? | [optional] 
**Bootstrap** | Pointer to **bool** | Can this experience be used to bootstrap another compatible experience? | [optional] 
**CatalogId** | **string** | Unique ID for this experience in the sales catalog. Must be unique. | 
**CatalogIdAlias** | Pointer to **string** | Human-readable identifier for the experience in the sales catalog (ex: LP-15). Must be unique. | [optional] 
**Category** | [**CategoryEnum**](CategoryEnum.md) | Functional group that this experience is targetting  * &#x60;AI&#x60; - AI * &#x60;Clara&#x60; - Clara * &#x60;Data Science&#x60; - Data Science * &#x60;3D Design Collaboration and Simulation&#x60; - 3D Design Collaboration and Simulation * &#x60;Developer&#x60; - Developer * &#x60;Infrastructure Optimization&#x60; - Infrastructure Optimization | 
**CollectionBranch** | Pointer to **string** | Ansible collection branch initialized within the pipeline | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Experience** | **string** | Experience name slug | 
**ExperienceBranch** | Pointer to **string** | Experience branch name used during deployment (default: origin/main) | [optional] 
**FcPlatform** | Pointer to [**FcPlatformEnum**](FcPlatformEnum.md) |  | [optional] 
**FcSupport** | Pointer to **bool** | Does the experience support Flight Control? | [optional] 
**GarageId** | Pointer to **string** | ID of the garage where nodes for this experience should be selected from | [optional] 
**GcBranch** | Pointer to **string** | Ground Control branch name (default: main) | [optional] 
**GpuCount** | Pointer to **int32** | Number of GPUs used | [optional] 
**GpuOs** | [**GpuOs**](GpuOs.md) |  | 
**Id** | **string** |  | [readonly] 
**InactivityMax** | Pointer to **int32** | Number of days without user interaction before the experience is torn down (default: 3) | [optional] 
**Lifetime** | Pointer to **int32** | Default number of days a provisioned experience should remain active (default: 3). A null lifetime will cause a deployment to remain active indefinitely. | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | Pointer to **int32** | Number of Nodes used | [optional] 
**Persona** | **string** |  | 
**Pipeline** | **int64** | Pipeline ID used for provisioning | 
**Platform** | [**PlatformEnum**](PlatformEnum.md) | Base platform that the experience will be provisioned onto  * &#x60;air&#x60; - NVIDIA Air * &#x60;flight_deck&#x60; - Flight Deck * &#x60;kvm_bastion&#x60; - KVM Bastion * &#x60;lp-vmware-platform&#x60; - lp-vmware-platform * &#x60;minimal&#x60; - minimal * &#x60;openshift&#x60; - OpenShift * &#x60;vsphere&#x60; - vSphere * &#x60;vsphere_horizon&#x60; - VMware Horizon * &#x60;vsphere7&#x60; - vSphere 7 * &#x60;vsphere8&#x60; - vSphere 8 | 
**Provider** | Pointer to **string** | If set, the experience must be provisioned to the given provider | [optional] 
**Published** | Pointer to [**PublishedEnum**](PublishedEnum.md) | Is the experience published for use?  * &#x60;draft&#x60; - draft * &#x60;no&#x60; - no * &#x60;yes&#x60; - yes | [optional] 
**Repo** | Pointer to **string** | URL of the repository for provisioning automation | [optional] 
**RequiredGpus** | Pointer to **[]string** | If set, the experience must be provisioned using one of the given GPU types. GPU requirements are evaluated in the order they are set (ex: [\&quot;{l40s_gpu_uuid}\&quot;, \&quot;{a100_gpu_uuid}\&quot;] will prefer an l40s). | [optional] 
**SaLab** | Pointer to **bool** | Is this a persistent experience for SAs? | [optional] 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | Required CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**Title** | **string** |  | 
**VgpuProfile** | Pointer to [**VgpuProfileEnum**](VgpuProfileEnum.md) |  | [optional] 

## Methods

### NewDeploymentExperience

`func NewDeploymentExperience(catalogId string, category CategoryEnum, created time.Time, experience string, gpuOs GpuOs, id string, modified time.Time, persona string, pipeline int64, platform PlatformEnum, systemArch SystemArchEnum, title string, ) *DeploymentExperience`

NewDeploymentExperience instantiates a new DeploymentExperience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentExperienceWithDefaults

`func NewDeploymentExperienceWithDefaults() *DeploymentExperience`

NewDeploymentExperienceWithDefaults instantiates a new DeploymentExperience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *DeploymentExperience) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *DeploymentExperience) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *DeploymentExperience) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *DeploymentExperience) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAutoapprove

`func (o *DeploymentExperience) GetAutoapprove() bool`

GetAutoapprove returns the Autoapprove field if non-nil, zero value otherwise.

### GetAutoapproveOk

`func (o *DeploymentExperience) GetAutoapproveOk() (*bool, bool)`

GetAutoapproveOk returns a tuple with the Autoapprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoapprove

`func (o *DeploymentExperience) SetAutoapprove(v bool)`

SetAutoapprove sets Autoapprove field to given value.

### HasAutoapprove

`func (o *DeploymentExperience) HasAutoapprove() bool`

HasAutoapprove returns a boolean if a field has been set.

### GetAutoprovision

`func (o *DeploymentExperience) GetAutoprovision() bool`

GetAutoprovision returns the Autoprovision field if non-nil, zero value otherwise.

### GetAutoprovisionOk

`func (o *DeploymentExperience) GetAutoprovisionOk() (*bool, bool)`

GetAutoprovisionOk returns a tuple with the Autoprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoprovision

`func (o *DeploymentExperience) SetAutoprovision(v bool)`

SetAutoprovision sets Autoprovision field to given value.

### HasAutoprovision

`func (o *DeploymentExperience) HasAutoprovision() bool`

HasAutoprovision returns a boolean if a field has been set.

### GetBootstrap

`func (o *DeploymentExperience) GetBootstrap() bool`

GetBootstrap returns the Bootstrap field if non-nil, zero value otherwise.

### GetBootstrapOk

`func (o *DeploymentExperience) GetBootstrapOk() (*bool, bool)`

GetBootstrapOk returns a tuple with the Bootstrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrap

`func (o *DeploymentExperience) SetBootstrap(v bool)`

SetBootstrap sets Bootstrap field to given value.

### HasBootstrap

`func (o *DeploymentExperience) HasBootstrap() bool`

HasBootstrap returns a boolean if a field has been set.

### GetCatalogId

`func (o *DeploymentExperience) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *DeploymentExperience) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *DeploymentExperience) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.


### GetCatalogIdAlias

`func (o *DeploymentExperience) GetCatalogIdAlias() string`

GetCatalogIdAlias returns the CatalogIdAlias field if non-nil, zero value otherwise.

### GetCatalogIdAliasOk

`func (o *DeploymentExperience) GetCatalogIdAliasOk() (*string, bool)`

GetCatalogIdAliasOk returns a tuple with the CatalogIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdAlias

`func (o *DeploymentExperience) SetCatalogIdAlias(v string)`

SetCatalogIdAlias sets CatalogIdAlias field to given value.

### HasCatalogIdAlias

`func (o *DeploymentExperience) HasCatalogIdAlias() bool`

HasCatalogIdAlias returns a boolean if a field has been set.

### GetCategory

`func (o *DeploymentExperience) GetCategory() CategoryEnum`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DeploymentExperience) GetCategoryOk() (*CategoryEnum, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DeploymentExperience) SetCategory(v CategoryEnum)`

SetCategory sets Category field to given value.


### GetCollectionBranch

`func (o *DeploymentExperience) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *DeploymentExperience) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *DeploymentExperience) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *DeploymentExperience) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### GetCreated

`func (o *DeploymentExperience) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentExperience) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentExperience) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *DeploymentExperience) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentExperience) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentExperience) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentExperience) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExperience

`func (o *DeploymentExperience) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *DeploymentExperience) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *DeploymentExperience) SetExperience(v string)`

SetExperience sets Experience field to given value.


### GetExperienceBranch

`func (o *DeploymentExperience) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *DeploymentExperience) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *DeploymentExperience) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *DeploymentExperience) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### GetFcPlatform

`func (o *DeploymentExperience) GetFcPlatform() FcPlatformEnum`

GetFcPlatform returns the FcPlatform field if non-nil, zero value otherwise.

### GetFcPlatformOk

`func (o *DeploymentExperience) GetFcPlatformOk() (*FcPlatformEnum, bool)`

GetFcPlatformOk returns a tuple with the FcPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPlatform

`func (o *DeploymentExperience) SetFcPlatform(v FcPlatformEnum)`

SetFcPlatform sets FcPlatform field to given value.

### HasFcPlatform

`func (o *DeploymentExperience) HasFcPlatform() bool`

HasFcPlatform returns a boolean if a field has been set.

### GetFcSupport

`func (o *DeploymentExperience) GetFcSupport() bool`

GetFcSupport returns the FcSupport field if non-nil, zero value otherwise.

### GetFcSupportOk

`func (o *DeploymentExperience) GetFcSupportOk() (*bool, bool)`

GetFcSupportOk returns a tuple with the FcSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSupport

`func (o *DeploymentExperience) SetFcSupport(v bool)`

SetFcSupport sets FcSupport field to given value.

### HasFcSupport

`func (o *DeploymentExperience) HasFcSupport() bool`

HasFcSupport returns a boolean if a field has been set.

### GetGarageId

`func (o *DeploymentExperience) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *DeploymentExperience) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *DeploymentExperience) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *DeploymentExperience) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### GetGcBranch

`func (o *DeploymentExperience) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *DeploymentExperience) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *DeploymentExperience) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *DeploymentExperience) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### GetGpuCount

`func (o *DeploymentExperience) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *DeploymentExperience) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *DeploymentExperience) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *DeploymentExperience) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuOs

`func (o *DeploymentExperience) GetGpuOs() GpuOs`

GetGpuOs returns the GpuOs field if non-nil, zero value otherwise.

### GetGpuOsOk

`func (o *DeploymentExperience) GetGpuOsOk() (*GpuOs, bool)`

GetGpuOsOk returns a tuple with the GpuOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOs

`func (o *DeploymentExperience) SetGpuOs(v GpuOs)`

SetGpuOs sets GpuOs field to given value.


### GetId

`func (o *DeploymentExperience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentExperience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentExperience) SetId(v string)`

SetId sets Id field to given value.


### GetInactivityMax

`func (o *DeploymentExperience) GetInactivityMax() int32`

GetInactivityMax returns the InactivityMax field if non-nil, zero value otherwise.

### GetInactivityMaxOk

`func (o *DeploymentExperience) GetInactivityMaxOk() (*int32, bool)`

GetInactivityMaxOk returns a tuple with the InactivityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityMax

`func (o *DeploymentExperience) SetInactivityMax(v int32)`

SetInactivityMax sets InactivityMax field to given value.

### HasInactivityMax

`func (o *DeploymentExperience) HasInactivityMax() bool`

HasInactivityMax returns a boolean if a field has been set.

### GetLifetime

`func (o *DeploymentExperience) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *DeploymentExperience) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *DeploymentExperience) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *DeploymentExperience) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetModified

`func (o *DeploymentExperience) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentExperience) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentExperience) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *DeploymentExperience) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *DeploymentExperience) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *DeploymentExperience) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *DeploymentExperience) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetPersona

`func (o *DeploymentExperience) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *DeploymentExperience) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *DeploymentExperience) SetPersona(v string)`

SetPersona sets Persona field to given value.


### GetPipeline

`func (o *DeploymentExperience) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DeploymentExperience) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DeploymentExperience) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.


### GetPlatform

`func (o *DeploymentExperience) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeploymentExperience) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeploymentExperience) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.


### GetProvider

`func (o *DeploymentExperience) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DeploymentExperience) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DeploymentExperience) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DeploymentExperience) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *DeploymentExperience) GetPublished() PublishedEnum`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DeploymentExperience) GetPublishedOk() (*PublishedEnum, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DeploymentExperience) SetPublished(v PublishedEnum)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *DeploymentExperience) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *DeploymentExperience) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *DeploymentExperience) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *DeploymentExperience) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *DeploymentExperience) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRequiredGpus

`func (o *DeploymentExperience) GetRequiredGpus() []string`

GetRequiredGpus returns the RequiredGpus field if non-nil, zero value otherwise.

### GetRequiredGpusOk

`func (o *DeploymentExperience) GetRequiredGpusOk() (*[]string, bool)`

GetRequiredGpusOk returns a tuple with the RequiredGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGpus

`func (o *DeploymentExperience) SetRequiredGpus(v []string)`

SetRequiredGpus sets RequiredGpus field to given value.

### HasRequiredGpus

`func (o *DeploymentExperience) HasRequiredGpus() bool`

HasRequiredGpus returns a boolean if a field has been set.

### GetSaLab

`func (o *DeploymentExperience) GetSaLab() bool`

GetSaLab returns the SaLab field if non-nil, zero value otherwise.

### GetSaLabOk

`func (o *DeploymentExperience) GetSaLabOk() (*bool, bool)`

GetSaLabOk returns a tuple with the SaLab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLab

`func (o *DeploymentExperience) SetSaLab(v bool)`

SetSaLab sets SaLab field to given value.

### HasSaLab

`func (o *DeploymentExperience) HasSaLab() bool`

HasSaLab returns a boolean if a field has been set.

### GetSystemArch

`func (o *DeploymentExperience) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *DeploymentExperience) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *DeploymentExperience) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetTitle

`func (o *DeploymentExperience) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeploymentExperience) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeploymentExperience) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVgpuProfile

`func (o *DeploymentExperience) GetVgpuProfile() VgpuProfileEnum`

GetVgpuProfile returns the VgpuProfile field if non-nil, zero value otherwise.

### GetVgpuProfileOk

`func (o *DeploymentExperience) GetVgpuProfileOk() (*VgpuProfileEnum, bool)`

GetVgpuProfileOk returns a tuple with the VgpuProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuProfile

`func (o *DeploymentExperience) SetVgpuProfile(v VgpuProfileEnum)`

SetVgpuProfile sets VgpuProfile field to given value.

### HasVgpuProfile

`func (o *DeploymentExperience) HasVgpuProfile() bool`

HasVgpuProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


