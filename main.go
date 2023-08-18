package main

import (
	"net/http"

	"github.com/vitorwdson/hercules/web"
)

func main() {
	web.SetupRoutes()
	http.ListenAndServe(":3000", nil)
}
