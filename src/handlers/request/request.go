package request

import "net/http"

//Request is a class utility for helping handling with http requests
type Request struct {
	request *http.Request
}

//CheckHeader is a utility for helping checking headers in requests
func (r *Request) CheckHeader(key string, value string) {

}

//NewRequest creates a new Request utility
func NewRequest(request *http.Request) *Request {
	return &Request{request: request}
}
