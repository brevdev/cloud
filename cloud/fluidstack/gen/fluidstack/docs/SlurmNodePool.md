# SlurmNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the slurm node pool | 
**Name** | **string** | Name of the node pool | 
**Type** | **string** | Instance type of the node pool | 
**Count** | **int32** | Number of nodes in the node pool | 
**Nodes** | [**[]SlurmNode**](SlurmNode.md) | List of nodes in the node pool | 

## Methods

### NewSlurmNodePool

`func NewSlurmNodePool(id string, name string, type_ string, count int32, nodes []SlurmNode, ) *SlurmNodePool`

NewSlurmNodePool instantiates a new SlurmNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmNodePoolWithDefaults

`func NewSlurmNodePoolWithDefaults() *SlurmNodePool`

NewSlurmNodePoolWithDefaults instantiates a new SlurmNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlurmNodePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlurmNodePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlurmNodePool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SlurmNodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmNodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmNodePool) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SlurmNodePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlurmNodePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlurmNodePool) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *SlurmNodePool) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmNodePool) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmNodePool) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNodes

`func (o *SlurmNodePool) GetNodes() []SlurmNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SlurmNodePool) GetNodesOk() (*[]SlurmNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SlurmNodePool) SetNodes(v []SlurmNode)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


