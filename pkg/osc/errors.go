package osc

import (
	"errors"
	"slices"
	"strconv"
)

func (e *ErrorResponse) Error() string {
	var msg string

	for i, v := range e.Errors {
		if i != 0 {
			msg += "\n"
		}
		msg += "[" + v.Code + "] " + v.Type
		if v.Details != "" {
			msg += ", " + v.Details
		}
	}

	return msg
}

func (e *ErrorResponse) GetCode() string {
	for _, v := range e.Errors {
		return v.Code
	}

	return ""
}

func AsErrorResponse(e error) *ErrorResponse {
	var err *ErrorResponse

	if errors.As(e, &err) {
		return err
	}

	return nil
}

func hasErrorCode(err error, codes []string) bool {
	if err := AsErrorResponse(err); err != nil {
		for _, e := range err.Errors {
			if slices.Contains(codes, e.Code) {
				return true
			}
		}
	}

	return false
}

func hasErrorCodeRange(err error, min, max int) bool {
	if err := AsErrorResponse(err); err != nil {
		for _, e := range err.Errors {
			c, err := strconv.Atoi(e.Code)
			if err == nil && c >= min && c <= max {
				return true
			}
		}
	}

	return false
}

func IsNotFound(err error) bool {
	return hasErrorCodeRange(err, 5000, 5999)
}

func IsConflict(err error) bool {
	return hasErrorCodeRange(err, 6000, 6999) || hasErrorCodeRange(err, 9000, 9999)
}

func IsAuthError(err error) bool {
	return hasErrorCode(err, []string{
		"1", "5", "7", "14", "20", "4120", "4000",
	})
}

func IsQuotaOrCapacity(err error) bool {
	return hasErrorCodeRange(err, 10000, 10999)
}
