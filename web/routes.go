package web

import (
	"github.com/gorilla/mux"
	"github.com/vitorwdson/hercules/web/controllers/auth"
)

func (s *Server) SetupRoutes() {
	router := mux.NewRouter()

	routes := map[string]RouteHandler{
        "/register": auth.RegisterIndex,
    }

	for path, handler := range routes {
		router.Handle(path, s.HandleErrors(handler))
	}

	s.Router = router
}
