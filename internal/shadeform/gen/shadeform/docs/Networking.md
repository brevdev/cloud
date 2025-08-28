# Networking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UfwRules** | Pointer to [**[]UfwRule**](UfwRule.md) | List of UFW (Uncomplicated Firewall) rules | [optional] 

## Methods

### NewNetworking

`func NewNetworking() *Networking`

NewNetworking instantiates a new Networking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingWithDefaults

`func NewNetworkingWithDefaults() *Networking`

NewNetworkingWithDefaults instantiates a new Networking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUfwRules

`func (o *Networking) GetUfwRules() []UfwRule`

GetUfwRules returns the UfwRules field if non-nil, zero value otherwise.

### GetUfwRulesOk

`func (o *Networking) GetUfwRulesOk() (*[]UfwRule, bool)`

GetUfwRulesOk returns a tuple with the UfwRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUfwRules

`func (o *Networking) SetUfwRules(v []UfwRule)`

SetUfwRules sets UfwRules field to given value.

### HasUfwRules

`func (o *Networking) HasUfwRules() bool`

HasUfwRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


