/*
	go test \
		-v \
		-run TestCiExecutionConfig \
		-account $HARNESS_ACCOUNT_ID \
		-api-key $HARNESS_PLATFORM_API_KEY \
		./harness/nextgen
*/

package nextgen

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type ciTestContext struct {
	t         *testing.T
	client    *APIClient
	ctx       context.Context
	infraType string
}

func TestCiExecutionConfig(t *testing.T) {
	client, ctx := getClientWithContext()

	tc := &ciTestContext{
		t:         t,
		client:    client,
		ctx:       ctx,
		infraType: "K8",
	}

	testUpdateImageType := "liteEngineTag"
	testUpdateImageTag := "harness/ci-lite-engine:updatedTmpNow"

	testCiGetDefaultConfig(tc)
	testCiGetCustomerConfigOverridesOnly(tc)

	testCiUpdateConfig(tc, testUpdateImageType, testUpdateImageTag)
	testCiGetCustomerConfigOverridesOnlyAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	testCiGetCustomerConfigAllAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)

	testCiResetConfig(tc, testUpdateImageType)
	testCiGetCustomerConfigOverridesOnlyAfterReset(tc, testUpdateImageType)
}

// pass: HTTP 200, status SUCCESS, data contains at least one image tag
// fail: any HTTP error, or data is nil (no images returned)
func testCiGetDefaultConfig(tc *ciTestContext) {
	tc.t.Log("=== GetDefaultConfig ===")
	resp, httpResp, err := tc.client.CiExecutionConfigApi.GetDefaultConfig(tc.ctx, tc.infraType)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
	require.NotNil(tc.t, resp.Data, "expected default config data, got nil")
	require.NotEmpty(tc.t, resp.Data.LiteEngineTag, "expected liteEngineTag to be set in default config")
}

// pass: HTTP 200, status SUCCESS, data nil or liteEngineTag empty (no overrides set yet)
// fail: HTTP error, or liteEngineTag already has an override (dirty state from a previous run)
func testCiGetCustomerConfigOverridesOnly(tc *ciTestContext) {
	tc.t.Log("=== GetCustomerConfig before update (overridesOnly=true, expect no overrides) ===")
	resp, httpResp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
	require.True(tc.t, resp.Data == nil || resp.Data.LiteEngineTag == "",
		"expected no liteEngineTag override before update, got %q", resp.Data.LiteEngineTag)
}

// pass: HTTP 200, status SUCCESS — the API accepted the override
// fail: any HTTP error (e.g. 400 bad field name, 401 auth, 403 forbidden)
func testCiUpdateConfig(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Logf("=== UpdateConfig (set %s = %s) ===", imageField, imageTag)
	body := []CiExecutionConfigUpdate{
		{Field: imageField, Value: imageTag},
	}
	resp, httpResp, err := tc.client.CiExecutionConfigApi.UpdateConfig(tc.ctx, tc.infraType, body)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
}

// pass: HTTP 200, data contains imageField set to imageTag (override is visible)
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testCiGetCustomerConfigOverridesOnlyAfterUpdate(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=true, expect override present) ===")
	resp, httpResp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
	checkCiField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, full config returned with imageField set to imageTag
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testCiGetCustomerConfigAllAfterUpdate(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=false, expect override present) ===")
	resp, httpResp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, false)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
	checkCiField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, status SUCCESS — the API accepted the reset
// fail: any HTTP error
func testCiResetConfig(tc *ciTestContext, imageField string) {
	tc.t.Logf("=== ResetConfig (reset %s) ===", imageField)
	body := []CiExecutionConfigUpdate{
		{Field: imageField},
	}
	resp, httpResp, err := tc.client.CiExecutionConfigApi.ResetConfig(tc.ctx, tc.infraType, body)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
}

// pass: HTTP 200, data nil or imageField empty — override was removed by reset
// fail: HTTP error, or imageField still holds the previously set override value
func testCiGetCustomerConfigOverridesOnlyAfterReset(tc *ciTestContext, imageField string) {
	tc.t.Log("=== GetCustomerConfig after reset (overridesOnly=true, expect no overrides) ===")
	resp, httpResp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, httpResp.StatusCode, err)
	require.True(tc.t, resp.Data == nil || resp.Data.LiteEngineTag == "",
		"%s still has override value %q after reset", imageField, resp.Data.LiteEngineTag)
	tc.t.Logf("PASS: %s override is gone after reset", imageField)
}

func checkCiField(t *testing.T, resp CiExecutionConfigResponse, imageField, imageTag string) {
	require.NotNil(t, resp.Data, "%s — response data is nil", imageField)
	require.Equal(t, imageTag, resp.Data.LiteEngineTag,
		"%s mismatch: got %q, expected %q", imageField, resp.Data.LiteEngineTag, imageTag)
	t.Logf("PASS: %s = %q", imageField, imageTag)
}

func printCiResult(t *testing.T, resp CiExecutionConfigResponse, statusCode int, err error) {
	require.NoError(t, err, "HTTP %d", statusCode)
	t.Logf("Status: %s (HTTP %d)", resp.Status, statusCode)
	if resp.Data != nil {
		b, _ := json.MarshalIndent(resp.Data, "", "  ")
		t.Log(string(b))
	} else {
		t.Log("Data: (empty — all defaults in use)")
	}
}
