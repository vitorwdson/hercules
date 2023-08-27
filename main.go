package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vitorwdson/hercules/db"
	"github.com/vitorwdson/hercules/web"
)

func main() {
	devMode := flag.Bool("dev", false, "Use develoment mode")
	runMigrations := flag.Bool("migrate", false, "Applies migrations and exits the program")
	flag.Parse()

	if *devMode {
		godotenv.Load()
	}

	dbConnection := db.GetDB()
	defer dbConnection.Close()

    if *runMigrations {
        db.RunMigrations()
        os.Exit(0)
    }

	web.SetupRoutes()
	http.ListenAndServe(":3000", nil)
}
