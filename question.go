package main

import (
	"database/sql"
	"fmt"
)

type question struct {
	key int64
	prompt string
	qtype int
	clist []*criterion
	nchan chan bool
}

func (q *question) New_response ()  *response {
	r := new (response)
	Init(r)//TODO: check if there's already a response for that responder
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
		return o.sclist(Db)
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
	return o.getclist(Db)
}

func (o *question) Pkey() int64{
	return o.key
}
func (o *question) Zkey(){
	o.key=0
}
//DB collections

func (o *question) sclist(Db *sql.DB) error {
	for _,r := range o.clist {
		err :=r.Store(Db)
		if err != nil {
			return err
		}
		stmt, err := Db.Prepare("replace into questionscriterion(okey,ikey) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(o.key,r.key)
		checkErr(err)
		if res == nil {
			fmt.Println("TODO: nothing")
		}
	}
	return nil
}
func (o *question) getclist(Db *sql.DB) error {
	rows, err := Db.Query("select ikey from questionscriterion where okey = ?", o.key)
	checkErr(err)
	defer rows.Close()
	i :=0
	for rows.Next() {
		var k int64
		r := new(criterion)
		err := rows.Scan(&k)
		checkErr(err)
		r.key = k
		err = r.Get(Db)
		o.clist = append(o.clist ,r)
		checkErr(err)
		i=i+1
	}
	return nil
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
