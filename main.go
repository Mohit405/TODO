package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func login(w http.ResponseWriter, r *http.Request) {
	a := r.Method
	if a == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Printf("username is: %s", r.Form["username"])
		fmt.Printf("password is:%s", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
