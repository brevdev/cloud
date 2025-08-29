# InstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryInGb** | **int32** | The amount of memory for the instance in gigabytes. Note that this is not VRAM which is determined by GPU type and the number of GPUs. | 
**StorageInGb** | **int32** | The amount of storage for the instance. If this storage is too low for the instance type, please email support@shadeform.ai as the storage may be adjustable. | 
**Vcpus** | **int32** | The number of vCPUs for the instance. | 
**NumGpus** | **int32** | The number of GPUs for the instance. | 
**GpuType** | **string** | The type of GPU for the instance. | 
**Interconnect** | **string** | The type of GPU interconnect. | 
**Nvlink** | Pointer to **bool** | If the instance has NVLink | [optional] 
**VramPerGpuInGb** | **int32** | The video memory per GPU for the instance in gigabytes. | 
**Os** | **string** | The operating system of the instance. | 

## Methods

### NewInstanceConfiguration

`func NewInstanceConfiguration(memoryInGb int32, storageInGb int32, vcpus int32, numGpus int32, gpuType string, interconnect string, vramPerGpuInGb int32, os string, ) *InstanceConfiguration`

NewInstanceConfiguration instantiates a new InstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigurationWithDefaults

`func NewInstanceConfigurationWithDefaults() *InstanceConfiguration`

NewInstanceConfigurationWithDefaults instantiates a new InstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryInGb

`func (o *InstanceConfiguration) GetMemoryInGb() int32`

GetMemoryInGb returns the MemoryInGb field if non-nil, zero value otherwise.

### GetMemoryInGbOk

`func (o *InstanceConfiguration) GetMemoryInGbOk() (*int32, bool)`

GetMemoryInGbOk returns a tuple with the MemoryInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInGb

`func (o *InstanceConfiguration) SetMemoryInGb(v int32)`

SetMemoryInGb sets MemoryInGb field to given value.


### GetStorageInGb

`func (o *InstanceConfiguration) GetStorageInGb() int32`

GetStorageInGb returns the StorageInGb field if non-nil, zero value otherwise.

### GetStorageInGbOk

`func (o *InstanceConfiguration) GetStorageInGbOk() (*int32, bool)`

GetStorageInGbOk returns a tuple with the StorageInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInGb

`func (o *InstanceConfiguration) SetStorageInGb(v int32)`

SetStorageInGb sets StorageInGb field to given value.


### GetVcpus

`func (o *InstanceConfiguration) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *InstanceConfiguration) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *InstanceConfiguration) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.


### GetNumGpus

`func (o *InstanceConfiguration) GetNumGpus() int32`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *InstanceConfiguration) GetNumGpusOk() (*int32, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *InstanceConfiguration) SetNumGpus(v int32)`

SetNumGpus sets NumGpus field to given value.


### GetGpuType

`func (o *InstanceConfiguration) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *InstanceConfiguration) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *InstanceConfiguration) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.


### GetInterconnect

`func (o *InstanceConfiguration) GetInterconnect() string`

GetInterconnect returns the Interconnect field if non-nil, zero value otherwise.

### GetInterconnectOk

`func (o *InstanceConfiguration) GetInterconnectOk() (*string, bool)`

GetInterconnectOk returns a tuple with the Interconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnect

`func (o *InstanceConfiguration) SetInterconnect(v string)`

SetInterconnect sets Interconnect field to given value.


### GetNvlink

`func (o *InstanceConfiguration) GetNvlink() bool`

GetNvlink returns the Nvlink field if non-nil, zero value otherwise.

### GetNvlinkOk

`func (o *InstanceConfiguration) GetNvlinkOk() (*bool, bool)`

GetNvlinkOk returns a tuple with the Nvlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvlink

`func (o *InstanceConfiguration) SetNvlink(v bool)`

SetNvlink sets Nvlink field to given value.

### HasNvlink

`func (o *InstanceConfiguration) HasNvlink() bool`

HasNvlink returns a boolean if a field has been set.

### GetVramPerGpuInGb

`func (o *InstanceConfiguration) GetVramPerGpuInGb() int32`

GetVramPerGpuInGb returns the VramPerGpuInGb field if non-nil, zero value otherwise.

### GetVramPerGpuInGbOk

`func (o *InstanceConfiguration) GetVramPerGpuInGbOk() (*int32, bool)`

GetVramPerGpuInGbOk returns a tuple with the VramPerGpuInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramPerGpuInGb

`func (o *InstanceConfiguration) SetVramPerGpuInGb(v int32)`

SetVramPerGpuInGb sets VramPerGpuInGb field to given value.


### GetOs

`func (o *InstanceConfiguration) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstanceConfiguration) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstanceConfiguration) SetOs(v string)`

SetOs sets Os field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


