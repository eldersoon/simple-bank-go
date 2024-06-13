package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSouce  = "postgres://root:root@localhost:5432/simple_bank_db?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSouce)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
