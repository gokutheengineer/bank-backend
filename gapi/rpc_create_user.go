package gapi

import (
	"context"

	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	"github.com/gokutheengineer/bank-backend/pb"
	"github.com/gokutheengineer/bank-backend/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword := util.HashPassword(req.GetPassword())
	createUserArgs := db.CreateUserParams{
		Username:       req.GetUsername(),
		Fullname:       req.GetFullname(),
		Email:          req.GetEmail(),
		PasswordHashed: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, createUserArgs)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "User is not unique: %v", &pqErr)
			}
		}
		return nil, status.Errorf(codes.AlreadyExists, "Failed to create user")
	}

	rsp := &pb.CreateUserResponse{
		User: &pb.User{
			Username:          user.Username,
			Fullname:          user.Fullname,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordUpdatedAt),
			CreatedAt:         timestamppb.New(user.CreatedAt),
		},
	}

	return rsp, nil
}
