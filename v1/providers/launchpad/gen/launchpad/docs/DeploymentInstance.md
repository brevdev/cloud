# DeploymentInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**InstanceId** | **string** | Unique ID for this instance assigned by its provider | [readonly] 
**Name** | **NullableString** | User-friendly name of this instance | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**State** | [**InstanceState**](InstanceState.md) | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | [readonly] 
**StateModified** | **time.Time** |  | [readonly] 
**Tags** | **[]string** | Tags associated with this instance | 

## Methods

### NewDeploymentInstance

`func NewDeploymentInstance(created time.Time, id string, instanceId string, name NullableString, modified time.Time, state InstanceState, stateModified time.Time, tags []string, ) *DeploymentInstance`

NewDeploymentInstance instantiates a new DeploymentInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentInstanceWithDefaults

`func NewDeploymentInstanceWithDefaults() *DeploymentInstance`

NewDeploymentInstanceWithDefaults instantiates a new DeploymentInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeploymentInstance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentInstance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentInstance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *DeploymentInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentInstance) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *DeploymentInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DeploymentInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DeploymentInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *DeploymentInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentInstance) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeploymentInstance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentInstance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetModified

`func (o *DeploymentInstance) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentInstance) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentInstance) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetState

`func (o *DeploymentInstance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentInstance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentInstance) SetState(v InstanceState)`

SetState sets State field to given value.


### GetStateModified

`func (o *DeploymentInstance) GetStateModified() time.Time`

GetStateModified returns the StateModified field if non-nil, zero value otherwise.

### GetStateModifiedOk

`func (o *DeploymentInstance) GetStateModifiedOk() (*time.Time, bool)`

GetStateModifiedOk returns a tuple with the StateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateModified

`func (o *DeploymentInstance) SetStateModified(v time.Time)`

SetStateModified sets StateModified field to given value.


### GetTags

`func (o *DeploymentInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeploymentInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeploymentInstance) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


