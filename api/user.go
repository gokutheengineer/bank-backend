package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=6,max=16"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	Fullname          string    `json:"fullname"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordUpdatedAt time.Time `json:"password_updated_at"`
}

func (server *Server) handleCreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// hash password and store
	hashedPassword, err := util.HashPasswordBcrypt(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	createUserArgs := db.CreateUserParams{
		Username:       req.Username,
		Fullname:       req.Fullname,
		PasswordHashed: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, createUserArgs)
	if err != nil {
		// TODO: check why the returned error can't be casted to *pq.Error
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := &createUserResponse{
		Username:          user.Username,
		Fullname:          user.Fullname,
		CreatedAt:         user.CreatedAt,
		PasswordUpdatedAt: user.PasswordUpdatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) handleLoginUser(ctx *gin.Context) {

}

type getUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (server *Server) handleGetUser(ctx *gin.Context) {
	// handlles get user op
	var req getUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}
