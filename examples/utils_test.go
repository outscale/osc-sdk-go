package examples_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/require"
)

type testingLogger struct {
	*testing.T
}

func (t *testingLogger) RequestHttp(ctx context.Context, req *http.Request) {
	var bodyString string

	if req.GetBody != nil {
		bodyReader, err := req.GetBody()
		if err == nil {
			bodyBytes, _ := io.ReadAll(bodyReader)
			bodyString = string(bodyBytes)
		}
	}

	t.Log(
		"[http request]",
		"method: ",
		req.Method,
		", url: ",
		req.URL.String(),
		" body: ",
		bodyString,
	)
}

func (t *testingLogger) ResponseHttp(ctx context.Context, resp *http.Response, d time.Duration) {
	// Copy response body
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		t.Log(
			"[http response error]",
			"status_code: ",
			resp.StatusCode,
			", duration: ",
			d.String(),
			", body: ",
			bodyString,
		)
	} else {
		t.Log("[http response]", "status_code: ", resp.StatusCode, ", duration: ", d.String(), ", body: ", bodyString)
	}
}

func (t *testingLogger) Request(ctx context.Context, req any) {
	json, err := json.Marshal(req)
	if err != nil {
		t.Log("[request]", "error :", err.Error())
	}

	if len(json) == 0 {
		t.Log("[request]", "body: ", string(json))
	}
}

func (t *testingLogger) Response(ctx context.Context, resp any) {
	json, err := json.Marshal(resp)
	if err != nil {
		t.Log("[response]", "error :", err.Error())
	}

	if len(json) == 0 {
		t.Log("[response]", "body :", string(json))
	}
}

func (t *testingLogger) Error(ctx context.Context, err error) {
	t.Log("[error]", err.Error())
}

func RandomString(length int) string {
	if length <= 0 {
		return ""
	}

	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = alphabet[rand.IntN(len(alphabet))]
	}

	return string(bytes)
}

func newOSCClient(t *testing.T) *osc.Client {
	t.Helper()

	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := osc.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	return client
}

func newOKSClient(t *testing.T) *oks.Client {
	t.Helper()

	userProfile, err := profile.New()
	require.NoError(t, err)

	if !slices.Contains([]string{"eu-west-2", "cloudgouv-eu-west-1"}, userProfile.Region) {
		t.Skip("OKS is not deployed in this region")
	}

	client, err := oks.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	return client
}

func skipIfOKSTestsDisabled(t *testing.T) {
	t.Helper()

	value := os.Getenv("SKIP_OKS_TESTS")
	if value == "" {
		return
	}

	skip, err := strconv.ParseBool(value)
	if err == nil && skip {
		t.Skip("skipping OKS tests because SKIP_OKS_TESTS is enabled")
	}
}
