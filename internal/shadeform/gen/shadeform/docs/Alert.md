# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateThreshold** | Pointer to **string** | RFC3339 date string | [optional] 
**SpendThreshold** | Pointer to **string** | Valid decimal representation of a dollar amount | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateThreshold

`func (o *Alert) GetDateThreshold() string`

GetDateThreshold returns the DateThreshold field if non-nil, zero value otherwise.

### GetDateThresholdOk

`func (o *Alert) GetDateThresholdOk() (*string, bool)`

GetDateThresholdOk returns a tuple with the DateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateThreshold

`func (o *Alert) SetDateThreshold(v string)`

SetDateThreshold sets DateThreshold field to given value.

### HasDateThreshold

`func (o *Alert) HasDateThreshold() bool`

HasDateThreshold returns a boolean if a field has been set.

### GetSpendThreshold

`func (o *Alert) GetSpendThreshold() string`

GetSpendThreshold returns the SpendThreshold field if non-nil, zero value otherwise.

### GetSpendThresholdOk

`func (o *Alert) GetSpendThresholdOk() (*string, bool)`

GetSpendThresholdOk returns a tuple with the SpendThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendThreshold

`func (o *Alert) SetSpendThreshold(v string)`

SetSpendThreshold sets SpendThreshold field to given value.

### HasSpendThreshold

`func (o *Alert) HasSpendThreshold() bool`

HasSpendThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


