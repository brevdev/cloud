# KubernetesNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the kubernetes node pool | 
**Name** | **string** | Name of the kubernetes node pool | 
**Cluster** | **string** | Name of the kubernetes cluster | 
**Type** | **string** | Instance type of the kubernetes node pool | 
**Count** | **int32** | Number of nodes in the kubernetes node pool | 
**Version** | **string** | Version of the kubernetes node pool | 
**Nodes** | Pointer to [**[]KubernetesNode**](KubernetesNode.md) | List of nodes in the kubernetes node pool | [optional] 

## Methods

### NewKubernetesNodePool

`func NewKubernetesNodePool(id string, name string, cluster string, type_ string, count int32, version string, ) *KubernetesNodePool`

NewKubernetesNodePool instantiates a new KubernetesNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolWithDefaults

`func NewKubernetesNodePoolWithDefaults() *KubernetesNodePool`

NewKubernetesNodePoolWithDefaults instantiates a new KubernetesNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesNodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodePool) SetName(v string)`

SetName sets Name field to given value.


### GetCluster

`func (o *KubernetesNodePool) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KubernetesNodePool) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KubernetesNodePool) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetType

`func (o *KubernetesNodePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodePool) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *KubernetesNodePool) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesNodePool) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesNodePool) SetCount(v int32)`

SetCount sets Count field to given value.


### GetVersion

`func (o *KubernetesNodePool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesNodePool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesNodePool) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetNodes

`func (o *KubernetesNodePool) GetNodes() []KubernetesNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesNodePool) GetNodesOk() (*[]KubernetesNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesNodePool) SetNodes(v []KubernetesNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *KubernetesNodePool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


