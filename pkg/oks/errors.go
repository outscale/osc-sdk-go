package oks

import (
	"errors"
	"fmt"
)

func (e *ErrorResponse) Error() string {
	var errror string

	for _, e := range e.Errors {
		errror += e.Error() + "\n"
	}

	return errror
}

func (e *ErrorItem) Error() string {
	var detail string

	if v, err := e.Details.AsErrorItemDetails0(); err == nil {
		detail = v
	}

	if v, err := e.Details.AsErrorItemDetails1(); err == nil {
		for _, d := range v {
			detail += fmt.Sprintf("%s: %s, ", d.Type, d.Msg)
		}
	}

	return fmt.Sprintf("Error %s: %s", e.Code, detail)
}

func AsErrorResponse(e error) *ErrorResponse {
	var err *ErrorResponse

	if errors.As(e, &err) {
		return err
	}

	return nil
}

func isErrorType(err error, code string) bool {
	if err := AsErrorResponse(err); err != nil {
		for _, error := range err.Errors {
			if error.Code == code {
				return true
			}
		}
	}

	return false
}

func IsNotFound(err error) bool {
	return isErrorType(err, "404")
}

func IsConflict(err error) bool {
	return isErrorType(err, "409")
}
