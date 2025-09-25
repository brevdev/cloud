# Gpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**FormFactor** | Pointer to [**InterconnectionTypeEnum**](InterconnectionTypeEnum.md) | GPU form factor  * &#x60;pcie&#x60; - PCIe * &#x60;sxm&#x60; - SXM | [optional] [default to InterconnectionTypePCIe]
**Id** | **string** |  | [readonly] 
**Memory** | Pointer to **int32** | Total GPU memory (in GB) | [optional] [default to 0]
**Model** | **string** | GPU model name | 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | Pointer to **int32** |  | [optional] [readonly] 
**Priority** | Pointer to **int32** | Weighted preference to use in selecting a GPU for a deployment. A higher priority means the GPU is in higher demand. Lower priority values will be preferred over higher ones during cluster selection. | [optional] 

## Methods

### NewGpu

`func NewGpu(created time.Time, id string, model string, modified time.Time, ) *Gpu`

NewGpu instantiates a new Gpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuWithDefaults

`func NewGpuWithDefaults() *Gpu`

NewGpuWithDefaults instantiates a new Gpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Gpu) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Gpu) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Gpu) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetFormFactor

`func (o *Gpu) GetFormFactor() InterconnectionTypeEnum`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *Gpu) GetFormFactorOk() (*InterconnectionTypeEnum, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *Gpu) SetFormFactor(v InterconnectionTypeEnum)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *Gpu) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetId

`func (o *Gpu) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gpu) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gpu) SetId(v string)`

SetId sets Id field to given value.


### GetMemory

`func (o *Gpu) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Gpu) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Gpu) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Gpu) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *Gpu) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Gpu) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Gpu) SetModel(v string)`

SetModel sets Model field to given value.


### GetModified

`func (o *Gpu) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Gpu) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Gpu) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *Gpu) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Gpu) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Gpu) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Gpu) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetPriority

`func (o *Gpu) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Gpu) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Gpu) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Gpu) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


