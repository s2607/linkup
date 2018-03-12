package main

import (
	"math/rand"
	"net/http"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)
func Statichandler(w http.ResponseWriter, r *http.Request) {//Should we really keep it?
	http.ServeFile(w, r, "./"+r.URL.Path)
	fmt.Print(r.URL.Path)
}
func sometestdata() {
	q := new(question)
	p := new(question)
	Init(q)
	q.prompt = "What is your mothers madin name"
	fmt.Println(q)
	checkErr(Sstore(q))
	Init(p)
	p.prompt = "What is your fathers madin name"
	fmt.Println(p)
	checkErr(Sstore(p))
	o := new(operator)
	var op storable
	o.Getbyname("swiley")//hehhhhhh
	if o.key == 0 {
		fmt.Println("Creating default user")
		Init(o)
		o.uname = "swiley"
		o.setpss("abc123")
		o.cser = new(service)
		o.cser.name = "cfaw"
		Init(o.cser)
		o.cser.qlist=append(o.cser.qlist,*p)
		o.cser.qlist=append(o.cser.qlist,*q)
		Sstore(o)
	}
	op = o
	fmt.Println(op)
}

func main() {//main always feels ugly and hacky
//this one in particular though 
	http.HandleFunc("/auth", Authhandler)
	http.HandleFunc("/newr", Ursession_handler)
	http.HandleFunc("/", Statichandler)
	http.HandleFunc("/index.html", Authhandler)
	http.HandleFunc("/qprompt", qprompt_handler)
	http.HandleFunc("/qprompt/", qprompt_handler)
	http.HandleFunc("/sugs/", sugg_handler)
	rand.Seed(4)//a random number
	fmt.Println("start")
	go Dbwriter()
	defer func(){Killchan <-true}()//lol
	sometestdata()
	fmt.Println("starting web server")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
