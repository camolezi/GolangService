package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/camolezi/MicroservicesGolang/src/debug"
	"github.com/camolezi/MicroservicesGolang/src/model"
)

//AccessInterface defines methods to access data
type AccessInterface interface {
	GetPost(id int64) (model.Post, error)
}

//CreateDataAccess returns a new data access
func CreateDataAccess() Access {
	return Access{
		database: &databaseAccess,
		log:      databaseAccess.log,
	}
}

//Access is used for accessing the database from services
type Access struct {
	database *databaseConnection
	log      debug.Logger
}

//GetPost retrieves a post from the database
func (a *Access) GetPost(id int64) (model.Post, error) {

	post := model.Post{}

	//this is okay because id is a int64- But still need to berufy security
	query := fmt.Sprintf("SELECT * FROM posts WHERE id=%v", id)

	err := a.database.connection.QueryRow(context.Background(), query).
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
