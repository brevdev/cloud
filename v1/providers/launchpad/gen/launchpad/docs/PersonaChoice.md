# PersonaChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **int32** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** |  | 

## Methods

### NewPersonaChoice

`func NewPersonaChoice(created time.Time, id int32, modified time.Time, name string, ) *PersonaChoice`

NewPersonaChoice instantiates a new PersonaChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonaChoiceWithDefaults

`func NewPersonaChoiceWithDefaults() *PersonaChoice`

NewPersonaChoiceWithDefaults instantiates a new PersonaChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PersonaChoice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PersonaChoice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PersonaChoice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *PersonaChoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonaChoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonaChoice) SetId(v int32)`

SetId sets Id field to given value.


### GetModified

`func (o *PersonaChoice) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PersonaChoice) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PersonaChoice) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *PersonaChoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonaChoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonaChoice) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


