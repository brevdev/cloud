# Cluster

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
**GpuAlias** | Pointer to **string** | Alias for GPU plan (i.e. installed GPU type and count) | [optional] [readonly] 
**GpuCount** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | **string** |  | [readonly] 
**Instances** | [**[]ClusterInstancesInner**](ClusterInstancesInner.md) |  | 
**LastUsed** | **NullableTime** | Timestamp of when the cluster was last in use | [readonly] 
**Maintenance** | Pointer to **bool** | Is the cluster in maintenance mode? | [optional] 
**MgmtIp** | Pointer to **string** | Management IP address | [optional] 
**MgmtMac** | Pointer to **NullableString** | Management MAC address | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Netmask** | Pointer to **NullableInt32** | The subnet mask of the cluster&#39;s public IP address in CIDR notation | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] [readonly] 
**Nodes** | Pointer to [**[]ClusterNodesInner**](ClusterNodesInner.md) |  | [optional] 
**Notes** | Pointer to **NullableString** | Administrative comments about the cluster | [optional] 
**Persist** | Pointer to **bool** | Is the cluster exempt from provisioning_state timeouts? Can be used to ensure the cluster persists after a provisioning failure. | [optional] 
**ProviderCapacity** | **bool** | Does the provider have capacity to provision this cluster? | [readonly] 
**ProvisionUser** | Pointer to **NullableString** | Username used for provisioning this cluster | [optional] 
**ProvisioningAttempts** | Pointer to **int32** | The number of attempts that have been made to provision this cluster. Automatically resets to 0 after successful provisioning. | [optional] 
**ProvisioningConfig** | Pointer to **NullableString** | Applied provisioning configuration for the cluster | [optional] 
**ProvisioningRequest** | Pointer to **NullableString** | Requested provisioning configuration for the cluster | [optional] 
**ProvisioningState** | Pointer to [**ProvisioningStateEnum**](ProvisioningStateEnum.md) | Is the cluster currently provisioned?  * &#x60;deployed&#x60; - Cluster is in use by a deployment * &#x60;deploying&#x60; - Provisioning is in progress * &#x60;destroying&#x60; - Cluster is being destroyed * &#x60;pending&#x60; - Provisioning will begin soon * &#x60;ready&#x60; - Provisioning has completed and is ready for a deployment * &#x60;reserved&#x60; - Cluster is unprovisioned but reserved for later use * &#x60;unprovisioned&#x60; - Cluster has not yet been provisioned | [optional] 
**PublicAddress** | Pointer to **NullableString** | Public IP address or fully-qualified domain name of this cluster | [optional] 
**RequestId** | Pointer to **NullableString** | The request ID for the lab that is currently provisioned on this cluster (ex: TRY-1234) | [optional] 
**Reservation** | Pointer to **bool** | Is the cluster a static reservation from its provider? | [optional] 
**TenantIds** | **[]string** | Tenant UUID(s) that have been generated for this cluster during provisioning | 
**VlanId** | Pointer to **int32** | VLAN number | [optional] 
**Workshop** | Pointer to **bool** | Is the cluster set aside for use in a workshop? | [optional] 
**WorkshopId** | Pointer to **NullableString** | Identifier of the workshop this cluster is set aside for | [optional] 

## Methods

### NewCluster

`func NewCluster(available bool, created time.Time, deployment ClusterDeployment, gpus []ClusterGpusInner, id string, instances []ClusterInstancesInner, lastUsed NullableTime, modified time.Time, providerCapacity bool, tenantIds []string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Cluster) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Cluster) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Cluster) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetBastionName

`func (o *Cluster) GetBastionName() string`

GetBastionName returns the BastionName field if non-nil, zero value otherwise.

### GetBastionNameOk

`func (o *Cluster) GetBastionNameOk() (*string, bool)`

GetBastionNameOk returns a tuple with the BastionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionName

`func (o *Cluster) SetBastionName(v string)`

SetBastionName sets BastionName field to given value.

### HasBastionName

`func (o *Cluster) HasBastionName() bool`

HasBastionName returns a boolean if a field has been set.

### SetBastionNameNil

`func (o *Cluster) SetBastionNameNil(b bool)`

 SetBastionNameNil sets the value for BastionName to be an explicit nil

### UnsetBastionName
`func (o *Cluster) UnsetBastionName()`

UnsetBastionName ensures that no value is present for BastionName, not even an explicit nil
### GetCreated

`func (o *Cluster) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Cluster) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Cluster) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *Cluster) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *Cluster) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *Cluster) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetEnabled

`func (o *Cluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Cluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Cluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Cluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExperience

`func (o *Cluster) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *Cluster) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *Cluster) SetExperience(v string)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *Cluster) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *Cluster) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *Cluster) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetFreeBy

`func (o *Cluster) GetFreeBy() time.Time`

GetFreeBy returns the FreeBy field if non-nil, zero value otherwise.

### GetFreeByOk

`func (o *Cluster) GetFreeByOk() (*time.Time, bool)`

GetFreeByOk returns a tuple with the FreeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBy

`func (o *Cluster) SetFreeBy(v time.Time)`

SetFreeBy sets FreeBy field to given value.

### HasFreeBy

`func (o *Cluster) HasFreeBy() bool`

HasFreeBy returns a boolean if a field has been set.

### SetFreeByNil

`func (o *Cluster) SetFreeByNil(b bool)`

 SetFreeByNil sets the value for FreeBy to be an explicit nil

### UnsetFreeBy
`func (o *Cluster) UnsetFreeBy()`

UnsetFreeBy ensures that no value is present for FreeBy, not even an explicit nil
### GetGpus

`func (o *Cluster) GetGpus() []ClusterGpusInner`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *Cluster) GetGpusOk() (*[]ClusterGpusInner, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *Cluster) SetGpus(v []ClusterGpusInner)`

SetGpus sets Gpus field to given value.


### GetGpuAlias

`func (o *Cluster) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *Cluster) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *Cluster) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *Cluster) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### GetGpuCount

`func (o *Cluster) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Cluster) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Cluster) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *Cluster) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *Cluster) GetInstances() []ClusterInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Cluster) GetInstancesOk() (*[]ClusterInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Cluster) SetInstances(v []ClusterInstancesInner)`

SetInstances sets Instances field to given value.


### GetLastUsed

`func (o *Cluster) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *Cluster) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *Cluster) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *Cluster) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *Cluster) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetMaintenance

`func (o *Cluster) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *Cluster) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *Cluster) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *Cluster) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetMgmtIp

`func (o *Cluster) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *Cluster) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *Cluster) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *Cluster) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *Cluster) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *Cluster) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *Cluster) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *Cluster) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### SetMgmtMacNil

`func (o *Cluster) SetMgmtMacNil(b bool)`

 SetMgmtMacNil sets the value for MgmtMac to be an explicit nil

### UnsetMgmtMac
`func (o *Cluster) UnsetMgmtMac()`

UnsetMgmtMac ensures that no value is present for MgmtMac, not even an explicit nil
### GetModified

`func (o *Cluster) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Cluster) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Cluster) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetmask

`func (o *Cluster) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Cluster) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Cluster) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Cluster) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *Cluster) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *Cluster) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetNodeCount

`func (o *Cluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Cluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Cluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Cluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodes

`func (o *Cluster) GetNodes() []ClusterNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Cluster) GetNodesOk() (*[]ClusterNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Cluster) SetNodes(v []ClusterNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Cluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNotes

`func (o *Cluster) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Cluster) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Cluster) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Cluster) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Cluster) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Cluster) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPersist

`func (o *Cluster) GetPersist() bool`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *Cluster) GetPersistOk() (*bool, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *Cluster) SetPersist(v bool)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *Cluster) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### GetProviderCapacity

`func (o *Cluster) GetProviderCapacity() bool`

GetProviderCapacity returns the ProviderCapacity field if non-nil, zero value otherwise.

### GetProviderCapacityOk

`func (o *Cluster) GetProviderCapacityOk() (*bool, bool)`

GetProviderCapacityOk returns a tuple with the ProviderCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCapacity

`func (o *Cluster) SetProviderCapacity(v bool)`

SetProviderCapacity sets ProviderCapacity field to given value.


### GetProvisionUser

`func (o *Cluster) GetProvisionUser() string`

GetProvisionUser returns the ProvisionUser field if non-nil, zero value otherwise.

### GetProvisionUserOk

`func (o *Cluster) GetProvisionUserOk() (*string, bool)`

GetProvisionUserOk returns a tuple with the ProvisionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionUser

`func (o *Cluster) SetProvisionUser(v string)`

SetProvisionUser sets ProvisionUser field to given value.

### HasProvisionUser

`func (o *Cluster) HasProvisionUser() bool`

HasProvisionUser returns a boolean if a field has been set.

### SetProvisionUserNil

`func (o *Cluster) SetProvisionUserNil(b bool)`

 SetProvisionUserNil sets the value for ProvisionUser to be an explicit nil

### UnsetProvisionUser
`func (o *Cluster) UnsetProvisionUser()`

UnsetProvisionUser ensures that no value is present for ProvisionUser, not even an explicit nil
### GetProvisioningAttempts

`func (o *Cluster) GetProvisioningAttempts() int32`

GetProvisioningAttempts returns the ProvisioningAttempts field if non-nil, zero value otherwise.

### GetProvisioningAttemptsOk

`func (o *Cluster) GetProvisioningAttemptsOk() (*int32, bool)`

GetProvisioningAttemptsOk returns a tuple with the ProvisioningAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAttempts

`func (o *Cluster) SetProvisioningAttempts(v int32)`

SetProvisioningAttempts sets ProvisioningAttempts field to given value.

### HasProvisioningAttempts

`func (o *Cluster) HasProvisioningAttempts() bool`

HasProvisioningAttempts returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *Cluster) GetProvisioningConfig() string`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *Cluster) GetProvisioningConfigOk() (*string, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *Cluster) SetProvisioningConfig(v string)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *Cluster) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### SetProvisioningConfigNil

`func (o *Cluster) SetProvisioningConfigNil(b bool)`

 SetProvisioningConfigNil sets the value for ProvisioningConfig to be an explicit nil

### UnsetProvisioningConfig
`func (o *Cluster) UnsetProvisioningConfig()`

UnsetProvisioningConfig ensures that no value is present for ProvisioningConfig, not even an explicit nil
### GetProvisioningRequest

`func (o *Cluster) GetProvisioningRequest() string`

GetProvisioningRequest returns the ProvisioningRequest field if non-nil, zero value otherwise.

### GetProvisioningRequestOk

`func (o *Cluster) GetProvisioningRequestOk() (*string, bool)`

GetProvisioningRequestOk returns a tuple with the ProvisioningRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningRequest

`func (o *Cluster) SetProvisioningRequest(v string)`

SetProvisioningRequest sets ProvisioningRequest field to given value.

### HasProvisioningRequest

`func (o *Cluster) HasProvisioningRequest() bool`

HasProvisioningRequest returns a boolean if a field has been set.

### SetProvisioningRequestNil

`func (o *Cluster) SetProvisioningRequestNil(b bool)`

 SetProvisioningRequestNil sets the value for ProvisioningRequest to be an explicit nil

### UnsetProvisioningRequest
`func (o *Cluster) UnsetProvisioningRequest()`

UnsetProvisioningRequest ensures that no value is present for ProvisioningRequest, not even an explicit nil
### GetProvisioningState

`func (o *Cluster) GetProvisioningState() ProvisioningStateEnum`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *Cluster) GetProvisioningStateOk() (*ProvisioningStateEnum, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *Cluster) SetProvisioningState(v ProvisioningStateEnum)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *Cluster) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### GetPublicAddress

`func (o *Cluster) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *Cluster) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *Cluster) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *Cluster) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### SetPublicAddressNil

`func (o *Cluster) SetPublicAddressNil(b bool)`

 SetPublicAddressNil sets the value for PublicAddress to be an explicit nil

### UnsetPublicAddress
`func (o *Cluster) UnsetPublicAddress()`

UnsetPublicAddress ensures that no value is present for PublicAddress, not even an explicit nil
### GetRequestId

`func (o *Cluster) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Cluster) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Cluster) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Cluster) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *Cluster) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *Cluster) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetReservation

`func (o *Cluster) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Cluster) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Cluster) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Cluster) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetTenantIds

`func (o *Cluster) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *Cluster) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *Cluster) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.


### GetVlanId

`func (o *Cluster) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Cluster) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Cluster) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Cluster) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetWorkshop

`func (o *Cluster) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *Cluster) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *Cluster) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *Cluster) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### GetWorkshopId

`func (o *Cluster) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *Cluster) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *Cluster) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *Cluster) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.

### SetWorkshopIdNil

`func (o *Cluster) SetWorkshopIdNil(b bool)`

 SetWorkshopIdNil sets the value for WorkshopId to be an explicit nil

### UnsetWorkshopId
`func (o *Cluster) UnsetWorkshopId()`

UnsetWorkshopId ensures that no value is present for WorkshopId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


