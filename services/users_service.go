package services

import (
	"github.com/uwaifo/bookstore_users_api/domain/users"
	"github.com/uwaifo/bookstore_users_api/utils"
	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//
var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

//CreateUser . . .
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//after the validatio we can now save the user
	user.DateCreated = utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	//return nil, nil
}

//GetUser . . . .
func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
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

//UpdateUser . . . .
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := UserService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	//noneed to validate since that is a constraint in the database
	/*
		if err := user.Validate(); err != nil {
			return nil, err
		}
	*/
	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}

	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email

	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

//DeleteUser / / /
func (s *userService) DeleteUser(userIDParam int64) *errors.RestErr {
	currentUser := &users.User{ID: userIDParam}
	return currentUser.Delete()

}

//Search  . . .
func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
