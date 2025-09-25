# ClusterNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BmcIp** | Pointer to **string** | IP address of the BMC | [optional] 
**BmcMac** | Pointer to **string** | MAC address of the BMC | [optional] 
**BmcPassword** | Pointer to **string** | Password for the BMC | [optional] 
**BmcUser** | Pointer to **string** | Username for the BMC | [optional] 
**Cluster** | Pointer to **string** | UUID of the node&#39;s parent cluster | [optional] 
**Cpu** | Pointer to **int32** | Number of CPU cores installed | [optional] 
**CpuManufacturer** | Pointer to [**CpuManufacturerEnum**](CpuManufacturerEnum.md) |  | [optional] 
**CpuModel** | Pointer to **string** | Model information for the node&#39;s CPU | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**GarageId** | Pointer to **string** | ID for the garage where a node is parked when unprovisioned | [optional] 
**Gpu** | [**ClusterGpusInner**](ClusterGpusInner.md) |  | 
**GpuAlias** | **string** | Alias for GPU plan (i.e. installed GPU type and count) | [readonly] 
**GpuCount** | Pointer to **int32** | Number of GPUs installed | [optional] 
**GpuModel** | **string** | Model of GPU(s) installed | [readonly] 
**GpuVbios** | Pointer to **string** | VBIOS version used by installed GPU(s) | [optional] 
**Id** | **string** |  | [readonly] 
**Location** | [**NodeLocation**](NodeLocation.md) |  | 
**Memory** | Pointer to **int32** | Amount of RAM installed (in GB) | [optional] 
**MgmtIp** | Pointer to **string** | Management IP address | [optional] 
**MgmtMac** | Pointer to **string** | Management MAC address | [optional] 
**Model** | Pointer to **string** | Hardware model of the node | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NetworkType** | Pointer to [**NetworkTypeEnum**](NetworkTypeEnum.md) | Type of networking technology used  * &#x60;ethernet&#x60; - Ethernet * &#x60;infiniband&#x60; - InfiniBand | [optional] 
**NicPrefixes** | Pointer to **[]string** | Prefixes for the node&#39;s network interface(s) | [optional] 
**Notes** | Pointer to **string** | Administrative comments about the node | [optional] 
**Oem** | Pointer to **string** | UUID of the node&#39;s OEM | [optional] 
**ProviderNodeId** | Pointer to **string** | Unique ID for this node assigned by its provider | [optional] 
**Rack** | Pointer to **string** | Physical rack identifier | [optional] 
**RackUnit** | Pointer to **int32** | The rack unit (RU) within the rack where the node is installed | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the node | [optional] 
**Storage** | [**[]NodeStorage**](NodeStorage.md) |  | 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**Tee** | Pointer to **bool** | Does the node support Trusted Execution Environment (TEE)? | [optional] 

## Methods

### NewClusterNodesInner

`func NewClusterNodesInner(created time.Time, gpu ClusterGpusInner, gpuAlias string, gpuModel string, id string, location NodeLocation, modified time.Time, storage []NodeStorage, systemArch SystemArchEnum, ) *ClusterNodesInner`

NewClusterNodesInner instantiates a new ClusterNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodesInnerWithDefaults

`func NewClusterNodesInnerWithDefaults() *ClusterNodesInner`

NewClusterNodesInnerWithDefaults instantiates a new ClusterNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmcIp

`func (o *ClusterNodesInner) GetBmcIp() string`

GetBmcIp returns the BmcIp field if non-nil, zero value otherwise.

### GetBmcIpOk

`func (o *ClusterNodesInner) GetBmcIpOk() (*string, bool)`

GetBmcIpOk returns a tuple with the BmcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcIp

`func (o *ClusterNodesInner) SetBmcIp(v string)`

SetBmcIp sets BmcIp field to given value.

### HasBmcIp

`func (o *ClusterNodesInner) HasBmcIp() bool`

HasBmcIp returns a boolean if a field has been set.

### GetBmcMac

`func (o *ClusterNodesInner) GetBmcMac() string`

GetBmcMac returns the BmcMac field if non-nil, zero value otherwise.

### GetBmcMacOk

`func (o *ClusterNodesInner) GetBmcMacOk() (*string, bool)`

GetBmcMacOk returns a tuple with the BmcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMac

`func (o *ClusterNodesInner) SetBmcMac(v string)`

SetBmcMac sets BmcMac field to given value.

### HasBmcMac

`func (o *ClusterNodesInner) HasBmcMac() bool`

HasBmcMac returns a boolean if a field has been set.

### GetBmcPassword

`func (o *ClusterNodesInner) GetBmcPassword() string`

GetBmcPassword returns the BmcPassword field if non-nil, zero value otherwise.

### GetBmcPasswordOk

`func (o *ClusterNodesInner) GetBmcPasswordOk() (*string, bool)`

GetBmcPasswordOk returns a tuple with the BmcPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcPassword

`func (o *ClusterNodesInner) SetBmcPassword(v string)`

SetBmcPassword sets BmcPassword field to given value.

### HasBmcPassword

`func (o *ClusterNodesInner) HasBmcPassword() bool`

HasBmcPassword returns a boolean if a field has been set.

### GetBmcUser

`func (o *ClusterNodesInner) GetBmcUser() string`

GetBmcUser returns the BmcUser field if non-nil, zero value otherwise.

### GetBmcUserOk

`func (o *ClusterNodesInner) GetBmcUserOk() (*string, bool)`

GetBmcUserOk returns a tuple with the BmcUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcUser

`func (o *ClusterNodesInner) SetBmcUser(v string)`

SetBmcUser sets BmcUser field to given value.

### HasBmcUser

`func (o *ClusterNodesInner) HasBmcUser() bool`

HasBmcUser returns a boolean if a field has been set.

### GetCluster

`func (o *ClusterNodesInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterNodesInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterNodesInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterNodesInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCpu

`func (o *ClusterNodesInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterNodesInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterNodesInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ClusterNodesInner) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpuManufacturer

`func (o *ClusterNodesInner) GetCpuManufacturer() CpuManufacturerEnum`

GetCpuManufacturer returns the CpuManufacturer field if non-nil, zero value otherwise.

### GetCpuManufacturerOk

`func (o *ClusterNodesInner) GetCpuManufacturerOk() (*CpuManufacturerEnum, bool)`

GetCpuManufacturerOk returns a tuple with the CpuManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuManufacturer

`func (o *ClusterNodesInner) SetCpuManufacturer(v CpuManufacturerEnum)`

SetCpuManufacturer sets CpuManufacturer field to given value.

### HasCpuManufacturer

`func (o *ClusterNodesInner) HasCpuManufacturer() bool`

HasCpuManufacturer returns a boolean if a field has been set.

### GetCpuModel

`func (o *ClusterNodesInner) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ClusterNodesInner) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ClusterNodesInner) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ClusterNodesInner) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCreated

`func (o *ClusterNodesInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterNodesInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterNodesInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetGarageId

`func (o *ClusterNodesInner) GetGarageId() string`

GetGarageId returns the GarageId field if non-nil, zero value otherwise.

### GetGarageIdOk

`func (o *ClusterNodesInner) GetGarageIdOk() (*string, bool)`

GetGarageIdOk returns a tuple with the GarageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarageId

`func (o *ClusterNodesInner) SetGarageId(v string)`

SetGarageId sets GarageId field to given value.

### HasGarageId

`func (o *ClusterNodesInner) HasGarageId() bool`

HasGarageId returns a boolean if a field has been set.

### GetGpu

`func (o *ClusterNodesInner) GetGpu() ClusterGpusInner`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *ClusterNodesInner) GetGpuOk() (*ClusterGpusInner, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *ClusterNodesInner) SetGpu(v ClusterGpusInner)`

SetGpu sets Gpu field to given value.


### GetGpuAlias

`func (o *ClusterNodesInner) GetGpuAlias() string`

GetGpuAlias returns the GpuAlias field if non-nil, zero value otherwise.

### GetGpuAliasOk

`func (o *ClusterNodesInner) GetGpuAliasOk() (*string, bool)`

GetGpuAliasOk returns a tuple with the GpuAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuAlias

`func (o *ClusterNodesInner) SetGpuAlias(v string)`

SetGpuAlias sets GpuAlias field to given value.


### GetGpuCount

`func (o *ClusterNodesInner) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ClusterNodesInner) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ClusterNodesInner) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ClusterNodesInner) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuModel

`func (o *ClusterNodesInner) GetGpuModel() string`

GetGpuModel returns the GpuModel field if non-nil, zero value otherwise.

### GetGpuModelOk

`func (o *ClusterNodesInner) GetGpuModelOk() (*string, bool)`

GetGpuModelOk returns a tuple with the GpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuModel

`func (o *ClusterNodesInner) SetGpuModel(v string)`

SetGpuModel sets GpuModel field to given value.


### GetGpuVbios

`func (o *ClusterNodesInner) GetGpuVbios() string`

GetGpuVbios returns the GpuVbios field if non-nil, zero value otherwise.

### GetGpuVbiosOk

`func (o *ClusterNodesInner) GetGpuVbiosOk() (*string, bool)`

GetGpuVbiosOk returns a tuple with the GpuVbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuVbios

`func (o *ClusterNodesInner) SetGpuVbios(v string)`

SetGpuVbios sets GpuVbios field to given value.

### HasGpuVbios

`func (o *ClusterNodesInner) HasGpuVbios() bool`

HasGpuVbios returns a boolean if a field has been set.

### GetId

`func (o *ClusterNodesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterNodesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterNodesInner) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *ClusterNodesInner) GetLocation() NodeLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterNodesInner) GetLocationOk() (*NodeLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterNodesInner) SetLocation(v NodeLocation)`

SetLocation sets Location field to given value.


### GetMemory

`func (o *ClusterNodesInner) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterNodesInner) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterNodesInner) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterNodesInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMgmtIp

`func (o *ClusterNodesInner) GetMgmtIp() string`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *ClusterNodesInner) GetMgmtIpOk() (*string, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *ClusterNodesInner) SetMgmtIp(v string)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *ClusterNodesInner) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### GetMgmtMac

`func (o *ClusterNodesInner) GetMgmtMac() string`

GetMgmtMac returns the MgmtMac field if non-nil, zero value otherwise.

### GetMgmtMacOk

`func (o *ClusterNodesInner) GetMgmtMacOk() (*string, bool)`

GetMgmtMacOk returns a tuple with the MgmtMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMac

`func (o *ClusterNodesInner) SetMgmtMac(v string)`

SetMgmtMac sets MgmtMac field to given value.

### HasMgmtMac

`func (o *ClusterNodesInner) HasMgmtMac() bool`

HasMgmtMac returns a boolean if a field has been set.

### GetModel

`func (o *ClusterNodesInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClusterNodesInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClusterNodesInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClusterNodesInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModified

`func (o *ClusterNodesInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterNodesInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterNodesInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNetworkType

`func (o *ClusterNodesInner) GetNetworkType() NetworkTypeEnum`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ClusterNodesInner) GetNetworkTypeOk() (*NetworkTypeEnum, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ClusterNodesInner) SetNetworkType(v NetworkTypeEnum)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *ClusterNodesInner) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetNicPrefixes

`func (o *ClusterNodesInner) GetNicPrefixes() []string`

GetNicPrefixes returns the NicPrefixes field if non-nil, zero value otherwise.

### GetNicPrefixesOk

`func (o *ClusterNodesInner) GetNicPrefixesOk() (*[]string, bool)`

GetNicPrefixesOk returns a tuple with the NicPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicPrefixes

`func (o *ClusterNodesInner) SetNicPrefixes(v []string)`

SetNicPrefixes sets NicPrefixes field to given value.

### HasNicPrefixes

`func (o *ClusterNodesInner) HasNicPrefixes() bool`

HasNicPrefixes returns a boolean if a field has been set.

### GetNotes

`func (o *ClusterNodesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ClusterNodesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ClusterNodesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ClusterNodesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOem

`func (o *ClusterNodesInner) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *ClusterNodesInner) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *ClusterNodesInner) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *ClusterNodesInner) HasOem() bool`

HasOem returns a boolean if a field has been set.

### GetProviderNodeId

`func (o *ClusterNodesInner) GetProviderNodeId() string`

GetProviderNodeId returns the ProviderNodeId field if non-nil, zero value otherwise.

### GetProviderNodeIdOk

`func (o *ClusterNodesInner) GetProviderNodeIdOk() (*string, bool)`

GetProviderNodeIdOk returns a tuple with the ProviderNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNodeId

`func (o *ClusterNodesInner) SetProviderNodeId(v string)`

SetProviderNodeId sets ProviderNodeId field to given value.

### HasProviderNodeId

`func (o *ClusterNodesInner) HasProviderNodeId() bool`

HasProviderNodeId returns a boolean if a field has been set.

### GetRack

`func (o *ClusterNodesInner) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *ClusterNodesInner) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *ClusterNodesInner) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *ClusterNodesInner) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetRackUnit

`func (o *ClusterNodesInner) GetRackUnit() int32`

GetRackUnit returns the RackUnit field if non-nil, zero value otherwise.

### GetRackUnitOk

`func (o *ClusterNodesInner) GetRackUnitOk() (*int32, bool)`

GetRackUnitOk returns a tuple with the RackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnit

`func (o *ClusterNodesInner) SetRackUnit(v int32)`

SetRackUnit sets RackUnit field to given value.

### HasRackUnit

`func (o *ClusterNodesInner) HasRackUnit() bool`

HasRackUnit returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ClusterNodesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ClusterNodesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ClusterNodesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ClusterNodesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStorage

`func (o *ClusterNodesInner) GetStorage() []NodeStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ClusterNodesInner) GetStorageOk() (*[]NodeStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ClusterNodesInner) SetStorage(v []NodeStorage)`

SetStorage sets Storage field to given value.


### GetSystemArch

`func (o *ClusterNodesInner) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *ClusterNodesInner) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *ClusterNodesInner) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetTee

`func (o *ClusterNodesInner) GetTee() bool`

GetTee returns the Tee field if non-nil, zero value otherwise.

### GetTeeOk

`func (o *ClusterNodesInner) GetTeeOk() (*bool, bool)`

GetTeeOk returns a tuple with the Tee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTee

`func (o *ClusterNodesInner) SetTee(v bool)`

SetTee sets Tee field to given value.

### HasTee

`func (o *ClusterNodesInner) HasTee() bool`

HasTee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


