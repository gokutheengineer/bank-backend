package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gokutheengineer/bank-backend/api"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/gapi"
	"github.com/gokutheengineer/bank-backend/pb"
	"github.com/gokutheengineer/bank-backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	runGrpcServer(store, config)
}

func runGrpcServer(store db.Store, config util.Config) {
	bankGrpcServer := gapi.NewServer(store)
	grpcServer := grpc.NewServer()
	pb.RegisterBankServer(grpcServer, bankGrpcServer)
	// allows grpc client to explore which rpcs' are available
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("can not create listener for grpc server")
	}

	fmt.Printf("starting grpc server at: %s\n", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can not start grpc server")
	}
}

func runGinServer(store db.Store, config util.Config) {
	server := api.NewServer(store)
	err := server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("can't start the server")
	}
	fmt.Printf("http server started at: %s", config.HttpServerAddress)
}
