//nolint:mnd // Стандартный HTTP код
package app

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/netnarkoticam/rest-api2.git/config"
	"github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
	"github.com/spf13/viper"
)

func Run() {
	cfg, err := config.Get(viper.New())
	if err != nil {
		log.Fatalf("Get config %v", cfg)
	}
	dbConn := getDB()
	defer dbConn.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":" + cfg.Server.HTTP); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func getDB() *sql.DB {
	cfg, err := config.Get(viper.New())
	if err != nil {
		log.Fatalf("Get config %v", cfg)
	}
	dsn := (cfg.Server.URL)
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
