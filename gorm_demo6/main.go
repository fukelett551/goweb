package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  uint8
}

func (User) TableName() string {
	return "user_demo6"
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("open db error ", err.Error())
	}

	db.AutoMigrate(&User{})
	/*	u1 := User{
			Name: "mario",
			Age:  18,
		}
		db.Create(&u1)
		u2 := User{
			Name: "kinopio",
			Age:  10,
		}
		db.Create(&u2)*/
	var u User
	u.ID = 1
	//软删除，更新删除时间，并非真正删除数据
	db.Delete(&u)
	var u1 User
	db.Debug().Where("name = ?", "kinopio").Delete(&u1)

}
