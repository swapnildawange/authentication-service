package repository

import (
	"context"
	"database/sql"

	"github.com/authentication-service/spec"
	"github.com/authentication-service/svcerr"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUserFromAuth(ctx context.Context, email string) (spec.AuthUserResponse, error)
	GetUser(ctx context.Context, userId int) (spec.User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetUserFromAuth(ctx context.Context, email string) (spec.AuthUserResponse, error) {
	var user spec.AuthUserResponse
	getUserQuery := `SELECT id,user_id,email,password from auth where email = $1`
	err := repo.db.GetContext(ctx, &user, getUserQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, svcerr.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (repo *repository) GetUser(ctx context.Context, userId int) (spec.User, error) {
	var user spec.User
	getUserQuery := `SELECT * FROM users WHERE id = $1`
	err := repo.db.GetContext(ctx, &user, getUserQuery, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, svcerr.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
