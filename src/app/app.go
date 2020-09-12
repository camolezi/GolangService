package app

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers"
	"github.com/camolezi/MicroservicesGolang/src/middleware"
	"github.com/camolezi/MicroservicesGolang/src/mux"
	"github.com/camolezi/MicroservicesGolang/src/services/data"
	// just for test
)

//Config defines a configuration for starting a App
type Config struct {
	ServerAddr    string
	LogLevel      debug.LogLevel
	LogRequest    bool
	JWTKey        []byte
	RefreshJTWKey []byte
	DBConfig      string
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
		&middleware.UserAuthMiddleware{
			JWTKey:        config.JWTKey,
			RefreshJTWKey: config.RefreshJTWKey},
		basicChain,
	)

	getPostHandler := middleware.NewChain(
		handlers.NewPostHandler(logger),
		basicChain,
	)

	loginHandler := middleware.NewChain(
		&handlers.LoginHandler{
			JWTKey:        config.JWTKey,
			RefreshJTWKey: config.RefreshJTWKey,
			Log:           logger},
		basicChain,
	)

	userHandler := middleware.NewChain(
		&handlers.UserHandler{Log: logger},
		basicChain,
	)

	refreshHandler := middleware.NewChain(
		&handlers.RefreshHandler{Log: logger, JWTKey: config.JWTKey},
		&middleware.UserAuthMiddleware{
			JWTKey:        config.JWTKey,
			RefreshJTWKey: config.RefreshJTWKey},
		basicChain,
	)

	serverMux := mux.CreateNewServeMux()
	serverMux.Get("/post/", getPostHandler)
	serverMux.Post("/post", postPostHandler)
	serverMux.Post("/login", loginHandler)
	serverMux.Post("/user", userHandler)
	serverMux.Post("/refresh", refreshHandler)

	//Connect to database
	data.InitializeDatabase(config.DBConfig, logger)
	defer data.CloseDatabase()

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
