# ClusterProvisionCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | [**ClusterPipeline**](ClusterPipeline.md) |  | [readonly] 
**ProvisioningConfig** | [**ProvisioningRequest**](ProvisioningRequest.md) |  | [readonly] 

## Methods

### NewClusterProvisionCreateResponse

`func NewClusterProvisionCreateResponse(pipeline ClusterPipeline, provisioningConfig ProvisioningRequest, ) *ClusterProvisionCreateResponse`

NewClusterProvisionCreateResponse instantiates a new ClusterProvisionCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterProvisionCreateResponseWithDefaults

`func NewClusterProvisionCreateResponseWithDefaults() *ClusterProvisionCreateResponse`

NewClusterProvisionCreateResponseWithDefaults instantiates a new ClusterProvisionCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *ClusterProvisionCreateResponse) GetPipeline() ClusterPipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *ClusterProvisionCreateResponse) GetPipelineOk() (*ClusterPipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *ClusterProvisionCreateResponse) SetPipeline(v ClusterPipeline)`

SetPipeline sets Pipeline field to given value.


### GetProvisioningConfig

`func (o *ClusterProvisionCreateResponse) GetProvisioningConfig() ProvisioningRequest`

GetProvisioningConfig returns the ProvisioningConfig field if non-nil, zero value otherwise.

### GetProvisioningConfigOk

`func (o *ClusterProvisionCreateResponse) GetProvisioningConfigOk() (*ProvisioningRequest, bool)`

GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConfig

`func (o *ClusterProvisionCreateResponse) SetProvisioningConfig(v ProvisioningRequest)`

SetProvisioningConfig sets ProvisioningConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


