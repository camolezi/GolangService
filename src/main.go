package main

import (
	"fmt"
	"net/http"
)

func printHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func log(function http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello, new request recived")
		function(writer, request)
	}
}

func main() {
	fmt.Println("Started server")

	http.HandleFunc("/", log(printHello))
	http.ListenAndServe(":8080", nil)

}
