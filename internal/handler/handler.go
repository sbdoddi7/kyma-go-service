package handler

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Kyma Go App 🚀")
}
