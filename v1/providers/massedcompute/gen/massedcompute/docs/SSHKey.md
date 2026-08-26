# SSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeys** | Pointer to [**[]SSHKeyItem**](SSHKeyItem.md) |  | [optional]

## Methods

### NewSSHKey

`func NewSSHKey() *SSHKey`

NewSSHKey instantiates a new SSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyWithDefaults

`func NewSSHKeyWithDefaults() *SSHKey`

NewSSHKeyWithDefaults instantiates a new SSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeys

`func (o *SSHKey) GetSshKeys() []SSHKeyItem`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *SSHKey) GetSshKeysOk() (*[]SSHKeyItem, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *SSHKey) SetSshKeys(v []SSHKeyItem)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *SSHKey) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


