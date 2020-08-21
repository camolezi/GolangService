package middleware

import (
	"log"
	"net/http"
)

//LogMiddleware log all the requests in the provided log
type LogMiddleware struct {
	Log *log.Logger
}

func (l *LogMiddleware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Log the request
		l.Log.Println(request.Method)
		//Call next function on the chain
		next(writer, request)
	}
}
