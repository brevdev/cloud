# ClusterInfoResponse

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

### NewClusterInfoResponse

`func NewClusterInfoResponse(id string, cloud string, name string, regionInfo ClusterRegionInfo, status string, createdAt time.Time, updatedAt time.Time, instances []Instance, hourlyPrice int32, ) *ClusterInfoResponse`

NewClusterInfoResponse instantiates a new ClusterInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoResponseWithDefaults

`func NewClusterInfoResponseWithDefaults() *ClusterInfoResponse`

NewClusterInfoResponseWithDefaults instantiates a new ClusterInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCloud

`func (o *ClusterInfoResponse) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ClusterInfoResponse) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ClusterInfoResponse) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetName

`func (o *ClusterInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInfoResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCloudClusterId

`func (o *ClusterInfoResponse) GetCloudClusterId() string`

GetCloudClusterId returns the CloudClusterId field if non-nil, zero value otherwise.

### GetCloudClusterIdOk

`func (o *ClusterInfoResponse) GetCloudClusterIdOk() (*string, bool)`

GetCloudClusterIdOk returns a tuple with the CloudClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudClusterId

`func (o *ClusterInfoResponse) SetCloudClusterId(v string)`

SetCloudClusterId sets CloudClusterId field to given value.

### HasCloudClusterId

`func (o *ClusterInfoResponse) HasCloudClusterId() bool`

HasCloudClusterId returns a boolean if a field has been set.

### SetCloudClusterIdNil

`func (o *ClusterInfoResponse) SetCloudClusterIdNil(b bool)`

 SetCloudClusterIdNil sets the value for CloudClusterId to be an explicit nil

### UnsetCloudClusterId
`func (o *ClusterInfoResponse) UnsetCloudClusterId()`

UnsetCloudClusterId ensures that no value is present for CloudClusterId, not even an explicit nil
### GetRegionInfo

`func (o *ClusterInfoResponse) GetRegionInfo() ClusterRegionInfo`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *ClusterInfoResponse) GetRegionInfoOk() (*ClusterRegionInfo, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *ClusterInfoResponse) SetRegionInfo(v ClusterRegionInfo)`

SetRegionInfo sets RegionInfo field to given value.


### GetStatus

`func (o *ClusterInfoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterInfoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterInfoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *ClusterInfoResponse) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *ClusterInfoResponse) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *ClusterInfoResponse) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *ClusterInfoResponse) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *ClusterInfoResponse) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *ClusterInfoResponse) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedAt

`func (o *ClusterInfoResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterInfoResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterInfoResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ClusterInfoResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterInfoResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterInfoResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInstances

`func (o *ClusterInfoResponse) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterInfoResponse) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterInfoResponse) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.


### GetHourlyPrice

`func (o *ClusterInfoResponse) GetHourlyPrice() int32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ClusterInfoResponse) GetHourlyPriceOk() (*int32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ClusterInfoResponse) SetHourlyPrice(v int32)`

SetHourlyPrice sets HourlyPrice field to given value.


### GetCostEstimate

`func (o *ClusterInfoResponse) GetCostEstimate() string`

GetCostEstimate returns the CostEstimate field if non-nil, zero value otherwise.

### GetCostEstimateOk

`func (o *ClusterInfoResponse) GetCostEstimateOk() (*string, bool)`

GetCostEstimateOk returns a tuple with the CostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimate

`func (o *ClusterInfoResponse) SetCostEstimate(v string)`

SetCostEstimate sets CostEstimate field to given value.

### HasCostEstimate

`func (o *ClusterInfoResponse) HasCostEstimate() bool`

HasCostEstimate returns a boolean if a field has been set.

### SetCostEstimateNil

`func (o *ClusterInfoResponse) SetCostEstimateNil(b bool)`

 SetCostEstimateNil sets the value for CostEstimate to be an explicit nil

### UnsetCostEstimate
`func (o *ClusterInfoResponse) UnsetCostEstimate()`

UnsetCostEstimate ensures that no value is present for CostEstimate, not even an explicit nil
### GetActiveAt

`func (o *ClusterInfoResponse) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *ClusterInfoResponse) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *ClusterInfoResponse) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *ClusterInfoResponse) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### SetActiveAtNil

`func (o *ClusterInfoResponse) SetActiveAtNil(b bool)`

 SetActiveAtNil sets the value for ActiveAt to be an explicit nil

### UnsetActiveAt
`func (o *ClusterInfoResponse) UnsetActiveAt()`

UnsetActiveAt ensures that no value is present for ActiveAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


