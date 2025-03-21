package apierror

type ServerError struct {
	StatusCode int
	Message    string
}

func NewServerError(message string) *ServerError {
	return &ServerError{400, message}
}

func (conflictError *ServerError) Error() string {
	return conflictError.Message
}
