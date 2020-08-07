package controllers

import (
	"net/http"
	"strconv"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type servicesInterface interface {
	GetPost(id uint64) (domain.Post, *utils.ErrorAPI)
}

var (
	service servicesInterface
)

type services struct{}

func (s *services) GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {
	return servicesPkg.GetPost(id)
}

func init() {
	service = &services{}
}

//GetPost is a function to handle GET requests at /users
func GetPost(writer http.ResponseWriter, request *http.Request) {

	//Get id from url- implement
	idString := request.URL.Query().Get("id")

	if idString == "" {
		//Serve default page
		writer.Write([]byte("Default post will be here, or all posts"))
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)

	//Type error- id needs to be a number
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("id needs to be a number"))
		return
	}

	post, apiError := service.GetPost(id)
	//Api Error
	if apiError != nil {
		writer.WriteHeader(apiError.ErrorCode) //Write error header
		writer.Write([]byte(apiError.ErrorMessage))
		return
	}

	writer.Write([]byte(post.Title))
}
