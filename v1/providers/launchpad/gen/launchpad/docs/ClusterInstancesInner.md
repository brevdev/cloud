# ClusterInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**DeploymentCluster**](DeploymentCluster.md) |  | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**InstanceId** | **string** | Unique ID for this instance assigned by its provider | 
**Name** | Pointer to **string** | User-friendly name of this instance | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**State** | Pointer to [**InstanceState**](InstanceState.md) | Current lifecycle state of this instance  * &#x60;running&#x60; - Instance is running * &#x60;starting&#x60; - Instance is starting * &#x60;stopped&#x60; - Instance is stopped * &#x60;stopping&#x60; - Instance is stopping * &#x60;unknown&#x60; - Instance state is currently unknown | [optional] 
**StateModified** | **time.Time** |  | [readonly] 
**Tags** | Pointer to **[]string** | Tags associated with this instance | [optional] 

## Methods

### NewClusterInstancesInner

`func NewClusterInstancesInner(cluster DeploymentCluster, created time.Time, id string, instanceId string, modified time.Time, stateModified time.Time, ) *ClusterInstancesInner`

NewClusterInstancesInner instantiates a new ClusterInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstancesInnerWithDefaults

`func NewClusterInstancesInnerWithDefaults() *ClusterInstancesInner`

NewClusterInstancesInnerWithDefaults instantiates a new ClusterInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ClusterInstancesInner) GetCluster() DeploymentCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterInstancesInner) GetClusterOk() (*DeploymentCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterInstancesInner) SetCluster(v DeploymentCluster)`

SetCluster sets Cluster field to given value.


### GetCreated

`func (o *ClusterInstancesInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterInstancesInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterInstancesInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *ClusterInstancesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterInstancesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterInstancesInner) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *ClusterInstancesInner) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClusterInstancesInner) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClusterInstancesInner) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *ClusterInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModified

`func (o *ClusterInstancesInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterInstancesInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterInstancesInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetState

`func (o *ClusterInstancesInner) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterInstancesInner) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterInstancesInner) SetState(v InstanceState)`

SetState sets State field to given value.

### HasState

`func (o *ClusterInstancesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateModified

`func (o *ClusterInstancesInner) GetStateModified() time.Time`

GetStateModified returns the StateModified field if non-nil, zero value otherwise.

### GetStateModifiedOk

`func (o *ClusterInstancesInner) GetStateModifiedOk() (*time.Time, bool)`

GetStateModifiedOk returns a tuple with the StateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateModified

`func (o *ClusterInstancesInner) SetStateModified(v time.Time)`

SetStateModified sets StateModified field to given value.


### GetTags

`func (o *ClusterInstancesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterInstancesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterInstancesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterInstancesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


