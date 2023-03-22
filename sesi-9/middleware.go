package main

import "net/http"

const (
	USERNAME = "batman"
	PASSWORD = "secret"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return false
	}

	if ok := (username == USERNAME) && (password == PASSWORD); !ok {
		w.Write([]byte("wrong username / password"))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("only GET is allowed"))
		return false
	}

	return true
}
