package nextgen

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type CiExecutionConfigApiService service

type CiExecutionConfigResponse struct {
	Status        string            `json:"status,omitempty"`
	Data          map[string]string `json:"data,omitempty"`
	MetaData      interface{}       `json:"metaData,omitempty"`
	CorrelationId string            `json:"correlationId,omitempty"`
}

// CiExecutionConfigUpdate represents a single image field update for the update-config endpoint
type CiExecutionConfigUpdate struct {
	Field string `json:"field"`
	Value string `json:"value,omitempty"`
}

const (
	ciExCfgGetDefaultConfigUrl  = "/ci/execution-config/get-default-config"
	ciExCfgGetCustomerConfigUrl = "/ci/execution-config/get-customer-config"
	ciExCfgUpdateConfigUrl      = "/ci/execution-config/update-config"
	ciExCfgResetConfigUrl       = "/ci/execution-config/reset-config"
)

func (a *CiExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (CiExecutionConfigResponse, *http.Response, error) {
	return a.executeCiRequest(ctx, http.MethodGet,
		ciExCfgGetDefaultConfigUrl, ciExCfgQueryParams(infraType, nil), []string{}, nil)
}

func (a *CiExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (CiExecutionConfigResponse, *http.Response, error) {
	return a.executeCiRequest(ctx, http.MethodGet,
		ciExCfgGetCustomerConfigUrl, ciExCfgQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *CiExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []CiExecutionConfigUpdate) (CiExecutionConfigResponse, *http.Response, error) {
	return a.executeCiRequest(ctx, http.MethodPost,
		ciExCfgUpdateConfigUrl, ciExCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *CiExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []CiExecutionConfigUpdate) (CiExecutionConfigResponse, *http.Response, error) {
	return a.executeCiRequest(ctx, http.MethodPost,
		ciExCfgResetConfigUrl, ciExCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *CiExecutionConfigApiService) executeCiRequest(ctx context.Context,
	httpMethod, path string, extraQueryParams url.Values,
	contentTypes []string, body interface{}) (CiExecutionConfigResponse, *http.Response, error) {

	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue CiExecutionConfigResponse
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

func ciExCfgQueryParams(infraType string, overridesOnly *bool) url.Values {
	params := url.Values{}
	params.Add("infra", parameterToString(infraType, ""))
	if overridesOnly != nil {
		params.Add("overridesOnly", parameterToString(*overridesOnly, ""))
	}
	return params
}
