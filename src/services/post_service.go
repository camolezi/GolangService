package services

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/services/data"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type domainInterface interface {
	GetPost(id int64) (model.Post, error)
}

var domainVar domainInterface

type domains struct{}

func (*domains) GetPost(id int64) (model.Post, error) {
	return data.GetPost(id)
}

func init() {
	domainVar = &domains{}
}

//GetPost return a post of the specified id
func GetPost(id int64) (model.Post, *utils.ErrorAPI) {

	//For test
	access := data.CreateDataAccess()
	post, resourceError := access.GetPost(id)

	//post, resourceError := domainVar.GetPost(id)

	//Post not found
	if resourceError != nil {
		return post, &utils.ErrorAPI{ErrorCode: http.StatusNotFound, ErrorMessage: resourceError.Error()}
	}

	return post, nil
}

//GetLatestPosts return a array with the latest posts created
func GetLatestPosts(numberOfPosts uint) ([]model.Post, error) {
	return data.GetLatestPosts(numberOfPosts)
}

//NewPost trys to Create a new post
func NewPost(id int64, post model.Post) error {
	return data.NewPost(id, post)
}
