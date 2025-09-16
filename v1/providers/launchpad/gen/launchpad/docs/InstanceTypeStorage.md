# InstanceTypeStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGb** | **int32** | System storage size (in GB) | 
**Type** | [**TypeEnum**](TypeEnum.md) | Disk type  * &#x60;nvme&#x60; - NVMe * &#x60;ssd&#x60; - ssd | 

## Methods

### NewInstanceTypeStorage

`func NewInstanceTypeStorage(sizeGb int32, type_ TypeEnum, ) *InstanceTypeStorage`

NewInstanceTypeStorage instantiates a new InstanceTypeStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeStorageWithDefaults

`func NewInstanceTypeStorageWithDefaults() *InstanceTypeStorage`

NewInstanceTypeStorageWithDefaults instantiates a new InstanceTypeStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGb

`func (o *InstanceTypeStorage) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *InstanceTypeStorage) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *InstanceTypeStorage) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetType

`func (o *InstanceTypeStorage) GetType() TypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceTypeStorage) GetTypeOk() (*TypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceTypeStorage) SetType(v TypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


