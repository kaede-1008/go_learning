package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"html/template"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil{
		log.Fatal(err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(writer, "Hello astaxie")
}
func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	if request.Method == "GET" {
		t, _ :=template.ParseFiles("login.gtpl")
		t.Execute(writer, nil)
	} else {
		// ログインデータがリクエストされ，ログインのロジック判断が実行される
		// request.Form["username"]はrequest.FormValue("username")ともかける
		// FormValueの場合はParseFormを呼ぶ必要はなくなる
		// FormValueはデータの中から1つめのものだけを返す
		// request.ParseForm()
		// fmt.Println("username:", request.Form["username"])
		// fmt.Println("password:", request.Form["password"])
		fmt.Println("username:", request.FormValue("username"))
		fmt.Println("password:", request.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}