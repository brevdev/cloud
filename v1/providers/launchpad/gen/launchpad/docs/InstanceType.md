# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **map[string]int32** | Key/value pairs of region name and total available capacity for the instance type | 
**Cloud** | **string** | Cloud provider name | 
**Cpu** | **int32** | Total number of CPUs/vCPUs available | 
**Gpu** | [**InstanceTypeGpu**](InstanceTypeGpu.md) |  | 
**MemoryGb** | **int32** | Total system memory (in GB) | 
**Price** | [**InstanceTypePrice**](InstanceTypePrice.md) |  | 
**Regions** | **[]string** | List of regions names that have available capacity for the instance type | 
**Storage** | [**[]InstanceTypeStorage**](InstanceTypeStorage.md) |  | 
**SystemArch** | [**SystemArchEnum**](SystemArchEnum.md) | CPU architecture  * &#x60;amd64&#x60; - amd64 * &#x60;arm64&#x60; - arm64 | 
**WorkshopId** | **string** | ID of the workshop this instance type is reserved for | 

## Methods

### NewInstanceType

`func NewInstanceType(capacity map[string]int32, cloud string, cpu int32, gpu InstanceTypeGpu, memoryGb int32, price InstanceTypePrice, regions []string, storage []InstanceTypeStorage, systemArch SystemArchEnum, workshopId string, ) *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *InstanceType) GetCapacity() map[string]int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *InstanceType) GetCapacityOk() (*map[string]int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *InstanceType) SetCapacity(v map[string]int32)`

SetCapacity sets Capacity field to given value.


### GetCloud

`func (o *InstanceType) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceType) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceType) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetCpu

`func (o *InstanceType) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceType) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceType) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetGpu

`func (o *InstanceType) GetGpu() InstanceTypeGpu`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *InstanceType) GetGpuOk() (*InstanceTypeGpu, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *InstanceType) SetGpu(v InstanceTypeGpu)`

SetGpu sets Gpu field to given value.


### GetMemoryGb

`func (o *InstanceType) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *InstanceType) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *InstanceType) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.


### GetPrice

`func (o *InstanceType) GetPrice() InstanceTypePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InstanceType) GetPriceOk() (*InstanceTypePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InstanceType) SetPrice(v InstanceTypePrice)`

SetPrice sets Price field to given value.


### GetRegions

`func (o *InstanceType) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *InstanceType) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *InstanceType) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetStorage

`func (o *InstanceType) GetStorage() []InstanceTypeStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceType) GetStorageOk() (*[]InstanceTypeStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceType) SetStorage(v []InstanceTypeStorage)`

SetStorage sets Storage field to given value.


### GetSystemArch

`func (o *InstanceType) GetSystemArch() SystemArchEnum`

GetSystemArch returns the SystemArch field if non-nil, zero value otherwise.

### GetSystemArchOk

`func (o *InstanceType) GetSystemArchOk() (*SystemArchEnum, bool)`

GetSystemArchOk returns a tuple with the SystemArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArch

`func (o *InstanceType) SetSystemArch(v SystemArchEnum)`

SetSystemArch sets SystemArch field to given value.


### GetWorkshopId

`func (o *InstanceType) GetWorkshopId() string`

GetWorkshopId returns the WorkshopId field if non-nil, zero value otherwise.

### GetWorkshopIdOk

`func (o *InstanceType) GetWorkshopIdOk() (*string, bool)`

GetWorkshopIdOk returns a tuple with the WorkshopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkshopId

`func (o *InstanceType) SetWorkshopId(v string)`

SetWorkshopId sets WorkshopId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


