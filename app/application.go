package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApplication is the starting point bein called by the main go file
func StartApplication() {
	mapUrls()
	router.Run(":8080")

}

/*
r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
*/
