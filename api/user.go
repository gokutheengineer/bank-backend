package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) handleCreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createUserArgs := db.CreateUserParams{
		Username:       req.Username,
		Fullname:       req.Fullname,
		PasswordHashed: req.Password,
	}

	user, err := server.store.CreateUser(ctx, createUserArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) handleLoginUser(ctx *gin.Context) {

}

type getUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (server *Server) handleGetUser(ctx *gin.Context) {
	// handlles get user op
	var req getUserRequest

	fmt.Println("in handleGetUser")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("in handleGetUser err", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}
