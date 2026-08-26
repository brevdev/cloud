# SshKeysPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the SSH key |
**PublicKey** | **string** | The public key associated with the SSH key |

## Methods

### NewSshKeysPostRequest

`func NewSshKeysPostRequest(name string, publicKey string, ) *SshKeysPostRequest`

NewSshKeysPostRequest instantiates a new SshKeysPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeysPostRequestWithDefaults

`func NewSshKeysPostRequestWithDefaults() *SshKeysPostRequest`

NewSshKeysPostRequestWithDefaults instantiates a new SshKeysPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshKeysPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeysPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeysPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *SshKeysPostRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshKeysPostRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshKeysPostRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


