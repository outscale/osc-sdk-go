package oks

import (
	"errors"
	"fmt"
	"strings"
)

func (e *ErrorResponse) Error() string {
	var errrorBuilder strings.Builder

	for _, e := range e.Errors {
		errrorBuilder.WriteString(e.Error() + "\n")
	}

	return errrorBuilder.String()
}

func (e *ErrorItem) Error() string {
	var detailBuilder strings.Builder

	if v, err := e.Details.AsErrorItemDetails0(); err == nil {
		detailBuilder.WriteString(v)
	}

	if v, err := e.Details.AsErrorItemDetails1(); err == nil {
		for _, d := range v {
			detailBuilder.WriteString(d.Type + ": " + d.Msg + ", ")
		}
	}

	return fmt.Sprintf("Error %s: %s", e.Code, detailBuilder.String())
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
