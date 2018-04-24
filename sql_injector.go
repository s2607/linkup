package main

import (
	"database/sql"
	"fmt"
)

func Sql_injector(p string, m *string) error {

	fmt.Println("got:" + p)
	if p == "" {
		return nil
	}

	nchan := make(chan error)
	DBchan <- func(Db *sql.DB) func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error with SQL STATEMENT! - ", err)
                *m = "Error With SQL Statement"
			}
		}()
		Db.Exec(p)
		//checkErr(err)
        *m = "Submitted"
        fmt.Println("msg: " + *m)
		return func() {
			nchan <- nil
		}
	}
	return <-nchan
}
