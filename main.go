package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"sqlcs.sqlc.dev/app/api"
	"sqlcs.sqlc.dev/app/sqlcs"
	"sqlcs.sqlc.dev/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlcs.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
