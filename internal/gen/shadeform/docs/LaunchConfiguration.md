# LaunchConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of launch configuration. See [Launch Configuration](/getting-started/concepts#launch-configuration) for more details. | 
**DockerConfiguration** | Pointer to [**DockerConfiguration**](DockerConfiguration.md) |  | [optional] 
**ScriptConfiguration** | Pointer to [**ScriptConfiguration**](ScriptConfiguration.md) |  | [optional] 

## Methods

### NewLaunchConfiguration

`func NewLaunchConfiguration(type_ string, ) *LaunchConfiguration`

NewLaunchConfiguration instantiates a new LaunchConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchConfigurationWithDefaults

`func NewLaunchConfigurationWithDefaults() *LaunchConfiguration`

NewLaunchConfigurationWithDefaults instantiates a new LaunchConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LaunchConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetDockerConfiguration

`func (o *LaunchConfiguration) GetDockerConfiguration() DockerConfiguration`

GetDockerConfiguration returns the DockerConfiguration field if non-nil, zero value otherwise.

### GetDockerConfigurationOk

`func (o *LaunchConfiguration) GetDockerConfigurationOk() (*DockerConfiguration, bool)`

GetDockerConfigurationOk returns a tuple with the DockerConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerConfiguration

`func (o *LaunchConfiguration) SetDockerConfiguration(v DockerConfiguration)`

SetDockerConfiguration sets DockerConfiguration field to given value.

### HasDockerConfiguration

`func (o *LaunchConfiguration) HasDockerConfiguration() bool`

HasDockerConfiguration returns a boolean if a field has been set.

### GetScriptConfiguration

`func (o *LaunchConfiguration) GetScriptConfiguration() ScriptConfiguration`

GetScriptConfiguration returns the ScriptConfiguration field if non-nil, zero value otherwise.

### GetScriptConfigurationOk

`func (o *LaunchConfiguration) GetScriptConfigurationOk() (*ScriptConfiguration, bool)`

GetScriptConfigurationOk returns a tuple with the ScriptConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptConfiguration

`func (o *LaunchConfiguration) SetScriptConfiguration(v ScriptConfiguration)`

SetScriptConfiguration sets ScriptConfiguration field to given value.

### HasScriptConfiguration

`func (o *LaunchConfiguration) HasScriptConfiguration() bool`

HasScriptConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


