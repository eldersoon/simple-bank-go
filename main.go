package main

import (
	"database/sql"
	"log"

	"github.com/eldersoon/simple-bank/api"
	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/eldersoon/simple-bank/utils"
	_ "github.com/lib/pq"
)


func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatalln("Cannot load config: ", err)
	}


	conn, err := sql.Open(config.DBDriver, config.DBSource)
	
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddres, conn)

	log.Fatal("Cannot start the server: ", err)
}