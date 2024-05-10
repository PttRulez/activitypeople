package config

import (
	"antiscoof/internal/store/pgstore"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpListenPort string
	Postgres       pgstore.PostgresConfig
	SessionSecret  string
	Strava         StravaConfig
	UserSessionKey string
}

type StravaConfig struct {
	ClientID         string
	ClientSecret     string
	OAuthRedirectUrl string
}

func MustLoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	httpListenPort := os.Getenv("HTTP_LISTEN_PORT")
	if httpListenPort == "" {
		panic("HTTP_LISTEN_PORT is not set")
	}
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		panic("SESSION_SECRET is not set")
	}
	userSessionKey := os.Getenv("USER_SESSION_KEY")
	if userSessionKey == "" {
		panic("USER_SESSION_KEY is not set")
	}

	// POSTGRESS SETTINGS
	pgDBName := os.Getenv("PG_DB_NAME")
	if pgDBName == "" {
		panic("PG_DB_NAME is not set")
	}
	pgHost := os.Getenv("PG_HOST")
	if pgHost == "" {
		panic("PG_HOST is not set")
	}
	pgPassword := os.Getenv("PG_PASSWORD")
	if pgPassword == "" {
		panic("PG_PASSWORD is not set")
	}
	pgPort := os.Getenv("PG_PORT")
	if pgPort == "" {
		panic("PG_PORT is not set")
	}
	pgSSLMode := os.Getenv("PG_SSLMODE")
	if pgSSLMode == "" {
		pgSSLMode = "disable"
	}
	pgUsername := os.Getenv("PG_USERNAME")
	if pgUsername == "" {
		panic("PG_USERNAME is not set")
	}

	var pg pgstore.PostgresConfig = pgstore.PostgresConfig{
		DBName:   pgDBName,
		Host:     pgHost,
		Password: pgPassword,
		Port:     pgPort,
		SSLMode:  pgSSLMode,
		Username: pgUsername,
	}

	return &Config{
		HttpListenPort: httpListenPort,
		Postgres:       pg,
		SessionSecret:  sessionSecret,
		UserSessionKey: userSessionKey,
	}
}
