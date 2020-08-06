package controllers

import (
	"fmt"
	"net/http"
)

//GetPost is a function to handle GET requests at /users
func GetPost(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HelloPost")
}
