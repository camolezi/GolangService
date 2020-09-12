package data

import (
	"github.com/camolezi/MicroservicesGolang/src/model"
)

//AccessInterface defines methods to access data
type AccessInterface interface {
	GetPost(id int64) (model.Post, error)
}

//Access is used for accessing the database from services
type Access struct {
	database *databaseConnection
}

//GetPost retrieves a post from the database
func (a *Access) GetPost(id int64) {

}
