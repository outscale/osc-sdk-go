package oks

import (
	"errors"
	"fmt"
)

func (e *ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Details)
}

func (e *HTTPValidationError) Error() string {
	var msg string

	for i, err := range *e.Errors {
		if i != 0 {
			msg += "\n"
		}
		msg += err.Error()
	}

	return msg
}

func HTTPValidationErrorHelper(e error) *HTTPValidationError {
	var err *HTTPValidationError

	if errors.As(e, &err) {
		return err
	}

	return nil
}

func isErrorType(err error, code string) bool {
	if err := HTTPValidationErrorHelper(err); err != nil {
		for _, error := range *err.Errors {
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
