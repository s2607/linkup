package main

import (
	"database/sql"
	"fmt"
)

type responder struct {
	key int64
	responses []*response
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
func Getallmatch (fname string,lname string,dob int,zip string) (error, []responder){
	nchan := make(chan error)
	r := make([]responder,0,0)
	DBchan <- func() {
		rows, err := Db.Query("select key from responder where fname = ? and lname = ? and dob = ? and zip = ?", fname,lname,dob,zip)
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			err := rows.Scan(&k)
			if err != nil {
				nchan <- err
			}
			r[i] = new(response)
			r[i].key = k
			err =r[i].Get(Db)
			if err != nil {
				nchan <- err
			}
			i=i+1
		}
		return func() {
			nchan <-err
		}
	}
	return <-nchan,r
}
func (o *responder) Tohtml() string {
	return "<div id=responder>"+o.fname+" ID:"+strconv.itoa(o.key)+"<form method=\"post\"> <input type=\"hidden\" name=\"rkey\" value=\""+strconv.itoa(o.key)+"\"><input type=submit></form>"+ "</div>"
}


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
		return o.sresponces(Db)
	}
	return nil
}
func (o *responder) sresponces(Db *sql.DB) error {
	for _,r := range o.responses {
		err :=r.Store(Db)
		if err != nil {
			return err
		}
		stmt, err := Db.Prepare("replace respondersresponse(okey,ikey) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(o.key,r.key)
		checkErr(err)
		if res == nil {
			return nil //happens if relation already stored
		}
	}
	return nil
}
func (o *responder) getresponses(Db *sql.DB) error {
	rows, err := Db.Query("select ikey from respondersresponse where okey = ?", o.key)
	checkErr(err)
	defer rows.Close()
	i :=0
	for rows.Next() {
		var k int64
		err := rows.Scan(&k)
		checkErr(err)
		o.responses[i] = new(response)
		o.responses[i].key = k
		err = o.responses[i].Get(Db)
		checkErr(err)
		i=i+1
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
	return o.getresponses(Db)
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
