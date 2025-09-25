# DeploymentTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**DeploymentTaskActionEnum**](DeploymentTaskActionEnum.md) | The action the task will perform  * &#x60;start_instances&#x60; - Start all instances in the deployment * &#x60;stop_instances&#x60; - Stop all instances in the deployment | 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Deployment** | [**ClusterDeployment**](ClusterDeployment.md) |  | 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Retries** | **int32** | Number of times the task has been retried | [readonly] 
**Status** | [**StatusEnum**](StatusEnum.md) | Current status of the task  * &#x60;completed&#x60; - completed * &#x60;failed&#x60; - failed * &#x60;pending&#x60; - pending * &#x60;processing&#x60; - processing * &#x60;retrying&#x60; - retrying | [readonly] 
**StatusText** | **string** | Messages related to the current task status | [readonly] 

## Methods

### NewDeploymentTask

`func NewDeploymentTask(action DeploymentTaskActionEnum, created time.Time, deployment ClusterDeployment, id string, modified time.Time, retries int32, status StatusEnum, statusText string, ) *DeploymentTask`

NewDeploymentTask instantiates a new DeploymentTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentTaskWithDefaults

`func NewDeploymentTaskWithDefaults() *DeploymentTask`

NewDeploymentTaskWithDefaults instantiates a new DeploymentTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeploymentTask) GetAction() DeploymentTaskActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentTask) GetActionOk() (*DeploymentTaskActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentTask) SetAction(v DeploymentTaskActionEnum)`

SetAction sets Action field to given value.


### GetCreated

`func (o *DeploymentTask) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentTask) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentTask) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *DeploymentTask) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *DeploymentTask) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *DeploymentTask) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetId

`func (o *DeploymentTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentTask) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *DeploymentTask) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentTask) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentTask) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetRetries

`func (o *DeploymentTask) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *DeploymentTask) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *DeploymentTask) SetRetries(v int32)`

SetRetries sets Retries field to given value.


### GetStatus

`func (o *DeploymentTask) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentTask) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentTask) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetStatusText

`func (o *DeploymentTask) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *DeploymentTask) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *DeploymentTask) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


