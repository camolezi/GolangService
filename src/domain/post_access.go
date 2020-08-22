package domain

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//Mock database for now
var dbMock = map[uint64]Post{
	0: {ID: 0, Title: "First post ever, you are a lucky guy for seeing this"},
}

//GetPost retrieves a post from the database
func GetPost(id uint64) (Post, error) {
	post, contain := dbMock[id]
	if !contain {
		return Post{}, &utils.ResourceError{ErrorMessage: "Post not Found"}
	}
	return post, nil
}

//NewPost creates a new post
func NewPost(id uint64, post Post) error {
	_, contain := dbMock[id]
	if contain {
		return errors.New("Post on this ID already exist")
	}

	dbMock[id] = post
	return nil
}
