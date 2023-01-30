package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

type user struct{
	First string
	Last string
	Loggedin bool
}
var db = map[string]user{}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}


func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil{
		id := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
	}


    var u user
	if req.Method == http.MethodPost{
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		li := req.FormValue("loggedin") == "on"
		u = user{f, l, li}
		db[c.Value]= u
	}
	
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil{
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u, ok := db[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}