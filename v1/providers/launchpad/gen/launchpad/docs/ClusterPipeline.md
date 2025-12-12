# ClusterPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PipelineAction**](PipelineAction.md) | Action for the pipeline to run  * &#x60;apply&#x60; - apply * &#x60;destroy&#x60; - destroy * &#x60;notify&#x60; - notify | [readonly] 
**Cluster** | [**ClusterPipelineCluster**](ClusterPipelineCluster.md) |  | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**PipelineId** | **int32** | GitLab pipeline ID | [readonly] 
**RequestId** | **string** | Request ID that the pipeline was triggered for | [readonly] 
**Url** | **string** | URL for the pipeline details | [readonly] 

## Methods

### NewClusterPipeline

`func NewClusterPipeline(action PipelineAction, cluster ClusterPipelineCluster, created time.Time, id string, modified time.Time, pipelineId int32, requestId string, url string, ) *ClusterPipeline`

NewClusterPipeline instantiates a new ClusterPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPipelineWithDefaults

`func NewClusterPipelineWithDefaults() *ClusterPipeline`

NewClusterPipelineWithDefaults instantiates a new ClusterPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ClusterPipeline) GetAction() PipelineAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterPipeline) GetActionOk() (*PipelineAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterPipeline) SetAction(v PipelineAction)`

SetAction sets Action field to given value.


### GetCluster

`func (o *ClusterPipeline) GetCluster() ClusterPipelineCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterPipeline) GetClusterOk() (*ClusterPipelineCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterPipeline) SetCluster(v ClusterPipelineCluster)`

SetCluster sets Cluster field to given value.


### GetCreated

`func (o *ClusterPipeline) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClusterPipeline) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClusterPipeline) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *ClusterPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPipeline) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *ClusterPipeline) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ClusterPipeline) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ClusterPipeline) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPipelineId

`func (o *ClusterPipeline) GetPipelineId() int32`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *ClusterPipeline) GetPipelineIdOk() (*int32, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *ClusterPipeline) SetPipelineId(v int32)`

SetPipelineId sets PipelineId field to given value.


### GetRequestId

`func (o *ClusterPipeline) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClusterPipeline) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClusterPipeline) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUrl

`func (o *ClusterPipeline) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterPipeline) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterPipeline) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


