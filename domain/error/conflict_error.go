package apierror

type ConflictError struct {
	StatusCode int
	Message    string
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{400, message}
}

func (conflictError *ConflictError) Error() string {
	return conflictError.Message
}
