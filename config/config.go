package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var ErrMissingRequiredConfig = errors.New("missing required config")

type Config struct {
	HTTP HTTP
	DB   DB
}

type HTTP struct {
	Port string
}

type DB struct {
	URL string
}

func Get(v *viper.Viper) (Config, error) {
	v.AutomaticEnv()

	db, err := getDB(v)
	if err != nil {
		return Config{}, err
	}

	http, err := getHTTP(v)
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTP: http,
		DB:   db,
	}, nil
}

func getDB(v *viper.Viper) (DB, error) {
	const urlKey = "PG_URL"

	var db DB

	if !v.IsSet(urlKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, urlKey)
	}

	db.URL = v.GetString(urlKey)

	return db, nil
}

func getHTTP(v *viper.Viper) (HTTP, error) {
	const portKey = "HTTP_PORT"

	var http HTTP

	if !v.IsSet(portKey) {
		return http, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, portKey)
	}

	http.Port = v.GetString(portKey)

	return http, nil
}
