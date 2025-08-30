# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the instance. Used in the instances for the /instances/{id}/info and /instances/{id}/delete APIs. | 
**Cloud** | [**Cloud**](Cloud.md) |  | 
**Region** | **string** | Specifies the region. | 
**ShadeInstanceType** | **string** | The Shadeform standardized instance type. See this [explanation](/getting-started/concepts#shade-instance-type-and-cloud-instance-type) for more details. | 
**CloudInstanceType** | **string** | The instance type for the underlying cloud provider. See this [explanation](/getting-started/concepts#shade-instance-type-and-cloud-instance-type) for more details. | 
**CloudAssignedId** | **string** | The unique identifier of the instance issued by the underlying cloud provider. | 
**ShadeCloud** | **bool** | Specifies if the instance is launched in [Shade Cloud](/getting-started/concepts#shade-cloud) or in a linked cloud account. | 
**Name** | **string** | The name of the instance | 
**Configuration** | [**InstanceConfiguration**](InstanceConfiguration.md) |  | 
**Ip** | **string** | The public IP address of the instance. In select cases, it may also be the DNS. | 
**SshUser** | **string** | The SSH user used to SSH into the instance. | 
**SshPort** | **int32** | The SSH port of the instance. In most cases, this will be port 22 but for some clouds, this may be a different port. | 
**Status** | [**Status**](Status.md) |  | 
**StatusDetails** | Pointer to **string** | Additional context for the status | [optional] 
**CostEstimate** | **string** | The cost incurred by the instance. This only the cost via Shadeform. If the instance is deployed in your own cloud account, then all billing is through your cloud provider. | 
**HourlyPrice** | Pointer to **int32** | The hourly price of the instance in cents. | [optional] 
**LaunchConfiguration** | Pointer to [**LaunchConfiguration**](LaunchConfiguration.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Add custom, searchable tags to instances. | [optional] 
**PortMappings** | Pointer to [**[]InstancePortMappings**](InstancePortMappings.md) | List of port mappings on an instance that a Cloud Provider might have automatically set up. | [optional] 
**ActiveAt** | Pointer to **time.Time** | The timestamp of when the instance was active in UTC. | [optional] 
**CreatedAt** | **time.Time** | The timestamp of when the instance was created in UTC. | 
**DeletedAt** | **time.Time** | The timestamp of when the instance was deleted in UTC. | 
**BootTime** | Pointer to [**BootTime**](BootTime.md) |  | [optional] 

## Methods

### NewInstance

`func NewInstance(id string, cloud Cloud, region string, shadeInstanceType string, cloudInstanceType string, cloudAssignedId string, shadeCloud bool, name string, configuration InstanceConfiguration, ip string, sshUser string, sshPort int32, status Status, costEstimate string, createdAt time.Time, deletedAt time.Time, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetCloud

`func (o *Instance) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Instance) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Instance) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *Instance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetShadeInstanceType

`func (o *Instance) GetShadeInstanceType() string`

GetShadeInstanceType returns the ShadeInstanceType field if non-nil, zero value otherwise.

### GetShadeInstanceTypeOk

`func (o *Instance) GetShadeInstanceTypeOk() (*string, bool)`

GetShadeInstanceTypeOk returns a tuple with the ShadeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeInstanceType

`func (o *Instance) SetShadeInstanceType(v string)`

SetShadeInstanceType sets ShadeInstanceType field to given value.


### GetCloudInstanceType

`func (o *Instance) GetCloudInstanceType() string`

GetCloudInstanceType returns the CloudInstanceType field if non-nil, zero value otherwise.

### GetCloudInstanceTypeOk

`func (o *Instance) GetCloudInstanceTypeOk() (*string, bool)`

GetCloudInstanceTypeOk returns a tuple with the CloudInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceType

`func (o *Instance) SetCloudInstanceType(v string)`

SetCloudInstanceType sets CloudInstanceType field to given value.


### GetCloudAssignedId

`func (o *Instance) GetCloudAssignedId() string`

GetCloudAssignedId returns the CloudAssignedId field if non-nil, zero value otherwise.

### GetCloudAssignedIdOk

`func (o *Instance) GetCloudAssignedIdOk() (*string, bool)`

GetCloudAssignedIdOk returns a tuple with the CloudAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssignedId

`func (o *Instance) SetCloudAssignedId(v string)`

SetCloudAssignedId sets CloudAssignedId field to given value.


### GetShadeCloud

`func (o *Instance) GetShadeCloud() bool`

GetShadeCloud returns the ShadeCloud field if non-nil, zero value otherwise.

### GetShadeCloudOk

`func (o *Instance) GetShadeCloudOk() (*bool, bool)`

GetShadeCloudOk returns a tuple with the ShadeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeCloud

`func (o *Instance) SetShadeCloud(v bool)`

SetShadeCloud sets ShadeCloud field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *Instance) GetConfiguration() InstanceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Instance) GetConfigurationOk() (*InstanceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Instance) SetConfiguration(v InstanceConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetIp

`func (o *Instance) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Instance) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Instance) SetIp(v string)`

SetIp sets Ip field to given value.


### GetSshUser

`func (o *Instance) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *Instance) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *Instance) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetSshPort

`func (o *Instance) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Instance) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Instance) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.


### GetStatus

`func (o *Instance) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *Instance) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Instance) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Instance) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Instance) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetCostEstimate

`func (o *Instance) GetCostEstimate() string`

GetCostEstimate returns the CostEstimate field if non-nil, zero value otherwise.

### GetCostEstimateOk

`func (o *Instance) GetCostEstimateOk() (*string, bool)`

GetCostEstimateOk returns a tuple with the CostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimate

`func (o *Instance) SetCostEstimate(v string)`

SetCostEstimate sets CostEstimate field to given value.


### GetHourlyPrice

`func (o *Instance) GetHourlyPrice() int32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *Instance) GetHourlyPriceOk() (*int32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *Instance) SetHourlyPrice(v int32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *Instance) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetLaunchConfiguration

`func (o *Instance) GetLaunchConfiguration() LaunchConfiguration`

GetLaunchConfiguration returns the LaunchConfiguration field if non-nil, zero value otherwise.

### GetLaunchConfigurationOk

`func (o *Instance) GetLaunchConfigurationOk() (*LaunchConfiguration, bool)`

GetLaunchConfigurationOk returns a tuple with the LaunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfiguration

`func (o *Instance) SetLaunchConfiguration(v LaunchConfiguration)`

SetLaunchConfiguration sets LaunchConfiguration field to given value.

### HasLaunchConfiguration

`func (o *Instance) HasLaunchConfiguration() bool`

HasLaunchConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *Instance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Instance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortMappings

`func (o *Instance) GetPortMappings() []InstancePortMappings`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *Instance) GetPortMappingsOk() (*[]InstancePortMappings, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *Instance) SetPortMappings(v []InstancePortMappings)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *Instance) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetActiveAt

`func (o *Instance) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *Instance) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *Instance) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *Instance) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Instance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Instance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Instance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *Instance) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Instance) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Instance) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetBootTime

`func (o *Instance) GetBootTime() BootTime`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *Instance) GetBootTimeOk() (*BootTime, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *Instance) SetBootTime(v BootTime)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *Instance) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


