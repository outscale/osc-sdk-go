package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

type Logger interface {
	Request(ctx context.Context, req any)
	Response(ctx context.Context, resp any)
	RequestHttp(ctx context.Context, req *http.Request)
	ResponseHttp(ctx context.Context, resp *http.Response, d time.Duration)
	Error(ctx context.Context, err error)
}

var (
	logger *defaultLogger
	lock   = &sync.Mutex{}
)

type defaultLogger struct {
	*slog.Logger
}

func Default() *defaultLogger {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()

		logger = &defaultLogger{
			slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		}
	}

	return logger
}

func (l *defaultLogger) RequestHttp(ctx context.Context, req *http.Request) {
	var bodyString string

	if req.GetBody != nil {
		bodyReader, err := req.GetBody()
		if err == nil {
			bodyBytes, _ := io.ReadAll(bodyReader)
			bodyString = string(bodyBytes)
		}
	}

	l.DebugContext(
		ctx,
		"http request",
		"method",
		req.Method,
		"url",
		req.URL.String(),
		"body",
		bodyString,
	)
}

func (l *defaultLogger) ResponseHttp(ctx context.Context, resp *http.Response, d time.Duration) {
	// Copy response body
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		l.ErrorContext(
			ctx,
			"http response error",
			"status_code",
			resp.StatusCode,
			"duration",
			d.String(),
			"body",
			bodyString,
		)
	} else {
		l.DebugContext(ctx, "http response", "status_code", resp.StatusCode, "duration", d.String(), "body", bodyString)
	}
}

func (l *defaultLogger) Request(ctx context.Context, req any) {
	l.DebugContext(ctx, "request", "body", req)
}

func (l *defaultLogger) Response(ctx context.Context, resp any) {
	l.DebugContext(ctx, "response", "body", resp)
}

func (l *defaultLogger) Error(ctx context.Context, err error) {
	l.ErrorContext(ctx, "error", "message", err.Error())
}
