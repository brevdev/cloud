# NodeStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Disk size (in GB) | [optional] [default to 0]
**Type** | [**TypeEnum**](TypeEnum.md) | Disk type  * &#x60;nvme&#x60; - NVMe * &#x60;ssd&#x60; - ssd | 

## Methods

### NewNodeStorage

`func NewNodeStorage(type_ TypeEnum, ) *NodeStorage`

NewNodeStorage instantiates a new NodeStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStorageWithDefaults

`func NewNodeStorageWithDefaults() *NodeStorage`

NewNodeStorageWithDefaults instantiates a new NodeStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *NodeStorage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NodeStorage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NodeStorage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *NodeStorage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *NodeStorage) GetType() TypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeStorage) GetTypeOk() (*TypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeStorage) SetType(v TypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


