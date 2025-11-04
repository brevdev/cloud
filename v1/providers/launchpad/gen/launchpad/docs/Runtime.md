# Runtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Runtime repository branch | [optional] 
**CnsAddonPack** | Pointer to **bool** | Include CNS add-ons? | [optional] 
**CnsDocker** | Pointer to **bool** | Include Docker with CNS? | [optional] 
**CnsDriverVersion** | Pointer to **NullableString** | GPU driver version | [optional] 
**CnsK8s** | Pointer to **bool** | Include Kubernetes with CNS? | [optional] 
**CnsNvidiaDriver** | Pointer to **bool** | Include NVIDIA driver with CNS? | [optional] 
**CnsVersion** | Pointer to **NullableString** | NVIDIA Cloud Native Stack version | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**Mig** | Pointer to **bool** | Include MIG support with CNS? | [optional] 
**MigProfile** | Pointer to **NullableString** | MIG profile name | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** | Human-readable name of the runtime | 
**Url** | **string** | URL of the runtime repository | 

## Methods

### NewRuntime

`func NewRuntime(created time.Time, id string, modified time.Time, name string, url string, ) *Runtime`

NewRuntime instantiates a new Runtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeWithDefaults

`func NewRuntimeWithDefaults() *Runtime`

NewRuntimeWithDefaults instantiates a new Runtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *Runtime) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Runtime) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Runtime) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Runtime) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCnsAddonPack

`func (o *Runtime) GetCnsAddonPack() bool`

GetCnsAddonPack returns the CnsAddonPack field if non-nil, zero value otherwise.

### GetCnsAddonPackOk

`func (o *Runtime) GetCnsAddonPackOk() (*bool, bool)`

GetCnsAddonPackOk returns a tuple with the CnsAddonPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsAddonPack

`func (o *Runtime) SetCnsAddonPack(v bool)`

SetCnsAddonPack sets CnsAddonPack field to given value.

### HasCnsAddonPack

`func (o *Runtime) HasCnsAddonPack() bool`

HasCnsAddonPack returns a boolean if a field has been set.

### GetCnsDocker

`func (o *Runtime) GetCnsDocker() bool`

GetCnsDocker returns the CnsDocker field if non-nil, zero value otherwise.

### GetCnsDockerOk

`func (o *Runtime) GetCnsDockerOk() (*bool, bool)`

GetCnsDockerOk returns a tuple with the CnsDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsDocker

`func (o *Runtime) SetCnsDocker(v bool)`

SetCnsDocker sets CnsDocker field to given value.

### HasCnsDocker

`func (o *Runtime) HasCnsDocker() bool`

HasCnsDocker returns a boolean if a field has been set.

### GetCnsDriverVersion

`func (o *Runtime) GetCnsDriverVersion() string`

GetCnsDriverVersion returns the CnsDriverVersion field if non-nil, zero value otherwise.

### GetCnsDriverVersionOk

`func (o *Runtime) GetCnsDriverVersionOk() (*string, bool)`

GetCnsDriverVersionOk returns a tuple with the CnsDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsDriverVersion

`func (o *Runtime) SetCnsDriverVersion(v string)`

SetCnsDriverVersion sets CnsDriverVersion field to given value.

### HasCnsDriverVersion

`func (o *Runtime) HasCnsDriverVersion() bool`

HasCnsDriverVersion returns a boolean if a field has been set.

### SetCnsDriverVersionNil

`func (o *Runtime) SetCnsDriverVersionNil(b bool)`

 SetCnsDriverVersionNil sets the value for CnsDriverVersion to be an explicit nil

### UnsetCnsDriverVersion
`func (o *Runtime) UnsetCnsDriverVersion()`

UnsetCnsDriverVersion ensures that no value is present for CnsDriverVersion, not even an explicit nil
### GetCnsK8s

`func (o *Runtime) GetCnsK8s() bool`

GetCnsK8s returns the CnsK8s field if non-nil, zero value otherwise.

### GetCnsK8sOk

`func (o *Runtime) GetCnsK8sOk() (*bool, bool)`

GetCnsK8sOk returns a tuple with the CnsK8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsK8s

`func (o *Runtime) SetCnsK8s(v bool)`

SetCnsK8s sets CnsK8s field to given value.

### HasCnsK8s

`func (o *Runtime) HasCnsK8s() bool`

HasCnsK8s returns a boolean if a field has been set.

### GetCnsNvidiaDriver

`func (o *Runtime) GetCnsNvidiaDriver() bool`

GetCnsNvidiaDriver returns the CnsNvidiaDriver field if non-nil, zero value otherwise.

### GetCnsNvidiaDriverOk

`func (o *Runtime) GetCnsNvidiaDriverOk() (*bool, bool)`

GetCnsNvidiaDriverOk returns a tuple with the CnsNvidiaDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsNvidiaDriver

`func (o *Runtime) SetCnsNvidiaDriver(v bool)`

SetCnsNvidiaDriver sets CnsNvidiaDriver field to given value.

### HasCnsNvidiaDriver

`func (o *Runtime) HasCnsNvidiaDriver() bool`

HasCnsNvidiaDriver returns a boolean if a field has been set.

### GetCnsVersion

`func (o *Runtime) GetCnsVersion() string`

GetCnsVersion returns the CnsVersion field if non-nil, zero value otherwise.

### GetCnsVersionOk

`func (o *Runtime) GetCnsVersionOk() (*string, bool)`

GetCnsVersionOk returns a tuple with the CnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsVersion

`func (o *Runtime) SetCnsVersion(v string)`

SetCnsVersion sets CnsVersion field to given value.

### HasCnsVersion

`func (o *Runtime) HasCnsVersion() bool`

HasCnsVersion returns a boolean if a field has been set.

### SetCnsVersionNil

`func (o *Runtime) SetCnsVersionNil(b bool)`

 SetCnsVersionNil sets the value for CnsVersion to be an explicit nil

### UnsetCnsVersion
`func (o *Runtime) UnsetCnsVersion()`

UnsetCnsVersion ensures that no value is present for CnsVersion, not even an explicit nil
### GetCreated

`func (o *Runtime) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Runtime) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Runtime) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *Runtime) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Runtime) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Runtime) SetId(v string)`

SetId sets Id field to given value.


### GetMig

`func (o *Runtime) GetMig() bool`

GetMig returns the Mig field if non-nil, zero value otherwise.

### GetMigOk

`func (o *Runtime) GetMigOk() (*bool, bool)`

GetMigOk returns a tuple with the Mig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMig

`func (o *Runtime) SetMig(v bool)`

SetMig sets Mig field to given value.

### HasMig

`func (o *Runtime) HasMig() bool`

HasMig returns a boolean if a field has been set.

### GetMigProfile

`func (o *Runtime) GetMigProfile() string`

GetMigProfile returns the MigProfile field if non-nil, zero value otherwise.

### GetMigProfileOk

`func (o *Runtime) GetMigProfileOk() (*string, bool)`

GetMigProfileOk returns a tuple with the MigProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigProfile

`func (o *Runtime) SetMigProfile(v string)`

SetMigProfile sets MigProfile field to given value.

### HasMigProfile

`func (o *Runtime) HasMigProfile() bool`

HasMigProfile returns a boolean if a field has been set.

### SetMigProfileNil

`func (o *Runtime) SetMigProfileNil(b bool)`

 SetMigProfileNil sets the value for MigProfile to be an explicit nil

### UnsetMigProfile
`func (o *Runtime) UnsetMigProfile()`

UnsetMigProfile ensures that no value is present for MigProfile, not even an explicit nil
### GetModified

`func (o *Runtime) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Runtime) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Runtime) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *Runtime) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Runtime) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Runtime) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Runtime) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Runtime) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Runtime) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


