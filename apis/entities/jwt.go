package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	jwt.RegisteredClaims
	Id                 int64     `json:"id,omitempty"`
	UserName           string    `json:"user_name,omitempty"`
	UpdatedAccountUser time.Time `json:"updated_at,omitempty"`
}

type JwtResponse struct {
	Token              string    `json:"token"`
	UserName           string    `json:"user_name"`
	UserId             int64     `json:"user_id"`
	UpdatedAccountUser time.Time `json:"updated_at,omitempty"`
}

type LoginResponse struct {
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
