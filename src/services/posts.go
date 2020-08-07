package services

import (
	"net/http"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//GetPost return a post of the id
func GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {

	post, resourceError := domain.GetPost(id)

	//Post not found
	if resourceError != nil {
		return post, &utils.ErrorAPI{ErrorCode: http.StatusNotFound, ErrorMessage: resourceError.Error()}
	}

	return post, nil
}
