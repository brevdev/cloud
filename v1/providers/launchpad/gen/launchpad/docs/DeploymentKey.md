# DeploymentKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp of when the object was created | [readonly] 
**Deployment** | [**ClusterDeployment**](ClusterDeployment.md) |  | 
**Id** | **string** |  | [readonly] 
**Modified** | **time.Time** | Timestamp of when the object was last modified | [readonly] 
**Name** | **string** | Descriptive name for the SSH public key | 
**PublicKey** | **string** | An SSH public key that should be authorized to access the deployment | 

## Methods

### NewDeploymentKey

`func NewDeploymentKey(created time.Time, deployment ClusterDeployment, id string, modified time.Time, name string, publicKey string, ) *DeploymentKey`

NewDeploymentKey instantiates a new DeploymentKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentKeyWithDefaults

`func NewDeploymentKeyWithDefaults() *DeploymentKey`

NewDeploymentKeyWithDefaults instantiates a new DeploymentKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeploymentKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeploymentKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeploymentKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeployment

`func (o *DeploymentKey) GetDeployment() ClusterDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *DeploymentKey) GetDeploymentOk() (*ClusterDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *DeploymentKey) SetDeployment(v ClusterDeployment)`

SetDeployment sets Deployment field to given value.


### GetId

`func (o *DeploymentKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentKey) SetId(v string)`

SetId sets Id field to given value.


### GetModified

`func (o *DeploymentKey) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DeploymentKey) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DeploymentKey) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *DeploymentKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentKey) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *DeploymentKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DeploymentKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DeploymentKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


