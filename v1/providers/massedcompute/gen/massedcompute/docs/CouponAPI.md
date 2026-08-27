# \CouponAPI

All URIs are relative to *https://vm.massedcompute.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CouponAcceptedProductsPost**](CouponAPI.md#CouponAcceptedProductsPost) | **Post** /coupon/accepted-products | Retrieve products that a coupon is valid for.
[**CouponInformationPost**](CouponAPI.md#CouponInformationPost) | **Post** /coupon/information | Retrieve information about a coupon.



## CouponAcceptedProductsPost

> RetrieveAcceptProductsV1 CouponAcceptedProductsPost(ctx).CouponInformationPostRequest(couponInformationPostRequest).Execute()

Retrieve products that a coupon is valid for.



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
	couponInformationPostRequest := *openapiclient.NewCouponInformationPostRequest() // CouponInformationPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponAPI.CouponAcceptedProductsPost(context.Background()).CouponInformationPostRequest(couponInformationPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponAPI.CouponAcceptedProductsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponAcceptedProductsPost`: RetrieveAcceptProductsV1
	fmt.Fprintf(os.Stdout, "Response from `CouponAPI.CouponAcceptedProductsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponAcceptedProductsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponInformationPostRequest** | [**CouponInformationPostRequest**](CouponInformationPostRequest.md) |  |

### Return type

[**RetrieveAcceptProductsV1**](RetrieveAcceptProductsV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponInformationPost

> RetrieveCouponInformationV1 CouponInformationPost(ctx).CouponInformationPostRequest(couponInformationPostRequest).Execute()

Retrieve information about a coupon.



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
	couponInformationPostRequest := *openapiclient.NewCouponInformationPostRequest() // CouponInformationPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponAPI.CouponInformationPost(context.Background()).CouponInformationPostRequest(couponInformationPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponAPI.CouponInformationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponInformationPost`: RetrieveCouponInformationV1
	fmt.Fprintf(os.Stdout, "Response from `CouponAPI.CouponInformationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponInformationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponInformationPostRequest** | [**CouponInformationPostRequest**](CouponInformationPostRequest.md) |  |

### Return type

[**RetrieveCouponInformationV1**](RetrieveCouponInformationV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

