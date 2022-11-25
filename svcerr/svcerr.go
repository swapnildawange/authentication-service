package svcerr

type CustomError interface {
	Error() string
}

type customErrString struct {
	s string
}

func NewCustomError(text string) CustomError {
	return &customErrString{text}
}

func (e *customErrString) Error() string {
	return e.s
}

var (
	ErrFailedToDecode      = NewCustomError("failed to decode request")
	ErrInvalidRequest      = NewCustomError("invalid request")
	ErrAlreadyExists       = NewCustomError("already exists")
	ErrNotFound            = NewCustomError("not found")
	ErrBadRouting          = NewCustomError("inconsistent mapping between route and handler (programmer error)")
	ErrNotAuthorized       = NewCustomError("user is not authorized to access the resources")
	ErrLoginFailed         = NewCustomError("failed to login user")
	ErrInvalidToken        = NewCustomError("invalid JWT token")
	ErrFailedToGenerateJWT = NewCustomError("failed to generate jwt token")

	ErrFailedToGenerateAccessToken  = NewCustomError("failed to generate access token")
	ErrFailedToGenerateRefreshToken = NewCustomError("failed to generate refresh token")
	ErrInvalidLoginCreds            = NewCustomError("email or password is wrong")

	ErrUserNotFound = NewCustomError("user not found")
)
