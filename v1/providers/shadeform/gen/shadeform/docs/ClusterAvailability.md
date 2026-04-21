# ClusterAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The region code. | 
**DisplayName** | **string** | Human-readable display name for the region. | 
**Available** | **bool** | Whether the cluster type is available in this region. | 

## Methods

### NewClusterAvailability

`func NewClusterAvailability(region string, displayName string, available bool, ) *ClusterAvailability`

NewClusterAvailability instantiates a new ClusterAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAvailabilityWithDefaults

`func NewClusterAvailabilityWithDefaults() *ClusterAvailability`

NewClusterAvailabilityWithDefaults instantiates a new ClusterAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *ClusterAvailability) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterAvailability) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterAvailability) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDisplayName

`func (o *ClusterAvailability) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterAvailability) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterAvailability) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAvailable

`func (o *ClusterAvailability) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ClusterAvailability) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ClusterAvailability) SetAvailable(v bool)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


