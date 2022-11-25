package spec

import "time"

type Config struct {
	WebPort int
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Id        int       `db:"id" json:"id"`
	Email     string    `db:"email" json:"email,omitempty"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Role      int       `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type AuthUserResponse struct {
	Id           int       `db:"id" json:"id"`
	UserId       int       `db:"user_id" json:"-"`
	Password     string    `db:"password" json:"-"`
	Email        string    `db:"email" json:"email"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastName     string    `db:"last_name" json:"last_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
