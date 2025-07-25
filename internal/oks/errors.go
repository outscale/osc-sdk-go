package oks

func (e *HTTPValidationError) Error() string {
	var msg string

	for i, err := range *e.Errors {
		d, _ := err.Details.AsHTTPValidationErrorErrorsDetails0()
		if d != nil {
			for _, v := range d {
				if i != 0 {
					msg += "\n"
				}
				msg += "[" + v.Type + "] " + v.Msg
			}
		}

		s, _ := err.Details.AsHTTPValidationErrorErrorsDetails1()
		if s != "" {
			msg += s
		}
	}

	return msg
}
