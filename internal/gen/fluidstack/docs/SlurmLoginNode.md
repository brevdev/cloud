# SlurmLoginNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the node | 
**PublicIP** | Pointer to **string** | Public IP address of the node | [optional] 

## Methods

### NewSlurmLoginNode

`func NewSlurmLoginNode(name string, ) *SlurmLoginNode`

NewSlurmLoginNode instantiates a new SlurmLoginNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmLoginNodeWithDefaults

`func NewSlurmLoginNodeWithDefaults() *SlurmLoginNode`

NewSlurmLoginNodeWithDefaults instantiates a new SlurmLoginNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlurmLoginNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmLoginNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmLoginNode) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIP

`func (o *SlurmLoginNode) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *SlurmLoginNode) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *SlurmLoginNode) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *SlurmLoginNode) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


