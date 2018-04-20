package main
import (
	"net/http"
	"strconv"
	"html/template"
	"database/sql"
	"fmt"
)
//TODO: edit
func opcreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
	    no := new(operator)
        if o == nil {
                webmessage(w,"Bad Session")
        } else if r.FormValue("uname") != "" {
		if r.FormValue("nokey") != "" {
			no.key,_ = strconv.ParseInt(r.FormValue("nokey"),10,64)
			Sget(no)
		}else { Init(no) }
		no.uname=r.FormValue("uname")
		no.cser = new(service)
		no.cser.key,_ = strconv.ParseInt(r.FormValue("skey"),10,64)
		Sget(no.cser)
		no.setpss(r.FormValue("pw"))
		Sstore(no)
        t := template.Must(template.ParseFiles("addop.html.tpl"))
            t.Execute(w,struct{Succ string; Anim string}{"Interviewer Successfully Added", "none"})
        }else {
		if r.FormValue("nokey") != "" {
			no.key,_ = strconv.ParseInt(r.FormValue("nokey"),10,64)
			Sget(no)
		}
		t := template.Must(template.ParseFiles("addop.html.tpl"))
            t.Execute(w,struct{Succ string; Anim string}{"", "fade_in"})
	}
}
func servicecreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
	ns := new(service)
    qid := r.FormValue("nqkey") //gets qid to put in nprompt when add is clicked on searchqid
    title := "Add" //Sets default title to add
    fmt.Println("Service Key is: " + r.FormValue("nskey"))
        if o == nil {
                webmessage(w,"Bad Session")
        } else if r.FormValue("name") != "" {
		if r.FormValue("nskey") != "" {
			ns.key,_ = strconv.ParseInt(r.FormValue("nskey"),10,64)
			Sget(ns)
		}else { Init(ns)}
		ns.name=r.FormValue("name")
		ns.url=r.FormValue("url")
		ns.description=r.FormValue("description")
		if r.FormValue("inv") !="" {
			c:= new (criterion)
			createc(c,r)
			ns.criteria = append(ns.criteria,c)
		}
		if r.FormValue("nprompt") != "" {
			q :=Get1q(r.FormValue("nprompt"))
			if q != nil {
				ns.qlist= append(ns.qlist,*q)
			}
		}
		Sstore(ns)
        webmessage(w,"ok")
        }else {
		if r.FormValue("nskey") != "" {
			ns.key,_ = strconv.ParseInt(r.FormValue("nskey"),10,64)
			Sget(ns)
            title = "Edit"
		}
		t := template.Must(template.ParseFiles("addserv.html.tpl"))
        t.Execute(w,struct{T string; Qid string; A string; O *service}{title, qid,"/newserv",ns})
	}
}

func cquestion(nq *question, r *http.Request) {
		nq.prompt=r.FormValue("prompt")
		nq.qtype,_=strconv.Atoi(r.FormValue("qtype"))
}
func questioncreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
	nq := new(question)
        if o == nil {
                webmessage(w,"Bad Session")
        } else if r.FormValue("prompt") != "" {
		if r.FormValue("nqkey") != "" {
			nq.key,_ = strconv.ParseInt(r.FormValue("nqkey"),10,64)
			Sget(nq)
		}else { Init(nq)}

		cquestion(nq,r)
		if r.FormValue("inv") !="" {
			c:= new (criterion)
			createc(c,r)
			nq.clist = append(nq.clist,c) //NOTE: questions being updated will not create multiple entries
		}
		Sstore(nq)
		webmessage(w,"ok")
        }else {
		if r.FormValue("nqkey") != "" {
			nq.key,_ = strconv.ParseInt(r.FormValue("nqkey"),10,64)
			Sget(nq)
		}
		t := template.Must(template.ParseFiles("addq.html.tpl"))
		t.Execute(w,struct{A string; O *question}{"/newq",nq})
	}
}
func ist(s string) bool{
	return s=="true"
}
func createc(nc *criterion,r *http.Request)error{
		nc.regex=r.FormValue("regex")//The one string
		nc.aval,_=strconv.Atoi(r.FormValue("aval"))
		nc.bval,_=strconv.Atoi(r.FormValue("bval"))
		nc.lval=ist(r.FormValue("lval"))
		nc.isnl=ist(r.FormValue("isnil"))
		nc.inv=ist(r.FormValue("inv"))
		nc.conj=ist(r.FormValue("conj"))
		return nil
}
func criterioncreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
	nc := new(criterion)
        if o == nil {
                webmessage(w,"Bad Session")
        } else if r.FormValue("uname") != "" {
		if r.FormValue("nckey") != "" {
			nc.key,_ = strconv.ParseInt(r.FormValue("nckey"),10,64)
			Sget(nc)
		}else { Init(nc)}
		createc(nc,r)
		Sstore(nc)
		webmessage(w,"ok")
        }else {
		if r.FormValue("nckey") != "" {
			nc.key,_ = strconv.ParseInt(r.FormValue("nckey"),10,64)
			Sget(nc)
		}
		t := template.Must(template.ParseFiles("addc.html.tpl"))
		t.Execute(w,nc)
	}
}
func delc_handler(w http.ResponseWriter, r *http.Request) {
	nchan := make (chan error)
	DBchan <- func(DB *sql.DB) func() {
		stmt,err := DB.Prepare("delete from ?scriterion where ikey = ?")
		if err != nil { nchan <- err }
		_,err =stmt.Exec(r.FormValue("composingtype"),r.FormValue("ckey"))
		if err != nil { nchan <- err }
		return func() {
			nchan <- err
		}
	}
	checkErr(<-nchan)
}
func delq_handler(w http.ResponseWriter, r *http.Request) {
	nchan := make (chan error)
	DBchan <- func(DB *sql.DB) func() {
		stmt,err := DB.Prepare("delete from servicesquestion where ikey = ? and okey = ?")
		if err != nil { nchan <- err }
		_,err =stmt.Exec(r.FormValue("qkey"))
		if err != nil { nchan <- err }
		return func() {
			nchan <- err
		}
	}
	checkErr(<-nchan)
}
//search

func searchq_handler(w http.ResponseWriter, r *http.Request) {

    if curop(r) != nil {
        fmt.Println("searchq"+r.FormValue("q"))
		err,s := Getallqbyname(r.FormValue("q"))
		fmt.Println(s)
		checkErr(err)
		t := template.Must(template.ParseFiles("searchq.html.tpl"))
		t.Execute(w,s)
    }else{
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte("<body>Bad Session</body>.\n"))
    }
}
func searchs_handler(w http.ResponseWriter, r *http.Request) {
    if curop(r) != nil {
        fmt.Println("searchs"+r.FormValue("q"))
		err,s := Getallsbyname(r.FormValue("q"))
		fmt.Println(s)
		checkErr(err)
		t := template.Must(template.ParseFiles("searchs.html.tpl"))
		t.Execute(w,s)
    }else{
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte("<body>Bad Session</body>.\n"))
    }

}
func searcho_handler(w http.ResponseWriter, r *http.Request) {
    if curop(r) != nil {
        fmt.Println("searcho"+r.FormValue("q"))
		err,s := Getallobyname(r.FormValue("q"))
		fmt.Println(s)
		checkErr(err)
		t := template.Must(template.ParseFiles("searcho.html.tpl"))
		t.Execute(w,s)
    }else{
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte("<body>Bad Session</body>.\n"))
    }
}

func searchqid_handler(w http.ResponseWriter, r *http.Request) {
    if curop(r) != nil {
        fmt.Println("searchq"+r.FormValue("q"))
        fmt.Println("Service Key is: " + r.FormValue("skey"))
		err,q := Getallqbyname(r.FormValue("q"))
		fmt.Println(q)
		checkErr(err)
        data := struct{
            Q []*question
            K string
        }{
            q,
            r.FormValue("skey"),
        }
		t := template.Must(template.ParseFiles("searchqid.html.tpl"))
		t.Execute(w,data)
    }else{
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte("<body>Bad Session</body>.\n"))
    }
}

//SQL Command Interface
func sql_handler(w http.ResponseWriter, r *http.Request) {
    if curop(r) != nil {
        fmt.Println("sql" + r.FormValue("q"))
        s := Sql_injector(r.FormValue("q"))
        fmt.Println(s)
        t := template.Must(template.ParseFiles("sql.html.tpl"))
        t.Execute(w, s)
    }else{
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte("<body>Bad Session</body>.\n"))
    }
}

