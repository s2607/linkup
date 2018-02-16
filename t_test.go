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
	q := new(operator)
	o.Init()
	if o.key ==0 {
		fmt.Println("key of zero")
		t.Fail()
	}
	fmt.Println("key:"+strconv.FormatInt(o.key,10))
	o.uname="swiley-test"
	o.Sstore()
	p.key=o.key
	p.Sget()
	fmt.Println("eh")
	if p.uname != o.uname {
		fmt.Println("Could not retrieve operator")
		t.Fail()
	}
	q.Getbyname(o.uname)
	fmt.Println("key:"+strconv.FormatInt(q.key,10))
	if q.key != o.key {
		fmt.Println("Could not retrieve operator by name")
		t.Fail()
	}
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
