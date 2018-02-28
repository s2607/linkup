package main

import (
	"database/sql"
	"math/rand"
	"fmt"
	"encoding/hex"
	"crypto/md5"
)

type operator struct {
	key int64
	pwhash [16]byte
	uname string
	cursessionid int
	nchan chan bool
	cresp *responder
}
func (o *operator) setpss(pw string) error {
	o.pwhash=md5.Sum([]byte(pw))
	return nil
}

func (o *operator) Auth(pw string) bool{
	//TODO: pss auth
	if md5.Sum([]byte(pw)) != o.pwhash {
		return false
	}
	o.cursessionid=rand.Int()
	return true
}
func (o *operator) Checksesh(sesh int) bool{
	if o.cursessionid == sesh { return true}
	return false
}
func (o *operator) Getbyname(name string) error{
	o.Readynchan()
	o.key = 0
	o.uname = name
	DBchan <- func (Db *sql.DB)func() {
		o.Get(Db)
		fmt.Print("Got an:")
		fmt.Println(o.key)
		return o.Notify
	}
	o.Wait()
	fmt.Println("Got get:"+o.uname)
	return nil
}

//DB stuff

func (o *operator) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into operator(key) values(NULL)")
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update operator set(pwhash, uname, cursessionid, cresp) = (?,?,?,?) where key = ?")
		checkErr(err)
		var rk int64
		if(o.cresp != nil) {
			rk = o.cresp.key
		}else {
			rk = 0
		}
		res, err := stmt.Exec(hex.EncodeToString(o.pwhash[:]),o.uname,o.cursessionid,rk,o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
		if o.cresp != nil {
			return o.cresp.Store(Db)
		}else {
			return nil
		}
	}
	return nil
}
func (o *operator) FuckGet(DB *sql.DB) error {
	var tempPass, tempCresp string
	if o.key == 0 {
		err :=  Db.QueryRow("select key from operator where uname = ?", o.uname).Scan(&o.key)a
	}
	err = GetData(map[string]*interface{}{
			"pwhash":tempPass,
			"cursessionid":&o.cursessionid,
			"cresp":tempCresp,
		}, "operator", o.key, DB)
	if err != nil {
		return err
	}
}

func (o *operator) Get(Db *sql.DB) error{
	//var pwhash string
	var rkey int64
	var phh string
	if o.key == 0 {
		fmt.Println("Getting "+o.uname)
		err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)
		fmt.Print("Got key: ")
		fmt.Println(o.key)
		if err !=nil {return err}
	} else {
		err := Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where key = ?", o.key).Scan(&o.key,  &o.uname, &o.cursessionid, &phh, &rkey)
		if err !=nil {o.key=0;return err}
	}
	ph, err := hex.DecodeString(phh)
	copy(o.pwhash[:],ph)
	if err !=nil {
		o.key = 0
		return err
	}
	if rkey != 0 {
		o.cresp = new(responder)
		o.cresp.key=rkey
		err = o.cresp.Get(Db)
		if err != nil {
			o.key = 0
			return err
		}
	}
	return nil
}
func (o *operator) Readynchan() {
	o.nchan = make(chan bool)
}

func (o *operator) Pkey() int64{
	return o.key
}
func (o *operator) Zkey(){
	o.key=0
}
//DB Sync stuff
func (o *operator) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}

func (o *operator) Notify() {
	fmt.Println("note")
	o.nchan <- true
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
