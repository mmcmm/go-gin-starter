package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres" // for postgress migrations
	_ "github.com/mattes/migrate/source/file"       // for postgress import from file
)

// Init the db connection & run the migrations
func Init() (conn *pgx.Conn) {
	absPath, _ := filepath.Abs("db/migrations")
	m, err := migrate.New("file://"+absPath, "postgres://username:password@localhost:5432/username?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to run migrations: %v\n", err)
		os.Exit(1)
	}
	m.Steps(2)

	return connect("main")
}

func connect(applicationName string) (conn *pgx.Conn) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              "username",
		Password:          "password",
		Host:              "localhost",
		Port:              5432,
		Database:          "username",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}
