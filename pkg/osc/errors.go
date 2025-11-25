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

type responseInterface interface {
	Status() string
	StatusCode() int
	DeepError() error
}

func ErrorResponseHelper(e error) *ErrorResponse {
	httpVal, ok := e.(*ErrorResponse)
	if !ok {
		return nil
	}

	return httpVal
}

func StatusHelper(e error) *string {
	resp, ok := e.(responseInterface)
	if !ok {
		return nil
	}

	status := resp.Status()
	return &status
}

func StatusCodeHelper(e error) *int {
	resp, ok := e.(responseInterface)
	if !ok {
		return nil
	}

	statusCode := resp.StatusCode()
	return &statusCode
}

func DeepErrorHelper(e error) error {
	resp, ok := e.(responseInterface)
	if !ok {
		return nil
	}

	return resp.DeepError()
}
