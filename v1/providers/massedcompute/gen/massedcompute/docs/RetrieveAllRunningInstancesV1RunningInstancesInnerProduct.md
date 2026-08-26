# RetrieveAllRunningInstancesV1RunningInstancesInnerProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**GpuCount** | Pointer to **int32** |  | [optional]
**Vcpu** | Pointer to **int32** |  | [optional]
**Ram** | Pointer to **int32** |  | [optional]
**Storage** | Pointer to **int32** |  | [optional]
**PriceHr** | Pointer to **string** |  | [optional]
**FinalPriceHr** | Pointer to **string** |  | [optional]

## Methods

### NewRetrieveAllRunningInstancesV1RunningInstancesInnerProduct

`func NewRetrieveAllRunningInstancesV1RunningInstancesInnerProduct() *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct`

NewRetrieveAllRunningInstancesV1RunningInstancesInnerProduct instantiates a new RetrieveAllRunningInstancesV1RunningInstancesInnerProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveAllRunningInstancesV1RunningInstancesInnerProductWithDefaults

`func NewRetrieveAllRunningInstancesV1RunningInstancesInnerProductWithDefaults() *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct`

NewRetrieveAllRunningInstancesV1RunningInstancesInnerProductWithDefaults instantiates a new RetrieveAllRunningInstancesV1RunningInstancesInnerProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGpuCount

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetVcpu

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetRam

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetPriceHr() string`

GetPriceHr returns the PriceHr field if non-nil, zero value otherwise.

### GetPriceHrOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetPriceHrOk() (*string, bool)`

GetPriceHrOk returns a tuple with the PriceHr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetPriceHr(v string)`

SetPriceHr sets PriceHr field to given value.

### HasPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasPriceHr() bool`

HasPriceHr returns a boolean if a field has been set.

### GetFinalPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetFinalPriceHr() string`

GetFinalPriceHr returns the FinalPriceHr field if non-nil, zero value otherwise.

### GetFinalPriceHrOk

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) GetFinalPriceHrOk() (*string, bool)`

GetFinalPriceHrOk returns a tuple with the FinalPriceHr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) SetFinalPriceHr(v string)`

SetFinalPriceHr sets FinalPriceHr field to given value.

### HasFinalPriceHr

`func (o *RetrieveAllRunningInstancesV1RunningInstancesInnerProduct) HasFinalPriceHr() bool`

HasFinalPriceHr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


