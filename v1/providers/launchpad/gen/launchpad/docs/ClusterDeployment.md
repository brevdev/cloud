# ClusterDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionOperatingSystem** | Pointer to **string** | Override bastion operating system provisioned and/or configured by Liftoff | [optional] 
**CatalogId** | Pointer to **string** | Unique ID for this experience in the sales catalog. Must be unique. | [optional] 
**CatalogIdAlias** | Pointer to **string** | Human-readable identifier for the experience in the sales catalog (ex: LP-15). Must be unique. | [optional] 
**Cluster** | Pointer to [**DeploymentCluster**](DeploymentCluster.md) |  | [optional] 
**CollectionBranch** | Pointer to **string** | Override the Ansible collection branch initialized within the pipeline | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Experience** | Pointer to [**DeploymentExperience**](DeploymentExperience.md) |  | [optional] 
**ExperienceBranch** | Pointer to **string** | Override the experience branch | [optional] 
**ExperienceId** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**FlightcontrolRelease** | Pointer to **string** | Override the image tag used for Flight Control | [optional] 
**GarageId** | Pointer to **string** | Require a cluster with nodes in the given garage | [optional] 
**GcBranch** | Pointer to **string** | Override the default Ground Control branch | [optional] 
**GpuAlias** | Pointer to **string** | Require a cluster with the given GPU alias | [optional] 
**GpuCount** | Pointer to **int32** | Require a cluster with the given number of GPUs | [optional] 
**GpuModel** | Pointer to **string** | Require a cluster with the given GPU model | [optional] 
**GpuOsName** | Pointer to **string** |  | [optional] 
**GpuOsRelease** | Pointer to **string** |  | [optional] 
**GpuOsVersion** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | [readonly] 
**IpAllowlist** | Pointer to **[]string** | Host IP addresses that should be allowed to access the deployment | [optional] 
**Lifetime** | Pointer to **int32** | Set expires_at value to be a given number of days from the current time. A value of 0 will cause a deployment to remain active indefinitely. | [optional] 
**MinGpuCount** | Pointer to **int32** | Require a cluster whose GPU count is greater than or equal to the given number | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | Pointer to **int32** | Require a cluster with the given number of nodes | [optional] 
**OemName** | Pointer to **string** | Require a cluster manufactured by the given OEM name | [optional] 
**OrgName** | **string** | Requester&#39;s organization name | 
**Overrides** | **interface{}** |  | 
**PersistOnFailure** | Pointer to **bool** | Override the default cleanup/destroy behavior when a provisioning failure occurs | [optional] 
**Persona** | Pointer to **string** |  | [optional] 
**Pipeline** | Pointer to **int64** | Override the pipeline ID that will be triggered for request fulfillment | [optional] 
**PipelineBranch** | Pointer to **string** | Override the default pipeline branch ref used when triggering a Fuselage pipeline | [optional] 
**Pipelines** | **[]string** |  | 
**Platform** | Pointer to [**PlatformEnum**](PlatformEnum.md) |  | [optional] 
**Priority** | Pointer to [**PriorityEnum**](PriorityEnum.md) | Priority level for the request  * &#x60;p0&#x60; - p0 * &#x60;p1&#x60; - p1 * &#x60;p2&#x60; - p2 * &#x60;p3&#x60; - p3 | [optional] 
**ProviderName** | Pointer to **string** | Require a cluster from the given provider name | [optional] 
**PublicKey** | Pointer to **string** | The initial or administrative public key used during deployment creation. Additional keys can be authorized for access using the &#x60;ssh-keys&#x60; endpoint. | [optional] 
**Region** | Pointer to **string** | Require a cluster located in the given region | [optional] 
**RequestId** | Pointer to **string** | Trial request ID (ex: TRY-1234) | [optional] 
**RequesterEmail** | **string** | Email address of the user requesting the experience | 
**RequesterName** | **string** | Name of the user requesting the experience | 
**RetryCount** | **int32** | Number of times the deployment has been retried | [readonly] 
**SalesCreatedDate** | Pointer to **time.Time** |  | [optional] 
**SalesId** | Pointer to **string** | Unique identifier for the requester&#39;s sales relationship | [optional] 
**SalesOwnerEmail** | Pointer to **string** | Email address of the sales contact associated with the requester | [optional] 
**SalesOwnerName** | Pointer to **string** | Name of the sales contact associated with the requester | [optional] 
**Services** | **[]string** |  | 
**SshPort** | **int32** |  | [readonly] 
**SshUser** | **string** |  | [readonly] 
**State** | [**DeploymentState**](DeploymentState.md) | Current state of the deployment  * &#x60;destroyed&#x60; - Deployment has been fully destroyed * &#x60;destroying&#x60; - Deployment is being destroyed * &#x60;error&#x60; - Deployment has encountered a fatal error and will not be retried * &#x60;failed&#x60; - Deployment has failed but may be retried * &#x60;paused&#x60; - Deployment is paused but may be retried later * &#x60;ready&#x60; - Deployment is ready and all instances are running * &#x60;retrying&#x60; - Deployment is retrying * &#x60;starting&#x60; - Deployment instances are starting * &#x60;stopped&#x60; - Deployment instances are stopped * &#x60;stopping&#x60; - Deployment instances are stopping * &#x60;waiting&#x60; - Waiting for deployment to be ready | [readonly] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Workshop** | Pointer to **bool** | Require a cluster whose workshop flag is set | [optional] 
**WorkshopId** | Pointer to **string** | Require a cluster with the given workshop ID | [optional] 
**WorkshopOverridePassword** | Pointer to **string** | Override the deployment&#39;s default authentication to use a static password. This is useful for workshops when you&#39;d like an identical password associated with a collection of environments. (LaunchPad Team only) | [optional] 

## Methods

### NewClusterDeployment

`func NewClusterDeployment(created time.Time, id string, modified time.Time, orgName string, overrides interface{}, pipelines []string, requesterEmail string, requesterName string, retryCount int32, services []string, sshPort int32, sshUser string, state DeploymentState, ) *ClusterDeployment`

NewClusterDeployment instantiates a new ClusterDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDeploymentWithDefaults

`func NewClusterDeploymentWithDefaults() *ClusterDeployment`

NewClusterDeploymentWithDefaults instantiates a new ClusterDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionOperatingSystem

`func (o *ClusterDeployment) GetBastionOperatingSystem() string`

GetBastionOperatingSystem returns the BastionOperatingSystem field if non-nil, zero value otherwise.

### GetBastionOperatingSystemOk

`func (o *ClusterDeployment) GetBastionOperatingSystemOk() (*string, bool)`

GetBastionOperatingSystemOk returns a tuple with the BastionOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionOperatingSystem

`func (o *ClusterDeployment) SetBastionOperatingSystem(v string)`

SetBastionOperatingSystem sets BastionOperatingSystem field to given value.

### HasBastionOperatingSystem

`func (o *ClusterDeployment) HasBastionOperatingSystem() bool`

HasBastionOperatingSystem returns a boolean if a field has been set.

### GetCatalogId

`func (o *ClusterDeployment) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *ClusterDeployment) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *ClusterDeployment) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *ClusterDeployment) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogIdAlias

`func (o *ClusterDeployment) GetCatalogIdAlias() string`

GetCatalogIdAlias returns the CatalogIdAlias field if non-nil, zero value otherwise.

### GetCatalogIdAliasOk

`func (o *ClusterDeployment) GetCatalogIdAliasOk() (*string, bool)`

GetCatalogIdAliasOk returns a tuple with the CatalogIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdAlias

`func (o *ClusterDeployment) SetCatalogIdAlias(v string)`

SetCatalogIdAlias sets CatalogIdAlias field to given value.

### HasCatalogIdAlias

`func (o *ClusterDeployment) HasCatalogIdAlias() bool`

HasCatalogIdAlias returns a boolean if a field has been set.

### GetCluster

`func (o *ClusterDeployment) GetCluster() DeploymentCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterDeployment) GetClusterOk() (*DeploymentCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterDeployment) SetCluster(v DeploymentCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterDeployment) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCollectionBranch

`func (o *ClusterDeployment) GetCollectionBranch() string`

GetCollectionBranch returns the CollectionBranch field if non-nil, zero value otherwise.

### GetCollectionBranchOk

`func (o *ClusterDeployment) GetCollectionBranchOk() (*string, bool)`

GetCollectionBranchOk returns a tuple with the CollectionBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionBranch

`func (o *ClusterDeployment) SetCollectionBranch(v string)`

SetCollectionBranch sets CollectionBranch field to given value.

### HasCollectionBranch

`func (o *ClusterDeployment) HasCollectionBranch() bool`

HasCollectionBranch returns a boolean if a field has been set.

### GetCreated

`func (o *ClusterDeployment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterDeployment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterDeployment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExperience

`func (o *ClusterDeployment) GetExperience() DeploymentExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ClusterDeployment) GetExperienceOk() (*DeploymentExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ClusterDeployment) SetExperience(v DeploymentExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *ClusterDeployment) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetExperienceBranch

`func (o *ClusterDeployment) GetExperienceBranch() string`

GetExperienceBranch returns the ExperienceBranch field if non-nil, zero value otherwise.

### GetExperienceBranchOk

`func (o *ClusterDeployment) GetExperienceBranchOk() (*string, bool)`

GetExperienceBranchOk returns a tuple with the ExperienceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceBranch

`func (o *ClusterDeployment) SetExperienceBranch(v string)`

SetExperienceBranch sets ExperienceBranch field to given value.

### HasExperienceBranch

`func (o *ClusterDeployment) HasExperienceBranch() bool`

HasExperienceBranch returns a boolean if a field has been set.

### GetExperienceId

`func (o *ClusterDeployment) GetExperienceId() string`

GetExperienceId returns the ExperienceId field if non-nil, zero value otherwise.

### GetExperienceIdOk

`func (o *ClusterDeployment) GetExperienceIdOk() (*string, bool)`

GetExperienceIdOk returns a tuple with the ExperienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceId

`func (o *ClusterDeployment) SetExperienceId(v string)`

SetExperienceId sets ExperienceId field to given value.

### HasExperienceId

`func (o *ClusterDeployment) HasExperienceId() bool`

HasExperienceId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ClusterDeployment) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ClusterDeployment) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ClusterDeployment) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ClusterDeployment) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFlightcontrolRelease

`func (o *ClusterDeployment) GetFlightcontrolRelease() string`

GetFlightcontrolRelease returns the FlightcontrolRelease field if non-nil, zero value otherwise.

### GetFlightcontrolReleaseOk

`func (o *ClusterDeployment) GetFlightcontrolReleaseOk() (*string, bool)`

GetFlightcontrolReleaseOk returns a tuple with the FlightcontrolRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightcontrolRelease

`func (o *ClusterDeployment) SetFlightcontrolRelease(v string)`

SetFlightcontrolRelease sets FlightcontrolRelease field to given value.

### HasFlightcontrolRelease

`func (o *ClusterDeployment) HasFlightcontrolRelease() bool`

HasFlightcontrolRelease returns a boolean if a field has been set.

### GetGarageId

`func (o *ClusterDeployment) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *ClusterDeployment) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *ClusterDeployment) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *ClusterDeployment) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### GetGcBranch

`func (o *ClusterDeployment) GetGcBranch() string`

GetGcBranch returns the GcBranch field if non-nil, zero value otherwise.

### GetGcBranchOk

`func (o *ClusterDeployment) GetGcBranchOk() (*string, bool)`

GetGcBranchOk returns a tuple with the GcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcBranch

`func (o *ClusterDeployment) SetGcBranch(v string)`

SetGcBranch sets GcBranch field to given value.

### HasGcBranch

`func (o *ClusterDeployment) HasGcBranch() bool`

HasGcBranch returns a boolean if a field has been set.

### GetGpuAlias

`func (o *ClusterDeployment) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *ClusterDeployment) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *ClusterDeployment) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *ClusterDeployment) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### GetGpuCount

`func (o *ClusterDeployment) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ClusterDeployment) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ClusterDeployment) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ClusterDeployment) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *ClusterDeployment) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ClusterDeployment) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ClusterDeployment) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.

### HasGpuModel

`func (o *ClusterDeployment) HasGpuModel() bool`

HasGpuModel returns a boolean if a field has been set.

### GetGpuOsName

`func (o *ClusterDeployment) GetGpuOsName() string`

GetGpuOsName returns the GpuOsName field if non-nil, zero value otherwise.

### GetGpuOsNameOk

`func (o *ClusterDeployment) GetGpuOsNameOk() (*string, bool)`

GetGpuOsNameOk returns a tuple with the GpuOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsName

`func (o *ClusterDeployment) SetGpuOsName(v string)`

SetGpuOsName sets GpuOsName field to given value.

### HasGpuOsName

`func (o *ClusterDeployment) HasGpuOsName() bool`

HasGpuOsName returns a boolean if a field has been set.

### GetGpuOsRelease

`func (o *ClusterDeployment) GetGpuOsRelease() string`

GetGpuOsRelease returns the GpuOsRelease field if non-nil, zero value otherwise.

### GetGpuOsReleaseOk

`func (o *ClusterDeployment) GetGpuOsReleaseOk() (*string, bool)`

GetGpuOsReleaseOk returns a tuple with the GpuOsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsRelease

`func (o *ClusterDeployment) SetGpuOsRelease(v string)`

SetGpuOsRelease sets GpuOsRelease field to given value.

### HasGpuOsRelease

`func (o *ClusterDeployment) HasGpuOsRelease() bool`

HasGpuOsRelease returns a boolean if a field has been set.

### GetGpuOsVersion

`func (o *ClusterDeployment) GetGpuOsVersion() string`

GetGpuOsVersion returns the GpuOsVersion field if non-nil, zero value otherwise.

### GetGpuOsVersionOk

`func (o *ClusterDeployment) GetGpuOsVersionOk() (*string, bool)`

GetGpuOsVersionOk returns a tuple with the GpuOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuOsVersion

`func (o *ClusterDeployment) SetGpuOsVersion(v string)`

SetGpuOsVersion sets GpuOsVersion field to given value.

### HasGpuOsVersion

`func (o *ClusterDeployment) HasGpuOsVersion() bool`

HasGpuOsVersion returns a boolean if a field has been set.

### GetId

`func (o *ClusterDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetIpAllowlist

`func (o *ClusterDeployment) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *ClusterDeployment) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *ClusterDeployment) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *ClusterDeployment) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetLifetime

`func (o *ClusterDeployment) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *ClusterDeployment) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *ClusterDeployment) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *ClusterDeployment) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetMinGpuCount

`func (o *ClusterDeployment) GetMinGpuCount() int32`

GetMinGpuCount returns the MinGpuCount field if non-nil, zero value otherwise.

### GetMinGpuCountOk

`func (o *ClusterDeployment) GetMinGpuCountOk() (*int32, bool)`

GetMinGpuCountOk returns a tuple with the MinGpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGpuCount

`func (o *ClusterDeployment) SetMinGpuCount(v int32)`

SetMinGpuCount sets MinGpuCount field to given value.

### HasMinGpuCount

`func (o *ClusterDeployment) HasMinGpuCount() bool`

HasMinGpuCount returns a boolean if a field has been set.

### GetModified

`func (o *ClusterDeployment) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterDeployment) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterDeployment) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *ClusterDeployment) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterDeployment) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterDeployment) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterDeployment) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetOemName

`func (o *ClusterDeployment) GetOemName() string`

GetOemName returns the OemName field if non-nil, zero value otherwise.

### GetOemNameOk

`func (o *ClusterDeployment) GetOemNameOk() (*string, bool)`

GetOemNameOk returns a tuple with the OemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemName

`func (o *ClusterDeployment) SetOemName(v string)`

SetOemName sets OemName field to given value.

### HasOemName

`func (o *ClusterDeployment) HasOemName() bool`

HasOemName returns a boolean if a field has been set.

### GetOrgName

`func (o *ClusterDeployment) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ClusterDeployment) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ClusterDeployment) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOverrides

`func (o *ClusterDeployment) GetOverrides() interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ClusterDeployment) GetOverridesOk() (*interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ClusterDeployment) SetOverrides(v interface{})`

SetOverrides sets Overrides field to given value.


### SetOverridesNil

`func (o *ClusterDeployment) SetOverridesNil(b bool)`

 SetOverridesNil sets the value for Overrides to be an explicit nil

### UnsetOverrides
`func (o *ClusterDeployment) UnsetOverrides()`

UnsetOverrides ensures that no value is present for Overrides, not even an explicit nil
### GetPersistOnFailure

`func (o *ClusterDeployment) GetPersistOnFailure() bool`

GetPersistOnFailure returns the PersistOnFailure field if non-nil, zero value otherwise.

### GetPersistOnFailureOk

`func (o *ClusterDeployment) GetPersistOnFailureOk() (*bool, bool)`

GetPersistOnFailureOk returns a tuple with the PersistOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOnFailure

`func (o *ClusterDeployment) SetPersistOnFailure(v bool)`

SetPersistOnFailure sets PersistOnFailure field to given value.

### HasPersistOnFailure

`func (o *ClusterDeployment) HasPersistOnFailure() bool`

HasPersistOnFailure returns a boolean if a field has been set.

### GetPersona

`func (o *ClusterDeployment) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *ClusterDeployment) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *ClusterDeployment) SetPersona(v string)`

SetPersona sets Persona field to given value.

### HasPersona

`func (o *ClusterDeployment) HasPersona() bool`

HasPersona returns a boolean if a field has been set.

### GetPipeline

`func (o *ClusterDeployment) GetPipeline() int64`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *ClusterDeployment) GetPipelineOk() (*int64, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *ClusterDeployment) SetPipeline(v int64)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *ClusterDeployment) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineBranch

`func (o *ClusterDeployment) GetPipelineBranch() string`

GetPipelineBranch returns the PipelineBranch field if non-nil, zero value otherwise.

### GetPipelineBranchOk

`func (o *ClusterDeployment) GetPipelineBranchOk() (*string, bool)`

GetPipelineBranchOk returns a tuple with the PipelineBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineBranch

`func (o *ClusterDeployment) SetPipelineBranch(v string)`

SetPipelineBranch sets PipelineBranch field to given value.

### HasPipelineBranch

`func (o *ClusterDeployment) HasPipelineBranch() bool`

HasPipelineBranch returns a boolean if a field has been set.

### GetPipelines

`func (o *ClusterDeployment) GetPipelines() []string`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *ClusterDeployment) GetPipelinesOk() (*[]string, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *ClusterDeployment) SetPipelines(v []string)`

SetPipelines sets Pipelines field to given value.


### GetPlatform

`func (o *ClusterDeployment) GetPlatform() PlatformEnum`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ClusterDeployment) GetPlatformOk() (*PlatformEnum, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ClusterDeployment) SetPlatform(v PlatformEnum)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ClusterDeployment) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPriority

`func (o *ClusterDeployment) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ClusterDeployment) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ClusterDeployment) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ClusterDeployment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *ClusterDeployment) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ClusterDeployment) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ClusterDeployment) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ClusterDeployment) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetPublicKey

`func (o *ClusterDeployment) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ClusterDeployment) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ClusterDeployment) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ClusterDeployment) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterDeployment) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterDeployment) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterDeployment) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterDeployment) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestId

`func (o *ClusterDeployment) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClusterDeployment) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClusterDeployment) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ClusterDeployment) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequesterEmail

`func (o *ClusterDeployment) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *ClusterDeployment) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *ClusterDeployment) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.


### GetRequesterName

`func (o *ClusterDeployment) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *ClusterDeployment) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *ClusterDeployment) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.


### GetRetryCount

`func (o *ClusterDeployment) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ClusterDeployment) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ClusterDeployment) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.


### GetSalesCreatedDate

`func (o *ClusterDeployment) GetSalesCreatedDate() time.Time`

GetSalesCreatedDate returns the SalesCreatedDate field if non-nil, zero value otherwise.

### GetSalesCreatedDateOk

`func (o *ClusterDeployment) GetSalesCreatedDateOk() (*time.Time, bool)`

GetSalesCreatedDateOk returns a tuple with the SalesCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCreatedDate

`func (o *ClusterDeployment) SetSalesCreatedDate(v time.Time)`

SetSalesCreatedDate sets SalesCreatedDate field to given value.

### HasSalesCreatedDate

`func (o *ClusterDeployment) HasSalesCreatedDate() bool`

HasSalesCreatedDate returns a boolean if a field has been set.

### GetSalesId

`func (o *ClusterDeployment) GetSalesId() string`

GetSalesId returns the SalesId field if non-nil, zero value otherwise.

### GetSalesIdOk

`func (o *ClusterDeployment) GetSalesIdOk() (*string, bool)`

GetSalesIdOk returns a tuple with the SalesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesId

`func (o *ClusterDeployment) SetSalesId(v string)`

SetSalesId sets SalesId field to given value.

### HasSalesId

`func (o *ClusterDeployment) HasSalesId() bool`

HasSalesId returns a boolean if a field has been set.

### GetSalesOwnerEmail

`func (o *ClusterDeployment) GetSalesOwnerEmail() string`

GetSalesOwnerEmail returns the SalesOwnerEmail field if non-nil, zero value otherwise.

### GetSalesOwnerEmailOk

`func (o *ClusterDeployment) GetSalesOwnerEmailOk() (*string, bool)`

GetSalesOwnerEmailOk returns a tuple with the SalesOwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerEmail

`func (o *ClusterDeployment) SetSalesOwnerEmail(v string)`

SetSalesOwnerEmail sets SalesOwnerEmail field to given value.

### HasSalesOwnerEmail

`func (o *ClusterDeployment) HasSalesOwnerEmail() bool`

HasSalesOwnerEmail returns a boolean if a field has been set.

### GetSalesOwnerName

`func (o *ClusterDeployment) GetSalesOwnerName() string`

GetSalesOwnerName returns the SalesOwnerName field if non-nil, zero value otherwise.

### GetSalesOwnerNameOk

`func (o *ClusterDeployment) GetSalesOwnerNameOk() (*string, bool)`

GetSalesOwnerNameOk returns a tuple with the SalesOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOwnerName

`func (o *ClusterDeployment) SetSalesOwnerName(v string)`

SetSalesOwnerName sets SalesOwnerName field to given value.

### HasSalesOwnerName

`func (o *ClusterDeployment) HasSalesOwnerName() bool`

HasSalesOwnerName returns a boolean if a field has been set.

### GetServices

`func (o *ClusterDeployment) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ClusterDeployment) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ClusterDeployment) SetServices(v []string)`

SetServices sets Services field to given value.


### GetSshPort

`func (o *ClusterDeployment) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ClusterDeployment) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ClusterDeployment) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.


### GetSshUser

`func (o *ClusterDeployment) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *ClusterDeployment) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *ClusterDeployment) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetState

`func (o *ClusterDeployment) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterDeployment) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterDeployment) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetTags

`func (o *ClusterDeployment) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterDeployment) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterDeployment) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterDeployment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ClusterDeployment) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ClusterDeployment) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetWorkshop

`func (o *ClusterDeployment) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *ClusterDeployment) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *ClusterDeployment) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *ClusterDeployment) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### GetWorkshopId

`func (o *ClusterDeployment) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *ClusterDeployment) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *ClusterDeployment) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *ClusterDeployment) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### GetWorkshopOverridePassword

`func (o *ClusterDeployment) GetWorkshopOverridePassword() string`

GetWorkshopOverridePassword returns the WorkshopOverridePassword field if non-nil, zero value otherwise.

### GetWorkshopOverridePasswordOk

`func (o *ClusterDeployment) GetWorkshopOverridePasswordOk() (*string, bool)`

GetWorkshopOverridePasswordOk returns a tuple with the WorkshopOverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopOverridePassword

`func (o *ClusterDeployment) SetWorkshopOverridePassword(v string)`

SetWorkshopOverridePassword sets WorkshopOverridePassword field to given value.

### HasWorkshopOverridePassword

`func (o *ClusterDeployment) HasWorkshopOverridePassword() bool`

HasWorkshopOverridePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


