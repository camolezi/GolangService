package services

import (
	"net/http"
	"testing"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type domainMock struct {
	GetPostMockFunction func(id uint64) (model.Post, error)
}

func (d *domainMock) GetPost(id uint64) (model.Post, error) {
	return d.GetPostMockFunction(id)
}

func TestGetPost_Error(t *testing.T) {
	mock := &domainMock{
		GetPostMockFunction: func(id uint64) (model.Post, error) {
			//Does not find the post with specified id
			return model.Post{}, &utils.ResourceError{ErrorMessage: "Error message"}
		},
	}

	domainVar = mock

	_, error := GetPost(0)

	if error == nil {
		t.Fatal("Error should not be nil in failed cases")
	}

	if error.ErrorMessage != "Error message" {
		t.Error("Message error modified")
	}

	if error.ErrorCode != http.StatusNotFound {
		t.Error("Status should be not found")
	}
}

func TestGetPost_Success(t *testing.T) {
	mock := &domainMock{
		GetPostMockFunction: func(id uint64) (model.Post, error) {
			//Does not find the post with specified id
			if id == 0 {
				return model.Post{ID: 0, Title: "This is not a real post"}, nil
			}
			return model.Post{}, &utils.ResourceError{ErrorMessage: "Error message"}
		},
	}
	domainVar = mock

	post, error := GetPost(0)

	if error != nil {
		t.Fatal("Error should be nil in success case")
	}

	if post.Title != "This is not a real post" {
		t.Error("Title of post is modified")
	}

}
