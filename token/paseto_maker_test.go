package token

import (
	"testing"
	"time"

	"github.com/gokutheengineer/bank-backend/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	issuedAt := time.Now()
	duration := 1 * time.Hour
	expiresAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(util.RandomOwner(), "accountOwner", duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotZero(t, payload.ID)
	require.Equal(t, payload.Role, "accountOwner")
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPaseto(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	duration := 1 * time.Hour
	token, payload, err := maker.CreateToken(util.RandomOwner(), "accountOwner", -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Empty(t, payload)
}
