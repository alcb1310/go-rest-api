package api

import (
	"net/http"
)

func (r router) AuthEndpoints() {
	// register
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("register"))
	})

	// login
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	})

	// logout
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("logout"))
	})

}
