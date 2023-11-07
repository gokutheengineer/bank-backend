package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Provides combination of all queries and transactions
type Store interface {
	Querier
}

// Embed Queries struct functionality into QueryStore struct
type QueryStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connectionPool *pgxpool.Pool) Store {
	return &QueryStore{
		Queries: New(connectionPool),
	}
}
