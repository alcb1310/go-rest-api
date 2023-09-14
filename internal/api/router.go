package api

import "net/http"

type router struct {
	*http.ServeMux
}

func NewRouter() router {
	return router{http.NewServeMux()}
}
