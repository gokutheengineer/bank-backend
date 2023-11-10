package main

import (
	"context"
	"log"

	"github.com/gokutheengineer/bank-backend/api"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig("configs")
	if err != nil {
		log.Fatal("can't load config file: ", err)
	}

	// establish connection to pgx
	pgxPool, err := pgxpool.New(context.Background(), config.DB_Source)
	if err != nil {
		log.Fatal("can't establish connection to postgres db: ", err)
	}

	store := db.NewStore(pgxPool)

	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start the server")
	}
}
