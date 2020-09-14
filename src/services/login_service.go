package services

import (
	"github.com/camolezi/MicroservicesGolang/src/services/data"
	"golang.org/x/crypto/bcrypt"
)

//CheckUserCredentials check to see if credentials of a user is correct
func CheckUserCredentials(username string, password []byte) error {
	//Maybe we want to do this asynchronous

	access := data.CreateDataAccess()

	//This is the password from the db
	user, err := access.GetUser(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, password)
	if err != nil {
		return err
	}

	//Correct credentials
	return nil

}
