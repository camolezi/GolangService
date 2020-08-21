package middleware

import "net/http"

//Middleware define the interface for middleware to implement
type Middleware interface {
	execute(next http.HandlerFunc) http.HandlerFunc
}

//A final handler function that does nothing
func doNothingHandler(writer http.ResponseWriter, request *http.Request) {}

//NewChain creates a new chain of middlewares with a final handler
/*
	Last is the last Middleware in call stack, secondLast is the secondLast and the order is defined by the order of the parameters
	Final handler is the original handler - not a middleware. This will be called last- after all the Middleware
*/
func NewChain(finalHandler http.Handler, last Middleware, chain ...Middleware) http.Handler {
	middlewareChain := last.execute(finalHandler.ServeHTTP)

	for _, middle := range chain {
		middlewareChain = middle.execute(middlewareChain)
	}

	return http.HandlerFunc(middlewareChain)
}

//NewMiddlewareChain creates a new chain of middlewares without a final handler
func NewMiddlewareChain(last Middleware, secondLast Middleware, chain ...Middleware) http.Handler {
	middlewareChain := secondLast.execute(last.execute(doNothingHandler))

	for _, middle := range chain {
		middlewareChain = middle.execute(middlewareChain)
	}

	return http.HandlerFunc(middlewareChain)
}
