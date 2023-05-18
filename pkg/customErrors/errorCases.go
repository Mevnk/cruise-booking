package customErrors

type Case string
type ScanErrorCase string

const (
	Creation Case = "CREATION"
	Update        = "UPDATE"
	NotFound      = "SEARCH"
	Deletion      = "DELETION"
)

const (
	Scan      ScanErrorCase = "SCANNING"
	Parse                   = "PARSING"
	Iteration               = "ITERATION"
)
