package utils

//ErrorAPI represents a error in Api search
type ErrorAPI struct {
	ErrorCode    int
	ErrorMessage string
}

func (e *ErrorAPI) Error() string {
	return e.ErrorMessage
}
