package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPasswordBlake(t *testing.T) {
	password := RandomString(6)
	hashedPassword := HashPasswordBlake(password)

	fmt.Println("Blake password", password, "hashedPassword", hashedPassword)

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

	isValid := VerifyPasswordBcrypt(password, hashedPassword)
	require.True(t, isValid)

	isValid = VerifyPasswordBcrypt(RandomString(6), hashedPassword)
	require.False(t, isValid)
}
