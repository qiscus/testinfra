package testinfra_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/qiscus/testinfra"
)

func TestPostgres(t *testing.T) {
	dsn, close := testinfra.Postgres("9.6")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal("failed to connect to db")
	}

	_, err = db.Query("select 1;")
	if err != nil {
		t.Error("cannot execute query", err)
	}

	close()

	_, err = db.Query("select 1;")
	if err == nil {
		t.Error("should got error database closed")
	}
}
