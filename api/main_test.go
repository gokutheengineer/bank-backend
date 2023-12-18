package api

import (
	"testing"

	db "github.com/gokutheengineer/bank-backend/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server := NewServer(store)
	return server
}
