package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/controllers"
)

//StartApp is the starting point of the application for now
func StartApp() {
	const port = ":8081"
	fmt.Println("Started server on port " + port)
	//http.HandleFunc("/", logHandler(printHello))

	postHandler := withAnalytics(controllers.GetPost)
	http.Handle("/post/", postHandler)
	http.HandleFunc("/post", controllers.GetPost)

	log.Fatal(http.ListenAndServe(port, nil))
}
