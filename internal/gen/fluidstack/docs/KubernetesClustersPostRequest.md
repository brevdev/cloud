# KubernetesClustersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The cluster name. | 
**Version** | Pointer to **string** | The cluster version. | [optional] 

## Methods

### NewKubernetesClustersPostRequest

`func NewKubernetesClustersPostRequest(name string, ) *KubernetesClustersPostRequest`

NewKubernetesClustersPostRequest instantiates a new KubernetesClustersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClustersPostRequestWithDefaults

`func NewKubernetesClustersPostRequestWithDefaults() *KubernetesClustersPostRequest`

NewKubernetesClustersPostRequestWithDefaults instantiates a new KubernetesClustersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClustersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClustersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClustersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *KubernetesClustersPostRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesClustersPostRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesClustersPostRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesClustersPostRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


