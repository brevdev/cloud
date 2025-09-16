# ModelChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | **string** |  | [readonly] 
**Changes** | [**[]ModelChangeChangesInner**](ModelChangeChangesInner.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **int32** |  | 
**Username** | **string** |  | 

## Methods

### NewModelChange

`func NewModelChange(changeType string, changes []ModelChangeChangesInner, created time.Time, id int32, username string, ) *ModelChange`

NewModelChange instantiates a new ModelChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelChangeWithDefaults

`func NewModelChangeWithDefaults() *ModelChange`

NewModelChangeWithDefaults instantiates a new ModelChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *ModelChange) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *ModelChange) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *ModelChange) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.


### GetChanges

`func (o *ModelChange) GetChanges() []ModelChangeChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ModelChange) GetChangesOk() (*[]ModelChangeChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ModelChange) SetChanges(v []ModelChangeChangesInner)`

SetChanges sets Changes field to given value.


### GetCreated

`func (o *ModelChange) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelChange) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelChange) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *ModelChange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelChange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelChange) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *ModelChange) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelChange) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelChange) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


