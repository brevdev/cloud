# InstanceTypePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency code | 
**OnDemandPerHour** | **float64** | Price per hour | 

## Methods

### NewInstanceTypePrice

`func NewInstanceTypePrice(currency string, onDemandPerHour float64, ) *InstanceTypePrice`

NewInstanceTypePrice instantiates a new InstanceTypePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypePriceWithDefaults

`func NewInstanceTypePriceWithDefaults() *InstanceTypePrice`

NewInstanceTypePriceWithDefaults instantiates a new InstanceTypePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InstanceTypePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InstanceTypePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InstanceTypePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetOnDemandPerHour

`func (o *InstanceTypePrice) GetOnDemandPerHour() float64`

GetOnDemandPerHour returns the OnDemandPerHour field if non-nil, zero value otherwise.

### GetOnDemandPerHourOk

`func (o *InstanceTypePrice) GetOnDemandPerHourOk() (*float64, bool)`

GetOnDemandPerHourOk returns a tuple with the OnDemandPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandPerHour

`func (o *InstanceTypePrice) SetOnDemandPerHour(v float64)`

SetOnDemandPerHour sets OnDemandPerHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


