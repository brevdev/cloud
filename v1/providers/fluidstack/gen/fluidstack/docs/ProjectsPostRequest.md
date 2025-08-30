# ProjectsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the project | 
**Tags** | Pointer to **map[string]string** | Tags for the project. Tag keys must be 1–63 characters, start and end with alphanumerics, and may include &#x60;-&#x60;, &#x60;_&#x60;, and &#x60;.&#x60;.  | [optional] 

## Methods

### NewProjectsPostRequest

`func NewProjectsPostRequest(name string, ) *ProjectsPostRequest`

NewProjectsPostRequest instantiates a new ProjectsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsPostRequestWithDefaults

`func NewProjectsPostRequestWithDefaults() *ProjectsPostRequest`

NewProjectsPostRequestWithDefaults instantiates a new ProjectsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ProjectsPostRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectsPostRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectsPostRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectsPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


