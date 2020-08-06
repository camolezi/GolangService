package errors

//ErrorAPI represents a error in Api search
type ErrorAPI struct {
	ErrorCode int
	Err       error
}
