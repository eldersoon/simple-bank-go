package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/eldersoon/simple-bank/utils"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../../../")

	if err != nil {
		log.Fatalln("Cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
