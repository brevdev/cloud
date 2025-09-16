# PaginatedPersonaChoiceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | [**[]PersonaChoice**](PersonaChoice.md) |  | 

## Methods

### NewPaginatedPersonaChoiceList

`func NewPaginatedPersonaChoiceList(count int32, results []PersonaChoice, ) *PaginatedPersonaChoiceList`

NewPaginatedPersonaChoiceList instantiates a new PaginatedPersonaChoiceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPersonaChoiceListWithDefaults

`func NewPaginatedPersonaChoiceListWithDefaults() *PaginatedPersonaChoiceList`

NewPaginatedPersonaChoiceListWithDefaults instantiates a new PaginatedPersonaChoiceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedPersonaChoiceList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedPersonaChoiceList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedPersonaChoiceList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedPersonaChoiceList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedPersonaChoiceList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedPersonaChoiceList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedPersonaChoiceList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedPersonaChoiceList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedPersonaChoiceList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedPersonaChoiceList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedPersonaChoiceList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedPersonaChoiceList) GetResults() []PersonaChoice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPersonaChoiceList) GetResultsOk() (*[]PersonaChoice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPersonaChoiceList) SetResults(v []PersonaChoice)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


