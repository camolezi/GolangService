package app

import (
	"flag"
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers"
)

//Config defines a configuration for starting a App
type Config struct {
}

//StartApp is the starting point of the application
func StartApp() {

	//Change this flag stuff for other place later
	addr := flag.String("addr", ":8080", "Define the port that the server will be listening and serving")
	logLevelString := flag.String("log", "debug",
		`Define the level of application logging: 
		error: only error logs.
		warning: errors and warning logs
		debug: errors, warnings and debug logs`,
	)

	flag.Parse()

	var logLevel debug.LogLevel
	switch *logLevelString {
	case "debug":
		logLevel = debug.DebugLevel
	case "warning":
		logLevel = debug.WarningLevel
	case "error":
		logLevel = debug.ErrorLevel
	default:
		log.Fatalln("Invalid Input for log flag, Options: debug,warning,error. Use -help for more info")
	}
	logger := debug.NewLogger(logLevel)

	serverMux := http.NewServeMux()
	serverMux.Handle("/post/", handlers.NewPostHandler(logger))

	httpServer := &http.Server{
		Addr:     *addr,
		Handler:  serverMux,
		ErrorLog: logger.Error(),
	}

	logger.Debug().Println("Starting server on port " + *addr)

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error().Fatal(err.Error())
	}
}
