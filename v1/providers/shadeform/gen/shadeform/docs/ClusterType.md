# ClusterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | The cloud provider. | 
**NumInstances** | **int32** | The number of instances in this cluster type. | 
**Availability** | [**ClusterAvailability**](ClusterAvailability.md) |  | 
**ClusterType** | [**InstanceType**](InstanceType.md) |  | 

## Methods

### NewClusterType

`func NewClusterType(cloud string, numInstances int32, availability ClusterAvailability, clusterType InstanceType, ) *ClusterType`

NewClusterType instantiates a new ClusterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTypeWithDefaults

`func NewClusterTypeWithDefaults() *ClusterType`

NewClusterTypeWithDefaults instantiates a new ClusterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *ClusterType) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ClusterType) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ClusterType) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetNumInstances

`func (o *ClusterType) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *ClusterType) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *ClusterType) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.


### GetAvailability

`func (o *ClusterType) GetAvailability() ClusterAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ClusterType) GetAvailabilityOk() (*ClusterAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ClusterType) SetAvailability(v ClusterAvailability)`

SetAvailability sets Availability field to given value.


### GetClusterType

`func (o *ClusterType) GetClusterType() InstanceType`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterType) GetClusterTypeOk() (*InstanceType, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterType) SetClusterType(v InstanceType)`

SetClusterType sets ClusterType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


