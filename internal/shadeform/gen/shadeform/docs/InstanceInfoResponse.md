# InstanceInfoResponse

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
**VolumeIds** | Pointer to **[]string** | List of volume IDs to be mounted. Currently only supports 1 volume at a time. | [optional] 
**SshKeyId** | Pointer to **string** | The ID of the SSH Key. | [optional] 
**LaunchConfiguration** | Pointer to [**LaunchConfiguration**](LaunchConfiguration.md) |  | [optional] 
**AutoDelete** | Pointer to [**AutoDelete**](AutoDelete.md) |  | [optional] 
**Alert** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**VolumeMount** | Pointer to [**VolumeMount**](VolumeMount.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Add custom, searchable tags to instances. | [optional] 
**Envs** | Pointer to [**[]Env**](Env.md) | List of environment variable name and values to automatically add to the instance | [optional] 
**PortMappings** | Pointer to [**[]InstancePortMappings**](InstancePortMappings.md) | List of port mappings on an instance that a Cloud Provider might have automatically set up. | [optional] 
**ActiveAt** | Pointer to **time.Time** | The timestamp of when the instance was active in UTC. | [optional] 
**CreatedAt** | **time.Time** | The timestamp of when the instance was created in UTC. | 
**DeletedAt** | **time.Time** | The timestamp of when the instance was deleted in UTC. | 
**BootTime** | Pointer to [**BootTime**](BootTime.md) |  | [optional] 

## Methods

### NewInstanceInfoResponse

`func NewInstanceInfoResponse(id string, cloud Cloud, region string, shadeInstanceType string, cloudInstanceType string, cloudAssignedId string, shadeCloud bool, name string, configuration InstanceConfiguration, ip string, sshUser string, sshPort int32, status Status, costEstimate string, createdAt time.Time, deletedAt time.Time, ) *InstanceInfoResponse`

NewInstanceInfoResponse instantiates a new InstanceInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInfoResponseWithDefaults

`func NewInstanceInfoResponseWithDefaults() *InstanceInfoResponse`

NewInstanceInfoResponseWithDefaults instantiates a new InstanceInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCloud

`func (o *InstanceInfoResponse) GetCloud() Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceInfoResponse) GetCloudOk() (*Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceInfoResponse) SetCloud(v Cloud)`

SetCloud sets Cloud field to given value.


### GetRegion

`func (o *InstanceInfoResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceInfoResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceInfoResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetShadeInstanceType

`func (o *InstanceInfoResponse) GetShadeInstanceType() string`

GetShadeInstanceType returns the ShadeInstanceType field if non-nil, zero value otherwise.

### GetShadeInstanceTypeOk

`func (o *InstanceInfoResponse) GetShadeInstanceTypeOk() (*string, bool)`

GetShadeInstanceTypeOk returns a tuple with the ShadeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeInstanceType

`func (o *InstanceInfoResponse) SetShadeInstanceType(v string)`

SetShadeInstanceType sets ShadeInstanceType field to given value.


### GetCloudInstanceType

`func (o *InstanceInfoResponse) GetCloudInstanceType() string`

GetCloudInstanceType returns the CloudInstanceType field if non-nil, zero value otherwise.

### GetCloudInstanceTypeOk

`func (o *InstanceInfoResponse) GetCloudInstanceTypeOk() (*string, bool)`

GetCloudInstanceTypeOk returns a tuple with the CloudInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceType

`func (o *InstanceInfoResponse) SetCloudInstanceType(v string)`

SetCloudInstanceType sets CloudInstanceType field to given value.


### GetCloudAssignedId

`func (o *InstanceInfoResponse) GetCloudAssignedId() string`

GetCloudAssignedId returns the CloudAssignedId field if non-nil, zero value otherwise.

### GetCloudAssignedIdOk

`func (o *InstanceInfoResponse) GetCloudAssignedIdOk() (*string, bool)`

GetCloudAssignedIdOk returns a tuple with the CloudAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssignedId

`func (o *InstanceInfoResponse) SetCloudAssignedId(v string)`

SetCloudAssignedId sets CloudAssignedId field to given value.


### GetShadeCloud

`func (o *InstanceInfoResponse) GetShadeCloud() bool`

GetShadeCloud returns the ShadeCloud field if non-nil, zero value otherwise.

### GetShadeCloudOk

`func (o *InstanceInfoResponse) GetShadeCloudOk() (*bool, bool)`

GetShadeCloudOk returns a tuple with the ShadeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadeCloud

`func (o *InstanceInfoResponse) SetShadeCloud(v bool)`

SetShadeCloud sets ShadeCloud field to given value.


### GetName

`func (o *InstanceInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceInfoResponse) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *InstanceInfoResponse) GetConfiguration() InstanceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InstanceInfoResponse) GetConfigurationOk() (*InstanceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InstanceInfoResponse) SetConfiguration(v InstanceConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetIp

`func (o *InstanceInfoResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceInfoResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceInfoResponse) SetIp(v string)`

SetIp sets Ip field to given value.


### GetSshUser

`func (o *InstanceInfoResponse) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *InstanceInfoResponse) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *InstanceInfoResponse) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.


### GetSshPort

`func (o *InstanceInfoResponse) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *InstanceInfoResponse) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *InstanceInfoResponse) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.


### GetStatus

`func (o *InstanceInfoResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceInfoResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceInfoResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *InstanceInfoResponse) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *InstanceInfoResponse) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *InstanceInfoResponse) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *InstanceInfoResponse) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetCostEstimate

`func (o *InstanceInfoResponse) GetCostEstimate() string`

GetCostEstimate returns the CostEstimate field if non-nil, zero value otherwise.

### GetCostEstimateOk

`func (o *InstanceInfoResponse) GetCostEstimateOk() (*string, bool)`

GetCostEstimateOk returns a tuple with the CostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimate

`func (o *InstanceInfoResponse) SetCostEstimate(v string)`

SetCostEstimate sets CostEstimate field to given value.


### GetHourlyPrice

`func (o *InstanceInfoResponse) GetHourlyPrice() int32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *InstanceInfoResponse) GetHourlyPriceOk() (*int32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *InstanceInfoResponse) SetHourlyPrice(v int32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *InstanceInfoResponse) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetVolumeIds

`func (o *InstanceInfoResponse) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *InstanceInfoResponse) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *InstanceInfoResponse) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.

### HasVolumeIds

`func (o *InstanceInfoResponse) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### GetSshKeyId

`func (o *InstanceInfoResponse) GetSshKeyId() string`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *InstanceInfoResponse) GetSshKeyIdOk() (*string, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *InstanceInfoResponse) SetSshKeyId(v string)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *InstanceInfoResponse) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.

### GetLaunchConfiguration

`func (o *InstanceInfoResponse) GetLaunchConfiguration() LaunchConfiguration`

GetLaunchConfiguration returns the LaunchConfiguration field if non-nil, zero value otherwise.

### GetLaunchConfigurationOk

`func (o *InstanceInfoResponse) GetLaunchConfigurationOk() (*LaunchConfiguration, bool)`

GetLaunchConfigurationOk returns a tuple with the LaunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfiguration

`func (o *InstanceInfoResponse) SetLaunchConfiguration(v LaunchConfiguration)`

SetLaunchConfiguration sets LaunchConfiguration field to given value.

### HasLaunchConfiguration

`func (o *InstanceInfoResponse) HasLaunchConfiguration() bool`

HasLaunchConfiguration returns a boolean if a field has been set.

### GetAutoDelete

`func (o *InstanceInfoResponse) GetAutoDelete() AutoDelete`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *InstanceInfoResponse) GetAutoDeleteOk() (*AutoDelete, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *InstanceInfoResponse) SetAutoDelete(v AutoDelete)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *InstanceInfoResponse) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAlert

`func (o *InstanceInfoResponse) GetAlert() Alert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *InstanceInfoResponse) GetAlertOk() (*Alert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *InstanceInfoResponse) SetAlert(v Alert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *InstanceInfoResponse) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetVolumeMount

`func (o *InstanceInfoResponse) GetVolumeMount() VolumeMount`

GetVolumeMount returns the VolumeMount field if non-nil, zero value otherwise.

### GetVolumeMountOk

`func (o *InstanceInfoResponse) GetVolumeMountOk() (*VolumeMount, bool)`

GetVolumeMountOk returns a tuple with the VolumeMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMount

`func (o *InstanceInfoResponse) SetVolumeMount(v VolumeMount)`

SetVolumeMount sets VolumeMount field to given value.

### HasVolumeMount

`func (o *InstanceInfoResponse) HasVolumeMount() bool`

HasVolumeMount returns a boolean if a field has been set.

### GetTags

`func (o *InstanceInfoResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceInfoResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceInfoResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceInfoResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvs

`func (o *InstanceInfoResponse) GetEnvs() []Env`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *InstanceInfoResponse) GetEnvsOk() (*[]Env, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *InstanceInfoResponse) SetEnvs(v []Env)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *InstanceInfoResponse) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetPortMappings

`func (o *InstanceInfoResponse) GetPortMappings() []InstancePortMappings`

GetPortMappings returns the PortMappings field if non-nil, zero value otherwise.

### GetPortMappingsOk

`func (o *InstanceInfoResponse) GetPortMappingsOk() (*[]InstancePortMappings, bool)`

GetPortMappingsOk returns a tuple with the PortMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappings

`func (o *InstanceInfoResponse) SetPortMappings(v []InstancePortMappings)`

SetPortMappings sets PortMappings field to given value.

### HasPortMappings

`func (o *InstanceInfoResponse) HasPortMappings() bool`

HasPortMappings returns a boolean if a field has been set.

### GetActiveAt

`func (o *InstanceInfoResponse) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *InstanceInfoResponse) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *InstanceInfoResponse) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *InstanceInfoResponse) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceInfoResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceInfoResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceInfoResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *InstanceInfoResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InstanceInfoResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InstanceInfoResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetBootTime

`func (o *InstanceInfoResponse) GetBootTime() BootTime`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *InstanceInfoResponse) GetBootTimeOk() (*BootTime, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *InstanceInfoResponse) SetBootTime(v BootTime)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *InstanceInfoResponse) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


