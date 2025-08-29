# SlurmNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the node | 
**InstanceID** | **string** | Unique identifier of the instance backing the node | 
**States** | **[]string** | States of the node | 
**SlurmdVersion** | **string** | Version of slurmd running on the node | 

## Methods

### NewSlurmNode

`func NewSlurmNode(name string, instanceID string, states []string, slurmdVersion string, ) *SlurmNode`

NewSlurmNode instantiates a new SlurmNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmNodeWithDefaults

`func NewSlurmNodeWithDefaults() *SlurmNode`

NewSlurmNodeWithDefaults instantiates a new SlurmNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlurmNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmNode) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceID

`func (o *SlurmNode) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *SlurmNode) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *SlurmNode) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.


### GetStates

`func (o *SlurmNode) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *SlurmNode) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *SlurmNode) SetStates(v []string)`

SetStates sets States field to given value.


### GetSlurmdVersion

`func (o *SlurmNode) GetSlurmdVersion() string`

GetSlurmdVersion returns the SlurmdVersion field if non-nil, zero value otherwise.

### GetSlurmdVersionOk

`func (o *SlurmNode) GetSlurmdVersionOk() (*string, bool)`

GetSlurmdVersionOk returns a tuple with the SlurmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurmdVersion

`func (o *SlurmNode) SetSlurmdVersion(v string)`

SetSlurmdVersion sets SlurmdVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


