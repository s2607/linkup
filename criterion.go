package main

import (
    "fmt"
    "database/sql"
    "errors"
)

type criterion struct {
	key int64
	aval int
	bval int
	regex string
	lval bool
	isnl bool
	inv bool
	conj bool
	nchan chan bool
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

//DB stuff
func (o *criterion) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into criterion(key) values(NULL)")

		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update criterion set(key , aval, bval, regex, lval, isnil, inv, conj) = (?,?,?,?,?,?,?,?) where key = ?")

		checkErr(err)
		res, err := stmt.Exec(o.key ,o.aval,o.bval,o.regex,o.lval,o.isnl,o.inv,o.conj,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		//TODO: composed collections
	}
	return nil
}
func (o *criterion) Get(Db *sql.DB) error{
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
		return errors.New("not implemented")
	} else {
		err := Db.QueryRow("select key , aval, bval, regex, lval, isnil, inv, conj from criterion where key = ?", o.key).Scan(&o.key,&o.aval,&o.bval,&o.regex,&o.lval,&o.isnl,&o.inv,&o.conj)

		if err !=nil {o.key = 0; return err}
	}
	return nil
}

func (o *criterion) Pkey() int64{
	return o.key
}
func (o *criterion) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *criterion) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *criterion) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *criterion) Readynchan() {
	o.nchan = make(chan bool)
}
