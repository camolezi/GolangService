package services

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/services/data"
	"golang.org/x/crypto/bcrypt"
)

//CheckUserCredentials check to see if credentials of a user is correct
func CheckUserCredentials(username string, password []byte) (bool, error) {
	//Maybe we want to do this asynchronous

	//This is the password from the db
	user, err := data.GetUser(username)

	if err != nil {
		return false, errors.New("User not found")
	}

	okay := bcrypt.CompareHashAndPassword(user.HashedPassword, password)

	if okay != nil {
		return false, errors.New("Incorrect credentials")
	}

	//Correct credentials
	return true, nil

}
