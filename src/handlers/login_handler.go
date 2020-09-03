package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/camolezi/MicroservicesGolang/src/claims"
	"github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/dgrijalva/jwt-go"
)

//LoginHandler handles login
type LoginHandler struct {
	JWTKey        []byte
	RefreshJTWKey []byte
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

func (p *LoginHandler) authenticate(writer http.ResponseWriter, request *http.Request) {

	//For now assume that the user credentials are correct

	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	bodyData, _ := ioutil.ReadAll(request.Body)

	err := json.Unmarshal(bodyData, &user)
	log.Printf("%#v", user)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = services.CheckUserCredentials(user.Login, []byte(user.Password))

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		log.Println(err.Error())
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
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
	}

	//For now does not have refresh token implemented
	tokenStruct := struct {
		AccessToken  string `json:"acessToken"`
		RefreshToken string `json:"refreshToken"`
	}{AccessToken: signedToken, RefreshToken: "Placeholder"}

	tokenJSON, err := json.Marshal(tokenStruct)

	if err != nil {
		log.Println(err.Error())
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write([]byte(tokenJSON))

}
