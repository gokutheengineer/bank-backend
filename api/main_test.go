package api

import (
	"testing"
	"time"

	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: 1 * time.Hour,
	}

	server, err := NewServer(store, config)
	require.NoError(t, err)
	return server
}
