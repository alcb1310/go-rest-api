package api

import (
	"fmt"
	"net/http"
)

type router struct {
	*http.ServeMux
}

func NewRouter() router {
	return router{http.NewServeMux()}
}

func (r *router) ApiVersion(v int) *router {
	mux := NewRouter()
	var handler http.Handler = mux

	apiPath := fmt.Sprintf("/api/v%d", v)
	r.Handle(apiPath+"/", http.StripPrefix(apiPath, handler))
	return &mux
}

func (r *router) Prefix(path string) router {
	mux := NewRouter()
	var handler http.Handler = mux

	r.Handle(path+"/", http.StripPrefix(path, handler))
	return mux
}
