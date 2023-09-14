package web

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorwdson/hercules/models/user"
)

func (s Server) testRoute(w http.ResponseWriter, r *http.Request, user *user.User) {
	io.WriteString(w, "Accessed")
}

func (s *Server) SetupRoutes() {
	router := mux.NewRouter()

	router.Handle("/", s.RequireAuthentication(s.testRoute))

	s.Router = router
}
