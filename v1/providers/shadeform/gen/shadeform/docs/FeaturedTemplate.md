# FeaturedTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the featured template | 
**Description** | **string** | Description of the featured template | 
**Public** | **bool** | Whether the template is publicly available | 
**LaunchConfiguration** | Pointer to **map[string]interface{}** | Configuration for launching the instance | [optional] 
**AutoDelete** | Pointer to **map[string]interface{}** | Auto-deletion configuration | [optional] 
**Alert** | Pointer to **map[string]interface{}** | Alert configuration | [optional] 
**VolumeMount** | Pointer to **map[string]interface{}** | Volume mount configuration | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the template | [optional] 
**Envs** | Pointer to [**[]Env**](Env.md) | Environment variables for the template | [optional] 
**Networking** | Pointer to [**Networking**](Networking.md) |  | [optional] 

## Methods

### NewFeaturedTemplate

`func NewFeaturedTemplate(name string, description string, public bool, ) *FeaturedTemplate`

NewFeaturedTemplate instantiates a new FeaturedTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturedTemplateWithDefaults

`func NewFeaturedTemplateWithDefaults() *FeaturedTemplate`

NewFeaturedTemplateWithDefaults instantiates a new FeaturedTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeaturedTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeaturedTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeaturedTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FeaturedTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeaturedTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeaturedTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPublic

`func (o *FeaturedTemplate) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FeaturedTemplate) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FeaturedTemplate) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetLaunchConfiguration

`func (o *FeaturedTemplate) GetLaunchConfiguration() map[string]interface{}`

GetLaunchConfiguration returns the LaunchConfiguration field if non-nil, zero value otherwise.

### GetLaunchConfigurationOk

`func (o *FeaturedTemplate) GetLaunchConfigurationOk() (*map[string]interface{}, bool)`

GetLaunchConfigurationOk returns a tuple with the LaunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfiguration

`func (o *FeaturedTemplate) SetLaunchConfiguration(v map[string]interface{})`

SetLaunchConfiguration sets LaunchConfiguration field to given value.

### HasLaunchConfiguration

`func (o *FeaturedTemplate) HasLaunchConfiguration() bool`

HasLaunchConfiguration returns a boolean if a field has been set.

### GetAutoDelete

`func (o *FeaturedTemplate) GetAutoDelete() map[string]interface{}`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *FeaturedTemplate) GetAutoDeleteOk() (*map[string]interface{}, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *FeaturedTemplate) SetAutoDelete(v map[string]interface{})`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *FeaturedTemplate) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAlert

`func (o *FeaturedTemplate) GetAlert() map[string]interface{}`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *FeaturedTemplate) GetAlertOk() (*map[string]interface{}, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *FeaturedTemplate) SetAlert(v map[string]interface{})`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *FeaturedTemplate) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetVolumeMount

`func (o *FeaturedTemplate) GetVolumeMount() map[string]interface{}`

GetVolumeMount returns the VolumeMount field if non-nil, zero value otherwise.

### GetVolumeMountOk

`func (o *FeaturedTemplate) GetVolumeMountOk() (*map[string]interface{}, bool)`

GetVolumeMountOk returns a tuple with the VolumeMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMount

`func (o *FeaturedTemplate) SetVolumeMount(v map[string]interface{})`

SetVolumeMount sets VolumeMount field to given value.

### HasVolumeMount

`func (o *FeaturedTemplate) HasVolumeMount() bool`

HasVolumeMount returns a boolean if a field has been set.

### GetTags

`func (o *FeaturedTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeaturedTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeaturedTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeaturedTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvs

`func (o *FeaturedTemplate) GetEnvs() []Env`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *FeaturedTemplate) GetEnvsOk() (*[]Env, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *FeaturedTemplate) SetEnvs(v []Env)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *FeaturedTemplate) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetNetworking

`func (o *FeaturedTemplate) GetNetworking() Networking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *FeaturedTemplate) GetNetworkingOk() (*Networking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *FeaturedTemplate) SetNetworking(v Networking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *FeaturedTemplate) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


