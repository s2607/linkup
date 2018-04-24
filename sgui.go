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
        msg := ""
        title := "Add"
        editing := false //used to know which popup text to use
        if o == nil {
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
        } else if r.FormValue("uname") != "" {
		if r.FormValue("nokey") != "0" {
			no.key,_ = strconv.ParseInt(r.FormValue("nokey"),10,64)
			Sget(no)
            msg = "Information Successfully Changed"
		}else { Init(no)
               msg = "Interviewer Successfully Added"
        }
        title = "Edit"
        editing = true
		no.uname=r.FormValue("uname")
		no.cser = new(service)
		no.cser.key,_ = strconv.ParseInt(r.FormValue("skey"),10,64)
		Sget(no.cser)
		no.setpss(r.FormValue("pw"))
		Sstore(no)
        //This if statement updates the cookie uname if the interviewer edited themself to keep session good and update name on home page
        if o.key == no.key {
            uc,_ := r.Cookie("uname")
            uc.Value = no.uname
            http.SetCookie(w, uc)
        }
        t := template.Must(template.ParseFiles("addop.html.tpl"))
        t.Execute(w,struct{E bool; I *operator; T string; Succ string; Anim string}{editing, no, title, msg, "none"})
        }else {
		if r.FormValue("nokey") != "" {
			no.key,_ = strconv.ParseInt(r.FormValue("nokey"),10,64)
			Sget(no)
            title = "Edit"
            editing = true
		}
		t := template.Must(template.ParseFiles("addop.html.tpl"))
            t.Execute(w,struct{E bool; I *operator; T string; Succ string; Anim string}{editing, no, title, msg, "fade_in"})
	}
}
func servicecreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
	ns := new(service)
    qid := r.FormValue("nqkey") //gets qid to put in nprompt when add is clicked on searchqid
    title := "Add" //Sets default title to add
    fmt.Println(r.FormValue("name"))
        if o == nil {
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
        }else {
		if r.FormValue("nskey") != "" && r.FormValue("nskey") != "0"{
			ns.key,_ = strconv.ParseInt(r.FormValue("nskey"),10,64)
			Sget(ns)
            title = "Edit"
			fmt.Println("editing older service"+ns.name)
		}else { Init(ns)}
		if r.FormValue("name") != "" {
			ns.name=r.FormValue("name")
			ns.url=r.FormValue("url")
			ns.description=r.FormValue("description")
		}
        fmt.Println(ns)
        if r.FormValue("qid") !="" {
			c:= new (criterion)
			createc(c,r)
			ns.criteria = append(ns.criteria,c)
		}
		if r.FormValue("nprompt") != "" {
			q :=new(question)
			q.key,_=strconv.ParseInt(r.FormValue("nprompt"),10,64)
			Sget(q)
			//q :=Get1q(r.FormValue("nprompt"))
			if q != nil {
				ns.qlist= append(ns.qlist,*q)
			}
		}
		if ns.key != 0 {
            Sstore(ns)
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
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
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
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
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
        outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
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
        outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
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
        outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
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
        outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
    }
}

//SQL Command Interface
func sql_handler(w http.ResponseWriter, r *http.Request) {
    if curop(r) != nil {
        msg := ""
        anim := ""
        color := ""
        fmt.Println("sql" + r.FormValue("q"))
        s := Sql_injector(r.FormValue("q"), &msg)
        fmt.Println(s)
        t := template.Must(template.ParseFiles("sql.html.tpl"))

        //used for ui messages and animations
        if msg != "" {
            anim = "animation: none"
            //if the message is an error message
            if msg != "Submitted" {
                color = "color: red"
            }
        }

        data := struct{
            M string
            A string
            C string
        }{
            msg,
            anim,
            color,
        }

        t.Execute(w, data)
    }else{
        outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
    }
}

