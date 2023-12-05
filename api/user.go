package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type userResponse struct {
	Username          string    `json:"Username" binding:"required,alphanum"`
	Email             string    `json:"Email" binding:"required,email"`
	Password          string    `json:"password" binding:"required"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
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
		Email:          req.Email,
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
