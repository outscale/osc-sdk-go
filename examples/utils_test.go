package examples_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/require"
)

func marshalRedacted(value any) ([]byte, error) {
	return json.Marshal(redactValue(reflect.ValueOf(value)))
}

func redactValue(value reflect.Value) any {
	if !value.IsValid() {
		return nil
	}
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		return redactValue(value.Elem())
	}

	switch value.Kind() {
	case reflect.Struct:
		result := make(map[string]any)
		typeOfValue := value.Type()
		for i := range value.NumField() {
			field := typeOfValue.Field(i)
			if field.PkgPath != "" {
				continue
			}
			jsonName, _, _ := strings.Cut(field.Tag.Get("json"), ",")
			if jsonName == "-" {
				continue
			}
			if jsonName == "" {
				jsonName = field.Name
			}
			if tag := field.Tag.Get("log"); tag == "sensitive" || tag == "pii" {
				result[jsonName] = "[REDACTED]"
			} else {
				result[jsonName] = redactValue(value.Field(i))
			}
		}
		return result
	case reflect.Slice, reflect.Array:
		result := make([]any, value.Len())
		for i := range result {
			result[i] = redactValue(value.Index(i))
		}
		return result
	case reflect.Map:
		result := make(map[string]any)
		for _, key := range value.MapKeys() {
			result[fmt.Sprint(key.Interface())] = redactValue(value.MapIndex(key))
		}
		return result
	default:
		return value.Interface()
	}
}

type testingLogger struct {
	*testing.T
}

func (t *testingLogger) RequestHttp(ctx context.Context, req *http.Request) {
	t.Log(
		"[http request]",
		"method: ",
		req.Method,
		", url: ",
		req.URL.String(),
	)
}

func (t *testingLogger) ResponseHttp(ctx context.Context, resp *http.Response, d time.Duration) {
	if resp.StatusCode != http.StatusOK {
		t.Log(
			"[http response error]",
			"status_code: ",
			resp.StatusCode,
			", duration: ",
			d.String(),
		)
	} else {
		t.Log("[http response]", "status_code: ", resp.StatusCode, ", duration: ", d.String())
	}
}

func (t *testingLogger) Request(ctx context.Context, req any) {
	jsonBody, err := marshalRedacted(req)
	if err != nil {
		t.Log("[request]", "error :", err.Error())
		return
	}
	t.Log("[request]", "body: ", string(jsonBody))
}

func (t *testingLogger) Response(ctx context.Context, resp any) {
	jsonBody, err := marshalRedacted(resp)
	if err != nil {
		t.Log("[response]", "error :", err.Error())
		return
	}
	t.Log("[response]", "body :", string(jsonBody))
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
