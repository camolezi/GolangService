package data

import (
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
