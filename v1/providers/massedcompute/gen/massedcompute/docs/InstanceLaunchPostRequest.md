# InstanceLaunchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **int32** | The ID of the image to deploy |
**ProductName** | **string** | The product name of the GPU instance you want to deploy. Example &#x3D; &#39;gpu_1x_l40&#39; |
**RegionName** | **string** | Set value equal to &#39;any&#39; |
**InstanceName** | Pointer to **string** | The name of the instance you want to deploy | [optional]
**Coupon** | Pointer to **string** | The coupon code you want to apply to the instance | [optional]
**Command** | Pointer to **string** | The command you want to run on startup | [optional]
**SshKeys** | Pointer to **[]string** | The SSH key you want to use to connect to the instance | [optional]

## Methods

### NewInstanceLaunchPostRequest

`func NewInstanceLaunchPostRequest(imageId int32, productName string, regionName string, ) *InstanceLaunchPostRequest`

NewInstanceLaunchPostRequest instantiates a new InstanceLaunchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceLaunchPostRequestWithDefaults

`func NewInstanceLaunchPostRequestWithDefaults() *InstanceLaunchPostRequest`

NewInstanceLaunchPostRequestWithDefaults instantiates a new InstanceLaunchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *InstanceLaunchPostRequest) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *InstanceLaunchPostRequest) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *InstanceLaunchPostRequest) SetImageId(v int32)`

SetImageId sets ImageId field to given value.


### GetProductName

`func (o *InstanceLaunchPostRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *InstanceLaunchPostRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *InstanceLaunchPostRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetRegionName

`func (o *InstanceLaunchPostRequest) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *InstanceLaunchPostRequest) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *InstanceLaunchPostRequest) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetInstanceName

`func (o *InstanceLaunchPostRequest) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *InstanceLaunchPostRequest) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *InstanceLaunchPostRequest) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *InstanceLaunchPostRequest) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetCoupon

`func (o *InstanceLaunchPostRequest) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *InstanceLaunchPostRequest) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *InstanceLaunchPostRequest) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *InstanceLaunchPostRequest) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetCommand

`func (o *InstanceLaunchPostRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *InstanceLaunchPostRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *InstanceLaunchPostRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *InstanceLaunchPostRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSshKeys

`func (o *InstanceLaunchPostRequest) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstanceLaunchPostRequest) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstanceLaunchPostRequest) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstanceLaunchPostRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


