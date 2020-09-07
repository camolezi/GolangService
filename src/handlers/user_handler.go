package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/handlers/request"
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

func (u *UserHandler) registerUser(defaultWriter http.ResponseWriter, defaultRequest *http.Request) {

	response := response.CreateResponse(defaultWriter, u.Log)
	request := request.CreateRequest(defaultRequest, u.Log)

	if !request.VerifyHeader("Content-Type", "application/json") {
		response.BadRequest("Content-Type must be JSON")
		return
	}

	user := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	requestBody := request.GetBody()
	err := json.NewDecoder(requestBody).Decode(&user)
	if err != nil {
		response.BadRequest(err.Error())
		return
	}

	err = services.CreateNewUser(model.User{Login: user.Login}, user.Password)
	if err != nil {
		u.Log.Error().Println(err.Error())
		response.BadRequest(err.Error())
		return
	}

	response.Created(nil)
}
