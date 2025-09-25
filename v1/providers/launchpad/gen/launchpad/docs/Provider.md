# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**DisplayName** | Pointer to **NullableString** | Human-friendly version of name. Used for display purposes. | [optional] 
**Id** | **string** |  | [readonly] 
**InstanceLimit** | Pointer to **int32** | The maximum number of provisionined instances allowed globally for the provider (0 &#x3D; unlimited) | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** |  | 
**Priority** | Pointer to **int32** | Weighted preference to use in selecting a provider for a deployment. Higher priority values will be preferred over lower ones. | [optional] 

## Methods

### NewProvider

`func NewProvider(created time.Time, id string, modified time.Time, name string, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Provider) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Provider) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Provider) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDisplayName

`func (o *Provider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Provider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Provider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Provider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Provider) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Provider) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *Provider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceLimit

`func (o *Provider) GetInstanceLimit() int32`

GetInstanceLimit returns the InstanceLimit field if non-nil, zero value otherwise.

### GetInstanceLimitOk

`func (o *Provider) GetInstanceLimitOk() (*int32, bool)`

GetInstanceLimitOk returns a tuple with the InstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLimit

`func (o *Provider) SetInstanceLimit(v int32)`

SetInstanceLimit sets InstanceLimit field to given value.

### HasInstanceLimit

`func (o *Provider) HasInstanceLimit() bool`

HasInstanceLimit returns a boolean if a field has been set.

### GetModified

`func (o *Provider) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Provider) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Provider) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *Provider) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Provider) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Provider) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Provider) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


