# SlurmNodePoolPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the node pool | 
**Type** | **string** | Instance type of the node pool | 
**Count** | **int32** | Number of nodes in the node pool | 

## Methods

### NewSlurmNodePoolPostRequest

`func NewSlurmNodePoolPostRequest(name string, type_ string, count int32, ) *SlurmNodePoolPostRequest`

NewSlurmNodePoolPostRequest instantiates a new SlurmNodePoolPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmNodePoolPostRequestWithDefaults

`func NewSlurmNodePoolPostRequestWithDefaults() *SlurmNodePoolPostRequest`

NewSlurmNodePoolPostRequestWithDefaults instantiates a new SlurmNodePoolPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlurmNodePoolPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmNodePoolPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmNodePoolPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SlurmNodePoolPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlurmNodePoolPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlurmNodePoolPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *SlurmNodePoolPostRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmNodePoolPostRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmNodePoolPostRequest) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


