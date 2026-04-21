# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the cluster. | 
**Cloud** | **string** | The cloud provider for the cluster. | 
**Name** | **string** | The name of the cluster. | 
**CloudClusterId** | Pointer to **NullableString** | The cloud provider assigned cluster ID. | [optional] 
**RegionInfo** | [**ClusterRegionInfo**](ClusterRegionInfo.md) |  | 
**Status** | **string** | The current status of the cluster. | 
**StatusDetails** | Pointer to **NullableString** | Additional details about the cluster status. | [optional] 
**CreatedAt** | **time.Time** | The timestamp when the cluster was created. | 
**UpdatedAt** | **time.Time** | The timestamp when the cluster was last updated. | 
**Instances** | [**[]Instance**](Instance.md) | Array of instances in the cluster. | 
**HourlyPrice** | **int32** | The hourly price of the cluster in cents. | 
**CostEstimate** | Pointer to **NullableString** | Estimated cost of the cluster. | [optional] 
**ActiveAt** | Pointer to **NullableTime** | The timestamp when the cluster became active. | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, cloud string, name string, regionInfo ClusterRegionInfo, status string, createdAt time.Time, updatedAt time.Time, instances []Instance, hourlyPrice int32, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetCloud

`func (o *Cluster) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Cluster) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Cluster) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetCloudClusterId

`func (o *Cluster) GetCloudClusterId() string`

GetCloudClusterId returns the CloudClusterId field if non-nil, zero value otherwise.

### GetCloudClusterIdOk

`func (o *Cluster) GetCloudClusterIdOk() (*string, bool)`

GetCloudClusterIdOk returns a tuple with the CloudClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudClusterId

`func (o *Cluster) SetCloudClusterId(v string)`

SetCloudClusterId sets CloudClusterId field to given value.

### HasCloudClusterId

`func (o *Cluster) HasCloudClusterId() bool`

HasCloudClusterId returns a boolean if a field has been set.

### SetCloudClusterIdNil

`func (o *Cluster) SetCloudClusterIdNil(b bool)`

 SetCloudClusterIdNil sets the value for CloudClusterId to be an explicit nil

### UnsetCloudClusterId
`func (o *Cluster) UnsetCloudClusterId()`

UnsetCloudClusterId ensures that no value is present for CloudClusterId, not even an explicit nil
### GetRegionInfo

`func (o *Cluster) GetRegionInfo() ClusterRegionInfo`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *Cluster) GetRegionInfoOk() (*ClusterRegionInfo, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *Cluster) SetRegionInfo(v ClusterRegionInfo)`

SetRegionInfo sets RegionInfo field to given value.


### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *Cluster) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Cluster) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Cluster) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Cluster) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *Cluster) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *Cluster) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInstances

`func (o *Cluster) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Cluster) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Cluster) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.


### GetHourlyPrice

`func (o *Cluster) GetHourlyPrice() int32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *Cluster) GetHourlyPriceOk() (*int32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *Cluster) SetHourlyPrice(v int32)`

SetHourlyPrice sets HourlyPrice field to given value.


### GetCostEstimate

`func (o *Cluster) GetCostEstimate() string`

GetCostEstimate returns the CostEstimate field if non-nil, zero value otherwise.

### GetCostEstimateOk

`func (o *Cluster) GetCostEstimateOk() (*string, bool)`

GetCostEstimateOk returns a tuple with the CostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimate

`func (o *Cluster) SetCostEstimate(v string)`

SetCostEstimate sets CostEstimate field to given value.

### HasCostEstimate

`func (o *Cluster) HasCostEstimate() bool`

HasCostEstimate returns a boolean if a field has been set.

### SetCostEstimateNil

`func (o *Cluster) SetCostEstimateNil(b bool)`

 SetCostEstimateNil sets the value for CostEstimate to be an explicit nil

### UnsetCostEstimate
`func (o *Cluster) UnsetCostEstimate()`

UnsetCostEstimate ensures that no value is present for CostEstimate, not even an explicit nil
### GetActiveAt

`func (o *Cluster) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *Cluster) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *Cluster) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *Cluster) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### SetActiveAtNil

`func (o *Cluster) SetActiveAtNil(b bool)`

 SetActiveAtNil sets the value for ActiveAt to be an explicit nil

### UnsetActiveAt
`func (o *Cluster) UnsetActiveAt()`

UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


