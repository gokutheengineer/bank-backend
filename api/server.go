package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	token "github.com/gokutheengineer/bank-backend/token"
	util "github.com/gokutheengineer/bank-backend/util"
)

// Server servers HTTP reqs
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

// NewServer creates a new htpp server
func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	// register 'currency' validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

// Start starts the http server on the provided address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error: ": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.handleCreateUser)
	router.POST("/users/login", server.handleLoginUser)
	router.GET("/users/get", server.handleGetUser)

	router.POST("/accounts", server.handleCreateAccount)
	router.GET("/accounts/:id", server.handleGetAccount)
	router.GET("/accounts", server.handleListAccounts)
	router.POST("/transfer", server.handleCreateTransfer)

	server.router = router
}
