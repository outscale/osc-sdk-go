// Package middleware provides a way to chain together
// middleware functions for an http.RoundTripper.
package middleware

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/logger"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware/logging"
)

type Middleware interface {
	Decorate(next http.RoundTripper) http.RoundTripper
}

type MiddlewareSlot = int

const (
	MiddlewareSlotAuth MiddlewareSlot = iota
	MiddlewareSlotRetry
	MiddlewareSlotRateLimit
	MiddlewareSlotUseragent
	middlewareChainSize
)

func defaultPooledTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
		MaxIdleConns:          10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
		ForceAttemptHTTP2:     true,
		MaxIdleConnsPerHost:   10,
	}
	return transport
}

type MiddlewareChain struct {
	slots    []Middleware
	Base     *http.Transport
	composed http.RoundTripper
	logger   logger.Logger
}

type MiddlewareChainOption func(*MiddlewareChain) error

func NewMiddlewareChain(opts ...MiddlewareChainOption) (*MiddlewareChain, error) {
	mc := MiddlewareChain{
		slots: make([]Middleware, middlewareChainSize),
		Base:  defaultPooledTransport(),
	}

	err := mc.applyOptions(opts...)
	if err != nil {
		return nil, err
	}

	mc.rebuild()
	return &mc, nil
}

func WithMiddleware(slot MiddlewareSlot, mw Middleware) MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		mc.slots[slot] = mw
		return nil
	}
}

func WithBaseTransport(base *http.Transport) MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		mc.Base = base
		return nil
	}
}

func WithLogger(log logger.Logger) MiddlewareChainOption {
	return func(mc *MiddlewareChain) error {
		mc.logger = log
		return nil
	}
}

func (mc *MiddlewareChain) WithOptions(opts ...MiddlewareChainOption) (*MiddlewareChain, error) {
	if len(opts) == 0 {
		return mc, nil
	}

	mcp := MiddlewareChain{
		slots:  make([]Middleware, middlewareChainSize),
		Base:   mc.Base,
		logger: mc.logger,
	}

	_ = copy(mcp.slots, mc.slots)
	err := mcp.applyOptions(opts...)
	if err != nil {
		return nil, err
	}

	mcp.rebuild()
	return &mcp, nil
}

func (mc *MiddlewareChain) applyOptions(opts ...MiddlewareChainOption) error {
	for _, o := range opts {
		if err := o(mc); err != nil {
			return err
		}
	}

	return nil
}

func (mc *MiddlewareChain) rebuild() {
	var composed http.RoundTripper
	composed = mc.Base

	if mc.logger != nil {
		tmp := logging.LoggingMiddleware{Logger: mc.logger}
		composed = tmp.Decorate(composed)
	}

	for i := middlewareChainSize - 1; i >= 0; i-- {
		if mc.slots[i] != nil {
			composed = mc.slots[i].Decorate(composed)
		}
	}

	mc.composed = composed
}

func (mc *MiddlewareChain) RoundTrip(req *http.Request) (*http.Response, error) {
	return mc.composed.RoundTrip(req)
}

// Proxy for pkg logger
func (mc *MiddlewareChain) LogRequest(ctx context.Context, req any) {
	if mc.logger == nil {
		return
	}

	mc.logger.Request(ctx, req)
}

func (mc *MiddlewareChain) LogResponse(ctx context.Context, resp any) {
	if mc.logger == nil {
		return
	}

	mc.logger.Response(ctx, resp)
}
