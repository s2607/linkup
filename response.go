package main

import (

	"database/sql"
	"fmt"
	"errors"
)
type response struct {
	key int64
	value string
	q *question
	nchan chan bool
}




//DB stuff

func (o *response) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into response(key) values(NULL)")
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		fmt.Println("storing"+o.value)
		stmt, err := Db.Prepare("update response set(value, qkey) = (?,?) where key = ?")
		checkErr(err)
		var rk int64
		if(o.q!= nil) {
			rk = o.q.key
		}else {
			rk = 0
		}
		res, err := stmt.Exec(o.value,rk,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		if o.q!= nil {
			//return o.q.Store(Db) //why ever?
			return nil
		}else {
			return nil
		}
	}
	return errors.New("bad store call")
}
func (o *response) Get(Db *sql.DB) error{
	//var pwhash string
	var rkey int64
	err := Db.QueryRow("select key, value, qkey from response where key = ?", o.key).Scan(&o.key,  &o.value, &rkey)
    if err !=nil {
        o.key=0
        fmt.Print(err.Error())
        return err
    }
	checkErr(err)
	if rkey != 0 {
		o.q= new(question)
		o.q.key=rkey
		err = o.q.Get(Db)
		if err != nil {
			o.key = 0
			return err
		}
	}
	return nil
}
func (o *response) Readynchan() {
	o.nchan = make(chan bool)
}

func (o *response) Pkey() int64{
	return o.key
}
func (o *response) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *response) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}

func (o *response) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
