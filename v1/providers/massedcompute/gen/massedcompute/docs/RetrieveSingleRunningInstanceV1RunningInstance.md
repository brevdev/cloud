# RetrieveSingleRunningInstanceV1RunningInstance

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
**Product** | Pointer to [**RetrieveSingleRunningInstanceV1RunningInstanceProduct**](RetrieveSingleRunningInstanceV1RunningInstanceProduct.md) |  | [optional]

## Methods

### NewRetrieveSingleRunningInstanceV1RunningInstance

`func NewRetrieveSingleRunningInstanceV1RunningInstance() *RetrieveSingleRunningInstanceV1RunningInstance`

NewRetrieveSingleRunningInstanceV1RunningInstance instantiates a new RetrieveSingleRunningInstanceV1RunningInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveSingleRunningInstanceV1RunningInstanceWithDefaults

`func NewRetrieveSingleRunningInstanceV1RunningInstanceWithDefaults() *RetrieveSingleRunningInstanceV1RunningInstance`

NewRetrieveSingleRunningInstanceV1RunningInstanceWithDefaults instantiates a new RetrieveSingleRunningInstanceV1RunningInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUsername

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOsBooted

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetOsBooted() int32`

GetOsBooted returns the OsBooted field if non-nil, zero value otherwise.

### GetOsBootedOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetOsBootedOk() (*int32, bool)`

GetOsBootedOk returns a tuple with the OsBooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBooted

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetOsBooted(v int32)`

SetOsBooted sets OsBooted field to given value.

### HasOsBooted

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasOsBooted() bool`

HasOsBooted returns a boolean if a field has been set.

### GetCommandStartup

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetCommandStartup() string`

GetCommandStartup returns the CommandStartup field if non-nil, zero value otherwise.

### GetCommandStartupOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetCommandStartupOk() (*string, bool)`

GetCommandStartupOk returns a tuple with the CommandStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandStartup

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetCommandStartup(v string)`

SetCommandStartup sets CommandStartup field to given value.

### HasCommandStartup

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasCommandStartup() bool`

HasCommandStartup returns a boolean if a field has been set.

### GetCreated

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetActive

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetImage

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetImage() RetrieveAllRunningInstancesV1RunningInstancesInnerImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetImageOk() (*RetrieveAllRunningInstancesV1RunningInstancesInnerImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetImage(v RetrieveAllRunningInstancesV1RunningInstancesInnerImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetProduct

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetProduct() RetrieveSingleRunningInstanceV1RunningInstanceProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) GetProductOk() (*RetrieveSingleRunningInstanceV1RunningInstanceProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) SetProduct(v RetrieveSingleRunningInstanceV1RunningInstanceProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *RetrieveSingleRunningInstanceV1RunningInstance) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


