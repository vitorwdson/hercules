package web

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vitorwdson/hercules/db"
	"github.com/vitorwdson/hercules/models/session"
	"github.com/vitorwdson/hercules/models/user"
)

type ProtectedHandler = func(http.ResponseWriter, *http.Request, *user.User)

func RequireAuthentication(handler ProtectedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        db := db.GetDB()
        id, err := uuid.Parse("01ae42a1-66ea-462b-96ab-c1e2e6ef906c")
        if err != nil {
            fmt.Println("uuid parse error")
            panic(err)
        }
        s, err := session.GetByUUID(db, id)
        if err != nil {
            fmt.Println("session retrieval error")
            panic(err)
        }


        handler(w, r, s.User)
    }
}
