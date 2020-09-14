package middleware

import "net/http"

//Middleware define the interface for middleware to implement
type Middleware interface {
	execute(next http.HandlerFunc) http.HandlerFunc
}

//NewChain creates a new chain of middlewares with a final handler
/*
	Last is the last Middleware in the call stack, the rest will be called in inverse order- (The last parameter is the first to be called)
	Final handler is the original handler - not a middleware. it will be called last- after all the Middleware
*/
func NewChain(finalHandler http.Handler, last Middleware, chain ...Middleware) http.Handler {
	middlewareChain := last.execute(finalHandler.ServeHTTP)

	for _, middle := range chain {
		middlewareChain = middle.execute(middlewareChain)
	}

	return http.HandlerFunc(middlewareChain)
}

//A final handler function that does nothing.
func doNothingHandler(writer http.ResponseWriter, request *http.Request) {}

//basicMiddlware is used as a intermediate middleware, before adding a Handler
type basicMiddlware struct {
	hadlerFunction http.HandlerFunc
}

func (b *basicMiddlware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		b.hadlerFunction(writer, request)
		//Call next function on the chain
		next(writer, request)
	}
}

//NewMiddlewareChain creates a new chain of middlewares without a final handler-
//This works in the same way as NewChain, but returns a Middleware instead of a http Handler
func NewMiddlewareChain(last Middleware, secondLast Middleware, chain ...Middleware) Middleware {
	middlewareChain := secondLast.execute(last.execute(doNothingHandler))

	for _, middle := range chain {
		middlewareChain = middle.execute(middlewareChain)
	}

	return &basicMiddlware{hadlerFunction: middlewareChain}
}
