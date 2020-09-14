package middleware

import (
	"net/http"
)

//SecurityHeadersMiddleware add recomended security headers in the http response
type SecurityHeadersMiddleware struct {
}

func (s *SecurityHeadersMiddleware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Add the security headers

		//X-Frame-Options
		writer.Header().Set("X-Frame-Options", "DENY")

		//X-Content-Type-Options
		writer.Header().Set("X-Content-Type-Options", "nosniff")

		//Content-Security-Policy
		writer.Header().Set("Content-Security-Policy", "frame-ancestors 'none'")

		//Cache-Control
		writer.Header().Set("Cache-Control", "no-store")

		//Strict-Transport-Security-
		//Implemente when we have https

		//Call next function on the chain
		next(writer, request)
	}
}
