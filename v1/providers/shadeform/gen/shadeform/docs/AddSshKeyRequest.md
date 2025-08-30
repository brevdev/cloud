# AddSshKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the SSH Key. | 
**PublicKey** | **string** | The public key. | 

## Methods

### NewAddSshKeyRequest

`func NewAddSshKeyRequest(name string, publicKey string, ) *AddSshKeyRequest`

NewAddSshKeyRequest instantiates a new AddSshKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSshKeyRequestWithDefaults

`func NewAddSshKeyRequestWithDefaults() *AddSshKeyRequest`

NewAddSshKeyRequestWithDefaults instantiates a new AddSshKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddSshKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSshKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSshKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *AddSshKeyRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AddSshKeyRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AddSshKeyRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


