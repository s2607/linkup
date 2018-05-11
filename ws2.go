package main

import (
	"math/rand"
	"net/http"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"net"
	"log"
)
func Statichandler(w http.ResponseWriter, r *http.Request) {//Should we really keep it?
	http.ServeFile(w, r, "./"+r.URL.Path)
	fmt.Print(r.URL.Path)
}
func sometestdata() {
	o := new(operator)
	var op storable
	o.Getbyname("swiley")//hehhhhhh
	if o.key == 0 {
		fmt.Println("Creating default user")
		Init(o)
		o.uname = "swiley"
		o.admin = true
		o.setpss("abc123")
        o.setAdmin("true")
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
	http.HandleFunc("/delc", delc_handler)
    http.HandleFunc("/delqc", delqc_handler)
	http.HandleFunc("/delq", delq_handler)
	http.HandleFunc("/searchq", searchq_handler)
	http.HandleFunc("/searchs", searchs_handler)
	http.HandleFunc("/searcho", searcho_handler)
    http.HandleFunc("/searchqid", searchqid_handler)
	http.HandleFunc("/sql", sql_handler)
	//Home and Help Button
	http.HandleFunc("/home", home_handler)
    http.HandleFunc("/help", help_handler)
	rand.Seed(4)//a random number
	fmt.Println("start")
	go Dbwriter()
	defer func(){Killchan <-true}()//lol
	sometestdata()
	fmt.Println("starting web server")
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS("0.0.0.0:8090", "server.crt", "server.key", nil)
//	server := &http.Server{Handler: nil}
//	l, err := net.Listen("tcp4", "0.0.0.0:8090")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	err = server.Serve(l)
	fmt.Println(err)
}
