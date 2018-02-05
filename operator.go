package main

import (
	"database/sql"
	"math/rand"
	"fmt"
)

type operator struct {
	key int64
	pwhash []byte
	uname string
	cursessionid int
	nchan chan bool
}

func (o *operator) Auth(pw string) bool{
	//TODO: pss auth
	o.cursessionid=rand.Int()
	if o.Sstore() != nil {
		return false
	}
	return true
}
func (o *operator) Checksesh(sesh int) bool{
	return false
}
func (o *operator) Getbyname(Db *sql.DB, name string) error{
	var b int
	//var pwhash string
	return  Db.QueryRow("select key, uname, cursessionid from operator where uname = ?", name).Scan(&o.key,  &o.uname, &b)
}

//DB stuff

func (o *operator) Store(Db *sql.DB) error{
	stmt, err := Db.Prepare("insert into operator(key, pwhash, uname, cursessionid) values(?,?,?,?)")
        checkErr(err)
        res, err := stmt.Exec(o.key,o.pwhash,o.uname,o.cursessionid)
        checkErr(err)
	if res == nil {//XXX
		panic(err)//never happens?
	}
	return nil
}
func (o *operator) Init(Db *sql.DB) error {
	stmt, err := Db.Prepare("insert into operator(key) values(NULL)")
        checkErr(err)
        res, err := stmt.Exec()
        checkErr(err)
	o.key, err = res.LastInsertId()
        checkErr(err)
	return err
}
func (o *operator) Getcomposed(Db *sql.DB) error{
	//Operator has no composed collections
	return nil
}
func (o *operator) Get(Db *sql.DB) error{
	var b int
	//var pwhash string
	return  Db.QueryRow("select key, uname, cursessionid from operator where key = ?", o.key).Scan(&o.key,  &o.uname, &b)
}

func (o *operator) PKey() int64{
	return o.key
}
//DB Sync stuff
func (o *operator) Sstore() error{
	o.nchan = make(chan bool)
	Wrchan <-o
	o.Wait()
	return nil
}
func (o *operator) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}

func (o *operator) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
