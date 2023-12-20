package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPasswordBlake(t *testing.T) {
	password := RandomString(6)
	hashedPassword := HashPasswordBlake(password)

	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	isValid := VerifyPasswordBlake(password, hashedPassword)
	require.True(t, isValid)

	isValid = VerifyPasswordBlake(RandomString(6), hashedPassword)
	require.False(t, isValid)
}

func TestHashPasswordBcrypt(t *testing.T) {
	password := RandomString(6)
	hashedPassword, err := HashPasswordBcrypt(password)

	fmt.Println("Bcrypt password", password, "hashedPassword", hashedPassword)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	err = VerifyPasswordBcrypt(password, hashedPassword)
	require.NoError(t, err)

	err = VerifyPasswordBcrypt(RandomString(6), hashedPassword)
	require.NoError(t, err)
}
