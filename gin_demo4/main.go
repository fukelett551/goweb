package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//使用map 等同于 gin.H{}
		data := map[string]interface{}{
			"name": "mario",
			"age":  "18",
		}
		data = data
		//使用结构体 tag用于反射返回前端解析
		type msg struct {
			Name    string `json:"name"`
			Age     int    `json:"age"`
			Address string `json:"address"`
		}

		m := msg{
			Name:    "mario",
			Age:     18,
			Address: "日本",
		}
		c.JSON(http.StatusOK, m)
	})

	r.Run()
}
