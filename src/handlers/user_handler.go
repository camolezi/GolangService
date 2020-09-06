package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/response"
	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/services"
)

//UserHandler is the handler for the user route
type UserHandler struct {
	Log debug.Logger
}

func (u *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	u.registerUser(writer, request)
}

func (u *UserHandler) registerUser(defaultWriter http.ResponseWriter, request *http.Request) {
	response := response.CreateResponse(defaultWriter, u.Log)

	u.Log.Debug().Println(request.Body)
	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.BadRequest(err.Error())
		return
	}

	err = services.CreateNewUser(model.User{Login: user.Login}, user.Password)

	if err != nil {
		response.ServerError(err.Error())
		return
	}
	response.Created(nil)
}
