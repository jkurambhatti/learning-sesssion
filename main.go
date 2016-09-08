package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var userDataBase = make(map[string]string)

var (
	cookieHandler = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32))
)

func main() {
	r := NewRouter()
	// r.Methods("GET").Path("/verfied").HandlerFunc(AuthorisedHandler)
	r.Path("/auth").HandlerFunc(VerifyHandler)
	r.Path("/login").HandlerFunc(LoginHandler)
	// r.Methods("GET").Path("/auth").HandlerFunc(VerifyHandler)
	r.Methods("GET").Path("/logout").HandlerFunc(LogoutHandler)

	r.HandleFunc("/", Homehandler)

	fmt.Println("listening at 8080")
	http.ListenAndServe(":8080", r)
}
