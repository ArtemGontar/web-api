package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tabels ...string) {
		if len(tabels) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tabels, ", "))
		}

		db.Close()
	}
}