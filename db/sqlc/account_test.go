package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createAccountParams := &CreateAccountParams{
		Owner:    "Gokhan",
		Balance:  100,
		Currency: "TRY",
	}

	account, err := testStore.CreateAccount(context.Background(), *createAccountParams)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	fmt.Println("created at: ", account.CreatedAt)
	fmt.Println("currency: ", account.Currency)
	fmt.Println("ID : ", account.ID)
}

func TestGetAccount(t *testing.T) {
	account_id := 3

	account, err := testStore.GetAccount(context.Background(), int64(account_id))
	require.NoError(t, err)
	require.NotEmpty(t, account)

	fmt.Println("created at: ", account.CreatedAt)
	fmt.Println("ID : ", account.Owner)

}
