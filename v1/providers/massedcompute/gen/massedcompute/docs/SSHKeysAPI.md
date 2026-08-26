# \SSHKeysAPI

All URIs are relative to *https://vm.massedcompute.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshKeysGet**](SSHKeysAPI.md#SshKeysGet) | **Get** /ssh-keys | Retrieve SSH keys associated with the account.
[**SshKeysIdDelete**](SSHKeysAPI.md#SshKeysIdDelete) | **Delete** /ssh-keys/{id} | Remove an SSH key from the account.
[**SshKeysPost**](SSHKeysAPI.md#SshKeysPost) | **Post** /ssh-keys | Add an SSH key to the account.



## SshKeysGet

> SSHKey SshKeysGet(ctx).Execute()

Retrieve SSH keys associated with the account.



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
	resp, r, err := apiClient.SSHKeysAPI.SshKeysGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysGet`: SSHKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeysGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysGetRequest struct via the builder pattern


### Return type

[**SSHKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysIdDelete

> SshKeysIdDelete200Response SshKeysIdDelete(ctx, id).Execute()

Remove an SSH key from the account.



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
	id := "id_example" // string | The unique identifier for the SSH key to be removed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.SshKeysIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeysIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysIdDelete`: SshKeysIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeysIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the SSH key to be removed |

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SshKeysIdDelete200Response**](SshKeysIdDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysPost

> POSTSSHKey SshKeysPost(ctx).SshKeysPostRequest(sshKeysPostRequest).Execute()

Add an SSH key to the account.



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
	sshKeysPostRequest := *openapiclient.NewSshKeysPostRequest("Name_example", "PublicKey_example") // SshKeysPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.SshKeysPost(context.Background()).SshKeysPostRequest(sshKeysPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.SshKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshKeysPost`: POSTSSHKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.SshKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeysPostRequest** | [**SshKeysPostRequest**](SshKeysPostRequest.md) |  |

### Return type

[**POSTSSHKey**](POSTSSHKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

