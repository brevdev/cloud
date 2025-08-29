# InstanceTypeConfiguration

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
**OsOptions** | **[]string** | The list of available operating systems for the instance. | 

## Methods

### NewInstanceTypeConfiguration

`func NewInstanceTypeConfiguration(memoryInGb int32, storageInGb int32, vcpus int32, numGpus int32, gpuType string, interconnect string, vramPerGpuInGb int32, osOptions []string, ) *InstanceTypeConfiguration`

NewInstanceTypeConfiguration instantiates a new InstanceTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeConfigurationWithDefaults

`func NewInstanceTypeConfigurationWithDefaults() *InstanceTypeConfiguration`

NewInstanceTypeConfigurationWithDefaults instantiates a new InstanceTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryInGb

`func (o *InstanceTypeConfiguration) GetMemoryInGb() int32`

GetMemoryInGb returns the MemoryInGb field if non-nil, zero value otherwise.

### GetMemoryInGbOk

`func (o *InstanceTypeConfiguration) GetMemoryInGbOk() (*int32, bool)`

GetMemoryInGbOk returns a tuple with the MemoryInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInGb

`func (o *InstanceTypeConfiguration) SetMemoryInGb(v int32)`

SetMemoryInGb sets MemoryInGb field to given value.


### GetStorageInGb

`func (o *InstanceTypeConfiguration) GetStorageInGb() int32`

GetStorageInGb returns the StorageInGb field if non-nil, zero value otherwise.

### GetStorageInGbOk

`func (o *InstanceTypeConfiguration) GetStorageInGbOk() (*int32, bool)`

GetStorageInGbOk returns a tuple with the StorageInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInGb

`func (o *InstanceTypeConfiguration) SetStorageInGb(v int32)`

SetStorageInGb sets StorageInGb field to given value.


### GetVcpus

`func (o *InstanceTypeConfiguration) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *InstanceTypeConfiguration) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *InstanceTypeConfiguration) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.


### GetNumGpus

`func (o *InstanceTypeConfiguration) GetNumGpus() int32`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *InstanceTypeConfiguration) GetNumGpusOk() (*int32, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *InstanceTypeConfiguration) SetNumGpus(v int32)`

SetNumGpus sets NumGpus field to given value.


### GetGpuType

`func (o *InstanceTypeConfiguration) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *InstanceTypeConfiguration) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *InstanceTypeConfiguration) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.


### GetInterconnect

`func (o *InstanceTypeConfiguration) GetInterconnect() string`

GetInterconnect returns the Interconnect field if non-nil, zero value otherwise.

### GetInterconnectOk

`func (o *InstanceTypeConfiguration) GetInterconnectOk() (*string, bool)`

GetInterconnectOk returns a tuple with the Interconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnect

`func (o *InstanceTypeConfiguration) SetInterconnect(v string)`

SetInterconnect sets Interconnect field to given value.


### GetNvlink

`func (o *InstanceTypeConfiguration) GetNvlink() bool`

GetNvlink returns the Nvlink field if non-nil, zero value otherwise.

### GetNvlinkOk

`func (o *InstanceTypeConfiguration) GetNvlinkOk() (*bool, bool)`

GetNvlinkOk returns a tuple with the Nvlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvlink

`func (o *InstanceTypeConfiguration) SetNvlink(v bool)`

SetNvlink sets Nvlink field to given value.

### HasNvlink

`func (o *InstanceTypeConfiguration) HasNvlink() bool`

HasNvlink returns a boolean if a field has been set.

### GetVramPerGpuInGb

`func (o *InstanceTypeConfiguration) GetVramPerGpuInGb() int32`

GetVramPerGpuInGb returns the VramPerGpuInGb field if non-nil, zero value otherwise.

### GetVramPerGpuInGbOk

`func (o *InstanceTypeConfiguration) GetVramPerGpuInGbOk() (*int32, bool)`

GetVramPerGpuInGbOk returns a tuple with the VramPerGpuInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramPerGpuInGb

`func (o *InstanceTypeConfiguration) SetVramPerGpuInGb(v int32)`

SetVramPerGpuInGb sets VramPerGpuInGb field to given value.


### GetOsOptions

`func (o *InstanceTypeConfiguration) GetOsOptions() []string`

GetOsOptions returns the OsOptions field if non-nil, zero value otherwise.

### GetOsOptionsOk

`func (o *InstanceTypeConfiguration) GetOsOptionsOk() (*[]string, bool)`

GetOsOptionsOk returns a tuple with the OsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsOptions

`func (o *InstanceTypeConfiguration) SetOsOptions(v []string)`

SetOsOptions sets OsOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


