package auth

import (
	"context"

	"github.com/authentication-service/security"
	"github.com/authentication-service/spec"
	"github.com/authentication-service/svcerr"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Endpoints struct {
	Login            endpoint.Endpoint
	GenerateJWTToken endpoint.Endpoint
}

func NewEndpoints(logger log.Logger, bl BL) Endpoints {
	return Endpoints{
		Login:            makeLoginHandler(logger, bl),
		GenerateJWTToken: makeGenerateJWT(logger, bl),
	}
}

func makeGenerateJWT(logger log.Logger, bl BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token, err := security.GenerateJWT(1, 1)
		if err != nil {
			return "", svcerr.ErrFailedToGenerateJWT
		}
		return token, nil
	}
}

func makeLoginHandler(logger log.Logger, bl BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		loginReq := request.(spec.LoginRequest)
		user, err := bl.Login(ctx, loginReq)
		if err != nil {
			logger.Log("[debug]", "failed to login", "err", err)
			_, ok := err.(svcerr.CustomError)
			if ok {
				return nil, err
			}
			return nil, svcerr.ErrLoginFailed
		}
		return user, nil
	}
}
