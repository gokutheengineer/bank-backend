package db

import (
	"context"
	"testing"

	"github.com/gokutheengineer/bank-backend/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) (user User) {
	createUserParams := &CreateUserParams{
		Username:       util.RandomOwner(),
		PasswordHashed: util.RandomPassword(),
		Fullname:       util.RandomFullname(),
	}

	user, err := testStore.CreateUser(context.Background(), *createUserParams)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Fullname, user.Fullname)
	require.Equal(t, createUserParams.PasswordHashed, user.PasswordHashed)
	require.NotZero(t, user.CreatedAt)

	return
}
