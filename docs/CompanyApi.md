# \CompanyApi

All URIs are relative to *http://}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CompanyPost**](CompanyApi.md#V1CompanyPost) | **Post** /v1/company | CREATE Company
[**V1CompanySkCompanyIdGet**](CompanyApi.md#V1CompanySkCompanyIdGet) | **Get** /v1/company/{skCompanyId} | READ Company



## V1CompanyPost

> V1CompanyPost(ctx).Body(body).Execute()

CREATE Company

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompanyApi.V1CompanyPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.V1CompanyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CompanySkCompanyIdGet

> V1CompanySkCompanyIdGet(ctx, skCompanyId).Execute()

READ Company

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    skCompanyId := "skCompanyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompanyApi.V1CompanySkCompanyIdGet(context.Background(), skCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.V1CompanySkCompanyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skCompanyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

