/*
DaVinci Copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// EnvironmentsApiService EnvironmentsApi service
type EnvironmentsApiService service

type ApiV1CompanySkCompanyIdStatsGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	skCompanyId string
}

func (r ApiV1CompanySkCompanyIdStatsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdStatsGetExecute(r)
}

/*
V1CompanySkCompanyIdStatsGet GET Current Environment Statistics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @return ApiV1CompanySkCompanyIdStatsGetRequest
*/
func (a *EnvironmentsApiService) V1CompanySkCompanyIdStatsGet(ctx context.Context, skCompanyId string) ApiV1CompanySkCompanyIdStatsGetRequest {
	return ApiV1CompanySkCompanyIdStatsGetRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
	}
}

// Execute executes the request
func (a *EnvironmentsApiService) V1CompanySkCompanyIdStatsGetExecute(r ApiV1CompanySkCompanyIdStatsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.V1CompanySkCompanyIdStatsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/stats"
	localVarPath = strings.Replace(localVarPath, "{"+"skCompanyId"+"}", url.PathEscape(parameterToString(r.skCompanyId, "")), -1)

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

type ApiV1CompanySkCompanyIdSwitchPutRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	skCompanyId string
	contentType *string
	body *map[string]interface{}
}

func (r ApiV1CompanySkCompanyIdSwitchPutRequest) ContentType(contentType string) ApiV1CompanySkCompanyIdSwitchPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiV1CompanySkCompanyIdSwitchPutRequest) Body(body map[string]interface{}) ApiV1CompanySkCompanyIdSwitchPutRequest {
	r.body = &body
	return r
}

func (r ApiV1CompanySkCompanyIdSwitchPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdSwitchPutExecute(r)
}

/*
V1CompanySkCompanyIdSwitchPut SET Current Environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @return ApiV1CompanySkCompanyIdSwitchPutRequest
*/
func (a *EnvironmentsApiService) V1CompanySkCompanyIdSwitchPut(ctx context.Context, skCompanyId string) ApiV1CompanySkCompanyIdSwitchPutRequest {
	return ApiV1CompanySkCompanyIdSwitchPutRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
	}
}

// Execute executes the request
func (a *EnvironmentsApiService) V1CompanySkCompanyIdSwitchPutExecute(r ApiV1CompanySkCompanyIdSwitchPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.V1CompanySkCompanyIdSwitchPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/switch"
	localVarPath = strings.Replace(localVarPath, "{"+"skCompanyId"+"}", url.PathEscape(parameterToString(r.skCompanyId, "")), -1)

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

type ApiV1CustomersMeGetRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
}

func (r ApiV1CustomersMeGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CustomersMeGetExecute(r)
}

/*
V1CustomersMeGet READ ALL Environments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1CustomersMeGetRequest
*/
func (a *EnvironmentsApiService) V1CustomersMeGet(ctx context.Context) ApiV1CustomersMeGetRequest {
	return ApiV1CustomersMeGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EnvironmentsApiService) V1CustomersMeGetExecute(r ApiV1CustomersMeGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.V1CustomersMeGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/customers/me"

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
