package api

import (
	"go/token"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
)

// Server servers HTTP reqs
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new htpp server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// register 'currency' validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.handleCreateUser)
	router.POST("/users/login", server.handleLoginUser)
	router.GET("/users/get", server.handleGetUser)

	router.POST("/accounts", server.handleCreateAccount)
	router.GET("/accounts/:id", server.handleGetAccount)
	router.GET("/accounts", server.handleListAccounts)
	router.POST("/transfer", server.handleCreateTransfer)

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
