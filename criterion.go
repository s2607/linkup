package main

import (
//    "fmt"
 //   "net/http"
  //  "database/sql"
 //   _ "github.com/mattn/go-sqlite3"
)

type criterion struct {
	key int
	aval int
	bval int
	regex string
	lval bool
	isnl bool
	inv bool
	conj bool
}

func (o *criterion) checkstr(v string) bool{
	return false
}
func (o *criterion) checkint(v int) bool{
	return false
}

func (o *criterion) checkbool(v bool) bool{
	return false
}

