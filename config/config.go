package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB   DB
	HTTP HTTP
}

type HTTP struct {
	HTTPport string
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

	http, err := getHttp(v)
	if err != nil {
		return Config{}, err
	}

	return Config{

		HTTP: http,
		DB:   db,
	}, nil
}

func getDB(v *viper.Viper) (DB, error) {
	const (
		dbHostPortKey = "POSTGRES_HOSTPORT"
		dbUserKey     = "POSTGRES_USER"
		dbPasswordKey = "POSTGRES_PASSWORD"
		dbNameKey     = "POSTGRES_DB"
		dbURLKey      = "PG_URL"
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
	if !v.IsSet(dbURLKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, dbURLKey)
	}

	db.HostPort = v.GetString(dbHostPortKey)
	db.User = v.GetString(dbUserKey)
	db.Password = v.GetString(dbPasswordKey)
	db.DBName = v.GetString(dbNameKey)

	return db, nil
}

func getHttp(v *viper.Viper) (HTTP, error) {
	const (
		htHTTPKey = "HTTP_PORT"
	)
	var ht HTTP
	if !v.IsSet(htHTTPKey) {
		return ht, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, htHTTPKey)
	}

	ht.HTTPport = v.GetString(htHTTPKey)

	return ht, nil
}
