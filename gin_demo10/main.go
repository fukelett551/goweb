package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/redirect_permanent", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "www.baidu.com")
	})

	r.GET("/redirect_temp", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect_permanent"
		r.HandleContext(c)
	})

	r.Run()
}
