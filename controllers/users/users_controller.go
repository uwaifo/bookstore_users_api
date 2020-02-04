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

//CreateUser . .
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
	fmt.Println(result)

}

//GetUser . .
func GetUser(c *gin.Context) {

	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
	//c.String(http.StatusNotImplemented, "imppliment me na !")

}

/*
//SearchUser . . .
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "imppliment me na !")

}
*/
