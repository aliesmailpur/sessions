package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)


func main(){
	
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie , err := req.Cookie("session")
	if err != nil{
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}

	fmt.Fprintln(res , cookie)
}