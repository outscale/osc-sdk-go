package osc

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
