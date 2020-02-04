package services

import (
	"github.com/uwaifo/bookstore_users_api/domain/users"
	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//CreateUser . . .
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//after the validatio we can now save the user
	//
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	//return nil, nil
}

//GetUser . . . .
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	// at this point we are certain that the argumat is an int 64 so i commnt the block
	/*
		if userID <= 0 {
			return nil, errors.NewBadRequest("invalid user id")
		}
	*/
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	//else we vall the dao gaints the db
	return result, nil

}