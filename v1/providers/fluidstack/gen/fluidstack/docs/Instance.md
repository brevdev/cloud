# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the instance | 
**Name** | **string** | Name of the instance | 
**Type** | **string** | Instance type | 
**Ip** | **NullableString** | Private IP address of the instance | 
**State** | [**InstanceState**](InstanceState.md) |  | 
**Image** | **string** | Operating system image for the instance | 
**Preemptible** | **bool** | Whether the instance is preemptible | 
**Ephemeral** | **bool** | Whether the instance is ephemeral | 
**UserData** | **NullableString** | Base64-encoded cloud-init user data for instance initialization | 
**Filesystems** | **[]string** | List of filesystem UUIDs attached to the instance | 
**Tags** | **map[string]string** | Tags for the instance | 

## Methods

### NewInstance

`func NewInstance(id string, name string, type_ string, ip NullableString, state InstanceState, image string, preemptible bool, ephemeral bool, userData NullableString, filesystems []string, tags map[string]string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Instance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v string)`

SetType sets Type field to given value.


### GetIp

`func (o *Instance) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Instance) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Instance) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *Instance) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Instance) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetState

`func (o *Instance) GetState() InstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*InstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v InstanceState)`

SetState sets State field to given value.


### GetImage

`func (o *Instance) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Instance) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Instance) SetImage(v string)`

SetImage sets Image field to given value.


### GetPreemptible

`func (o *Instance) GetPreemptible() bool`

GetPreemptible returns the Preemptible field if non-nil, zero value otherwise.

### GetPreemptibleOk

`func (o *Instance) GetPreemptibleOk() (*bool, bool)`

GetPreemptibleOk returns a tuple with the Preemptible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptible

`func (o *Instance) SetPreemptible(v bool)`

SetPreemptible sets Preemptible field to given value.


### GetEphemeral

`func (o *Instance) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *Instance) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *Instance) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.


### GetUserData

`func (o *Instance) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *Instance) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *Instance) SetUserData(v string)`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *Instance) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *Instance) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetFilesystems

`func (o *Instance) GetFilesystems() []string`

GetFilesystems returns the Filesystems field if non-nil, zero value otherwise.

### GetFilesystemsOk

`func (o *Instance) GetFilesystemsOk() (*[]string, bool)`

GetFilesystemsOk returns a tuple with the Filesystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystems

`func (o *Instance) SetFilesystems(v []string)`

SetFilesystems sets Filesystems field to given value.


### GetTags

`func (o *Instance) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v map[string]string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


