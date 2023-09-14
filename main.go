package main

import (
	"flag"
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

	dbConnection := db.Connect()
	defer dbConnection.Close()

    if *runMigrations {
        db.RunMigrations(dbConnection)
        os.Exit(0)
    }

    server := web.Server {
        DB: dbConnection,
        Port: ":3000",
        DevMode: *devMode,
    }
	server.SetupRoutes()
    server.ServeStaticFiles()
    server.StartServer()
}
