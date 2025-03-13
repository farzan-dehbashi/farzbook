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
	r.mux.HandleFunc("/login", farzbook.Login)
	r.mux.HandleFunc("/profile", farzbook.Profile)
	r.mux.HandleFunc("/logout", farzbook.Logout)
	r.mux.HandleFunc("/add", farzbook.Logout)
	r.mux.HandleFunc("/sub", farzbook.Logout)
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
