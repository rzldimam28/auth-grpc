package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rzldimam28/auth-grpc/src/entity"
	"github.com/rzldimam28/auth-grpc/src/model"
	"github.com/rzldimam28/auth-grpc/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	db *sql.DB
	validate *validator.Validate
	repository repository.Repository
}

func New(db *sql.DB, validate *validator.Validate, repository repository.Repository) Service {
	return &service{
		db: db,
		validate: validate,
		repository: repository,
	}
}

func (ths *service) Login(ctx context.Context, req model.LoginRequest) (*model.UserResponse, error) {
	log.Info().Msg("service.Login invoked...")

	err := ths.validate.Struct(req)
	if err != nil {
		log.Error().Err(err).Msg("service.Login returns err when validate...")
		return nil, err
	}

	tx, err := ths.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("service.Login returns err starting tx...")
		return nil, err
	}

	user, err := ths.repository.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		log.Error().Err(err).Msg("service.Login returns err when calling repository...")
		tx.Rollback()
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil || req.Email != user.Email {
		log.Error().Err(err).Msg("service.Login returns err when comparing password...")
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	res := &model.UserResponse{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	log.Info().Str("res", res.String()).Msg("service.Login returns response...")
	return res, nil
}

func (ths *service) Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error) {
	log.Info().Msg("service.Register invoked...")

	err := ths.validate.Struct(req)
	if err != nil {
		log.Error().Err(err).Msg("service.Register returns err when validate...")
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		log.Error().Err(err).Msg("service.Register returns err when hashing password...")
		return nil, err
	}

	id := uuid.NewString()
	user := entity.User{
		ID: id,
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
	}

	tx, err := ths.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("service.Register returns err starting tx...")
		return nil, err
	}

	newUser, err := ths.repository.InsertUser(ctx, tx, user)
	if err != nil {
		log.Error().Err(err).Msg("service.Register returns err when calling repository...")
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	res := &model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
		Username: newUser.Username,
		Password: newUser.Password,
	}

	log.Info().Str("res", res.String()).Msg("service.Register returns response...")
	return res, nil
}