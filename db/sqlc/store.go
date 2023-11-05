package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

type QueryStore struct {
	*Queries
}

func NewStore(connectionPool *pgxpool.Pool) Store {
	return &QueryStore{
		Queries: New(connectionPool),
	}
}
