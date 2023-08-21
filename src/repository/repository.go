package repository

import (
	"context"
	"database/sql"

	"github.com/rzldimam28/auth-grpc/src/entity"
)

type Repository interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error)
	InsertUser(ctx context.Context, tx *sql.Tx, user entity.User) (*entity.User, error)
}