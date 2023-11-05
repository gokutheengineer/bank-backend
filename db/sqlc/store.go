package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store interface defines functions to execute queries and transactions
type Store interface {
	Querier
	//CreateUserTx()
}

type SQLStore struct {
	connectionPool *pgxpool.Pool
	*Queries
}

func NewStore(connectionPool *pgxpool.Pool) Store {
	return &SQLStore{
		connectionPool: connectionPool,
		Queries:        New(connectionPool),
	}
}
