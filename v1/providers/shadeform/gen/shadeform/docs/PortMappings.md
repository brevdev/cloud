# PortMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPort** | **int32** | Port of the host. | 
**ContainerPort** | **int32** | Port of the container. | 

## Methods

### NewPortMappings

`func NewPortMappings(hostPort int32, containerPort int32, ) *PortMappings`

NewPortMappings instantiates a new PortMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMappingsWithDefaults

`func NewPortMappingsWithDefaults() *PortMappings`

NewPortMappingsWithDefaults instantiates a new PortMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostPort

`func (o *PortMappings) GetHostPort() int32`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *PortMappings) GetHostPortOk() (*int32, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *PortMappings) SetHostPort(v int32)`

SetHostPort sets HostPort field to given value.


### GetContainerPort

`func (o *PortMappings) GetContainerPort() int32`

GetContainerPort returns the ContainerPort field if non-nil, zero value otherwise.

### GetContainerPortOk

`func (o *PortMappings) GetContainerPortOk() (*int32, bool)`

GetContainerPortOk returns a tuple with the ContainerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPort

`func (o *PortMappings) SetContainerPort(v int32)`

SetContainerPort sets ContainerPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


