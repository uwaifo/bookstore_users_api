package users

import (
	"strings"

	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//User domain entity definition
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate method style
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower((user.Email)))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email addres")
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
