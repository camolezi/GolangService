package handlers

//This file manages dependencies with services

import (
	"github.com/camolezi/MicroservicesGolang/src/domain"
	servicesPkg "github.com/camolezi/MicroservicesGolang/src/services"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//Interface for the services used by this handlers
type servicesInterface interface {
	GetPost(id uint64) (domain.Post, *utils.ErrorAPI)
}

//servicesWrapper is a wrapper for the actual services
type servicesWrapper struct{}

func (s *servicesWrapper) GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {
	return servicesPkg.GetPost(id)
}
