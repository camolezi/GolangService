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

//WriteStatusCode manualy sets the status code
//Only call this after seeting all headers in the header map
func (r *Response) writeStatusCode(status int) {
	r.writer.WriteHeader(status)
}

func (r *Response) write(data []byte) {
	_, err := r.writer.Write(data)
	if err != nil {
		r.log.Error().Println(err)
	}
}

//WriteHeader write a header to a response(prefer not using this directly)- call this before WriteJSON
func (r *Response) WriteHeader(headerKey string, value string) {
	r.writer.Header().Set(headerKey, value)
}

//WriteJSON write json to the response and set header type to json- status code is 200ok
func (r *Response) WriteJSON(data []byte) {
	r.WriteHeader("Content-Type", "application/json")
	r.write(data)
}

//WriteJSONWithStatusCode write a json and a status code to the response
func (r *Response) WriteJSONWithStatusCode(data []byte, status int) {
	r.WriteHeader("Content-Type", "application/json")
	r.writeStatusCode(status)
	r.write(data)
}

//WriteError writes a generic error to the reponse
func (r *Response) WriteError(status int, err string) {
	errorStruct := struct {
		Status      string `json:"status"`
		Description string `json:"description"`
	}{Status: http.StatusText(status), Description: err}

	errorJSON, errMarshal := json.Marshal(errorStruct)

	if errMarshal != nil {
		r.log.Error().Println(errMarshal.Error())
		return
	}

	r.WriteJSONWithStatusCode(errorJSON, status)
}

//Standard http responses- Use these

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
	r.WriteJSONWithStatusCode(data, http.StatusCreated)
}

//CreateResponse returns a new response object
func CreateResponse(writer http.ResponseWriter, logger debug.Logger) *Response {
	return &Response{writer: writer, log: logger}
}
