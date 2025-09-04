# BootTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinBootInSec** | Pointer to **int32** | boot time minimum estimate | [optional] 
**MaxBootInSec** | Pointer to **int32** | boot time maximum estimate | [optional] 

## Methods

### NewBootTime

`func NewBootTime() *BootTime`

NewBootTime instantiates a new BootTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootTimeWithDefaults

`func NewBootTimeWithDefaults() *BootTime`

NewBootTimeWithDefaults instantiates a new BootTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinBootInSec

`func (o *BootTime) GetMinBootInSec() int32`

GetMinBootInSec returns the MinBootInSec field if non-nil, zero value otherwise.

### GetMinBootInSecOk

`func (o *BootTime) GetMinBootInSecOk() (*int32, bool)`

GetMinBootInSecOk returns a tuple with the MinBootInSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBootInSec

`func (o *BootTime) SetMinBootInSec(v int32)`

SetMinBootInSec sets MinBootInSec field to given value.

### HasMinBootInSec

`func (o *BootTime) HasMinBootInSec() bool`

HasMinBootInSec returns a boolean if a field has been set.

### GetMaxBootInSec

`func (o *BootTime) GetMaxBootInSec() int32`

GetMaxBootInSec returns the MaxBootInSec field if non-nil, zero value otherwise.

### GetMaxBootInSecOk

`func (o *BootTime) GetMaxBootInSecOk() (*int32, bool)`

GetMaxBootInSecOk returns a tuple with the MaxBootInSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBootInSec

`func (o *BootTime) SetMaxBootInSec(v int32)`

SetMaxBootInSec sets MaxBootInSec field to given value.

### HasMaxBootInSec

`func (o *BootTime) HasMaxBootInSec() bool`

HasMaxBootInSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


