
package main

import (
    "fmt"
    "database/sql"
    "errors"
)

type validresponse struct {
	key int64
	qkey int64
	text string
	nchan chan bool
}


//DB stuff
func (o *validresponse) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into validresponse(key) values(NULL)")

		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update validresponse set(key , text, qkey) = (?,?,?) where key = ?")
		checkErr(err)
		res, err := stmt.Exec(o.key ,o.text,o.qkey,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		//TODO: composed collections
	}
	return nil
}
func (o *validresponse) Get(Db *sql.DB) error{
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
		return errors.New("not implemented")
	} else {
		err := Db.QueryRow("select key , text, qkey from validresponse where key = ?", o.key).Scan(&o.key,&o.text,&o.qkey,&o.key)

		if err !=nil {o.key = 0; return err}
	}
	return nil
}

func (o *validresponse) Pkey() int64{
	return o.key
}
func (o *validresponse) Ptext() string{
	return o.text
}
func (o *validresponse) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *validresponse) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *validresponse) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *validresponse) Readynchan() {
	o.nchan = make(chan bool)
}
