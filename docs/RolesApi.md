# \RolesApi

All URIs are relative to *http://}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CompanySkCompanyIdRolesGet**](RolesApi.md#V1CompanySkCompanyIdRolesGet) | **Get** /v1/company/{skCompanyId}/roles | READ Roles
[**V1CompanySkCompanyIdRolesPost**](RolesApi.md#V1CompanySkCompanyIdRolesPost) | **Post** /v1/company/{skCompanyId}/roles | CREATE Role
[**V1CompanySkCompanyIdRolesSkRoleNameDelete**](RolesApi.md#V1CompanySkCompanyIdRolesSkRoleNameDelete) | **Delete** /v1/company/{skCompanyId}/roles/{skRoleName} | DELETE Role
[**V1CompanySkCompanyIdRolesSkRoleNamePut**](RolesApi.md#V1CompanySkCompanyIdRolesSkRoleNamePut) | **Put** /v1/company/{skCompanyId}/roles/{skRoleName} | UPDATE Role



## V1CompanySkCompanyIdRolesGet

> V1CompanySkCompanyIdRolesGet(ctx, skCompanyId).Execute()

READ Roles

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
    resp, r, err := apiClient.RolesApi.V1CompanySkCompanyIdRolesGet(context.Background(), skCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.V1CompanySkCompanyIdRolesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdRolesGetRequest struct via the builder pattern


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


## V1CompanySkCompanyIdRolesPost

> V1CompanySkCompanyIdRolesPost(ctx, skCompanyId).Body(body).Execute()

CREATE Role

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
    resp, r, err := apiClient.RolesApi.V1CompanySkCompanyIdRolesPost(context.Background(), skCompanyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.V1CompanySkCompanyIdRolesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdRolesPostRequest struct via the builder pattern


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


## V1CompanySkCompanyIdRolesSkRoleNameDelete

> V1CompanySkCompanyIdRolesSkRoleNameDelete(ctx, skCompanyId, skRoleName).Execute()

DELETE Role

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
    skRoleName := "skRoleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.V1CompanySkCompanyIdRolesSkRoleNameDelete(context.Background(), skCompanyId, skRoleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.V1CompanySkCompanyIdRolesSkRoleNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skCompanyId** | **string** |  | 
**skRoleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdRolesSkRoleNameDeleteRequest struct via the builder pattern


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


## V1CompanySkCompanyIdRolesSkRoleNamePut

> V1CompanySkCompanyIdRolesSkRoleNamePut(ctx, skCompanyId, skRoleName).Body(body).Execute()

UPDATE Role

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
    skRoleName := "skRoleName_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.V1CompanySkCompanyIdRolesSkRoleNamePut(context.Background(), skCompanyId, skRoleName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.V1CompanySkCompanyIdRolesSkRoleNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skCompanyId** | **string** |  | 
**skRoleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CompanySkCompanyIdRolesSkRoleNamePutRequest struct via the builder pattern


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

