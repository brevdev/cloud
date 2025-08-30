# \DefaultAPI

All URIs are relative to *https://api.shadeform.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Instances**](DefaultAPI.md#Instances) | **Get** /instances | /instances
[**InstancesCreate**](DefaultAPI.md#InstancesCreate) | **Post** /instances/create | /instances/create
[**InstancesDelete**](DefaultAPI.md#InstancesDelete) | **Post** /instances/{id}/delete | /instances/{id}/delete
[**InstancesInfo**](DefaultAPI.md#InstancesInfo) | **Get** /instances/{id}/info | /instances/{id}/info
[**InstancesRestart**](DefaultAPI.md#InstancesRestart) | **Post** /instances/{id}/restart | /instances/{id}/restart
[**InstancesTypes**](DefaultAPI.md#InstancesTypes) | **Get** /instances/types | /instances/types
[**InstancesUpdate**](DefaultAPI.md#InstancesUpdate) | **Post** /instances/{id}/update | /instances/{id}/update
[**SshKeys**](DefaultAPI.md#SshKeys) | **Get** /sshkeys | /sshkeys
[**SshKeysAdd**](DefaultAPI.md#SshKeysAdd) | **Post** /sshkeys/add | /sshkeys/add
[**SshKeysDelete**](DefaultAPI.md#SshKeysDelete) | **Post** /sshkeys/{id}/delete | /sshkeys/{id}/delete
[**SshKeysInfo**](DefaultAPI.md#SshKeysInfo) | **Get** /sshkeys/{id}/info | /sshkeys/{id}/info
[**SshKeysSetDefault**](DefaultAPI.md#SshKeysSetDefault) | **Post** /sshkeys/{id}/setdefault | /sshkeys/{id}/setdefault
[**Templates**](DefaultAPI.md#Templates) | **Get** /templates | /templates
[**TemplatesDelete**](DefaultAPI.md#TemplatesDelete) | **Post** /templates/{template_id}/delete | /templates/{template_id}/delete
[**TemplatesFeatured**](DefaultAPI.md#TemplatesFeatured) | **Get** /templates/featured | /templates/featured
[**TemplatesInfo**](DefaultAPI.md#TemplatesInfo) | **Get** /templates/{template_id}/info | /templates/{template_id}/info
[**TemplatesSave**](DefaultAPI.md#TemplatesSave) | **Post** /templates/save | /templates/save
[**TemplatesUpdate**](DefaultAPI.md#TemplatesUpdate) | **Post** /templates/{template_id}/update | /templates/{template_id}/update
[**Volumes**](DefaultAPI.md#Volumes) | **Get** /volumes | /volumes
[**VolumesCreate**](DefaultAPI.md#VolumesCreate) | **Post** /volumes/create | /volumes/create
[**VolumesDelete**](DefaultAPI.md#VolumesDelete) | **Post** /volumes/{id}/delete | /volumes/{id}/delete
[**VolumesInfo**](DefaultAPI.md#VolumesInfo) | **Get** /volumes/{id}/info | /volumes/{id}/info
[**VolumesTypes**](DefaultAPI.md#VolumesTypes) | **Get** /volumes/types | /volumes/types



## Instances

> InstancesResponse Instances(ctx).Execute()

/instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Instances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Instances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Instances`: InstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Instances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesRequest struct via the builder pattern


### Return type

[**InstancesResponse**](InstancesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesCreate

> CreateResponse InstancesCreate(ctx).CreateRequest(createRequest).Execute()

/instances/create



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	createRequest := *openapiclient.NewCreateRequest(openapiclient.Cloud("lambdalabs"), "canada-1", "A6000", true, "cool-gpu-server") // CreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InstancesCreate(context.Background()).CreateRequest(createRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstancesCreate`: CreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**CreateRequest**](CreateRequest.md) |  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesDelete

> InstancesDelete(ctx, id).Execute()

/instances/{id}/delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.InstancesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInfo

> InstanceInfoResponse InstancesInfo(ctx, id).Execute()

/instances/{id}/info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InstancesInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstancesInfo`: InstanceInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InstancesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceInfoResponse**](InstanceInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesRestart

> InstancesRestart(ctx, id).Execute()

/instances/{id}/restart



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.InstancesRestart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesTypes

> InstanceTypesResponse InstancesTypes(ctx).Cloud(cloud).Region(region).NumGpus(numGpus).GpuType(gpuType).ShadeInstanceType(shadeInstanceType).Available(available).Sort(sort).Execute()

/instances/types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	cloud := "aws" // string | Filter the instance type results by cloud. (optional)
	region := "us-east-1a" // string | Filter the instance type results by region. (optional)
	numGpus := "1" // string | Filter the instance type results by the number of gpus. (optional)
	gpuType := "A100_80G" // string | Filter the instance type results by gpu type. (optional)
	shadeInstanceType := "A100_80G" // string | Filter the instance type results by the shade instance type. (optional)
	available := true // bool | Filter the instance type results by availability. (optional)
	sort := "price" // string | Sort the order of the instance type results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InstancesTypes(context.Background()).Cloud(cloud).Region(region).NumGpus(numGpus).GpuType(gpuType).ShadeInstanceType(shadeInstanceType).Available(available).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstancesTypes`: InstanceTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InstancesTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstancesTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** | Filter the instance type results by cloud. | 
 **region** | **string** | Filter the instance type results by region. | 
 **numGpus** | **string** | Filter the instance type results by the number of gpus. | 
 **gpuType** | **string** | Filter the instance type results by gpu type. | 
 **shadeInstanceType** | **string** | Filter the instance type results by the shade instance type. | 
 **available** | **bool** | Filter the instance type results by availability. | 
 **sort** | **string** | Sort the order of the instance type results. | 

### Return type

[**InstanceTypesResponse**](InstanceTypesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesUpdate

> InstancesUpdate(ctx, id).UpdateRequest(updateRequest).Execute()

/instances/{id}/update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The instance id
	updateRequest := *openapiclient.NewUpdateRequest() // UpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.InstancesUpdate(context.Background(), id).UpdateRequest(updateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeys

> SshKeysResponse SshKeys(ctx).Execute()

/sshkeys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SshKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeys`: SshKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SshKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysRequest struct via the builder pattern


### Return type

[**SshKeysResponse**](SshKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysAdd

> AddSshKeyResponse SshKeysAdd(ctx).AddSshKeyRequest(addSshKeyRequest).Execute()

/sshkeys/add



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	addSshKeyRequest := *openapiclient.NewAddSshKeyRequest("My ssh key", "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSU GPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3 Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XA t3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/En mZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbx NrRFi9wrf+M7Q== schacon@mylaptop.local") // AddSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SshKeysAdd(context.Background()).AddSshKeyRequest(addSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SshKeysAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysAdd`: AddSshKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SshKeysAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSshKeyRequest** | [**AddSshKeyRequest**](AddSshKeyRequest.md) |  | 

### Return type

[**AddSshKeyResponse**](AddSshKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysDelete

> SshKeysDelete(ctx, id).Execute()

/sshkeys/{id}/delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The ssh key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SshKeysDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SshKeysDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ssh key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysInfo

> SshKey SshKeysInfo(ctx, id).Execute()

/sshkeys/{id}/info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The ssh key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SshKeysInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SshKeysInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysInfo`: SshKey
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SshKeysInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ssh key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SshKey**](SshKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysSetDefault

> SshKeysSetDefault(ctx, id).Execute()

/sshkeys/{id}/setdefault



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The ssh key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SshKeysSetDefault(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SshKeysSetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ssh key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysSetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Templates

> TemplatesResponse Templates(ctx).Execute()

/templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Templates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Templates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Templates`: TemplatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Templates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesRequest struct via the builder pattern


### Return type

[**TemplatesResponse**](TemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesDelete

> TemplatesDelete(ctx, templateId).Execute()

/templates/{template_id}/delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	templateId := "templateId_example" // string | The template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.TemplatesDelete(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TemplatesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesFeatured

> TemplatesResponse TemplatesFeatured(ctx).Execute()

/templates/featured



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TemplatesFeatured(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TemplatesFeatured``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesFeatured`: TemplatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TemplatesFeatured`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesFeaturedRequest struct via the builder pattern


### Return type

[**TemplatesResponse**](TemplatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesInfo

> Template TemplatesInfo(ctx, templateId).Execute()

/templates/{template_id}/info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	templateId := "templateId_example" // string | The template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TemplatesInfo(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TemplatesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesInfo`: Template
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TemplatesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesSave

> TemplateCreateResponse TemplatesSave(ctx).TemplateSaveRequest(templateSaveRequest).Execute()

/templates/save



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	templateSaveRequest := *openapiclient.NewTemplateSaveRequest("My Template") // TemplateSaveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TemplatesSave(context.Background()).TemplateSaveRequest(templateSaveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TemplatesSave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplatesSave`: TemplateCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TemplatesSave`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateSaveRequest** | [**TemplateSaveRequest**](TemplateSaveRequest.md) |  | 

### Return type

[**TemplateCreateResponse**](TemplateCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdate

> TemplatesUpdate(ctx, templateId).TemplateSaveRequest(templateSaveRequest).Execute()

/templates/{template_id}/update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	templateId := "templateId_example" // string | The template ID
	templateSaveRequest := *openapiclient.NewTemplateSaveRequest("My Template") // TemplateSaveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.TemplatesUpdate(context.Background(), templateId).TemplateSaveRequest(templateSaveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TemplatesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateSaveRequest** | [**TemplateSaveRequest**](TemplateSaveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Volumes

> VolumesResponse Volumes(ctx).Execute()

/volumes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Volumes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Volumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Volumes`: VolumesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Volumes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesRequest struct via the builder pattern


### Return type

[**VolumesResponse**](VolumesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesCreate

> CreateVolumeResponse VolumesCreate(ctx).CreateVolumeRequest(createVolumeRequest).Execute()

/volumes/create



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	createVolumeRequest := *openapiclient.NewCreateVolumeRequest(openapiclient.Cloud("lambdalabs"), "canada-1", int32(100), "My storage volume") // CreateVolumeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VolumesCreate(context.Background()).CreateVolumeRequest(createVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VolumesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesCreate`: CreateVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VolumesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md) |  | 

### Return type

[**CreateVolumeResponse**](CreateVolumeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesDelete

> VolumesDelete(ctx, id).Execute()

/volumes/{id}/delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The volume id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.VolumesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VolumesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The volume id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesInfo

> Volume VolumesInfo(ctx, id).Execute()

/volumes/{id}/info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // string | The volume id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VolumesInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VolumesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesInfo`: Volume
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VolumesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The volume id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesTypes

> VolumesTypesResponse VolumesTypes(ctx).Execute()

/volumes/types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/brevdev/cloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VolumesTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VolumesTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesTypes`: VolumesTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VolumesTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesTypesRequest struct via the builder pattern


### Return type

[**VolumesTypesResponse**](VolumesTypesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

