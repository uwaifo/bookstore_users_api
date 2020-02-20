package users

import (
	"fmt"
	//"errors"

	"github.com/uwaifo/bookstore_users_api/datasource/mysql/usersdb"
	"github.com/uwaifo/bookstore_users_api/logger"
	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//used for interacting with the database

const (
	indexUniqueEmail  = "email_UNIQUE"
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser      = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUdpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser   = "DELETE FROM users WHERE id=?;"
	queryUserByStatus = "SELECT id, first_name,last_name, email, date_created, status FROM users WHERE status=?;"
	errorNoRows       = "no rows in result set"
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
		logger.Error("error when trying to prepaire get user statment", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		logger.Error("error attempting to get user by id", err)
		return errors.NewInternalServerError("database error")
		//fmt.Println(err)
		/*UNSURE
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))

		}

		return errors.NewInternalServerError(
			fmt.Sprintf("error attempting to get user %d: %s", user.ID, err.Error()))
		*/
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

//Delete  . . .
func (user *User) Delete() *errors.RestErr {
	//prepare and execute the delete query
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//
	if _, err = stmt.Exec(user.ID); err != nil {
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

	//user.DateCreated = utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

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

//FindByStatus method . . . .
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {

	stmt, err := usersdb.Client.Prepare(queryUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())

	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError((fmt.Sprintf("no users matching status %s", status)))
	}
	return results, nil

}
