# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Groups** | **[]string** |  | 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**PreferredTheme** | Pointer to [**NullablePreferredThemeEnum**](PreferredThemeEnum.md) |  | [optional] 
**PreferredViewSettings** | Pointer to **interface{}** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewProfile

`func NewProfile(created time.Time, groups []string, id string, modified time.Time, username string, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Profile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Profile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Profile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetGroups

`func (o *Profile) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Profile) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Profile) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *Profile) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Profile) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Profile) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPreferredTheme

`func (o *Profile) GetPreferredTheme() PreferredThemeEnum`

GetPreferredTheme returns the PreferredTheme field if non-nil, zero value otherwise.

### GetPreferredThemeOk

`func (o *Profile) GetPreferredThemeOk() (*PreferredThemeEnum, bool)`

GetPreferredThemeOk returns a tuple with the PreferredTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTheme

`func (o *Profile) SetPreferredTheme(v PreferredThemeEnum)`

SetPreferredTheme sets PreferredTheme field to given value.

### HasPreferredTheme

`func (o *Profile) HasPreferredTheme() bool`

HasPreferredTheme returns a boolean if a field has been set.

### SetPreferredThemeNil

`func (o *Profile) SetPreferredThemeNil(b bool)`

 SetPreferredThemeNil sets the value for PreferredTheme to be an explicit nil

### UnsetPreferredTheme
`func (o *Profile) UnsetPreferredTheme()`

UnsetPreferredTheme ensures that no value is present for PreferredTheme, not even an explicit nil
### GetPreferredViewSettings

`func (o *Profile) GetPreferredViewSettings() interface{}`

GetPreferredViewSettings returns the PreferredViewSettings field if non-nil, zero value otherwise.

### GetPreferredViewSettingsOk

`func (o *Profile) GetPreferredViewSettingsOk() (*interface{}, bool)`

GetPreferredViewSettingsOk returns a tuple with the PreferredViewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredViewSettings

`func (o *Profile) SetPreferredViewSettings(v interface{})`

SetPreferredViewSettings sets PreferredViewSettings field to given value.

### HasPreferredViewSettings

`func (o *Profile) HasPreferredViewSettings() bool`

HasPreferredViewSettings returns a boolean if a field has been set.

### SetPreferredViewSettingsNil

`func (o *Profile) SetPreferredViewSettingsNil(b bool)`

 SetPreferredViewSettingsNil sets the value for PreferredViewSettings to be an explicit nil

### UnsetPreferredViewSettings
`func (o *Profile) UnsetPreferredViewSettings()`

UnsetPreferredViewSettings ensures that no value is present for PreferredViewSettings, not even an explicit nil
### GetUsername

`func (o *Profile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Profile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Profile) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


