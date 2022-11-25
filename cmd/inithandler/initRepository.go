package inithandler

import (
	"github.com/authentication-service/auth/repository"

	"github.com/jmoiron/sqlx"
)

func InitRepository(db *sqlx.DB) repository.Repository {
	// initate repository
	return repository.NewRepository(db)
}
