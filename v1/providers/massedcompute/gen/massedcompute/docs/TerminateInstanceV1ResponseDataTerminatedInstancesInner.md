# TerminateInstanceV1ResponseDataTerminatedInstancesInner

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

### NewTerminateInstanceV1ResponseDataTerminatedInstancesInner

`func NewTerminateInstanceV1ResponseDataTerminatedInstancesInner() *TerminateInstanceV1ResponseDataTerminatedInstancesInner`

NewTerminateInstanceV1ResponseDataTerminatedInstancesInner instantiates a new TerminateInstanceV1ResponseDataTerminatedInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateInstanceV1ResponseDataTerminatedInstancesInnerWithDefaults

`func NewTerminateInstanceV1ResponseDataTerminatedInstancesInnerWithDefaults() *TerminateInstanceV1ResponseDataTerminatedInstancesInner`

NewTerminateInstanceV1ResponseDataTerminatedInstancesInnerWithDefaults instantiates a new TerminateInstanceV1ResponseDataTerminatedInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetStatus

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSshKeyNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetSshKeyNames() []string`

GetSshKeyNames returns the SshKeyNames field if non-nil, zero value otherwise.

### GetSshKeyNamesOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetSshKeyNamesOk() (*[]string, bool)`

GetSshKeyNamesOk returns a tuple with the SshKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetSshKeyNames(v []string)`

SetSshKeyNames sets SshKeyNames field to given value.

### HasSshKeyNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasSshKeyNames() bool`

HasSshKeyNames returns a boolean if a field has been set.

### GetFileSystemNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetFileSystemNames() []string`

GetFileSystemNames returns the FileSystemNames field if non-nil, zero value otherwise.

### GetFileSystemNamesOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetFileSystemNamesOk() (*[]string, bool)`

GetFileSystemNamesOk returns a tuple with the FileSystemNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetFileSystemNames(v []string)`

SetFileSystemNames sets FileSystemNames field to given value.

### HasFileSystemNames

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasFileSystemNames() bool`

HasFileSystemNames returns a boolean if a field has been set.

### GetRegion

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetRegion() RestartInstanceV1ResponseInnerRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetRegionOk() (*RestartInstanceV1ResponseInnerRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetRegion(v RestartInstanceV1ResponseInnerRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetInstanceType

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetInstanceType() RestartInstanceV1ResponseInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetInstanceTypeOk() (*RestartInstanceV1ResponseInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetInstanceType(v RestartInstanceV1ResponseInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetJupyterToken

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetJupyterToken() string`

GetJupyterToken returns the JupyterToken field if non-nil, zero value otherwise.

### GetJupyterTokenOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetJupyterTokenOk() (*string, bool)`

GetJupyterTokenOk returns a tuple with the JupyterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJupyterToken

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetJupyterToken(v string)`

SetJupyterToken sets JupyterToken field to given value.

### HasJupyterToken

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasJupyterToken() bool`

HasJupyterToken returns a boolean if a field has been set.

### GetJupyterUrl

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetJupyterUrl() string`

GetJupyterUrl returns the JupyterUrl field if non-nil, zero value otherwise.

### GetJupyterUrlOk

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) GetJupyterUrlOk() (*string, bool)`

GetJupyterUrlOk returns a tuple with the JupyterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJupyterUrl

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) SetJupyterUrl(v string)`

SetJupyterUrl sets JupyterUrl field to given value.

### HasJupyterUrl

`func (o *TerminateInstanceV1ResponseDataTerminatedInstancesInner) HasJupyterUrl() bool`

HasJupyterUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


