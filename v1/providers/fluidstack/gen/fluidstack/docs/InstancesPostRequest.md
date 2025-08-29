# InstancesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the instance | 
**Type** | **string** | Instance type | 
**Image** | Pointer to **string** | Operating system image for the instance | [optional] [default to "image://ubuntu22.04"]
**Preemptible** | Pointer to **bool** | Whether the instance is preemptible | [optional] [default to false]
**Ephemeral** | Pointer to **bool** | Whether the instance is ephemeral | [optional] [default to false]
**UserData** | Pointer to **NullableString** | Base64-encoded cloud-init user data for instance initialization | [optional] 
**Filesystems** | Pointer to **[]string** | List of filesystem UUIDs attached to the instance | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for the instance. Tag keys must be 1–63 characters, start and end with alphanumerics, and may include &#x60;-&#x60;, &#x60;_&#x60;, and &#x60;.&#x60;.  | [optional] 

## Methods

### NewInstancesPostRequest

`func NewInstancesPostRequest(name string, type_ string, ) *InstancesPostRequest`

NewInstancesPostRequest instantiates a new InstancesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesPostRequestWithDefaults

`func NewInstancesPostRequestWithDefaults() *InstancesPostRequest`

NewInstancesPostRequestWithDefaults instantiates a new InstancesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstancesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstancesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstancesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InstancesPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstancesPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstancesPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetImage

`func (o *InstancesPostRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstancesPostRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstancesPostRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *InstancesPostRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPreemptible

`func (o *InstancesPostRequest) GetPreemptible() bool`

GetPreemptible returns the Preemptible field if non-nil, zero value otherwise.

### GetPreemptibleOk

`func (o *InstancesPostRequest) GetPreemptibleOk() (*bool, bool)`

GetPreemptibleOk returns a tuple with the Preemptible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptible

`func (o *InstancesPostRequest) SetPreemptible(v bool)`

SetPreemptible sets Preemptible field to given value.

### HasPreemptible

`func (o *InstancesPostRequest) HasPreemptible() bool`

HasPreemptible returns a boolean if a field has been set.

### GetEphemeral

`func (o *InstancesPostRequest) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *InstancesPostRequest) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *InstancesPostRequest) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *InstancesPostRequest) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetUserData

`func (o *InstancesPostRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstancesPostRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstancesPostRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstancesPostRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *InstancesPostRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *InstancesPostRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetFilesystems

`func (o *InstancesPostRequest) GetFilesystems() []string`

GetFilesystems returns the Filesystems field if non-nil, zero value otherwise.

### GetFilesystemsOk

`func (o *InstancesPostRequest) GetFilesystemsOk() (*[]string, bool)`

GetFilesystemsOk returns a tuple with the Filesystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystems

`func (o *InstancesPostRequest) SetFilesystems(v []string)`

SetFilesystems sets Filesystems field to given value.

### HasFilesystems

`func (o *InstancesPostRequest) HasFilesystems() bool`

HasFilesystems returns a boolean if a field has been set.

### GetTags

`func (o *InstancesPostRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstancesPostRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstancesPostRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstancesPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


