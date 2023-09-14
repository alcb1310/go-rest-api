package api

import "net/http"

type router struct {
	*http.ServeMux
}

func NewRouter() router {
	return router{http.NewServeMux()}
}

func (r *router) Prefix(path string) {
	mux := NewRouter()
	var handler http.Handler = mux

	r.Handle(path+"/", http.StripPrefix(path, handler))
}
