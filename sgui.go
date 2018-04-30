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
        if !o.admin{return}
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
        no.setAdmin(r.FormValue("admin"))
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
    if !o.admin{return}
	ns := new(service)
    q := new(question)
    qid := r.FormValue("nqkey") //gets qid to put in nprompt when add is clicked on searchqid
    title := "Add" //Sets default title to add
    msg := "" //Success message
    anim := "" //to stop animation
    alreadyCriterion := false //prevents two criterion to one question in service
    alreadyQuestion := false //prevents two of same question being added

    //bools to show/hide forms
    editing := false
    associateQuestion := false
    nonemptyCriterion := false
    addCriterion := false
    nonemptyQList := false
    numQ := false
    boolQ := false
    errMsg := false

    fmt.Println(r.FormValue("name"))
        if o == nil {
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
        }else {
		if r.FormValue("nskey") != "" && r.FormValue("nskey") != "0"{
			ns.key,_ = strconv.ParseInt(r.FormValue("nskey"),10,64)
			Sget(ns)
            title = "Edit"
            editing = true //allows it to show question/service part of form
			fmt.Println("editing older service"+ns.name)
		}else { Init(ns)}

        //if coming back from question set the message and put id in associate question box
        if r.FormValue("nqkey") != "" {
            qid = r.FormValue("nqkey")
            associateQuestion = true
        }

        //takes to page to create new question if that button was clicked
        if r.FormValue("createq") == "true" {
            questioncreate_handler(w,r)
            return
        }

        //shows form to associate questions if that button was clicked
        if r.FormValue("assoc") == "true" {
            associateQuestion = true
            anim = "animation: none"
        }

        //creates/updates a service of the top form was clicked
		if r.FormValue("name") != "" {
			ns.name=r.FormValue("name")
			ns.url=r.FormValue("url")
			ns.description=r.FormValue("description")
            if title == "Edit" {
                msg = "Service Updated"
            }else{
                msg = "Service Created"
            }
            anim = "animation: none"
		}
        fmt.Println(ns)

        //Shows add a criterion form if not already a criterion for the question
        if r.FormValue("questionid") != "" {

            x,_ := strconv.ParseInt(r.FormValue("questionid"),10,64)
            for _, cq := range ns.criteria{
                if cq.q.key == x{
                    alreadyCriterion = true
                }
            }

            if !alreadyCriterion{
                addCriterion = true

                //get question and output the specific form for its type
			     q.key,_=strconv.ParseInt(r.FormValue("questionid"),10,64)
			     Sget(q)
                switch q.qtype {
                    case 0: //Do nothing as form will default to text
                    case 1: numQ = true
                    case 2: boolQ = true
                }
            }else{
                msg = "Question Already Has A Criterion"
                errMsg = true
            }

            anim = "animation: none"

        }

        //creates criteria
        if r.FormValue("qid") !="" {
			c:= new (criterion)
			createc(c,r)
            Sstore(c)
            ns.criteria = append(ns.criteria,c)
            msg = "Criterion Created"
            anim = "animation: none"
		}

        //if remove question button was clicked
        if r.FormValue("ikey") != "" {
            msg = "Question Removed"
            anim = "animation: none"
        }

        //if remove criterion button was clicked
        if r.FormValue("nckey") != "" {
            msg = "Criterion Removed"
            anim = "animation: none"
        }

        //if there ARE criteria then show the remove criteria form
        if ns.criteria != nil {
            nonemptyCriterion = true
        }

        //if associating a question to the service, add it to the services qlist
		if r.FormValue("nprompt") != "" {
			q.key,_=strconv.ParseInt(r.FormValue("nprompt"),10,64)

            for _, sq := range ns.qlist{
                if sq.key == q.key{
                    alreadyQuestion = true
                }
            }

            if !alreadyQuestion{
                Sget(q)
			     //q :=Get1q(r.FormValue("nprompt"))
			     if q != nil && q.key != 0 {
				    ns.qlist= append(ns.qlist,*q)
			     }
                msg = "Question Associated"
            }else{
                msg = "Question Already Associated"
                errMsg = true
            }

            anim = "animation: none"
		}

        //if services qlist is not empty, show the remove question form
        if ns.qlist != nil {
            nonemptyQList = true
        }

		if ns.key != 0 {
            Sstore(ns)
        }
		t := template.Must(template.ParseFiles("addserv.html.tpl"))

        data := struct{
            M string
            T string
            Qid string
            A string
            O *service
            E bool
            Assoc bool
            Nec bool
            Q *question
            C bool
            QList bool
            NumQ bool
            BoolQ bool
            Err bool
        }{
            msg,
            title,
            qid,
            anim,
            ns,
            editing,
            associateQuestion,
            nonemptyCriterion,
            q,
            addCriterion,
            nonemptyQList,
            numQ,
            boolQ,
            errMsg,
        }

		t.Execute(w,data)
        }
}

func cquestion(nq *question, r *http.Request) {
		nq.prompt=r.FormValue("prompt")
		nq.qtype,_=strconv.Atoi(r.FormValue("qtype"))
}
func questioncreate_handler(w http.ResponseWriter, r *http.Request) {
        o := curop(r)
    if !o.admin{return}
	   nq := new(question)
        msg := ""
        anim := ""
        title := "Add"
        backToServe := false
        servKey := r.FormValue("nskey") //used to send back to editing service if came from there

        //bools to show/hide forms
        editing := false
        numQ := false
        boolQ := false
        nonEmptyCriterion := false

        if o == nil {
                outpage("auth.html.tpl",w,map[string]string{"err":"Bad Session"})
        } else {

		if r.FormValue("nqkey") != "" {
			nq.key,_ = strconv.ParseInt(r.FormValue("nqkey"),10,64)
			Sget(nq)
            title = "Edit"
            editing = true

            switch nq.qtype {
                case 0: //Do nothing as form will default to text
		        case 1: numQ = true
                case 2: boolQ = true
            }

		}else {
            Init(nq)
        }

        if r.FormValue("prompt") != "" {
            cquestion(nq,r)

            if r.FormValue("editing") == "Edit" {
                msg = "Question Updated"
            }else{
                msg = "Question Created With ID: " + strconv.FormatInt(nq.key, 10)
            }

            anim = "animation: none"
            title = "Edit"

            switch nq.qtype {
                case 0: //Do nothing as form will default to text
		        case 1: numQ = true
                case 2: boolQ = true
            }

            //If they came directly from selecting a service, go back to it. Note: this must only appear after the question is created or they will return to the service with a question id for a nil question
            if r.FormValue("nskey") != ""{
                backToServe = true
            }
        }

		if r.FormValue("qkey") !="" {
			c:= new (criterion)
			createc(c,r)
			nq.clist = append(nq.clist,c) //NOTE: questions being updated will not create multiple entries
            msg = "Criterion Created"
            anim = "animation: none"
            title = "Edit"
            if r.FormValue("nskey") != ""{
                backToServe = true
            }
		}

        if nq.clist != nil {
            nonEmptyCriterion = true
        }

        if r.FormValue("nckey") != "" {
            msg = "Criterion Removed"
            anim = "animation: none"
            if r.FormValue("nskey") != ""{
                backToServe = true
            }
        }

        if nq.key != 0 {
            Sstore(nq)
        }

        t := template.Must(template.ParseFiles("addq.html.tpl"))

        data := struct{
            A string
            M string
            O *question
            T string
            S string
            E bool
            Back bool
            N bool
            B bool
            CList bool
        }{
            anim,
            msg,
            nq,
            title,
            servKey,
            editing,
            backToServe,
            numQ,
            boolQ,
            nonEmptyCriterion,
        }

	   t.Execute(w,data)
    }
}
func ist(s string) bool{
	return s=="true"
}
func createc(nc *criterion,r *http.Request)error{
        Init(nc)
		nc.regex=r.FormValue("regex")//The one string
		nc.aval,_=strconv.ParseFloat(r.FormValue("aval"),64)
		nc.bval,_=strconv.ParseFloat(r.FormValue("bval"),64)
		nc.lval=ist(r.FormValue("lval"))
		nc.isnl=ist(r.FormValue("isnil"))
		nc.inv=ist(r.FormValue("inv"))
		nc.conj=ist(r.FormValue("conj"))
        nc.onlyint=ist(r.FormValue("dec"))
        nc.onlypos=ist(r.FormValue("pos"))
        nc.exclusive=ist(r.FormValue("exc"))
        q := new (question)
        q.key,_ = strconv.ParseInt(r.FormValue("qid"),10,64)
        Sget(q)
        nc.q = q
		return nil
}

func delqc_handler(w http.ResponseWriter, r *http.Request) {
	nchan := make (chan error)
	DBchan <- func(DB *sql.DB) func() {
		stmt,err := DB.Prepare("delete from questionscriterion where okey = ? and ikey = ?")
		if err != nil { nchan <- err }
		okey, _ := strconv.Atoi(r.FormValue("nqkey"))
        ikey, _ := strconv.Atoi(r.FormValue("nckey"))
        _,err = stmt.Exec(okey, ikey)
		if err != nil { nchan <- err }
		return func() {
			nchan <- err
		}
	}
	checkErr(<-nchan)

    //go to addserv page again
    questioncreate_handler(w,r)
}


func delc_handler(w http.ResponseWriter, r *http.Request) {
	nchan := make (chan error)
	DBchan <- func(DB *sql.DB) func() {
		stmt,err := DB.Prepare("delete from servicescriterion where okey = ? and ikey = ?")
		if err != nil { nchan <- err }
		okey, _ := strconv.Atoi(r.FormValue("nskey"))
        ikey, _ := strconv.Atoi(r.FormValue("nckey"))
        _,err = stmt.Exec(okey, ikey)
		if err != nil { nchan <- err }
		return func() {
			nchan <- err
		}
	}
	checkErr(<-nchan)

    //go to addserv page again
    servicecreate_handler(w,r)
}
func delq_handler(w http.ResponseWriter, r *http.Request) {
	nchan := make (chan error)
	DBchan <- func(DB *sql.DB) func() {
		stmt,err := DB.Prepare("delete from servicesquestion where okey = ? and ikey = ?")
		if err != nil { nchan <- err }
        okey, _ := strconv.Atoi(r.FormValue("nskey"))
        ikey, _ := strconv.Atoi(r.FormValue("ikey"))
        _,err = stmt.Exec(okey, ikey)
		if err != nil { nchan <- err }
		return func() {
			nchan <- err
		}
	}
	checkErr(<-nchan)

    //go to addserv page again
    servicecreate_handler(w,r)
}

//search
func searchq_handler(w http.ResponseWriter, r *http.Request) {

    if curop(r) != nil {
        if !curop(r).admin{return}
        fmt.Println("searchq"+r.FormValue("q"))
		err,s := Getallqbynamefuzz(r.FormValue("q"))
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
        if !curop(r).admin{return}
        fmt.Println("searchs"+r.FormValue("q"))
		err,s := Getallsbynamefuzz(r.FormValue("q"))
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
        if !curop(r).admin{return}
        fmt.Println("searcho"+r.FormValue("q"))
		err,s := Getallobynamefuzz(r.FormValue("q"))
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
        if !curop(r).admin{return}
        fmt.Println("searchq"+r.FormValue("q"))
        fmt.Println("Service Key is: " + r.FormValue("skey"))
		err,q := Getallqbynamefuzz(r.FormValue("q"))
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
        if !curop(r).admin{return}
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

