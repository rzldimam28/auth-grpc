package service

import (
	"context"

	"github.com/rzldimam28/auth-grpc/src/model"
)

type Service interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.UserResponse, error)
	Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error)
	// Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
	// Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)
}