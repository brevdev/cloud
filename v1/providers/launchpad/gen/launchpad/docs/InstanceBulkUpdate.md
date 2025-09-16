# InstanceBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**DeploymentCluster**](DeploymentCluster.md) |  | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**InstanceId** | **string** | Unique ID for this instance assigned by its provider | [readonly] 
**Name** | Pointer to **NullableString** | User-friendly name of this instance | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**State** | Pointer to [**InstanceState**](InstanceState.md) | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | [optional] 
**StateModified** | **time.Time** |  | [readonly] 
**Tags** | Pointer to **[]string** | Tags associated with this instance | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewInstanceBulkUpdate

`func NewInstanceBulkUpdate(cluster DeploymentCluster, created time.Time, id string, instanceId string, modified time.Time, stateModified time.Time, count int32, ids []string, result string, ) *InstanceBulkUpdate`

NewInstanceBulkUpdate instantiates a new InstanceBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceBulkUpdateWithDefaults

`func NewInstanceBulkUpdateWithDefaults() *InstanceBulkUpdate`

NewInstanceBulkUpdateWithDefaults instantiates a new InstanceBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *InstanceBulkUpdate) GetCluster() DeploymentCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InstanceBulkUpdate) GetClusterOk() (*DeploymentCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InstanceBulkUpdate) SetCluster(v DeploymentCluster)`

SetCluster sets Cluster field to given value.


### GetCreated

`func (o *InstanceBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InstanceBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InstanceBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *InstanceBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *InstanceBulkUpdate) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceBulkUpdate) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceBulkUpdate) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *InstanceBulkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceBulkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceBulkUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceBulkUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceBulkUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceBulkUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetModified

`func (o *InstanceBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InstanceBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InstanceBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetState

`func (o *InstanceBulkUpdate) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceBulkUpdate) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceBulkUpdate) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *InstanceBulkUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateModified

`func (o *InstanceBulkUpdate) GetStateModified() time.Time`

GetStateModified returns the StateModified field if non-nil, zero value otherwise.

### GetStateModifiedOk

`func (o *InstanceBulkUpdate) GetStateModifiedOk() (*time.Time, bool)`

GetStateModifiedOk returns a tuple with the StateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateModified

`func (o *InstanceBulkUpdate) SetStateModified(v time.Time)`

SetStateModified sets StateModified field to given value.


### GetTags

`func (o *InstanceBulkUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceBulkUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceBulkUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceBulkUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCount

`func (o *InstanceBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InstanceBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InstanceBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *InstanceBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InstanceBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InstanceBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *InstanceBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InstanceBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InstanceBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


