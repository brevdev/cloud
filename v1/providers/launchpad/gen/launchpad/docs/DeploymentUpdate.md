# DeploymentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionOperatingSystem** | **NullableString** | Override bastion operating system provisioned and/or configured by Liftoff | [readonly] 
**Cluster** | **NullableString** | The cluster where the experience has been deployed | [readonly] 
**CollectionBranch** | Pointer to **NullableString** | Override the Ansible collection branch initialized within the pipeline | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Experience** | [**DeploymentExperience**](DeploymentExperience.md) |  | 
**ExperienceBranch** | Pointer to **NullableString** | Override the experience branch | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FlightcontrolRelease** | Pointer to **NullableString** | Override the image tag used for Flight Control | [optional] 
**GarageId** | **NullableString** | Require a cluster with nodes in the given garage | [readonly] 
**GcBranch** | Pointer to **NullableString** | Override the default Ground Control branch | [optional] 
**GpuAlias** | **NullableString** | Require a cluster with the given GPU alias | [readonly] 
**GpuCount** | **NullableInt32** | Require a cluster with the given number of GPUs | [readonly] 
**GpuModel** | **NullableString** | Require a cluster with the given GPU model | [readonly] 
**GpuOsName** | **NullableString** | Override the GPU node operating system name | [readonly] 
**GpuOsRelease** | **NullableString** | Override the GPU node operating system release | [readonly] 
**GpuOsVersion** | **NullableString** | Override the GPU node operating system version | [readonly] 
**Id** | **string** |  | [readonly] 
**IpAllowlist** | Pointer to **[]string** | Host IP addresses that should be allowed to access the deployment | [optional] 
**Lifetime** | **NullableInt32** | Set expires_at value to be a given number of days from the current time. A value of 0 will cause a deployment to remain active indefinitely. | [readonly] 
**MinGpuCount** | **NullableInt32** | Require a cluster whose GPU count is greater than or equal to the given number | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | **NullableInt32** | Require a cluster with the given number of nodes | [readonly] 
**OemName** | **NullableString** | Require a cluster manufactured by the given OEM name | [readonly] 
**OrgName** | Pointer to **string** | Requester&#39;s organization name | [optional] 
**Overrides** | **interface{}** |  | 
**PersistOnFailure** | Pointer to **NullableBool** | Override the default cleanup/destroy behavior when a provisioning failure occurs | [optional] 
**Persona** | Pointer to **NullableString** | Override the defined persona in the experience | [optional] 
**Pipeline** | Pointer to **NullableInt64** | Override the pipeline ID that will be triggered for request fulfillment | [optional] 
**PipelineBranch** | Pointer to **NullableString** | Override the default pipeline branch ref used when triggering a Fuselage pipeline | [optional] 
**Pipelines** | **[]string** |  | 
**Platform** | Pointer to [**NullablePlatformEnum**](PlatformEnum.md) |  | [optional] 
**Priority** | [**PriorityEnum**](PriorityEnum.md) | Priority level for the request  * &#x60;p0&#x60; - p0 * &#x60;p1&#x60; - p1 * &#x60;p2&#x60; - p2 * &#x60;p3&#x60; - p3 | [readonly] 
**ProviderName** | **NullableString** | Require a cluster from the given provider name | [readonly] 
**PublicKey** | **NullableString** | The initial or administrative public key used during deployment creation. Additional keys can be authorized for access using the &#x60;ssh-keys&#x60; endpoint. | [readonly] 
**Region** | **NullableString** | Require a cluster located in the given region | [readonly] 
**RequestId** | **string** | Trial request ID (ex: TRY-1234) | [readonly] 
**RequesterEmail** | Pointer to **string** | Email address of the user requesting the experience | [optional] 
**RequesterName** | Pointer to **string** | Name of the user requesting the experience | [optional] 
**RetryCount** | Pointer to **int32** | Number of times the deployment has been retried | [optional] 
**SalesCreatedDate** | **NullableTime** | Timestamp when the requester&#39;s sales relationship was created | [readonly] 
**SalesId** | **NullableString** | Unique identifier for the requester&#39;s sales relationship | [readonly] 
**SalesOwnerEmail** | Pointer to **NullableString** | Email address of the sales contact associated with the requester | [optional] 
**SalesOwnerName** | Pointer to **NullableString** | Name of the sales contact associated with the requester | [optional] 
**Services** | **[]string** |  | 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) | Current state of the deployment  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Workshop** | **NullableBool** | Require a cluster whose workshop flag is set | [readonly] 
**WorkshopId** | **NullableString** | Require a cluster with the given workshop ID | [readonly] 
**WorkshopOverridePassword** | **NullableString** | Override the deployment&#39;s default authentication to use a static password. This is useful for workshops when you&#39;d like an identical password associated with a collection of environments. (LaunchPad Team only) | [readonly] 

## Methods

### NewDeploymentUpdate

`func NewDeploymentUpdate(bastionOperatingSystem NullableString, cluster NullableString, created time.Time, experience DeploymentExperience, garageId NullableString, gpuAlias NullableString, gpuCount NullableInt32, gpuModel NullableString, gpuOsName NullableString, gpuOsRelease NullableString, gpuOsVersion NullableString, id string, lifetime NullableInt32, minGpuCount NullableInt32, modified time.Time, nodeCount NullableInt32, oemName NullableString, overrides interface{}, pipelines []string, priority PriorityEnum, providerName NullableString, publicKey NullableString, region NullableString, requestId string, salesCreatedDate NullableTime, salesId NullableString, services []string, workshop NullableBool, workshopId NullableString, workshopOverridePassword NullableString, ) *DeploymentUpdate`

NewDeploymentUpdate instantiates a new DeploymentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentUpdateWithDefaults

`func NewDeploymentUpdateWithDefaults() *DeploymentUpdate`

NewDeploymentUpdateWithDefaults instantiates a new DeploymentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionOperatingSystem

`func (o *DeploymentUpdate) GetBastionOperatingSystem() string`

GetBastionOperatingSystem returns the BastionOperatingSystem field if non-nil, zero value otherwise.

### GetBastionOperatingSystemOk

`func (o *DeploymentUpdate) GetBastionOperatingSystemOk() (*string, bool)`

GetBastionOperatingSystemOk returns a tuple with the BastionOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionOperatingSystem

`func (o *DeploymentUpdate) SetBastionOperatingSystem(v string)`

SetBastionOperatingSystem sets BastionOperatingSystem field to given value.


### SetBastionOperatingSystemNil

`func (o *DeploymentUpdate) SetBastionOperatingSystemNil(b bool)`

 SetBastionOperatingSystemNil sets the value for BastionOperatingSystem to be an explicit nil

### UnsetBastionOperatingSystem
`func (o *DeploymentUpdate) UnsetBastionOperatingSystem()`

UnsetBastionOperatingSystem ensures that no value is present for BastionOperatingSystem, not even an explicit nil
### GetCluster

`func (o *DeploymentUpdate) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeploymentUpdate) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeploymentUpdate) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### SetClusterNil

`func (o *DeploymentUpdate) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeploymentUpdate) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetCollectionBranch

`func (o *DeploymentUpdate) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *DeploymentUpdate) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *DeploymentUpdate) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *DeploymentUpdate) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### SetCollectionBranchNil

`func (o *DeploymentUpdate) SetCollectionBranchNil(b bool)`

 SetCollectionBranchNil sets the value for CollectionBranch to be an explicit nil

### UnsetCollectionBranch
`func (o *DeploymentUpdate) UnsetCollectionBranch()`

UnsetCollectionBranch ensures that no value is present for CollectionBranch, not even an explicit nil
### GetCreated

`func (o *DeploymentUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExperience

`func (o *DeploymentUpdate) GetExperience() DeploymentExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *DeploymentUpdate) GetExperienceOk() (*DeploymentExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *DeploymentUpdate) SetExperience(v DeploymentExperience)`

SetExperience sets Experience field to given value.


### GetExperienceBranch

`func (o *DeploymentUpdate) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *DeploymentUpdate) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *DeploymentUpdate) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *DeploymentUpdate) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### SetExperienceBranchNil

`func (o *DeploymentUpdate) SetExperienceBranchNil(b bool)`

 SetExperienceBranchNil sets the value for ExperienceBranch to be an explicit nil

### UnsetExperienceBranch
`func (o *DeploymentUpdate) UnsetExperienceBranch()`

UnsetExperienceBranch ensures that no value is present for ExperienceBranch, not even an explicit nil
### GetExpiresAt

`func (o *DeploymentUpdate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DeploymentUpdate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DeploymentUpdate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DeploymentUpdate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *DeploymentUpdate) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DeploymentUpdate) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFlightcontrolRelease

`func (o *DeploymentUpdate) GetFlightcontrolRelease() string`

GetFlightcontrolRelease returns the FlightcontrolRelease field if non-nil, zero value otherwise.

### GetFlightcontrolReleaseOk

`func (o *DeploymentUpdate) GetFlightcontrolReleaseOk() (*string, bool)`

GetFlightcontrolReleaseOk returns a tuple with the FlightcontrolRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightcontrolRelease

`func (o *DeploymentUpdate) SetFlightcontrolRelease(v string)`

SetFlightcontrolRelease sets FlightcontrolRelease field to given value.

### HasFlightcontrolRelease

`func (o *DeploymentUpdate) HasFlightcontrolRelease() bool`

HasFlightcontrolRelease returns a boolean if a field has been set.

### SetFlightcontrolReleaseNil

`func (o *DeploymentUpdate) SetFlightcontrolReleaseNil(b bool)`

 SetFlightcontrolReleaseNil sets the value for FlightcontrolRelease to be an explicit nil

### UnsetFlightcontrolRelease
`func (o *DeploymentUpdate) UnsetFlightcontrolRelease()`

UnsetFlightcontrolRelease ensures that no value is present for FlightcontrolRelease, not even an explicit nil
### GetGarageId

`func (o *DeploymentUpdate) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *DeploymentUpdate) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *DeploymentUpdate) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.


### SetGarageIdNil

`func (o *DeploymentUpdate) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *DeploymentUpdate) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGcBranch

`func (o *DeploymentUpdate) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *DeploymentUpdate) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *DeploymentUpdate) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *DeploymentUpdate) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### SetGcBranchNil

`func (o *DeploymentUpdate) SetGcBranchNil(b bool)`

 SetGcBranchNil sets the value for GcBranch to be an explicit nil

### UnsetGcBranch
`func (o *DeploymentUpdate) UnsetGcBranch()`

UnsetGcBranch ensures that no value is present for GcBranch, not even an explicit nil
### GetGpuAlias

`func (o *DeploymentUpdate) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *DeploymentUpdate) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *DeploymentUpdate) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.


### SetGpuAliasNil

`func (o *DeploymentUpdate) SetGpuAliasNil(b bool)`

 SetGpuAliasNil sets the value for GpuAlias to be an explicit nil

### UnsetGpuAlias
`func (o *DeploymentUpdate) UnsetGpuAlias()`

UnsetGpuAlias ensures that no value is present for GpuAlias, not even an explicit nil
### GetGpuCount

`func (o *DeploymentUpdate) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *DeploymentUpdate) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *DeploymentUpdate) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### SetGpuCountNil

`func (o *DeploymentUpdate) SetGpuCountNil(b bool)`

 SetGpuCountNil sets the value for GpuCount to be an explicit nil

### UnsetGpuCount
`func (o *DeploymentUpdate) UnsetGpuCount()`

UnsetGpuCount ensures that no value is present for GpuCount, not even an explicit nil
### GetGpuModel

`func (o *DeploymentUpdate) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *DeploymentUpdate) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *DeploymentUpdate) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.


### SetGpuModelNil

`func (o *DeploymentUpdate) SetGpuModelNil(b bool)`

 SetGpuModelNil sets the value for GpuModel to be an explicit nil

### UnsetGpuModel
`func (o *DeploymentUpdate) UnsetGpuModel()`

UnsetGpuModel ensures that no value is present for GpuModel, not even an explicit nil
### GetGpuOsName

`func (o *DeploymentUpdate) GetGpuOsName() string`

GetGpuOsName returns the GpuOsName field if non-nil, zero value otherwise.

### GetGpuOsNameOk

`func (o *DeploymentUpdate) GetGpuOsNameOk() (*string, bool)`

GetGpuOsNameOk returns a tuple with the GpuOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsName

`func (o *DeploymentUpdate) SetGpuOsName(v string)`

SetGpuOsName sets GpuOsName field to given value.


### SetGpuOsNameNil

`func (o *DeploymentUpdate) SetGpuOsNameNil(b bool)`

 SetGpuOsNameNil sets the value for GpuOsName to be an explicit nil

### UnsetGpuOsName
`func (o *DeploymentUpdate) UnsetGpuOsName()`

UnsetGpuOsName ensures that no value is present for GpuOsName, not even an explicit nil
### GetGpuOsRelease

`func (o *DeploymentUpdate) GetGpuOsRelease() string`

GetGpuOsRelease returns the GpuOsRelease field if non-nil, zero value otherwise.

### GetGpuOsReleaseOk

`func (o *DeploymentUpdate) GetGpuOsReleaseOk() (*string, bool)`

GetGpuOsReleaseOk returns a tuple with the GpuOsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsRelease

`func (o *DeploymentUpdate) SetGpuOsRelease(v string)`

SetGpuOsRelease sets GpuOsRelease field to given value.


### SetGpuOsReleaseNil

`func (o *DeploymentUpdate) SetGpuOsReleaseNil(b bool)`

 SetGpuOsReleaseNil sets the value for GpuOsRelease to be an explicit nil

### UnsetGpuOsRelease
`func (o *DeploymentUpdate) UnsetGpuOsRelease()`

UnsetGpuOsRelease ensures that no value is present for GpuOsRelease, not even an explicit nil
### GetGpuOsVersion

`func (o *DeploymentUpdate) GetGpuOsVersion() string`

GetGpuOsVersion returns the GpuOsVersion field if non-nil, zero value otherwise.

### GetGpuOsVersionOk

`func (o *DeploymentUpdate) GetGpuOsVersionOk() (*string, bool)`

GetGpuOsVersionOk returns a tuple with the GpuOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsVersion

`func (o *DeploymentUpdate) SetGpuOsVersion(v string)`

SetGpuOsVersion sets GpuOsVersion field to given value.


### SetGpuOsVersionNil

`func (o *DeploymentUpdate) SetGpuOsVersionNil(b bool)`

 SetGpuOsVersionNil sets the value for GpuOsVersion to be an explicit nil

### UnsetGpuOsVersion
`func (o *DeploymentUpdate) UnsetGpuOsVersion()`

UnsetGpuOsVersion ensures that no value is present for GpuOsVersion, not even an explicit nil
### GetId

`func (o *DeploymentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetIpAllowlist

`func (o *DeploymentUpdate) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *DeploymentUpdate) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *DeploymentUpdate) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *DeploymentUpdate) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetLifetime

`func (o *DeploymentUpdate) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *DeploymentUpdate) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *DeploymentUpdate) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.


### SetLifetimeNil

`func (o *DeploymentUpdate) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *DeploymentUpdate) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetMinGpuCount

`func (o *DeploymentUpdate) GetMinGpuCount() int32`

GetMinGpuCount returns the MinGpuCount field if non-nil, zero value otherwise.

### GetMinGpuCountOk

`func (o *DeploymentUpdate) GetMinGpuCountOk() (*int32, bool)`

GetMinGpuCountOk returns a tuple with the MinGpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGpuCount

`func (o *DeploymentUpdate) SetMinGpuCount(v int32)`

SetMinGpuCount sets MinGpuCount field to given value.


### SetMinGpuCountNil

`func (o *DeploymentUpdate) SetMinGpuCountNil(b bool)`

 SetMinGpuCountNil sets the value for MinGpuCount to be an explicit nil

### UnsetMinGpuCount
`func (o *DeploymentUpdate) UnsetMinGpuCount()`

UnsetMinGpuCount ensures that no value is present for MinGpuCount, not even an explicit nil
### GetModified

`func (o *DeploymentUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *DeploymentUpdate) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *DeploymentUpdate) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *DeploymentUpdate) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### SetNodeCountNil

`func (o *DeploymentUpdate) SetNodeCountNil(b bool)`

 SetNodeCountNil sets the value for NodeCount to be an explicit nil

### UnsetNodeCount
`func (o *DeploymentUpdate) UnsetNodeCount()`

UnsetNodeCount ensures that no value is present for NodeCount, not even an explicit nil
### GetOemName

`func (o *DeploymentUpdate) GetOemName() string`

GetOemName returns the OemName field if non-nil, zero value otherwise.

### GetOemNameOk

`func (o *DeploymentUpdate) GetOemNameOk() (*string, bool)`

GetOemNameOk returns a tuple with the OemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemName

`func (o *DeploymentUpdate) SetOemName(v string)`

SetOemName sets OemName field to given value.


### SetOemNameNil

`func (o *DeploymentUpdate) SetOemNameNil(b bool)`

 SetOemNameNil sets the value for OemName to be an explicit nil

### UnsetOemName
`func (o *DeploymentUpdate) UnsetOemName()`

UnsetOemName ensures that no value is present for OemName, not even an explicit nil
### GetOrgName

`func (o *DeploymentUpdate) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DeploymentUpdate) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DeploymentUpdate) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DeploymentUpdate) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOverrides

`func (o *DeploymentUpdate) GetOverrides() interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *DeploymentUpdate) GetOverridesOk() (*interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *DeploymentUpdate) SetOverrides(v interface{})`

SetOverrides sets Overrides field to given value.


### SetOverridesNil

`func (o *DeploymentUpdate) SetOverridesNil(b bool)`

 SetOverridesNil sets the value for Overrides to be an explicit nil

### UnsetOverrides
`func (o *DeploymentUpdate) UnsetOverrides()`

UnsetOverrides ensures that no value is present for Overrides, not even an explicit nil
### GetPersistOnFailure

`func (o *DeploymentUpdate) GetPersistOnFailure() bool`

GetPersistOnFailure returns the PersistOnFailure field if non-nil, zero value otherwise.

### GetPersistOnFailureOk

`func (o *DeploymentUpdate) GetPersistOnFailureOk() (*bool, bool)`

GetPersistOnFailureOk returns a tuple with the PersistOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOnFailure

`func (o *DeploymentUpdate) SetPersistOnFailure(v bool)`

SetPersistOnFailure sets PersistOnFailure field to given value.

### HasPersistOnFailure

`func (o *DeploymentUpdate) HasPersistOnFailure() bool`

HasPersistOnFailure returns a boolean if a field has been set.

### SetPersistOnFailureNil

`func (o *DeploymentUpdate) SetPersistOnFailureNil(b bool)`

 SetPersistOnFailureNil sets the value for PersistOnFailure to be an explicit nil

### UnsetPersistOnFailure
`func (o *DeploymentUpdate) UnsetPersistOnFailure()`

UnsetPersistOnFailure ensures that no value is present for PersistOnFailure, not even an explicit nil
### GetPersona

`func (o *DeploymentUpdate) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *DeploymentUpdate) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *DeploymentUpdate) SetPersona(v string)`

SetPersona sets Persona field to given value.

### HasPersona

`func (o *DeploymentUpdate) HasPersona() bool`

HasPersona returns a boolean if a field has been set.

### SetPersonaNil

`func (o *DeploymentUpdate) SetPersonaNil(b bool)`

 SetPersonaNil sets the value for Persona to be an explicit nil

### UnsetPersona
`func (o *DeploymentUpdate) UnsetPersona()`

UnsetPersona ensures that no value is present for Persona, not even an explicit nil
### GetPipeline

`func (o *DeploymentUpdate) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DeploymentUpdate) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DeploymentUpdate) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DeploymentUpdate) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *DeploymentUpdate) SetPipelineNil(b bool)`

 SetPipelineNil sets the value for Pipeline to be an explicit nil

### UnsetPipeline
`func (o *DeploymentUpdate) UnsetPipeline()`

UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
### GetPipelineBranch

`func (o *DeploymentUpdate) GetPipelineBranch() string`

GetPipelineBranch returns the PipelineBranch field if non-nil, zero value otherwise.

### GetPipelineBranchOk

`func (o *DeploymentUpdate) GetPipelineBranchOk() (*string, bool)`

GetPipelineBranchOk returns a tuple with the PipelineBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineBranch

`func (o *DeploymentUpdate) SetPipelineBranch(v string)`

SetPipelineBranch sets PipelineBranch field to given value.

### HasPipelineBranch

`func (o *DeploymentUpdate) HasPipelineBranch() bool`

HasPipelineBranch returns a boolean if a field has been set.

### SetPipelineBranchNil

`func (o *DeploymentUpdate) SetPipelineBranchNil(b bool)`

 SetPipelineBranchNil sets the value for PipelineBranch to be an explicit nil

### UnsetPipelineBranch
`func (o *DeploymentUpdate) UnsetPipelineBranch()`

UnsetPipelineBranch ensures that no value is present for PipelineBranch, not even an explicit nil
### GetPipelines

`func (o *DeploymentUpdate) GetPipelines() []string`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *DeploymentUpdate) GetPipelinesOk() (*[]string, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *DeploymentUpdate) SetPipelines(v []string)`

SetPipelines sets Pipelines field to given value.


### GetPlatform

`func (o *DeploymentUpdate) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeploymentUpdate) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeploymentUpdate) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeploymentUpdate) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *DeploymentUpdate) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *DeploymentUpdate) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPriority

`func (o *DeploymentUpdate) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeploymentUpdate) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeploymentUpdate) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.


### GetProviderName

`func (o *DeploymentUpdate) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *DeploymentUpdate) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *DeploymentUpdate) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### SetProviderNameNil

`func (o *DeploymentUpdate) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *DeploymentUpdate) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetPublicKey

`func (o *DeploymentUpdate) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DeploymentUpdate) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DeploymentUpdate) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *DeploymentUpdate) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *DeploymentUpdate) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetRegion

`func (o *DeploymentUpdate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentUpdate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentUpdate) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *DeploymentUpdate) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *DeploymentUpdate) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRequestId

`func (o *DeploymentUpdate) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeploymentUpdate) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeploymentUpdate) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequesterEmail

`func (o *DeploymentUpdate) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *DeploymentUpdate) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *DeploymentUpdate) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.

### HasRequesterEmail

`func (o *DeploymentUpdate) HasRequesterEmail() bool`

HasRequesterEmail returns a boolean if a field has been set.

### GetRequesterName

`func (o *DeploymentUpdate) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *DeploymentUpdate) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *DeploymentUpdate) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *DeploymentUpdate) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetRetryCount

`func (o *DeploymentUpdate) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *DeploymentUpdate) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *DeploymentUpdate) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *DeploymentUpdate) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetSalesCreatedDate

`func (o *DeploymentUpdate) GetSalesCreatedDate() time.Time`

GetSalesCreatedDate returns the SalesCreatedDate field if non-nil, zero value otherwise.

### GetSalesCreatedDateOk

`func (o *DeploymentUpdate) GetSalesCreatedDateOk() (*time.Time, bool)`

GetSalesCreatedDateOk returns a tuple with the SalesCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCreatedDate

`func (o *DeploymentUpdate) SetSalesCreatedDate(v time.Time)`

SetSalesCreatedDate sets SalesCreatedDate field to given value.


### SetSalesCreatedDateNil

`func (o *DeploymentUpdate) SetSalesCreatedDateNil(b bool)`

 SetSalesCreatedDateNil sets the value for SalesCreatedDate to be an explicit nil

### UnsetSalesCreatedDate
`func (o *DeploymentUpdate) UnsetSalesCreatedDate()`

UnsetSalesCreatedDate ensures that no value is present for SalesCreatedDate, not even an explicit nil
### GetSalesId

`func (o *DeploymentUpdate) GetSalesId() string`

GetSalesId returns the SalesId field if non-nil, zero value otherwise.

### GetSalesIdOk

`func (o *DeploymentUpdate) GetSalesIdOk() (*string, bool)`

GetSalesIdOk returns a tuple with the SalesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesId

`func (o *DeploymentUpdate) SetSalesId(v string)`

SetSalesId sets SalesId field to given value.


### SetSalesIdNil

`func (o *DeploymentUpdate) SetSalesIdNil(b bool)`

 SetSalesIdNil sets the value for SalesId to be an explicit nil

### UnsetSalesId
`func (o *DeploymentUpdate) UnsetSalesId()`

UnsetSalesId ensures that no value is present for SalesId, not even an explicit nil
### GetSalesOwnerEmail

`func (o *DeploymentUpdate) GetSalesOwnerEmail() string`

GetSalesOwnerEmail returns the SalesOwnerEmail field if non-nil, zero value otherwise.

### GetSalesOwnerEmailOk

`func (o *DeploymentUpdate) GetSalesOwnerEmailOk() (*string, bool)`

GetSalesOwnerEmailOk returns a tuple with the SalesOwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerEmail

`func (o *DeploymentUpdate) SetSalesOwnerEmail(v string)`

SetSalesOwnerEmail sets SalesOwnerEmail field to given value.

### HasSalesOwnerEmail

`func (o *DeploymentUpdate) HasSalesOwnerEmail() bool`

HasSalesOwnerEmail returns a boolean if a field has been set.

### SetSalesOwnerEmailNil

`func (o *DeploymentUpdate) SetSalesOwnerEmailNil(b bool)`

 SetSalesOwnerEmailNil sets the value for SalesOwnerEmail to be an explicit nil

### UnsetSalesOwnerEmail
`func (o *DeploymentUpdate) UnsetSalesOwnerEmail()`

UnsetSalesOwnerEmail ensures that no value is present for SalesOwnerEmail, not even an explicit nil
### GetSalesOwnerName

`func (o *DeploymentUpdate) GetSalesOwnerName() string`

GetSalesOwnerName returns the SalesOwnerName field if non-nil, zero value otherwise.

### GetSalesOwnerNameOk

`func (o *DeploymentUpdate) GetSalesOwnerNameOk() (*string, bool)`

GetSalesOwnerNameOk returns a tuple with the SalesOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerName

`func (o *DeploymentUpdate) SetSalesOwnerName(v string)`

SetSalesOwnerName sets SalesOwnerName field to given value.

### HasSalesOwnerName

`func (o *DeploymentUpdate) HasSalesOwnerName() bool`

HasSalesOwnerName returns a boolean if a field has been set.

### SetSalesOwnerNameNil

`func (o *DeploymentUpdate) SetSalesOwnerNameNil(b bool)`

 SetSalesOwnerNameNil sets the value for SalesOwnerName to be an explicit nil

### UnsetSalesOwnerName
`func (o *DeploymentUpdate) UnsetSalesOwnerName()`

UnsetSalesOwnerName ensures that no value is present for SalesOwnerName, not even an explicit nil
### GetServices

`func (o *DeploymentUpdate) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeploymentUpdate) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeploymentUpdate) SetServices(v []string)`

SetServices sets Services field to given value.


### GetState

`func (o *DeploymentUpdate) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentUpdate) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentUpdate) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *DeploymentUpdate) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeploymentUpdate) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeploymentUpdate) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeploymentUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeploymentUpdate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeploymentUpdate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetWorkshop

`func (o *DeploymentUpdate) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *DeploymentUpdate) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *DeploymentUpdate) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.


### SetWorkshopNil

`func (o *DeploymentUpdate) SetWorkshopNil(b bool)`

 SetWorkshopNil sets the value for Workshop to be an explicit nil

### UnsetWorkshop
`func (o *DeploymentUpdate) UnsetWorkshop()`

UnsetWorkshop ensures that no value is present for Workshop, not even an explicit nil
### GetWorkshopId

`func (o *DeploymentUpdate) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *DeploymentUpdate) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *DeploymentUpdate) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.


### SetWorkshopIdNil

`func (o *DeploymentUpdate) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *DeploymentUpdate) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil
### GetWorkshopOverridePassword

`func (o *DeploymentUpdate) GetWorkshopOverridePassword() string`

GetWorkshopOverridePassword returns the WorkshopOverridePassword field if non-nil, zero value otherwise.

### GetWorkshopOverridePasswordOk

`func (o *DeploymentUpdate) GetWorkshopOverridePasswordOk() (*string, bool)`

GetWorkshopOverridePasswordOk returns a tuple with the WorkshopOverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopOverridePassword

`func (o *DeploymentUpdate) SetWorkshopOverridePassword(v string)`

SetWorkshopOverridePassword sets WorkshopOverridePassword field to given value.


### SetWorkshopOverridePasswordNil

`func (o *DeploymentUpdate) SetWorkshopOverridePasswordNil(b bool)`

 SetWorkshopOverridePasswordNil sets the value for WorkshopOverridePassword to be an explicit nil

### UnsetWorkshopOverridePassword
`func (o *DeploymentUpdate) UnsetWorkshopOverridePassword()`

UnsetWorkshopOverridePassword ensures that no value is present for WorkshopOverridePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


