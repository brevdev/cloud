# \AccountAPI

All URIs are relative to *https://vm.massedcompute.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountBillingGet**](AccountAPI.md#AccountBillingGet) | **Get** /account/billing | Retrieve billing information.
[**AccountTokenValidationPost**](AccountAPI.md#AccountTokenValidationPost) | **Post** /account/token/validation | Validate an API token.



## AccountBillingGet

> RetrieveBillingInformationV1 AccountBillingGet(ctx).Execute()

Retrieve billing information.



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
	resp, r, err := apiClient.AccountAPI.AccountBillingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountBillingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountBillingGet`: RetrieveBillingInformationV1
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountBillingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountBillingGetRequest struct via the builder pattern


### Return type

[**RetrieveBillingInformationV1**](RetrieveBillingInformationV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountTokenValidationPost

> AccountTokenValidationPost200Response AccountTokenValidationPost(ctx).Execute()

Validate an API token.



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
	resp, r, err := apiClient.AccountAPI.AccountTokenValidationPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountTokenValidationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountTokenValidationPost`: AccountTokenValidationPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountTokenValidationPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountTokenValidationPostRequest struct via the builder pattern


### Return type

[**AccountTokenValidationPost200Response**](AccountTokenValidationPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

