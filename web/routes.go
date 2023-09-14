package web

import (
	"github.com/gorilla/mux"
)

func (s *Server) SetupRoutes() {
	router := mux.NewRouter()

	routes := map[string]RouteHandler{}

	for path, handler := range routes {
		router.Handle(path, s.HandleErrors(handler))
	}

	s.Router = router
}
