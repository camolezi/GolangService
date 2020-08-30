package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/middleware"
	"github.com/dgrijalva/jwt-go"
)

//LoginHandler handles login
type LoginHandler struct {
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
	//For not assume that the user credentials are correct

	user := domain.User{}
	bodyData, _ := ioutil.ReadAll(request.Body)

	user.FromJSON(bodyData)
	log.Printf("%#v", user)

	claims := middleware.Claims{Login: user.Login}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(middleware.Secrekey))

	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write([]byte(signedToken))

}
