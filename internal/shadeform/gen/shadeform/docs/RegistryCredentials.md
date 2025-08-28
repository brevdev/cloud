# RegistryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username for the docker registry. | [optional] 
**Password** | Pointer to **string** | The password for the docker registry. | [optional] 

## Methods

### NewRegistryCredentials

`func NewRegistryCredentials() *RegistryCredentials`

NewRegistryCredentials instantiates a new RegistryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryCredentialsWithDefaults

`func NewRegistryCredentialsWithDefaults() *RegistryCredentials`

NewRegistryCredentialsWithDefaults instantiates a new RegistryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RegistryCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegistryCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegistryCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegistryCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *RegistryCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegistryCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegistryCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegistryCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


