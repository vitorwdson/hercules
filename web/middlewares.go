package web

import (
	"net/http"

	"github.com/vitorwdson/hercules/models/user"
)

type ProtectedHandler = func(http.ResponseWriter, *http.Request, *user.User)

func RequireAuthentication(handler ProtectedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        user := user.User{}

        handler(w, r, &user)
    }
}
