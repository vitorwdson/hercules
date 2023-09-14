package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/vitorwdson/hercules/models/session"
	"github.com/vitorwdson/hercules/models/user"
)

type RouteHandler = func(http.ResponseWriter, *http.Request) error
type ProtectedHandler = func(http.ResponseWriter, *http.Request, *user.User) error

func (s Server) HandleErrors(handler RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			panic(err)
		}
	}
}

func (s Server) RequireAuthentication(handler ProtectedHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := uuid.Parse("01ae42a1-66ea-462b-96ab-c1e2e6ef906c")
		if err != nil {
			return err
		}
		s, err := session.GetByUUID(s.DB, id)
		if err != nil {
			return err
		}

		err = handler(w, r, s.User)

		return err
	}
}
