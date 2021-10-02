package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable user=postgres password=cce16cc03cfb49c9b247fc6faff58fa7"
	}

	os.Exit(m.Run())
}
