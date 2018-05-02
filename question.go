package main

import (
	"database/sql"
    "strconv"
	"fmt"
)

type question struct {
	key int64
	prompt string
	qtype int
	clist []*criterion
	nchan chan bool
}

func (q *question) New_response (rr *responder)  *response {
	nr := new (response)
	fmt.Println("newresponse")
    q.delold(rr)
	for i,r :=range rr.responses {
		fmt.Print("checking")
		fmt.Println(i)
		if r.q.key == q.key { nr.key = r.q.key }
	}
	if nr.key == 0 {
		Init(nr)
		nr.q=q
		rr.responses=append(rr.responses,nr)
		Sstore(nr)
	}
	nr.q = q
	Sstore(nr)
	Sstore(rr)
	fmt.Print("newresponsekey")
	fmt.Println(nr.key)
	return nr
}
//Should have been better with visibility...
//template getters
func (o *question)Pprompt() string {
	return o.prompt
}

func (q *question)Ptype() int {return q.qtype}
func (q *question)Pclist() []*criterion {return q.clist}
func (q *question)Ptext() string {return q.prompt}
func (q *question) Pvalue(c *criterion) string{
    s := ""

    switch q.qtype {
        case 0: s = c.regex
        case 1: if c.inv {
                    if c.exclusive{
                        s = "Less Than Or Equal To " + strconv.FormatFloat(c.aval, 'f', 0, 64) + " And Greater Than Or Equal To " + strconv.FormatFloat(c.bval, 'f', 0, 64)
                    }else{
                        s = "Less Than " + strconv.FormatFloat(c.aval, 'f', 0, 64) + " And Greater Than " + strconv.FormatFloat(c.bval, 'f', 0, 64)
                    }
                }else{
                    if c.exclusive {
                        s = strconv.FormatFloat(c.aval, 'f', 0, 64) + " to " + strconv.FormatFloat(c.bval, 'f', 0, 64) + "\r\n- Exclusive"
                    }else{
                        s = strconv.FormatFloat(c.aval, 'f', 0, 64) + " to " + strconv.FormatFloat(c.bval, 'f', 0, 64)
                    }
                }
                if c.onlyint{
                        s += "\r\n- No Decimals"
                }
                if c.onlypos{
                    s += "\r\n- Only Positives"
                }
        case 2: if c.lval {
                s = "Yes"
            }else{
                s = "No"
            }
    }

    return s
}
//DB stuff

func Getallquestions() (error, []*question){
	nchan := make(chan error)
	var r []*question
	DBchan <- func(Db *sql.DB)func() {
		rows, err := Db.Query("select key from question")
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			q := new(question)
			rows.Scan(&k)
			//checkErr(err)
			if k == 0 {continue}
			r=append(r,q)
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

func Getallqbyname (p string) (error, []*question){
	fmt.Println("got:"+p)
	if p=="" {return nil,nil}

	nchan := make(chan error)
	var r []*question
	DBchan <- func(Db *sql.DB)func() {

		rows, err := Db.Query("select key from question where prompt like '%?%'",p)
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			s := new(question)
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

func Getallprompts (p string) (error, []*question){
	nchan := make(chan error)
	var r []*question
	DBchan <- func(Db *sql.DB)func() {
		rows, err := Db.Query("select key from question where prompt = ?", p)//TODO regex
		checkErr(err)
		defer rows.Close()
		i :=0
		for rows.Next() {
			var k int64
			err := rows.Scan(&k)
			if err != nil {
				nchan <- err
			}
			r = append(r,new(question))
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

func Getallqbynamefuzz (p string) (error, []*question){
	fmt.Println("got:"+p)
	if p=="" {return nil,nil}
	var r []*question
	search := "%" + p + "%"
	nchan := make(chan error)
	DBchan <- func(Db *sql.DB)func() {
		rows, err := Db.Query("select key from question where prompt like ?", search)
		checkErr(err)
		defer rows.Close()
		for i := 0; rows.Next(); i++ {
			var k int64
			rows.Scan(&k)
			if k == 0 {continue}
			r = append(r, new(question))
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

func Get1q(p string) (*question) {
	err,qs := Getallprompts(p)
	if err != nil || len(qs)<1{return nil} else{
		return qs[0]
	}
}
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

func (o *question) delold(rr *responder) {
    fmt.Print("delold")
    fmt.Println(len(rr.responses))
	for i,r :=range rr.responses {
		if r.q.key == o.key {
			o.Readynchan()
			DBchan <- func (Db *sql.DB) func() {
				stmt,err := Db.Prepare("delete from respondersresponse where ikey = ? and okey = ?")
				checkErr(err)
				_,err = stmt.Exec(rr.key,r.key)
				checkErr(err)
                stmt,err = Db.Prepare("delete from response where key = ?")
				checkErr(err)
				_,err = stmt.Exec(r.key)
				checkErr(err)
				return o.Notify
			}
			o.Wait()
            fmt.Println(rr.responses)
            if i==len(rr.responses) {rr.responses = rr.responses[:i];
			} else {
				rr.responses = append(rr.responses[:i],rr.responses[i+1:]...)
			}
            fmt.Println(rr.responses)
		}
	}

    Sstore(rr)
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
