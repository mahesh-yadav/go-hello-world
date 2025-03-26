package main

import (
	"net/http"
	"strconv"

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

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func QueryParamsHandler(c *gin.Context) {
	name := c.Query("name")
	ageStr := c.Query("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age parameter"})
		return
	}

	c.JSON(http.StatusOK, User{
		Name: name,
		Age:  age,
	})
}

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)

	router.GET("/:name", PathParamsHandler)

	router.GET("/query", QueryParamsHandler)

	router.Run(":3000")
}
