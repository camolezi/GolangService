package mux

import (
	"log"
	"net/http"
)

//If this get too complex or inefficient, we should replace it for a 3rd party router- like gorilla mux. In that case, it should be a drag and drop replacement.

//ServeMux in custom type of mux
type ServeMux struct {
	httpMux map[string]*http.ServeMux
	log     *log.Logger
}

func (s *ServeMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if mux, supported := s.httpMux[request.Method]; supported {
		//Redirect
		mux.ServeHTTP(writer, request)
		return
	}

	//Default
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

//Get will serve get requests
func (s *ServeMux) Get(path string, handler http.Handler) {
	if mux, supported := s.httpMux[http.MethodGet]; supported {
		mux.Handle(path, handler)
	}
}

//Post will serve post requests
func (s *ServeMux) Post(path string, handler http.Handler) {
	if mux, supported := s.httpMux[http.MethodPost]; supported {
		mux.Handle(path, handler)
	}
}

func (s *ServeMux) init() {
	//Should add all used http verbs here
	s.httpMux[http.MethodGet] = http.NewServeMux()
	s.httpMux[http.MethodPost] = http.NewServeMux()
}

//CreateNewServeMux return a new serve mux
func CreateNewServeMux() *ServeMux {
	newMux := &ServeMux{httpMux: make(map[string]*http.ServeMux)}
	newMux.init()
	return newMux
}
