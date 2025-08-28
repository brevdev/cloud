# VolumeMounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPath** | **string** | Filepath of the host. | 
**ContainerPath** | **string** | Filepath of the container. | 

## Methods

### NewVolumeMounts

`func NewVolumeMounts(hostPath string, containerPath string, ) *VolumeMounts`

NewVolumeMounts instantiates a new VolumeMounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeMountsWithDefaults

`func NewVolumeMountsWithDefaults() *VolumeMounts`

NewVolumeMountsWithDefaults instantiates a new VolumeMounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostPath

`func (o *VolumeMounts) GetHostPath() string`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *VolumeMounts) GetHostPathOk() (*string, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *VolumeMounts) SetHostPath(v string)`

SetHostPath sets HostPath field to given value.


### GetContainerPath

`func (o *VolumeMounts) GetContainerPath() string`

GetContainerPath returns the ContainerPath field if non-nil, zero value otherwise.

### GetContainerPathOk

`func (o *VolumeMounts) GetContainerPathOk() (*string, bool)`

GetContainerPathOk returns a tuple with the ContainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPath

`func (o *VolumeMounts) SetContainerPath(v string)`

SetContainerPath sets ContainerPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


