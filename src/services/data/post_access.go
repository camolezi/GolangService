package data

import (
	"context"
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//GetPost retrieves a post from the database
func (a *Access) GetPost(id int64) (model.Post, error) {
	post := model.Post{}

	//this is okay because id is a int64- But still need to verify security
	query := "SELECT * FROM posts WHERE id=$1"

	err := a.database.connection.QueryRow(context.Background(), query, id).
		Scan(
			&post.ID,
			&post.Title,
			&post.CreatedAt,
			&post.UserLogin,
			&post.Body,
		)

	if err != nil {
		a.log.Warning().Printf("QueryRow failed: %v\n", err)
		return model.Post{}, errors.New("post not found")
	}

	return post, nil
}

//CreatePost creates a new post
func (a *Access) CreatePost(post model.Post) error {
	query := "INSERT INTO posts (createdAt,title,userLogin,body) VALUES(NOW(),$1,$2,$3)"

	tag, err := a.database.connection.Exec(context.Background(), query,
		post.Title,
		post.UserLogin,
		post.Body,
	)

	a.log.Debug().Println(tag)

	if err != nil {
		a.log.Warning().Printf("Inserted failed: %v\n", err)
		return err
	}

	return nil
}

//GetLatestPosts for now will return posts ordered by data(probably there is a more efficient way to do this)
func (a *Access) GetLatestPosts(size uint) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	//this is okay because id is a int64- But still need to verify security
	query := "SELECT * FROM posts ORDER BY createdAt DESC LIMIT $1"

	rows, err := a.database.connection.Query(context.Background(), query, size)
	defer rows.Close()

	if err != nil {
		a.log.Error().Printf("Latest posts query failed %v\n", err)
		return nil, err
	}

	for rows.Next() {

		post := model.Post{}
		err := rows.
			Scan(
				&post.ID,
				&post.Title,
				&post.CreatedAt,
				&post.UserLogin,
				&post.Body,
			)

		if err != nil {
			a.log.Warning().Printf("QueryRow failed: %v\n", err)
			return posts, errors.New("Error querying some of the posts")
		}

		posts = append(posts, post)
	}

	//Need to compare tag to see if the number of posts is correct
	return posts, nil

}

//PostAccess - Maybe create separated interfaces latter
type PostAccess interface {
}

//Mock

//Mock database for now
var dbMock = map[int64]model.Post{
	0: {ID: 0, Title: "First post ever, you are a lucky guy for seeing this"},
	1: {ID: 1, Title: "Second post ever, you are a lucky guy for seeing this"},
	2: {ID: 2, Title: "third post ever, you are a lucky guy for seeing this"},
	3: {ID: 3, Title: "forth post ever, you are a lucky guy for seeing this"},
	4: {ID: 4, Title: "fifth post ever, you are a lucky guy for seeing this"},
}

//GetPost retrieves a post from the database
func GetPost(id int64) (model.Post, error) {
	post, contain := dbMock[id]
	if !contain {
		return model.Post{}, &utils.ResourceError{ErrorMessage: "Post not Found"}
	}
	return post, nil
}

//NewPost creates a new post
func NewPost(id int64, post model.Post) error {
	_, contain := dbMock[id]
	if contain {
		return errors.New("Post on this ID already exist")
	}

	dbMock[id] = post
	return nil
}

//GetLatestPosts for now will return all posts in the map
func GetLatestPosts(size uint) ([]model.Post, error) {
	postSlice := make([]model.Post, 0)
	for _, post := range dbMock {
		postSlice = append(postSlice, post)
	}

	return postSlice, nil
}
