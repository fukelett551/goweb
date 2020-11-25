package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/login",func (c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//defaultform 设置参数默认值
		//getpostform 获取参数的值 不存在返回false
		
		c.JSON(http.StatusOK,gin.H{
			"username" : username,
			"password" : password,
		})
	})
	r.Run()
}
