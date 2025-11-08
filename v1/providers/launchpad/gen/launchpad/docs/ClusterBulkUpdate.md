# ClusterBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Is the cluster currently available for provisioning? | [readonly] 
**BastionName** | Pointer to **NullableString** | Name of the bastion assigned to the cluster | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Deployment** | [**ClusterDeployment**](ClusterDeployment.md) |  | 
**Enabled** | Pointer to **bool** | Is the cluster administratively enabled? | [optional] 
**Experience** | Pointer to **NullableString** | Experience provisioned onto this cluster | [optional] 
**FreeBy** | Pointer to **NullableTime** |  | [optional] 
**Gpus** | [**[]ClusterGpusInner**](ClusterGpusInner.md) |  | 
**GpuAlias** | **string** | Alias for GPU plan (i.e. installed GPU type and count) | [readonly] 
**GpuCount** | **int32** |  | [readonly] 
**Id** | **string** |  | [readonly] 
**Instances** | [**[]ClusterInstancesInner**](ClusterInstancesInner.md) |  | 
**LastUsed** | **NullableTime** | Timestamp of when the cluster was last in use | [readonly] 
**Maintenance** | Pointer to **bool** | Is the cluster in maintenance mode? | [optional] 
**MgmtIp** | Pointer to **string** | Management IP address | [optional] 
**MgmtMac** | Pointer to **NullableString** | Management MAC address | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Netmask** | Pointer to **NullableInt32** | The subnet mask of the cluster&#39;s public IP address in CIDR notation | [optional] 
**NodeCount** | **int32** |  | [readonly] 
**Nodes** | Pointer to [**[]ClusterNodesInner**](ClusterNodesInner.md) |  | [optional] 
**Notes** | Pointer to **NullableString** | Administrative comments about the cluster | [optional] 
**Persist** | Pointer to **bool** | Is the cluster exempt from provisioning_state timeouts? Can be used to ensure the cluster persists after a provisioning failure. | [optional] 
**ProviderCapacity** | **bool** | Does the provider have capacity to provision this cluster? | [readonly] 
**ProvisionUser** | Pointer to **NullableString** | Username used for provisioning this cluster | [optional] 
**ProvisioningAttempts** | Pointer to **int32** | The number of attempts that have been made to provision this cluster. Automatically resets to 0 after successful provisioning. | [optional] 
**ProvisioningConfig** | Pointer to **NullableString** | Applied provisioning configuration for the cluster | [optional] 
**ProvisioningRequest** | Pointer to **NullableString** | Requested provisioning configuration for the cluster | [optional] 
**ProvisioningState** | Pointer to [**ProvisioningStateEnum**](ProvisioningStateEnum.md) | Is the cluster currently provisioned?  * &#x60;deployed&#x60; - Cluster is in use by a deployment * &#x60;deploying&#x60; - Provisioning is in progress * &#x60;destroying&#x60; - Cluster is being destroyed * &#x60;pending&#x60; - Provisioning will begin soon * &#x60;ready&#x60; - Provisioning has completed and is ready for a deployment * &#x60;reserved&#x60; - Cluster is unprovisioned but reserved for later use * &#x60;unprovisioned&#x60; - Cluster has not yet been provisioned | [optional] 
**PublicAddress** | **NullableString** | Public IP address or fully-qualified domain name of this cluster | [readonly] 
**RequestId** | **NullableString** | The request ID for the lab that is currently provisioned on this cluster (ex: TRY-1234) | [readonly] 
**Reservation** | Pointer to **bool** | Is the cluster a static reservation from its provider? | [optional] 
**TenantIds** | **[]string** | Tenant UUID(s) that have been generated for this cluster during provisioning | 
**VlanId** | Pointer to **int32** | VLAN number | [optional] 
**Workshop** | Pointer to **bool** | Is the cluster set aside for use in a workshop? | [optional] 
**WorkshopId** | Pointer to **NullableString** | Identifier of the workshop this cluster is set aside for | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewClusterBulkUpdate

`func NewClusterBulkUpdate(available bool, created time.Time, deployment ClusterDeployment, gpus []ClusterGpusInner, gpuAlias string, gpuCount int32, id string, instances []ClusterInstancesInner, lastUsed NullableTime, modified time.Time, nodeCount int32, providerCapacity bool, publicAddress NullableString, requestId NullableString, tenantIds []string, count int32, ids []string, result string, ) *ClusterBulkUpdate`

NewClusterBulkUpdate instantiates a new ClusterBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterBulkUpdateWithDefaults

`func NewClusterBulkUpdateWithDefaults() *ClusterBulkUpdate`

NewClusterBulkUpdateWithDefaults instantiates a new ClusterBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ClusterBulkUpdate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ClusterBulkUpdate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ClusterBulkUpdate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetBastionName

`func (o *ClusterBulkUpdate) GetBastionName() string`

GetBastionName returns the BastionName field if non-nil, zero value otherwise.

### GetBastionNameOk

`func (o *ClusterBulkUpdate) GetBastionNameOk() (*string, bool)`

GetBastionNameOk returns a tuple with the BastionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionName

`func (o *ClusterBulkUpdate) SetBastionName(v string)`

SetBastionName sets BastionName field to given value.

### HasBastionName

`func (o *ClusterBulkUpdate) HasBastionName() bool`

HasBastionName returns a boolean if a field has been set.

### SetBastionNameNil

`func (o *ClusterBulkUpdate) SetBastionNameNil(b bool)`

 SetBastionNameNil sets the value for BastionName to be an explicit nil

### UnsetBastionName
`func (o *ClusterBulkUpdate) UnsetBastionName()`

UnsetBastionName ensures that no value is present for BastionName, not even an explicit nil
### GetCreated

`func (o *ClusterBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *ClusterBulkUpdate) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ClusterBulkUpdate) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ClusterBulkUpdate) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetEnabled

`func (o *ClusterBulkUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterBulkUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterBulkUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterBulkUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExperience

`func (o *ClusterBulkUpdate) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ClusterBulkUpdate) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ClusterBulkUpdate) SetExperience(v string)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *ClusterBulkUpdate) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *ClusterBulkUpdate) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *ClusterBulkUpdate) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetFreeBy

`func (o *ClusterBulkUpdate) GetFreeBy() time.Time`

GetFreeBy returns the FreeBy field if non-nil, zero value otherwise.

### GetFreeByOk

`func (o *ClusterBulkUpdate) GetFreeByOk() (*time.Time, bool)`

GetFreeByOk returns a tuple with the FreeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBy

`func (o *ClusterBulkUpdate) SetFreeBy(v time.Time)`

SetFreeBy sets FreeBy field to given value.

### HasFreeBy

`func (o *ClusterBulkUpdate) HasFreeBy() bool`

HasFreeBy returns a boolean if a field has been set.

### SetFreeByNil

`func (o *ClusterBulkUpdate) SetFreeByNil(b bool)`

 SetFreeByNil sets the value for FreeBy to be an explicit nil

### UnsetFreeBy
`func (o *ClusterBulkUpdate) UnsetFreeBy()`

UnsetFreeBy ensures that no value is present for FreeBy, not even an explicit nil
### GetGpus

`func (o *ClusterBulkUpdate) GetGpus() []ClusterGpusInner`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *ClusterBulkUpdate) GetGpusOk() (*[]ClusterGpusInner, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *ClusterBulkUpdate) SetGpus(v []ClusterGpusInner)`

SetGpus sets Gpus field to given value.


### GetGpuAlias

`func (o *ClusterBulkUpdate) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *ClusterBulkUpdate) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *ClusterBulkUpdate) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.


### GetGpuCount

`func (o *ClusterBulkUpdate) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ClusterBulkUpdate) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ClusterBulkUpdate) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.


### GetId

`func (o *ClusterBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *ClusterBulkUpdate) GetInstances() []ClusterInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterBulkUpdate) GetInstancesOk() (*[]ClusterInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterBulkUpdate) SetInstances(v []ClusterInstancesInner)`

SetInstances sets Instances field to given value.


### GetLastUsed

`func (o *ClusterBulkUpdate) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ClusterBulkUpdate) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ClusterBulkUpdate) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *ClusterBulkUpdate) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *ClusterBulkUpdate) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetMaintenance

`func (o *ClusterBulkUpdate) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *ClusterBulkUpdate) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *ClusterBulkUpdate) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *ClusterBulkUpdate) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetMgmtIp

`func (o *ClusterBulkUpdate) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *ClusterBulkUpdate) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *ClusterBulkUpdate) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *ClusterBulkUpdate) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *ClusterBulkUpdate) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *ClusterBulkUpdate) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *ClusterBulkUpdate) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *ClusterBulkUpdate) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### SetMgmtMacNil

`func (o *ClusterBulkUpdate) SetMgmtMacNil(b bool)`

 SetMgmtMacNil sets the value for MgmtMac to be an explicit nil

### UnsetMgmtMac
`func (o *ClusterBulkUpdate) UnsetMgmtMac()`

UnsetMgmtMac ensures that no value is present for MgmtMac, not even an explicit nil
### GetModified

`func (o *ClusterBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetmask

`func (o *ClusterBulkUpdate) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ClusterBulkUpdate) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ClusterBulkUpdate) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ClusterBulkUpdate) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *ClusterBulkUpdate) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *ClusterBulkUpdate) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetNodeCount

`func (o *ClusterBulkUpdate) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterBulkUpdate) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterBulkUpdate) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetNodes

`func (o *ClusterBulkUpdate) GetNodes() []ClusterNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterBulkUpdate) GetNodesOk() (*[]ClusterNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterBulkUpdate) SetNodes(v []ClusterNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterBulkUpdate) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNotes

`func (o *ClusterBulkUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ClusterBulkUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ClusterBulkUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ClusterBulkUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ClusterBulkUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ClusterBulkUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPersist

`func (o *ClusterBulkUpdate) GetPersist() bool`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *ClusterBulkUpdate) GetPersistOk() (*bool, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *ClusterBulkUpdate) SetPersist(v bool)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *ClusterBulkUpdate) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### GetProviderCapacity

`func (o *ClusterBulkUpdate) GetProviderCapacity() bool`

GetProviderCapacity returns the ProviderCapacity field if non-nil, zero value otherwise.

### GetProviderCapacityOk

`func (o *ClusterBulkUpdate) GetProviderCapacityOk() (*bool, bool)`

GetProviderCapacityOk returns a tuple with the ProviderCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCapacity

`func (o *ClusterBulkUpdate) SetProviderCapacity(v bool)`

SetProviderCapacity sets ProviderCapacity field to given value.


### GetProvisionUser

`func (o *ClusterBulkUpdate) GetProvisionUser() string`

GetProvisionUser returns the ProvisionUser field if non-nil, zero value otherwise.

### GetProvisionUserOk

`func (o *ClusterBulkUpdate) GetProvisionUserOk() (*string, bool)`

GetProvisionUserOk returns a tuple with the ProvisionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionUser

`func (o *ClusterBulkUpdate) SetProvisionUser(v string)`

SetProvisionUser sets ProvisionUser field to given value.

### HasProvisionUser

`func (o *ClusterBulkUpdate) HasProvisionUser() bool`

HasProvisionUser returns a boolean if a field has been set.

### SetProvisionUserNil

`func (o *ClusterBulkUpdate) SetProvisionUserNil(b bool)`

 SetProvisionUserNil sets the value for ProvisionUser to be an explicit nil

### UnsetProvisionUser
`func (o *ClusterBulkUpdate) UnsetProvisionUser()`

UnsetProvisionUser ensures that no value is present for ProvisionUser, not even an explicit nil
### GetProvisioningAttempts

`func (o *ClusterBulkUpdate) GetProvisioningAttempts() int32`

GetProvisioningAttempts returns the ProvisioningAttempts field if non-nil, zero value otherwise.

### GetProvisioningAttemptsOk

`func (o *ClusterBulkUpdate) GetProvisioningAttemptsOk() (*int32, bool)`

GetProvisioningAttemptsOk returns a tuple with the ProvisioningAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAttempts

`func (o *ClusterBulkUpdate) SetProvisioningAttempts(v int32)`

SetProvisioningAttempts sets ProvisioningAttempts field to given value.

### HasProvisioningAttempts

`func (o *ClusterBulkUpdate) HasProvisioningAttempts() bool`

HasProvisioningAttempts returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *ClusterBulkUpdate) GetProvisioningConfig() string`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ClusterBulkUpdate) GetProvisioningConfigOk() (*string, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ClusterBulkUpdate) SetProvisioningConfig(v string)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *ClusterBulkUpdate) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### SetProvisioningConfigNil

`func (o *ClusterBulkUpdate) SetProvisioningConfigNil(b bool)`

 SetProvisioningConfigNil sets the value for ProvisioningConfig to be an explicit nil

### UnsetProvisioningConfig
`func (o *ClusterBulkUpdate) UnsetProvisioningConfig()`

UnsetProvisioningConfig ensures that no value is present for ProvisioningConfig, not even an explicit nil
### GetProvisioningRequest

`func (o *ClusterBulkUpdate) GetProvisioningRequest() string`

GetProvisioningRequest returns the ProvisioningRequest field if non-nil, zero value otherwise.

### GetProvisioningRequestOk

`func (o *ClusterBulkUpdate) GetProvisioningRequestOk() (*string, bool)`

GetProvisioningRequestOk returns a tuple with the ProvisioningRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningRequest

`func (o *ClusterBulkUpdate) SetProvisioningRequest(v string)`

SetProvisioningRequest sets ProvisioningRequest field to given value.

### HasProvisioningRequest

`func (o *ClusterBulkUpdate) HasProvisioningRequest() bool`

HasProvisioningRequest returns a boolean if a field has been set.

### SetProvisioningRequestNil

`func (o *ClusterBulkUpdate) SetProvisioningRequestNil(b bool)`

 SetProvisioningRequestNil sets the value for ProvisioningRequest to be an explicit nil

### UnsetProvisioningRequest
`func (o *ClusterBulkUpdate) UnsetProvisioningRequest()`

UnsetProvisioningRequest ensures that no value is present for ProvisioningRequest, not even an explicit nil
### GetProvisioningState

`func (o *ClusterBulkUpdate) GetProvisioningState() ProvisioningStateEnum`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *ClusterBulkUpdate) GetProvisioningStateOk() (*ProvisioningStateEnum, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *ClusterBulkUpdate) SetProvisioningState(v ProvisioningStateEnum)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *ClusterBulkUpdate) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### GetPublicAddress

`func (o *ClusterBulkUpdate) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *ClusterBulkUpdate) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *ClusterBulkUpdate) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.


### SetPublicAddressNil

`func (o *ClusterBulkUpdate) SetPublicAddressNil(b bool)`

 SetPublicAddressNil sets the value for PublicAddress to be an explicit nil

### UnsetPublicAddress
`func (o *ClusterBulkUpdate) UnsetPublicAddress()`

UnsetPublicAddress ensures that no value is present for PublicAddress, not even an explicit nil
### GetRequestId

`func (o *ClusterBulkUpdate) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClusterBulkUpdate) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClusterBulkUpdate) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *ClusterBulkUpdate) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *ClusterBulkUpdate) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetReservation

`func (o *ClusterBulkUpdate) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *ClusterBulkUpdate) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *ClusterBulkUpdate) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *ClusterBulkUpdate) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetTenantIds

`func (o *ClusterBulkUpdate) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *ClusterBulkUpdate) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *ClusterBulkUpdate) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.


### GetVlanId

`func (o *ClusterBulkUpdate) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ClusterBulkUpdate) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ClusterBulkUpdate) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ClusterBulkUpdate) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetWorkshop

`func (o *ClusterBulkUpdate) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *ClusterBulkUpdate) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *ClusterBulkUpdate) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *ClusterBulkUpdate) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### GetWorkshopId

`func (o *ClusterBulkUpdate) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *ClusterBulkUpdate) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *ClusterBulkUpdate) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *ClusterBulkUpdate) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### SetWorkshopIdNil

`func (o *ClusterBulkUpdate) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *ClusterBulkUpdate) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil
### GetCount

`func (o *ClusterBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClusterBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClusterBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *ClusterBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ClusterBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ClusterBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *ClusterBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ClusterBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ClusterBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


