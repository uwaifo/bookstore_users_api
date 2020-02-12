package users

import (
	"strings"

	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//
const (
	StatusActive = "active"
)

//User domain entity definition
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` // makes it an internal field
}

//Users (multiples of the User type) here is a slice of User
//this is helpfull for our search function in the contollers that respond with iterative objects
type Users []User

//Validate method style
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower((user.Email)))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email addres")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequest("INVALID PASSWORD")
	}

	return nil
}

/*
//function style

func Validate(user *User) *errors.RestErr {
	use.Email = string.TrimSpace(strings.Tolower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil

}
*/
