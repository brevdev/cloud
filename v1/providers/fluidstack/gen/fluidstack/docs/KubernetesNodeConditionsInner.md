# KubernetesNodeConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the condition | [optional] 
**Status** | Pointer to **string** | Status of the condition | [optional] 
**LastHeartbeatTime** | Pointer to **time.Time** | Last time the condition was updated | [optional] 
**LastTransitionTime** | Pointer to **time.Time** | Last time the condition transitioned from one status to another | [optional] 
**Reason** | Pointer to **string** | Reason for the condition&#39;s last transition | [optional] 
**Message** | Pointer to **string** | Human-readable message indicating details about the condition | [optional] 

## Methods

### NewKubernetesNodeConditionsInner

`func NewKubernetesNodeConditionsInner() *KubernetesNodeConditionsInner`

NewKubernetesNodeConditionsInner instantiates a new KubernetesNodeConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeConditionsInnerWithDefaults

`func NewKubernetesNodeConditionsInnerWithDefaults() *KubernetesNodeConditionsInner`

NewKubernetesNodeConditionsInnerWithDefaults instantiates a new KubernetesNodeConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KubernetesNodeConditionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodeConditionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodeConditionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodeConditionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesNodeConditionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesNodeConditionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesNodeConditionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesNodeConditionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastHeartbeatTime

`func (o *KubernetesNodeConditionsInner) GetLastHeartbeatTime() time.Time`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *KubernetesNodeConditionsInner) GetLastHeartbeatTimeOk() (*time.Time, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *KubernetesNodeConditionsInner) SetLastHeartbeatTime(v time.Time)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.

### HasLastHeartbeatTime

`func (o *KubernetesNodeConditionsInner) HasLastHeartbeatTime() bool`

HasLastHeartbeatTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *KubernetesNodeConditionsInner) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *KubernetesNodeConditionsInner) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *KubernetesNodeConditionsInner) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *KubernetesNodeConditionsInner) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReason

`func (o *KubernetesNodeConditionsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubernetesNodeConditionsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubernetesNodeConditionsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubernetesNodeConditionsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *KubernetesNodeConditionsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubernetesNodeConditionsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubernetesNodeConditionsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubernetesNodeConditionsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


