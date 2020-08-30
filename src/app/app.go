package app

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers"
	"github.com/camolezi/MicroservicesGolang/src/middleware"
)

//Config defines a configuration for starting a App
type Config struct {
	ServerAddr string
	LogLevel   debug.LogLevel
	LogRequest bool
}

//StartApp is the starting point of the application
func StartApp(config Config) {

	logger := debug.NewLogger(config.LogLevel)

	//Create the middleware chain for posts
	postHandler := middleware.NewChain(handlers.NewPostHandler(logger),
		&middleware.UserAuthMiddleware{},
		&middleware.SecurityHeadersMiddleware{},
		&middleware.LogMiddleware{Log: logger.Debug()})

	loginHandler := middleware.NewChain(&handlers.LoginHandler{},
		&middleware.SecurityHeadersMiddleware{},
		&middleware.LogMiddleware{Log: logger.Debug()})

	serverMux := http.NewServeMux()
	serverMux.Handle("/post/", postHandler)
	serverMux.Handle("/login", loginHandler)

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
