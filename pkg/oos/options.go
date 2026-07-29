package oos

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
)

var _ middleware.InitializeMiddleware = (*retryTimeoutMiddleware)(nil)

type retryTimeoutMiddleware struct {
	timeout time.Duration
}

func (*retryTimeoutMiddleware) ID() string {
	return "OOSRetryTimeout"
}

func (m *retryTimeoutMiddleware) HandleInitialize(
	ctx context.Context,
	in middleware.InitializeInput,
	next middleware.InitializeHandler,
) (middleware.InitializeOutput, middleware.Metadata, error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	return next.HandleInitialize(ctx, in)
}

func WithRetryTimeout(timeout time.Duration) func(*s3.Options) {
	return func(options *s3.Options) {
		options.APIOptions = append(
			options.APIOptions,
			func(stack *middleware.Stack) error {
				return stack.Initialize.Add(
					&retryTimeoutMiddleware{timeout: timeout},
					middleware.Before,
				)
			},
		)
	}
}
