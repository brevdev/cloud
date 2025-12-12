# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterPipelineCluster**](ClusterPipelineCluster.md) |  | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**InstanceId** | **string** | Unique ID for this instance assigned by its provider | 
**Name** | Pointer to **NullableString** | User-friendly name of this instance | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**State** | Pointer to [**InstanceState**](InstanceState.md) | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | [optional] 
**StateModified** | **time.Time** |  | [readonly] 
**Tags** | Pointer to **[]string** | Tags associated with this instance | [optional] 

## Methods

### NewInstance

`func NewInstance(cluster ClusterPipelineCluster, created time.Time, id string, instanceId string, modified time.Time, stateModified time.Time, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *Instance) GetCluster() ClusterPipelineCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Instance) GetClusterOk() (*ClusterPipelineCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Instance) SetCluster(v ClusterPipelineCluster)`

SetCluster sets Cluster field to given value.


### GetCreated

`func (o *Instance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Instance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Instance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Instance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Instance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetModified

`func (o *Instance) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Instance) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Instance) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetState

`func (o *Instance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *Instance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateModified

`func (o *Instance) GetStateModified() time.Time`

GetStateModified returns the StateModified field if non-nil, zero value otherwise.

### GetStateModifiedOk

`func (o *Instance) GetStateModifiedOk() (*time.Time, bool)`

GetStateModifiedOk returns a tuple with the StateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateModified

`func (o *Instance) SetStateModified(v time.Time)`

SetStateModified sets StateModified field to given value.


### GetTags

`func (o *Instance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Instance) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


