package pgstore

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	DBName   string
	Host     string
	Password string
	Port     string
	SSLMode  string
	Username string
}

func CreatePGConnection(cfg PostgresConfig) *sql.DB {
	connStr := fmt.Sprintf(`postgresql://%v:%v@%v:%v/%v?sslmode=%v`,
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("sql.Open err", err)
	}
	return db
}
