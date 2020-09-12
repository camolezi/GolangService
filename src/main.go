package main

import (
	"flag"
	"log"

	"github.com/camolezi/MicroservicesGolang/src/app"
	"github.com/camolezi/MicroservicesGolang/src/debug"
)

//Main parses the command line flags and configure the application accordingly
func main() {

	//Change this flag stuff for other place later
	addr := flag.String("addr", ":8080", "Define the port that the server will be listening and serving")
	logLevelString := flag.String("log", "debug",
		`Define the level of application logging: 
		error: only error logs.
		warning: errors and warning logs
		debug: errors, warnings and debug logs`,
	)

	dbconfig := flag.String("dbconfig",
		"user=postgres password=userpassword host=localhost port=5432 dbname=db_name pool_max_conns=10",
		"String defining the configuration for the database connection",
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

	//this is obviously a placeholder
	Secrekey := []byte("mysuperscretekey")

	appConfig := app.Config{
		ServerAddr: *addr,
		LogLevel:   logLevel,
		JWTKey:     Secrekey,
		DBConfig:   *dbconfig,
	}

	app.StartApp(appConfig)
}
