# GpuOsChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **int32** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** | Name of the OS to be provisioned onto GPU Node(s) (ex: ubuntu) | 
**Release** | **string** | Release name of the OS to be provisioned onto GPU Node(s) (ex: jammy) | 
**Version** | **string** | Version number of the OS to be provisioned onto GPU Node(s) (ex: 22.04) | 

## Methods

### NewGpuOsChoice

`func NewGpuOsChoice(created time.Time, id int32, modified time.Time, name string, release string, version string, ) *GpuOsChoice`

NewGpuOsChoice instantiates a new GpuOsChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuOsChoiceWithDefaults

`func NewGpuOsChoiceWithDefaults() *GpuOsChoice`

NewGpuOsChoiceWithDefaults instantiates a new GpuOsChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GpuOsChoice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GpuOsChoice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GpuOsChoice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *GpuOsChoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpuOsChoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpuOsChoice) SetId(v int32)`

SetId sets Id field to given value.


### GetModified

`func (o *GpuOsChoice) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GpuOsChoice) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GpuOsChoice) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *GpuOsChoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpuOsChoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpuOsChoice) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *GpuOsChoice) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *GpuOsChoice) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *GpuOsChoice) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetVersion

`func (o *GpuOsChoice) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GpuOsChoice) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GpuOsChoice) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


