package web

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	DB      *sql.DB
	Port    string
	DevMode bool
	Router  *mux.Router
}

func (s Server) ServeStaticFiles() {
	if !s.DevMode {
		return
	}

	http.Handle(
		"/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))),
	)
}

func (s Server) StartServer() {
	http.ListenAndServe(s.Port, s.Router)
}
