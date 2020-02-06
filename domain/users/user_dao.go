package users

import (
	"fmt"
	//"errors"
	"strings"

	"github.com/uwaifo/bookstore_users_api/datasource/mysql/usersdb"
	"github.com/uwaifo/bookstore_users_api/utils"
	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//used for interacting with the database

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUdpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	errorNoRows      = "no rows in result set"
)

/*var (
	userDB = make(map[int64]*User)
)
*/

//Get method .Here we use a pointer to ensure we are not working with a copy but raher
//the actual User object/struct from our user_dto.go
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		//fmt.Println(err)
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))

		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error attempting to get user %d: %s", user.ID, err.Error()))
	}

	return nil

}

//Update . . .
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUdpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.ParseError(err)

	}
	return nil

}

//Save method
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//very important to defer and cloe statement if we have an error
	defer stmt.Close()

	user.DateCreated = utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return errors.ParseError(saveErr)
		/*
			if strings.Contains(err.Error(), indexUniqueEmail) {
				return errors.NewBadRequest(
					fmt.Sprintf("email %s already exists", user.Email))
			}
			return errors.NewInternalServerError(fmt.Sprintf("error while attempting to save user: %s", err.Error()))
		*/
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.ParseError(saveErr)
		//return errors.NewInternalServerError(fmt.Sprintf("error while attempting to save user: %s", err.Error()))
	}

	/*
		currentUser := userDB[user.ID]
		if currentUser != nil {
			if currentUser.Email == user.Email {
				return errors.NewBadRequest(fmt.Sprintf("A user with the email %s is already registered", user.Email))
			}
			return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.ID))

		}
		user.DateCreated = utils.GetNowString()
		//now := time.Now()
		//user.DateCreated = now.Format("02-01-2006T15:04:05Z")
		userDB[user.ID] = user
	*/
	user.ID = userID
	return nil

}
