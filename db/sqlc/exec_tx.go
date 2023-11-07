package db

import (
	"context"
	"fmt"
)

// execTx takes a callback function and executes it inside a database transaction
func (store *QueryStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	// create a new Queries struct with the transaction
	queries := New(tx)
	err = fn(queries)

	// if error occurs, rollback transaction
	if err != nil {

		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}

		return err
	}

	// if no error occurs, commit transaction
	return tx.Commit(ctx)
}
