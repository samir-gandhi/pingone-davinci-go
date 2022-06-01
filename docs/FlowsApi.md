# \FlowsApi

All URIs are relative to *http://}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FlowsAppsFlowsVersionsGet**](FlowsApi.md#V1FlowsAppsFlowsVersionsGet) | **Get** /v1/flows/apps/flowsVersions | READ Flow Version Details
[**V1FlowsGet**](FlowsApi.md#V1FlowsGet) | **Get** /v1/flows | READ ALL Flows
[**V1FlowsPost**](FlowsApi.md#V1FlowsPost) | **Post** /v1/flows | CREATE Flow
[**V1FlowsSkFlowIdDelete**](FlowsApi.md#V1FlowsSkFlowIdDelete) | **Delete** /v1/flows/{skFlowId} | DELETE Flow
[**V1FlowsSkFlowIdDeployPut**](FlowsApi.md#V1FlowsSkFlowIdDeployPut) | **Put** /v1/flows/{skFlowId}/deploy | DEPLOY Flow
[**V1FlowsSkFlowIdGet**](FlowsApi.md#V1FlowsSkFlowIdGet) | **Get** /v1/flows/{skFlowId} | READ One Flow
[**V1FlowsSkFlowIdPut**](FlowsApi.md#V1FlowsSkFlowIdPut) | **Put** /v1/flows/{skFlowId} | UPDATE Flow
[**V1FlowsSkFlowIdVersionsGet**](FlowsApi.md#V1FlowsSkFlowIdVersionsGet) | **Get** /v1/flows/{skFlowId}/versions | READ Flow Versions



## V1FlowsAppsFlowsVersionsGet

> V1FlowsAppsFlowsVersionsGet(ctx).FlowsVersions(flowsVersions).Execute()

READ Flow Version Details

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
    flowsVersions := "{{skFlowsVersion}}" // string | Base 64 Encoded format - [{\"flowId\":\"<your flow id>\", \"versionId\":<your version>}] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsAppsFlowsVersionsGet(context.Background()).FlowsVersions(flowsVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsAppsFlowsVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsAppsFlowsVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowsVersions** | **string** | Base 64 Encoded format - [{\&quot;flowId\&quot;:\&quot;&lt;your flow id&gt;\&quot;, \&quot;versionId\&quot;:&lt;your version&gt;}] | 

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


## V1FlowsGet

> V1FlowsGet(ctx).Execute()

READ ALL Flows

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsGetRequest struct via the builder pattern


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


## V1FlowsPost

> V1FlowsPost(ctx).ContentType(contentType).Body(body).Execute()

CREATE Flow

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsPost(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
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


## V1FlowsSkFlowIdDelete

> V1FlowsSkFlowIdDelete(ctx, skFlowId).Execute()

DELETE Flow

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
    skFlowId := "skFlowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsSkFlowIdDelete(context.Background(), skFlowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsSkFlowIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsSkFlowIdDeleteRequest struct via the builder pattern


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


## V1FlowsSkFlowIdDeployPut

> V1FlowsSkFlowIdDeployPut(ctx, skFlowId).ContentType(contentType).Body(body).Execute()

DEPLOY Flow

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
    skFlowId := "skFlowId_example" // string | 
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsSkFlowIdDeployPut(context.Background(), skFlowId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsSkFlowIdDeployPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsSkFlowIdDeployPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
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


## V1FlowsSkFlowIdGet

> V1FlowsSkFlowIdGet(ctx, skFlowId).Execute()

READ One Flow

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
    skFlowId := "skFlowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsSkFlowIdGet(context.Background(), skFlowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsSkFlowIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsSkFlowIdGetRequest struct via the builder pattern


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


## V1FlowsSkFlowIdPut

> V1FlowsSkFlowIdPut(ctx, skFlowId).ContentType(contentType).Body(body).Execute()

UPDATE Flow

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
    skFlowId := "skFlowId_example" // string | 
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsSkFlowIdPut(context.Background(), skFlowId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsSkFlowIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsSkFlowIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
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


## V1FlowsSkFlowIdVersionsGet

> V1FlowsSkFlowIdVersionsGet(ctx, skFlowId).Execute()

READ Flow Versions

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
    skFlowId := "skFlowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowsApi.V1FlowsSkFlowIdVersionsGet(context.Background(), skFlowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowsApi.V1FlowsSkFlowIdVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FlowsSkFlowIdVersionsGetRequest struct via the builder pattern


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

