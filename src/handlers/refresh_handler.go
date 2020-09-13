package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/response"
	"github.com/camolezi/MicroservicesGolang/src/utils/claims"
	"github.com/dgrijalva/jwt-go"
)

//RefreshHandler is the handler to refresh access token
type RefreshHandler struct {
	Log    debug.Logger
	JWTKey []byte
}

func (u *RefreshHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	u.refreshToken(writer, request)
}

func (u *RefreshHandler) refreshToken(defaultWriter http.ResponseWriter, request *http.Request) {

	response := response.CreateResponse(defaultWriter, u.Log)
	//request := request.CreateRequest(defaultRequest, u.Log)

	userLogin, ok := request.Context().Value(claims.KeyLogin).(string)

	if !ok {
		response.ServerError("Expected KeyLogin in context")
	}

	//Create token
	_claims := claims.Claims{
		Login: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, _claims)
	signedToken, err := token.SignedString(u.JWTKey)

	if err != nil {
		response.ServerError(err.Error())
		return
	}

	//Return the token
	tokenStruct := struct {
		AccessToken string `json:"acessToken"`
		Login       string `json:"login"`
	}{AccessToken: signedToken, Login: userLogin}

	tokenJSON, err := json.Marshal(tokenStruct)
	if err != nil {
		response.ServerError(err.Error())
		return
	}

	response.WriteJSON(tokenJSON)

}
