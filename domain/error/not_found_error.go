package apierror

type NotFoundError struct {
	StatusCode int
	Message    string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{404, message}
}

func (conflictError *NotFoundError) Error() string {
	return conflictError.Message
}
