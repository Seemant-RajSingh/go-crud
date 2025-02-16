package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Seemant-RajSingh/go-crud/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct { // 2nd field same as Store
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter() // pointer to mux Router returned

	userStore := user.NewStore((s.db))        // user is package name, pointer to Store {db *sql.DB} returned
	userHnadler := user.NewHandler(userStore) // pointer to Hanlder (struct with field of type types.UserStore), can add more services like this
	userHnadler.RegisterRouter(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
