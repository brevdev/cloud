# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AuthLoginCreate**](AuthAPI.md#V1AuthLoginCreate) | **Post** /v1/auth/login/ | 
[**V1AuthLogoutCreate**](AuthAPI.md#V1AuthLogoutCreate) | **Post** /v1/auth/logout/ | 
[**V1AuthPasswordChangeCreate**](AuthAPI.md#V1AuthPasswordChangeCreate) | **Post** /v1/auth/password-change/ | 
[**V1AuthRedirectRetrieve**](AuthAPI.md#V1AuthRedirectRetrieve) | **Get** /v1/auth/redirect/ | 



## V1AuthLoginCreate

> AuthToken V1AuthLoginCreate(ctx).AuthCode(authCode).Password(password).Username(username).Execute()





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
	authCode := "authCode_example" // string |  (optional)
	password := "password_example" // string |  (optional)
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.V1AuthLoginCreate(context.Background()).AuthCode(authCode).Password(password).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthLoginCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthLoginCreate`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthLoginCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthLoginCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCode** | **string** |  | 
 **password** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthLogoutCreate

> Logout V1AuthLogoutCreate(ctx).Execute()





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
	resp, r, err := apiClient.AuthAPI.V1AuthLogoutCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthLogoutCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthLogoutCreate`: Logout
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthLogoutCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthLogoutCreateRequest struct via the builder pattern


### Return type

[**Logout**](Logout.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthPasswordChangeCreate

> PasswordChange V1AuthPasswordChangeCreate(ctx).Username(username).Password(password).NewPassword(newPassword).Result(result).Execute()





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
	username := "username_example" // string | 
	password := "password_example" // string | 
	newPassword := "newPassword_example" // string | 
	result := "result_example" // string |  (default to "Password changed successfully.")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.V1AuthPasswordChangeCreate(context.Background()).Username(username).Password(password).NewPassword(newPassword).Result(result).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthPasswordChangeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthPasswordChangeCreate`: PasswordChange
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthPasswordChangeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthPasswordChangeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **newPassword** | **string** |  | 
 **result** | **string** |  | [default to &quot;Password changed successfully.&quot;]

### Return type

[**PasswordChange**](PasswordChange.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthRedirectRetrieve

> AuthRedirect V1AuthRedirectRetrieve(ctx).AsJson(asJson).Execute()





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
	asJson := true // bool | If \"false\" or unset, receive an HTTP 302 with the OAuth URL returned in the Location header. If \"true\", receive an HTTP 200 with the OAuth URL returned in the JSON payload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.V1AuthRedirectRetrieve(context.Background()).AsJson(asJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.V1AuthRedirectRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AuthRedirectRetrieve`: AuthRedirect
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.V1AuthRedirectRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthRedirectRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asJson** | **bool** | If \&quot;false\&quot; or unset, receive an HTTP 302 with the OAuth URL returned in the Location header. If \&quot;true\&quot;, receive an HTTP 200 with the OAuth URL returned in the JSON payload. | 

### Return type

[**AuthRedirect**](AuthRedirect.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

