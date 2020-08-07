package controllers

import (
	"testing"

	"github.com/camolezi/MicroservicesGolang/src/domain"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

type servicesTest struct {
	testGetPostFunction func(uint64) (domain.Post, *utils.ErrorAPI)
}

func (s *servicesTest) GetPost(id uint64) (domain.Post, *utils.ErrorAPI) {
	return s.testGetPostFunction(id)
}

func TestGetPost_Error(t *testing.T) {
	//Mock a function that returns a error
	serviceMock := &servicesTest{
		testGetPostFunction: func(uint64) (domain.Post, *utils.ErrorAPI) {
			return domain.Post{}, &utils.ErrorAPI{}
		},
	}

	service = serviceMock

}

func TestGetPost_Success(t *testing.T) {

}
