package oks

import "errors"

func (e *HTTPValidationError) Error() string {
	var msg string

	for i, err := range *e.Errors {
		d, _ := err.Details.AsHTTPValidationErrorErrorsDetails0()
		for _, v := range d {
			if i != 0 {
				msg += "\n"
			}
			msg += "[" + v.Type + "] " + v.Msg
		}

		s, _ := err.Details.AsHTTPValidationErrorErrorsDetails1()
		if s != "" {
			msg += s
		}
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
			details, convErr := error.Details.AsHTTPValidationErrorErrorsDetails0()
			if convErr != nil {
				continue
			}

			for _, detail := range details {
				if detail.Type == code {
					return true
				}
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
