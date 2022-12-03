package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", get)
	router.POST("/", post)
	router.PUT("/", put)
	router.DELETE("/", delete)

	router.Run("localhost:8080")
}

func get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "get called")
}

func post(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "post called")
}

func put(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, "put called")
}

func delete(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "delete called")
}
