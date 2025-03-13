package router

import (
	"net/http"

	"farzbook.com/m/handler"
)

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) RegisterRoutes(farzbook *handler.FarzbookHandler) {
	r.mux.HandleFunc("/", farzbook.Home)
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
