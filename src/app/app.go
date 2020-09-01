package app

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers"
	"github.com/camolezi/MicroservicesGolang/src/middleware"
	"github.com/camolezi/MicroservicesGolang/src/mux"
)

//Config defines a configuration for starting a App
type Config struct {
	ServerAddr string
	LogLevel   debug.LogLevel
	LogRequest bool
	JWTKey     []byte
}

//StartApp is the starting point of the application
func StartApp(config Config) {

	logger := debug.NewLogger(config.LogLevel)

	//Create the middleware chain for posts
	postHandler := middleware.NewChain(handlers.NewPostHandler(logger),
		&middleware.UserAuthMiddleware{JWTKey: config.JWTKey},
		&middleware.SecurityHeadersMiddleware{},
		&middleware.LogMiddleware{Log: logger.Debug()})

	loginHandler := middleware.NewChain(&handlers.LoginHandler{JWTKey: config.JWTKey},
		&middleware.SecurityHeadersMiddleware{},
		&middleware.LogMiddleware{Log: logger.Debug()})

	serverMux := mux.CreateNewServeMux()
	serverMux.Get("/post/", postHandler)
	serverMux.Post("/login", loginHandler)

	httpServer := &http.Server{
		Addr:     config.ServerAddr,
		Handler:  serverMux,
		ErrorLog: logger.Error(),
	}

	logger.Debug().Println("Starting server on port " + config.ServerAddr)

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error().Fatal(err.Error())
	}
}
