package handler

import (
	"net/http"
)

type FarzbookHandler struct{}

func NewFarzbookHandler() *FarzbookHandler {
	return &FarzbookHandler{}
}

func (h *FarzbookHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
