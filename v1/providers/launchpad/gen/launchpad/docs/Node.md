# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BmcIp** | Pointer to **NullableString** | IP address of the BMC | [optional] 
**BmcMac** | Pointer to **NullableString** | MAC address of the BMC | [optional] 
**BmcPassword** | Pointer to **NullableString** | Password for the BMC | [optional] 
**BmcUser** | Pointer to **NullableString** | Username for the BMC | [optional] 
**Cluster** | Pointer to **NullableString** | UUID of the node&#39;s parent cluster | [optional] 
**Cpu** | Pointer to **int32** | Number of CPU cores installed | [optional] 
**CpuManufacturer** | Pointer to [**NullableCpuManufacturerEnum**](CpuManufacturerEnum.md) |  | [optional] 
**CpuModel** | Pointer to **NullableString** | Model information for the node&#39;s CPU | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**GarageId** | Pointer to **NullableString** | ID for the garage where a node is parked when unprovisioned | [optional] 
**Gpu** | [**ClusterGpusInner**](ClusterGpusInner.md) |  | 
**GpuAlias** | **string** | Alias for GPU plan (i.e. installed GPU type and count) | [readonly] 
**GpuCount** | Pointer to **int32** | Number of GPUs installed | [optional] 
**GpuModel** | **string** | Model of GPU(s) installed | [readonly] 
**GpuVbios** | Pointer to **NullableString** | VBIOS version used by installed GPU(s) | [optional] 
**Id** | **string** |  | [readonly] 
**Location** | [**NodeLocation**](NodeLocation.md) |  | 
**Memory** | Pointer to **int32** | Amount of RAM installed (in GB) | [optional] 
**MgmtIp** | Pointer to **string** | Management IP address | [optional] 
**MgmtMac** | Pointer to **NullableString** | Management MAC address | [optional] 
**Model** | Pointer to **NullableString** | Hardware model of the node | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NetworkType** | Pointer to [**NetworkTypeEnum**](NetworkTypeEnum.md) | Type of networking technology used  * &#x60;ethernet&#x60; - Ethernet * &#x60;infiniband&#x60; - InfiniBand | [optional] 
**NicPrefixes** | Pointer to **[]string** | Prefixes for the node&#39;s network interface(s) | [optional] 
**Notes** | Pointer to **NullableString** | Administrative comments about the node | [optional] 
**Oem** | Pointer to **NullableString** | UUID of the node&#39;s OEM | [optional] 
**ProviderNodeId** | Pointer to **NullableString** | Unique ID for this node assigned by its provider | [optional] 
**Rack** | Pointer to **NullableString** | Physical rack identifier | [optional] 
**RackUnit** | Pointer to **NullableInt32** | The rack unit (RU) within the rack where the node is installed | [optional] 
**SerialNumber** | Pointer to **NullableString** | Serial number of the node | [optional] 
**Storage** | [**[]NodeStorage**](NodeStorage.md) |  | 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**Tee** | Pointer to **bool** | Does the node support Trusted Execution Environment (TEE)? | [optional] 

## Methods

### NewNode

`func NewNode(created time.Time, gpu ClusterGpusInner, gpuAlias string, gpuModel string, id string, location NodeLocation, modified time.Time, storage []NodeStorage, systemArch SystemArchEnum, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmcIp

`func (o *Node) GetBmcIp() string`

GetBmcIp returns the BmcIp field if non-nil, zero value otherwise.

### GetBmcIpOk

`func (o *Node) GetBmcIpOk() (*string, bool)`

GetBmcIpOk returns a tuple with the BmcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcIp

`func (o *Node) SetBmcIp(v string)`

SetBmcIp sets BmcIp field to given value.

### HasBmcIp

`func (o *Node) HasBmcIp() bool`

HasBmcIp returns a boolean if a field has been set.

### SetBmcIpNil

`func (o *Node) SetBmcIpNil(b bool)`

 SetBmcIpNil sets the value for BmcIp to be an explicit nil

### UnsetBmcIp
`func (o *Node) UnsetBmcIp()`

UnsetBmcIp ensures that no value is present for BmcIp, not even an explicit nil
### GetBmcMac

`func (o *Node) GetBmcMac() string`

GetBmcMac returns the BmcMac field if non-nil, zero value otherwise.

### GetBmcMacOk

`func (o *Node) GetBmcMacOk() (*string, bool)`

GetBmcMacOk returns a tuple with the BmcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMac

`func (o *Node) SetBmcMac(v string)`

SetBmcMac sets BmcMac field to given value.

### HasBmcMac

`func (o *Node) HasBmcMac() bool`

HasBmcMac returns a boolean if a field has been set.

### SetBmcMacNil

`func (o *Node) SetBmcMacNil(b bool)`

 SetBmcMacNil sets the value for BmcMac to be an explicit nil

### UnsetBmcMac
`func (o *Node) UnsetBmcMac()`

UnsetBmcMac ensures that no value is present for BmcMac, not even an explicit nil
### GetBmcPassword

`func (o *Node) GetBmcPassword() string`

GetBmcPassword returns the BmcPassword field if non-nil, zero value otherwise.

### GetBmcPasswordOk

`func (o *Node) GetBmcPasswordOk() (*string, bool)`

GetBmcPasswordOk returns a tuple with the BmcPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcPassword

`func (o *Node) SetBmcPassword(v string)`

SetBmcPassword sets BmcPassword field to given value.

### HasBmcPassword

`func (o *Node) HasBmcPassword() bool`

HasBmcPassword returns a boolean if a field has been set.

### SetBmcPasswordNil

`func (o *Node) SetBmcPasswordNil(b bool)`

 SetBmcPasswordNil sets the value for BmcPassword to be an explicit nil

### UnsetBmcPassword
`func (o *Node) UnsetBmcPassword()`

UnsetBmcPassword ensures that no value is present for BmcPassword, not even an explicit nil
### GetBmcUser

`func (o *Node) GetBmcUser() string`

GetBmcUser returns the BmcUser field if non-nil, zero value otherwise.

### GetBmcUserOk

`func (o *Node) GetBmcUserOk() (*string, bool)`

GetBmcUserOk returns a tuple with the BmcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcUser

`func (o *Node) SetBmcUser(v string)`

SetBmcUser sets BmcUser field to given value.

### HasBmcUser

`func (o *Node) HasBmcUser() bool`

HasBmcUser returns a boolean if a field has been set.

### SetBmcUserNil

`func (o *Node) SetBmcUserNil(b bool)`

 SetBmcUserNil sets the value for BmcUser to be an explicit nil

### UnsetBmcUser
`func (o *Node) UnsetBmcUser()`

UnsetBmcUser ensures that no value is present for BmcUser, not even an explicit nil
### GetCluster

`func (o *Node) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Node) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Node) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Node) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *Node) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *Node) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetCpu

`func (o *Node) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Node) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Node) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Node) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpuManufacturer

`func (o *Node) GetCpuManufacturer() CpuManufacturerEnum`

GetCpuManufacturer returns the CpuManufacturer field if non-nil, zero value otherwise.

### GetCpuManufacturerOk

`func (o *Node) GetCpuManufacturerOk() (*CpuManufacturerEnum, bool)`

GetCpuManufacturerOk returns a tuple with the CpuManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuManufacturer

`func (o *Node) SetCpuManufacturer(v CpuManufacturerEnum)`

SetCpuManufacturer sets CpuManufacturer field to given value.

### HasCpuManufacturer

`func (o *Node) HasCpuManufacturer() bool`

HasCpuManufacturer returns a boolean if a field has been set.

### SetCpuManufacturerNil

`func (o *Node) SetCpuManufacturerNil(b bool)`

 SetCpuManufacturerNil sets the value for CpuManufacturer to be an explicit nil

### UnsetCpuManufacturer
`func (o *Node) UnsetCpuManufacturer()`

UnsetCpuManufacturer ensures that no value is present for CpuManufacturer, not even an explicit nil
### GetCpuModel

`func (o *Node) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *Node) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *Node) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *Node) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### SetCpuModelNil

`func (o *Node) SetCpuModelNil(b bool)`

 SetCpuModelNil sets the value for CpuModel to be an explicit nil

### UnsetCpuModel
`func (o *Node) UnsetCpuModel()`

UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
### GetCreated

`func (o *Node) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Node) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Node) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetGarageId

`func (o *Node) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *Node) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *Node) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *Node) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *Node) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *Node) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGpu

`func (o *Node) GetGpu() ClusterGpusInner`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *Node) GetGpuOk() (*ClusterGpusInner, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *Node) SetGpu(v ClusterGpusInner)`

SetGpu sets Gpu field to given value.


### GetGpuAlias

`func (o *Node) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *Node) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *Node) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.


### GetGpuCount

`func (o *Node) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *Node) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *Node) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *Node) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *Node) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *Node) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *Node) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.


### GetGpuVbios

`func (o *Node) GetGpuVbios() string`

GetGpuVbios returns the GpuVbios field if non-nil, zero value otherwise.

### GetGpuVbiosOk

`func (o *Node) GetGpuVbiosOk() (*string, bool)`

GetGpuVbiosOk returns a tuple with the GpuVbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuVbios

`func (o *Node) SetGpuVbios(v string)`

SetGpuVbios sets GpuVbios field to given value.

### HasGpuVbios

`func (o *Node) HasGpuVbios() bool`

HasGpuVbios returns a boolean if a field has been set.

### SetGpuVbiosNil

`func (o *Node) SetGpuVbiosNil(b bool)`

 SetGpuVbiosNil sets the value for GpuVbios to be an explicit nil

### UnsetGpuVbios
`func (o *Node) UnsetGpuVbios()`

UnsetGpuVbios ensures that no value is present for GpuVbios, not even an explicit nil
### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *Node) GetLocation() NodeLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Node) GetLocationOk() (*NodeLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Node) SetLocation(v NodeLocation)`

SetLocation sets Location field to given value.


### GetMemory

`func (o *Node) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Node) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Node) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Node) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMgmtIp

`func (o *Node) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *Node) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *Node) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *Node) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *Node) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *Node) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *Node) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *Node) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### SetMgmtMacNil

`func (o *Node) SetMgmtMacNil(b bool)`

 SetMgmtMacNil sets the value for MgmtMac to be an explicit nil

### UnsetMgmtMac
`func (o *Node) UnsetMgmtMac()`

UnsetMgmtMac ensures that no value is present for MgmtMac, not even an explicit nil
### GetModel

`func (o *Node) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Node) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Node) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Node) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Node) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Node) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetModified

`func (o *Node) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Node) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Node) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetworkType

`func (o *Node) GetNetworkType() NetworkTypeEnum`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Node) GetNetworkTypeOk() (*NetworkTypeEnum, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Node) SetNetworkType(v NetworkTypeEnum)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Node) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNicPrefixes

`func (o *Node) GetNicPrefixes() []string`

GetNicPrefixes returns the NicPrefixes field if non-nil, zero value otherwise.

### GetNicPrefixesOk

`func (o *Node) GetNicPrefixesOk() (*[]string, bool)`

GetNicPrefixesOk returns a tuple with the NicPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrefixes

`func (o *Node) SetNicPrefixes(v []string)`

SetNicPrefixes sets NicPrefixes field to given value.

### HasNicPrefixes

`func (o *Node) HasNicPrefixes() bool`

HasNicPrefixes returns a boolean if a field has been set.

### GetNotes

`func (o *Node) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Node) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Node) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Node) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Node) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Node) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOem

`func (o *Node) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *Node) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *Node) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *Node) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *Node) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *Node) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetProviderNodeId

`func (o *Node) GetProviderNodeId() string`

GetProviderNodeId returns the ProviderNodeId field if non-nil, zero value otherwise.

### GetProviderNodeIdOk

`func (o *Node) GetProviderNodeIdOk() (*string, bool)`

GetProviderNodeIdOk returns a tuple with the ProviderNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNodeId

`func (o *Node) SetProviderNodeId(v string)`

SetProviderNodeId sets ProviderNodeId field to given value.

### HasProviderNodeId

`func (o *Node) HasProviderNodeId() bool`

HasProviderNodeId returns a boolean if a field has been set.

### SetProviderNodeIdNil

`func (o *Node) SetProviderNodeIdNil(b bool)`

 SetProviderNodeIdNil sets the value for ProviderNodeId to be an explicit nil

### UnsetProviderNodeId
`func (o *Node) UnsetProviderNodeId()`

UnsetProviderNodeId ensures that no value is present for ProviderNodeId, not even an explicit nil
### GetRack

`func (o *Node) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Node) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Node) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Node) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *Node) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *Node) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetRackUnit

`func (o *Node) GetRackUnit() int32`

GetRackUnit returns the RackUnit field if non-nil, zero value otherwise.

### GetRackUnitOk

`func (o *Node) GetRackUnitOk() (*int32, bool)`

GetRackUnitOk returns a tuple with the RackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnit

`func (o *Node) SetRackUnit(v int32)`

SetRackUnit sets RackUnit field to given value.

### HasRackUnit

`func (o *Node) HasRackUnit() bool`

HasRackUnit returns a boolean if a field has been set.

### SetRackUnitNil

`func (o *Node) SetRackUnitNil(b bool)`

 SetRackUnitNil sets the value for RackUnit to be an explicit nil

### UnsetRackUnit
`func (o *Node) UnsetRackUnit()`

UnsetRackUnit ensures that no value is present for RackUnit, not even an explicit nil
### GetSerialNumber

`func (o *Node) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Node) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Node) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Node) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *Node) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *Node) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStorage

`func (o *Node) GetStorage() []NodeStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Node) GetStorageOk() (*[]NodeStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Node) SetStorage(v []NodeStorage)`

SetStorage sets Storage field to given value.


### GetSystemArch

`func (o *Node) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *Node) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *Node) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetTee

`func (o *Node) GetTee() bool`

GetTee returns the Tee field if non-nil, zero value otherwise.

### GetTeeOk

`func (o *Node) GetTeeOk() (*bool, bool)`

GetTeeOk returns a tuple with the Tee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTee

`func (o *Node) SetTee(v bool)`

SetTee sets Tee field to given value.

### HasTee

`func (o *Node) HasTee() bool`

HasTee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


