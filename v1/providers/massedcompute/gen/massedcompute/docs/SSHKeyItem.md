# SSHKeyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the SSH key | [optional]
**Name** | Pointer to **string** | The name of the SSH key | [optional]
**PublicKey** | Pointer to **string** | The public key associated with the SSH key | [optional]

## Methods

### NewSSHKeyItem

`func NewSSHKeyItem() *SSHKeyItem`

NewSSHKeyItem instantiates a new SSHKeyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyItemWithDefaults

`func NewSSHKeyItemWithDefaults() *SSHKeyItem`

NewSSHKeyItemWithDefaults instantiates a new SSHKeyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSHKeyItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHKeyItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHKeyItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SSHKeyItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SSHKeyItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSHKeyItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSHKeyItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSHKeyItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *SSHKeyItem) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SSHKeyItem) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SSHKeyItem) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SSHKeyItem) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


