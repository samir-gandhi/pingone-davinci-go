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


// ApplicationsApiService ApplicationsApi service
type ApplicationsApiService service

type ApiV1AppsGetRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
}

func (r ApiV1AppsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsGetExecute(r)
}

/*
V1AppsGet READ ALL Apps

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1AppsGetRequest
*/
func (a *ApplicationsApiService) V1AppsGet(ctx context.Context) ApiV1AppsGetRequest {
	return ApiV1AppsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsGetExecute(r ApiV1AppsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps"

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

type ApiV1AppsPostRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1AppsPostRequest) ContentType(contentType string) ApiV1AppsPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1AppsPostRequest) Body(body map[string]interface{}) ApiV1AppsPostRequest {
	r.body = &body
	return r
}

func (r ApiV1AppsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsPostExecute(r)
}

/*
V1AppsPost CREATE App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1AppsPostRequest
*/
func (a *ApplicationsApiService) V1AppsPost(ctx context.Context) ApiV1AppsPostRequest {
	return ApiV1AppsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsPostExecute(r ApiV1AppsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps"

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

type ApiV1AppsSkAppIdDeleteRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	skAppId string
	contentType *string
}

func (r ApiV1AppsSkAppIdDeleteRequest) ContentType(contentType string) ApiV1AppsSkAppIdDeleteRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1AppsSkAppIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsSkAppIdDeleteExecute(r)
}

/*
V1AppsSkAppIdDelete DELETE App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skAppId
 @return ApiV1AppsSkAppIdDeleteRequest
*/
func (a *ApplicationsApiService) V1AppsSkAppIdDelete(ctx context.Context, skAppId string) ApiV1AppsSkAppIdDeleteRequest {
	return ApiV1AppsSkAppIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		skAppId: skAppId,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsSkAppIdDeleteExecute(r ApiV1AppsSkAppIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsSkAppIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps/{skAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skAppId"+"}", url.PathEscape(parameterToString(r.skAppId, "")), -1)

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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiV1AppsSkAppIdGetRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	skAppId string
}

func (r ApiV1AppsSkAppIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsSkAppIdGetExecute(r)
}

/*
V1AppsSkAppIdGet READ One App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skAppId
 @return ApiV1AppsSkAppIdGetRequest
*/
func (a *ApplicationsApiService) V1AppsSkAppIdGet(ctx context.Context, skAppId string) ApiV1AppsSkAppIdGetRequest {
	return ApiV1AppsSkAppIdGetRequest{
		ApiService: a,
		ctx: ctx,
		skAppId: skAppId,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsSkAppIdGetExecute(r ApiV1AppsSkAppIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsSkAppIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps/{skAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skAppId"+"}", url.PathEscape(parameterToString(r.skAppId, "")), -1)

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

type ApiV1AppsSkAppIdPolicyPostRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	skAppId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1AppsSkAppIdPolicyPostRequest) ContentType(contentType string) ApiV1AppsSkAppIdPolicyPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1AppsSkAppIdPolicyPostRequest) Body(body map[string]interface{}) ApiV1AppsSkAppIdPolicyPostRequest {
	r.body = &body
	return r
}

func (r ApiV1AppsSkAppIdPolicyPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsSkAppIdPolicyPostExecute(r)
}

/*
V1AppsSkAppIdPolicyPost CREATE App Flow Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skAppId
 @return ApiV1AppsSkAppIdPolicyPostRequest
*/
func (a *ApplicationsApiService) V1AppsSkAppIdPolicyPost(ctx context.Context, skAppId string) ApiV1AppsSkAppIdPolicyPostRequest {
	return ApiV1AppsSkAppIdPolicyPostRequest{
		ApiService: a,
		ctx: ctx,
		skAppId: skAppId,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsSkAppIdPolicyPostExecute(r ApiV1AppsSkAppIdPolicyPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsSkAppIdPolicyPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps/{skAppId}/policy"
	localVarPath = strings.Replace(localVarPath, "{"+"skAppId"+"}", url.PathEscape(parameterToString(r.skAppId, "")), -1)

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

type ApiV1AppsSkAppIdPutRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	skAppId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1AppsSkAppIdPutRequest) ContentType(contentType string) ApiV1AppsSkAppIdPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1AppsSkAppIdPutRequest) Body(body map[string]interface{}) ApiV1AppsSkAppIdPutRequest {
	r.body = &body
	return r
}

func (r ApiV1AppsSkAppIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AppsSkAppIdPutExecute(r)
}

/*
V1AppsSkAppIdPut UPDATE App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skAppId
 @return ApiV1AppsSkAppIdPutRequest
*/
func (a *ApplicationsApiService) V1AppsSkAppIdPut(ctx context.Context, skAppId string) ApiV1AppsSkAppIdPutRequest {
	return ApiV1AppsSkAppIdPutRequest{
		ApiService: a,
		ctx: ctx,
		skAppId: skAppId,
	}
}

// Execute executes the request
func (a *ApplicationsApiService) V1AppsSkAppIdPutExecute(r ApiV1AppsSkAppIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.V1AppsSkAppIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/apps/{skAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skAppId"+"}", url.PathEscape(parameterToString(r.skAppId, "")), -1)

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
