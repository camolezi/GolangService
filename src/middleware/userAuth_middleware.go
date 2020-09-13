package middleware

import (
	"context"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/utils/claims"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserAuthMiddleware authenticate a user
type UserAuthMiddleware struct {
	JWTKey        []byte
	RefreshJTWKey []byte
	Log           debug.Logger
}

func (u *UserAuthMiddleware) execute(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//See user authorization here
		jtwToken := request.Header.Get("Authorization")

		token, err := jwt.ParseWithClaims(
			jtwToken,
			&claims.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return u.JWTKey, nil
			})

		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		ourclaims, ok := token.Claims.(*claims.Claims)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		loginCtx := context.WithValue(request.Context(), claims.KeyLogin, ourclaims.Login)

		//Call next function on the chain
		next(writer, request.WithContext(loginCtx))
	}
}
