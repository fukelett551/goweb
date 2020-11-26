package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
}

func (User) TableName() string {
	return "user_demo4"
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("open db error ", err.Error())
	}
	var user User
	//db.First(&user)
	db.FirstOrInit(&user, &User{
		Name: "kinopiko",
		Age:  10,
	})
	db.Create(&user)
	fmt.Println(user)
}
