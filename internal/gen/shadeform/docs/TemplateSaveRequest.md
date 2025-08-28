# TemplateSaveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template | 
**Description** | Pointer to **string** | Description of the template | [optional] 
**Public** | Pointer to **bool** | Whether the template is publicly available | [optional] 
**LaunchConfiguration** | Pointer to [**LaunchConfiguration**](LaunchConfiguration.md) |  | [optional] 
**AutoDelete** | Pointer to [**AutoDelete**](AutoDelete.md) |  | [optional] 
**Alert** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**VolumeMount** | Pointer to [**VolumeMount**](VolumeMount.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the template | [optional] 
**Envs** | Pointer to [**[]Env**](Env.md) | Environment variables for the template | [optional] 
**Networking** | Pointer to [**Networking**](Networking.md) |  | [optional] 

## Methods

### NewTemplateSaveRequest

`func NewTemplateSaveRequest(name string, ) *TemplateSaveRequest`

NewTemplateSaveRequest instantiates a new TemplateSaveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSaveRequestWithDefaults

`func NewTemplateSaveRequestWithDefaults() *TemplateSaveRequest`

NewTemplateSaveRequestWithDefaults instantiates a new TemplateSaveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateSaveRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateSaveRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateSaveRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TemplateSaveRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateSaveRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateSaveRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateSaveRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublic

`func (o *TemplateSaveRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *TemplateSaveRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *TemplateSaveRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *TemplateSaveRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetLaunchConfiguration

`func (o *TemplateSaveRequest) GetLaunchConfiguration() LaunchConfiguration`

GetLaunchConfiguration returns the LaunchConfiguration field if non-nil, zero value otherwise.

### GetLaunchConfigurationOk

`func (o *TemplateSaveRequest) GetLaunchConfigurationOk() (*LaunchConfiguration, bool)`

GetLaunchConfigurationOk returns a tuple with the LaunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfiguration

`func (o *TemplateSaveRequest) SetLaunchConfiguration(v LaunchConfiguration)`

SetLaunchConfiguration sets LaunchConfiguration field to given value.

### HasLaunchConfiguration

`func (o *TemplateSaveRequest) HasLaunchConfiguration() bool`

HasLaunchConfiguration returns a boolean if a field has been set.

### GetAutoDelete

`func (o *TemplateSaveRequest) GetAutoDelete() AutoDelete`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *TemplateSaveRequest) GetAutoDeleteOk() (*AutoDelete, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *TemplateSaveRequest) SetAutoDelete(v AutoDelete)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *TemplateSaveRequest) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAlert

`func (o *TemplateSaveRequest) GetAlert() Alert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *TemplateSaveRequest) GetAlertOk() (*Alert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *TemplateSaveRequest) SetAlert(v Alert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *TemplateSaveRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetVolumeMount

`func (o *TemplateSaveRequest) GetVolumeMount() VolumeMount`

GetVolumeMount returns the VolumeMount field if non-nil, zero value otherwise.

### GetVolumeMountOk

`func (o *TemplateSaveRequest) GetVolumeMountOk() (*VolumeMount, bool)`

GetVolumeMountOk returns a tuple with the VolumeMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMount

`func (o *TemplateSaveRequest) SetVolumeMount(v VolumeMount)`

SetVolumeMount sets VolumeMount field to given value.

### HasVolumeMount

`func (o *TemplateSaveRequest) HasVolumeMount() bool`

HasVolumeMount returns a boolean if a field has been set.

### GetTags

`func (o *TemplateSaveRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateSaveRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateSaveRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateSaveRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvs

`func (o *TemplateSaveRequest) GetEnvs() []Env`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *TemplateSaveRequest) GetEnvsOk() (*[]Env, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *TemplateSaveRequest) SetEnvs(v []Env)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *TemplateSaveRequest) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetNetworking

`func (o *TemplateSaveRequest) GetNetworking() Networking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *TemplateSaveRequest) GetNetworkingOk() (*Networking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *TemplateSaveRequest) SetNetworking(v Networking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *TemplateSaveRequest) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


