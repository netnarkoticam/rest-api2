//nolint:mnd // Стандартный HTTP код
package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
)

func Run() {
	dbConn := getDB()
	defer dbConn.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func getDB() *sql.DB {
	dsn := os.Getenv("PG_URL")
	if dsn == "" {
		log.Fatal("Environment variable PG_URL is not set")
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err = migrate.RunMigrations(dbConn, "./migrations"); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	return dbConn
}
