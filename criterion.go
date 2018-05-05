package main

import (
    "fmt"
    "strconv"
    "database/sql"
    "errors"
    "regexp"
    "strings"
)

type criterion struct {
	key int64
	aval float64
	bval float64
	regex string
	lval bool
	isnl bool
	inv bool//not redundant.
	conj bool
    onlyint bool
    onlypos bool
    exclusive bool
    apresent bool
    bpresent bool
	q *question
	nchan chan bool
}

func (o *criterion) checkstr(v string) bool{
    v = strings.ToLower(v)
	fmt.Println("check str:"+v+" "+o.regex)
	regex,_ := regexp.Compile(strings.ToLower(o.regex))
	//return o.checkbool(v==o.regex)//regex sans the +,*and () operators
	return o.checkbool(regex.MatchString(v))
}
func (o *criterion) checkint(x string) bool{
	fmt.Print("check int:"+x+" ")
	fmt.Print(o.aval)
	fmt.Println(o.bval)
    v,err := strconv.ParseFloat(x,64)
    if o.onlypos && v < 0 {return o.checkbool(false)}
	if o.onlyint && v!=float64(int64(v)) {return o.checkbool(false)}
    if o.exclusive && (v<=o.aval || v>=o.bval) {return o.checkbool(false)}
	if err != nil { return o.checkbool(false)}
	return o.checkbool(((!o.apresent)||v>=o.aval)&&((!o.bpresent)||v<=o.bval))
}

func (o *criterion) checkbool(v bool) bool{//other check methods call this one 
	var x bool
	//x = !(v!=o.lval!=o.inv)
	//XXX
	x = (v!=o.inv)
	return x
}

func Validate(re []*response, cr []*criterion) bool {
	fmt.Println("validate")
	//this checks the user's responses to questions it depends on.
	//the dependancy can be checked by q.something
	vals := make(map[int64] int)//yes, for the conjunctive critria
    conjs :=make(map[int64] bool)
	for _,c := range cr {
		vals[c.key] = -1
	}
	for _,c := range cr {
			fmt.Println("check"+c.regex)
		for _,r :=range re {
			fmt.Println("against "+c.regex+" "+r.value)
			var xj bool
			xj = false
			if c.q == nil || c.q.key == r.q.key {
				switch r.q.qtype {
					case 0: xj = c.checkstr(r.value)
					case 1: xj = c.checkint(r.value)
					case 2: xj,_ = strconv.ParseBool(r.value)
					default: panic("nope")
				}
				if c.conj&&(vals[c.key]!=0) {
                    conjs[c.key] = true
					if xj {
						vals[c.key]= 1
					}else {
						vals[c.key]=0
					}
				} else if !c.conj&&vals[c.key] != 1{
                    conjs[c.key] =false
					if xj {
						vals[c.key]= 1
					} else {
						vals[c.key]=0
					}
				}

			}else {
				fmt.Println("bad question")
			}

		}
	}
	for _,c := range cr {
		if vals[c.key]<1&&conjs[c.key] {
			return false
		}
	}

	hasdis := false
	for _,c := range cr {
        if !conjs[c.key]{hasdis = true}
		if vals[c.key]>0&&!conjs[c.key] {
			return true
		}
	}
	return !hasdis
}

//DB stuff
func (o *criterion) Store(Db *sql.DB) error{
	if o.key == 0 { //init
		stmt, err := Db.Prepare("insert into criterion(key) values(NULL)")

		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		o.key, err = res.LastInsertId()
		checkErr(err)
	} else  { //store
		stmt, err := Db.Prepare("update criterion set(key , aval, bval, regex, lval, isnil, inv, exc, pos, dec, conj, qkey, apresent, bpresent) = (?,?,?,?,?,?,?,?,?,?,?,?,?,?) where key = ?")

		checkErr(err)
		var qk int64
		if(o.q != nil) {
			qk = o.q.key
		} else {
			qk = 0
		}
		res, err := stmt.Exec(o.key, o.aval, o.bval, o.regex, o.lval, o.isnl, o.inv, o.exclusive, o.onlypos, o.onlyint, o.conj, qk, o.apresent, o.bpresent, o.key)
		checkErr(err)
		if res == nil {//XXX
			panic(err)//never happens?
		}
	}
	return nil
}
func (o *criterion) Get(Db *sql.DB) error{
	var qkey int64
	if o.key == 0 {
		//err :=  Db.QueryRow("select key, uname, cursessionid, pwhash, cresp from operator where uname = ?", o.uname).Scan(&o.key,  &o.uname, &o.cursessionid, &phh,&rkey)//TODO
		//if err !=nil {return err}
		return errors.New("not implemented")
	} else {
		err := Db.QueryRow("select key , aval, bval, regex, lval, isnil, inv, exc, pos, dec, conj, qkey, apresent, bpresent from criterion where key = ?", o.key).Scan(&o.key,&o.aval,&o.bval,&o.regex,&o.lval,&o.isnl,&o.inv,&o.exclusive,&o.onlypos,&o.onlyint,&o.conj,&qkey, &o.apresent, &o.bpresent)
		fmt.Print("retrived ")
		fmt.Println(o)
		if err !=nil {o.key = 0; return err}
		if qkey != 0 {
			o.q = new(question)
			o.q.key = qkey
			err = o.q.Get(Db)
			if err != nil {
				o.key = 0
				return err
			}
			return nil
		}
		return nil
	}
	return nil
}

//Template getters
func (o *criterion) Pkey() int64{
	return o.key
}
func (o *criterion) Zkey(){
	o.key=0
}
func (o *criterion) Qkey() int64{
    if o.q == nil {
        return 0
    }else{
        return o.q.key
    }
}
func (o *criterion) Pvalue() string{
    s := ""
    if o.q == nil {
        //do nothing to return empty string
    }else{
        switch o.q.qtype {
            case 0: s = o.regex
            case 1: if o.inv {
                        if !o.apresent && !o.bpresent{
                            s = "Any Value"
                        }else{
                            if !o.apresent && o.bpresent{
                                s = "Greater Than Or Equal To " + strconv.FormatFloat(o.bval, 'f', 0, 64)
                            } else {
                                if o.apresent && !o.bpresent{
                                    s = "Less Than Or Equal To " + strconv.FormatFloat(o.aval, 'f', 0, 64)
                                } else {
                                    if o.exclusive{
                                        //TODO: add checks for apresent and bpresent again
                                        s = "Less Than Or Equal To " +  strconv.FormatFloat(o.aval, 'f', 0, 64) + " And Greater Than Or Equal To " + strconv.FormatFloat(o.bval, 'f', 0, 64)
                                    }else{
                                        s = "Less Than " +  strconv.FormatFloat(o.aval, 'f', 0, 64) + " And Greater Than " + strconv.FormatFloat(o.bval, 'f', 0, 64)
                                    }
                                }
                            }
                        }
                    }else{
                         if !o.apresent && !o.bpresent{
                            s = "Any Value"
                        }else{
                            if !o.apresent && o.bpresent{
                                s = "Less Than Or Equal To " + strconv.FormatFloat(o.bval, 'f', 0, 64)
                            } else {
                                if o.apresent && !o.bpresent{
                                    s = "Greater Than Or Equal To " + strconv.FormatFloat(o.aval, 'f', 0, 64)
                                } else {
                                    if o.exclusive{
                                        s =  strconv.FormatFloat(o.aval, 'f', 0, 64) + " to " + strconv.FormatFloat(o.bval, 'f', 0, 64) + "\r\n- Exclusive"
                                    }else{
                                        s =  strconv.FormatFloat(o.aval, 'f', 0, 64) + " to " + strconv.FormatFloat(o.bval, 'f', 0, 64)
                                    }
                                }
                            }
                        }
                    }
                    if o.onlyint{
                        s += "\r\n- No Decimals"
                    }
                    if o.onlypos{
                        s += "\r\n- Only Positives"
                    }
            case 2: if o.lval {
                        s = "Yes"
                    }else{
                        s = "No"
                    }
        }

        //adds conjunctive attribute
        if o.conj {
            s += "\r\n- Non-Optional"
        }

    }

    return s
}

//DB Sync stuff
func (o *criterion) Wait() {//NOTE: multiple threads cannot use this on the same object
	fmt.Println("waiting...")
	<-o.nchan
}
func (o *criterion) Notify() {
	fmt.Println("note")
	o.nchan <- true
}
func (o *criterion) Readynchan() {
	o.nchan = make(chan bool)
}
