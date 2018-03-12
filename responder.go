package main

import (
	"database/sql"
	"fmt"
	"strconv"
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

func (r *responder) update_suggestions() error {
	err,se :=Getallservices()
	checkErr(err)
	r.suggestions = nil //release for garbage collection
	for _,s := range se {
		fmt.Print("checking service:"+s.name)
		if Validate(r.responses,s.criteria) {
			fmt.Println(" yes")
			r.suggestions = append(r.suggestions,*s)
		}else {
			fmt.Println(" no")
		}
	}
	return nil
}

//DB stuff
func Getallmatch (fname string,lname string,dob int,zip string) (error, []*responder){
	nchan := make(chan error)
	var r []*responder
	DBchan <- func(Db *sql.DB)func() {
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
			r = append(r,new(responder))
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
	return "<div class='responder_entry'>"+o.fname+" ID:"+strconv.FormatInt(o.key,10)+"<form method=\"post\"> <input type=\"hidden\" name=\"rkey\" value=\""+strconv.FormatInt(o.key,10)+"\"><input type=submit id='submit_button' value='Select'></form>"+ "</div>"
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
		stmt, err := Db.Prepare("replace into respondersresponse(okey,ikey) values(?,?)")
		checkErr(err)
		res, err := stmt.Exec(o.key,r.key)
		checkErr(err)
		if res == nil {
			fmt.Println("TODO: nothing")
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
		r := new(response)
		err := rows.Scan(&k)
		checkErr(err)
		r.key = k
		err = r.Get(Db)
		o.responses = append(o.responses,r)
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
