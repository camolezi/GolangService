package services

import (
	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/services/data"
	"golang.org/x/crypto/bcrypt"
)

//CreateNewUser is the service that creates new users - for now only
func CreateNewUser(user model.User, password string) (bool, error) {
	//Maybe we want to do this asynchronous

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return false, err
	}

	user.HashedPassword = hash

	//This is the password from the db
	err = data.CreateUser(user)

	if err != nil {
		return false, err
	}

	//Created
	return true, nil

}
