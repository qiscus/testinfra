package testinfra

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

func Postgres(version string) (string, func()) {
	if version == "" {
		version = "latest"
	}

	var (
		dbName     = "testing"
		dbUser     = "user"
		dbPassword = "password"
	)
	resource, err := StartContainer("postgres", version, []string{
		fmt.Sprintf("POSTGRES_DB=%s", dbName),
		fmt.Sprintf("POSTGRES_USER=%s", dbUser),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", dbPassword),
	})
	if err != nil {
		log.Panicln("can't start postgres, err:", err)
	}

	var dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword,
		GetIP(resource, "5432/tcp"), GetPort(resource, "5432/tcp"),
		dbName,
	)
	err = waitContainer(func() error {
		return PostgresCheckFunc(dsn)
	})
	if err != nil {
		log.Panicln("failed to wait postgres to be ready, err:", err)
	}

	return dsn, func() {
		pool.Purge(resource)
	}
}

func PostgresCheckFunc(dsn string) error {
	db, _ := sql.Open("postgres", dsn)
	_, err := db.Exec("select 1;")
	if err == nil {
		db.Close()
		return nil
	}
	db.Close()
	return errors.New("can't query postgres")
}
