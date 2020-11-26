package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   uint64
	Name string
	Age  uint
}

func (User) TableName() string {
	return "user_demo3"
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("open db error ", err.Error())
	}

	//db.Table("user_demo3").AutoMigrate(&User{})

	u := User{
		Name: "mario",
		Age:  10,
	}

	fmt.Println(db.NewRecord(&u))
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u))

}
