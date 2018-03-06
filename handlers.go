package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "strconv"
    "html/template"
    //"math/rand"
    _ "github.com/mattn/go-sqlite3"
)


//var Db *sql.DB

func mktab (d *sql.DB, name string, cols map[string]string) {
	s := "CREATE TABLE IF NOT EXISTS " + name +" ("
	i := 0
	for k,v := range cols {
		i = i + 1
		if(k == "key") {
			v = "integer primary key asc" //See sqlite documentation.
		}
		s += k + " " + v
		if i != len(cols) {
			s += ","
		}
	}
	s += ")"
//	fmt.Println(s)
	stmt, err := d.Prepare(s)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
//	affect, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println(affect)
	checkErr(err)
}
func Getdb() *sql.DB {
	d, err := sql.Open("sqlite3","./wdb.db")
	checkErr(err)
	return d
}
func initdb () *sql.DB{
	d := Getdb()
	mktab(d,"service",map[string]string{
		"key":"int",
		"name":"string",
		"description":"string",
		"url":"string",
	})
	mktab(d,"operator",map[string]string{
		"key":"int",
		"uname":"string",
		"pwhash":"string",
		"cursessionid":"int",
		"cresp":"int",
	})
	mktab(d,"responder",map[string]string{
		"key":"int",
		"fname":"string",
		"lname":"string",
		"zip":"string",
		"pwhash":"string",
		"dob":"int",//TODO: time
		"iqkey":"int",
	})
	mktab(d,"response",map[string]string{
		"key":"int",
		"value":"string",
		"qkey":"int",
	})
	mktab(d,"question",map[string]string{
		"key":"int",
		"prompt":"string",
		"qtype":"int",
	})
	mktab(d,"criterion",map[string]string{
		"key":"int",
		"aval":"int",
		"bval":"int",
		"lval":"bool",
		"isnil":"bool",
		"inv":"bool",
		"conj":"bool",
		"regex":"string",
		"qtype":"int",
	})
	//relation tables
	mktab(d,"respondersresponse",map[string]string{
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"respondersservice",map[string]string{
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"questionscriterion",map[string]string{
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"servicesquestion",map[string]string{
		"okey":"int",
		"ikey":"int",
	})
	mktab(d,"servicesquestioncriterion",map[string]string{
		"okey":"int",
		"iqkey":"int",
		"ickey":"int",
	})
	return d
}
func webmessage (w http.ResponseWriter,m string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<body>"+m+"</body>.\n"))
}
func outpage(f string , w http.ResponseWriter, d map[string]string){

	t := template.Must(template.ParseFiles(f))
	t.Execute(w,d)
}
func qlist([]question) string {
	t := template.Must(template.Parse(`
	<div id=\"qlist\"> 
	{{range .}} 
	<a href=\"/qprompt/{{q.key}}\">{{q.prompt}}</a><br>\n
	{{end}}
	`)
}
func Authhandler(w http.ResponseWriter, r *http.Request) {
//TODO: rename to sessionhandler
	o := new(operator)
	fmt.Println("")
	fmt.Println("got:"+r.FormValue("uname")+" "+r.FormValue("pw"))
	if o.Getbyname(r.FormValue("uname")) == nil && o.key != 0 {
		if o.Auth(r.FormValue("pw")) {
		//we set two session cookies: one has the sessionid and the
		//other has the username
			uc := http.Cookie{
				Name: "uname",
			}
			uc.Value = o.uname
			sc := http.Cookie{
				Name: "sessionid",
			}
			sc.Value= strconv.Itoa(o.cursessionid)
			http.SetCookie(w, &uc)
			http.SetCookie(w, &sc)

			//w.Header().Set("Content-Type", "text/html")
			//w.Write([]byte("<body>Auth successfull!<br><a href=\"/addresponder.html\">add a responder</a></body>\n"))
            outpage("addresponder.html",w,map[string]string{"err": "",})
			Sstore(o)
		}else {
			outpage("auth.html.tpl",w,map[string]string{"err":"Bad Secret",})
		}
	} else {
		fmt.Println("key:"+strconv.FormatInt(o.key,10))
		//webmessage(w,"No identity to auth!\n")
		if r.FormValue("uname") != ""|| r.FormValue("pw") != "" {
			outpage("auth.html.tpl",w,map[string]string{"err":"Please enter a username",})
		} else {
			outpage("auth.html.tpl",w,map[string]string{"err":""})
		}
	}

}
func curop(r *http.Request) *operator {
	uc, err := r.Cookie("uname")
	fmt.Println("logged in as"+uc.Value)
	if err != nil {
		return nil
	}
	o := new (operator)
	o.Getbyname(uc.Value)
	if o.key == 0 {
		fmt.Println("nil op")
		return nil
	}
	sc, err := r.Cookie("sessionid")
	if err != nil {
		return nil
	}
	s,_:=strconv.Atoi(sc.Value)
	if o.Checksesh(s) {
		fmt.Println("session good")
		return o
	}
	fmt.Println("session bad")
	return nil

}
func Ursession_handler(w http.ResponseWriter, r *http.Request) {
	o := curop(r)
	if o == nil {
		webmessage(w,"Bad Session")
	} else {
		if r.FormValue("rkey") =="" {
			fmt.Println("got:"+r.FormValue("fname")+" "+r.FormValue("lname")+" "+r.FormValue("dob")+" "+r.FormValue("zip"))
			dob,_:=strconv.Atoi(r.FormValue("dob"))
			err,rs :=Getallmatch(r.FormValue("fname"),r.FormValue("lname"),dob,r.FormValue("zip"))
			checkErr(err)

			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<body>Select a responder\n"))
			for _,r := range rs {
				w.Write([]byte(r.Tohtml()+"<br>\n"))
			}
			s:="<br> Or create a fresh one:<form method=\"post\">"
			s+="<input type=\"hidden\" name=\"fname\" value=\"" +r.FormValue("fname")+"\">"
			s+="<input type=\"hidden\" name=\"lname\" value=\"" +r.FormValue("lname")+"\">"
			s+="<input type=\"hidden\" name=\"dob\" value=\"" +r.FormValue("dob")+"\">"
			s+="<input type=\"hidden\" name=\"zip\" value=\"" +r.FormValue("zip")+"\">"
			s+="<input type=\"hidden\" name=\"rkey\" value=\"0\">"
			s+="<input type=submit></form>"
			w.Write([]byte(s))
		} else {
			if r.FormValue("rkey") == "0" {
				o.cresp = new(responder)
				Init(o.cresp)
				o.cresp.fname=r.FormValue("fname")
				o.cresp.lname=r.FormValue("lname")
				o.cresp.dob,_=strconv.Atoi(r.FormValue("dob"))
				o.cresp.zip=r.FormValue("zip")
				Sstore(o)
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte("<body>New responder created!</body>\n"))
			}else {
				o.cresp.key,_ = strconv.ParseInt(r.FormValue("rkey"),10,64)
				Sget(o.cresp)
				Sstore(o)//TODO:check errors
			}
		}
	}
}

func qprompt_handler(w http.ResponseWriter, r *http.Request) {
	o := curop(r)
	if o == nil {
		webmessage(w,"Bad Session")
	} else {
		k,err := strconv.ParseInt(r.URL[8:])
		if k==0 || err != nil {
			//no question, list them.
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte("<html><body>\n"
					+"<h1>Questions</h1> :D D: :D\n"
					+qlist(o.cser.qlist)
					+"or whatever</body></html>"))

		}
	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("The application has encounterd an unrecoverable error")
		panic(err)
	}
}

