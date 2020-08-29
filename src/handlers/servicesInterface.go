package handlers

//This file manages dependencies with services

import (
	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/services"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//Interface for the services used by this handlers
type servicesInterface interface {
	GetPost(id uint64) (domain.Post, *utils.ErrorAPI)
	NewPost(id uint64, post domain.Post) error
	GetLatestPosts(numberOfPosts uint) ([]domain.Post, error)
}

//servicesWrapper is a wrapper for the actual services
type servicesWrapper struct{}

func (s *servicesWrapper) GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {
	return servicesPkg.GetPost(id)
}

func (s *servicesWrapper) NewPost(id uint64, post domain.Post) error {
	return services.NewPost(id, post)
}

func (s *servicesWrapper) GetLatestPosts(numberOfPosts uint) ([]domain.Post, error) {
	return services.GetLatestPosts(numberOfPosts)
}
