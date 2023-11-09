package db

import (
	"context"
	"fmt"
)

type TransferTxInputParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs transfer among two accounts
func (store *QueryStore) TransferTx(ctx context.Context, inputParams TransferTxInputParams) (TransferTxResult, error) {
	result := TransferTxResult{}
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: inputParams.FromAccountID,
			ToAccountID:   inputParams.ToAccountID,
			Amount:        inputParams.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: inputParams.FromAccountID,
			Amount:    -inputParams.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: inputParams.ToAccountID,
			Amount:    inputParams.Amount,
		})
		if err != nil {
			return err
		}

		// To avoid deadlock, we make sure two concourrent transactions always lock the accounts in the same order
		// This can be done by sorting the account IDs and lock them in order, always lock smaller ID first
		if inputParams.FromAccountID < inputParams.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, inputParams.FromAccountID, -inputParams.Amount, result.Transfer.ToAccountID, inputParams.Amount)
			if err != nil {
				return fmt.Errorf("addMoney is failed: %v", err)
			}
		} else {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, inputParams.ToAccountID, inputParams.Amount, result.Transfer.FromAccountID, -inputParams.Amount)
			if err != nil {
				return fmt.Errorf("addMoney is failed: %v", err)
			}
		}

		return err
	})

	return result, err
}

func addMoney(ctx context.Context, queries *Queries, fromAccountID, amount1, toAccountID, amount2 int64) (fromAccount, toAccount Account, err error) {
	fromAccount, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount1,
		ID:     fromAccountID,
	})
	if err != nil {
		return
	}

	toAccount, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount2,
		ID:     toAccountID,
	})

	return
}
