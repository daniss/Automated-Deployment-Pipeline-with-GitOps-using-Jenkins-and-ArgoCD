package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		user := []string{"user1", "user2", "user3"}
		c.JSON(200, gin.H{
			"users": user,
		})
	})
	r.Run(":8081")
}