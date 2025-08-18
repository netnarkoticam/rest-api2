//nolint:mnd // Стандартный HTTP код
package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/netnarkoticam/rest-api2.git/config"
	"github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	cfg, err := config.Get(viper.New())
	if err != nil {
		log.Error().Err(err).Msg("function error")
		return
	}

	dbConn := getDB(cfg)
	defer dbConn.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":" + cfg.HTTP.Port); err != nil {
		log.Error().Err(err).Msg("server start error")
	}
}

func getDB(cfg config.Config) *sql.DB {
	dsn := (cfg.DB.URL)
	if dsn == "" {
		log.Error().Msg("url cant be found")
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("cant connect to DB")
	}

	if err = migrate.RunMigrations(dbConn, "./migrations"); err != nil {
		log.Error().Err(err).Msg("cant summon migrations")
	}

	return dbConn
}
