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

type domains struct{}

func (*domains) GetPost(id uint64) (domainPkg.Post, error) {
	return domainPkg.GetPost(id)
}

func init() {
	domainVar = &domains{}
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

//GetLatestPosts return a array with the latest posts created
func GetLatestPosts(numberOfPosts uint) ([]domainPkg.Post, error) {
	return domainPkg.GetLatestPosts(numberOfPosts)
}

//NewPost trys to Create a new post
func NewPost(id uint64, post domainPkg.Post) error {
	return domainPkg.NewPost(id, post)
}
