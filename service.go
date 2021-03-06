package main

import (
	"fmt"
	"database/sql"
//	"regexp"
)

type service struct {
	key int64
	criteria []*criterion
	qlist []question
	name string
	description string
	url string
	nchan chan bool
}

func (s *service) check(rid  int) bool {
	return true
}
func (s *service) Pdesc() string {
	return s.description
}
func (s *service) Pclist() []*criterion {
	return s.criteria
}
func (s *service) Pqlist() []question{
	return s.qlist
}
//Template stuff
func (s *service)Pname() string {return s.name}
func (s *service)Pdescription() string {return s.description}
func (s *service)Purl() string {return s.url}

//DB stuff

func Getallsbyname (p string) (error, []*service){
	fmt.Println("got:"+p)
	if p=="" {return nil,nil}

	nchan := make(chan error)
	var r []*service
	DBchan <- func(Db *sql.DB)func() {
		//p="\"%"+p+"\"%"
		//rows, err := Db.Query("select key from service where name like \"%?%\"",p)
		rows, err := Db.Query("select key from service")
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			s := new(service)
			rows.Scan(&k)
			//checkErr(err)
			if k == 0 {continue}
			r=append(r,s)
			r[i].key = k
			r[i].Get(Db)
			//checkErr(err)
			i=i+1
		}
		return func() {
			nchan <-nil
		}
	}
	return <-nchan,r
}
func Getallsbynamefuzz (p string) (error, []*service){
	fmt.Println("got:"+p)
	if p=="" {return nil,nil}
	var r []*service
	search := "%" + p + "%"
	nchan := make(chan error)
	DBchan <- func(Db *sql.DB)func() {
		rows, err := Db.Query("select key from service where name like ? or description like ?", search, search)
		checkErr(err)
		defer rows.Close()
		for i := 0; rows.Next(); i++ {
			var k int64
			rows.Scan(&k)
			if k == 0 {continue}
			r = append(r, new(service))
			r[i].key = k
			err = r[i].Get(Db)
			if err != nil {
				fmt.Println("error")
				nchan <- err
			}
		}
		return func() {
			nchan <- nil
		}
	}
	return <-nchan, r
}

//DB stuff

func Getallservices() (error, []*service){
	nchan := make(chan error)
	var r []*service
	DBchan <- func(Db *sql.DB)func() {
		rows, err := Db.Query("select key from service")
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			s := new(service)
			rows.Scan(&k)
			//checkErr(err)
			if k == 0 {continue}
			r=append(r,s)
			r[i].key = k
			r[i].Get(Db)
			//checkErr(err)
			i=i+1
		}
		return func() {
			nchan <-nil
		}
	}
	return <-nchan,r
}
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
		err = o.sqlist(Db)
		if err != nil {
			return err
		}
		return o.sclist(Db)
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
	err := o.getqlist(Db)
	if err != nil {
		o.key = 0
		return err
	}
	return o.getclist(Db)

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
func (o *service) sclist(Db *sql.DB) error {
	for _,r := range o.criteria {
		err :=r.Store(Db)
		if err != nil {
			return err
		}
		fmt.Println("service lelc")
		stmt, err := Db.Prepare("replace into servicescriterion(okey,ikey) values(?,?)")
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
func (o *service) getclist(Db *sql.DB) error {
	rows, err := Db.Query("select ikey from servicescriterion where okey = ?", o.key)
	checkErr(err)
	defer rows.Close()
	i :=0
	for rows.Next() {
		var k int64
		k = 0
		q := new(criterion)
		//err := rows.Scan(&k)
		rows.Scan(&k)
		fmt.Println(k)
		//checkErr(err)
		if(k==0){continue}
		o.criteria=append(o.criteria,q)
		o.criteria[i].key = k
		err = o.criteria[i].Get(Db)
		//checkErr(err)//XXX XXX
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
