package domain

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//temporary
var userdbmock = make(map[string]User)

//CreateUser creates a new user
func CreateUser(newUser User) error {
	_, contain := userdbmock[newUser.Login]
	if contain {
		return errors.New("User login already exists")
	}

	userdbmock[newUser.Login] = newUser
	return nil
}

//GetUser gets a user from the database
func GetUser(login string) (User, error) {
	user, contain := userdbmock[login]
	if !contain {
		return User{}, &utils.ResourceError{ErrorMessage: "User not Found"}
	}
	return user, nil

}
