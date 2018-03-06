package main

import (
	"fmt"
	"database/sql"
)

type service struct {
	key int64
	criteria map [question] []criterion
	qlist []question
	name string
	description string
	url string
	nchan chan bool
}

func (s *service) check(rid  int) bool {
	return true
}

//DB stuff
func (o *service) Store(Db *sql.DB) error{
	fmt.Println("storing service")
	fmt.Print("qlist len")
	fmt.Println(len( o.qlist))
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into service(key) values(NULL)")
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update service set(key, name, description, url) = (?,?,?,?) where key = ?")
		checkErr(err)
		res, err := stmt.Exec( o.key, o.name, o.description, o.url, o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		//TODO: criteria
		return o.sqlist(Db)
	}
	return nil
}
func (o *service ) Get(Db *sql.DB) error{
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
	} else {
		err := Db.QueryRow("select key, name, description, url from service where key = ?", o.key).Scan(&o.key, &o.name, &o.description, &o.url)
		if err !=nil {o.key = 0; return err}
	}
	return o.getqlist(Db)
}

func (o *service) Pkey() int64{
	return o.key
}
func (o *service) Zkey(){
	o.key=0
}
//DB Collections 

func (o *service) sqlist(Db *sql.DB) error {
	for _,r := range o.qlist {
		err :=r.Store(Db)
		if err != nil {
			return err
		}
		fmt.Println("service lelm")
		stmt, err := Db.Prepare("replace into servicesquestion(okey,ikey) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(o.key,r.key)
		checkErr(err)
		if res == nil {
			fmt.Println("TODO:nothing")
		}
	}
	return nil
}

func (o *service) getqlist(Db *sql.DB) error {
	rows, err := Db.Query("select ikey from servicesquestion where okey = ?", o.key)
	checkErr(err)
	defer rows.Close()
	i :=0
	for rows.Next() {
		var k int64
		var q question
		err := rows.Scan(&k)
		checkErr(err)
		o.qlist=append(o.qlist,q)
		o.qlist[i].key = k
		err = o.qlist[i].Get(Db)
		checkErr(err)
		i=i+1
	}
	return nil
}
//DB Sync stuff
func (o *service) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *service) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *service) Readynchan() {
	o.nchan = make(chan bool)
}
