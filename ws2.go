package main

import (
    "fmt"
    "net/http"
//    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/insert", ih)
	http.HandleFunc("/view", vh)
	http.HandleFunc("/useradd", uah)
	Db = initdb()
	defer Db.Close()
	http.ListenAndServe(":8080", nil)
}
func updateanswer(key string, val string, user string) {
	stmt, err := Db.Prepare("update answers set ?=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(key,val,user)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}
