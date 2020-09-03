package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/claims"
	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/response"
	"github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/dgrijalva/jwt-go"
)

//LoginHandler handles login
type LoginHandler struct {
	JWTKey        []byte
	RefreshJTWKey []byte
	Log           debug.Logger
}

func (p *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		p.authenticate(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (p *LoginHandler) authenticate(defaultWriter http.ResponseWriter, request *http.Request) {
	response := response.CreateResponse(defaultWriter, p.Log)

	//For now assume that the user credentials are correct
	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	bodyData, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(bodyData, &user)

	if err != nil {
		response.BadRequest(err.Error())
		return
	}

	_, err = services.CheckUserCredentials(user.Login, []byte(user.Password))

	if err != nil {
		response.WriteStatusCode(http.StatusUnauthorized)
		return
	}

	claims := claims.Claims{
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Audience  string `json:"aud,omitempty"`
	// ExpiresAt int64  `json:"exp,omitempty"`
	// Id        string `json:"jti,omitempty"`
	// IssuedAt  int64  `json:"iat,omitempty"`
	// Issuer    string `json:"iss,omitempty"`
	// NotBefore int64  `json:"nbf,omitempty"`
	// Subject   string `json:"sub,omitempty"`

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(p.JWTKey)

	if err != nil {
		response.ServerError(err.Error())
	}

	//For now does not have refresh token implemented
	tokenStruct := struct {
		AccessToken  string `json:"acessToken"`
		RefreshToken string `json:"refreshToken"`
	}{AccessToken: signedToken, RefreshToken: "Placeholder"}

	tokenJSON, err := json.Marshal(tokenStruct)

	if err != nil {
		response.ServerError(err.Error())
		return
	}

	response.WriteJSON(tokenJSON)

}
