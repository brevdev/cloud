# Accelerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the hardware accelerator | 
**Product** | **string** | Product name of the accelerator | 
**DriverVersion** | **string** | Version of the driver for the accelerator | 

## Methods

### NewAccelerator

`func NewAccelerator(type_ string, product string, driverVersion string, ) *Accelerator`

NewAccelerator instantiates a new Accelerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorWithDefaults

`func NewAcceleratorWithDefaults() *Accelerator`

NewAcceleratorWithDefaults instantiates a new Accelerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Accelerator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Accelerator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Accelerator) SetType(v string)`

SetType sets Type field to given value.


### GetProduct

`func (o *Accelerator) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Accelerator) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Accelerator) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetDriverVersion

`func (o *Accelerator) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *Accelerator) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *Accelerator) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


