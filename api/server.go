package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
)

// Server servers HTTP reqs
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new htpp server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router

	return server
}

// Start starts the http server on the provided address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error: ": err.Error()}
}
