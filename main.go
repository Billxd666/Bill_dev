package main

import (
	"database/sql"
	"log"

	"github.com/Billxd666/bill_dev.git/api"
	db "github.com/Billxd666/bill_dev.git/db/sqlc"
	"github.com/Billxd666/bill_dev.git/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	//load env variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	//database conection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	//store and database conection initialized
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	//run server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
