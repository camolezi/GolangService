package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/services"
)

//UserHandler is the handler for the user route
type UserHandler struct {
}

func (u *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	u.registerUser(writer, request)
}

func (u *UserHandler) registerUser(writer http.ResponseWriter, request *http.Request) {

	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	bodyData, _ := ioutil.ReadAll(request.Body)

	err := json.Unmarshal(bodyData, &user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	services.CreateNewUser(model.User{Login: user.Login}, user.Password)
	writer.WriteHeader(http.StatusCreated)
}
