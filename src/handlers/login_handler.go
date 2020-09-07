package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/claims"
	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/request"
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

func (p *LoginHandler) authenticate(defaultWriter http.ResponseWriter, defaultRequest *http.Request) {

	response := response.CreateResponse(defaultWriter, p.Log)
	request := request.CreateRequest(defaultRequest, p.Log)

	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(request.GetBody()).Decode(&user)
	if err != nil {
		response.BadRequest(err.Error())
		return
	}

	err = services.CheckUserCredentials(user.Login, []byte(user.Password))
	if err != nil {
		response.WriteError(http.StatusUnauthorized, "Credentials not accepted")
		return
	}

	// Audience  string `json:"aud,omitempty"`
	// ExpiresAt int64  `json:"exp,omitempty"`
	// Id        string `json:"jti,omitempty"`
	// IssuedAt  int64  `json:"iat,omitempty"`
	// Issuer    string `json:"iss,omitempty"`
	// NotBefore int64  `json:"nbf,omitempty"`
	// Subject   string `json:"sub,omitempty"`

	//Create token
	claims := claims.Claims{
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(p.JWTKey)

	if err != nil {
		response.ServerError(err.Error())
		return
	}

	//Return the token
	tokenStruct := struct {
		Login       string `json:"login"`
		AccessToken string `json:"acessToken"`
	}{AccessToken: signedToken, Login: user.Login}

	tokenJSON, err := json.Marshal(tokenStruct)

	if err != nil {
		response.ServerError(err.Error())
		return
	}

	response.WriteJSON(tokenJSON)

}
