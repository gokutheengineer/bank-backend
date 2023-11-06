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

func TestUpdateAccount(t *testing.T) {
	account_created := createTestAccount(t)

	updateAccountParams := &UpdateAccountParams{
		ID:      account_created.ID,
		Balance: account_created.Balance + 1000,
	}

	account_updated, err := testStore.UpdateAccount(context.Background(), *updateAccountParams)
	require.NoError(t, err)
	require.NotEmpty(t, account_updated)

	require.Equal(t, account_created.ID, account_updated.ID)
	require.NotEqual(t, account_created, account_updated.Balance)
	require.Equal(t, updateAccountParams.Balance, account_updated.Balance)
	require.Equal(t, account_created.Currency, account_updated.Currency)
	require.Equal(t, account_created.Owner, account_updated.Owner)
	require.WithinDuration(t, account_created.CreatedAt, account_updated.CreatedAt, 0)
}

func TestDeleteAccount(t *testing.T) {
	account_created := createTestAccount(t)

	err := testStore.DeleteAccount(context.Background(), account_created.ID)
	require.NoError(t, err)

	account_retrieved, err := testStore.GetAccount(context.Background(), account_created.ID)
	require.Error(t, err)
	require.Empty(t, account_retrieved)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestAccount(t)
	}

	accounts, err := testStore.ListAccounts(context.Background(), ListAccountsParams{Limit: 7, Offset: 3})
	require.NoError(t, err)
	require.Len(t, accounts, 7)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
