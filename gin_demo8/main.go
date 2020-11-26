package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userinfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		var user userinfo
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.Run()
}
