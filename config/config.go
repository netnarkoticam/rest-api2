package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB     DB
	Server Server
}

type Server struct {
	HTTP string
	URL  string
}

type DB struct {
	HostPort string
	User     string
	Password string
	DBName   string
}

var ErrMissingRequiredConfig = errors.New("missing required config")

func getDB(v *viper.Viper) (DB, error) {
	const (
		hostPortKey = "POSTGRES_PORT"
		userKey     = "POSTGRES_USER"
		passwordKey = "POSTGRES_PASSWORD"
		nameKey     = "POSTGRES_DB"
	)

	var db DB

	if !v.IsSet(hostPortKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, hostPortKey)
	}
	if !v.IsSet(userKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, userKey)
	}
	if !v.IsSet(passwordKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, passwordKey)
	}
	if !v.IsSet(nameKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, nameKey)
	}

	db.HostPort = v.GetString(hostPortKey)
	db.User = v.GetString(userKey)
	db.Password = v.GetString(passwordKey)
	db.DBName = v.GetString(nameKey)

	return db, nil
}

func getServer(v *viper.Viper) (Server, error) {
	const (
		http = "HTTP_PORT"
		url  = "PG_URL"
	)
	var sv Server
	if !v.IsSet(http) {
		return sv, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, http)
	}
	if !v.IsSet(url) {
		return sv, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, url)
	}

	sv.HTTP = v.GetString(http)
	sv.URL = v.GetString(url)

	return sv, nil
}

func Get(v *viper.Viper) (Config, error) {
	v.AutomaticEnv()

	db, err := getDB(v)
	if err != nil {
		return Config{}, err
	}

	server, err := getServer(v)
	if err != nil {
		return Config{}, err
	}

	return Config{

		Server: server,
		DB:     db,
	}, nil
}
