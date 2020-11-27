package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB
var err error

func initMySQLDB() error {
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open db error ", err.Error())
	}
	return DB.DB().Ping()
}

func main() {

	if err := initMySQLDB(); err != nil {
		fmt.Println("connect db error")
		return
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./template/*")
	//index
	r.GET("/index", indexHandler)

	//待办事项todo
	v1group := r.Group("/v1")
	{
		v1group.POST("/todo", createHandler)
		v1group.DELETE("/todo/:id", deleteHandler)
		v1group.GET("/todo", queryAllHandler)
		v1group.GET("/todo/:id", queryOneHandler)
		v1group.PUT("/todo/:id", updateHandler)
	}

	r.Run()
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func createHandler(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})

	} else {
		if err := DB.Debug().Create(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}
func deleteHandler(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := DB.Debug().Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	} else {
		c.ShouldBind(&todo)
		if err := DB.Delete(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}
func queryAllHandler(c *gin.Context) {
	var todos []Todo
	if err := DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}
func queryOneHandler(c *gin.Context) {

}
func updateHandler(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := DB.Debug().Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	} else {
		c.ShouldBind(&todo)
		if err := DB.Save(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}
