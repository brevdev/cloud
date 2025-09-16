# \IdentityProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProfilePartialUpdate**](IdentityProfileAPI.md#IdentityProfilePartialUpdate) | **Patch** /v1/identity/profile/ | 
[**IdentityProfileRetrieve**](IdentityProfileAPI.md#IdentityProfileRetrieve) | **Get** /v1/identity/profile/ | 



## IdentityProfilePartialUpdate

> Profile IdentityProfilePartialUpdate(ctx).Profile(profile).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/brevdev/cloud"
)

func main() {
	profile := *openapiclient.NewProfile(time.Now(), []string{"Groups_example"}, "Id_example", time.Now(), "Username_example") // Profile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProfileAPI.IdentityProfilePartialUpdate(context.Background()).Profile(profile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfileAPI.IdentityProfilePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProfilePartialUpdate`: Profile
	fmt.Fprintf(os.Stdout, "Response from `IdentityProfileAPI.IdentityProfilePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProfilePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | [**Profile**](Profile.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProfileRetrieve

> Profile IdentityProfileRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.IdentityProfileAPI.IdentityProfileRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfileAPI.IdentityProfileRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentityProfileRetrieve`: Profile
	fmt.Fprintf(os.Stdout, "Response from `IdentityProfileAPI.IdentityProfileRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProfileRetrieveRequest struct via the builder pattern


### Return type

[**Profile**](Profile.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

