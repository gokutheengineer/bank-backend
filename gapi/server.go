package gapi

import (
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/pb"
)

// Server servers HTTP reqs
type Server struct {
	pb.UnimplementedBankServer
	store db.Store
}

// NewServer creates a new GRPC server.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	return server
}
