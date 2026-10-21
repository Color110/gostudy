package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r:= gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	user:=r.Group("/user")
	{
		user.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login",
			})
		user.GET("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "register",
			})
		})
		})	

	}
	r.Run(":8080") // listen and serve on

	
}