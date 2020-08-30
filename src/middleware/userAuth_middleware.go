package middleware

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

//UserAuthMiddleware authenticate a user
type UserAuthMiddleware struct {
}

//Secrekey is obviously a placeholder
const Secrekey = "mysuperscretekey"

//Claims define jtw clains
type Claims struct {
	Login string `json:"username"`
	jwt.StandardClaims
}

func (u *UserAuthMiddleware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//See user authorization here
		jtwToken := request.Header.Get("Authorization")
		log.Println(jtwToken)

		token, err := jwt.ParseWithClaims(
			jtwToken,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(Secrekey), nil
			})

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println(claims.Login)

		//Call next function on the chain
		next(writer, request)
	}
}
