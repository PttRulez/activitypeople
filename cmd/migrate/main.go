package main

import (
	"antiscoof/internal/config"
	"antiscoof/internal/store/pgstore"
	"os"
)

func main() {
	cfg := config.MustLoadConfig()

	pgConn := pgstore.CreatePGConnection(cfg.Postgres)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		pgstore.CreateUsersTable(pgConn)
		pgstore.CreateStravaInfoTable(pgConn)
	}
	if cmd == "down" {
		pgstore.DropUsersTable(pgConn)
		pgstore.DropStravaInfoTable(pgConn)
	}
}
