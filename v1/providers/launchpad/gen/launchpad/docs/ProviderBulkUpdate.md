# ProviderBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**DisplayName** | Pointer to **NullableString** | Human-friendly version of name. Used for display purposes. | [optional] 
**Id** | **string** |  | [readonly] 
**InstanceLimit** | Pointer to **int32** | The maximum number of provisionined instances allowed globally for the provider (0 &#x3D; unlimited) | [optional] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** |  | [readonly] 
**Priority** | Pointer to **int32** | Weighted preference to use in selecting a provider for a deployment. Higher priority values will be preferred over lower ones. | [optional] 
**Count** | **int32** |  | [readonly] 
**Ids** | **[]string** |  | 
**Result** | **string** |  | [readonly] 

## Methods

### NewProviderBulkUpdate

`func NewProviderBulkUpdate(created time.Time, id string, modified time.Time, name string, count int32, ids []string, result string, ) *ProviderBulkUpdate`

NewProviderBulkUpdate instantiates a new ProviderBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderBulkUpdateWithDefaults

`func NewProviderBulkUpdateWithDefaults() *ProviderBulkUpdate`

NewProviderBulkUpdateWithDefaults instantiates a new ProviderBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ProviderBulkUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProviderBulkUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProviderBulkUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDisplayName

`func (o *ProviderBulkUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProviderBulkUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProviderBulkUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProviderBulkUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ProviderBulkUpdate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ProviderBulkUpdate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *ProviderBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderBulkUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceLimit

`func (o *ProviderBulkUpdate) GetInstanceLimit() int32`

GetInstanceLimit returns the InstanceLimit field if non-nil, zero value otherwise.

### GetInstanceLimitOk

`func (o *ProviderBulkUpdate) GetInstanceLimitOk() (*int32, bool)`

GetInstanceLimitOk returns a tuple with the InstanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLimit

`func (o *ProviderBulkUpdate) SetInstanceLimit(v int32)`

SetInstanceLimit sets InstanceLimit field to given value.

### HasInstanceLimit

`func (o *ProviderBulkUpdate) HasInstanceLimit() bool`

HasInstanceLimit returns a boolean if a field has been set.

### GetModified

`func (o *ProviderBulkUpdate) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ProviderBulkUpdate) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ProviderBulkUpdate) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *ProviderBulkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderBulkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderBulkUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *ProviderBulkUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProviderBulkUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProviderBulkUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProviderBulkUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCount

`func (o *ProviderBulkUpdate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProviderBulkUpdate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProviderBulkUpdate) SetCount(v int32)`

SetCount sets Count field to given value.


### GetIds

`func (o *ProviderBulkUpdate) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ProviderBulkUpdate) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ProviderBulkUpdate) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetResult

`func (o *ProviderBulkUpdate) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProviderBulkUpdate) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProviderBulkUpdate) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


