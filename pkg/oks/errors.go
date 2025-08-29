package oks

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

type responseInterface interface {
	Status() string
	StatusCode() int
	DeepError() error
}

func HTTPValidationErrorHelper(e error) *HTTPValidationError {
	httpVal, ok := e.(*HTTPValidationError)
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
