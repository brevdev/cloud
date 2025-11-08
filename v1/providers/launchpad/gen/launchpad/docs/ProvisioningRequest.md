# ProvisioningRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionOperatingSystem** | Pointer to **NullableString** | Override bastion operating system provisioned and/or configured by Liftoff | [optional] 
**CatalogId** | Pointer to **string** | Catalog ID of the experience provisioned to the cluster | [optional] 
**CatalogIdAlias** | Pointer to **string** | Catalog ID alias of the experience provisioned to the cluster | [optional] 
**Cluster** | Pointer to **NullableString** | Target a specific cluster for provisioning | [optional] 
**CollectionBranch** | Pointer to **NullableString** | Override the Ansible collection branch initialized within the pipeline | [optional] 
**Experience** | **NullableString** | The experience being deployed for use | [readonly] 
**ExperienceBranch** | Pointer to **NullableString** | Override the experience branch | [optional] 
**ExperienceId** | Pointer to **string** | UUID of the experience provisioned to the cluster | [optional] 
**FlightcontrolRelease** | Pointer to **NullableString** | Override the image tag used for Flight Control | [optional] 
**GarageId** | Pointer to **NullableString** | Require a cluster with nodes in the given garage | [optional] 
**GcBranch** | Pointer to **NullableString** | Override the default Ground Control branch | [optional] 
**GpuAlias** | Pointer to **NullableString** | Require a cluster with the given GPU alias | [optional] 
**GpuCount** | Pointer to **int32** | Require a cluster with the given number of GPUs | [optional] 
**GpuModel** | Pointer to **NullableString** | Require a cluster with the given GPU model | [optional] 
**GpuOsName** | Pointer to **NullableString** | Override the GPU node operating system name | [optional] 
**GpuOsRelease** | Pointer to **NullableString** | Override the GPU node operating system release | [optional] 
**GpuOsVersion** | Pointer to **NullableString** | Override the GPU node operating system version | [optional] 
**MinGpuCount** | Pointer to **int32** | Require a cluster whose GPU count is greater than or equal to the given number | [optional] 
**NodeCount** | Pointer to **int32** | Require a cluster with the given number of nodes | [optional] 
**OemName** | Pointer to **NullableString** | Require a cluster manufactured by the given OEM name | [optional] 
**PersistOnFailure** | Pointer to **NullableBool** | Override the default cleanup/destroy behavior when a provisioning failure occurs | [optional] 
**Persona** | Pointer to **NullableString** | Override the defined persona in the experience | [optional] 
**Pipeline** | Pointer to **int64** | Override the pipeline ID that will be triggered for request fulfillment | [optional] 
**PipelineBranch** | Pointer to **NullableString** | Override the default pipeline branch ref used when triggering a Fuselage pipeline | [optional] 
**Platform** | Pointer to [**NullablePlatformEnum**](PlatformEnum.md) |  | [optional] 
**ProviderName** | Pointer to **NullableString** | Require a cluster from the given provider name | [optional] 
**Region** | Pointer to **NullableString** | Require a cluster located in the given region | [optional] 
**RequestId** | **string** | Trial request ID (ex: TRY-1234) | 
**Runtime** | Pointer to **NullableString** | Use the presets of the given runtime when provisioning this experience | [optional] 
**RuntimeBranch** | Pointer to **NullableString** | Override the runtime repository branch | [optional] 
**RuntimeCnsAddonPack** | Pointer to **NullableBool** | Override the runtime&#39;s CNS add-ons flag | [optional] 
**RuntimeCnsDocker** | Pointer to **NullableBool** | Override the runtime&#39;s Docker with CNS flag | [optional] 
**RuntimeCnsDriverVersion** | Pointer to **NullableString** | Override the runtime&#39;s GPU driver version | [optional] 
**RuntimeCnsK8s** | Pointer to **NullableBool** | Override the runtime&#39;s Kubernetes with CNS flag | [optional] 
**RuntimeCnsNvidiaDriver** | Pointer to **NullableBool** | Override the runtime&#39;s NVIDIA driver with CNS flag | [optional] 
**RuntimeCnsVersion** | Pointer to **NullableString** | Override the runtime&#39;s Cloud Native Stack version | [optional] 
**RuntimeMig** | Pointer to **NullableBool** | Override the runtime&#39;s MIG support with CNS flag | [optional] 
**RuntimeMigProfile** | Pointer to **NullableString** | Override the runtime&#39;s MIG profile name | [optional] 
**RuntimeUrl** | Pointer to **NullableString** | Override the URL of the runtime repository | [optional] 
**Workshop** | Pointer to **NullableBool** | Require a cluster whose workshop flag is set | [optional] 
**WorkshopId** | Pointer to **NullableString** | Require a cluster with the given workshop ID | [optional] 
**WorkshopOverridePassword** | Pointer to **NullableString** | Override the deployment&#39;s default authentication to use a static password. This is useful for workshops when you&#39;d like an identical password associated with a collection of environments. (LaunchPad Team only) | [optional] 

## Methods

### NewProvisioningRequest

`func NewProvisioningRequest(experience NullableString, requestId string, ) *ProvisioningRequest`

NewProvisioningRequest instantiates a new ProvisioningRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningRequestWithDefaults

`func NewProvisioningRequestWithDefaults() *ProvisioningRequest`

NewProvisioningRequestWithDefaults instantiates a new ProvisioningRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionOperatingSystem

`func (o *ProvisioningRequest) GetBastionOperatingSystem() string`

GetBastionOperatingSystem returns the BastionOperatingSystem field if non-nil, zero value otherwise.

### GetBastionOperatingSystemOk

`func (o *ProvisioningRequest) GetBastionOperatingSystemOk() (*string, bool)`

GetBastionOperatingSystemOk returns a tuple with the BastionOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionOperatingSystem

`func (o *ProvisioningRequest) SetBastionOperatingSystem(v string)`

SetBastionOperatingSystem sets BastionOperatingSystem field to given value.

### HasBastionOperatingSystem

`func (o *ProvisioningRequest) HasBastionOperatingSystem() bool`

HasBastionOperatingSystem returns a boolean if a field has been set.

### SetBastionOperatingSystemNil

`func (o *ProvisioningRequest) SetBastionOperatingSystemNil(b bool)`

 SetBastionOperatingSystemNil sets the value for BastionOperatingSystem to be an explicit nil

### UnsetBastionOperatingSystem
`func (o *ProvisioningRequest) UnsetBastionOperatingSystem()`

UnsetBastionOperatingSystem ensures that no value is present for BastionOperatingSystem, not even an explicit nil
### GetCatalogId

`func (o *ProvisioningRequest) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ProvisioningRequest) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ProvisioningRequest) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *ProvisioningRequest) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogIdAlias

`func (o *ProvisioningRequest) GetCatalogIdAlias() string`

GetCatalogIdAlias returns the CatalogIdAlias field if non-nil, zero value otherwise.

### GetCatalogIdAliasOk

`func (o *ProvisioningRequest) GetCatalogIdAliasOk() (*string, bool)`

GetCatalogIdAliasOk returns a tuple with the CatalogIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdAlias

`func (o *ProvisioningRequest) SetCatalogIdAlias(v string)`

SetCatalogIdAlias sets CatalogIdAlias field to given value.

### HasCatalogIdAlias

`func (o *ProvisioningRequest) HasCatalogIdAlias() bool`

HasCatalogIdAlias returns a boolean if a field has been set.

### GetCluster

`func (o *ProvisioningRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ProvisioningRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ProvisioningRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ProvisioningRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *ProvisioningRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *ProvisioningRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetCollectionBranch

`func (o *ProvisioningRequest) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *ProvisioningRequest) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *ProvisioningRequest) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *ProvisioningRequest) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### SetCollectionBranchNil

`func (o *ProvisioningRequest) SetCollectionBranchNil(b bool)`

 SetCollectionBranchNil sets the value for CollectionBranch to be an explicit nil

### UnsetCollectionBranch
`func (o *ProvisioningRequest) UnsetCollectionBranch()`

UnsetCollectionBranch ensures that no value is present for CollectionBranch, not even an explicit nil
### GetExperience

`func (o *ProvisioningRequest) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ProvisioningRequest) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ProvisioningRequest) SetExperience(v string)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *ProvisioningRequest) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *ProvisioningRequest) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetExperienceBranch

`func (o *ProvisioningRequest) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *ProvisioningRequest) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *ProvisioningRequest) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *ProvisioningRequest) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### SetExperienceBranchNil

`func (o *ProvisioningRequest) SetExperienceBranchNil(b bool)`

 SetExperienceBranchNil sets the value for ExperienceBranch to be an explicit nil

### UnsetExperienceBranch
`func (o *ProvisioningRequest) UnsetExperienceBranch()`

UnsetExperienceBranch ensures that no value is present for ExperienceBranch, not even an explicit nil
### GetExperienceId

`func (o *ProvisioningRequest) GetExperienceId() string`

GetExperienceId returns the ExperienceId field if non-nil, zero value otherwise.

### GetExperienceIdOk

`func (o *ProvisioningRequest) GetExperienceIdOk() (*string, bool)`

GetExperienceIdOk returns a tuple with the ExperienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceId

`func (o *ProvisioningRequest) SetExperienceId(v string)`

SetExperienceId sets ExperienceId field to given value.

### HasExperienceId

`func (o *ProvisioningRequest) HasExperienceId() bool`

HasExperienceId returns a boolean if a field has been set.

### GetFlightcontrolRelease

`func (o *ProvisioningRequest) GetFlightcontrolRelease() string`

GetFlightcontrolRelease returns the FlightcontrolRelease field if non-nil, zero value otherwise.

### GetFlightcontrolReleaseOk

`func (o *ProvisioningRequest) GetFlightcontrolReleaseOk() (*string, bool)`

GetFlightcontrolReleaseOk returns a tuple with the FlightcontrolRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightcontrolRelease

`func (o *ProvisioningRequest) SetFlightcontrolRelease(v string)`

SetFlightcontrolRelease sets FlightcontrolRelease field to given value.

### HasFlightcontrolRelease

`func (o *ProvisioningRequest) HasFlightcontrolRelease() bool`

HasFlightcontrolRelease returns a boolean if a field has been set.

### SetFlightcontrolReleaseNil

`func (o *ProvisioningRequest) SetFlightcontrolReleaseNil(b bool)`

 SetFlightcontrolReleaseNil sets the value for FlightcontrolRelease to be an explicit nil

### UnsetFlightcontrolRelease
`func (o *ProvisioningRequest) UnsetFlightcontrolRelease()`

UnsetFlightcontrolRelease ensures that no value is present for FlightcontrolRelease, not even an explicit nil
### GetGarageId

`func (o *ProvisioningRequest) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *ProvisioningRequest) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *ProvisioningRequest) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *ProvisioningRequest) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *ProvisioningRequest) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *ProvisioningRequest) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGcBranch

`func (o *ProvisioningRequest) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *ProvisioningRequest) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *ProvisioningRequest) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *ProvisioningRequest) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### SetGcBranchNil

`func (o *ProvisioningRequest) SetGcBranchNil(b bool)`

 SetGcBranchNil sets the value for GcBranch to be an explicit nil

### UnsetGcBranch
`func (o *ProvisioningRequest) UnsetGcBranch()`

UnsetGcBranch ensures that no value is present for GcBranch, not even an explicit nil
### GetGpuAlias

`func (o *ProvisioningRequest) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *ProvisioningRequest) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *ProvisioningRequest) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *ProvisioningRequest) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### SetGpuAliasNil

`func (o *ProvisioningRequest) SetGpuAliasNil(b bool)`

 SetGpuAliasNil sets the value for GpuAlias to be an explicit nil

### UnsetGpuAlias
`func (o *ProvisioningRequest) UnsetGpuAlias()`

UnsetGpuAlias ensures that no value is present for GpuAlias, not even an explicit nil
### GetGpuCount

`func (o *ProvisioningRequest) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ProvisioningRequest) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ProvisioningRequest) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ProvisioningRequest) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *ProvisioningRequest) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ProvisioningRequest) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ProvisioningRequest) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *ProvisioningRequest) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### SetGpuModelNil

`func (o *ProvisioningRequest) SetGpuModelNil(b bool)`

 SetGpuModelNil sets the value for GpuModel to be an explicit nil

### UnsetGpuModel
`func (o *ProvisioningRequest) UnsetGpuModel()`

UnsetGpuModel ensures that no value is present for GpuModel, not even an explicit nil
### GetGpuOsName

`func (o *ProvisioningRequest) GetGpuOsName() string`

GetGpuOsName returns the GpuOsName field if non-nil, zero value otherwise.

### GetGpuOsNameOk

`func (o *ProvisioningRequest) GetGpuOsNameOk() (*string, bool)`

GetGpuOsNameOk returns a tuple with the GpuOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsName

`func (o *ProvisioningRequest) SetGpuOsName(v string)`

SetGpuOsName sets GpuOsName field to given value.

### HasGpuOsName

`func (o *ProvisioningRequest) HasGpuOsName() bool`

HasGpuOsName returns a boolean if a field has been set.

### SetGpuOsNameNil

`func (o *ProvisioningRequest) SetGpuOsNameNil(b bool)`

 SetGpuOsNameNil sets the value for GpuOsName to be an explicit nil

### UnsetGpuOsName
`func (o *ProvisioningRequest) UnsetGpuOsName()`

UnsetGpuOsName ensures that no value is present for GpuOsName, not even an explicit nil
### GetGpuOsRelease

`func (o *ProvisioningRequest) GetGpuOsRelease() string`

GetGpuOsRelease returns the GpuOsRelease field if non-nil, zero value otherwise.

### GetGpuOsReleaseOk

`func (o *ProvisioningRequest) GetGpuOsReleaseOk() (*string, bool)`

GetGpuOsReleaseOk returns a tuple with the GpuOsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsRelease

`func (o *ProvisioningRequest) SetGpuOsRelease(v string)`

SetGpuOsRelease sets GpuOsRelease field to given value.

### HasGpuOsRelease

`func (o *ProvisioningRequest) HasGpuOsRelease() bool`

HasGpuOsRelease returns a boolean if a field has been set.

### SetGpuOsReleaseNil

`func (o *ProvisioningRequest) SetGpuOsReleaseNil(b bool)`

 SetGpuOsReleaseNil sets the value for GpuOsRelease to be an explicit nil

### UnsetGpuOsRelease
`func (o *ProvisioningRequest) UnsetGpuOsRelease()`

UnsetGpuOsRelease ensures that no value is present for GpuOsRelease, not even an explicit nil
### GetGpuOsVersion

`func (o *ProvisioningRequest) GetGpuOsVersion() string`

GetGpuOsVersion returns the GpuOsVersion field if non-nil, zero value otherwise.

### GetGpuOsVersionOk

`func (o *ProvisioningRequest) GetGpuOsVersionOk() (*string, bool)`

GetGpuOsVersionOk returns a tuple with the GpuOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsVersion

`func (o *ProvisioningRequest) SetGpuOsVersion(v string)`

SetGpuOsVersion sets GpuOsVersion field to given value.

### HasGpuOsVersion

`func (o *ProvisioningRequest) HasGpuOsVersion() bool`

HasGpuOsVersion returns a boolean if a field has been set.

### SetGpuOsVersionNil

`func (o *ProvisioningRequest) SetGpuOsVersionNil(b bool)`

 SetGpuOsVersionNil sets the value for GpuOsVersion to be an explicit nil

### UnsetGpuOsVersion
`func (o *ProvisioningRequest) UnsetGpuOsVersion()`

UnsetGpuOsVersion ensures that no value is present for GpuOsVersion, not even an explicit nil
### GetMinGpuCount

`func (o *ProvisioningRequest) GetMinGpuCount() int32`

GetMinGpuCount returns the MinGpuCount field if non-nil, zero value otherwise.

### GetMinGpuCountOk

`func (o *ProvisioningRequest) GetMinGpuCountOk() (*int32, bool)`

GetMinGpuCountOk returns a tuple with the MinGpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGpuCount

`func (o *ProvisioningRequest) SetMinGpuCount(v int32)`

SetMinGpuCount sets MinGpuCount field to given value.

### HasMinGpuCount

`func (o *ProvisioningRequest) HasMinGpuCount() bool`

HasMinGpuCount returns a boolean if a field has been set.

### GetNodeCount

`func (o *ProvisioningRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ProvisioningRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ProvisioningRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ProvisioningRequest) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetOemName

`func (o *ProvisioningRequest) GetOemName() string`

GetOemName returns the OemName field if non-nil, zero value otherwise.

### GetOemNameOk

`func (o *ProvisioningRequest) GetOemNameOk() (*string, bool)`

GetOemNameOk returns a tuple with the OemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemName

`func (o *ProvisioningRequest) SetOemName(v string)`

SetOemName sets OemName field to given value.

### HasOemName

`func (o *ProvisioningRequest) HasOemName() bool`

HasOemName returns a boolean if a field has been set.

### SetOemNameNil

`func (o *ProvisioningRequest) SetOemNameNil(b bool)`

 SetOemNameNil sets the value for OemName to be an explicit nil

### UnsetOemName
`func (o *ProvisioningRequest) UnsetOemName()`

UnsetOemName ensures that no value is present for OemName, not even an explicit nil
### GetPersistOnFailure

`func (o *ProvisioningRequest) GetPersistOnFailure() bool`

GetPersistOnFailure returns the PersistOnFailure field if non-nil, zero value otherwise.

### GetPersistOnFailureOk

`func (o *ProvisioningRequest) GetPersistOnFailureOk() (*bool, bool)`

GetPersistOnFailureOk returns a tuple with the PersistOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOnFailure

`func (o *ProvisioningRequest) SetPersistOnFailure(v bool)`

SetPersistOnFailure sets PersistOnFailure field to given value.

### HasPersistOnFailure

`func (o *ProvisioningRequest) HasPersistOnFailure() bool`

HasPersistOnFailure returns a boolean if a field has been set.

### SetPersistOnFailureNil

`func (o *ProvisioningRequest) SetPersistOnFailureNil(b bool)`

 SetPersistOnFailureNil sets the value for PersistOnFailure to be an explicit nil

### UnsetPersistOnFailure
`func (o *ProvisioningRequest) UnsetPersistOnFailure()`

UnsetPersistOnFailure ensures that no value is present for PersistOnFailure, not even an explicit nil
### GetPersona

`func (o *ProvisioningRequest) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *ProvisioningRequest) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *ProvisioningRequest) SetPersona(v string)`

SetPersona sets Persona field to given value.

### HasPersona

`func (o *ProvisioningRequest) HasPersona() bool`

HasPersona returns a boolean if a field has been set.

### SetPersonaNil

`func (o *ProvisioningRequest) SetPersonaNil(b bool)`

 SetPersonaNil sets the value for Persona to be an explicit nil

### UnsetPersona
`func (o *ProvisioningRequest) UnsetPersona()`

UnsetPersona ensures that no value is present for Persona, not even an explicit nil
### GetPipeline

`func (o *ProvisioningRequest) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *ProvisioningRequest) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *ProvisioningRequest) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *ProvisioningRequest) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineBranch

`func (o *ProvisioningRequest) GetPipelineBranch() string`

GetPipelineBranch returns the PipelineBranch field if non-nil, zero value otherwise.

### GetPipelineBranchOk

`func (o *ProvisioningRequest) GetPipelineBranchOk() (*string, bool)`

GetPipelineBranchOk returns a tuple with the PipelineBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineBranch

`func (o *ProvisioningRequest) SetPipelineBranch(v string)`

SetPipelineBranch sets PipelineBranch field to given value.

### HasPipelineBranch

`func (o *ProvisioningRequest) HasPipelineBranch() bool`

HasPipelineBranch returns a boolean if a field has been set.

### SetPipelineBranchNil

`func (o *ProvisioningRequest) SetPipelineBranchNil(b bool)`

 SetPipelineBranchNil sets the value for PipelineBranch to be an explicit nil

### UnsetPipelineBranch
`func (o *ProvisioningRequest) UnsetPipelineBranch()`

UnsetPipelineBranch ensures that no value is present for PipelineBranch, not even an explicit nil
### GetPlatform

`func (o *ProvisioningRequest) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ProvisioningRequest) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ProvisioningRequest) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ProvisioningRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ProvisioningRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ProvisioningRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetProviderName

`func (o *ProvisioningRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProvisioningRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProvisioningRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProvisioningRequest) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProvisioningRequest) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProvisioningRequest) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetRegion

`func (o *ProvisioningRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProvisioningRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProvisioningRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProvisioningRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ProvisioningRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ProvisioningRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRequestId

`func (o *ProvisioningRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProvisioningRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProvisioningRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRuntime

`func (o *ProvisioningRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ProvisioningRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ProvisioningRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ProvisioningRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *ProvisioningRequest) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *ProvisioningRequest) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetRuntimeBranch

`func (o *ProvisioningRequest) GetRuntimeBranch() string`

GetRuntimeBranch returns the RuntimeBranch field if non-nil, zero value otherwise.

### GetRuntimeBranchOk

`func (o *ProvisioningRequest) GetRuntimeBranchOk() (*string, bool)`

GetRuntimeBranchOk returns a tuple with the RuntimeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeBranch

`func (o *ProvisioningRequest) SetRuntimeBranch(v string)`

SetRuntimeBranch sets RuntimeBranch field to given value.

### HasRuntimeBranch

`func (o *ProvisioningRequest) HasRuntimeBranch() bool`

HasRuntimeBranch returns a boolean if a field has been set.

### SetRuntimeBranchNil

`func (o *ProvisioningRequest) SetRuntimeBranchNil(b bool)`

 SetRuntimeBranchNil sets the value for RuntimeBranch to be an explicit nil

### UnsetRuntimeBranch
`func (o *ProvisioningRequest) UnsetRuntimeBranch()`

UnsetRuntimeBranch ensures that no value is present for RuntimeBranch, not even an explicit nil
### GetRuntimeCnsAddonPack

`func (o *ProvisioningRequest) GetRuntimeCnsAddonPack() bool`

GetRuntimeCnsAddonPack returns the RuntimeCnsAddonPack field if non-nil, zero value otherwise.

### GetRuntimeCnsAddonPackOk

`func (o *ProvisioningRequest) GetRuntimeCnsAddonPackOk() (*bool, bool)`

GetRuntimeCnsAddonPackOk returns a tuple with the RuntimeCnsAddonPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsAddonPack

`func (o *ProvisioningRequest) SetRuntimeCnsAddonPack(v bool)`

SetRuntimeCnsAddonPack sets RuntimeCnsAddonPack field to given value.

### HasRuntimeCnsAddonPack

`func (o *ProvisioningRequest) HasRuntimeCnsAddonPack() bool`

HasRuntimeCnsAddonPack returns a boolean if a field has been set.

### SetRuntimeCnsAddonPackNil

`func (o *ProvisioningRequest) SetRuntimeCnsAddonPackNil(b bool)`

 SetRuntimeCnsAddonPackNil sets the value for RuntimeCnsAddonPack to be an explicit nil

### UnsetRuntimeCnsAddonPack
`func (o *ProvisioningRequest) UnsetRuntimeCnsAddonPack()`

UnsetRuntimeCnsAddonPack ensures that no value is present for RuntimeCnsAddonPack, not even an explicit nil
### GetRuntimeCnsDocker

`func (o *ProvisioningRequest) GetRuntimeCnsDocker() bool`

GetRuntimeCnsDocker returns the RuntimeCnsDocker field if non-nil, zero value otherwise.

### GetRuntimeCnsDockerOk

`func (o *ProvisioningRequest) GetRuntimeCnsDockerOk() (*bool, bool)`

GetRuntimeCnsDockerOk returns a tuple with the RuntimeCnsDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDocker

`func (o *ProvisioningRequest) SetRuntimeCnsDocker(v bool)`

SetRuntimeCnsDocker sets RuntimeCnsDocker field to given value.

### HasRuntimeCnsDocker

`func (o *ProvisioningRequest) HasRuntimeCnsDocker() bool`

HasRuntimeCnsDocker returns a boolean if a field has been set.

### SetRuntimeCnsDockerNil

`func (o *ProvisioningRequest) SetRuntimeCnsDockerNil(b bool)`

 SetRuntimeCnsDockerNil sets the value for RuntimeCnsDocker to be an explicit nil

### UnsetRuntimeCnsDocker
`func (o *ProvisioningRequest) UnsetRuntimeCnsDocker()`

UnsetRuntimeCnsDocker ensures that no value is present for RuntimeCnsDocker, not even an explicit nil
### GetRuntimeCnsDriverVersion

`func (o *ProvisioningRequest) GetRuntimeCnsDriverVersion() string`

GetRuntimeCnsDriverVersion returns the RuntimeCnsDriverVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsDriverVersionOk

`func (o *ProvisioningRequest) GetRuntimeCnsDriverVersionOk() (*string, bool)`

GetRuntimeCnsDriverVersionOk returns a tuple with the RuntimeCnsDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDriverVersion

`func (o *ProvisioningRequest) SetRuntimeCnsDriverVersion(v string)`

SetRuntimeCnsDriverVersion sets RuntimeCnsDriverVersion field to given value.

### HasRuntimeCnsDriverVersion

`func (o *ProvisioningRequest) HasRuntimeCnsDriverVersion() bool`

HasRuntimeCnsDriverVersion returns a boolean if a field has been set.

### SetRuntimeCnsDriverVersionNil

`func (o *ProvisioningRequest) SetRuntimeCnsDriverVersionNil(b bool)`

 SetRuntimeCnsDriverVersionNil sets the value for RuntimeCnsDriverVersion to be an explicit nil

### UnsetRuntimeCnsDriverVersion
`func (o *ProvisioningRequest) UnsetRuntimeCnsDriverVersion()`

UnsetRuntimeCnsDriverVersion ensures that no value is present for RuntimeCnsDriverVersion, not even an explicit nil
### GetRuntimeCnsK8s

`func (o *ProvisioningRequest) GetRuntimeCnsK8s() bool`

GetRuntimeCnsK8s returns the RuntimeCnsK8s field if non-nil, zero value otherwise.

### GetRuntimeCnsK8sOk

`func (o *ProvisioningRequest) GetRuntimeCnsK8sOk() (*bool, bool)`

GetRuntimeCnsK8sOk returns a tuple with the RuntimeCnsK8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsK8s

`func (o *ProvisioningRequest) SetRuntimeCnsK8s(v bool)`

SetRuntimeCnsK8s sets RuntimeCnsK8s field to given value.

### HasRuntimeCnsK8s

`func (o *ProvisioningRequest) HasRuntimeCnsK8s() bool`

HasRuntimeCnsK8s returns a boolean if a field has been set.

### SetRuntimeCnsK8sNil

`func (o *ProvisioningRequest) SetRuntimeCnsK8sNil(b bool)`

 SetRuntimeCnsK8sNil sets the value for RuntimeCnsK8s to be an explicit nil

### UnsetRuntimeCnsK8s
`func (o *ProvisioningRequest) UnsetRuntimeCnsK8s()`

UnsetRuntimeCnsK8s ensures that no value is present for RuntimeCnsK8s, not even an explicit nil
### GetRuntimeCnsNvidiaDriver

`func (o *ProvisioningRequest) GetRuntimeCnsNvidiaDriver() bool`

GetRuntimeCnsNvidiaDriver returns the RuntimeCnsNvidiaDriver field if non-nil, zero value otherwise.

### GetRuntimeCnsNvidiaDriverOk

`func (o *ProvisioningRequest) GetRuntimeCnsNvidiaDriverOk() (*bool, bool)`

GetRuntimeCnsNvidiaDriverOk returns a tuple with the RuntimeCnsNvidiaDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsNvidiaDriver

`func (o *ProvisioningRequest) SetRuntimeCnsNvidiaDriver(v bool)`

SetRuntimeCnsNvidiaDriver sets RuntimeCnsNvidiaDriver field to given value.

### HasRuntimeCnsNvidiaDriver

`func (o *ProvisioningRequest) HasRuntimeCnsNvidiaDriver() bool`

HasRuntimeCnsNvidiaDriver returns a boolean if a field has been set.

### SetRuntimeCnsNvidiaDriverNil

`func (o *ProvisioningRequest) SetRuntimeCnsNvidiaDriverNil(b bool)`

 SetRuntimeCnsNvidiaDriverNil sets the value for RuntimeCnsNvidiaDriver to be an explicit nil

### UnsetRuntimeCnsNvidiaDriver
`func (o *ProvisioningRequest) UnsetRuntimeCnsNvidiaDriver()`

UnsetRuntimeCnsNvidiaDriver ensures that no value is present for RuntimeCnsNvidiaDriver, not even an explicit nil
### GetRuntimeCnsVersion

`func (o *ProvisioningRequest) GetRuntimeCnsVersion() string`

GetRuntimeCnsVersion returns the RuntimeCnsVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsVersionOk

`func (o *ProvisioningRequest) GetRuntimeCnsVersionOk() (*string, bool)`

GetRuntimeCnsVersionOk returns a tuple with the RuntimeCnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsVersion

`func (o *ProvisioningRequest) SetRuntimeCnsVersion(v string)`

SetRuntimeCnsVersion sets RuntimeCnsVersion field to given value.

### HasRuntimeCnsVersion

`func (o *ProvisioningRequest) HasRuntimeCnsVersion() bool`

HasRuntimeCnsVersion returns a boolean if a field has been set.

### SetRuntimeCnsVersionNil

`func (o *ProvisioningRequest) SetRuntimeCnsVersionNil(b bool)`

 SetRuntimeCnsVersionNil sets the value for RuntimeCnsVersion to be an explicit nil

### UnsetRuntimeCnsVersion
`func (o *ProvisioningRequest) UnsetRuntimeCnsVersion()`

UnsetRuntimeCnsVersion ensures that no value is present for RuntimeCnsVersion, not even an explicit nil
### GetRuntimeMig

`func (o *ProvisioningRequest) GetRuntimeMig() bool`

GetRuntimeMig returns the RuntimeMig field if non-nil, zero value otherwise.

### GetRuntimeMigOk

`func (o *ProvisioningRequest) GetRuntimeMigOk() (*bool, bool)`

GetRuntimeMigOk returns a tuple with the RuntimeMig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMig

`func (o *ProvisioningRequest) SetRuntimeMig(v bool)`

SetRuntimeMig sets RuntimeMig field to given value.

### HasRuntimeMig

`func (o *ProvisioningRequest) HasRuntimeMig() bool`

HasRuntimeMig returns a boolean if a field has been set.

### SetRuntimeMigNil

`func (o *ProvisioningRequest) SetRuntimeMigNil(b bool)`

 SetRuntimeMigNil sets the value for RuntimeMig to be an explicit nil

### UnsetRuntimeMig
`func (o *ProvisioningRequest) UnsetRuntimeMig()`

UnsetRuntimeMig ensures that no value is present for RuntimeMig, not even an explicit nil
### GetRuntimeMigProfile

`func (o *ProvisioningRequest) GetRuntimeMigProfile() string`

GetRuntimeMigProfile returns the RuntimeMigProfile field if non-nil, zero value otherwise.

### GetRuntimeMigProfileOk

`func (o *ProvisioningRequest) GetRuntimeMigProfileOk() (*string, bool)`

GetRuntimeMigProfileOk returns a tuple with the RuntimeMigProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMigProfile

`func (o *ProvisioningRequest) SetRuntimeMigProfile(v string)`

SetRuntimeMigProfile sets RuntimeMigProfile field to given value.

### HasRuntimeMigProfile

`func (o *ProvisioningRequest) HasRuntimeMigProfile() bool`

HasRuntimeMigProfile returns a boolean if a field has been set.

### SetRuntimeMigProfileNil

`func (o *ProvisioningRequest) SetRuntimeMigProfileNil(b bool)`

 SetRuntimeMigProfileNil sets the value for RuntimeMigProfile to be an explicit nil

### UnsetRuntimeMigProfile
`func (o *ProvisioningRequest) UnsetRuntimeMigProfile()`

UnsetRuntimeMigProfile ensures that no value is present for RuntimeMigProfile, not even an explicit nil
### GetRuntimeUrl

`func (o *ProvisioningRequest) GetRuntimeUrl() string`

GetRuntimeUrl returns the RuntimeUrl field if non-nil, zero value otherwise.

### GetRuntimeUrlOk

`func (o *ProvisioningRequest) GetRuntimeUrlOk() (*string, bool)`

GetRuntimeUrlOk returns a tuple with the RuntimeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUrl

`func (o *ProvisioningRequest) SetRuntimeUrl(v string)`

SetRuntimeUrl sets RuntimeUrl field to given value.

### HasRuntimeUrl

`func (o *ProvisioningRequest) HasRuntimeUrl() bool`

HasRuntimeUrl returns a boolean if a field has been set.

### SetRuntimeUrlNil

`func (o *ProvisioningRequest) SetRuntimeUrlNil(b bool)`

 SetRuntimeUrlNil sets the value for RuntimeUrl to be an explicit nil

### UnsetRuntimeUrl
`func (o *ProvisioningRequest) UnsetRuntimeUrl()`

UnsetRuntimeUrl ensures that no value is present for RuntimeUrl, not even an explicit nil
### GetWorkshop

`func (o *ProvisioningRequest) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *ProvisioningRequest) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *ProvisioningRequest) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *ProvisioningRequest) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### SetWorkshopNil

`func (o *ProvisioningRequest) SetWorkshopNil(b bool)`

 SetWorkshopNil sets the value for Workshop to be an explicit nil

### UnsetWorkshop
`func (o *ProvisioningRequest) UnsetWorkshop()`

UnsetWorkshop ensures that no value is present for Workshop, not even an explicit nil
### GetWorkshopId

`func (o *ProvisioningRequest) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *ProvisioningRequest) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *ProvisioningRequest) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *ProvisioningRequest) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### SetWorkshopIdNil

`func (o *ProvisioningRequest) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *ProvisioningRequest) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil
### GetWorkshopOverridePassword

`func (o *ProvisioningRequest) GetWorkshopOverridePassword() string`

GetWorkshopOverridePassword returns the WorkshopOverridePassword field if non-nil, zero value otherwise.

### GetWorkshopOverridePasswordOk

`func (o *ProvisioningRequest) GetWorkshopOverridePasswordOk() (*string, bool)`

GetWorkshopOverridePasswordOk returns a tuple with the WorkshopOverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopOverridePassword

`func (o *ProvisioningRequest) SetWorkshopOverridePassword(v string)`

SetWorkshopOverridePassword sets WorkshopOverridePassword field to given value.

### HasWorkshopOverridePassword

`func (o *ProvisioningRequest) HasWorkshopOverridePassword() bool`

HasWorkshopOverridePassword returns a boolean if a field has been set.

### SetWorkshopOverridePasswordNil

`func (o *ProvisioningRequest) SetWorkshopOverridePasswordNil(b bool)`

 SetWorkshopOverridePasswordNil sets the value for WorkshopOverridePassword to be an explicit nil

### UnsetWorkshopOverridePassword
`func (o *ProvisioningRequest) UnsetWorkshopOverridePassword()`

UnsetWorkshopOverridePassword ensures that no value is present for WorkshopOverridePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


