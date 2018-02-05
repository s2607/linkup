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

func main() {
	http.HandleFunc("/auth", Authhandler)
	http.HandleFunc("/", Statichandler)
	http.HandleFunc("/index.html", Statichandler)
	rand.Seed(4)//a random number
	Db = initdb()
	defer Db.Close()
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
