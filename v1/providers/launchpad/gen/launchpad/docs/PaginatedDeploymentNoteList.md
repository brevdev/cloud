# PaginatedDeploymentNoteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | [**[]DeploymentNote**](DeploymentNote.md) |  | 

## Methods

### NewPaginatedDeploymentNoteList

`func NewPaginatedDeploymentNoteList(count int32, results []DeploymentNote, ) *PaginatedDeploymentNoteList`

NewPaginatedDeploymentNoteList instantiates a new PaginatedDeploymentNoteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDeploymentNoteListWithDefaults

`func NewPaginatedDeploymentNoteListWithDefaults() *PaginatedDeploymentNoteList`

NewPaginatedDeploymentNoteListWithDefaults instantiates a new PaginatedDeploymentNoteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedDeploymentNoteList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedDeploymentNoteList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedDeploymentNoteList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedDeploymentNoteList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDeploymentNoteList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDeploymentNoteList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDeploymentNoteList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedDeploymentNoteList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedDeploymentNoteList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedDeploymentNoteList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedDeploymentNoteList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedDeploymentNoteList) GetResults() []DeploymentNote`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDeploymentNoteList) GetResultsOk() (*[]DeploymentNote, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDeploymentNoteList) SetResults(v []DeploymentNote)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


