package handler

import (
	"context"

	"github.com/rzldimam28/auth-grpc/src/delivery/pb"
	"github.com/rzldimam28/auth-grpc/src/model"
	"github.com/rzldimam28/auth-grpc/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/rs/zerolog/log"

)

type handler struct {
	service service.Service
	pb.UnimplementedUserHandlerServer
}

func NewUserGrpcServer(gserver *grpc.Server, service service.Service) {
	h := &handler{
		service: service,
	}
	pb.RegisterUserHandlerServer(gserver, h)
	reflection.Register(gserver)
}

func (ths *handler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Info().Msg("handler.Login invoked...")

	req := model.LoginRequest{
		Email: in.Email,
		Password: in.Password,
	}
	
	user, err := ths.service.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("handler.Login returns err...")
		res := &pb.LoginResponse{
			Success: "false",
			User: nil,
		}
		return res, nil
	}

	userResponse := &pb.User{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}

	res := &pb.LoginResponse{
		Success: "true",
		User: userResponse,
	}
	return res, nil
}

func (ths *handler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Info().Msg("handler.Register invoked...")

	req := model.RegisterRequest{
		Username: in.Username,
		Email: in.Email,
		Password: in.Password,
	}

	user, err := ths.service.Register(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("handler.Register returns err...")
		res := &pb.RegisterResponse{
			Success: "false",
			User: nil,
		}
		return res, nil
	}

	userResponse := &pb.User{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}

	res := &pb.RegisterResponse{
		Success: "true",
		User: userResponse,
	}
	return res, nil
}