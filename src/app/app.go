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

	//This will be used in all routes
	basicChain := middleware.NewMiddlewareChain(
		&middleware.SecurityHeadersMiddleware{},
		&middleware.LogMiddleware{Log: logger.Debug()},
	)

	//Create the middleware chain for posts
	postPostHandler := middleware.NewChain(
		handlers.NewPostHandler(logger),
		&middleware.UserAuthMiddleware{JWTKey: config.JWTKey},
		basicChain,
	)

	getPostHandler := middleware.NewChain(
		handlers.NewPostHandler(logger),
		basicChain,
	)

	loginHandler := middleware.NewChain(
		&handlers.LoginHandler{JWTKey: config.JWTKey},
		basicChain,
	)

	serverMux := mux.CreateNewServeMux()
	serverMux.Get("/post/", getPostHandler)
	serverMux.Post("/post", postPostHandler)
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
