package services

import (
	"net/http"

	domainPkg "github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type domainInterface interface {
	GetPost(id uint64) (domainPkg.Post, error)
}

var domainVar domainInterface

type domain struct{}

func (*domain) GetPost(id uint64) (domainPkg.Post, error) {
	return domainPkg.GetPost(id)
}

func init() {
	domainVar = &domain{}
}

//GetPost return a post of the specified id
func GetPost(id uint64) (domainPkg.Post, *utils.ErrorAPI) {

	post, resourceError := domainVar.GetPost(id)

	//Post not found
	if resourceError != nil {
		return post, &utils.ErrorAPI{ErrorCode: http.StatusNotFound, ErrorMessage: resourceError.Error()}
	}

	return post, nil
}

func addPost() {

}
