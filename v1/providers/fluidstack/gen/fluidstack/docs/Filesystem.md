# Filesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the filesystem | 
**Name** | **string** | Name of the filesystem | 
**Size** | **string** | Size of the filesystem in GB | 

## Methods

### NewFilesystem

`func NewFilesystem(id string, name string, size string, ) *Filesystem`

NewFilesystem instantiates a new Filesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemWithDefaults

`func NewFilesystemWithDefaults() *Filesystem`

NewFilesystemWithDefaults instantiates a new Filesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Filesystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Filesystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Filesystem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Filesystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filesystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filesystem) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *Filesystem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Filesystem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Filesystem) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


