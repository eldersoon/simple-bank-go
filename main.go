package main

import (
	"database/sql"
	"log"

	"github.com/eldersoon/simple-bank/api"
	db "github.com/eldersoon/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDrive = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank_db?sslmode=disable"
	serverAddress = "0.0.0.0:5000"
)



func main() {
	conn, err := sql.Open(dbDrive, dbSource)
	
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress, conn)

	log.Fatal("Cannot start the server: ", err)
}