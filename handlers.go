package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Homehandler(rw http.ResponseWriter, req *http.Request) {
	if isCookie(req) {
		http.Redirect(rw, req, "/auth", http.StatusTemporaryRedirect)
		return
	}

	rw.Header().Add("Content Type", "text/html")
	body, _ := ioutil.ReadFile("public/index.html")
	rw.Write(body)
}

func LogoutHandler(rw http.ResponseWriter, req *http.Request) {
	clearsession(rw)
	http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
}

func VerifyHandler(rw http.ResponseWriter, req *http.Request) {
	if isCookie(req) {
		rw.Header().Add("Content Type", "text/html")
		body, _ := ioutil.ReadFile("public/logout.html")
		rw.Write(body)
		return
	}

	username := req.FormValue("uname")
	passwd := req.FormValue("pswd")

	if checkDataBase(username, passwd) {
		userDataBase[username] = passwd

	}

	if username != "zxcv" || passwd != "123" {
		http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
		return
	}

	if setsession(username, rw) {
		fmt.Println("cookie has been set")
		rw.Header().Add("Content Type", "text/html")
		body, _ := ioutil.ReadFile("public/logout.html")
		rw.Write(body)
	} else {
		fmt.Println("error setting session")
		return
	}
}

func setsession(uname string, r http.ResponseWriter) bool {
	value := make(map[string]string)
	value["name"] = uname

	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: 300,
		}

		http.SetCookie(r, cookie)
		return true
	}
	return false
}

func clearsession(r http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(r, cookie)
}

// checks if cookie is present in the request
func isCookie(req *http.Request) bool {
	if cookie, err := req.Cookie("session"); err == nil {
		cookievalue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookievalue); err == nil {
			fmt.Println("cookie found")
			checkDataBase(cookievalue["name"], "")
		}
	}
	return true
}

// verifies if username and pasword matches
func checkDataBase(username string, password string) bool {
	if p, ok := userDataBase[username]; ok {
		if p == password {
			return true
		}
		return false
	}
	return false
}
