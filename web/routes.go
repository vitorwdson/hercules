package web

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorwdson/hercules/models/user"
)

func (s Server) testRoute(w http.ResponseWriter, r *http.Request, user *user.User) error {
	io.WriteString(w, "Accessed")

    return nil
}

func (s *Server) SetupRoutes() {
	router := mux.NewRouter()
    
    routes := map[string]RouteHandler {
        "/": s.RequireAuthentication(s.testRoute),
    }


    for path, handler := range routes {
        router.Handle(path, s.HandleErrors(handler))
    }

	s.Router = router
}
