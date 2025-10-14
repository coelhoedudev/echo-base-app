package util

type HttpError struct {
	Message string
	Status  int
}

func NewHttpError(message string, status int) *HttpError {
	return &HttpError{
		Message: message,
		Status:  status,
	}
}
