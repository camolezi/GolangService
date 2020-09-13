package claims

import (
	"github.com/dgrijalva/jwt-go"
)

//Claims define jtw clains
type Claims struct {
	Login string `json:"username"`
	jwt.StandardClaims
}

//ContextKey Are used for passing values between middleware and handlers in context
type contextKey int

const (
	//KeyLogin is for passing login information
	KeyLogin contextKey = iota
)
