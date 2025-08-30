# UfwRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** | The UFW rule action | [optional] 
**FromIp** | Pointer to **NullableString** | Source IP address or CIDR range | [optional] 
**ToIp** | Pointer to **NullableString** | Destination IP address or CIDR range | [optional] 
**Port** | Pointer to **NullableString** | Port number or range (e.g. &#39;80&#39; or &#39;8000:8999&#39;) | [optional] 
**Proto** | Pointer to **NullableString** | Protocol (e.g. tcp, udp) | [optional] 

## Methods

### NewUfwRule

`func NewUfwRule() *UfwRule`

NewUfwRule instantiates a new UfwRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfwRuleWithDefaults

`func NewUfwRuleWithDefaults() *UfwRule`

NewUfwRuleWithDefaults instantiates a new UfwRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *UfwRule) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UfwRule) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UfwRule) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *UfwRule) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetFromIp

`func (o *UfwRule) GetFromIp() string`

GetFromIp returns the FromIp field if non-nil, zero value otherwise.

### GetFromIpOk

`func (o *UfwRule) GetFromIpOk() (*string, bool)`

GetFromIpOk returns a tuple with the FromIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIp

`func (o *UfwRule) SetFromIp(v string)`

SetFromIp sets FromIp field to given value.

### HasFromIp

`func (o *UfwRule) HasFromIp() bool`

HasFromIp returns a boolean if a field has been set.

### SetFromIpNil

`func (o *UfwRule) SetFromIpNil(b bool)`

 SetFromIpNil sets the value for FromIp to be an explicit nil

### UnsetFromIp
`func (o *UfwRule) UnsetFromIp()`

UnsetFromIp ensures that no value is present for FromIp, not even an explicit nil
### GetToIp

`func (o *UfwRule) GetToIp() string`

GetToIp returns the ToIp field if non-nil, zero value otherwise.

### GetToIpOk

`func (o *UfwRule) GetToIpOk() (*string, bool)`

GetToIpOk returns a tuple with the ToIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIp

`func (o *UfwRule) SetToIp(v string)`

SetToIp sets ToIp field to given value.

### HasToIp

`func (o *UfwRule) HasToIp() bool`

HasToIp returns a boolean if a field has been set.

### SetToIpNil

`func (o *UfwRule) SetToIpNil(b bool)`

 SetToIpNil sets the value for ToIp to be an explicit nil

### UnsetToIp
`func (o *UfwRule) UnsetToIp()`

UnsetToIp ensures that no value is present for ToIp, not even an explicit nil
### GetPort

`func (o *UfwRule) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UfwRule) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UfwRule) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UfwRule) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *UfwRule) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *UfwRule) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProto

`func (o *UfwRule) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *UfwRule) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *UfwRule) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *UfwRule) HasProto() bool`

HasProto returns a boolean if a field has been set.

### SetProtoNil

`func (o *UfwRule) SetProtoNil(b bool)`

 SetProtoNil sets the value for Proto to be an explicit nil

### UnsetProto
`func (o *UfwRule) UnsetProto()`

UnsetProto ensures that no value is present for Proto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


