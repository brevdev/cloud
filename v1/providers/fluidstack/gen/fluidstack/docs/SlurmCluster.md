# SlurmCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the slurm cluster | 
**Name** | **string** | Name of the slurm cluster | 
**Version** | **string** | The cluster version. | 
**LoginNodes** | Pointer to [**[]SlurmLoginNode**](SlurmLoginNode.md) | Login nodes for the cluster | [optional] 

## Methods

### NewSlurmCluster

`func NewSlurmCluster(id string, name string, version string, ) *SlurmCluster`

NewSlurmCluster instantiates a new SlurmCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmClusterWithDefaults

`func NewSlurmClusterWithDefaults() *SlurmCluster`

NewSlurmClusterWithDefaults instantiates a new SlurmCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlurmCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlurmCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlurmCluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SlurmCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmCluster) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *SlurmCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SlurmCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SlurmCluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLoginNodes

`func (o *SlurmCluster) GetLoginNodes() []SlurmLoginNode`

GetLoginNodes returns the LoginNodes field if non-nil, zero value otherwise.

### GetLoginNodesOk

`func (o *SlurmCluster) GetLoginNodesOk() (*[]SlurmLoginNode, bool)`

GetLoginNodesOk returns a tuple with the LoginNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginNodes

`func (o *SlurmCluster) SetLoginNodes(v []SlurmLoginNode)`

SetLoginNodes sets LoginNodes field to given value.

### HasLoginNodes

`func (o *SlurmCluster) HasLoginNodes() bool`

HasLoginNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


