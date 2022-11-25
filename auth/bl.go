package auth

import (
	"context"
	"time"

	"github.com/authentication-service/auth/repository"
	"github.com/authentication-service/security"
	"github.com/authentication-service/spec"
	"github.com/authentication-service/svcerr"
	"github.com/go-kit/log"
	"golang.org/x/crypto/bcrypt"
)

type BL interface {
	Login(ctx context.Context, loginRequest spec.LoginRequest) (spec.AuthUserResponse, error)
}

type bl struct {
	log  log.Logger
	repo repository.Repository
}

func NewBL(log log.Logger, repo repository.Repository) BL {
	return bl{
		log:  log,
		repo: repo,
	}
}

func (bl bl) Login(ctx context.Context, loginRequest spec.LoginRequest) (spec.AuthUserResponse, error) {
	var (
		user        spec.AuthUserResponse
		userDetails spec.User
		err         error
	)
	// close the request after 3 sec
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err = bl.repo.GetUserFromAuth(ctx, loginRequest.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return user, svcerr.ErrInvalidLoginCreds
	}
	userDetails, err = bl.repo.GetUser(ctx, user.UserId)
	if err != nil {
		return user, err
	}
	// generatetoken
	token, err := security.GenerateJWT(user.UserId, userDetails.Role)
	if err != nil {
		return user, err
	}
	user.Id = userDetails.Id
	user.AccessToken = token.AccessToken
	user.RefreshToken = token.RefreshToken
	user.FirstName = userDetails.FirstName
	user.LastName = userDetails.LastName
	user.CreatedAt = userDetails.CreatedAt
	user.UpdatedAt = userDetails.UpdatedAt

	return user, nil
}
