package db

import (
	"context"
	"testing"

	"github.com/gokutheengineer/bank-backend/util"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	fromAccount := createTestAccount(t)
	toAccount := createTestAccount(t)

	noOfConcurrentTransferTXs := 10
	errors := make(chan error)
	transferResults := make(chan TransferTxResult)
	transferAmount := util.RandomInt(0, fromAccount.Balance/int64(noOfConcurrentTransferTXs))
	for i := 0; i < noOfConcurrentTransferTXs; i++ {
		go func() {
			inputParams := TransferTxInputParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        transferAmount,
			}
			transferResult, err := testStore.TransferTx(context.Background(), inputParams)
			errors <- err
			transferResults <- transferResult
		}()
	}

	// check transaction results
	for i := 0; i < noOfConcurrentTransferTXs; i++ {
		err := <-errors
		require.NoError(t, err)
		result := <-transferResults
		require.NotEmpty(t, result)

		// account info validation
		require.NotEmpty(t, result.FromAccount)
		require.Equal(t, fromAccount.ID, result.FromAccount.ID)
		require.NotEmpty(t, result.ToAccount)
		require.Equal(t, toAccount.ID, result.ToAccount.ID)

		balanceDifference := transferAmount * int64(i+1)
		require.Equal(t, fromAccount.Balance-balanceDifference, result.FromAccount.Balance)
		require.Equal(t, toAccount.Balance+balanceDifference, result.ToAccount.Balance)

		// transfer validation
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.Amount, transferAmount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// test getting transaction from the store
		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// entries validation
		require.NotEmpty(t, result.FromEntry)
		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, result.FromEntry.Amount, -transferAmount)
		require.Equal(t, result.ToEntry.Amount, transferAmount)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)
		require.NotZero(t, result.ToEntry.ID)
		require.NotZero(t, result.ToEntry.CreatedAt)
		require.Equal(t, result.FromEntry.AccountID, fromAccount.ID)
		require.Equal(t, result.ToEntry.AccountID, toAccount.ID)

		// test getting from entry from the store
		_, err = testStore.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)

		// test getting from entry from the store
		_, err = testStore.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)
	}
}
