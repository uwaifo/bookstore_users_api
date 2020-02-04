package app

import (
	"github.com/uwaifo/bookstore_users_api/controllers/ping"
	"github.com/uwaifo/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	// notice that the Ping function is refrenced but not executed/called with ()

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)

}
