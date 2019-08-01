package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8080", "address")
)

type WelcomeDate struct {
	MyDate string
	MyTime string
}

func main() {
	flag.Parse()
	http.HandleFunc("/home", HomePage)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(*addr, nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	thedate := &WelcomeDate{
		MyDate: time.Now().Format("2006-01-02"),
		MyTime: time.Now().Format("15:04:05"),
	}
	t, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		panic("parse home page template err")
	}
	t.Execute(w, thedate)

}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "please login")
}
