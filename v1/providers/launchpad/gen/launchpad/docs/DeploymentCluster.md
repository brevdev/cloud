# DeploymentCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Is the cluster currently available for provisioning? | [readonly] 
**BastionName** | Pointer to **string** | Name of the bastion assigned to the cluster | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Deployment** | [**ClusterDeployment**](ClusterDeployment.md) |  | 
**Enabled** | Pointer to **bool** | Is the cluster administratively enabled? | [optional] 
**Experience** | Pointer to **string** | Experience provisioned onto this cluster | [optional] 
**FreeBy** | Pointer to **time.Time** |  | [optional] 
**Gpus** | [**[]ClusterGpusInner**](ClusterGpusInner.md) |  | 
**GpuAlias** | Pointer to **string** | Alias for GPU plan (i.e. installed GPU type and count) | [optional] [readonly] 
**GpuCount** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | **string** |  | [readonly] 
**Instances** | [**[]ClusterInstancesInner**](ClusterInstancesInner.md) |  | 
**LastUsed** | **time.Time** | Timestamp of when the cluster was last in use | [readonly] 
**Maintenance** | Pointer to **bool** | Is the cluster in maintenance mode? | [optional] 
**MgmtIp** | Pointer to **string** | Management IP address | [optional] 
**MgmtMac** | Pointer to **string** | Management MAC address | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Netmask** | Pointer to **int32** | The subnet mask of the cluster&#39;s public IP address in CIDR notation | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] [readonly] 
**Nodes** | Pointer to [**[]ClusterNodesInner**](ClusterNodesInner.md) |  | [optional] 
**Notes** | Pointer to **string** | Administrative comments about the cluster | [optional] 
**Persist** | Pointer to **bool** | Is the cluster exempt from provisioning_state timeouts? Can be used to ensure the cluster persists after a provisioning failure. | [optional] 
**ProviderCapacity** | **bool** | Does the provider have capacity to provision this cluster? | [readonly] 
**ProvisionUser** | Pointer to **string** | Username used for provisioning this cluster | [optional] 
**ProvisioningAttempts** | Pointer to **int32** | The number of attempts that have been made to provision this cluster. Automatically resets to 0 after successful provisioning. | [optional] 
**ProvisioningConfig** | Pointer to **string** | Applied provisioning configuration for the cluster | [optional] 
**ProvisioningRequest** | Pointer to **string** | Requested provisioning configuration for the cluster | [optional] 
**ProvisioningState** | Pointer to [**ProvisioningStateEnum**](ProvisioningStateEnum.md) | Is the cluster currently provisioned?  * &#x60;deployed&#x60; - Cluster is in use by a deployment * &#x60;deploying&#x60; - Provisioning is in progress * &#x60;destroying&#x60; - Cluster is being destroyed * &#x60;pending&#x60; - Provisioning will begin soon * &#x60;ready&#x60; - Provisioning has completed and is ready for a deployment * &#x60;reserved&#x60; - Cluster is unprovisioned but reserved for later use * &#x60;unprovisioned&#x60; - Cluster has not yet been provisioned | [optional] 
**PublicAddress** | Pointer to **string** | Public IP address or fully-qualified domain name of this cluster | [optional] 
**RequestId** | Pointer to **string** | The request ID for the lab that is currently provisioned on this cluster (ex: TRY-1234) | [optional] 
**Reservation** | Pointer to **bool** | Is the cluster a static reservation from its provider? | [optional] 
**TenantIds** | **[]string** | Tenant UUID(s) that have been generated for this cluster during provisioning | 
**VlanId** | Pointer to **int32** | VLAN number | [optional] 
**Workshop** | Pointer to **bool** | Is the cluster set aside for use in a workshop? | [optional] 
**WorkshopId** | Pointer to **string** | Identifier of the workshop this cluster is set aside for | [optional] 

## Methods

### NewDeploymentCluster

`func NewDeploymentCluster(available bool, created time.Time, deployment ClusterDeployment, gpus []ClusterGpusInner, id string, instances []ClusterInstancesInner, lastUsed time.Time, modified time.Time, providerCapacity bool, tenantIds []string, ) *DeploymentCluster`

NewDeploymentCluster instantiates a new DeploymentCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentClusterWithDefaults

`func NewDeploymentClusterWithDefaults() *DeploymentCluster`

NewDeploymentClusterWithDefaults instantiates a new DeploymentCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *DeploymentCluster) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DeploymentCluster) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DeploymentCluster) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetBastionName

`func (o *DeploymentCluster) GetBastionName() string`

GetBastionName returns the BastionName field if non-nil, zero value otherwise.

### GetBastionNameOk

`func (o *DeploymentCluster) GetBastionNameOk() (*string, bool)`

GetBastionNameOk returns a tuple with the BastionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionName

`func (o *DeploymentCluster) SetBastionName(v string)`

SetBastionName sets BastionName field to given value.

### HasBastionName

`func (o *DeploymentCluster) HasBastionName() bool`

HasBastionName returns a boolean if a field has been set.

### GetCreated

`func (o *DeploymentCluster) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentCluster) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentCluster) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *DeploymentCluster) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *DeploymentCluster) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *DeploymentCluster) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetEnabled

`func (o *DeploymentCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeploymentCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeploymentCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeploymentCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExperience

`func (o *DeploymentCluster) GetExperience() string`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *DeploymentCluster) GetExperienceOk() (*string, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *DeploymentCluster) SetExperience(v string)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *DeploymentCluster) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetFreeBy

`func (o *DeploymentCluster) GetFreeBy() time.Time`

GetFreeBy returns the FreeBy field if non-nil, zero value otherwise.

### GetFreeByOk

`func (o *DeploymentCluster) GetFreeByOk() (*time.Time, bool)`

GetFreeByOk returns a tuple with the FreeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBy

`func (o *DeploymentCluster) SetFreeBy(v time.Time)`

SetFreeBy sets FreeBy field to given value.

### HasFreeBy

`func (o *DeploymentCluster) HasFreeBy() bool`

HasFreeBy returns a boolean if a field has been set.

### GetGpus

`func (o *DeploymentCluster) GetGpus() []ClusterGpusInner`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *DeploymentCluster) GetGpusOk() (*[]ClusterGpusInner, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *DeploymentCluster) SetGpus(v []ClusterGpusInner)`

SetGpus sets Gpus field to given value.


### GetGpuAlias

`func (o *DeploymentCluster) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *DeploymentCluster) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *DeploymentCluster) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.

### HasGpuAlias

`func (o *DeploymentCluster) HasGpuAlias() bool`

HasGpuAlias returns a boolean if a field has been set.

### GetGpuCount

`func (o *DeploymentCluster) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *DeploymentCluster) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *DeploymentCluster) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *DeploymentCluster) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetId

`func (o *DeploymentCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentCluster) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *DeploymentCluster) GetInstances() []ClusterInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DeploymentCluster) GetInstancesOk() (*[]ClusterInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DeploymentCluster) SetInstances(v []ClusterInstancesInner)`

SetInstances sets Instances field to given value.


### GetLastUsed

`func (o *DeploymentCluster) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *DeploymentCluster) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *DeploymentCluster) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetMaintenance

`func (o *DeploymentCluster) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *DeploymentCluster) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *DeploymentCluster) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *DeploymentCluster) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetMgmtIp

`func (o *DeploymentCluster) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *DeploymentCluster) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *DeploymentCluster) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *DeploymentCluster) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *DeploymentCluster) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *DeploymentCluster) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *DeploymentCluster) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *DeploymentCluster) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### GetModified

`func (o *DeploymentCluster) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentCluster) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentCluster) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetmask

`func (o *DeploymentCluster) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *DeploymentCluster) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *DeploymentCluster) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *DeploymentCluster) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNodeCount

`func (o *DeploymentCluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *DeploymentCluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *DeploymentCluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *DeploymentCluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodes

`func (o *DeploymentCluster) GetNodes() []ClusterNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *DeploymentCluster) GetNodesOk() (*[]ClusterNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *DeploymentCluster) SetNodes(v []ClusterNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *DeploymentCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNotes

`func (o *DeploymentCluster) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeploymentCluster) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeploymentCluster) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DeploymentCluster) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPersist

`func (o *DeploymentCluster) GetPersist() bool`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *DeploymentCluster) GetPersistOk() (*bool, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *DeploymentCluster) SetPersist(v bool)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *DeploymentCluster) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### GetProviderCapacity

`func (o *DeploymentCluster) GetProviderCapacity() bool`

GetProviderCapacity returns the ProviderCapacity field if non-nil, zero value otherwise.

### GetProviderCapacityOk

`func (o *DeploymentCluster) GetProviderCapacityOk() (*bool, bool)`

GetProviderCapacityOk returns a tuple with the ProviderCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCapacity

`func (o *DeploymentCluster) SetProviderCapacity(v bool)`

SetProviderCapacity sets ProviderCapacity field to given value.


### GetProvisionUser

`func (o *DeploymentCluster) GetProvisionUser() string`

GetProvisionUser returns the ProvisionUser field if non-nil, zero value otherwise.

### GetProvisionUserOk

`func (o *DeploymentCluster) GetProvisionUserOk() (*string, bool)`

GetProvisionUserOk returns a tuple with the ProvisionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionUser

`func (o *DeploymentCluster) SetProvisionUser(v string)`

SetProvisionUser sets ProvisionUser field to given value.

### HasProvisionUser

`func (o *DeploymentCluster) HasProvisionUser() bool`

HasProvisionUser returns a boolean if a field has been set.

### GetProvisioningAttempts

`func (o *DeploymentCluster) GetProvisioningAttempts() int32`

GetProvisioningAttempts returns the ProvisioningAttempts field if non-nil, zero value otherwise.

### GetProvisioningAttemptsOk

`func (o *DeploymentCluster) GetProvisioningAttemptsOk() (*int32, bool)`

GetProvisioningAttemptsOk returns a tuple with the ProvisioningAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAttempts

`func (o *DeploymentCluster) SetProvisioningAttempts(v int32)`

SetProvisioningAttempts sets ProvisioningAttempts field to given value.

### HasProvisioningAttempts

`func (o *DeploymentCluster) HasProvisioningAttempts() bool`

HasProvisioningAttempts returns a boolean if a field has been set.

### GetProvisioningConfig

`func (o *DeploymentCluster) GetProvisioningConfig() string`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *DeploymentCluster) GetProvisioningConfigOk() (*string, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *DeploymentCluster) SetProvisioningConfig(v string)`

SetProvisioningConfig sets ProvisioningConfig field to given value.

### HasProvisioningConfig

`func (o *DeploymentCluster) HasProvisioningConfig() bool`

HasProvisioningConfig returns a boolean if a field has been set.

### GetProvisioningRequest

`func (o *DeploymentCluster) GetProvisioningRequest() string`

GetProvisioningRequest returns the ProvisioningRequest field if non-nil, zero value otherwise.

### GetProvisioningRequestOk

`func (o *DeploymentCluster) GetProvisioningRequestOk() (*string, bool)`

GetProvisioningRequestOk returns a tuple with the ProvisioningRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningRequest

`func (o *DeploymentCluster) SetProvisioningRequest(v string)`

SetProvisioningRequest sets ProvisioningRequest field to given value.

### HasProvisioningRequest

`func (o *DeploymentCluster) HasProvisioningRequest() bool`

HasProvisioningRequest returns a boolean if a field has been set.

### GetProvisioningState

`func (o *DeploymentCluster) GetProvisioningState() ProvisioningStateEnum`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *DeploymentCluster) GetProvisioningStateOk() (*ProvisioningStateEnum, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *DeploymentCluster) SetProvisioningState(v ProvisioningStateEnum)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *DeploymentCluster) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### GetPublicAddress

`func (o *DeploymentCluster) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *DeploymentCluster) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *DeploymentCluster) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *DeploymentCluster) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### GetRequestId

`func (o *DeploymentCluster) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeploymentCluster) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeploymentCluster) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DeploymentCluster) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetReservation

`func (o *DeploymentCluster) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *DeploymentCluster) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *DeploymentCluster) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *DeploymentCluster) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetTenantIds

`func (o *DeploymentCluster) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *DeploymentCluster) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *DeploymentCluster) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.


### GetVlanId

`func (o *DeploymentCluster) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DeploymentCluster) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DeploymentCluster) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DeploymentCluster) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetWorkshop

`func (o *DeploymentCluster) GetWorkshop() bool`

GetWorkshop returns the Workshop field if non-nil, zero value otherwise.

### GetWorkshopOk

`func (o *DeploymentCluster) GetWorkshopOk() (*bool, bool)`

GetWorkshopOk returns a tuple with the Workshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshop

`func (o *DeploymentCluster) SetWorkshop(v bool)`

SetWorkshop sets Workshop field to given value.

### HasWorkshop

`func (o *DeploymentCluster) HasWorkshop() bool`

HasWorkshop returns a boolean if a field has been set.

### GetWorkshopId

`func (o *DeploymentCluster) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *DeploymentCluster) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *DeploymentCluster) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.

### HasWorkshopId

`func (o *DeploymentCluster) HasWorkshopId() bool`

HasWorkshopId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


