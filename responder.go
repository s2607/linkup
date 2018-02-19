package main

import (
	"database/sql"
	"fmt"
)

type responder struct {
	key int64
	responses []response
	fname string
	lname string
	dob int//TODO: time
	zip string
	suggestions []service
	nchan chan bool


}

func (r responder) update_suggestions() error {
	return nil
}

//DB stuff


func (o *responder) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into responder(key) values(NULL)")
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update responder set(fname, lname, dob,zip) = (?,?,?,?) where key = ?")
		checkErr(err)
		res, err := stmt.Exec(o.fname,o.lname,o.dob,o.zip,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		//TODO: composed collections
	}
	return nil
}
func (o *responder) sresponces(Db *sql.DB) error {
	for _,r := range o.responses {
		err :=r.Store(Db)
		if err != nil {
			return err
		}
		stmt, err := Db.Prepare("replace respondersresponses(okey,ikey) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(o.key,r.key)
		checkErr(err)
		if res == nil {
			return nil //happens if relation already stored
		}
	}
	return nil
}
func (o *responder) Get(Db *sql.DB) error{
	//var pwhash string
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
	} else {
		err := Db.QueryRow("select key,fname,lname,dob,zip from responder where key = ?", o.key).Scan(&o.key,  &o.fname, &o.lname, &o.dob, &o.zip)
		if err !=nil {o.key = 0; return err}
	}
	return nil
}

func (o *responder) Pkey() int64{
	return o.key
}
func (o *responder) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *responder) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *responder) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *responder) Readynchan() {
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
