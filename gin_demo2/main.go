package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var (
		err error
	)
	http.HandleFunc("/templatedemo1", tempalteDemo1)
	http.HandleFunc("/templateDemo2", tempalteDemo2)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func tempalteDemo1(w http.ResponseWriter, r *http.Request) {
	var (
		templ *template.Template
		err   error
		name  string
	)
	templ, err = template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	name = "mario"
	templ.Execute(w, name)
}

func tempalteDemo2(w http.ResponseWriter, r *http.Request) {
	var (
		templ *template.Template
		err   error
		user  User
	)

	templ, err = template.ParseFiles("./templateUser.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user = User{
		Name:    "mario",
		Age:     10,
		Address: "日本",
	}
	templ.Execute(w, user)
}
