package osc

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

func ErrorResponseHelper(e error) *ErrorResponse {
	httpVal, ok := e.(*ErrorResponse)
	if !ok {
		return nil
	}

	return httpVal
}
