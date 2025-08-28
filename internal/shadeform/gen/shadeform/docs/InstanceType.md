# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | [**Cloud**](Cloud.md) |  | 
**ShadeInstanceType** | **string** | The Shadeform standardized instance type. See this [explanation](/getting-started/concepts#shade-instance-type-and-cloud-instance-type) for more details. | 
**CloudInstanceType** | **string** | The instance type for the underlying cloud provider. See this [explanation](/getting-started/concepts#shade-instance-type-and-cloud-instance-type) for more details. | 
**Configuration** | [**InstanceTypeConfiguration**](InstanceTypeConfiguration.md) |  | 
**HourlyPrice** | **int32** | The hourly price of the instance in cents. | 
**DeploymentType** | **string** | Whether the instance is a \&quot;vm\&quot; (virtual machine), a \&quot;container\&quot; (docker container), or \&quot;baremetal\&quot; (bare metal server). | 
**Availability** | [**[]Availability**](Availability.md) |  | 
**BootTime** | Pointer to [**BootTime**](BootTime.md) |  | [optional] 

## Methods

### NewInstanceType

`func NewInstanceType(cloud Cloud, shadeInstanceType string, cloudInstanceType string, configuration InstanceTypeConfiguration, hourlyPrice int32, deploymentType string, availability []Availability, ) *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *InstanceType) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceType) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceType) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetShadeInstanceType

`func (o *InstanceType) GetShadeInstanceType() string`

GetShadeInstanceType returns the ShadeInstanceType field if non-nil, zero value otherwise.

### GetShadeInstanceTypeOk

`func (o *InstanceType) GetShadeInstanceTypeOk() (*string, bool)`

GetShadeInstanceTypeOk returns a tuple with the ShadeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeInstanceType

`func (o *InstanceType) SetShadeInstanceType(v string)`

SetShadeInstanceType sets ShadeInstanceType field to given value.


### GetCloudInstanceType

`func (o *InstanceType) GetCloudInstanceType() string`

GetCloudInstanceType returns the CloudInstanceType field if non-nil, zero value otherwise.

### GetCloudInstanceTypeOk

`func (o *InstanceType) GetCloudInstanceTypeOk() (*string, bool)`

GetCloudInstanceTypeOk returns a tuple with the CloudInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceType

`func (o *InstanceType) SetCloudInstanceType(v string)`

SetCloudInstanceType sets CloudInstanceType field to given value.


### GetConfiguration

`func (o *InstanceType) GetConfiguration() InstanceTypeConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InstanceType) GetConfigurationOk() (*InstanceTypeConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InstanceType) SetConfiguration(v InstanceTypeConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetHourlyPrice

`func (o *InstanceType) GetHourlyPrice() int32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *InstanceType) GetHourlyPriceOk() (*int32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *InstanceType) SetHourlyPrice(v int32)`

SetHourlyPrice sets HourlyPrice field to given value.


### GetDeploymentType

`func (o *InstanceType) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *InstanceType) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *InstanceType) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetAvailability

`func (o *InstanceType) GetAvailability() []Availability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceType) GetAvailabilityOk() (*[]Availability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceType) SetAvailability(v []Availability)`

SetAvailability sets Availability field to given value.


### GetBootTime

`func (o *InstanceType) GetBootTime() BootTime`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *InstanceType) GetBootTimeOk() (*BootTime, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *InstanceType) SetBootTime(v BootTime)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *InstanceType) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


