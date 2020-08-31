package middleware

import (
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/claims"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserAuthMiddleware authenticate a user
type UserAuthMiddleware struct {
	JWTKey []byte
}

func (u *UserAuthMiddleware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//See user authorization here
		jtwToken := request.Header.Get("Authorization")
		log.Println(jtwToken)

		token, err := jwt.ParseWithClaims(
			jtwToken,
			&claims.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return u.JWTKey, nil
			})

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*claims.Claims)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println(claims.Login)

		//Call next function on the chain
		next(writer, request)
	}
}
