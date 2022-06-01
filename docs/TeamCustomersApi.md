# \TeamCustomersApi

All URIs are relative to *http://}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CompanySkCompanyIdCustomersGet**](TeamCustomersApi.md#V1CompanySkCompanyIdCustomersGet) | **Get** /v1/company/{skCompanyId}/customers | READ ALL Customers
[**V1CompanySkCompanyIdCustomersPost**](TeamCustomersApi.md#V1CompanySkCompanyIdCustomersPost) | **Post** /v1/company/{skCompanyId}/customers | CREATE Customer
[**V1CompanySkCompanyIdCustomersSkCustomerIdDelete**](TeamCustomersApi.md#V1CompanySkCompanyIdCustomersSkCustomerIdDelete) | **Delete** /v1/company/{skCompanyId}/customers/{skCustomerId} | DELETE Customer
[**V1CompanySkCompanyIdCustomersSkCustomerIdPut**](TeamCustomersApi.md#V1CompanySkCompanyIdCustomersSkCustomerIdPut) | **Put** /v1/company/{skCompanyId}/customers/{skCustomerId} | UPDATE Customer



## V1CompanySkCompanyIdCustomersGet

> V1CompanySkCompanyIdCustomersGet(ctx, skCompanyId).Execute()

READ ALL Customers

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
    resp, r, err := apiClient.TeamCustomersApi.V1CompanySkCompanyIdCustomersGet(context.Background(), skCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamCustomersApi.V1CompanySkCompanyIdCustomersGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdCustomersGetRequest struct via the builder pattern


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


## V1CompanySkCompanyIdCustomersPost

> V1CompanySkCompanyIdCustomersPost(ctx, skCompanyId).Body(body).Execute()

CREATE Customer

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamCustomersApi.V1CompanySkCompanyIdCustomersPost(context.Background(), skCompanyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamCustomersApi.V1CompanySkCompanyIdCustomersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdCustomersPostRequest struct via the builder pattern


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


## V1CompanySkCompanyIdCustomersSkCustomerIdDelete

> V1CompanySkCompanyIdCustomersSkCustomerIdDelete(ctx, skCompanyId, skCustomerId).Execute()

DELETE Customer

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
    skCustomerId := "skCustomerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamCustomersApi.V1CompanySkCompanyIdCustomersSkCustomerIdDelete(context.Background(), skCompanyId, skCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamCustomersApi.V1CompanySkCompanyIdCustomersSkCustomerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skCompanyId** | **string** |  | 
**skCustomerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest struct via the builder pattern


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


## V1CompanySkCompanyIdCustomersSkCustomerIdPut

> V1CompanySkCompanyIdCustomersSkCustomerIdPut(ctx, skCompanyId, skCustomerId).Body(body).Execute()

UPDATE Customer

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
    skCustomerId := "skCustomerId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamCustomersApi.V1CompanySkCompanyIdCustomersSkCustomerIdPut(context.Background(), skCompanyId, skCustomerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamCustomersApi.V1CompanySkCompanyIdCustomersSkCustomerIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skCompanyId** | **string** |  | 
**skCustomerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest struct via the builder pattern


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

