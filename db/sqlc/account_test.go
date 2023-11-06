package db

import (
	"context"
	"testing"

	"github.com/gokutheengineer/bank-backend/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) (account Account) {
	createAccountParams := &CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoneyAmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), *createAccountParams)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, createAccountParams.Owner, account.Owner)
	require.Equal(t, createAccountParams.Balance, account.Balance)
	require.Equal(t, createAccountParams.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account_created := createTestAccount(t)

	account_retrieved, err := testStore.GetAccount(context.Background(), account_created.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account_retrieved)

	require.Equal(t, account_created.ID, account_retrieved.ID)
	require.Equal(t, account_created.Balance, account_retrieved.Balance)
	require.Equal(t, account_created.Currency, account_retrieved.Currency)
	require.Equal(t, account_created.Owner, account_retrieved.Owner)
	require.WithinDuration(t, account_created.CreatedAt, account_retrieved.CreatedAt, 0)
}
