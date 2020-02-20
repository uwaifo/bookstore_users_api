package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uwaifo/bookstore_users_api/domain/users"
	"github.com/uwaifo/bookstore_users_api/services"
	"github.com/uwaifo/bookstore_users_api/utils/errors"
)

//test

func TestServiceInterface() {
	//services.UserService.D

}

//get a user by id
func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequest("user id should be a number")
	}
	return userID, nil

}

//CreateUser . .
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
	//fmt.Println(result)

}

//GetUser . .
func GetUser(c *gin.Context) {

	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	/*
		userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if userErr != nil {
			err := errors.NewBadRequest("invalid user id")
			c.JSON(err.Status, err)
			return
		}
	*/

	user, getErr := services.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
	//c.String(http.StatusNotImplemented, "imppliment me na !")

}

//UpdateUser . .
func UpdateUser(c *gin.Context) {
	//Do the same body parsing and validation as in GetUser function
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	/*
		userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if userErr != nil {
			err := errors.NewBadRequest("invalid user id")
			c.JSON(err.Status, err)
			return
		}
	*/
	//same as in SaveUser
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	//NOW we add a user id to the use json object to be updated
	user.ID = userID

	//Swith http method if the request is a PATCH
	isPartial := c.Request.Method == http.MethodPatch

	//pass the properly composed json object to the UpdateUser service
	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
	fmt.Println(result)

}

//Delete . .. .
func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.UserService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	/*VERY IMPORTANT TO RETURN THE SAME CONTENT TYPE IN
	SUCCESS AS IN ERRORS
	IE. if we return c.JSON... for failure we can not later use
	 c.String for success becouse we realize we not have a payload
	 They must be the same.

	*/

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

//Search . . . .
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}

/*
//SearchUser . . .
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "imppliment me na !")

}
*/
