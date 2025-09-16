# InstanceTypeGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Total number of GPUs | 
**Family** | **string** | GPU family name | 
**InterconnectionType** | [**InterconnectionTypeEnum**](InterconnectionTypeEnum.md) | GPU form factor  * &#x60;pcie&#x60; - PCIe * &#x60;sxm&#x60; - SXM | 
**Manufacturer** | **string** | GPU manufacturer name | 
**MemoryGb** | **int32** | Total GPU memory (in GB) | 
**Model** | **string** | GPU model name | 

## Methods

### NewInstanceTypeGpu

`func NewInstanceTypeGpu(count int32, family string, interconnectionType InterconnectionTypeEnum, manufacturer string, memoryGb int32, model string, ) *InstanceTypeGpu`

NewInstanceTypeGpu instantiates a new InstanceTypeGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeGpuWithDefaults

`func NewInstanceTypeGpuWithDefaults() *InstanceTypeGpu`

NewInstanceTypeGpuWithDefaults instantiates a new InstanceTypeGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InstanceTypeGpu) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InstanceTypeGpu) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InstanceTypeGpu) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFamily

`func (o *InstanceTypeGpu) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *InstanceTypeGpu) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *InstanceTypeGpu) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetInterconnectionType

`func (o *InstanceTypeGpu) GetInterconnectionType() InterconnectionTypeEnum`

GetInterconnectionType returns the InterconnectionType field if non-nil, zero value otherwise.

### GetInterconnectionTypeOk

`func (o *InstanceTypeGpu) GetInterconnectionTypeOk() (*InterconnectionTypeEnum, bool)`

GetInterconnectionTypeOk returns a tuple with the InterconnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectionType

`func (o *InstanceTypeGpu) SetInterconnectionType(v InterconnectionTypeEnum)`

SetInterconnectionType sets InterconnectionType field to given value.


### GetManufacturer

`func (o *InstanceTypeGpu) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InstanceTypeGpu) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InstanceTypeGpu) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetMemoryGb

`func (o *InstanceTypeGpu) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *InstanceTypeGpu) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *InstanceTypeGpu) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.


### GetModel

`func (o *InstanceTypeGpu) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InstanceTypeGpu) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InstanceTypeGpu) SetModel(v string)`

SetModel sets Model field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


