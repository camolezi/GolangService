package handlers

//This file manages dependencies with services

import (
	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/services"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//Interface for the services used by this handlers
type servicesInterface interface {
	GetPost(id uint64) (model.Post, *utils.ErrorAPI)
	NewPost(id uint64, post model.Post) error
	GetLatestPosts(numberOfPosts uint) ([]model.Post, error)
}

//servicesWrapper is a wrapper for the actual services
type servicesWrapper struct{}

func (s *servicesWrapper) GetPost(id uint64) (model.Post, *utils.ErrorAPI) {
	return servicesPkg.GetPost(id)
}

func (s *servicesWrapper) NewPost(id uint64, post model.Post) error {
	return services.NewPost(id, post)
}

func (s *servicesWrapper) GetLatestPosts(numberOfPosts uint) ([]model.Post, error) {
	return services.GetLatestPosts(numberOfPosts)
}
