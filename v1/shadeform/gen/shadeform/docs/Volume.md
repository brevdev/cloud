# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the storage volume. | 
**Cloud** | [**Cloud**](Cloud.md) |  | 
**CloudAssignedId** | **string** | The unique identifier of the storage volume issued by the underlying cloud provider. | 
**Region** | **string** | Specifies the region. | 
**Name** | **string** | The name of the storage volume. | 
**FixedSize** | **bool** | Denotes whether the volume is fixed in size or elastically scaling. | 
**SizeInGb** | **int32** | Storage volume size in GB | 
**CostEstimate** | **string** | The cost incurred by the instance. This only the cost via Shadeform. If the instance is deployed in your own cloud account, then all billing is through your cloud provider. | 
**SupportsMultiMount** | **bool** | Denotes whether the volume supports multiple instances mounting to it at the same time. | 
**MountedBy** | **string** | The ID of the instance that is currently mounting the volume. | 

## Methods

### NewVolume

`func NewVolume(id string, cloud Cloud, cloudAssignedId string, region string, name string, fixedSize bool, sizeInGb int32, costEstimate string, supportsMultiMount bool, mountedBy string, ) *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.


### GetCloud

`func (o *Volume) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Volume) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Volume) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetCloudAssignedId

`func (o *Volume) GetCloudAssignedId() string`

GetCloudAssignedId returns the CloudAssignedId field if non-nil, zero value otherwise.

### GetCloudAssignedIdOk

`func (o *Volume) GetCloudAssignedIdOk() (*string, bool)`

GetCloudAssignedIdOk returns a tuple with the CloudAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssignedId

`func (o *Volume) SetCloudAssignedId(v string)`

SetCloudAssignedId sets CloudAssignedId field to given value.


### GetRegion

`func (o *Volume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Volume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Volume) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.


### GetFixedSize

`func (o *Volume) GetFixedSize() bool`

GetFixedSize returns the FixedSize field if non-nil, zero value otherwise.

### GetFixedSizeOk

`func (o *Volume) GetFixedSizeOk() (*bool, bool)`

GetFixedSizeOk returns a tuple with the FixedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSize

`func (o *Volume) SetFixedSize(v bool)`

SetFixedSize sets FixedSize field to given value.


### GetSizeInGb

`func (o *Volume) GetSizeInGb() int32`

GetSizeInGb returns the SizeInGb field if non-nil, zero value otherwise.

### GetSizeInGbOk

`func (o *Volume) GetSizeInGbOk() (*int32, bool)`

GetSizeInGbOk returns a tuple with the SizeInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInGb

`func (o *Volume) SetSizeInGb(v int32)`

SetSizeInGb sets SizeInGb field to given value.


### GetCostEstimate

`func (o *Volume) GetCostEstimate() string`

GetCostEstimate returns the CostEstimate field if non-nil, zero value otherwise.

### GetCostEstimateOk

`func (o *Volume) GetCostEstimateOk() (*string, bool)`

GetCostEstimateOk returns a tuple with the CostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimate

`func (o *Volume) SetCostEstimate(v string)`

SetCostEstimate sets CostEstimate field to given value.


### GetSupportsMultiMount

`func (o *Volume) GetSupportsMultiMount() bool`

GetSupportsMultiMount returns the SupportsMultiMount field if non-nil, zero value otherwise.

### GetSupportsMultiMountOk

`func (o *Volume) GetSupportsMultiMountOk() (*bool, bool)`

GetSupportsMultiMountOk returns a tuple with the SupportsMultiMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMultiMount

`func (o *Volume) SetSupportsMultiMount(v bool)`

SetSupportsMultiMount sets SupportsMultiMount field to given value.


### GetMountedBy

`func (o *Volume) GetMountedBy() string`

GetMountedBy returns the MountedBy field if non-nil, zero value otherwise.

### GetMountedByOk

`func (o *Volume) GetMountedByOk() (*string, bool)`

GetMountedByOk returns a tuple with the MountedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedBy

`func (o *Volume) SetMountedBy(v string)`

SetMountedBy sets MountedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


