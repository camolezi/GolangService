package data

import (
	"context"
	"errors"
	"log"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//GetUser gets a user from the database
func (a *Access) GetUser(login string) (model.User, error) {
	user := model.User{}

	//still need to verify security
	query := "SELECT login,createdAt,email,userPass FROM account WHERE login=$1"

	err := a.database.connection.QueryRow(context.Background(), query, login).
		Scan(
			&user.Login,
			&user.CreatedAt,
			&user.Email,
			&user.HashedPassword,
		)

	if err != nil {
		a.log.Warning().Printf("QueryRow failed: %v\n", err)
		return user, err
	}

	return user, nil
}

//CreateUser tries to create a new user
func (a *Access) CreateUser(newUser model.User) error {
	query := "INSERT INTO account (createdAt,login,email,userPass) VALUES(NOW(),$1,$2,$3)"

	tag, err := a.database.connection.Exec(context.Background(), query,
		newUser.Login,
		newUser.Email,
		newUser.HashedPassword)

	a.log.Debug().Printf("%#v", tag.String())

	if err != nil {
		a.log.Warning().Printf("Inserted failed: %v\n", err)
		return err
	}

	return nil
}

// -- from here on, its just temporary mock---------
var userdbmock = make(map[string]model.User)

//CreateUser creates a new user
func CreateUser(newUser model.User) error {
	_, contain := userdbmock[newUser.Login]
	if contain {
		return errors.New("User login already exists")
	}

	userdbmock[newUser.Login] = newUser
	PrintUsers()
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

//PrintUsers is just debug for now
func PrintUsers() {
	log.Printf("%#v", userdbmock)
}
