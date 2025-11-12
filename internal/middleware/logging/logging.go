package logging

import (
	"net/http"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/logger"
)

type LoggingMiddleware struct {
	logger.Logger
}

func (l *LoggingMiddleware) Decorate(next http.RoundTripper) http.RoundTripper {
	return &innerLogger{inner: next, logger: l.Logger}
}

type innerLogger struct {
	inner  http.RoundTripper
	logger logger.Logger
}

func (l *innerLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	context := req.Context()

	l.logger.RequestHttp(context, req)

	start := time.Now()
	resp, err := l.inner.RoundTrip(req)
	duration := time.Since(start)
	if err != nil {
		l.logger.Error(context, err)
		return nil, err
	}

	l.logger.ResponseHttp(context, resp, duration)

	return resp, nil
}
