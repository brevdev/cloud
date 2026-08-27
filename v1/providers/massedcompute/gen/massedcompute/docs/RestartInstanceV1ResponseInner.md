# RestartInstanceV1ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**Ip** | Pointer to **string** |  | [optional]
**Status** | Pointer to **string** |  | [optional]
**SshKeyNames** | Pointer to **[]string** |  | [optional]
**FileSystemNames** | Pointer to **[]string** |  | [optional]
**Region** | Pointer to [**RestartInstanceV1ResponseInnerRegion**](RestartInstanceV1ResponseInnerRegion.md) |  | [optional]
**InstanceType** | Pointer to [**RestartInstanceV1ResponseInnerInstanceType**](RestartInstanceV1ResponseInnerInstanceType.md) |  | [optional]
**JupyterToken** | Pointer to **string** |  | [optional]
**JupyterUrl** | Pointer to **string** |  | [optional]

## Methods

### NewRestartInstanceV1ResponseInner

`func NewRestartInstanceV1ResponseInner() *RestartInstanceV1ResponseInner`

NewRestartInstanceV1ResponseInner instantiates a new RestartInstanceV1ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartInstanceV1ResponseInnerWithDefaults

`func NewRestartInstanceV1ResponseInnerWithDefaults() *RestartInstanceV1ResponseInner`

NewRestartInstanceV1ResponseInnerWithDefaults instantiates a new RestartInstanceV1ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestartInstanceV1ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestartInstanceV1ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestartInstanceV1ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestartInstanceV1ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RestartInstanceV1ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestartInstanceV1ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestartInstanceV1ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestartInstanceV1ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *RestartInstanceV1ResponseInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RestartInstanceV1ResponseInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RestartInstanceV1ResponseInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RestartInstanceV1ResponseInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetStatus

`func (o *RestartInstanceV1ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestartInstanceV1ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestartInstanceV1ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestartInstanceV1ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSshKeyNames

`func (o *RestartInstanceV1ResponseInner) GetSshKeyNames() []string`

GetSshKeyNames returns the SshKeyNames field if non-nil, zero value otherwise.

### GetSshKeyNamesOk

`func (o *RestartInstanceV1ResponseInner) GetSshKeyNamesOk() (*[]string, bool)`

GetSshKeyNamesOk returns a tuple with the SshKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyNames

`func (o *RestartInstanceV1ResponseInner) SetSshKeyNames(v []string)`

SetSshKeyNames sets SshKeyNames field to given value.

### HasSshKeyNames

`func (o *RestartInstanceV1ResponseInner) HasSshKeyNames() bool`

HasSshKeyNames returns a boolean if a field has been set.

### GetFileSystemNames

`func (o *RestartInstanceV1ResponseInner) GetFileSystemNames() []string`

GetFileSystemNames returns the FileSystemNames field if non-nil, zero value otherwise.

### GetFileSystemNamesOk

`func (o *RestartInstanceV1ResponseInner) GetFileSystemNamesOk() (*[]string, bool)`

GetFileSystemNamesOk returns a tuple with the FileSystemNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemNames

`func (o *RestartInstanceV1ResponseInner) SetFileSystemNames(v []string)`

SetFileSystemNames sets FileSystemNames field to given value.

### HasFileSystemNames

`func (o *RestartInstanceV1ResponseInner) HasFileSystemNames() bool`

HasFileSystemNames returns a boolean if a field has been set.

### GetRegion

`func (o *RestartInstanceV1ResponseInner) GetRegion() RestartInstanceV1ResponseInnerRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RestartInstanceV1ResponseInner) GetRegionOk() (*RestartInstanceV1ResponseInnerRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RestartInstanceV1ResponseInner) SetRegion(v RestartInstanceV1ResponseInnerRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RestartInstanceV1ResponseInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetInstanceType

`func (o *RestartInstanceV1ResponseInner) GetInstanceType() RestartInstanceV1ResponseInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RestartInstanceV1ResponseInner) GetInstanceTypeOk() (*RestartInstanceV1ResponseInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RestartInstanceV1ResponseInner) SetInstanceType(v RestartInstanceV1ResponseInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *RestartInstanceV1ResponseInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetJupyterToken

`func (o *RestartInstanceV1ResponseInner) GetJupyterToken() string`

GetJupyterToken returns the JupyterToken field if non-nil, zero value otherwise.

### GetJupyterTokenOk

`func (o *RestartInstanceV1ResponseInner) GetJupyterTokenOk() (*string, bool)`

GetJupyterTokenOk returns a tuple with the JupyterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJupyterToken

`func (o *RestartInstanceV1ResponseInner) SetJupyterToken(v string)`

SetJupyterToken sets JupyterToken field to given value.

### HasJupyterToken

`func (o *RestartInstanceV1ResponseInner) HasJupyterToken() bool`

HasJupyterToken returns a boolean if a field has been set.

### GetJupyterUrl

`func (o *RestartInstanceV1ResponseInner) GetJupyterUrl() string`

GetJupyterUrl returns the JupyterUrl field if non-nil, zero value otherwise.

### GetJupyterUrlOk

`func (o *RestartInstanceV1ResponseInner) GetJupyterUrlOk() (*string, bool)`

GetJupyterUrlOk returns a tuple with the JupyterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJupyterUrl

`func (o *RestartInstanceV1ResponseInner) SetJupyterUrl(v string)`

SetJupyterUrl sets JupyterUrl field to given value.

### HasJupyterUrl

`func (o *RestartInstanceV1ResponseInner) HasJupyterUrl() bool`

HasJupyterUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


