package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/netnarkoticam/rest-api2.git/internal/app"
	db "github.com/netnarkoticam/rest-api2.git/internal/service/pgdb"
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
