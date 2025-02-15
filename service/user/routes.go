package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct { // can take any dependencies
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
