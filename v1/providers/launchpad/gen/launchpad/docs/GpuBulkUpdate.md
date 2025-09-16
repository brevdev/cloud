# GpuBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**FormFactor** | Pointer to [**InterconnectionTypeEnum**](InterconnectionTypeEnum.md) | GPU form factor  * &#x60;pcie&#x60; - PCIe * &#x60;sxm&#x60; - SXM | [optional] [default to InterconnectionTypePCIe]
**Id** | **string** |  | [readonly] 
**Memory** | Pointer to **int32** | Total GPU memory (in GB) | [optional] [default to 0]
**Model** | **string** | GPU model name | 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**NodeCount** | **int32** |  | [readonly] 
**Priority** | Pointer to **int32** | Weighted preference to use in selecting a GPU for a deployment. A higher priority means the GPU is in higher demand. Lower priority values will be preferred over higher ones during cluster selection. | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewGpuBulkUpdate

`func NewGpuBulkUpdate(created time.Time, id string, model string, modified time.Time, nodeCount int32, count int32, ids []string, result string, ) *GpuBulkUpdate`

NewGpuBulkUpdate instantiates a new GpuBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuBulkUpdateWithDefaults

`func NewGpuBulkUpdateWithDefaults() *GpuBulkUpdate`

NewGpuBulkUpdateWithDefaults instantiates a new GpuBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GpuBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GpuBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GpuBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetFormFactor

`func (o *GpuBulkUpdate) GetFormFactor() InterconnectionTypeEnum`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *GpuBulkUpdate) GetFormFactorOk() (*InterconnectionTypeEnum, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *GpuBulkUpdate) SetFormFactor(v InterconnectionTypeEnum)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *GpuBulkUpdate) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetId

`func (o *GpuBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpuBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpuBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetMemory

`func (o *GpuBulkUpdate) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GpuBulkUpdate) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GpuBulkUpdate) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GpuBulkUpdate) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *GpuBulkUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GpuBulkUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GpuBulkUpdate) SetModel(v string)`

SetModel sets Model field to given value.


### GetModified

`func (o *GpuBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GpuBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GpuBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *GpuBulkUpdate) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *GpuBulkUpdate) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *GpuBulkUpdate) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetPriority

`func (o *GpuBulkUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GpuBulkUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GpuBulkUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GpuBulkUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCount

`func (o *GpuBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GpuBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GpuBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *GpuBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GpuBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GpuBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *GpuBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GpuBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GpuBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


