# KubernetesNodePoolPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the kubernetes node pool | 
**Type** | **string** | Instance type of the kubernetes node pool | 
**Count** | **int32** | Number of nodes in the kubernetes node pool | 
**Version** | Pointer to **string** | Version of the kubernetes node pool | [optional] 

## Methods

### NewKubernetesNodePoolPostRequest

`func NewKubernetesNodePoolPostRequest(name string, type_ string, count int32, ) *KubernetesNodePoolPostRequest`

NewKubernetesNodePoolPostRequest instantiates a new KubernetesNodePoolPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolPostRequestWithDefaults

`func NewKubernetesNodePoolPostRequestWithDefaults() *KubernetesNodePoolPostRequest`

NewKubernetesNodePoolPostRequestWithDefaults instantiates a new KubernetesNodePoolPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNodePoolPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodePoolPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodePoolPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *KubernetesNodePoolPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodePoolPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodePoolPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *KubernetesNodePoolPostRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesNodePoolPostRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesNodePoolPostRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetVersion

`func (o *KubernetesNodePoolPostRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesNodePoolPostRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesNodePoolPostRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesNodePoolPostRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


