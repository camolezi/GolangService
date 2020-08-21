package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/handlers"
)

//StartApp is the starting point of the application
func StartApp() {
	const port = ":8081"

	serverMux := http.NewServeMux()
	fmt.Println("Started server on port " + port)
	//http.HandleFunc("/", logHandler(printHello))
	//ostHandler := withAnalytics(handlers.NewPostHandler())

	serverMux.Handle("/post/", handlers.NewPostHandler())
	//http.HandleFunc("/post", controllers.GetPost)

	httpServer := &http.Server{Addr: port, Handler: serverMux}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
