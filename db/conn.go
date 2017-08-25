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
func Init() (conn *pgx.Conn) {
	cred := config.DbCredentials()
	absPath, _ := filepath.Abs("db/migrations")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		cred["user"], cred["password"], cred["host"], cred["user"])

	m, err := migrate.New("file://"+absPath, dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to run migrations: %v\n", err)
		os.Exit(1)
	}
	m.Steps(2)
	return connect("main", cred)
}

func connect(applicationName string, cred map[string]string) (conn *pgx.Conn) {
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
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}
