package data

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//temporary
var userdbmock = make(map[string]model.User)

//CreateUser creates a new user
func CreateUser(newUser model.User) error {
	_, contain := userdbmock[newUser.Login]
	if contain {
		return errors.New("User login already exists")
	}

	userdbmock[newUser.Login] = newUser
	return nil
}

//GetUser gets a user from the database
func GetUser(login string) (model.User, error) {
	user, contain := userdbmock[login]
	if !contain {
		return model.User{}, &utils.ResourceError{ErrorMessage: "User not Found"}
	}
	return user, nil

}
