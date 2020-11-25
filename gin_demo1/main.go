package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.GET("/book",getRequest)
	r.POST("/book",postRequest)
	r.PUT("/book",putRequest)
	r.DELETE("/book",deleteRequest)
	r.Run()
}

func getRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "get",
	})
}
func postRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "post",
	})
}
func putRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "put",
	})
}
func deleteRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "delete",
	})
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello gin",
	})
}
