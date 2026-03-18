package examples_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand/v2"
	"net/http"
	"testing"
	"time"
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

func lenOrZero[T any](items *[]T) int {
	if items == nil {
		return 0
	}

	return len(*items)
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
