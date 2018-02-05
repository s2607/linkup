package main

import (
    "fmt"
    "net/http"
    "database/sql"
    //"math/rand"
    _ "github.com/mattn/go-sqlite3"
)


var Db *sql.DB


func mktab (d *sql.DB, name string, cols map[string]string) {
	s := "CREATE TABLE IF NOT EXISTS " + name +" ("
	i := 0
	for k,v := range cols {
		i = i + 1
		if(k == "key") {
			v = "integer primary key asc" //See sqlite documentation.
		}
		s += k + " " + v
		if i != len(cols) {
			s += ","
		}
	}
	s += ")"
	fmt.Println(s)
	stmt, err := d.Prepare(s)
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	checkErr(err)
}
func initdb () *sql.DB {
	d, err := sql.Open("sqlite3","./wdb.db")
	checkErr(err)
	o := new(operator)
	mktab(d,"service",map[string]string{
		"key":"int",
		"name":"string",
		"description":"string",
		"url":"string",
	})
	mktab(d,"operator",map[string]string{
		"key":"int",
		"uname":"string",
		"pwhash":"string",
		"cursessionid":"int",
	})
	mktab(d,"responder",map[string]string{
		"key":"int",
		"fname":"string",
		"lname":"string",
		"zip":"string",
		"pwhash":"string",
		"dob":"int",//TODO: time
		"iqkey":"int",
	})
	mktab(d,"response",map[string]string{
		"key":"int",
		"value":"string",
		"qkey":"int",
	})
	mktab(d,"question",map[string]string{
		"key":"int",
		"prompt":"string",
		"qtype":"int",
	})
	mktab(d,"criterion",map[string]string{
		"key":"int",
		"aval":"int",
		"bval":"int",
		"lval":"bool",
		"isnil":"bool",
		"inv":"bool",
		"conj":"bool",
		"regex":"string",
		"qtype":"int",
	})
	//relation tables
	mktab(d,"respondersresponse",map[string]string{
		"key":"int",
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"respondersservice",map[string]string{
		"key":"int",
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"questionscriterion",map[string]string{
		"key":"int",
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"servicesquestioncriterion",map[string]string{
		"key":"int",
		"okey":"int",
		"iqkey":"int",
		"ickey":"int",
	})
//	go Dbwriter(d)
	//o.Init(d)
	o.key=8
	//o.uname = "swiley"
	err = o.Get(d)
	//err = o.Store(d)
	checkErr(err)
	o.uname = "swiley2"
//	o.Sstore()
	fmt.Println(o)
	return d
}
func Authhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handling auth", r.URL.Path[1:])
	var o operator
	if o.Getbyname(Db,r.FormValue("id")) != nil {

	}
	if o.Auth(r.FormValue("id")) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<body>Auth successfull.</body>.\n"))
	}else {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<body>Auth failed.</body>.\n"))
	}

}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

