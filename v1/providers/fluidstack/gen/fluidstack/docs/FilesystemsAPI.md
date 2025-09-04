# \FilesystemsAPI

All URIs are relative to *https://&lt;region&gt;.atlas.fluidstack.io/api/v1alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilesystem**](FilesystemsAPI.md#CreateFilesystem) | **Post** /filesystems | Create a filesystem
[**DeleteFilesystem**](FilesystemsAPI.md#DeleteFilesystem) | **Delete** /filesystems/{id} | Delete a filesystem
[**GetFilesystem**](FilesystemsAPI.md#GetFilesystem) | **Get** /filesystems/{id} | Get a filesystem
[**ListFilesystems**](FilesystemsAPI.md#ListFilesystems) | **Get** /filesystems | List all filesystems



## CreateFilesystem

> Filesystem CreateFilesystem(ctx).XPROJECTID(xPROJECTID).FilesystemsPostRequest(filesystemsPostRequest).Execute()

Create a filesystem

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	filesystemsPostRequest := *openapiclient.NewFilesystemsPostRequest("fs-1", "512Gi") // FilesystemsPostRequest | Details of the filesystem to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemsAPI.CreateFilesystem(context.Background()).XPROJECTID(xPROJECTID).FilesystemsPostRequest(filesystemsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemsAPI.CreateFilesystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFilesystem`: Filesystem
	fmt.Fprintf(os.Stdout, "Response from `FilesystemsAPI.CreateFilesystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilesystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 
 **filesystemsPostRequest** | [**FilesystemsPostRequest**](FilesystemsPostRequest.md) | Details of the filesystem to create | 

### Return type

[**Filesystem**](Filesystem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFilesystem

> DeleteFilesystem(ctx, id).XPROJECTID(xPROJECTID).Execute()

Delete a filesystem

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the filesystem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesystemsAPI.DeleteFilesystem(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemsAPI.DeleteFilesystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the filesystem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilesystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesystem

> Filesystem GetFilesystem(ctx, id).XPROJECTID(xPROJECTID).Execute()

Get a filesystem

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the filesystem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemsAPI.GetFilesystem(context.Background(), id).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemsAPI.GetFilesystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesystem`: Filesystem
	fmt.Fprintf(os.Stdout, "Response from `FilesystemsAPI.GetFilesystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the filesystem | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 


### Return type

[**Filesystem**](Filesystem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilesystems

> []Filesystem ListFilesystems(ctx).XPROJECTID(xPROJECTID).Execute()

List all filesystems

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
	xPROJECTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project identifier passed as a header

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesystemsAPI.ListFilesystems(context.Background()).XPROJECTID(xPROJECTID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesystemsAPI.ListFilesystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFilesystems`: []Filesystem
	fmt.Fprintf(os.Stdout, "Response from `FilesystemsAPI.ListFilesystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFilesystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPROJECTID** | **string** | Project identifier passed as a header | 

### Return type

[**[]Filesystem**](Filesystem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

