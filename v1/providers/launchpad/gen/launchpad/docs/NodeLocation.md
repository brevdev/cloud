# NodeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** |  | 
**Provider** | [**LocationProvider**](LocationProvider.md) |  | 
**Region** | Pointer to **string** | Name of the region the location is in | [optional] 

## Methods

### NewNodeLocation

`func NewNodeLocation(created time.Time, id string, modified time.Time, name string, provider LocationProvider, ) *NodeLocation`

NewNodeLocation instantiates a new NodeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeLocationWithDefaults

`func NewNodeLocationWithDefaults() *NodeLocation`

NewNodeLocationWithDefaults instantiates a new NodeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *NodeLocation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NodeLocation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NodeLocation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *NodeLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeLocation) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *NodeLocation) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NodeLocation) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NodeLocation) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *NodeLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeLocation) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *NodeLocation) GetProvider() LocationProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NodeLocation) GetProviderOk() (*LocationProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NodeLocation) SetProvider(v LocationProvider)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *NodeLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NodeLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NodeLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NodeLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


