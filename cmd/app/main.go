package main

import (
	"database/sql"
	db "github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
	"github.com/netnarkoticam/rest-api2.git/internal/app"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("PG_URL")
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB")
	}
	defer dbConn.Close()

	db.RunMigrations(dbConn, "/migrations")

	app.RunServer()
}
