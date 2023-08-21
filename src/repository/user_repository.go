package repository

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"github.com/rzldimam28/auth-grpc/src/entity"
)

var (
	findByEmailQuery = "SELECT u.id, u.username, u.email, u.password FROM tb_user u WHERE u.email = ?"
	insertUserQuery = "INSERT INTO tb_user (id, username, email, password) VALUES (?, ?, ?, ?)"
)

type repository struct {
}

func New() Repository {
	return &repository{}
}

func (ths *repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error) {
	log.Info().Msg("repository.FindByEmail invoked...")

	var user entity.User
	err := tx.QueryRowContext(ctx, findByEmailQuery, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Error().Err(err).Msg("repository.FindByEmail returns err...")
		return nil, err
	}

	log.Info().Str("user_id", user.ID).Msg("repository.FindByEmail returns user...")

	return &user, nil
}

func (ths *repository) InsertUser(ctx context.Context, tx *sql.Tx, user entity.User) (*entity.User, error) {
	log.Info().Msg("repository.InsertUser invoked...")
	
	_, err := tx.ExecContext(ctx, insertUserQuery, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("repository.InsertUser returns err...")
		return nil, err
	}

	log.Info().Str("user_id", user.ID).Msg("repository.InsertUser returns user...")

	return &user, nil
}