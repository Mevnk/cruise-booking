package customErrors

type UpdateError struct {
	message string
}

func NewUpdateError(message string) *UpdateError {
	return &UpdateError{
		message: message,
	}
}

func (e *UpdateError) Error() string {
	return e.message
}
