# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionOperatingSystem** | Pointer to **NullableString** | Override bastion operating system provisioned and/or configured by Liftoff | [optional] 
**CatalogId** | Pointer to **string** | Catalog ID of the experience provisioned to the cluster | [optional] 
**CatalogIdAlias** | Pointer to **string** | Catalog ID alias of the experience provisioned to the cluster | [optional] 
**Cluster** | Pointer to [**ClusterPipelineCluster**](ClusterPipelineCluster.md) |  | [optional] 
**CollectionBranch** | Pointer to **NullableString** | Override the Ansible collection branch initialized within the pipeline | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Experience** | Pointer to **NullableString** | The experience being deployed for use | [optional] [readonly] 
**ExperienceBranch** | Pointer to **NullableString** | Override the experience branch | [optional] 
**ExperienceId** | Pointer to **string** | UUID of the experience provisioned to the cluster | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FlightcontrolRelease** | Pointer to **NullableString** | Override the image tag used for Flight Control | [optional] 
**GarageId** | Pointer to **NullableString** | Require a cluster with nodes in the given garage | [optional] 
**GcBranch** | Pointer to **NullableString** | Override the default Ground Control branch | [optional] 
**GpuAlias** | Pointer to **NullableString** | Require a cluster with the given GPU alias | [optional] 
**GpuCount** | Pointer to **NullableInt32** | Require a cluster with the given number of GPUs | [optional] [default to 0]
**GpuModel** | Pointer to **NullableString** | Require a cluster with the given GPU model | [optional] 
**GpuOsName** | Pointer to **NullableString** | Override the GPU node operating system name | [optional] 
**GpuOsRelease** | Pointer to **NullableString** | Override the GPU node operating system release | [optional] 
**GpuOsVersion** | Pointer to **NullableString** | Override the GPU node operating system version | [optional] 
**Id** | **string** |  | [readonly] 
**IpAllowlist** | Pointer to **[]string** | Host IP addresses that should be allowed to access the deployment | [optional] 
**Lifetime** | Pointer to **NullableInt32** | Set expires_at value to be a given number of days from the current time. A value of 0 will cause a deployment to remain active indefinitely. | [optional] 
**MinGpuCount** | Pointer to **NullableInt32** | Require a cluster whose GPU count is greater than or equal to the given number | [optional] [default to 0]
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | Pointer to **NullableInt32** | Require a cluster with the given number of nodes | [optional] [default to 0]
**OemName** | Pointer to **NullableString** | Require a cluster manufactured by the given OEM name | [optional] 
**OrgName** | **string** | Requester&#39;s organization name | 
**Overrides** | [**Overrides**](Overrides.md) | Overriden values from the original deployment request | [readonly] 
**PersistOnFailure** | Pointer to **NullableBool** | Override the default cleanup/destroy behavior when a provisioning failure occurs | [optional] [default to false]
**Persona** | Pointer to **NullableString** | Override the defined persona in the experience | [optional] 
**Pipeline** | Pointer to **NullableInt32** | Override the pipeline ID that will be triggered for request fulfillment | [optional] [default to 0]
**PipelineBranch** | Pointer to **NullableString** | Override the default pipeline branch ref used when triggering a Fuselage pipeline | [optional] 
**Pipelines** | **[]string** |  | 
**Platform** | Pointer to [**NullablePlatformEnum**](PlatformEnum.md) |  | [optional] 
**Priority** | Pointer to [**PriorityEnum**](PriorityEnum.md) | Priority level for the request  * &#x60;p0&#x60; - p0 * &#x60;p1&#x60; - p1 * &#x60;p2&#x60; - p2 * &#x60;p3&#x60; - p3 | [optional] 
**ProviderName** | Pointer to **NullableString** | Require a cluster from the given provider name | [optional] 
**PublicKey** | Pointer to **NullableString** | The initial or administrative public key used during deployment creation. Additional keys can be authorized for access using the &#x60;ssh-keys&#x60; endpoint. | [optional] 
**Region** | Pointer to **NullableString** | Require a cluster located in the given region | [optional] 
**RequestId** | Pointer to **string** | Trial request ID (ex: TRY-1234) | [optional] 
**RequesterEmail** | **string** | Email address of the user requesting the experience | 
**RequesterName** | **string** | Name of the user requesting the experience | 
**RetryCount** | **int32** | Number of times the deployment has been retried | [readonly] 
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
**SalesCreatedDate** | Pointer to **NullableTime** |  | [optional] 
**SalesId** | Pointer to **NullableString** | Unique identifier for the requester&#39;s sales relationship | [optional] 
**SalesOwnerEmail** | Pointer to **NullableString** | Email address of the sales contact associated with the requester | [optional] 
**SalesOwnerName** | Pointer to **NullableString** | Name of the sales contact associated with the requester | [optional] 
**Services** | **[]string** |  | 
**SshPort** | **int32** |  | [readonly] 
**SshUser** | **string** |  | [readonly] 
**State** | [**DeploymentState**](DeploymentState.md) | Current state of the deployment  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | [readonly] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Workshop** | Pointer to **NullableBool** | Require a cluster whose workshop flag is set | [optional] [default to false]
**WorkshopId** | Pointer to **NullableString** | Require a cluster with the given workshop ID | [optional] 
**WorkshopOverridePassword** | Pointer to **NullableString** | Override the deployment&#39;s default authentication to use a static password. This is useful for workshops when you&#39;d like an identical password associated with a collection of environments. (LaunchPad Team only) | [optional] 

## Methods

### NewDeployment

`func NewDeployment(created time.Time, id string, modified time.Time, orgName string, overrides Overrides, pipelines []string, requesterEmail string, requesterName string, retryCount int32, services []string, sshPort int32, sshUser string, state DeploymentState, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionOperatingSystem

`func (o *Deployment) GetBastionOperatingSystem() string`

GetBastionOperatingSystem returns the BastionOperatingSystem field if non-nil, zero value otherwise.

### GetBastionOperatingSystemOk

`func (o *Deployment) GetBastionOperatingSystemOk() (*string, bool)`

GetBastionOperatingSystemOk returns a tuple with the BastionOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionOperatingSystem

`func (o *Deployment) SetBastionOperatingSystem(v string)`

SetBastionOperatingSystem sets BastionOperatingSystem field to given value.

### HasBastionOperatingSystem

`func (o *Deployment) HasBastionOperatingSystem() bool`

HasBastionOperatingSystem returns a boolean if a field has been set.

### SetBastionOperatingSystemNil

`func (o *Deployment) SetBastionOperatingSystemNil(b bool)`

 SetBastionOperatingSystemNil sets the value for BastionOperatingSystem to be an explicit nil

### UnsetBastionOperatingSystem
`func (o *Deployment) UnsetBastionOperatingSystem()`

UnsetBastionOperatingSystem ensures that no value is present for BastionOperatingSystem, not even an explicit nil
### GetCatalogId

`func (o *Deployment) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *Deployment) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *Deployment) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *Deployment) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogIdAlias

`func (o *Deployment) GetCatalogIdAlias() string`

GetCatalogIdAlias returns the CatalogIdAlias field if non-nil, zero value otherwise.

### GetCatalogIdAliasOk

`func (o *Deployment) GetCatalogIdAliasOk() (*string, bool)`

GetCatalogIdAliasOk returns a tuple with the CatalogIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdAlias

`func (o *Deployment) SetCatalogIdAlias(v string)`

SetCatalogIdAlias sets CatalogIdAlias field to given value.

### HasCatalogIdAlias

`func (o *Deployment) HasCatalogIdAlias() bool`

HasCatalogIdAlias returns a boolean if a field has been set.

### GetCluster

`func (o *Deployment) GetCluster() ClusterPipelineCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Deployment) GetClusterOk() (*ClusterPipelineCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Deployment) SetCluster(v ClusterPipelineCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Deployment) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCollectionBranch

`func (o *Deployment) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *Deployment) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *Deployment) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *Deployment) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### SetCollectionBranchNil

`func (o *Deployment) SetCollectionBranchNil(b bool)`

 SetCollectionBranchNil sets the value for CollectionBranch to be an explicit nil

### UnsetCollectionBranch
`func (o *Deployment) UnsetCollectionBranch()`

UnsetCollectionBranch ensures that no value is present for CollectionBranch, not even an explicit nil
### GetCreated

`func (o *Deployment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Deployment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Deployment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExperience

`func (o *Deployment) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *Deployment) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *Deployment) SetExperience(v string)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *Deployment) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *Deployment) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *Deployment) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetExperienceBranch

`func (o *Deployment) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *Deployment) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *Deployment) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *Deployment) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### SetExperienceBranchNil

`func (o *Deployment) SetExperienceBranchNil(b bool)`

 SetExperienceBranchNil sets the value for ExperienceBranch to be an explicit nil

### UnsetExperienceBranch
`func (o *Deployment) UnsetExperienceBranch()`

UnsetExperienceBranch ensures that no value is present for ExperienceBranch, not even an explicit nil
### GetExperienceId

`func (o *Deployment) GetExperienceId() string`

GetExperienceId returns the ExperienceId field if non-nil, zero value otherwise.

### GetExperienceIdOk

`func (o *Deployment) GetExperienceIdOk() (*string, bool)`

GetExperienceIdOk returns a tuple with the ExperienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceId

`func (o *Deployment) SetExperienceId(v string)`

SetExperienceId sets ExperienceId field to given value.

### HasExperienceId

`func (o *Deployment) HasExperienceId() bool`

HasExperienceId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Deployment) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Deployment) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Deployment) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Deployment) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *Deployment) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Deployment) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFlightcontrolRelease

`func (o *Deployment) GetFlightcontrolRelease() string`

GetFlightcontrolRelease returns the FlightcontrolRelease field if non-nil, zero value otherwise.

### GetFlightcontrolReleaseOk

`func (o *Deployment) GetFlightcontrolReleaseOk() (*string, bool)`

GetFlightcontrolReleaseOk returns a tuple with the FlightcontrolRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightcontrolRelease

`func (o *Deployment) SetFlightcontrolRelease(v string)`

SetFlightcontrolRelease sets FlightcontrolRelease field to given value.

### HasFlightcontrolRelease

`func (o *Deployment) HasFlightcontrolRelease() bool`

HasFlightcontrolRelease returns a boolean if a field has been set.

### SetFlightcontrolReleaseNil

`func (o *Deployment) SetFlightcontrolReleaseNil(b bool)`

 SetFlightcontrolReleaseNil sets the value for FlightcontrolRelease to be an explicit nil

### UnsetFlightcontrolRelease
`func (o *Deployment) UnsetFlightcontrolRelease()`

UnsetFlightcontrolRelease ensures that no value is present for FlightcontrolRelease, not even an explicit nil
### GetGarageId

`func (o *Deployment) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *Deployment) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *Deployment) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *Deployment) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *Deployment) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *Deployment) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGcBranch

`func (o *Deployment) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *Deployment) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *Deployment) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *Deployment) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### SetGcBranchNil

`func (o *Deployment) SetGcBranchNil(b bool)`

 SetGcBranchNil sets the value for GcBranch to be an explicit nil

### UnsetGcBranch
`func (o *Deployment) UnsetGcBranch()`

UnsetGcBranch ensures that no value is present for GcBranch, not even an explicit nil
### GetGpuAlias

`func (o *Deployment) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *Deployment) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *Deployment) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *Deployment) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### SetGpuAliasNil

`func (o *Deployment) SetGpuAliasNil(b bool)`

 SetGpuAliasNil sets the value for GpuAlias to be an explicit nil

### UnsetGpuAlias
`func (o *Deployment) UnsetGpuAlias()`

UnsetGpuAlias ensures that no value is present for GpuAlias, not even an explicit nil
### GetGpuCount

`func (o *Deployment) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Deployment) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Deployment) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *Deployment) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### SetGpuCountNil

`func (o *Deployment) SetGpuCountNil(b bool)`

 SetGpuCountNil sets the value for GpuCount to be an explicit nil

### UnsetGpuCount
`func (o *Deployment) UnsetGpuCount()`

UnsetGpuCount ensures that no value is present for GpuCount, not even an explicit nil
### GetGpuModel

`func (o *Deployment) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *Deployment) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *Deployment) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *Deployment) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### SetGpuModelNil

`func (o *Deployment) SetGpuModelNil(b bool)`

 SetGpuModelNil sets the value for GpuModel to be an explicit nil

### UnsetGpuModel
`func (o *Deployment) UnsetGpuModel()`

UnsetGpuModel ensures that no value is present for GpuModel, not even an explicit nil
### GetGpuOsName

`func (o *Deployment) GetGpuOsName() string`

GetGpuOsName returns the GpuOsName field if non-nil, zero value otherwise.

### GetGpuOsNameOk

`func (o *Deployment) GetGpuOsNameOk() (*string, bool)`

GetGpuOsNameOk returns a tuple with the GpuOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsName

`func (o *Deployment) SetGpuOsName(v string)`

SetGpuOsName sets GpuOsName field to given value.

### HasGpuOsName

`func (o *Deployment) HasGpuOsName() bool`

HasGpuOsName returns a boolean if a field has been set.

### SetGpuOsNameNil

`func (o *Deployment) SetGpuOsNameNil(b bool)`

 SetGpuOsNameNil sets the value for GpuOsName to be an explicit nil

### UnsetGpuOsName
`func (o *Deployment) UnsetGpuOsName()`

UnsetGpuOsName ensures that no value is present for GpuOsName, not even an explicit nil
### GetGpuOsRelease

`func (o *Deployment) GetGpuOsRelease() string`

GetGpuOsRelease returns the GpuOsRelease field if non-nil, zero value otherwise.

### GetGpuOsReleaseOk

`func (o *Deployment) GetGpuOsReleaseOk() (*string, bool)`

GetGpuOsReleaseOk returns a tuple with the GpuOsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsRelease

`func (o *Deployment) SetGpuOsRelease(v string)`

SetGpuOsRelease sets GpuOsRelease field to given value.

### HasGpuOsRelease

`func (o *Deployment) HasGpuOsRelease() bool`

HasGpuOsRelease returns a boolean if a field has been set.

### SetGpuOsReleaseNil

`func (o *Deployment) SetGpuOsReleaseNil(b bool)`

 SetGpuOsReleaseNil sets the value for GpuOsRelease to be an explicit nil

### UnsetGpuOsRelease
`func (o *Deployment) UnsetGpuOsRelease()`

UnsetGpuOsRelease ensures that no value is present for GpuOsRelease, not even an explicit nil
### GetGpuOsVersion

`func (o *Deployment) GetGpuOsVersion() string`

GetGpuOsVersion returns the GpuOsVersion field if non-nil, zero value otherwise.

### GetGpuOsVersionOk

`func (o *Deployment) GetGpuOsVersionOk() (*string, bool)`

GetGpuOsVersionOk returns a tuple with the GpuOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsVersion

`func (o *Deployment) SetGpuOsVersion(v string)`

SetGpuOsVersion sets GpuOsVersion field to given value.

### HasGpuOsVersion

`func (o *Deployment) HasGpuOsVersion() bool`

HasGpuOsVersion returns a boolean if a field has been set.

### SetGpuOsVersionNil

`func (o *Deployment) SetGpuOsVersionNil(b bool)`

 SetGpuOsVersionNil sets the value for GpuOsVersion to be an explicit nil

### UnsetGpuOsVersion
`func (o *Deployment) UnsetGpuOsVersion()`

UnsetGpuOsVersion ensures that no value is present for GpuOsVersion, not even an explicit nil
### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.


### GetIpAllowlist

`func (o *Deployment) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *Deployment) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *Deployment) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *Deployment) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetLifetime

`func (o *Deployment) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *Deployment) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *Deployment) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *Deployment) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### SetLifetimeNil

`func (o *Deployment) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *Deployment) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetMinGpuCount

`func (o *Deployment) GetMinGpuCount() int32`

GetMinGpuCount returns the MinGpuCount field if non-nil, zero value otherwise.

### GetMinGpuCountOk

`func (o *Deployment) GetMinGpuCountOk() (*int32, bool)`

GetMinGpuCountOk returns a tuple with the MinGpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGpuCount

`func (o *Deployment) SetMinGpuCount(v int32)`

SetMinGpuCount sets MinGpuCount field to given value.

### HasMinGpuCount

`func (o *Deployment) HasMinGpuCount() bool`

HasMinGpuCount returns a boolean if a field has been set.

### SetMinGpuCountNil

`func (o *Deployment) SetMinGpuCountNil(b bool)`

 SetMinGpuCountNil sets the value for MinGpuCount to be an explicit nil

### UnsetMinGpuCount
`func (o *Deployment) UnsetMinGpuCount()`

UnsetMinGpuCount ensures that no value is present for MinGpuCount, not even an explicit nil
### GetModified

`func (o *Deployment) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Deployment) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Deployment) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *Deployment) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Deployment) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Deployment) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Deployment) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### SetNodeCountNil

`func (o *Deployment) SetNodeCountNil(b bool)`

 SetNodeCountNil sets the value for NodeCount to be an explicit nil

### UnsetNodeCount
`func (o *Deployment) UnsetNodeCount()`

UnsetNodeCount ensures that no value is present for NodeCount, not even an explicit nil
### GetOemName

`func (o *Deployment) GetOemName() string`

GetOemName returns the OemName field if non-nil, zero value otherwise.

### GetOemNameOk

`func (o *Deployment) GetOemNameOk() (*string, bool)`

GetOemNameOk returns a tuple with the OemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemName

`func (o *Deployment) SetOemName(v string)`

SetOemName sets OemName field to given value.

### HasOemName

`func (o *Deployment) HasOemName() bool`

HasOemName returns a boolean if a field has been set.

### SetOemNameNil

`func (o *Deployment) SetOemNameNil(b bool)`

 SetOemNameNil sets the value for OemName to be an explicit nil

### UnsetOemName
`func (o *Deployment) UnsetOemName()`

UnsetOemName ensures that no value is present for OemName, not even an explicit nil
### GetOrgName

`func (o *Deployment) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Deployment) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Deployment) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOverrides

`func (o *Deployment) GetOverrides() Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *Deployment) GetOverridesOk() (*Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *Deployment) SetOverrides(v Overrides)`

SetOverrides sets Overrides field to given value.


### GetPersistOnFailure

`func (o *Deployment) GetPersistOnFailure() bool`

GetPersistOnFailure returns the PersistOnFailure field if non-nil, zero value otherwise.

### GetPersistOnFailureOk

`func (o *Deployment) GetPersistOnFailureOk() (*bool, bool)`

GetPersistOnFailureOk returns a tuple with the PersistOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOnFailure

`func (o *Deployment) SetPersistOnFailure(v bool)`

SetPersistOnFailure sets PersistOnFailure field to given value.

### HasPersistOnFailure

`func (o *Deployment) HasPersistOnFailure() bool`

HasPersistOnFailure returns a boolean if a field has been set.

### SetPersistOnFailureNil

`func (o *Deployment) SetPersistOnFailureNil(b bool)`

 SetPersistOnFailureNil sets the value for PersistOnFailure to be an explicit nil

### UnsetPersistOnFailure
`func (o *Deployment) UnsetPersistOnFailure()`

UnsetPersistOnFailure ensures that no value is present for PersistOnFailure, not even an explicit nil
### GetPersona

`func (o *Deployment) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *Deployment) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *Deployment) SetPersona(v string)`

SetPersona sets Persona field to given value.

### HasPersona

`func (o *Deployment) HasPersona() bool`

HasPersona returns a boolean if a field has been set.

### SetPersonaNil

`func (o *Deployment) SetPersonaNil(b bool)`

 SetPersonaNil sets the value for Persona to be an explicit nil

### UnsetPersona
`func (o *Deployment) UnsetPersona()`

UnsetPersona ensures that no value is present for Persona, not even an explicit nil
### GetPipeline

`func (o *Deployment) GetPipeline() int32`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *Deployment) GetPipelineOk() (*int32, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *Deployment) SetPipeline(v int32)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *Deployment) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *Deployment) SetPipelineNil(b bool)`

 SetPipelineNil sets the value for Pipeline to be an explicit nil

### UnsetPipeline
`func (o *Deployment) UnsetPipeline()`

UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
### GetPipelineBranch

`func (o *Deployment) GetPipelineBranch() string`

GetPipelineBranch returns the PipelineBranch field if non-nil, zero value otherwise.

### GetPipelineBranchOk

`func (o *Deployment) GetPipelineBranchOk() (*string, bool)`

GetPipelineBranchOk returns a tuple with the PipelineBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineBranch

`func (o *Deployment) SetPipelineBranch(v string)`

SetPipelineBranch sets PipelineBranch field to given value.

### HasPipelineBranch

`func (o *Deployment) HasPipelineBranch() bool`

HasPipelineBranch returns a boolean if a field has been set.

### SetPipelineBranchNil

`func (o *Deployment) SetPipelineBranchNil(b bool)`

 SetPipelineBranchNil sets the value for PipelineBranch to be an explicit nil

### UnsetPipelineBranch
`func (o *Deployment) UnsetPipelineBranch()`

UnsetPipelineBranch ensures that no value is present for PipelineBranch, not even an explicit nil
### GetPipelines

`func (o *Deployment) GetPipelines() []string`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *Deployment) GetPipelinesOk() (*[]string, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *Deployment) SetPipelines(v []string)`

SetPipelines sets Pipelines field to given value.


### GetPlatform

`func (o *Deployment) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Deployment) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Deployment) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Deployment) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Deployment) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Deployment) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPriority

`func (o *Deployment) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Deployment) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Deployment) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Deployment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *Deployment) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *Deployment) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *Deployment) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *Deployment) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *Deployment) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *Deployment) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetPublicKey

`func (o *Deployment) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Deployment) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Deployment) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Deployment) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *Deployment) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *Deployment) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetRegion

`func (o *Deployment) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Deployment) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Deployment) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Deployment) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Deployment) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Deployment) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRequestId

`func (o *Deployment) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Deployment) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Deployment) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Deployment) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequesterEmail

`func (o *Deployment) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *Deployment) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *Deployment) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.


### GetRequesterName

`func (o *Deployment) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *Deployment) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *Deployment) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.


### GetRetryCount

`func (o *Deployment) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Deployment) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Deployment) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.


### GetRuntime

`func (o *Deployment) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Deployment) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Deployment) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Deployment) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *Deployment) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *Deployment) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetRuntimeBranch

`func (o *Deployment) GetRuntimeBranch() string`

GetRuntimeBranch returns the RuntimeBranch field if non-nil, zero value otherwise.

### GetRuntimeBranchOk

`func (o *Deployment) GetRuntimeBranchOk() (*string, bool)`

GetRuntimeBranchOk returns a tuple with the RuntimeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeBranch

`func (o *Deployment) SetRuntimeBranch(v string)`

SetRuntimeBranch sets RuntimeBranch field to given value.

### HasRuntimeBranch

`func (o *Deployment) HasRuntimeBranch() bool`

HasRuntimeBranch returns a boolean if a field has been set.

### SetRuntimeBranchNil

`func (o *Deployment) SetRuntimeBranchNil(b bool)`

 SetRuntimeBranchNil sets the value for RuntimeBranch to be an explicit nil

### UnsetRuntimeBranch
`func (o *Deployment) UnsetRuntimeBranch()`

UnsetRuntimeBranch ensures that no value is present for RuntimeBranch, not even an explicit nil
### GetRuntimeCnsAddonPack

`func (o *Deployment) GetRuntimeCnsAddonPack() bool`

GetRuntimeCnsAddonPack returns the RuntimeCnsAddonPack field if non-nil, zero value otherwise.

### GetRuntimeCnsAddonPackOk

`func (o *Deployment) GetRuntimeCnsAddonPackOk() (*bool, bool)`

GetRuntimeCnsAddonPackOk returns a tuple with the RuntimeCnsAddonPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsAddonPack

`func (o *Deployment) SetRuntimeCnsAddonPack(v bool)`

SetRuntimeCnsAddonPack sets RuntimeCnsAddonPack field to given value.

### HasRuntimeCnsAddonPack

`func (o *Deployment) HasRuntimeCnsAddonPack() bool`

HasRuntimeCnsAddonPack returns a boolean if a field has been set.

### SetRuntimeCnsAddonPackNil

`func (o *Deployment) SetRuntimeCnsAddonPackNil(b bool)`

 SetRuntimeCnsAddonPackNil sets the value for RuntimeCnsAddonPack to be an explicit nil

### UnsetRuntimeCnsAddonPack
`func (o *Deployment) UnsetRuntimeCnsAddonPack()`

UnsetRuntimeCnsAddonPack ensures that no value is present for RuntimeCnsAddonPack, not even an explicit nil
### GetRuntimeCnsDocker

`func (o *Deployment) GetRuntimeCnsDocker() bool`

GetRuntimeCnsDocker returns the RuntimeCnsDocker field if non-nil, zero value otherwise.

### GetRuntimeCnsDockerOk

`func (o *Deployment) GetRuntimeCnsDockerOk() (*bool, bool)`

GetRuntimeCnsDockerOk returns a tuple with the RuntimeCnsDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDocker

`func (o *Deployment) SetRuntimeCnsDocker(v bool)`

SetRuntimeCnsDocker sets RuntimeCnsDocker field to given value.

### HasRuntimeCnsDocker

`func (o *Deployment) HasRuntimeCnsDocker() bool`

HasRuntimeCnsDocker returns a boolean if a field has been set.

### SetRuntimeCnsDockerNil

`func (o *Deployment) SetRuntimeCnsDockerNil(b bool)`

 SetRuntimeCnsDockerNil sets the value for RuntimeCnsDocker to be an explicit nil

### UnsetRuntimeCnsDocker
`func (o *Deployment) UnsetRuntimeCnsDocker()`

UnsetRuntimeCnsDocker ensures that no value is present for RuntimeCnsDocker, not even an explicit nil
### GetRuntimeCnsDriverVersion

`func (o *Deployment) GetRuntimeCnsDriverVersion() string`

GetRuntimeCnsDriverVersion returns the RuntimeCnsDriverVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsDriverVersionOk

`func (o *Deployment) GetRuntimeCnsDriverVersionOk() (*string, bool)`

GetRuntimeCnsDriverVersionOk returns a tuple with the RuntimeCnsDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsDriverVersion

`func (o *Deployment) SetRuntimeCnsDriverVersion(v string)`

SetRuntimeCnsDriverVersion sets RuntimeCnsDriverVersion field to given value.

### HasRuntimeCnsDriverVersion

`func (o *Deployment) HasRuntimeCnsDriverVersion() bool`

HasRuntimeCnsDriverVersion returns a boolean if a field has been set.

### SetRuntimeCnsDriverVersionNil

`func (o *Deployment) SetRuntimeCnsDriverVersionNil(b bool)`

 SetRuntimeCnsDriverVersionNil sets the value for RuntimeCnsDriverVersion to be an explicit nil

### UnsetRuntimeCnsDriverVersion
`func (o *Deployment) UnsetRuntimeCnsDriverVersion()`

UnsetRuntimeCnsDriverVersion ensures that no value is present for RuntimeCnsDriverVersion, not even an explicit nil
### GetRuntimeCnsK8s

`func (o *Deployment) GetRuntimeCnsK8s() bool`

GetRuntimeCnsK8s returns the RuntimeCnsK8s field if non-nil, zero value otherwise.

### GetRuntimeCnsK8sOk

`func (o *Deployment) GetRuntimeCnsK8sOk() (*bool, bool)`

GetRuntimeCnsK8sOk returns a tuple with the RuntimeCnsK8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsK8s

`func (o *Deployment) SetRuntimeCnsK8s(v bool)`

SetRuntimeCnsK8s sets RuntimeCnsK8s field to given value.

### HasRuntimeCnsK8s

`func (o *Deployment) HasRuntimeCnsK8s() bool`

HasRuntimeCnsK8s returns a boolean if a field has been set.

### SetRuntimeCnsK8sNil

`func (o *Deployment) SetRuntimeCnsK8sNil(b bool)`

 SetRuntimeCnsK8sNil sets the value for RuntimeCnsK8s to be an explicit nil

### UnsetRuntimeCnsK8s
`func (o *Deployment) UnsetRuntimeCnsK8s()`

UnsetRuntimeCnsK8s ensures that no value is present for RuntimeCnsK8s, not even an explicit nil
### GetRuntimeCnsNvidiaDriver

`func (o *Deployment) GetRuntimeCnsNvidiaDriver() bool`

GetRuntimeCnsNvidiaDriver returns the RuntimeCnsNvidiaDriver field if non-nil, zero value otherwise.

### GetRuntimeCnsNvidiaDriverOk

`func (o *Deployment) GetRuntimeCnsNvidiaDriverOk() (*bool, bool)`

GetRuntimeCnsNvidiaDriverOk returns a tuple with the RuntimeCnsNvidiaDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsNvidiaDriver

`func (o *Deployment) SetRuntimeCnsNvidiaDriver(v bool)`

SetRuntimeCnsNvidiaDriver sets RuntimeCnsNvidiaDriver field to given value.

### HasRuntimeCnsNvidiaDriver

`func (o *Deployment) HasRuntimeCnsNvidiaDriver() bool`

HasRuntimeCnsNvidiaDriver returns a boolean if a field has been set.

### SetRuntimeCnsNvidiaDriverNil

`func (o *Deployment) SetRuntimeCnsNvidiaDriverNil(b bool)`

 SetRuntimeCnsNvidiaDriverNil sets the value for RuntimeCnsNvidiaDriver to be an explicit nil

### UnsetRuntimeCnsNvidiaDriver
`func (o *Deployment) UnsetRuntimeCnsNvidiaDriver()`

UnsetRuntimeCnsNvidiaDriver ensures that no value is present for RuntimeCnsNvidiaDriver, not even an explicit nil
### GetRuntimeCnsVersion

`func (o *Deployment) GetRuntimeCnsVersion() string`

GetRuntimeCnsVersion returns the RuntimeCnsVersion field if non-nil, zero value otherwise.

### GetRuntimeCnsVersionOk

`func (o *Deployment) GetRuntimeCnsVersionOk() (*string, bool)`

GetRuntimeCnsVersionOk returns a tuple with the RuntimeCnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCnsVersion

`func (o *Deployment) SetRuntimeCnsVersion(v string)`

SetRuntimeCnsVersion sets RuntimeCnsVersion field to given value.

### HasRuntimeCnsVersion

`func (o *Deployment) HasRuntimeCnsVersion() bool`

HasRuntimeCnsVersion returns a boolean if a field has been set.

### SetRuntimeCnsVersionNil

`func (o *Deployment) SetRuntimeCnsVersionNil(b bool)`

 SetRuntimeCnsVersionNil sets the value for RuntimeCnsVersion to be an explicit nil

### UnsetRuntimeCnsVersion
`func (o *Deployment) UnsetRuntimeCnsVersion()`

UnsetRuntimeCnsVersion ensures that no value is present for RuntimeCnsVersion, not even an explicit nil
### GetRuntimeMig

`func (o *Deployment) GetRuntimeMig() bool`

GetRuntimeMig returns the RuntimeMig field if non-nil, zero value otherwise.

### GetRuntimeMigOk

`func (o *Deployment) GetRuntimeMigOk() (*bool, bool)`

GetRuntimeMigOk returns a tuple with the RuntimeMig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMig

`func (o *Deployment) SetRuntimeMig(v bool)`

SetRuntimeMig sets RuntimeMig field to given value.

### HasRuntimeMig

`func (o *Deployment) HasRuntimeMig() bool`

HasRuntimeMig returns a boolean if a field has been set.

### SetRuntimeMigNil

`func (o *Deployment) SetRuntimeMigNil(b bool)`

 SetRuntimeMigNil sets the value for RuntimeMig to be an explicit nil

### UnsetRuntimeMig
`func (o *Deployment) UnsetRuntimeMig()`

UnsetRuntimeMig ensures that no value is present for RuntimeMig, not even an explicit nil
### GetRuntimeMigProfile

`func (o *Deployment) GetRuntimeMigProfile() string`

GetRuntimeMigProfile returns the RuntimeMigProfile field if non-nil, zero value otherwise.

### GetRuntimeMigProfileOk

`func (o *Deployment) GetRuntimeMigProfileOk() (*string, bool)`

GetRuntimeMigProfileOk returns a tuple with the RuntimeMigProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMigProfile

`func (o *Deployment) SetRuntimeMigProfile(v string)`

SetRuntimeMigProfile sets RuntimeMigProfile field to given value.

### HasRuntimeMigProfile

`func (o *Deployment) HasRuntimeMigProfile() bool`

HasRuntimeMigProfile returns a boolean if a field has been set.

### SetRuntimeMigProfileNil

`func (o *Deployment) SetRuntimeMigProfileNil(b bool)`

 SetRuntimeMigProfileNil sets the value for RuntimeMigProfile to be an explicit nil

### UnsetRuntimeMigProfile
`func (o *Deployment) UnsetRuntimeMigProfile()`

UnsetRuntimeMigProfile ensures that no value is present for RuntimeMigProfile, not even an explicit nil
### GetRuntimeUrl

`func (o *Deployment) GetRuntimeUrl() string`

GetRuntimeUrl returns the RuntimeUrl field if non-nil, zero value otherwise.

### GetRuntimeUrlOk

`func (o *Deployment) GetRuntimeUrlOk() (*string, bool)`

GetRuntimeUrlOk returns a tuple with the RuntimeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUrl

`func (o *Deployment) SetRuntimeUrl(v string)`

SetRuntimeUrl sets RuntimeUrl field to given value.

### HasRuntimeUrl

`func (o *Deployment) HasRuntimeUrl() bool`

HasRuntimeUrl returns a boolean if a field has been set.

### SetRuntimeUrlNil

`func (o *Deployment) SetRuntimeUrlNil(b bool)`

 SetRuntimeUrlNil sets the value for RuntimeUrl to be an explicit nil

### UnsetRuntimeUrl
`func (o *Deployment) UnsetRuntimeUrl()`

UnsetRuntimeUrl ensures that no value is present for RuntimeUrl, not even an explicit nil
### GetSalesCreatedDate

`func (o *Deployment) GetSalesCreatedDate() time.Time`

GetSalesCreatedDate returns the SalesCreatedDate field if non-nil, zero value otherwise.

### GetSalesCreatedDateOk

`func (o *Deployment) GetSalesCreatedDateOk() (*time.Time, bool)`

GetSalesCreatedDateOk returns a tuple with the SalesCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCreatedDate

`func (o *Deployment) SetSalesCreatedDate(v time.Time)`

SetSalesCreatedDate sets SalesCreatedDate field to given value.

### HasSalesCreatedDate

`func (o *Deployment) HasSalesCreatedDate() bool`

HasSalesCreatedDate returns a boolean if a field has been set.

### SetSalesCreatedDateNil

`func (o *Deployment) SetSalesCreatedDateNil(b bool)`

 SetSalesCreatedDateNil sets the value for SalesCreatedDate to be an explicit nil

### UnsetSalesCreatedDate
`func (o *Deployment) UnsetSalesCreatedDate()`

UnsetSalesCreatedDate ensures that no value is present for SalesCreatedDate, not even an explicit nil
### GetSalesId

`func (o *Deployment) GetSalesId() string`

GetSalesId returns the SalesId field if non-nil, zero value otherwise.

### GetSalesIdOk

`func (o *Deployment) GetSalesIdOk() (*string, bool)`

GetSalesIdOk returns a tuple with the SalesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesId

`func (o *Deployment) SetSalesId(v string)`

SetSalesId sets SalesId field to given value.

### HasSalesId

`func (o *Deployment) HasSalesId() bool`

HasSalesId returns a boolean if a field has been set.

### SetSalesIdNil

`func (o *Deployment) SetSalesIdNil(b bool)`

 SetSalesIdNil sets the value for SalesId to be an explicit nil

### UnsetSalesId
`func (o *Deployment) UnsetSalesId()`

UnsetSalesId ensures that no value is present for SalesId, not even an explicit nil
### GetSalesOwnerEmail

`func (o *Deployment) GetSalesOwnerEmail() string`

GetSalesOwnerEmail returns the SalesOwnerEmail field if non-nil, zero value otherwise.

### GetSalesOwnerEmailOk

`func (o *Deployment) GetSalesOwnerEmailOk() (*string, bool)`

GetSalesOwnerEmailOk returns a tuple with the SalesOwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerEmail

`func (o *Deployment) SetSalesOwnerEmail(v string)`

SetSalesOwnerEmail sets SalesOwnerEmail field to given value.

### HasSalesOwnerEmail

`func (o *Deployment) HasSalesOwnerEmail() bool`

HasSalesOwnerEmail returns a boolean if a field has been set.

### SetSalesOwnerEmailNil

`func (o *Deployment) SetSalesOwnerEmailNil(b bool)`

 SetSalesOwnerEmailNil sets the value for SalesOwnerEmail to be an explicit nil

### UnsetSalesOwnerEmail
`func (o *Deployment) UnsetSalesOwnerEmail()`

UnsetSalesOwnerEmail ensures that no value is present for SalesOwnerEmail, not even an explicit nil
### GetSalesOwnerName

`func (o *Deployment) GetSalesOwnerName() string`

GetSalesOwnerName returns the SalesOwnerName field if non-nil, zero value otherwise.

### GetSalesOwnerNameOk

`func (o *Deployment) GetSalesOwnerNameOk() (*string, bool)`

GetSalesOwnerNameOk returns a tuple with the SalesOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerName

`func (o *Deployment) SetSalesOwnerName(v string)`

SetSalesOwnerName sets SalesOwnerName field to given value.

### HasSalesOwnerName

`func (o *Deployment) HasSalesOwnerName() bool`

HasSalesOwnerName returns a boolean if a field has been set.

### SetSalesOwnerNameNil

`func (o *Deployment) SetSalesOwnerNameNil(b bool)`

 SetSalesOwnerNameNil sets the value for SalesOwnerName to be an explicit nil

### UnsetSalesOwnerName
`func (o *Deployment) UnsetSalesOwnerName()`

UnsetSalesOwnerName ensures that no value is present for SalesOwnerName, not even an explicit nil
### GetServices

`func (o *Deployment) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Deployment) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Deployment) SetServices(v []string)`

SetServices sets Services field to given value.


### GetSshPort

`func (o *Deployment) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Deployment) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Deployment) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.


### GetSshUser

`func (o *Deployment) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *Deployment) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *Deployment) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetState

`func (o *Deployment) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetTags

`func (o *Deployment) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Deployment) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Deployment) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Deployment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Deployment) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Deployment) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetWorkshop

`func (o *Deployment) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *Deployment) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *Deployment) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *Deployment) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### SetWorkshopNil

`func (o *Deployment) SetWorkshopNil(b bool)`

 SetWorkshopNil sets the value for Workshop to be an explicit nil

### UnsetWorkshop
`func (o *Deployment) UnsetWorkshop()`

UnsetWorkshop ensures that no value is present for Workshop, not even an explicit nil
### GetWorkshopId

`func (o *Deployment) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *Deployment) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *Deployment) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *Deployment) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### SetWorkshopIdNil

`func (o *Deployment) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *Deployment) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil
### GetWorkshopOverridePassword

`func (o *Deployment) GetWorkshopOverridePassword() string`

GetWorkshopOverridePassword returns the WorkshopOverridePassword field if non-nil, zero value otherwise.

### GetWorkshopOverridePasswordOk

`func (o *Deployment) GetWorkshopOverridePasswordOk() (*string, bool)`

GetWorkshopOverridePasswordOk returns a tuple with the WorkshopOverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopOverridePassword

`func (o *Deployment) SetWorkshopOverridePassword(v string)`

SetWorkshopOverridePassword sets WorkshopOverridePassword field to given value.

### HasWorkshopOverridePassword

`func (o *Deployment) HasWorkshopOverridePassword() bool`

HasWorkshopOverridePassword returns a boolean if a field has been set.

### SetWorkshopOverridePasswordNil

`func (o *Deployment) SetWorkshopOverridePasswordNil(b bool)`

 SetWorkshopOverridePasswordNil sets the value for WorkshopOverridePassword to be an explicit nil

### UnsetWorkshopOverridePassword
`func (o *Deployment) UnsetWorkshopOverridePassword()`

UnsetWorkshopOverridePassword ensures that no value is present for WorkshopOverridePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


