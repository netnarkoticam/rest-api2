package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(db *sql.DB, migrationsDir string) {
	if err := goose.SetDialect("postgress"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Failed to apply migrations")
	}
}
