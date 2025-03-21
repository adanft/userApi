package apierror

type BadRequestError struct {
	StatusCode int
	Message    string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{400, message}
}

func (conflictError *BadRequestError) Error() string {
	return conflictError.Message
}
