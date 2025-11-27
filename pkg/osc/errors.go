package osc

import "errors"

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
