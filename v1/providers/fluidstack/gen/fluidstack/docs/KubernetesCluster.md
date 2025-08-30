# KubernetesCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the cluster | 
**Name** | **string** | The cluster name. | 
**Version** | **string** | The cluster version. | 
**Endpoint** | Pointer to **string** | The API server endpoint of the kubernetes cluster | [optional] 

## Methods

### NewKubernetesCluster

`func NewKubernetesCluster(id string, name string, version string, ) *KubernetesCluster`

NewKubernetesCluster instantiates a new KubernetesCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterWithDefaults

`func NewKubernetesClusterWithDefaults() *KubernetesCluster`

NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesCluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCluster) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *KubernetesCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesCluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEndpoint

`func (o *KubernetesCluster) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *KubernetesCluster) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *KubernetesCluster) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *KubernetesCluster) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


