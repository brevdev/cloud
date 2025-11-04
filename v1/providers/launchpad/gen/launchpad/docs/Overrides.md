# Overrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionOperatingSystem** | Pointer to **NullableString** | Override bastion operating system provisioned and/or configured by Liftoff | [optional] 
**Cluster** | Pointer to **NullableString** | Target a specific cluster for provisioning | [optional] 
**CollectionBranch** | Pointer to **NullableString** | Override the Ansible collection branch initialized within the pipeline | [optional] 
**ExperienceBranch** | Pointer to **NullableString** | Override the experience branch | [optional] 
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

### NewOverrides

`func NewOverrides() *Overrides`

NewOverrides instantiates a new Overrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverridesWithDefaults

`func NewOverridesWithDefaults() *Overrides`

NewOverridesWithDefaults instantiates a new Overrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionOperatingSystem

`func (o *Overrides) GetBastionOperatingSystem() string`

GetBastionOperatingSystem returns the BastionOperatingSystem field if non-nil, zero value otherwise.

### GetBastionOperatingSystemOk

`func (o *Overrides) GetBastionOperatingSystemOk() (*string, bool)`

GetBastionOperatingSystemOk returns a tuple with the BastionOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionOperatingSystem

`func (o *Overrides) SetBastionOperatingSystem(v string)`

SetBastionOperatingSystem sets BastionOperatingSystem field to given value.

### HasBastionOperatingSystem

`func (o *Overrides) HasBastionOperatingSystem() bool`

HasBastionOperatingSystem returns a boolean if a field has been set.

### SetBastionOperatingSystemNil

`func (o *Overrides) SetBastionOperatingSystemNil(b bool)`

 SetBastionOperatingSystemNil sets the value for BastionOperatingSystem to be an explicit nil

### UnsetBastionOperatingSystem
`func (o *Overrides) UnsetBastionOperatingSystem()`

UnsetBastionOperatingSystem ensures that no value is present for BastionOperatingSystem, not even an explicit nil
### GetCluster

`func (o *Overrides) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Overrides) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Overrides) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Overrides) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *Overrides) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *Overrides) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetCollectionBranch

`func (o *Overrides) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *Overrides) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *Overrides) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *Overrides) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### SetCollectionBranchNil

`func (o *Overrides) SetCollectionBranchNil(b bool)`

 SetCollectionBranchNil sets the value for CollectionBranch to be an explicit nil

### UnsetCollectionBranch
`func (o *Overrides) UnsetCollectionBranch()`

UnsetCollectionBranch ensures that no value is present for CollectionBranch, not even an explicit nil
### GetExperienceBranch

`func (o *Overrides) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *Overrides) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *Overrides) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *Overrides) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### SetExperienceBranchNil

`func (o *Overrides) SetExperienceBranchNil(b bool)`

 SetExperienceBranchNil sets the value for ExperienceBranch to be an explicit nil

### UnsetExperienceBranch
`func (o *Overrides) UnsetExperienceBranch()`

UnsetExperienceBranch ensures that no value is present for ExperienceBranch, not even an explicit nil
### GetFlightcontrolRelease

`func (o *Overrides) GetFlightcontrolRelease() string`

GetFlightcontrolRelease returns the FlightcontrolRelease field if non-nil, zero value otherwise.

### GetFlightcontrolReleaseOk

`func (o *Overrides) GetFlightcontrolReleaseOk() (*string, bool)`

GetFlightcontrolReleaseOk returns a tuple with the FlightcontrolRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightcontrolRelease

`func (o *Overrides) SetFlightcontrolRelease(v string)`

SetFlightcontrolRelease sets FlightcontrolRelease field to given value.

### HasFlightcontrolRelease

`func (o *Overrides) HasFlightcontrolRelease() bool`

HasFlightcontrolRelease returns a boolean if a field has been set.

### SetFlightcontrolReleaseNil

`func (o *Overrides) SetFlightcontrolReleaseNil(b bool)`

 SetFlightcontrolReleaseNil sets the value for FlightcontrolRelease to be an explicit nil

### UnsetFlightcontrolRelease
`func (o *Overrides) UnsetFlightcontrolRelease()`

UnsetFlightcontrolRelease ensures that no value is present for FlightcontrolRelease, not even an explicit nil
### GetGarageId

`func (o *Overrides) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *Overrides) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *Overrides) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *Overrides) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *Overrides) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *Overrides) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGcBranch

`func (o *Overrides) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *Overrides) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *Overrides) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *Overrides) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### SetGcBranchNil

`func (o *Overrides) SetGcBranchNil(b bool)`

 SetGcBranchNil sets the value for GcBranch to be an explicit nil

### UnsetGcBranch
`func (o *Overrides) UnsetGcBranch()`

UnsetGcBranch ensures that no value is present for GcBranch, not even an explicit nil
### GetGpuAlias

`func (o *Overrides) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *Overrides) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *Overrides) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *Overrides) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### SetGpuAliasNil

`func (o *Overrides) SetGpuAliasNil(b bool)`

 SetGpuAliasNil sets the value for GpuAlias to be an explicit nil

### UnsetGpuAlias
`func (o *Overrides) UnsetGpuAlias()`

UnsetGpuAlias ensures that no value is present for GpuAlias, not even an explicit nil
### GetGpuCount

`func (o *Overrides) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Overrides) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Overrides) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *Overrides) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *Overrides) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *Overrides) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *Overrides) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *Overrides) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### SetGpuModelNil

`func (o *Overrides) SetGpuModelNil(b bool)`

 SetGpuModelNil sets the value for GpuModel to be an explicit nil

### UnsetGpuModel
`func (o *Overrides) UnsetGpuModel()`

UnsetGpuModel ensures that no value is present for GpuModel, not even an explicit nil
### GetGpuOsName

`func (o *Overrides) GetGpuOsName() string`

GetGpuOsName returns the GpuOsName field if non-nil, zero value otherwise.

### GetGpuOsNameOk

`func (o *Overrides) GetGpuOsNameOk() (*string, bool)`

GetGpuOsNameOk returns a tuple with the GpuOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsName

`func (o *Overrides) SetGpuOsName(v string)`

SetGpuOsName sets GpuOsName field to given value.

### HasGpuOsName

`func (o *Overrides) HasGpuOsName() bool`

HasGpuOsName returns a boolean if a field has been set.

### SetGpuOsNameNil

`func (o *Overrides) SetGpuOsNameNil(b bool)`

 SetGpuOsNameNil sets the value for GpuOsName to be an explicit nil

### UnsetGpuOsName
`func (o *Overrides) UnsetGpuOsName()`

UnsetGpuOsName ensures that no value is present for GpuOsName, not even an explicit nil
### GetGpuOsRelease

`func (o *Overrides) GetGpuOsRelease() string`

GetGpuOsRelease returns the GpuOsRelease field if non-nil, zero value otherwise.

### GetGpuOsReleaseOk

`func (o *Overrides) GetGpuOsReleaseOk() (*string, bool)`

GetGpuOsReleaseOk returns a tuple with the GpuOsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsRelease

`func (o *Overrides) SetGpuOsRelease(v string)`

SetGpuOsRelease sets GpuOsRelease field to given value.

### HasGpuOsRelease

`func (o *Overrides) HasGpuOsRelease() bool`

HasGpuOsRelease returns a boolean if a field has been set.

### SetGpuOsReleaseNil

`func (o *Overrides) SetGpuOsReleaseNil(b bool)`

 SetGpuOsReleaseNil sets the value for GpuOsRelease to be an explicit nil

### UnsetGpuOsRelease
`func (o *Overrides) UnsetGpuOsRelease()`

UnsetGpuOsRelease ensures that no value is present for GpuOsRelease, not even an explicit nil
### GetGpuOsVersion

`func (o *Overrides) GetGpuOsVersion() string`

GetGpuOsVersion returns the GpuOsVersion field if non-nil, zero value otherwise.

### GetGpuOsVersionOk

`func (o *Overrides) GetGpuOsVersionOk() (*string, bool)`

GetGpuOsVersionOk returns a tuple with the GpuOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsVersion

`func (o *Overrides) SetGpuOsVersion(v string)`

SetGpuOsVersion sets GpuOsVersion field to given value.

### HasGpuOsVersion

`func (o *Overrides) HasGpuOsVersion() bool`

HasGpuOsVersion returns a boolean if a field has been set.

### SetGpuOsVersionNil

`func (o *Overrides) SetGpuOsVersionNil(b bool)`

 SetGpuOsVersionNil sets the value for GpuOsVersion to be an explicit nil

### UnsetGpuOsVersion
`func (o *Overrides) UnsetGpuOsVersion()`

UnsetGpuOsVersion ensures that no value is present for GpuOsVersion, not even an explicit nil
### GetMinGpuCount

`func (o *Overrides) GetMinGpuCount() int32`

GetMinGpuCount returns the MinGpuCount field if non-nil, zero value otherwise.

### GetMinGpuCountOk

`func (o *Overrides) GetMinGpuCountOk() (*int32, bool)`

GetMinGpuCountOk returns a tuple with the MinGpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGpuCount

`func (o *Overrides) SetMinGpuCount(v int32)`

SetMinGpuCount sets MinGpuCount field to given value.

### HasMinGpuCount

`func (o *Overrides) HasMinGpuCount() bool`

HasMinGpuCount returns a boolean if a field has been set.

### GetNodeCount

`func (o *Overrides) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Overrides) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Overrides) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Overrides) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetOemName

`func (o *Overrides) GetOemName() string`

GetOemName returns the OemName field if non-nil, zero value otherwise.

### GetOemNameOk

`func (o *Overrides) GetOemNameOk() (*string, bool)`

GetOemNameOk returns a tuple with the OemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemName

`func (o *Overrides) SetOemName(v string)`

SetOemName sets OemName field to given value.

### HasOemName

`func (o *Overrides) HasOemName() bool`

HasOemName returns a boolean if a field has been set.

### SetOemNameNil

`func (o *Overrides) SetOemNameNil(b bool)`

 SetOemNameNil sets the value for OemName to be an explicit nil

### UnsetOemName
`func (o *Overrides) UnsetOemName()`

UnsetOemName ensures that no value is present for OemName, not even an explicit nil
### GetPersistOnFailure

`func (o *Overrides) GetPersistOnFailure() bool`

GetPersistOnFailure returns the PersistOnFailure field if non-nil, zero value otherwise.

### GetPersistOnFailureOk

`func (o *Overrides) GetPersistOnFailureOk() (*bool, bool)`

GetPersistOnFailureOk returns a tuple with the PersistOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOnFailure

`func (o *Overrides) SetPersistOnFailure(v bool)`

SetPersistOnFailure sets PersistOnFailure field to given value.

### HasPersistOnFailure

`func (o *Overrides) HasPersistOnFailure() bool`

HasPersistOnFailure returns a boolean if a field has been set.

### SetPersistOnFailureNil

`func (o *Overrides) SetPersistOnFailureNil(b bool)`

 SetPersistOnFailureNil sets the value for PersistOnFailure to be an explicit nil

### UnsetPersistOnFailure
`func (o *Overrides) UnsetPersistOnFailure()`

UnsetPersistOnFailure ensures that no value is present for PersistOnFailure, not even an explicit nil
### GetPersona

`func (o *Overrides) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *Overrides) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *Overrides) SetPersona(v string)`

SetPersona sets Persona field to given value.

### HasPersona

`func (o *Overrides) HasPersona() bool`

HasPersona returns a boolean if a field has been set.

### SetPersonaNil

`func (o *Overrides) SetPersonaNil(b bool)`

 SetPersonaNil sets the value for Persona to be an explicit nil

### UnsetPersona
`func (o *Overrides) UnsetPersona()`

UnsetPersona ensures that no value is present for Persona, not even an explicit nil
### GetPipeline

`func (o *Overrides) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *Overrides) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *Overrides) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *Overrides) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineBranch

`func (o *Overrides) GetPipelineBranch() string`

GetPipelineBranch returns the PipelineBranch field if non-nil, zero value otherwise.

### GetPipelineBranchOk

`func (o *Overrides) GetPipelineBranchOk() (*string, bool)`

GetPipelineBranchOk returns a tuple with the PipelineBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineBranch

`func (o *Overrides) SetPipelineBranch(v string)`

SetPipelineBranch sets PipelineBranch field to given value.

### HasPipelineBranch

`func (o *Overrides) HasPipelineBranch() bool`

HasPipelineBranch returns a boolean if a field has been set.

### SetPipelineBranchNil

`func (o *Overrides) SetPipelineBranchNil(b bool)`

 SetPipelineBranchNil sets the value for PipelineBranch to be an explicit nil

### UnsetPipelineBranch
`func (o *Overrides) UnsetPipelineBranch()`

UnsetPipelineBranch ensures that no value is present for PipelineBranch, not even an explicit nil
### GetPlatform

`func (o *Overrides) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Overrides) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Overrides) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Overrides) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Overrides) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Overrides) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetProviderName

`func (o *Overrides) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *Overrides) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *Overrides) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *Overrides) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *Overrides) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *Overrides) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetRegion

`func (o *Overrides) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Overrides) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Overrides) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Overrides) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Overrides) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Overrides) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRuntime

`func (o *Overrides) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Overrides) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Overrides) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Overrides) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *Overrides) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *Overrides) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetRuntimeBranch

`func (o *Overrides) GetRuntimeBranch() string`

GetRuntimeBranch returns the RuntimeBranch field if non-nil, zero value otherwise.

### GetRuntimeBranchOk

`func (o *Overrides) GetRuntimeBranchOk() (*string, bool)`

GetRuntimeBranchOk returns a tuple with the RuntimeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeBranch

`func (o *Overrides) SetRuntimeBranch(v string)`

SetRuntimeBranch sets RuntimeBranch field to given value.

### HasRuntimeBranch

`func (o *Overrides) HasRuntimeBranch() bool`

HasRuntimeBranch returns a boolean if a field has been set.

### SetRuntimeBranchNil

`func (o *Overrides) SetRuntimeBranchNil(b bool)`

 SetRuntimeBranchNil sets the value for RuntimeBranch to be an explicit nil

### UnsetRuntimeBranch
`func (o *Overrides) UnsetRuntimeBranch()`

UnsetRuntimeBranch ensures that no value is present for RuntimeBranch, not even an explicit nil
### GetRuntimeCnsAddonPack

`func (o *Overrides) GetRuntimeCnsAddonPack() bool`

GetRuntimeCnsAddonPack returns the RuntimeCnsAddonPack field if non-nil, zero value otherwise.

### GetRuntimeCnsAddonPackOk

`func (o *Overrides) GetRuntimeCnsAddonPackOk() (*bool, bool)`

GetRuntimeCnsAddonPackOk returns a tuple with the RuntimeCnsAddonPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsAddonPack

`func (o *Overrides) SetRuntimeCnsAddonPack(v bool)`

SetRuntimeCnsAddonPack sets RuntimeCnsAddonPack field to given value.

### HasRuntimeCnsAddonPack

`func (o *Overrides) HasRuntimeCnsAddonPack() bool`

HasRuntimeCnsAddonPack returns a boolean if a field has been set.

### SetRuntimeCnsAddonPackNil

`func (o *Overrides) SetRuntimeCnsAddonPackNil(b bool)`

 SetRuntimeCnsAddonPackNil sets the value for RuntimeCnsAddonPack to be an explicit nil

### UnsetRuntimeCnsAddonPack
`func (o *Overrides) UnsetRuntimeCnsAddonPack()`

UnsetRuntimeCnsAddonPack ensures that no value is present for RuntimeCnsAddonPack, not even an explicit nil
### GetRuntimeCnsDocker

`func (o *Overrides) GetRuntimeCnsDocker() bool`

GetRuntimeCnsDocker returns the RuntimeCnsDocker field if non-nil, zero value otherwise.

### GetRuntimeCnsDockerOk

`func (o *Overrides) GetRuntimeCnsDockerOk() (*bool, bool)`

GetRuntimeCnsDockerOk returns a tuple with the RuntimeCnsDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDocker

`func (o *Overrides) SetRuntimeCnsDocker(v bool)`

SetRuntimeCnsDocker sets RuntimeCnsDocker field to given value.

### HasRuntimeCnsDocker

`func (o *Overrides) HasRuntimeCnsDocker() bool`

HasRuntimeCnsDocker returns a boolean if a field has been set.

### SetRuntimeCnsDockerNil

`func (o *Overrides) SetRuntimeCnsDockerNil(b bool)`

 SetRuntimeCnsDockerNil sets the value for RuntimeCnsDocker to be an explicit nil

### UnsetRuntimeCnsDocker
`func (o *Overrides) UnsetRuntimeCnsDocker()`

UnsetRuntimeCnsDocker ensures that no value is present for RuntimeCnsDocker, not even an explicit nil
### GetRuntimeCnsDriverVersion

`func (o *Overrides) GetRuntimeCnsDriverVersion() string`

GetRuntimeCnsDriverVersion returns the RuntimeCnsDriverVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsDriverVersionOk

`func (o *Overrides) GetRuntimeCnsDriverVersionOk() (*string, bool)`

GetRuntimeCnsDriverVersionOk returns a tuple with the RuntimeCnsDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDriverVersion

`func (o *Overrides) SetRuntimeCnsDriverVersion(v string)`

SetRuntimeCnsDriverVersion sets RuntimeCnsDriverVersion field to given value.

### HasRuntimeCnsDriverVersion

`func (o *Overrides) HasRuntimeCnsDriverVersion() bool`

HasRuntimeCnsDriverVersion returns a boolean if a field has been set.

### SetRuntimeCnsDriverVersionNil

`func (o *Overrides) SetRuntimeCnsDriverVersionNil(b bool)`

 SetRuntimeCnsDriverVersionNil sets the value for RuntimeCnsDriverVersion to be an explicit nil

### UnsetRuntimeCnsDriverVersion
`func (o *Overrides) UnsetRuntimeCnsDriverVersion()`

UnsetRuntimeCnsDriverVersion ensures that no value is present for RuntimeCnsDriverVersion, not even an explicit nil
### GetRuntimeCnsK8s

`func (o *Overrides) GetRuntimeCnsK8s() bool`

GetRuntimeCnsK8s returns the RuntimeCnsK8s field if non-nil, zero value otherwise.

### GetRuntimeCnsK8sOk

`func (o *Overrides) GetRuntimeCnsK8sOk() (*bool, bool)`

GetRuntimeCnsK8sOk returns a tuple with the RuntimeCnsK8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsK8s

`func (o *Overrides) SetRuntimeCnsK8s(v bool)`

SetRuntimeCnsK8s sets RuntimeCnsK8s field to given value.

### HasRuntimeCnsK8s

`func (o *Overrides) HasRuntimeCnsK8s() bool`

HasRuntimeCnsK8s returns a boolean if a field has been set.

### SetRuntimeCnsK8sNil

`func (o *Overrides) SetRuntimeCnsK8sNil(b bool)`

 SetRuntimeCnsK8sNil sets the value for RuntimeCnsK8s to be an explicit nil

### UnsetRuntimeCnsK8s
`func (o *Overrides) UnsetRuntimeCnsK8s()`

UnsetRuntimeCnsK8s ensures that no value is present for RuntimeCnsK8s, not even an explicit nil
### GetRuntimeCnsNvidiaDriver

`func (o *Overrides) GetRuntimeCnsNvidiaDriver() bool`

GetRuntimeCnsNvidiaDriver returns the RuntimeCnsNvidiaDriver field if non-nil, zero value otherwise.

### GetRuntimeCnsNvidiaDriverOk

`func (o *Overrides) GetRuntimeCnsNvidiaDriverOk() (*bool, bool)`

GetRuntimeCnsNvidiaDriverOk returns a tuple with the RuntimeCnsNvidiaDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsNvidiaDriver

`func (o *Overrides) SetRuntimeCnsNvidiaDriver(v bool)`

SetRuntimeCnsNvidiaDriver sets RuntimeCnsNvidiaDriver field to given value.

### HasRuntimeCnsNvidiaDriver

`func (o *Overrides) HasRuntimeCnsNvidiaDriver() bool`

HasRuntimeCnsNvidiaDriver returns a boolean if a field has been set.

### SetRuntimeCnsNvidiaDriverNil

`func (o *Overrides) SetRuntimeCnsNvidiaDriverNil(b bool)`

 SetRuntimeCnsNvidiaDriverNil sets the value for RuntimeCnsNvidiaDriver to be an explicit nil

### UnsetRuntimeCnsNvidiaDriver
`func (o *Overrides) UnsetRuntimeCnsNvidiaDriver()`

UnsetRuntimeCnsNvidiaDriver ensures that no value is present for RuntimeCnsNvidiaDriver, not even an explicit nil
### GetRuntimeCnsVersion

`func (o *Overrides) GetRuntimeCnsVersion() string`

GetRuntimeCnsVersion returns the RuntimeCnsVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsVersionOk

`func (o *Overrides) GetRuntimeCnsVersionOk() (*string, bool)`

GetRuntimeCnsVersionOk returns a tuple with the RuntimeCnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsVersion

`func (o *Overrides) SetRuntimeCnsVersion(v string)`

SetRuntimeCnsVersion sets RuntimeCnsVersion field to given value.

### HasRuntimeCnsVersion

`func (o *Overrides) HasRuntimeCnsVersion() bool`

HasRuntimeCnsVersion returns a boolean if a field has been set.

### SetRuntimeCnsVersionNil

`func (o *Overrides) SetRuntimeCnsVersionNil(b bool)`

 SetRuntimeCnsVersionNil sets the value for RuntimeCnsVersion to be an explicit nil

### UnsetRuntimeCnsVersion
`func (o *Overrides) UnsetRuntimeCnsVersion()`

UnsetRuntimeCnsVersion ensures that no value is present for RuntimeCnsVersion, not even an explicit nil
### GetRuntimeMig

`func (o *Overrides) GetRuntimeMig() bool`

GetRuntimeMig returns the RuntimeMig field if non-nil, zero value otherwise.

### GetRuntimeMigOk

`func (o *Overrides) GetRuntimeMigOk() (*bool, bool)`

GetRuntimeMigOk returns a tuple with the RuntimeMig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMig

`func (o *Overrides) SetRuntimeMig(v bool)`

SetRuntimeMig sets RuntimeMig field to given value.

### HasRuntimeMig

`func (o *Overrides) HasRuntimeMig() bool`

HasRuntimeMig returns a boolean if a field has been set.

### SetRuntimeMigNil

`func (o *Overrides) SetRuntimeMigNil(b bool)`

 SetRuntimeMigNil sets the value for RuntimeMig to be an explicit nil

### UnsetRuntimeMig
`func (o *Overrides) UnsetRuntimeMig()`

UnsetRuntimeMig ensures that no value is present for RuntimeMig, not even an explicit nil
### GetRuntimeMigProfile

`func (o *Overrides) GetRuntimeMigProfile() string`

GetRuntimeMigProfile returns the RuntimeMigProfile field if non-nil, zero value otherwise.

### GetRuntimeMigProfileOk

`func (o *Overrides) GetRuntimeMigProfileOk() (*string, bool)`

GetRuntimeMigProfileOk returns a tuple with the RuntimeMigProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMigProfile

`func (o *Overrides) SetRuntimeMigProfile(v string)`

SetRuntimeMigProfile sets RuntimeMigProfile field to given value.

### HasRuntimeMigProfile

`func (o *Overrides) HasRuntimeMigProfile() bool`

HasRuntimeMigProfile returns a boolean if a field has been set.

### SetRuntimeMigProfileNil

`func (o *Overrides) SetRuntimeMigProfileNil(b bool)`

 SetRuntimeMigProfileNil sets the value for RuntimeMigProfile to be an explicit nil

### UnsetRuntimeMigProfile
`func (o *Overrides) UnsetRuntimeMigProfile()`

UnsetRuntimeMigProfile ensures that no value is present for RuntimeMigProfile, not even an explicit nil
### GetRuntimeUrl

`func (o *Overrides) GetRuntimeUrl() string`

GetRuntimeUrl returns the RuntimeUrl field if non-nil, zero value otherwise.

### GetRuntimeUrlOk

`func (o *Overrides) GetRuntimeUrlOk() (*string, bool)`

GetRuntimeUrlOk returns a tuple with the RuntimeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUrl

`func (o *Overrides) SetRuntimeUrl(v string)`

SetRuntimeUrl sets RuntimeUrl field to given value.

### HasRuntimeUrl

`func (o *Overrides) HasRuntimeUrl() bool`

HasRuntimeUrl returns a boolean if a field has been set.

### SetRuntimeUrlNil

`func (o *Overrides) SetRuntimeUrlNil(b bool)`

 SetRuntimeUrlNil sets the value for RuntimeUrl to be an explicit nil

### UnsetRuntimeUrl
`func (o *Overrides) UnsetRuntimeUrl()`

UnsetRuntimeUrl ensures that no value is present for RuntimeUrl, not even an explicit nil
### GetWorkshop

`func (o *Overrides) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *Overrides) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *Overrides) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *Overrides) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### SetWorkshopNil

`func (o *Overrides) SetWorkshopNil(b bool)`

 SetWorkshopNil sets the value for Workshop to be an explicit nil

### UnsetWorkshop
`func (o *Overrides) UnsetWorkshop()`

UnsetWorkshop ensures that no value is present for Workshop, not even an explicit nil
### GetWorkshopId

`func (o *Overrides) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *Overrides) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *Overrides) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *Overrides) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### SetWorkshopIdNil

`func (o *Overrides) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *Overrides) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil
### GetWorkshopOverridePassword

`func (o *Overrides) GetWorkshopOverridePassword() string`

GetWorkshopOverridePassword returns the WorkshopOverridePassword field if non-nil, zero value otherwise.

### GetWorkshopOverridePasswordOk

`func (o *Overrides) GetWorkshopOverridePasswordOk() (*string, bool)`

GetWorkshopOverridePasswordOk returns a tuple with the WorkshopOverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopOverridePassword

`func (o *Overrides) SetWorkshopOverridePassword(v string)`

SetWorkshopOverridePassword sets WorkshopOverridePassword field to given value.

### HasWorkshopOverridePassword

`func (o *Overrides) HasWorkshopOverridePassword() bool`

HasWorkshopOverridePassword returns a boolean if a field has been set.

### SetWorkshopOverridePasswordNil

`func (o *Overrides) SetWorkshopOverridePasswordNil(b bool)`

 SetWorkshopOverridePasswordNil sets the value for WorkshopOverridePassword to be an explicit nil

### UnsetWorkshopOverridePassword
`func (o *Overrides) UnsetWorkshopOverridePassword()`

UnsetWorkshopOverridePassword ensures that no value is present for WorkshopOverridePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


