package testinfra

import (
	"database/sql"
	"net/url"

	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
)

func RunDBMigration(dsn string, assetsDir string) {
	u, err := url.Parse(dsn)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(u.Scheme, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	source := &migrate.PackrMigrationSource{
		Box: packr.NewBox(assetsDir),
	}
	migrate.SetTable("migrations")

	_, err = migrate.Exec(db, u.Scheme, source, migrate.Down)
	if err != nil {
		panic(err)
	}

	_, err = migrate.Exec(db, u.Scheme, source, migrate.Up)
	if err != nil {
		panic(err)
	}
}
