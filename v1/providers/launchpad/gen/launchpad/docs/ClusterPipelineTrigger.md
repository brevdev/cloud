# ClusterPipelineTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PipelineAction**](PipelineAction.md) | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy * &#x60;notify&#x60; - notify | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**PipelineId** | **int32** | GitLab pipeline ID | [readonly] 
**RequestId** | **string** | Request ID that the pipeline was triggered for | [readonly] 
**Url** | **string** | URL for the pipeline details | [readonly] 

## Methods

### NewClusterPipelineTrigger

`func NewClusterPipelineTrigger(action PipelineAction, created time.Time, id string, modified time.Time, pipelineId int32, requestId string, url string, ) *ClusterPipelineTrigger`

NewClusterPipelineTrigger instantiates a new ClusterPipelineTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPipelineTriggerWithDefaults

`func NewClusterPipelineTriggerWithDefaults() *ClusterPipelineTrigger`

NewClusterPipelineTriggerWithDefaults instantiates a new ClusterPipelineTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ClusterPipelineTrigger) GetAction() PipelineAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterPipelineTrigger) GetActionOk() (*PipelineAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterPipelineTrigger) SetAction(v PipelineAction)`

SetAction sets Action field to given value.


### GetCreated

`func (o *ClusterPipelineTrigger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterPipelineTrigger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterPipelineTrigger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *ClusterPipelineTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPipelineTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPipelineTrigger) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *ClusterPipelineTrigger) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterPipelineTrigger) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterPipelineTrigger) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPipelineId

`func (o *ClusterPipelineTrigger) GetPipelineId() int32`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *ClusterPipelineTrigger) GetPipelineIdOk() (*int32, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *ClusterPipelineTrigger) SetPipelineId(v int32)`

SetPipelineId sets PipelineId field to given value.


### GetRequestId

`func (o *ClusterPipelineTrigger) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClusterPipelineTrigger) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClusterPipelineTrigger) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUrl

`func (o *ClusterPipelineTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterPipelineTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterPipelineTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


