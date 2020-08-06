package controllers

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/errors"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
)

type servicesInterface interface {
	GetPost(id uint64) (domain.Post, *errors.ErrorAPI)
}

var (
	service servicesInterface
)

type services struct{}

func (s *services) GetPost(id uint64) (domain.Post, *errors.ErrorAPI) {
	return servicesPkg.GetPost(id)
}

func init() {
	service = &services{}
}

//GetPost is a function to handle GET requests at /users
func GetPost(writer http.ResponseWriter, request *http.Request) {

	//Get id from url- implement
	id := 0
	post, apiError := service.GetPost(uint64(id))

	//Error
	if apiError != nil {
		//Respond here with 404 not found
		writer.WriteHeader(404) //Write error hero
	}

	writer.Write([]byte(post.Title))
}
