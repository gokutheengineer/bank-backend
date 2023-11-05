package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gokutheengineer/bank-backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../configs")
	if err != nil {
		log.Fatal("can't load config file: ", err)
	}

	// establish connection to pgx
	pgxPool, err := pgxpool.New(context.Background(), config.DB_Source)
	if err != nil {
		log.Fatal("can't establish connection to postgres db: ", err)
	}

	testStore = NewStore(pgxPool)
	os.Exit(m.Run())

}
