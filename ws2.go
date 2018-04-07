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
	l := new(criterion)
	m := new(criterion)
	Init(m)
	m.q=nil//implies same question to prevent infinitely nested get calls
	m.regex="mom"
	p.clist=append(p.clist,m)
	fmt.Println(m)
	Init(q)
	q.prompt = "What is your mothers madin name"
	fmt.Println(q)
	checkErr(Sstore(q))
	Init(p)
	p.prompt = "What is your fathers madin name"
	fmt.Println(p)
	Init(l)
	l.q=q
	l.regex="dad"
	fmt.Println(l)
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
		o.cser.url= "http://www.liberty.edu"
		o.cser.description= "college for a weekend"
		Init(o.cser)
		o.cser.qlist=append(o.cser.qlist,*p)
		o.cser.qlist=append(o.cser.qlist,*q)
		o.cser.criteria=append(o.cser.criteria,l)
		Sstore(o)
	}
	op = o
	fmt.Println(op)
}

func main() {//main always feels ugly and hacky
	http.HandleFunc("/auth", Authhandler)
	http.HandleFunc("/newr", Ursession_handler)
	http.HandleFunc("/", Statichandler)
	http.HandleFunc("/index.html", Authhandler)
	http.HandleFunc("/qprompt", qprompt_handler)
	http.HandleFunc("/qprompt/", qprompt_handler)
	http.HandleFunc("/sugs/", sugg_handler)
	//"service GUI"
	http.HandleFunc("/newop", opcreate_handler)
	http.HandleFunc("/newserv", servicecreate_handler)
	http.HandleFunc("/newq", questioncreate_handler)
	http.HandleFunc("/newc", criterioncreate_handler)
	http.HandleFunc("/delc", delc_handler)
	http.HandleFunc("/delq", delq_handler)
	http.HandleFunc("/searchq", searchq_handler)
	http.HandleFunc("/searchs", searchs_handler)
	http.HandleFunc("/searcho", searchs_handler)
	http.HandleFunc("/sql", sql_handler)
	rand.Seed(4)//a random number
	fmt.Println("start")
	go Dbwriter()
	defer func(){Killchan <-true}()//lol
	sometestdata()
	fmt.Println("starting web server")
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	fmt.Println(err)
}
