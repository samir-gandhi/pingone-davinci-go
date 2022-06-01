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


// TeamCustomersApiService TeamCustomersApi service
type TeamCustomersApiService service

type ApiV1CompanySkCompanyIdCustomersGetRequest struct {
	ctx context.Context
	ApiService *TeamCustomersApiService
	skCompanyId string
}

func (r ApiV1CompanySkCompanyIdCustomersGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdCustomersGetExecute(r)
}

/*
V1CompanySkCompanyIdCustomersGet READ ALL Customers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @return ApiV1CompanySkCompanyIdCustomersGetRequest
*/
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersGet(ctx context.Context, skCompanyId string) ApiV1CompanySkCompanyIdCustomersGetRequest {
	return ApiV1CompanySkCompanyIdCustomersGetRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
	}
}

// Execute executes the request
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersGetExecute(r ApiV1CompanySkCompanyIdCustomersGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamCustomersApiService.V1CompanySkCompanyIdCustomersGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/customers"
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

type ApiV1CompanySkCompanyIdCustomersPostRequest struct {
	ctx context.Context
	ApiService *TeamCustomersApiService
	skCompanyId string
	body *map[string]interface{}
}

func (r ApiV1CompanySkCompanyIdCustomersPostRequest) Body(body map[string]interface{}) ApiV1CompanySkCompanyIdCustomersPostRequest {
	r.body = &body
	return r
}

func (r ApiV1CompanySkCompanyIdCustomersPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdCustomersPostExecute(r)
}

/*
V1CompanySkCompanyIdCustomersPost CREATE Customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @return ApiV1CompanySkCompanyIdCustomersPostRequest
*/
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersPost(ctx context.Context, skCompanyId string) ApiV1CompanySkCompanyIdCustomersPostRequest {
	return ApiV1CompanySkCompanyIdCustomersPostRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
	}
}

// Execute executes the request
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersPostExecute(r ApiV1CompanySkCompanyIdCustomersPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamCustomersApiService.V1CompanySkCompanyIdCustomersPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/customers"
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

type ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest struct {
	ctx context.Context
	ApiService *TeamCustomersApiService
	skCompanyId string
	skCustomerId string
}

func (r ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdCustomersSkCustomerIdDeleteExecute(r)
}

/*
V1CompanySkCompanyIdCustomersSkCustomerIdDelete DELETE Customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @param skCustomerId
 @return ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest
*/
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersSkCustomerIdDelete(ctx context.Context, skCompanyId string, skCustomerId string) ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest {
	return ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
		skCustomerId: skCustomerId,
	}
}

// Execute executes the request
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersSkCustomerIdDeleteExecute(r ApiV1CompanySkCompanyIdCustomersSkCustomerIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamCustomersApiService.V1CompanySkCompanyIdCustomersSkCustomerIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/customers/{skCustomerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skCompanyId"+"}", url.PathEscape(parameterToString(r.skCompanyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skCustomerId"+"}", url.PathEscape(parameterToString(r.skCustomerId, "")), -1)

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

type ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest struct {
	ctx context.Context
	ApiService *TeamCustomersApiService
	skCompanyId string
	skCustomerId string
	body *map[string]interface{}
}

func (r ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest) Body(body map[string]interface{}) ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest {
	r.body = &body
	return r
}

func (r ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CompanySkCompanyIdCustomersSkCustomerIdPutExecute(r)
}

/*
V1CompanySkCompanyIdCustomersSkCustomerIdPut UPDATE Customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skCompanyId
 @param skCustomerId
 @return ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest
*/
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersSkCustomerIdPut(ctx context.Context, skCompanyId string, skCustomerId string) ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest {
	return ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest{
		ApiService: a,
		ctx: ctx,
		skCompanyId: skCompanyId,
		skCustomerId: skCustomerId,
	}
}

// Execute executes the request
func (a *TeamCustomersApiService) V1CompanySkCompanyIdCustomersSkCustomerIdPutExecute(r ApiV1CompanySkCompanyIdCustomersSkCustomerIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamCustomersApiService.V1CompanySkCompanyIdCustomersSkCustomerIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/company/{skCompanyId}/customers/{skCustomerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skCompanyId"+"}", url.PathEscape(parameterToString(r.skCompanyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skCustomerId"+"}", url.PathEscape(parameterToString(r.skCustomerId, "")), -1)

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
