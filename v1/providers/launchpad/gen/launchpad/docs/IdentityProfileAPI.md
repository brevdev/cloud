# \IdentityProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1IdentityProfilePartialUpdate**](IdentityProfileAPI.md#V1IdentityProfilePartialUpdate) | **Patch** /v1/identity/profile/ | 
[**V1IdentityProfileRetrieve**](IdentityProfileAPI.md#V1IdentityProfileRetrieve) | **Get** /v1/identity/profile/ | 



## V1IdentityProfilePartialUpdate

> Profile V1IdentityProfilePartialUpdate(ctx).Profile(profile).Execute()



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
	profile := *openapiclient.NewProfile(time.Now(), "DisplayName_example", []string{"Groups_example"}, "Id_example", time.Now(), "Username_example") // Profile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProfileAPI.V1IdentityProfilePartialUpdate(context.Background()).Profile(profile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfileAPI.V1IdentityProfilePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IdentityProfilePartialUpdate`: Profile
	fmt.Fprintf(os.Stdout, "Response from `IdentityProfileAPI.V1IdentityProfilePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IdentityProfilePartialUpdateRequest struct via the builder pattern


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


## V1IdentityProfileRetrieve

> Profile V1IdentityProfileRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.IdentityProfileAPI.V1IdentityProfileRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProfileAPI.V1IdentityProfileRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IdentityProfileRetrieve`: Profile
	fmt.Fprintf(os.Stdout, "Response from `IdentityProfileAPI.V1IdentityProfileRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1IdentityProfileRetrieveRequest struct via the builder pattern


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

