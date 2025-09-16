# NodeBulkUpdate

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
**SerialNumber** | **NullableString** | Serial number of the node | [readonly] 
**Storage** | [**[]NodeStorage**](NodeStorage.md) |  | 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**Tee** | Pointer to **bool** | Does the node support Trusted Execution Environment (TEE)? | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewNodeBulkUpdate

`func NewNodeBulkUpdate(created time.Time, gpu ClusterGpusInner, gpuAlias string, gpuModel string, id string, location NodeLocation, modified time.Time, serialNumber NullableString, storage []NodeStorage, systemArch SystemArchEnum, count int32, ids []string, result string, ) *NodeBulkUpdate`

NewNodeBulkUpdate instantiates a new NodeBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBulkUpdateWithDefaults

`func NewNodeBulkUpdateWithDefaults() *NodeBulkUpdate`

NewNodeBulkUpdateWithDefaults instantiates a new NodeBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmcIp

`func (o *NodeBulkUpdate) GetBmcIp() string`

GetBmcIp returns the BmcIp field if non-nil, zero value otherwise.

### GetBmcIpOk

`func (o *NodeBulkUpdate) GetBmcIpOk() (*string, bool)`

GetBmcIpOk returns a tuple with the BmcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcIp

`func (o *NodeBulkUpdate) SetBmcIp(v string)`

SetBmcIp sets BmcIp field to given value.

### HasBmcIp

`func (o *NodeBulkUpdate) HasBmcIp() bool`

HasBmcIp returns a boolean if a field has been set.

### SetBmcIpNil

`func (o *NodeBulkUpdate) SetBmcIpNil(b bool)`

 SetBmcIpNil sets the value for BmcIp to be an explicit nil

### UnsetBmcIp
`func (o *NodeBulkUpdate) UnsetBmcIp()`

UnsetBmcIp ensures that no value is present for BmcIp, not even an explicit nil
### GetBmcMac

`func (o *NodeBulkUpdate) GetBmcMac() string`

GetBmcMac returns the BmcMac field if non-nil, zero value otherwise.

### GetBmcMacOk

`func (o *NodeBulkUpdate) GetBmcMacOk() (*string, bool)`

GetBmcMacOk returns a tuple with the BmcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMac

`func (o *NodeBulkUpdate) SetBmcMac(v string)`

SetBmcMac sets BmcMac field to given value.

### HasBmcMac

`func (o *NodeBulkUpdate) HasBmcMac() bool`

HasBmcMac returns a boolean if a field has been set.

### SetBmcMacNil

`func (o *NodeBulkUpdate) SetBmcMacNil(b bool)`

 SetBmcMacNil sets the value for BmcMac to be an explicit nil

### UnsetBmcMac
`func (o *NodeBulkUpdate) UnsetBmcMac()`

UnsetBmcMac ensures that no value is present for BmcMac, not even an explicit nil
### GetBmcPassword

`func (o *NodeBulkUpdate) GetBmcPassword() string`

GetBmcPassword returns the BmcPassword field if non-nil, zero value otherwise.

### GetBmcPasswordOk

`func (o *NodeBulkUpdate) GetBmcPasswordOk() (*string, bool)`

GetBmcPasswordOk returns a tuple with the BmcPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcPassword

`func (o *NodeBulkUpdate) SetBmcPassword(v string)`

SetBmcPassword sets BmcPassword field to given value.

### HasBmcPassword

`func (o *NodeBulkUpdate) HasBmcPassword() bool`

HasBmcPassword returns a boolean if a field has been set.

### SetBmcPasswordNil

`func (o *NodeBulkUpdate) SetBmcPasswordNil(b bool)`

 SetBmcPasswordNil sets the value for BmcPassword to be an explicit nil

### UnsetBmcPassword
`func (o *NodeBulkUpdate) UnsetBmcPassword()`

UnsetBmcPassword ensures that no value is present for BmcPassword, not even an explicit nil
### GetBmcUser

`func (o *NodeBulkUpdate) GetBmcUser() string`

GetBmcUser returns the BmcUser field if non-nil, zero value otherwise.

### GetBmcUserOk

`func (o *NodeBulkUpdate) GetBmcUserOk() (*string, bool)`

GetBmcUserOk returns a tuple with the BmcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcUser

`func (o *NodeBulkUpdate) SetBmcUser(v string)`

SetBmcUser sets BmcUser field to given value.

### HasBmcUser

`func (o *NodeBulkUpdate) HasBmcUser() bool`

HasBmcUser returns a boolean if a field has been set.

### SetBmcUserNil

`func (o *NodeBulkUpdate) SetBmcUserNil(b bool)`

 SetBmcUserNil sets the value for BmcUser to be an explicit nil

### UnsetBmcUser
`func (o *NodeBulkUpdate) UnsetBmcUser()`

UnsetBmcUser ensures that no value is present for BmcUser, not even an explicit nil
### GetCluster

`func (o *NodeBulkUpdate) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NodeBulkUpdate) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NodeBulkUpdate) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NodeBulkUpdate) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *NodeBulkUpdate) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *NodeBulkUpdate) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetCpu

`func (o *NodeBulkUpdate) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeBulkUpdate) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeBulkUpdate) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeBulkUpdate) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpuManufacturer

`func (o *NodeBulkUpdate) GetCpuManufacturer() CpuManufacturerEnum`

GetCpuManufacturer returns the CpuManufacturer field if non-nil, zero value otherwise.

### GetCpuManufacturerOk

`func (o *NodeBulkUpdate) GetCpuManufacturerOk() (*CpuManufacturerEnum, bool)`

GetCpuManufacturerOk returns a tuple with the CpuManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuManufacturer

`func (o *NodeBulkUpdate) SetCpuManufacturer(v CpuManufacturerEnum)`

SetCpuManufacturer sets CpuManufacturer field to given value.

### HasCpuManufacturer

`func (o *NodeBulkUpdate) HasCpuManufacturer() bool`

HasCpuManufacturer returns a boolean if a field has been set.

### SetCpuManufacturerNil

`func (o *NodeBulkUpdate) SetCpuManufacturerNil(b bool)`

 SetCpuManufacturerNil sets the value for CpuManufacturer to be an explicit nil

### UnsetCpuManufacturer
`func (o *NodeBulkUpdate) UnsetCpuManufacturer()`

UnsetCpuManufacturer ensures that no value is present for CpuManufacturer, not even an explicit nil
### GetCpuModel

`func (o *NodeBulkUpdate) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *NodeBulkUpdate) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *NodeBulkUpdate) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *NodeBulkUpdate) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### SetCpuModelNil

`func (o *NodeBulkUpdate) SetCpuModelNil(b bool)`

 SetCpuModelNil sets the value for CpuModel to be an explicit nil

### UnsetCpuModel
`func (o *NodeBulkUpdate) UnsetCpuModel()`

UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
### GetCreated

`func (o *NodeBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NodeBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NodeBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetGarageId

`func (o *NodeBulkUpdate) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *NodeBulkUpdate) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *NodeBulkUpdate) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *NodeBulkUpdate) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### SetGarageIdNil

`func (o *NodeBulkUpdate) SetGarageIdNil(b bool)`

 SetGarageIdNil sets the value for GarageId to be an explicit nil

### UnsetGarageId
`func (o *NodeBulkUpdate) UnsetGarageId()`

UnsetGarageId ensures that no value is present for GarageId, not even an explicit nil
### GetGpu

`func (o *NodeBulkUpdate) GetGpu() ClusterGpusInner`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *NodeBulkUpdate) GetGpuOk() (*ClusterGpusInner, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *NodeBulkUpdate) SetGpu(v ClusterGpusInner)`

SetGpu sets Gpu field to given value.


### GetGpuAlias

`func (o *NodeBulkUpdate) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *NodeBulkUpdate) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *NodeBulkUpdate) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.


### GetGpuCount

`func (o *NodeBulkUpdate) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *NodeBulkUpdate) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *NodeBulkUpdate) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *NodeBulkUpdate) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *NodeBulkUpdate) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *NodeBulkUpdate) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *NodeBulkUpdate) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.


### GetGpuVbios

`func (o *NodeBulkUpdate) GetGpuVbios() string`

GetGpuVbios returns the GpuVbios field if non-nil, zero value otherwise.

### GetGpuVbiosOk

`func (o *NodeBulkUpdate) GetGpuVbiosOk() (*string, bool)`

GetGpuVbiosOk returns a tuple with the GpuVbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuVbios

`func (o *NodeBulkUpdate) SetGpuVbios(v string)`

SetGpuVbios sets GpuVbios field to given value.

### HasGpuVbios

`func (o *NodeBulkUpdate) HasGpuVbios() bool`

HasGpuVbios returns a boolean if a field has been set.

### SetGpuVbiosNil

`func (o *NodeBulkUpdate) SetGpuVbiosNil(b bool)`

 SetGpuVbiosNil sets the value for GpuVbios to be an explicit nil

### UnsetGpuVbios
`func (o *NodeBulkUpdate) UnsetGpuVbios()`

UnsetGpuVbios ensures that no value is present for GpuVbios, not even an explicit nil
### GetId

`func (o *NodeBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *NodeBulkUpdate) GetLocation() NodeLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NodeBulkUpdate) GetLocationOk() (*NodeLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NodeBulkUpdate) SetLocation(v NodeLocation)`

SetLocation sets Location field to given value.


### GetMemory

`func (o *NodeBulkUpdate) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeBulkUpdate) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeBulkUpdate) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NodeBulkUpdate) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMgmtIp

`func (o *NodeBulkUpdate) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *NodeBulkUpdate) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *NodeBulkUpdate) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *NodeBulkUpdate) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *NodeBulkUpdate) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *NodeBulkUpdate) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *NodeBulkUpdate) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *NodeBulkUpdate) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### SetMgmtMacNil

`func (o *NodeBulkUpdate) SetMgmtMacNil(b bool)`

 SetMgmtMacNil sets the value for MgmtMac to be an explicit nil

### UnsetMgmtMac
`func (o *NodeBulkUpdate) UnsetMgmtMac()`

UnsetMgmtMac ensures that no value is present for MgmtMac, not even an explicit nil
### GetModel

`func (o *NodeBulkUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NodeBulkUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NodeBulkUpdate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NodeBulkUpdate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *NodeBulkUpdate) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *NodeBulkUpdate) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetModified

`func (o *NodeBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NodeBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NodeBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetworkType

`func (o *NodeBulkUpdate) GetNetworkType() NetworkTypeEnum`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *NodeBulkUpdate) GetNetworkTypeOk() (*NetworkTypeEnum, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *NodeBulkUpdate) SetNetworkType(v NetworkTypeEnum)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *NodeBulkUpdate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNicPrefixes

`func (o *NodeBulkUpdate) GetNicPrefixes() []string`

GetNicPrefixes returns the NicPrefixes field if non-nil, zero value otherwise.

### GetNicPrefixesOk

`func (o *NodeBulkUpdate) GetNicPrefixesOk() (*[]string, bool)`

GetNicPrefixesOk returns a tuple with the NicPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrefixes

`func (o *NodeBulkUpdate) SetNicPrefixes(v []string)`

SetNicPrefixes sets NicPrefixes field to given value.

### HasNicPrefixes

`func (o *NodeBulkUpdate) HasNicPrefixes() bool`

HasNicPrefixes returns a boolean if a field has been set.

### GetNotes

`func (o *NodeBulkUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *NodeBulkUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *NodeBulkUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *NodeBulkUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *NodeBulkUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *NodeBulkUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOem

`func (o *NodeBulkUpdate) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *NodeBulkUpdate) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *NodeBulkUpdate) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *NodeBulkUpdate) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *NodeBulkUpdate) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *NodeBulkUpdate) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetProviderNodeId

`func (o *NodeBulkUpdate) GetProviderNodeId() string`

GetProviderNodeId returns the ProviderNodeId field if non-nil, zero value otherwise.

### GetProviderNodeIdOk

`func (o *NodeBulkUpdate) GetProviderNodeIdOk() (*string, bool)`

GetProviderNodeIdOk returns a tuple with the ProviderNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNodeId

`func (o *NodeBulkUpdate) SetProviderNodeId(v string)`

SetProviderNodeId sets ProviderNodeId field to given value.

### HasProviderNodeId

`func (o *NodeBulkUpdate) HasProviderNodeId() bool`

HasProviderNodeId returns a boolean if a field has been set.

### SetProviderNodeIdNil

`func (o *NodeBulkUpdate) SetProviderNodeIdNil(b bool)`

 SetProviderNodeIdNil sets the value for ProviderNodeId to be an explicit nil

### UnsetProviderNodeId
`func (o *NodeBulkUpdate) UnsetProviderNodeId()`

UnsetProviderNodeId ensures that no value is present for ProviderNodeId, not even an explicit nil
### GetRack

`func (o *NodeBulkUpdate) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *NodeBulkUpdate) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *NodeBulkUpdate) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *NodeBulkUpdate) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *NodeBulkUpdate) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *NodeBulkUpdate) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetRackUnit

`func (o *NodeBulkUpdate) GetRackUnit() int32`

GetRackUnit returns the RackUnit field if non-nil, zero value otherwise.

### GetRackUnitOk

`func (o *NodeBulkUpdate) GetRackUnitOk() (*int32, bool)`

GetRackUnitOk returns a tuple with the RackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnit

`func (o *NodeBulkUpdate) SetRackUnit(v int32)`

SetRackUnit sets RackUnit field to given value.

### HasRackUnit

`func (o *NodeBulkUpdate) HasRackUnit() bool`

HasRackUnit returns a boolean if a field has been set.

### SetRackUnitNil

`func (o *NodeBulkUpdate) SetRackUnitNil(b bool)`

 SetRackUnitNil sets the value for RackUnit to be an explicit nil

### UnsetRackUnit
`func (o *NodeBulkUpdate) UnsetRackUnit()`

UnsetRackUnit ensures that no value is present for RackUnit, not even an explicit nil
### GetSerialNumber

`func (o *NodeBulkUpdate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NodeBulkUpdate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NodeBulkUpdate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### SetSerialNumberNil

`func (o *NodeBulkUpdate) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *NodeBulkUpdate) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStorage

`func (o *NodeBulkUpdate) GetStorage() []NodeStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *NodeBulkUpdate) GetStorageOk() (*[]NodeStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *NodeBulkUpdate) SetStorage(v []NodeStorage)`

SetStorage sets Storage field to given value.


### GetSystemArch

`func (o *NodeBulkUpdate) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *NodeBulkUpdate) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *NodeBulkUpdate) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetTee

`func (o *NodeBulkUpdate) GetTee() bool`

GetTee returns the Tee field if non-nil, zero value otherwise.

### GetTeeOk

`func (o *NodeBulkUpdate) GetTeeOk() (*bool, bool)`

GetTeeOk returns a tuple with the Tee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTee

`func (o *NodeBulkUpdate) SetTee(v bool)`

SetTee sets Tee field to given value.

### HasTee

`func (o *NodeBulkUpdate) HasTee() bool`

HasTee returns a boolean if a field has been set.

### GetCount

`func (o *NodeBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NodeBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NodeBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *NodeBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *NodeBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *NodeBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *NodeBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *NodeBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *NodeBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


