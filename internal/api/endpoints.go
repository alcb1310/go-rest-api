package api

import "net/http"

func (r router) AuthEndpoints() {
	// register
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("register"))
	})

	// login

	// logout
}
