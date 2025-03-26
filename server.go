package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Go Gin API!",
	})
}

func PathParamsHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello, %s!", name)
}

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)

	router.GET("/:name", PathParamsHandler)

	router.Run(":3000")
}
