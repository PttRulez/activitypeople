package main

import (
	"antiscoof/internal/config"
	"antiscoof/internal/store/pgstore"
)

func main() {
	cfg := config.MustLoadConfig()

	pgConn := pgstore.CreatePGConnection(cfg.Postgres)

	pgstore.CreateUsersTable(pgConn)
}
