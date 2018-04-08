package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "strconv"
    "html/template"
    "unicode"
    "time"
    "strings"
    //"math/rand"
    _ "github.com/mattn/go-sqlite3"
)


//var Db *sql.DB

func mktab (d *sql.DB, name string, cols map[string]string) {
	s := "CREATE TABLE IF NOT EXISTS " + name +" ("
	i := 0
	var ikey string
	var okey string
	for k,v := range cols {
		i = i + 1
		if(k == "key") {
			v = "integer primary key asc" //See sqlite documentation.
		}
		if(k == "ikey") {
			ikey = k
		}
		if(k == "okey") {
			okey = k
		}
		s += k + " " + v
		if i != len(cols) {
			s += ","
		}
	}
	if len(okey)>1 && len(ikey) >1 {
		s += ", PRIMARY KEY ( ikey, okey )"
	}
	s += ")"
	fmt.Println(s)
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
		"cser":"int",
	})
	mktab(d,"responder",map[string]string{
		"key":"int",
		"fname":"string",
		"lname":"string",
		"zip":"string",
		"pwhash":"string",
		"dob":"int",//TODO: time
		//"iqkey":"int",//???
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
        "qkey":"int",
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
    mktab(d,"servicescriterion",map[string]string{
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
func qlist(w http.ResponseWriter, q []question){

    t := template.Must(template.ParseFiles("disp_qlist.html.tpl"))
    err := t.Execute(w,q)
    checkErr(err)

    /*t,err := template.New("dispt").Parse(`
	<div id="qlist">
	{{range .}} 
	<a href="/qprompt/{{.Pkey}}">{{.Pprompt}}</a><br>
	{{end}}
	</div>
	`)
	checkErr(err)
	err = t.Execute(w,q)
	checkErr(err)*/
}
func qdisp(w http.ResponseWriter, k int64) {
	q := new(question)
	q.key = k
	err := Sget(q)//TODO: check errors
	checkErr(err)

    t := template.Must(template.ParseFiles("disp_question.html.tpl"))

	/*t,err := template.New("dispt").Parse(`
	<div id="question">
	<form method="post">
	{{.Pprompt}}
	<input name="qanswer" value="" >
	<input type=submit></form>
	</div>
	`)//TODO: different form types*/
	err = t.Execute(w,q)
    checkErr(err)

}
func qanswer(k int64, s string, ur *responder) error {//TODO: error checking this whole function
	q := new(question)
	q.key = k
	err := Sget(q)//TODO: check errors
	checkErr(err)
	r:=q.New_response()
	r.value=s//TODO:validate
	ur.responses=append(ur.responses,r)
	return nil
}
func showsug(w http.ResponseWriter, r responder){
	/*t,err := template.New("dispt").Parse(`
	<div id="slist">
	{{range .}}
	<div id="suggestion"> <a href="{{.Purl}}">{{.Pname}}</a><p>{{.Pdesc}}</p> </div><br>
	{{end}}
	</div>
	<a href="/qprompt">return</a>
	`)
	checkErr(err)*/

    t := template.Must(template.ParseFiles("disp_sugs.html.tpl"))

    err := t.Execute(w,r.suggestions)
    checkErr(err)

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


            w.Header().Set("Content-Type", "text/html")
            w.Write([]byte(`<html><head>
            <title>LinkUp</title>
            <link rel="icon" href="imgs/chevron.png" type="image/x-icon">
            <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
            <meta charset="UTF-8">
            <meta name="description" content="">
            <meta name="keywords" content="">

            <!-- Fonts -->
            <link href="https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans" rel="stylesheet">

            <!-- Stylesheet -->
            <link href="css/survey_stylesheet.css" rel="stylesheet">

            </head>
            <body><div id="container"><div id="top_bar">
            <img id="logo" src="imgs/logo.svg" alt="LinkUp">
            </div>

            <div id="title">
                <h1>Select An Action</h1>
            </div>
            <div style="text-align: center; margin-top: 50px; animation: form_fade_in; animation-duration: 1s; animation-timing-function: ease-in-out;">
				<a href="/addresponder.html">Add A Responder</a><br /><br />
				<a href="/newop">Add An Operator</a><br /><br />
				<a href="/newserv">Add A Service Program</a><br /><br />
				<a href="/newq">Add A Question</a><br /><br />
				<a href="/searchq">Search For A Question</a><br /><br />
				<a href="/searchop">Search For An Operator</a><br /><br />
				<a href="/searchs">Search For A Service Program</a><br />
            </div>
            </div>
			</body>`))
			Sstore(o)
		}else {
			outpage("auth.html.tpl",w,map[string]string{"err":"Invalid Password",})
		}
	} else {
		fmt.Println("key:"+strconv.FormatInt(o.key,10))
		//webmessage(w,"No identity to auth!\n")
		if r.FormValue("uname") != ""|| r.FormValue("pw") != "" {
			outpage("auth.html.tpl",w,map[string]string{"err":"Please Enter A Valid Username",})
		} else {
			outpage("auth.html.tpl",w,map[string]string{"err":"Please Enter Credentials"})
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
        return
	}



    if r.FormValue("rkey") =="" {
        fmt.Println("got:"+r.FormValue("fname")+" "+r.FormValue("lname")+" "+r.FormValue("dob")+" "+r.FormValue("zip"))

        if(!validateResponderId(w,r)){
            return
        }

        dob,_:=strconv.Atoi(r.FormValue("dob"))
		err,rs :=Getallmatch(r.FormValue("fname"),r.FormValue("lname"),dob,r.FormValue("zip"))
		checkErr(err)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<head><title>LinkUp</title> <link rel='icon' href='imgs/chevron.png' type='image/x-icon'> <link href='https://fonts.googleapis.com/css?family=Roboto+Condensed|Nunito+Sans' rel='stylesheet'> <link href='css/survey_stylesheet.css' rel='stylesheet'></head><body><div id='top_bar'>         <img id='logo' src='imgs/logo.svg' alt='LinkUp'></div><div id='title'><h1>Select A Responder</h1></div>"))

        if rs == nil{
            w.Write([]byte("<div class='responder_entry'>No Matching Entries </div><br>"))
        }else{
		  for _,r := range rs {
		      w.Write([]byte(r.Tohtml()))
		  }
        }
		s:="<br><br><form id='form' method=\"post\">"
		s+="<input type=\"hidden\" name=\"fname\" value=\"" +r.FormValue("fname")+"\">"
		s+="<input type=\"hidden\" name=\"lname\" value=\"" +r.FormValue("lname")+"\">"
		s+="<input type=\"hidden\" name=\"dob\" value=\"" +r.FormValue("dob")+"\">"
		s+="<input type=\"hidden\" name=\"zip\" value=\"" +r.FormValue("zip")+"\">"
		s+="<input type=\"hidden\" name=\"rkey\" value=\"0\">"
        s+="<input type=submit id='submit_button' value='Create New Entry' style='width: 140px;'></form>"
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
			w.Write([]byte("<body>New responder created!<a href=\"/qprompt\">Answer questions</a></body>\n"))
        }else {
            o.cresp.key,_ = strconv.ParseInt(r.FormValue("rkey"),10,64)
			Sget(o.cresp)
			Sstore(o)//TODO:check errors
		}
    }
}

func qprompt_handler(w http.ResponseWriter, r *http.Request) {
	o := curop(r)
	if o == nil {
		webmessage(w,"Bad Session")
	} else {
		var k int64
		var err error
		fmt.Println("qprompt: got url:"+r.URL.Path)
		if len(r.URL.Path) >9 {
		k,err = strconv.ParseInt(r.URL.Path[9:],10,64)
		} else { k = 0 }
		if k==0 || err != nil {
			//no question, list them.
		/*		w.Header().Set("Content-Type", "text/html")
				w.Write([]byte("<html><body>\n"
					+"<h1>Questions</h1> :D D: :D\n"
					+qlist(o.cser.qlist)
					+"or whatever</body></html>"))
		*/
		fmt.Println(o)
		fmt.Println(o.cser)
		fmt.Println(o.cser.qlist)
			qlist(w,o.cser.qlist)//Note that curop calls Sget

		}else {
			if r.FormValue("qanswer") == "" {
				qdisp(w,k)
			}else {
				if qanswer(k,r.FormValue("qanswer"),o.cresp) != nil {
					w.Header().Set("Content-Type", "text/html")
					w.Write([]byte("<body>Failed</body>\n"))
				}else {
					err := Sstore(o)
					checkErr(err)
					w.Header().Set("Content-Type", "text/html")
					w.Write([]byte("<body>Response stored!</body>\n"))
				}
			}
		}
	}
}

func sugg_handler(w http.ResponseWriter, r *http.Request) {
	o := curop(r)
	if o == nil {
		webmessage(w,"Bad Session")
	} else {
		o.cresp.update_suggestions()
		showsug(w,*(o.cresp))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("The application has encounterd an unrecoverable error")
		panic(err)
	}
}

func validateResponderId(w http.ResponseWriter, r *http.Request) bool {
    //Validate Input
    if r.FormValue("fname") == "" || r.FormValue("lname") == "" ||
        r.FormValue("dob") == "" || r.FormValue("zip") == "" {
         outpage("addresponder.html.tpl",w,map[string]string{"err":"Please Complete All Fields", "fname":r.FormValue("fname"), "lname":r.FormValue("lname"), "dob":r.FormValue("dob"), "zip":r.FormValue("zip")})
        return false
    }

    if !checkTextInput(r.FormValue("fname")){
        outpage("addresponder.html.tpl",w,map[string]string{"err":"Invalid First Name", "fname":r.FormValue("fname"), "lname":r.FormValue("lname"), "dob":r.FormValue("dob"), "zip":r.FormValue("zip")})
        return false
    }

    if !checkTextInput(r.FormValue("lname")){
        outpage("addresponder.html.tpl",w,map[string]string{"err":"Invalid Last Name", "fname":r.FormValue("fname"), "lname":r.FormValue("lname"), "dob":r.FormValue("dob"), "zip":r.FormValue("zip")})
        return false
    }

    //Add this after type for DOB is correct
    if !checkDOBInput(r.FormValue("dob")){
        outpage("addresponder.html.tpl",w,map[string]string{"err":"Invalid Date Of Birth", "fname":r.FormValue("fname"), "lname":r.FormValue("lname"), "dob":r.FormValue("dob"), "zip":r.FormValue("zip")})
        return false
    }

    if !checkNumberInput(r.FormValue("zip")) || len(r.FormValue("zip")) != 5 {
        outpage("addresponder.html.tpl",w,map[string]string{"err":"Invalid ZIP Code", "fname":r.FormValue("fname"), "lname":r.FormValue("lname"), "dob":r.FormValue("dob"), "zip":r.FormValue("zip")})
        return false
    }

    return true
    //End validation
}

func checkTextInput(s string) bool{
    for _, r := range s{
        if !unicode.IsLetter(r){
            return false
        }
    }

    return true
}

func checkDOBInput(s string) bool{

    var slashCount int = 0

    for _, r:= range s {

        if !unicode.IsDigit(r) && r != '/' {
            return false
        }

        if r == '/'{
            slashCount++
        }

    }

    //format requires exactly 2 slashes
    if slashCount != 2 {
        return false
    }


    //Split string to month, day, and year strings
    date := strings.Split(s, "/")

    //make sure each has a value
    for _, d := range date {
        if d == "" {
            return false
        }
    }

    month,_ := strconv.Atoi(date[0])
    day,_ := strconv.Atoi(date[1])
    year,_ := strconv.Atoi(date[2])



    //Make sure month is 1-12
    if month < 1 || month > 12 {
        return false
    }

    // Make sure day is 1-31 for months 1,3,5,7,8,10,12
    if month == 1 || month == 3 || month == 5 || month == 7 ||
    month == 8 || month == 10 || month == 12{
        if day < 1 || day > 31 {
            return false
        }
    }

    // Make sure day is 1-30 for months 4,6,9,11
    if month == 4 || month == 6 || month == 9 || month == 11{
        if day < 1 || day > 30 {
            return false
        }
    }

    // Make sure day is between 1-28 for month 2
    if month == 2 {
        //If year is a leap year then range is 1-29
        if year % 4 == 0 {
            if day < 1 || day > 29 {
                return false
            }
        }else{
            if day < 1 || day > 28 {
                return false
            }
        }
    }


    //Make sure year is less than current year
    if year < 1850 || year > time.Now().Year() {
        return false
    }

    // Make sure not past the current date in current year
    if year == time.Now().Year(){
        if month > int(time.Now().Month()){
            return false
        }

        if month == int(time.Now().Month()){
            if day > time.Now().Day(){
                return false
            }
        }
    }

    //If date is valid
    return true
}

func checkNumberInput(s string) bool{
    for _, r := range s{
        if !unicode.IsDigit(r){
            return false
        }
    }

    return true
}

