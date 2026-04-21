# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the cluster. | 
**Cloud** | **string** | The cloud provider for the cluster. | 
**Region** | **string** | The region where the cluster will be deployed. | 
**ClusterType** | **string** | The type of GPU cluster to create. | 
**NumInstances** | **int32** | The number of instances in the cluster. | 
**SshKeyId** | Pointer to **string** | The SSH key ID to use for the cluster instances. | [optional] 
**Os** | Pointer to **string** | The operating system for the cluster. | [optional] 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest(name string, cloud string, region string, clusterType string, numInstances int32, ) *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCloud

`func (o *CreateClusterRequest) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CreateClusterRequest) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CreateClusterRequest) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *CreateClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetClusterType

`func (o *CreateClusterRequest) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *CreateClusterRequest) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *CreateClusterRequest) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetNumInstances

`func (o *CreateClusterRequest) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *CreateClusterRequest) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *CreateClusterRequest) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.


### GetSshKeyId

`func (o *CreateClusterRequest) GetSshKeyId() string`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *CreateClusterRequest) GetSshKeyIdOk() (*string, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *CreateClusterRequest) SetSshKeyId(v string)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *CreateClusterRequest) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetOs

`func (o *CreateClusterRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CreateClusterRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CreateClusterRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CreateClusterRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


