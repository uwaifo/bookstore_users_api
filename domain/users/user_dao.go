package users

import (
	"fmt"

	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//used for interacting with the database

var (
	userDB = make(map[int64]*User)
)

//Get method .Here we use a pointer to ensure we are not working with a copy but raher
//the actual User object/struct from our user_dto.go
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email

	return nil

}

//Save method
func (user *User) Save() *errors.RestErr {
	currentUser := userDB[user.ID]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("A user with the email %s is already registered", user.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.ID))

	}
	userDB[user.ID] = user

	return nil

}
