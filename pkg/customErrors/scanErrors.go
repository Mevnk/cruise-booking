package customErrors

type ScanError struct {
	message string
}

func NewScanError(entity Entity, errorCase ScanErrorCase) *ScanError {
	return &ScanError{
		message: string(entity) + " " + string(errorCase) + " FAIL",
	}
}

func (e *ScanError) Error() string {
	return e.message
}
