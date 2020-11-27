package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name   string
	Age    uint8
	Active bool
}

func (User) TableName() string {
	return "user_demo5"
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("open db error ", err.Error())
	}

	db.AutoMigrate(&User{})
	/*	u1 := &User{
			Name:   "mario",
			Age:    18,
			Active: false,
		}
		db.Create(&u1)
		u2 := &User{
			Name:   "kinopio",
			Age:    10,
			Active: true,
		}
		db.Create(&u2)
		var user User
		db.First(&user)
		user.Name = "kinopiko"
		//默认会修改所有字段 id，create_at ...
		db.Save(user)*/
	var user User
	db.Debug().Model(&user).Where("age = ? ", 18).Update("age", 33)

}
