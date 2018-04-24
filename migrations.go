package testinfra

import (
	"database/sql"

	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
)

func RunDBMigration(db *sql.DB, dialect, assetsDir string) {
	source := &migrate.PackrMigrationSource{
		Box: packr.NewBox(assetsDir),
	}
	migrate.SetTable("migrations")

	_, err := migrate.Exec(db, dialect, source, migrate.Down)
	if err != nil {
		panic(err)
	}

	_, err = migrate.Exec(db, dialect, source, migrate.Up)
	if err != nil {
		panic(err)
	}
}
