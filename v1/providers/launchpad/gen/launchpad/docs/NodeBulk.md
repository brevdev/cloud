# NodeBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsvFile** | **string** |  | 
**Nodes** | [**[]Node**](Node.md) |  | 

## Methods

### NewNodeBulk

`func NewNodeBulk(csvFile string, nodes []Node, ) *NodeBulk`

NewNodeBulk instantiates a new NodeBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBulkWithDefaults

`func NewNodeBulkWithDefaults() *NodeBulk`

NewNodeBulkWithDefaults instantiates a new NodeBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsvFile

`func (o *NodeBulk) GetCsvFile() string`

GetCsvFile returns the CsvFile field if non-nil, zero value otherwise.

### GetCsvFileOk

`func (o *NodeBulk) GetCsvFileOk() (*string, bool)`

GetCsvFileOk returns a tuple with the CsvFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFile

`func (o *NodeBulk) SetCsvFile(v string)`

SetCsvFile sets CsvFile field to given value.


### GetNodes

`func (o *NodeBulk) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodeBulk) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodeBulk) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


