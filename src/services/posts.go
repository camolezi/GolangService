package services

import (
	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/errors"
)

//GetPost return a post of the id
func GetPost(id uint64) (domain.Post, *errors.ErrorAPI) {
	return domain.Post{
		ID:    0,
		Title: "My First Post",
	}, nil
}
