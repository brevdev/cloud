# ClusterGpusInner

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

### NewClusterGpusInner

`func NewClusterGpusInner(created time.Time, id string, model string, modified time.Time, ) *ClusterGpusInner`

NewClusterGpusInner instantiates a new ClusterGpusInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterGpusInnerWithDefaults

`func NewClusterGpusInnerWithDefaults() *ClusterGpusInner`

NewClusterGpusInnerWithDefaults instantiates a new ClusterGpusInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ClusterGpusInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterGpusInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterGpusInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetFormFactor

`func (o *ClusterGpusInner) GetFormFactor() InterconnectionTypeEnum`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *ClusterGpusInner) GetFormFactorOk() (*InterconnectionTypeEnum, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *ClusterGpusInner) SetFormFactor(v InterconnectionTypeEnum)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *ClusterGpusInner) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetId

`func (o *ClusterGpusInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterGpusInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterGpusInner) SetId(v string)`

SetId sets Id field to given value.


### GetMemory

`func (o *ClusterGpusInner) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterGpusInner) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterGpusInner) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterGpusInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetModel

`func (o *ClusterGpusInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClusterGpusInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClusterGpusInner) SetModel(v string)`

SetModel sets Model field to given value.


### GetModified

`func (o *ClusterGpusInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterGpusInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterGpusInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetNodeCount

`func (o *ClusterGpusInner) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterGpusInner) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterGpusInner) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterGpusInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetPriority

`func (o *ClusterGpusInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ClusterGpusInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ClusterGpusInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ClusterGpusInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


