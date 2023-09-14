package api

import (
	"net/http"
)

func (r router) UserEndpoints() {
	// profile
	r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user profile"))
	})

}
