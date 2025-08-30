# CreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | [**Cloud**](Cloud.md) |  | 
**Region** | **string** | Specifies the region. | 
**ShadeInstanceType** | **string** | The Shadeform standardized instance type. See this [explanation](/getting-started/concepts#shade-instance-type-and-cloud-instance-type) for more details. | 
**ShadeCloud** | **bool** | Specifies if the instance is launched in [Shade Cloud](/getting-started/concepts#shade-cloud) or in a linked cloud account. | 
**Name** | **string** | The name of the instance | 
**Os** | Pointer to **string** | The operating system of the instance. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the template to use for this instance | [optional] 
**LaunchConfiguration** | Pointer to [**LaunchConfiguration**](LaunchConfiguration.md) |  | [optional] 
**VolumeIds** | Pointer to **[]string** | List of volume IDs to be mounted. Currently only supports 1 volume at a time. | [optional] 
**SshKeyId** | Pointer to **string** | The ID of the SSH Key. | [optional] 
**AutoDelete** | Pointer to [**AutoDelete**](AutoDelete.md) |  | [optional] 
**Alert** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**VolumeMount** | Pointer to [**VolumeMount**](VolumeMount.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Add custom, searchable tags to instances. | [optional] 
**Envs** | Pointer to [**[]Env**](Env.md) | List of environment variable name and values to automatically add to the instance | [optional] 

## Methods

### NewCreateRequest

`func NewCreateRequest(cloud Cloud, region string, shadeInstanceType string, shadeCloud bool, name string, ) *CreateRequest`

NewCreateRequest instantiates a new CreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestWithDefaults

`func NewCreateRequestWithDefaults() *CreateRequest`

NewCreateRequestWithDefaults instantiates a new CreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *CreateRequest) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CreateRequest) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CreateRequest) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *CreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetShadeInstanceType

`func (o *CreateRequest) GetShadeInstanceType() string`

GetShadeInstanceType returns the ShadeInstanceType field if non-nil, zero value otherwise.

### GetShadeInstanceTypeOk

`func (o *CreateRequest) GetShadeInstanceTypeOk() (*string, bool)`

GetShadeInstanceTypeOk returns a tuple with the ShadeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeInstanceType

`func (o *CreateRequest) SetShadeInstanceType(v string)`

SetShadeInstanceType sets ShadeInstanceType field to given value.


### GetShadeCloud

`func (o *CreateRequest) GetShadeCloud() bool`

GetShadeCloud returns the ShadeCloud field if non-nil, zero value otherwise.

### GetShadeCloudOk

`func (o *CreateRequest) GetShadeCloudOk() (*bool, bool)`

GetShadeCloudOk returns a tuple with the ShadeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeCloud

`func (o *CreateRequest) SetShadeCloud(v bool)`

SetShadeCloud sets ShadeCloud field to given value.


### GetName

`func (o *CreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOs

`func (o *CreateRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CreateRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CreateRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CreateRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetLaunchConfiguration

`func (o *CreateRequest) GetLaunchConfiguration() LaunchConfiguration`

GetLaunchConfiguration returns the LaunchConfiguration field if non-nil, zero value otherwise.

### GetLaunchConfigurationOk

`func (o *CreateRequest) GetLaunchConfigurationOk() (*LaunchConfiguration, bool)`

GetLaunchConfigurationOk returns a tuple with the LaunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfiguration

`func (o *CreateRequest) SetLaunchConfiguration(v LaunchConfiguration)`

SetLaunchConfiguration sets LaunchConfiguration field to given value.

### HasLaunchConfiguration

`func (o *CreateRequest) HasLaunchConfiguration() bool`

HasLaunchConfiguration returns a boolean if a field has been set.

### GetVolumeIds

`func (o *CreateRequest) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *CreateRequest) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *CreateRequest) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.

### HasVolumeIds

`func (o *CreateRequest) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### GetSshKeyId

`func (o *CreateRequest) GetSshKeyId() string`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *CreateRequest) GetSshKeyIdOk() (*string, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *CreateRequest) SetSshKeyId(v string)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *CreateRequest) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetAutoDelete

`func (o *CreateRequest) GetAutoDelete() AutoDelete`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *CreateRequest) GetAutoDeleteOk() (*AutoDelete, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *CreateRequest) SetAutoDelete(v AutoDelete)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *CreateRequest) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAlert

`func (o *CreateRequest) GetAlert() Alert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *CreateRequest) GetAlertOk() (*Alert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *CreateRequest) SetAlert(v Alert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *CreateRequest) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetVolumeMount

`func (o *CreateRequest) GetVolumeMount() VolumeMount`

GetVolumeMount returns the VolumeMount field if non-nil, zero value otherwise.

### GetVolumeMountOk

`func (o *CreateRequest) GetVolumeMountOk() (*VolumeMount, bool)`

GetVolumeMountOk returns a tuple with the VolumeMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMount

`func (o *CreateRequest) SetVolumeMount(v VolumeMount)`

SetVolumeMount sets VolumeMount field to given value.

### HasVolumeMount

`func (o *CreateRequest) HasVolumeMount() bool`

HasVolumeMount returns a boolean if a field has been set.

### GetTags

`func (o *CreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvs

`func (o *CreateRequest) GetEnvs() []Env`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *CreateRequest) GetEnvsOk() (*[]Env, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *CreateRequest) SetEnvs(v []Env)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *CreateRequest) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


