package osc

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func (e *ErrorResponse) Error() string {
	var msg string

	for i, v := range *e.Errors {
		if i != 0 {
			msg += "\n"
		}
		msg += "[" + *v.Code + "] " + *v.Type
		if v.Details != nil {
			msg += ", " + *v.Details
		}
	}

	return msg
}

func ErrorHelper(e error) string {
	err, ok := e.(*ErrorResponse)
	if !ok {
		return "not a OAPI error"
	}

	for _, err := range *err.Errors {
		return *err.Code
	}

	return "no error"
}

func RetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	shouldRetry, err := retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	if shouldRetry {
		return shouldRetry, err
	}

	if resp.StatusCode == http.StatusConflict {
		return true, nil
	}

	return shouldRetry, nil
}
