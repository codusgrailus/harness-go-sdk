package nextgen

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type IacmExecutionConfigApiService service

const (
	iacmExCfgGetDefaultConfigUrl  = "/iacm-manager/execution-config/get-default-config"
	iacmExCfgGetCustomerConfigUrl = "/iacm-manager/execution-config/get-customer-config"
	iacmExCfgUpdateConfigUrl      = "/iacm-manager/execution-config/update-config"
	iacmExCfgResetConfigUrl       = "/iacm-manager/execution-config/reset-config"
)

func (a *IacmExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (IacmExecutionConfigResponse, *http.Response, error) {
	return a.executeIacmRequest(ctx, http.MethodGet,
		iacmExCfgGetDefaultConfigUrl, iacmExCfgQueryParams(infraType, nil), []string{}, nil)
}

func (a *IacmExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (IacmExecutionConfigResponse, *http.Response, error) {
	return a.executeIacmRequest(ctx, http.MethodGet,
		iacmExCfgGetCustomerConfigUrl, iacmExCfgQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *IacmExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []IacmExecutionConfigUpdate) (IacmExecutionConfigResponse, *http.Response, error) {
	return a.executeIacmRequest(ctx, http.MethodPost,
		iacmExCfgUpdateConfigUrl, iacmExCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *IacmExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []IacmExecutionConfigUpdate) (IacmExecutionConfigResponse, *http.Response, error) {
	return a.executeIacmRequest(ctx, http.MethodPost,
		iacmExCfgResetConfigUrl, iacmExCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *IacmExecutionConfigApiService) executeIacmRequest(ctx context.Context,
	httpMethod, path string, extraQueryParams url.Values,
	contentTypes []string, body interface{}) (IacmExecutionConfigResponse, *http.Response, error) {

	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue IacmExecutionConfigResponse
	)

	localVarPath := a.client.cfg.BasePath + path

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

func iacmExCfgQueryParams(infraType string, overridesOnly *bool) url.Values {
	params := url.Values{}
	params.Add("infra", parameterToString(infraType, ""))
	if overridesOnly != nil {
		params.Add("overridesOnly", parameterToString(*overridesOnly, ""))
	}
	return params
}
