package request

import (
	"io"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
)

//Request is a class utility for helping handling with http requests
type Request struct {
	request *http.Request
	log     debug.Logger
}

//VerifyHeader is a utility for helping checking headers in requests
func (r *Request) VerifyHeader(key string, value string) bool {
	if header := r.request.Header.Get(key); header == value {
		return true
	}

	return false
}

//GetBody returns http request body
func (r *Request) GetBody() io.ReadCloser {
	return r.request.Body
}

//NewRequest creates a new Request utility
func CreateRequest(request *http.Request, log debug.Logger) *Request {
	return &Request{request: request, log: log}
}
