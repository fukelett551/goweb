package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("open db error", err.Error())
		return
	}
	fmt.Println("open db success")
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{
		Name:   "mario",
		Gender: "male",
		Hobby:  "jump",
	}
	//增
	db.Create(&u1)
	//查
	var u2 UserInfo
	db.First(&u2)
	fmt.Println(u2)
	//改
	db.Model(&u2).Update("hobby", "shot")
	//删
	db.Delete(&u2)
}
