package utils

//ResourceError represents a error in locating resource
type ResourceError struct {
	ErrorMessage string
}

func (e *ResourceError) Error() string {
	return e.ErrorMessage
}
