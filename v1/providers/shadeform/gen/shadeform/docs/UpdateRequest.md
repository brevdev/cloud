# UpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the instance | [optional] 
**AutoDelete** | Pointer to [**AutoDelete**](AutoDelete.md) |  | [optional] 
**Alert** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Add custom, searchable tags to instances. | [optional] 

## Methods

### NewUpdateRequest

`func NewUpdateRequest() *UpdateRequest`

NewUpdateRequest instantiates a new UpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestWithDefaults

`func NewUpdateRequestWithDefaults() *UpdateRequest`

NewUpdateRequestWithDefaults instantiates a new UpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoDelete

`func (o *UpdateRequest) GetAutoDelete() AutoDelete`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *UpdateRequest) GetAutoDeleteOk() (*AutoDelete, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *UpdateRequest) SetAutoDelete(v AutoDelete)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *UpdateRequest) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAlert

`func (o *UpdateRequest) GetAlert() Alert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UpdateRequest) GetAlertOk() (*Alert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UpdateRequest) SetAlert(v Alert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *UpdateRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetTags

`func (o *UpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


