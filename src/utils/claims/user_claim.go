package claims

import "github.com/dgrijalva/jwt-go"

//Claims define jtw clains
type Claims struct {
	Login string `json:"username"`
	jwt.StandardClaims
}
