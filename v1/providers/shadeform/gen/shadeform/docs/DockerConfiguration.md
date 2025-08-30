# DockerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Specifies the docker image to be pulled and run on the instance at startup. | 
**Args** | Pointer to **string** | Specifies the container arguments passed into the image at runtime. | [optional] 
**SharedMemoryInGb** | Pointer to **int32** | Describes the amount of shared memory allocated for the container. Equivalent to using the --shm-size flag in the docker cli. If shared_memory_in_gb is not specified, then the container will use the host namespace which is the equivalent of --ipc&#x3D;host. | [optional] 
**Envs** | Pointer to [**[]Env**](Env.md) | List of environment variable name-value pairs that will be passed to the docker container. | [optional] 
**PortMappings** | Pointer to [**[]PortMappings**](PortMappings.md) | List of port mappings between the host instance and the docker container. Equivalent of -p flag for docker run command. | [optional] 
**VolumeMounts** | Pointer to [**[]VolumeMounts**](VolumeMounts.md) | List of volume mounts between the host instance and the docker container. Equivalent of -v flag for docker run command. | [optional] 
**RegistryCredentials** | Pointer to [**RegistryCredentials**](RegistryCredentials.md) |  | [optional] 

## Methods

### NewDockerConfiguration

`func NewDockerConfiguration(image string, ) *DockerConfiguration`

NewDockerConfiguration instantiates a new DockerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerConfigurationWithDefaults

`func NewDockerConfigurationWithDefaults() *DockerConfiguration`

NewDockerConfigurationWithDefaults instantiates a new DockerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *DockerConfiguration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DockerConfiguration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DockerConfiguration) SetImage(v string)`

SetImage sets Image field to given value.


### GetArgs

`func (o *DockerConfiguration) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *DockerConfiguration) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *DockerConfiguration) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *DockerConfiguration) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetSharedMemoryInGb

`func (o *DockerConfiguration) GetSharedMemoryInGb() int32`

GetSharedMemoryInGb returns the SharedMemoryInGb field if non-nil, zero value otherwise.

### GetSharedMemoryInGbOk

`func (o *DockerConfiguration) GetSharedMemoryInGbOk() (*int32, bool)`

GetSharedMemoryInGbOk returns a tuple with the SharedMemoryInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMemoryInGb

`func (o *DockerConfiguration) SetSharedMemoryInGb(v int32)`

SetSharedMemoryInGb sets SharedMemoryInGb field to given value.

### HasSharedMemoryInGb

`func (o *DockerConfiguration) HasSharedMemoryInGb() bool`

HasSharedMemoryInGb returns a boolean if a field has been set.

### GetEnvs

`func (o *DockerConfiguration) GetEnvs() []Env`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *DockerConfiguration) GetEnvsOk() (*[]Env, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *DockerConfiguration) SetEnvs(v []Env)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *DockerConfiguration) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetPortMappings

`func (o *DockerConfiguration) GetPortMappings() []PortMappings`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *DockerConfiguration) GetPortMappingsOk() (*[]PortMappings, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *DockerConfiguration) SetPortMappings(v []PortMappings)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *DockerConfiguration) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *DockerConfiguration) GetVolumeMounts() []VolumeMounts`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *DockerConfiguration) GetVolumeMountsOk() (*[]VolumeMounts, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *DockerConfiguration) SetVolumeMounts(v []VolumeMounts)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *DockerConfiguration) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.

### GetRegistryCredentials

`func (o *DockerConfiguration) GetRegistryCredentials() RegistryCredentials`

GetRegistryCredentials returns the RegistryCredentials field if non-nil, zero value otherwise.

### GetRegistryCredentialsOk

`func (o *DockerConfiguration) GetRegistryCredentialsOk() (*RegistryCredentials, bool)`

GetRegistryCredentialsOk returns a tuple with the RegistryCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryCredentials

`func (o *DockerConfiguration) SetRegistryCredentials(v RegistryCredentials)`

SetRegistryCredentials sets RegistryCredentials field to given value.

### HasRegistryCredentials

`func (o *DockerConfiguration) HasRegistryCredentials() bool`

HasRegistryCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


