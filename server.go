package main

import "github.com/gin-gonic/gin"

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Go Gin API!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)

	router.Run(":3000")
}
