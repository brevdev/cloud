# \CatalogPersonaChoicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CatalogPersonaChoicesCreate**](CatalogPersonaChoicesAPI.md#V1CatalogPersonaChoicesCreate) | **Post** /v1/catalog/persona-choices/ | 
[**V1CatalogPersonaChoicesList**](CatalogPersonaChoicesAPI.md#V1CatalogPersonaChoicesList) | **Get** /v1/catalog/persona-choices/ | 



## V1CatalogPersonaChoicesCreate

> PersonaChoice V1CatalogPersonaChoicesCreate(ctx).PersonaChoice(personaChoice).Execute()



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
	personaChoice := *openapiclient.NewPersonaChoice(time.Now(), int32(123), time.Now(), "Name_example") // PersonaChoice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesCreate(context.Background()).PersonaChoice(personaChoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogPersonaChoicesCreate`: PersonaChoice
	fmt.Fprintf(os.Stdout, "Response from `CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogPersonaChoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personaChoice** | [**PersonaChoice**](PersonaChoice.md) |  | 

### Return type

[**PersonaChoice**](PersonaChoice.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CatalogPersonaChoicesList

> PaginatedPersonaChoiceList V1CatalogPersonaChoicesList(ctx).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	fields := "fields_example" // string | Include only the specified fields in the response (optional)
	omit := "omit_example" // string | Exclude the specified fields in the response (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | Search for personachoices by id, name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesList(context.Background()).Fields(fields).Omit(omit).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CatalogPersonaChoicesList`: PaginatedPersonaChoiceList
	fmt.Fprintf(os.Stdout, "Response from `CatalogPersonaChoicesAPI.V1CatalogPersonaChoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CatalogPersonaChoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Include only the specified fields in the response | 
 **omit** | **string** | Exclude the specified fields in the response | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | Search for personachoices by id, name | 

### Return type

[**PaginatedPersonaChoiceList**](PaginatedPersonaChoiceList.md)

### Authorization

[TokenAuthentication](../README.md#TokenAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

