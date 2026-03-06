package idp

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ExecutionConfigApiService service

// ExecutionConfigResponse is the API response wrapper for the IDP execution config endpoint.
// Data is a dynamic map of image field names to image tags, e.g. {"registerCatalog": "harness/registercatalog:1.9.0"}.
// Using a map avoids struct changes when the API adds or removes image types.
type ExecutionConfigResponse struct {
	Status        string            `json:"status,omitempty"`
	Data          map[string]string `json:"data,omitempty"`
	MetaData      interface{}       `json:"metaData,omitempty"`
	CorrelationId string            `json:"correlationId,omitempty"`
}

// ExecutionConfigUpdate represents a single image field update for the update-config endpoint
type ExecutionConfigUpdate struct {
	Field string `json:"field"`
	Value string `json:"value,omitempty"`
}

const (
	execCfgGetDefaultConfigUrl  = "/idp/execution-config/get-default-config"
	execCfgGetCustomerConfigUrl = "/idp/execution-config/get-customer-config"
	execCfgUpdateConfigUrl      = "/idp/execution-config/update-config"
	execCfgResetConfigUrl       = "/idp/execution-config/reset-config"
)

func (a *ExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (ExecutionConfigResponse, *http.Response, error) {
	return a.executeRequest(ctx, http.MethodGet,
		execCfgGetDefaultConfigUrl, execCfgQueryParams(infraType, nil), []string{}, nil)
}

func (a *ExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (ExecutionConfigResponse, *http.Response, error) {
	return a.executeRequest(ctx, http.MethodGet,
		execCfgGetCustomerConfigUrl, execCfgQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *ExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, *http.Response, error) {
	return a.executeRequest(ctx, http.MethodPost,
		execCfgUpdateConfigUrl, execCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *ExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, *http.Response, error) {
	return a.executeRequest(ctx, http.MethodPost,
		execCfgResetConfigUrl, execCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

// gatewayBase returns the gateway base URL without the /v1 suffix that the IDP
// client appends by default. The execution-config endpoints live at /idp/...,
// not under /v1.
func (a *ExecutionConfigApiService) gatewayBase() string {
	return strings.TrimSuffix(a.client.cfg.BasePath, "/v1")
}

func (a *ExecutionConfigApiService) executeRequest(ctx context.Context,
	httpMethod, path string, extraQueryParams url.Values,
	contentTypes []string, body interface{}) (ExecutionConfigResponse, *http.Response, error) {

	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ExecutionConfigResponse
	)

	localVarPath := a.gatewayBase() + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(a.client.AccountId, ""))
	for k, vs := range extraQueryParams {
		for _, v := range vs {
			localVarQueryParams.Add(k, v)
		}
	}

	localVarHttpContentType := selectHeaderContentType(contentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	localVarHttpHeaderAccept := selectHeaderAccept([]string{"application/json"})
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	if ctx != nil {
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}

	if body != nil {
		localVarPostBody = body
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, httpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer func() {
		_ = localVarHttpResponse.Body.Close()
	}()

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func execCfgQueryParams(infraType string, overridesOnly *bool) url.Values {
	params := url.Values{}
	params.Add("infra", parameterToString(infraType, ""))
	if overridesOnly != nil {
		params.Add("overridesOnly", parameterToString(*overridesOnly, ""))
	}
	return params
}
