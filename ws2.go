package main

import (
	"math/rand"
	"net/http"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)
func Statichandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
	fmt.Print(r.URL.Path)
}

func main() {//main always feels ugly and hacky
//this one in particular though 
	http.HandleFunc("/auth", Authhandler)
	http.HandleFunc("/", Statichandler)
	http.HandleFunc("/index.html", Statichandler)
	rand.Seed(4)//a random number
	fmt.Println("start")
	go Dbwriter()
	defer func(){Killchan <-true}()//lol
	o := new(operator)
	var op storable
	o.Getbyname("swiley")//hehhhhhh
	if o.key == 0 {
		fmt.Println("Creating default user")
		Init(o)
		o.uname = "swiley"
		o.setpss("abc123")
		Sstore(o)
	}
	op = o
	fmt.Println(op)
	//Db = initdb()
	//defer Db.Close()
	fmt.Println("starting web server")
	err := http.ListenAndServe(":80", nil)
	fmt.Println(err)
}
