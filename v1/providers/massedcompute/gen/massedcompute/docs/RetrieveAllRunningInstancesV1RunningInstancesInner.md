# RetrieveAllRunningInstancesV1RunningInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**Ip** | Pointer to **string** |  | [optional]
**Username** | Pointer to **string** |  | [optional]
**Password** | Pointer to **string** |  | [optional]
**Status** | Pointer to **string** |  | [optional]
**OsBooted** | Pointer to **int32** |  | [optional]
**CommandStartup** | Pointer to **string** |  | [optional]
**Created** | Pointer to **string** |  | [optional]
**Active** | Pointer to **int32** |  | [optional]
**Image** | Pointer to [**RetrieveAllRunningInstancesV1RunningInstancesInnerImage**](RetrieveAllRunningInstancesV1RunningInstancesInnerImage.md) |  | [optional]
**Product** | Pointer to [**RetrieveAllRunningInstancesV1RunningInstancesInnerProduct**](RetrieveAllRunningInstancesV1RunningInstancesInnerProduct.md) |  | [optional]

## Methods

### NewRetrieveAllRunningInstancesV1RunningInstancesInner

`func NewRetrieveAllRunningInstancesV1RunningInstancesInner() *RetrieveAllRunningInstancesV1RunningInstancesInner`

NewRetrieveAllRunningInstancesV1RunningInstancesInner instantiates a new RetrieveAllRunningInstancesV1RunningInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveAllRunningInstancesV1RunningInstancesInnerWithDefaults

`func NewRetrieveAllRunningInstancesV1RunningInstancesInnerWithDefaults() *RetrieveAllRunningInstancesV1RunningInstancesInner`

NewRetrieveAllRunningInstancesV1RunningInstancesInnerWithDefaults instantiates a new RetrieveAllRunningInstancesV1RunningInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUsername

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOsBooted

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetOsBooted() int32`

GetOsBooted returns the OsBooted field if non-nil, zero value otherwise.

### GetOsBootedOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetOsBootedOk() (*int32, bool)`

GetOsBootedOk returns a tuple with the OsBooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBooted

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetOsBooted(v int32)`

SetOsBooted sets OsBooted field to given value.

### HasOsBooted

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasOsBooted() bool`

HasOsBooted returns a boolean if a field has been set.

### GetCommandStartup

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetCommandStartup() string`

GetCommandStartup returns the CommandStartup field if non-nil, zero value otherwise.

### GetCommandStartupOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetCommandStartupOk() (*string, bool)`

GetCommandStartupOk returns a tuple with the CommandStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandStartup

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetCommandStartup(v string)`

SetCommandStartup sets CommandStartup field to given value.

### HasCommandStartup

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasCommandStartup() bool`

HasCommandStartup returns a boolean if a field has been set.

### GetCreated

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetActive

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetImage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetImage() RetrieveAllRunningInstancesV1RunningInstancesInnerImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetImageOk() (*RetrieveAllRunningInstancesV1RunningInstancesInnerImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetImage(v RetrieveAllRunningInstancesV1RunningInstancesInnerImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetProduct

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetProduct() RetrieveAllRunningInstancesV1RunningInstancesInnerProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) GetProductOk() (*RetrieveAllRunningInstancesV1RunningInstancesInnerProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) SetProduct(v RetrieveAllRunningInstancesV1RunningInstancesInnerProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInner) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


