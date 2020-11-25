package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/param", func(c *gin.Context) {
		//通过query获取请求参数值
		str := c.Query("name")
		str = str
		//使用defaultquery设置默认值
		str1 := c.DefaultQuery("name", "kinopiko")
		str1 = str1
		//使用 getquery 获取是否传递某参数
		str2, ok := c.GetQuery("name")
		str2, ok = str2, ok
		//使用
		c.String(http.StatusOK, "the param is %s", str)
	})

	r.Run()
}
