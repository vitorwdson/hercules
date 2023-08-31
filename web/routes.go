package web

import (
	"io"
	"net/http"

	"github.com/vitorwdson/hercules/models/user"
)

func testRoute(w http.ResponseWriter, r *http.Request, user *user.User) {
	io.WriteString(w, "Accessed")
}

func SetupRoutes() {
	http.HandleFunc("/", RequireAuthentication(testRoute))

	http.Handle(
		"/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))),
	)

}
