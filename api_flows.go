/*
DaVinci Copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingonedavinci

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// FlowsApiService FlowsApi service
type FlowsApiService service

type ApiV1FlowsAppsFlowsVersionsGetRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	flowsVersions *string
}

// Base 64 Encoded format - [{\&quot;flowId\&quot;:\&quot;&lt;your flow id&gt;\&quot;, \&quot;versionId\&quot;:&lt;your version&gt;}]
func (r ApiV1FlowsAppsFlowsVersionsGetRequest) FlowsVersions(flowsVersions string) ApiV1FlowsAppsFlowsVersionsGetRequest {
	r.flowsVersions = &flowsVersions
	return r
}

func (r ApiV1FlowsAppsFlowsVersionsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsAppsFlowsVersionsGetExecute(r)
}

/*
V1FlowsAppsFlowsVersionsGet READ Flow Version Details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1FlowsAppsFlowsVersionsGetRequest
*/
func (a *FlowsApiService) V1FlowsAppsFlowsVersionsGet(ctx context.Context) ApiV1FlowsAppsFlowsVersionsGetRequest {
	return ApiV1FlowsAppsFlowsVersionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsAppsFlowsVersionsGetExecute(r ApiV1FlowsAppsFlowsVersionsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsAppsFlowsVersionsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/apps/flowsVersions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.flowsVersions != nil {
		localVarQueryParams.Add("flowsVersions", parameterToString(*r.flowsVersions, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsGetRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
}

func (r ApiV1FlowsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsGetExecute(r)
}

/*
V1FlowsGet READ ALL Flows

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1FlowsGetRequest
*/
func (a *FlowsApiService) V1FlowsGet(ctx context.Context) ApiV1FlowsGetRequest {
	return ApiV1FlowsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsGetExecute(r ApiV1FlowsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsPostRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1FlowsPostRequest) ContentType(contentType string) ApiV1FlowsPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1FlowsPostRequest) Body(body map[string]interface{}) ApiV1FlowsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1FlowsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsPostExecute(r)
}

/*
V1FlowsPost CREATE Flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1FlowsPostRequest
*/
func (a *FlowsApiService) V1FlowsPost(ctx context.Context) ApiV1FlowsPostRequest {
	return ApiV1FlowsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsPostExecute(r ApiV1FlowsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsSkFlowIdDeleteRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	skFlowId string
}

func (r ApiV1FlowsSkFlowIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsSkFlowIdDeleteExecute(r)
}

/*
V1FlowsSkFlowIdDelete DELETE Flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skFlowId
 @return ApiV1FlowsSkFlowIdDeleteRequest
*/
func (a *FlowsApiService) V1FlowsSkFlowIdDelete(ctx context.Context, skFlowId string) ApiV1FlowsSkFlowIdDeleteRequest {
	return ApiV1FlowsSkFlowIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		skFlowId: skFlowId,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsSkFlowIdDeleteExecute(r ApiV1FlowsSkFlowIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsSkFlowIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/{skFlowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skFlowId"+"}", url.PathEscape(parameterToString(r.skFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsSkFlowIdDeployPutRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	skFlowId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1FlowsSkFlowIdDeployPutRequest) ContentType(contentType string) ApiV1FlowsSkFlowIdDeployPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1FlowsSkFlowIdDeployPutRequest) Body(body map[string]interface{}) ApiV1FlowsSkFlowIdDeployPutRequest {
	r.body = &body
	return r
}

func (r ApiV1FlowsSkFlowIdDeployPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsSkFlowIdDeployPutExecute(r)
}

/*
V1FlowsSkFlowIdDeployPut DEPLOY Flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skFlowId
 @return ApiV1FlowsSkFlowIdDeployPutRequest
*/
func (a *FlowsApiService) V1FlowsSkFlowIdDeployPut(ctx context.Context, skFlowId string) ApiV1FlowsSkFlowIdDeployPutRequest {
	return ApiV1FlowsSkFlowIdDeployPutRequest{
		ApiService: a,
		ctx: ctx,
		skFlowId: skFlowId,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsSkFlowIdDeployPutExecute(r ApiV1FlowsSkFlowIdDeployPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsSkFlowIdDeployPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/{skFlowId}/deploy"
	localVarPath = strings.Replace(localVarPath, "{"+"skFlowId"+"}", url.PathEscape(parameterToString(r.skFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsSkFlowIdGetRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	skFlowId string
}

func (r ApiV1FlowsSkFlowIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsSkFlowIdGetExecute(r)
}

/*
V1FlowsSkFlowIdGet READ One Flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skFlowId
 @return ApiV1FlowsSkFlowIdGetRequest
*/
func (a *FlowsApiService) V1FlowsSkFlowIdGet(ctx context.Context, skFlowId string) ApiV1FlowsSkFlowIdGetRequest {
	return ApiV1FlowsSkFlowIdGetRequest{
		ApiService: a,
		ctx: ctx,
		skFlowId: skFlowId,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsSkFlowIdGetExecute(r ApiV1FlowsSkFlowIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsSkFlowIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/{skFlowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skFlowId"+"}", url.PathEscape(parameterToString(r.skFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsSkFlowIdPutRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	skFlowId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1FlowsSkFlowIdPutRequest) ContentType(contentType string) ApiV1FlowsSkFlowIdPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1FlowsSkFlowIdPutRequest) Body(body map[string]interface{}) ApiV1FlowsSkFlowIdPutRequest {
	r.body = &body
	return r
}

func (r ApiV1FlowsSkFlowIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsSkFlowIdPutExecute(r)
}

/*
V1FlowsSkFlowIdPut UPDATE Flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skFlowId
 @return ApiV1FlowsSkFlowIdPutRequest
*/
func (a *FlowsApiService) V1FlowsSkFlowIdPut(ctx context.Context, skFlowId string) ApiV1FlowsSkFlowIdPutRequest {
	return ApiV1FlowsSkFlowIdPutRequest{
		ApiService: a,
		ctx: ctx,
		skFlowId: skFlowId,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsSkFlowIdPutExecute(r ApiV1FlowsSkFlowIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsSkFlowIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/{skFlowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skFlowId"+"}", url.PathEscape(parameterToString(r.skFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1FlowsSkFlowIdVersionsGetRequest struct {
	ctx context.Context
	ApiService *FlowsApiService
	skFlowId string
}

func (r ApiV1FlowsSkFlowIdVersionsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1FlowsSkFlowIdVersionsGetExecute(r)
}

/*
V1FlowsSkFlowIdVersionsGet READ Flow Versions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skFlowId
 @return ApiV1FlowsSkFlowIdVersionsGetRequest
*/
func (a *FlowsApiService) V1FlowsSkFlowIdVersionsGet(ctx context.Context, skFlowId string) ApiV1FlowsSkFlowIdVersionsGetRequest {
	return ApiV1FlowsSkFlowIdVersionsGetRequest{
		ApiService: a,
		ctx: ctx,
		skFlowId: skFlowId,
	}
}

// Execute executes the request
func (a *FlowsApiService) V1FlowsSkFlowIdVersionsGetExecute(r ApiV1FlowsSkFlowIdVersionsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlowsApiService.V1FlowsSkFlowIdVersionsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/flows/{skFlowId}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"skFlowId"+"}", url.PathEscape(parameterToString(r.skFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
