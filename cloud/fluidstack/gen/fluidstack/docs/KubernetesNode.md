# KubernetesNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the kubernetes node | 
**InstanceID** | **string** | Unique identifier of the instance backing the kubernetes node | 
**State** | **string** | State of the kubernetes node | 
**KubeletVersion** | **string** | Version of the kubelet running on the kubernetes node | 
**CreatedAt** | **time.Time** | Creation time of the kubernetes node | 
**Conditions** | [**[]KubernetesNodeConditionsInner**](KubernetesNodeConditionsInner.md) | Conditions of the kubernetes node | 
**Accelerator** | Pointer to [**Accelerator**](Accelerator.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** | Labels assigned to the kubernetes node | [optional] 
**Taints** | Pointer to [**[]KubernetesNodeTaintsInner**](KubernetesNodeTaintsInner.md) | Taints applied to the kubernetes node | [optional] 

## Methods

### NewKubernetesNode

`func NewKubernetesNode(name string, instanceID string, state string, kubeletVersion string, createdAt time.Time, conditions []KubernetesNodeConditionsInner, ) *KubernetesNode`

NewKubernetesNode instantiates a new KubernetesNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeWithDefaults

`func NewKubernetesNodeWithDefaults() *KubernetesNode`

NewKubernetesNodeWithDefaults instantiates a new KubernetesNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNode) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceID

`func (o *KubernetesNode) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *KubernetesNode) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *KubernetesNode) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.


### GetState

`func (o *KubernetesNode) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KubernetesNode) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KubernetesNode) SetState(v string)`

SetState sets State field to given value.


### GetKubeletVersion

`func (o *KubernetesNode) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *KubernetesNode) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *KubernetesNode) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.


### GetCreatedAt

`func (o *KubernetesNode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesNode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesNode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetConditions

`func (o *KubernetesNode) GetConditions() []KubernetesNodeConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KubernetesNode) GetConditionsOk() (*[]KubernetesNodeConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KubernetesNode) SetConditions(v []KubernetesNodeConditionsInner)`

SetConditions sets Conditions field to given value.


### GetAccelerator

`func (o *KubernetesNode) GetAccelerator() Accelerator`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *KubernetesNode) GetAcceleratorOk() (*Accelerator, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *KubernetesNode) SetAccelerator(v Accelerator)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *KubernetesNode) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesNode) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNode) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNode) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesNode) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaints

`func (o *KubernetesNode) GetTaints() []KubernetesNodeTaintsInner`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *KubernetesNode) GetTaintsOk() (*[]KubernetesNodeTaintsInner, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *KubernetesNode) SetTaints(v []KubernetesNodeTaintsInner)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *KubernetesNode) HasTaints() bool`

HasTaints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


