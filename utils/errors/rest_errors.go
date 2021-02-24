package errors

type RestError struct {
	Message    string
	StatusCode int
	Error      string
}

func NewBadRequest(message string) *RestError {
	restError := RestError{
		Message:    message,
		StatusCode: 400,
		Error:      "bad requset",
	}
	return &restError
}

func NotFoundRequest(message string) *RestError {
	restError := RestError{
		Message:    message,
		StatusCode: 400,
		Error:      "not found",
	}
	return &restError
}
