package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres" // for postgress migrations
	_ "github.com/mattes/migrate/source/file"       // for postgress migrate from file
	"github.com/mtdx/case-api/config"
)

// Init the db connection & run the migrations
func Init() (conn *pgx.ConnPool) {
	cred := config.DbCredentials()
	absPath, _ := filepath.Abs("db/migrations")
	pgURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		cred["user"], cred["password"], cred["host"], cred["user"])

	m, err := migrate.New("file://"+absPath, pgURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to run migrations: %v\n", err)
		os.Exit(1)
	}
	m.Steps(2)
	return connectPool("main", cred)
}

func connectPool(applicationName string, cred map[string]string) (conn *pgx.ConnPool) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              cred["user"],
		Password:          cred["password"],
		Host:              cred["host"],
		Port:              5432,
		Database:          cred["user"],
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 30,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection pool: %v\n", err)
		os.Exit(1)
	}
	return pool
}
