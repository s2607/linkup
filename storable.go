package main

import (
//	"reflect"
	"database/sql"
//	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
type storable interface {
	Store() error
	Init() error
	Getcomposed() error
	Get() error
	PKey() int64
	Notify() error
	Wait() error
}


/*func Getallmatch(name string, o *storable, cols map[string] string )[]*storable {
	//WARNING: the side affect of this is that all data in the reciver is overwritten with the last matching row
	//TODO: sanatize? there's probably a better way...
	//Because of the abouve, this is really only a utility method that storable objects should use
	//to implement more concrete selection methods.
	s := "select * from " + name +" where "
        i := 0
        for k,v := range cols {
                i = i + 1
                s += k + " = " + v
                if i != len(cols) {
                        s += " and "
                }
        }
	rows, err := Db.Query("s")
	checkErr(err)
	rs := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(*o)),0,0)
	rsp := new([]*storable)
	for rows.Next() {
		(*o).Fromrow(rows)
		append(rs,o)
		append(*rsp,(rs[len(rs)-1]))
	}
	rows.Close()//closed explicitly so we can recurse
	for i = 0; i<len(rs); i++ {
		rs[i].getcomposed()
	}
	return *rsp
}*/

