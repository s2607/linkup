package main

import (
	"database/sql"
	"fmt"
)

type question struct {
	key int64
	prompt string
	qtype int
	nchan chan bool
}

func (q *question) New_response ()  *response {
	r := new (response)
	Init(r)
	r.q=q
	return r
}
//Should have been better with visibility...
func (o *question)Pprompt() string {
	return o.prompt
}

//DB stuff
func (o *question) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into question(key) values(NULL)")
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update question set(prompt, qtype) = (?,?) where key = ?")
		checkErr(err)
		res, err := stmt.Exec(o.prompt,o.qtype,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		//TODO: composed collections
	}
	return nil
}
func (o *question) Get(Db *sql.DB) error{
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
	} else {
		err := Db.QueryRow("select key,prompt,qtype from question where key = ?", o.key).Scan(&o.key,  &o.prompt, &o.qtype)
		if err !=nil {o.key = 0; return err}
	}
	return nil
}

func (o *question) Pkey() int64{
	return o.key
}
func (o *question) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *question) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *question) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *question) Readynchan() {
	o.nchan = make(chan bool)
}
//old stuff
/*func (o *operator) Sstore() error{
	o.nchan = make(chan bool)
	//Wrchan <-o
	DBchan <- func (Db *sql.DB)func() {
		o.Store(Db)
		return o.Notify
	}
	o.Wait()
	return nil
}*/
/*
func (o *operator) Sget() error {
	o.nchan = make(chan bool)
	DBchan <- func (Db *sql.DB)func() {
		o.Get(Db)
		return o.Notify
	}
	o.Wait()
	return nil
}/*
func (o *operator) Init() error {
	o.key = 0
	return o.Sstore()
}
**/
