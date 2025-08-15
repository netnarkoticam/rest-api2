package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB       DB
	Postgres Postgres
}

type Postgres struct {
	HTTP string
	URL  string
}

type DB struct {
	HostPort string
	User     string
	Password string
	DBName   string
	URL      string
}

var ErrMissingRequiredConfig = errors.New("missing required config")

func Get(v *viper.Viper) (Config, error) {
	v.AutomaticEnv()

	db, err := getDB(v)
	if err != nil {
		return Config{}, err
	}

	postgres, err := getPostgres(v)
	if err != nil {
		return Config{}, err
	}

	return Config{

		Postgres: postgres,
		DB:       db,
	}, nil
}

func getDB(v *viper.Viper) (DB, error) {
	const (
		dbHostPortKey = "POSTGRES_HOSTPORT"
		dbUserKey     = "POSTGRES_USER"
		dbPasswordKey = "POSTGRES_PASSWORD"
		dbNameKey     = "POSTGRES_DB"
	)

	var db DB

	if !v.IsSet(dbHostPortKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, dbHostPortKey)
	}
	if !v.IsSet(dbUserKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, dbUserKey)
	}
	if !v.IsSet(dbPasswordKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, dbPasswordKey)
	}
	if !v.IsSet(dbNameKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, dbNameKey)
	}

	db.HostPort = v.GetString(dbHostPortKey)
	db.User = v.GetString(dbUserKey)
	db.Password = v.GetString(dbPasswordKey)
	db.DBName = v.GetString(dbNameKey)

	return db, nil
}

func getPostgres(v *viper.Viper) (Postgres, error) {
	const (
		pgHTTPKey = "HTTP_PORT"
		pgURLKey  = "PG_URL"
	)
	var pg Postgres
	if !v.IsSet(pgHTTPKey) {
		return pg, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, pgHTTPKey)
	}
	if !v.IsSet(pgURLKey) {
		return pg, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, pgURLKey)
	}

	pg.HTTP = v.GetString(pgHTTPKey)
	pg.URL = v.GetString(pgURLKey)

	return pg, nil
}
