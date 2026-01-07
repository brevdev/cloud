# DeploymentPipelineTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PipelineAction**](PipelineAction.md) | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy * &#x60;notify&#x60; - notify | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**PipelineId** | **int32** | GitLab pipeline ID | [readonly] 
**Url** | **string** | URL for the pipeline details | [readonly] 

## Methods

### NewDeploymentPipelineTrigger

`func NewDeploymentPipelineTrigger(action PipelineAction, created time.Time, id string, modified time.Time, pipelineId int32, url string, ) *DeploymentPipelineTrigger`

NewDeploymentPipelineTrigger instantiates a new DeploymentPipelineTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPipelineTriggerWithDefaults

`func NewDeploymentPipelineTriggerWithDefaults() *DeploymentPipelineTrigger`

NewDeploymentPipelineTriggerWithDefaults instantiates a new DeploymentPipelineTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentPipelineTrigger) GetAction() PipelineAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentPipelineTrigger) GetActionOk() (*PipelineAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentPipelineTrigger) SetAction(v PipelineAction)`

SetAction sets Action field to given value.


### GetCreated

`func (o *DeploymentPipelineTrigger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentPipelineTrigger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentPipelineTrigger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *DeploymentPipelineTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentPipelineTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentPipelineTrigger) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *DeploymentPipelineTrigger) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentPipelineTrigger) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentPipelineTrigger) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPipelineId

`func (o *DeploymentPipelineTrigger) GetPipelineId() int32`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *DeploymentPipelineTrigger) GetPipelineIdOk() (*int32, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *DeploymentPipelineTrigger) SetPipelineId(v int32)`

SetPipelineId sets PipelineId field to given value.


### GetUrl

`func (o *DeploymentPipelineTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentPipelineTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentPipelineTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


