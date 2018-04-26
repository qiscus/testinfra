package testinfra_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/qiscus/testinfra"
)

func TestRunDBMigration(t *testing.T) {
	dsn, close := testinfra.Postgres("")
	defer close()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal("cannot create connection to database")
	}

	// run the first migration
	testinfra.RunDBMigration(dsn, "_fixtures/assets")

	_, err = db.Exec(`
		insert into users
		values ('user1', 'password');
	`)
	if err != nil {
		t.Error("error while querying database", err)
	}

	querySelect := `
		select * from users
		where username='user1';
	`
	rows, err := db.Query(querySelect)
	if err != nil {
		t.Error("error while querying database", err)
	}
	count := 0
	for rows.Next() {
		count++
	}
	if count != 1 {
		t.Error("should have 1 row")
	}

	// reset database
	testinfra.RunDBMigration(dsn, "_fixtures/assets")

	rows, err = db.Query(querySelect)
	if err != nil {
		t.Error("error while querying database", err)
	}
	count = 0
	for rows.Next() {
		count++
	}
	if count != 0 {
		t.Error("database must be cleared")
	}
}
