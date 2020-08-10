package app

import (
	"fmt"
	"net/http"
	"strconv"
)

type basicAnalytics struct {
	numberOfAccess uint64
	handleFunction func(http.ResponseWriter, *http.Request)
}

func (b *basicAnalytics) getNumberOfAcess() uint64 {
	return b.numberOfAccess
}

func (b *basicAnalytics) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b.handleFunction(writer, request)
}

//THis is just for debbuging should be left out produnction
func logHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello, new request recived" + request.RequestURI)
		function(writer, request)
	}
}

func analyticsHandler(function http.HandlerFunc, analyticsHandler *basicAnalytics) http.HandlerFunc {

	//basicAnalytics.numberOfAccess++

	return func(writer http.ResponseWriter, request *http.Request) {
		analyticsHandler.numberOfAccess++
		fmt.Println("We had:" + strconv.FormatUint(analyticsHandler.numberOfAccess, 10))
		function(writer, request)
	}
}

func withAnalytics(function http.HandlerFunc) *basicAnalytics {

	//Need to protect this from racing condition
	analyticsHandlerObject := basicAnalytics{numberOfAccess: 0}
	analyticsHandlerObject.handleFunction = analyticsHandler(function, &analyticsHandlerObject)

	return &analyticsHandlerObject
}
