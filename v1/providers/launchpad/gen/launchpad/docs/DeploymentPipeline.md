# DeploymentPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PipelineAction**](PipelineAction.md) | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy * &#x60;notify&#x60; - notify | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Deployment** | [**ClusterDeployment**](ClusterDeployment.md) |  | 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**PipelineId** | **int64** | GitLab pipeline ID | 
**Url** | **string** | URL for the pipeline details | 

## Methods

### NewDeploymentPipeline

`func NewDeploymentPipeline(action PipelineAction, created time.Time, deployment ClusterDeployment, id string, modified time.Time, pipelineId int64, url string, ) *DeploymentPipeline`

NewDeploymentPipeline instantiates a new DeploymentPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPipelineWithDefaults

`func NewDeploymentPipelineWithDefaults() *DeploymentPipeline`

NewDeploymentPipelineWithDefaults instantiates a new DeploymentPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentPipeline) GetAction() PipelineAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentPipeline) GetActionOk() (*PipelineAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentPipeline) SetAction(v PipelineAction)`

SetAction sets Action field to given value.


### GetCreated

`func (o *DeploymentPipeline) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentPipeline) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentPipeline) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *DeploymentPipeline) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *DeploymentPipeline) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *DeploymentPipeline) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetId

`func (o *DeploymentPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentPipeline) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *DeploymentPipeline) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentPipeline) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentPipeline) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPipelineId

`func (o *DeploymentPipeline) GetPipelineId() int64`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *DeploymentPipeline) GetPipelineIdOk() (*int64, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *DeploymentPipeline) SetPipelineId(v int64)`

SetPipelineId sets PipelineId field to given value.


### GetUrl

`func (o *DeploymentPipeline) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentPipeline) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentPipeline) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


