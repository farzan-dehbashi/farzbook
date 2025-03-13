package handler

import (
	"fmt"
	"net/http"
)

type FarzbookHandler struct{
	counter int
}

func NewFarzbookHandler() *FarzbookHandler {
	return &FarzbookHandler{}
}

func (h *FarzbookHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("home"))
	fmt.Println("home")
}

func (h *FarzbookHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login"))
	fmt.Println("login")
}

func (h *FarzbookHandler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("profile"))
	fmt.Println("profile")
}

func (h *FarzbookHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logout"))
	fmt.Println("logout")
}

func (h *FarzbookHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	h.counter ++
	fmt.Println("add", h.counter)

	w.Write([]byte("add"))
}

func (h *FarzbookHandler) Sub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	h.counter --
	fmt.Println("sub", h.counter)

	w.Write([]byte("sub"))
}
