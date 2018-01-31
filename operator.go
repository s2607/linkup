package main

import (
	"database/sql"
)

type operator struct {
	key int64
	pwhash []byte
	uname string
	cursessionid int
}

func (o *operator) Auth(pw string) bool{
	return false
}
func (o *operator) Checksesh(sesh int) bool{
	return false
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
func (o *operator) Getcomposed() error{
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



