package domain

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//Mock database for now
var dbMock = map[uint64]Post{
	0: {ID: 0, Title: "First post ever, you are a lucky guy for seeing this", Body: PostBody{Text: "This is the body for first post"}},
	1: {ID: 1, Title: "Second post ever, you are a lucky guy for seeing this", Body: PostBody{Text: "This is the body for Second post"}},
	2: {ID: 2, Title: "third post ever, you are a lucky guy for seeing this", Body: PostBody{Text: "This is the body for third post"}},
	3: {ID: 3, Title: "forth post ever, you are a lucky guy for seeing this", Body: PostBody{Text: "This is the body for forth post"}},
	4: {ID: 4, Title: "fifth post ever, you are a lucky guy for seeing this", Body: PostBody{Text: "This is the body for fifth post"}},
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

//GetLatestPosts for now will return all posts in the map
func GetLatestPosts(size uint) ([]Post, error) {
	postSlice := make([]Post, 0)
	for _, post := range dbMock {
		postSlice = append(postSlice, post)
	}

	return postSlice, nil
}
