package response

import (
	"encoding/json"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
)

//Response is used to facilitate writing http responses
type Response struct {
	writer http.ResponseWriter
	log    debug.Logger
}

//WriteHeader write a header to a response(prefer not using this directly)
func (r *Response) WriteHeader(headerKey string, value string) {
	r.writer.Header().Set(headerKey, value)
}

//WriteStatusCode manualy sets the status code
func (r *Response) WriteStatusCode(status int) {
	r.writer.WriteHeader(status)
}

func (r *Response) write(data []byte) {
	_, err := r.writer.Write(data)
	if err != nil {
		r.log.Error().Println(err)
	}
}

//WriteJSON write json to the response and set header type to json
func (r *Response) WriteJSON(data []byte) {
	r.WriteHeader("Content-Type", "application/json")
	r.write(data)
}

//WriteError writes a generic error to the reponse
func (r *Response) WriteError(status int, err string) {
	r.WriteStatusCode(status)

	errorStruct := struct {
		status      string
		description string
	}{status: http.StatusText(status), description: err}

	errorJSON, errMarshal := json.Marshal(errorStruct)

	if errMarshal != nil {
		r.log.Error().Println(errMarshal)
		return
	}

	r.WriteJSON(errorJSON)
}

//Standard http responses

//BadRequest Writes a status badrequest and a err message
func (r *Response) BadRequest(err string) {
	r.WriteError(http.StatusBadRequest, err)
}

//ServerError returns internalServerError status code and log the error
func (r *Response) ServerError(err string) {
	r.WriteError(http.StatusInternalServerError, "")
	r.log.Error().Println(err)
}

//Created returns http created status code and write the newly created object
func (r *Response) Created(data []byte) {
	r.WriteStatusCode(http.StatusCreated)
	r.WriteJSON(data)
}

//CreateResponse returns a new response object
func CreateResponse(writer http.ResponseWriter, logger debug.Logger) *Response {
	return &Response{writer: writer, log: logger}
}
