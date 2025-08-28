# CreateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | [**Cloud**](Cloud.md) |  | 
**Region** | **string** | Specifies the region. | 
**SizeInGb** | **int32** | Storage volume size in GB | 
**Name** | **string** | The name of the storage volume. | 

## Methods

### NewCreateVolumeRequest

`func NewCreateVolumeRequest(cloud Cloud, region string, sizeInGb int32, name string, ) *CreateVolumeRequest`

NewCreateVolumeRequest instantiates a new CreateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeRequestWithDefaults

`func NewCreateVolumeRequestWithDefaults() *CreateVolumeRequest`

NewCreateVolumeRequestWithDefaults instantiates a new CreateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *CreateVolumeRequest) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CreateVolumeRequest) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CreateVolumeRequest) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *CreateVolumeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateVolumeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateVolumeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSizeInGb

`func (o *CreateVolumeRequest) GetSizeInGb() int32`

GetSizeInGb returns the SizeInGb field if non-nil, zero value otherwise.

### GetSizeInGbOk

`func (o *CreateVolumeRequest) GetSizeInGbOk() (*int32, bool)`

GetSizeInGbOk returns a tuple with the SizeInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInGb

`func (o *CreateVolumeRequest) SetSizeInGb(v int32)`

SetSizeInGb sets SizeInGb field to given value.


### GetName

`func (o *CreateVolumeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumeRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


