# ExperienceNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Text content of the note | [optional] 
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**CreatedBy** | **string** |  | [readonly] 
**Experience** | [**ExperienceNoteExperience**](ExperienceNoteExperience.md) |  | 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**ModifiedBy** | **string** |  | [readonly] 

## Methods

### NewExperienceNote

`func NewExperienceNote(created time.Time, createdBy string, experience ExperienceNoteExperience, id string, modified time.Time, modifiedBy string, ) *ExperienceNote`

NewExperienceNote instantiates a new ExperienceNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceNoteWithDefaults

`func NewExperienceNoteWithDefaults() *ExperienceNote`

NewExperienceNoteWithDefaults instantiates a new ExperienceNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ExperienceNote) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExperienceNote) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExperienceNote) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExperienceNote) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *ExperienceNote) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExperienceNote) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExperienceNote) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *ExperienceNote) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExperienceNote) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExperienceNote) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetExperience

`func (o *ExperienceNote) GetExperience() ExperienceNoteExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ExperienceNote) GetExperienceOk() (*ExperienceNoteExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ExperienceNote) SetExperience(v ExperienceNoteExperience)`

SetExperience sets Experience field to given value.


### GetId

`func (o *ExperienceNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperienceNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperienceNote) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *ExperienceNote) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ExperienceNote) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ExperienceNote) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetModifiedBy

`func (o *ExperienceNote) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ExperienceNote) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ExperienceNote) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


