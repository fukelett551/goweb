package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware1, middleware2)
	r.GET("/index", indexHandler)
	r.Run()
}

func indexHandler(c *gin.Context) {
	value, _ := c.Get("key")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    value,
	})
}

func middleware1(c *gin.Context) {
	fmt.Println("middleware1")
	start := time.Now()
	c.Set("key", "value1")
	//先去执行后面的handler然后回来
	c.Next()
	cost := time.Since(start)
	fmt.Println("cost time", cost)
}
func middleware2(c *gin.Context) {
	fmt.Println("middleware2")
	start := time.Now()
	c.Set("key", "value2")
	//先去执行后面的handler然后回来
	c.Next()
	//后续的不执行了，所以ok没返回
	//c.Abort()
	cost := time.Since(start)
	fmt.Println("cost time", cost)
}
