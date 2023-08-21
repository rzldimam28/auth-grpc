package main

import (
	"fmt"
	"net"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	_ "github.com/rzldimam28/auth-grpc/config"
	"github.com/rzldimam28/auth-grpc/config/db"
	handler "github.com/rzldimam28/auth-grpc/src/delivery"
	"github.com/rzldimam28/auth-grpc/src/repository"
	"github.com/rzldimam28/auth-grpc/src/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	PORT = viper.GetString("app_port")
)

func main() {

	db := db.Connect()
	validate := validator.New()

	repo := repository.New()
	usecase := service.New(db, validate, repo)
	grpcServer := grpc.NewServer()

	handler.NewUserGrpcServer(grpcServer, usecase)

	netListen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(fmt.Errorf("failed to listen: %w", err))
	}

	log.Info().Msg("Server running on port " + PORT)

	if err = grpcServer.Serve(netListen); err != nil {
		panic(fmt.Errorf("failed to serve: %w", err))
	}
}