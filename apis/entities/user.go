package entities

import (
	"time"
)

type CreateUserRequest struct {
	UserName    string `json:"user_name,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type GetProfile struct {
	Id          int64     `json:"id,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
