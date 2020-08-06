package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/controllers"
)

func printHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func logHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello, new request recived" + request.RequestURI)
		function(writer, request)
	}
}

//StartApp is the starting point of the application for now
func StartApp() {
	fmt.Println("Started server")
	http.HandleFunc("/", logHandler(printHello))
	http.HandleFunc("/post/", controllers.GetPost)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
