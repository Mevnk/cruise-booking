package customErrors

type CustomError struct {
	message string
}

func NewCustomError(entity Entity, errorCase Case) *CustomError {
	return &CustomError{
		message: string(entity) + " " + string(errorCase) + " FAIL",
	}
}

func (e *CustomError) Error() string {
	return e.message
}
