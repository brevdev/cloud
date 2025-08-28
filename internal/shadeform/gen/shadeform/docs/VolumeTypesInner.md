# VolumeTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | [**Cloud**](Cloud.md) |  | 
**Region** | **string** | Specifies the region. | 
**SupportsMultiMount** | **bool** | Denotes whether the volume supports multiple instances mounting to it at the same time. | 
**FixedSize** | **bool** | Denotes whether the volume is fixed in size or elastically scaling. | 
**PricePerGbPerHour** | **string** | Price of the specified volume type per GB per hour | 

## Methods

### NewVolumeTypesInner

`func NewVolumeTypesInner(cloud Cloud, region string, supportsMultiMount bool, fixedSize bool, pricePerGbPerHour string, ) *VolumeTypesInner`

NewVolumeTypesInner instantiates a new VolumeTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTypesInnerWithDefaults

`func NewVolumeTypesInnerWithDefaults() *VolumeTypesInner`

NewVolumeTypesInnerWithDefaults instantiates a new VolumeTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *VolumeTypesInner) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VolumeTypesInner) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VolumeTypesInner) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *VolumeTypesInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VolumeTypesInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VolumeTypesInner) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSupportsMultiMount

`func (o *VolumeTypesInner) GetSupportsMultiMount() bool`

GetSupportsMultiMount returns the SupportsMultiMount field if non-nil, zero value otherwise.

### GetSupportsMultiMountOk

`func (o *VolumeTypesInner) GetSupportsMultiMountOk() (*bool, bool)`

GetSupportsMultiMountOk returns a tuple with the SupportsMultiMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMultiMount

`func (o *VolumeTypesInner) SetSupportsMultiMount(v bool)`

SetSupportsMultiMount sets SupportsMultiMount field to given value.


### GetFixedSize

`func (o *VolumeTypesInner) GetFixedSize() bool`

GetFixedSize returns the FixedSize field if non-nil, zero value otherwise.

### GetFixedSizeOk

`func (o *VolumeTypesInner) GetFixedSizeOk() (*bool, bool)`

GetFixedSizeOk returns a tuple with the FixedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSize

`func (o *VolumeTypesInner) SetFixedSize(v bool)`

SetFixedSize sets FixedSize field to given value.


### GetPricePerGbPerHour

`func (o *VolumeTypesInner) GetPricePerGbPerHour() string`

GetPricePerGbPerHour returns the PricePerGbPerHour field if non-nil, zero value otherwise.

### GetPricePerGbPerHourOk

`func (o *VolumeTypesInner) GetPricePerGbPerHourOk() (*string, bool)`

GetPricePerGbPerHourOk returns a tuple with the PricePerGbPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerGbPerHour

`func (o *VolumeTypesInner) SetPricePerGbPerHour(v string)`

SetPricePerGbPerHour sets PricePerGbPerHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


