# GPUInventoryV1GpuInventoryValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to [**GPUInventoryV1GpuInventoryValueInstanceType**](GPUInventoryV1GpuInventoryValueInstanceType.md) |  | [optional]
**RegionsWithCapacityAvailable** | Pointer to [**[]GPUInventoryV1GpuInventoryValueRegionsWithCapacityAvailableInner**](GPUInventoryV1GpuInventoryValueRegionsWithCapacityAvailableInner.md) |  | [optional]
**CapacityAvailable** | Pointer to **int32** |  | [optional]

## Methods

### NewGPUInventoryV1GpuInventoryValue

`func NewGPUInventoryV1GpuInventoryValue() *GPUInventoryV1GpuInventoryValue`

NewGPUInventoryV1GpuInventoryValue instantiates a new GPUInventoryV1GpuInventoryValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPUInventoryV1GpuInventoryValueWithDefaults

`func NewGPUInventoryV1GpuInventoryValueWithDefaults() *GPUInventoryV1GpuInventoryValue`

NewGPUInventoryV1GpuInventoryValueWithDefaults instantiates a new GPUInventoryV1GpuInventoryValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *GPUInventoryV1GpuInventoryValue) GetInstanceType() GPUInventoryV1GpuInventoryValueInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GPUInventoryV1GpuInventoryValue) GetInstanceTypeOk() (*GPUInventoryV1GpuInventoryValueInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GPUInventoryV1GpuInventoryValue) SetInstanceType(v GPUInventoryV1GpuInventoryValueInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GPUInventoryV1GpuInventoryValue) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetRegionsWithCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) GetRegionsWithCapacityAvailable() []GPUInventoryV1GpuInventoryValueRegionsWithCapacityAvailableInner`

GetRegionsWithCapacityAvailable returns the RegionsWithCapacityAvailable field if non-nil, zero value otherwise.

### GetRegionsWithCapacityAvailableOk

`func (o *GPUInventoryV1GpuInventoryValue) GetRegionsWithCapacityAvailableOk() (*[]GPUInventoryV1GpuInventoryValueRegionsWithCapacityAvailableInner, bool)`

GetRegionsWithCapacityAvailableOk returns a tuple with the RegionsWithCapacityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionsWithCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) SetRegionsWithCapacityAvailable(v []GPUInventoryV1GpuInventoryValueRegionsWithCapacityAvailableInner)`

SetRegionsWithCapacityAvailable sets RegionsWithCapacityAvailable field to given value.

### HasRegionsWithCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) HasRegionsWithCapacityAvailable() bool`

HasRegionsWithCapacityAvailable returns a boolean if a field has been set.

### GetCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) GetCapacityAvailable() int32`

GetCapacityAvailable returns the CapacityAvailable field if non-nil, zero value otherwise.

### GetCapacityAvailableOk

`func (o *GPUInventoryV1GpuInventoryValue) GetCapacityAvailableOk() (*int32, bool)`

GetCapacityAvailableOk returns a tuple with the CapacityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) SetCapacityAvailable(v int32)`

SetCapacityAvailable sets CapacityAvailable field to given value.

### HasCapacityAvailable

`func (o *GPUInventoryV1GpuInventoryValue) HasCapacityAvailable() bool`

HasCapacityAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


