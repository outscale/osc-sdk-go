package logging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
}

type LoggingMiddleware struct {
	Logger Logger
}

func (l *LoggingMiddleware) Decorate(next http.RoundTripper) http.RoundTripper {
	return &innerLogger{inner: next, logger: l.Logger}
}

type innerLogger struct {
	inner  http.RoundTripper
	logger Logger
}

func (l *innerLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	context := req.Context()

	if getBody := req.GetBody; getBody != nil {
		reader, _ := getBody()
		l.logger.DebugContext(context, "request", req.Method, req.URL, reader)
	} else {
		l.logger.DebugContext(context, "request", req.Method, req.URL)
	}

	start := time.Now()
	resp, err := l.inner.RoundTrip(req)
	duration := time.Since(start)
	if err != nil {
		l.logger.ErrorContext(context, "error", err.Error())
	}

	respBodyBytes, _ := io.ReadAll(resp.Body)
	l.logger.InfoContext(
		context,
		"response",
		req.Method,
		req.URL,
		resp.StatusCode,
		duration,
		respBodyBytes,
	)
	resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))

	return resp, err
}
