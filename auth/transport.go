package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/authentication-service/spec"
	svcerror "github.com/authentication-service/svcerr"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(logger log.Logger, endpoint Endpoints) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
	loginUserHandler := httptransport.NewServer(endpoint.Login, decodeLoginRequest, encodeResponse, options...)

	jwtTokenHandler := httptransport.NewServer(
		endpoint.GenerateJWTToken,
		decodeGenerateTokenReq,
		encodeResponse,
		options...,
	)
	var r = mux.NewRouter()
	r.Methods(http.MethodPost).Path(spec.LoginRequestPath).Handler(loginUserHandler)
	r.Methods(http.MethodPost).Path(spec.GenerateJWTRequestPath).Handler(jwtTokenHandler)

	return r
}

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ = uni.GetTranslator("en")
}

func validateLoginRequest(request spec.LoginRequest) error {

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	err := validate.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(trans))
		}
		return err
	}

	return nil
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var loginReq spec.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		return nil, svcerror.ErrFailedToDecode
	}
	if err = validateLoginRequest(loginReq); err != nil {
		return nil, err
	}
	return loginReq, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err == svcerror.ErrInvalidLoginCreds || err == svcerror.ErrNotAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeGenerateTokenReq(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}
