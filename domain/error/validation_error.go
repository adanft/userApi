package apierror

type ValidationError struct {
	StatusCode int
	Message    string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{400, message}
}

func (conflictError *ValidationError) Error() string {
	return conflictError.Message
}
