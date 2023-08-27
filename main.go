package main

import (
	"flag"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vitorwdson/hercules/db"
	"github.com/vitorwdson/hercules/web"
)

func main() {
	devFlag := flag.Bool("dev", false, "Use develoment mode")
    flag.Parse()

	if *devFlag {
		godotenv.Load()
	}

	db := db.GetDB()
	defer db.Close()

	web.SetupRoutes()
	http.ListenAndServe(":3000", nil)
}
