//nolint:mnd // Стандартный HTTP код
package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"github.com/netnarkoticam/rest-api2.git/config"
	"github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
	"github.com/netnarkoticam/rest-api2.git/internal/handlesrs"
	"github.com/netnarkoticam/rest-api2.git/internal/repo/pgdb"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	
 cfg, err := config.Get(viper.New())
    if err != nil {
        log.Error().Err(err).Msg("function error")
        return
    }
    dbConn, err := sql.Open("postgres", cfg.DB.URL)
    if err != nil {
        log.Error().Err(err).Msg("cant connect to DB")
        return
    }
    if err = dbConn.Ping(); err != nil {
        log.Error().Err(err).Msg("db ping failed")
        return
    }
    defer dbConn.Close()

    if err = migrate.RunMigrations(dbConn, "./migrations"); err != nil {
        log.Error().Err(err).Msg("cant run migrations")
        return
    }

    repo := pgdb.NewUserRepo(dbConn)
    r := gin.Default()
    v1 := r.Group("/api/v1")
    {
        v1.POST("/users", handlesrs.RegisterUser(repo))
    }
    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

    if err := r.Run(":" + cfg.HTTP.Port); err != nil {
        log.Error().Err(err).Msg("server start error")
    }
}

func getDB(cfg config.Config) *sql.DB {
	dsn := (cfg.DB.URL)
	if dsn == "" {
		log.Error().Msg("url can't be found")
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
