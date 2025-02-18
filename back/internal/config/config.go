package config

import (
	"os"

	"github.com/pttrulez/activitypeople/internal/infra/store/pgstore"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpListenPort string
	Postgres       pgstore.PostgresConfig
	JwtSecret      string
	Strava         StravaConfig
}

type StravaConfig struct {
	ClientID     string
	ClientSecret string
	OAuthLink    string
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
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET is not set")
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

	stravaClientID := os.Getenv("STRAVA_CLIENT_ID")
	if stravaClientID == "" {
		panic("STRAVA_CLIENT_ID is not set")
	}
	stravaClientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	if pgUsername == "" {
		panic("STRAVA_CLIENT_SECRET is not set")
	}
	stravaOAuthLink := os.Getenv("STRAVA_OAUTH_LINK")
	if pgUsername == "" {
		panic("STRAVA_OAUTH_LINK is not set")
	}

	var strava StravaConfig = StravaConfig{
		ClientID:     stravaClientID,
		ClientSecret: stravaClientSecret,
		OAuthLink:    stravaOAuthLink,
	}

	return &Config{
		HttpListenPort: httpListenPort,
		Postgres:       pg,
		JwtSecret:      jwtSecret,
		Strava:         strava,
	}
}
