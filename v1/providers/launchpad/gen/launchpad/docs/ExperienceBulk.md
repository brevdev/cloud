# ExperienceBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsvFile** | **string** |  | 
**Experiences** | [**[]Experience**](Experience.md) |  | 

## Methods

### NewExperienceBulk

`func NewExperienceBulk(csvFile string, experiences []Experience, ) *ExperienceBulk`

NewExperienceBulk instantiates a new ExperienceBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceBulkWithDefaults

`func NewExperienceBulkWithDefaults() *ExperienceBulk`

NewExperienceBulkWithDefaults instantiates a new ExperienceBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsvFile

`func (o *ExperienceBulk) GetCsvFile() string`

GetCsvFile returns the CsvFile field if non-nil, zero value otherwise.

### GetCsvFileOk

`func (o *ExperienceBulk) GetCsvFileOk() (*string, bool)`

GetCsvFileOk returns a tuple with the CsvFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFile

`func (o *ExperienceBulk) SetCsvFile(v string)`

SetCsvFile sets CsvFile field to given value.


### GetExperiences

`func (o *ExperienceBulk) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *ExperienceBulk) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *ExperienceBulk) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


