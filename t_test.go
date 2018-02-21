package main
import (
	"fmt"
	"testing"
	"database/sql"
	"strconv"

func TestNewop(t *testing.T) {
	go Dbwriter()
	o := new(operator)
	p := new(operator)
	//q := new(operator)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.uname="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println("eh")
	if p.uname != o.uname {
		fmt.Println("Could not retrieve operator")
		t.Fail()
	}
//	q.Getbyname(o.uname)
//	fmt.Println("key:"+strconv.FormatInt(q.key,10))
//	if q.key != o.key {
//		fmt.Println("Could not retrieve operator by name")
//		t.Fail()
//	} //XXX Broken?
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from operator where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}

func TestNewres(t *testing.T) {
	go Dbwriter()
	o := new(responder)
	p := new(responder)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.fname="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println("eh")
	if p.fname != o.fname {
		fmt.Println("Could not retrieve responder")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from responder where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
func TestNewresponse(t *testing.T) {
	go Dbwriter()
	o := new(response)
	p := new(response)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.value="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println(p.key)
	if p.value != o.value {
		fmt.Println("Could not retrieve response")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from response where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
func TestNewquestion(t *testing.T) {
	go Dbwriter()
	o := new(question)
	p := new(question)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.prompt="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println(p.key)
	if p.prompt!= o.prompt{
		fmt.Println("Could not retrieve question")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from question where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
func TestNewcriterion(t *testing.T) {
	go Dbwriter()
	o := new(criterion)
	p := new(criterion)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.regex="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println(p.key)
	if p.regex!= o.regex{
		fmt.Println("Could not retrieve criterion")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from question where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
func TestNewservice(t *testing.T) {
	go Dbwriter()
	o := new(service)
	p := new(service)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.name="swiley-test"
	Sstore(o)
	p.key=o.key
	Sget(p)
	fmt.Println(p.key)
	if p.name!= o.name{
		fmt.Println("Could not retrieve service")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from question where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
func Testcompresponder(t *testing.T) {
	go Dbwriter()
	o := new(responder)
	p := new(responder)
	q := new(response)
	Init(o)
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	Init(q)
	if q.key ==0 {
		fmt.Println("composed key of zero")
		t.Fail()
	}
	Sstore(o)
	p.key=o.key
	Sget(p)
	if p.responses[0].key!= o.responses[0].key {
		fmt.Println("Could not retrieve responses")
		t.Fail()
	}
	DBchan <- func (Db *sql.DB)func() {
		stmt, err := Db.Prepare("delete from responses where key = ?")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		stmt.Exec(o.key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		return func(){}//Killchan provides needed syncronisation
	}
	Killchan<-true
}
