# SlurmClustersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The cluster name. | 
**Version** | Pointer to **string** | The cluster version. | [optional] 

## Methods

### NewSlurmClustersPostRequest

`func NewSlurmClustersPostRequest(name string, ) *SlurmClustersPostRequest`

NewSlurmClustersPostRequest instantiates a new SlurmClustersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmClustersPostRequestWithDefaults

`func NewSlurmClustersPostRequestWithDefaults() *SlurmClustersPostRequest`

NewSlurmClustersPostRequestWithDefaults instantiates a new SlurmClustersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlurmClustersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmClustersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmClustersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *SlurmClustersPostRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SlurmClustersPostRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SlurmClustersPostRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SlurmClustersPostRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


