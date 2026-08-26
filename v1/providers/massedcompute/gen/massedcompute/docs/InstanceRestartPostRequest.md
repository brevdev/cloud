# InstanceRestartPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceUuids** | **[]string** | The ID or IDs of instances to restart |

## Methods

### NewInstanceRestartPostRequest

`func NewInstanceRestartPostRequest(instanceUuids []string, ) *InstanceRestartPostRequest`

NewInstanceRestartPostRequest instantiates a new InstanceRestartPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceRestartPostRequestWithDefaults

`func NewInstanceRestartPostRequestWithDefaults() *InstanceRestartPostRequest`

NewInstanceRestartPostRequestWithDefaults instantiates a new InstanceRestartPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceUuids

`func (o *InstanceRestartPostRequest) GetInstanceUuids() []string`

GetInstanceUuids returns the InstanceUuids field if non-nil, zero value otherwise.

### GetInstanceUuidsOk

`func (o *InstanceRestartPostRequest) GetInstanceUuidsOk() (*[]string, bool)`

GetInstanceUuidsOk returns a tuple with the InstanceUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuids

`func (o *InstanceRestartPostRequest) SetInstanceUuids(v []string)`

SetInstanceUuids sets InstanceUuids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


