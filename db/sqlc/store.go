package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Provides combination of all queries and transactions
type Store interface {
	TransferTx(context.Context, TransferTxInputParams) (TransferTxResult, error)
	Querier
}

// Embed Queries struct functionality into QueryStore struct
type QueryStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connectionPool *pgxpool.Pool) Store {
	return &QueryStore{
		Queries:  New(connectionPool),
		connPool: connectionPool,
	}
}
