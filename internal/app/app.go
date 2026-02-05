//nolint:mnd // Стандартный HTTP код
package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq" // иначе не работает
	"github.com/netnarkoticam/rest-api2.git/config"
	"github.com/netnarkoticam/rest-api2.git/internal/app/migrate"
	"github.com/netnarkoticam/rest-api2.git/internal/handlers"
	"github.com/netnarkoticam/rest-api2.git/internal/repo/pgdb"
	service "github.com/netnarkoticam/rest-api2.git/internal/service/pgdb"
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
	userService := service.NewUserService(repo)
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", handlers.RegisterUser(userService))
		v1.GET("/users/:id", handlers.GetUser(userService))
		v1.PUT("/users/:id", handlers.UpdateUser(userService))
		v1.DELETE("/users/:id", handlers.DeleteUser(userService))
	}
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	if binderr := r.Run(":" + cfg.HTTP.Port); binderr != nil {
		log.Error().Err(err).Msg("server start error")
	}
}
