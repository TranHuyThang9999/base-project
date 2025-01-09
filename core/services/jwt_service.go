package services

import (
	"context"
	"fmt"
	"rices/common/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	config *configs.Configs
}

func NewJwtService(config *configs.Configs) *JwtService {
	return &JwtService{
		config: config,
	}
}

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

func (u *JwtService) GenToken(ctx context.Context, userName string, userId int64, updatedAt time.Time) (*JwtResponse, error) {
	expirationDuration, err := time.ParseDuration(u.config.ExpireAccess)
	if err != nil {
		return nil, fmt.Errorf("invalid expiration duration: %v", err)
	}

	claims := User{
		UserName:           userName,
		Id:                 userId,
		UpdatedAccountUser: updatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.config.AccessSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %v", err)
	}

	return &JwtResponse{
		Token:              tokenString,
		UserName:           userName,
		UserId:             userId,
		UpdatedAccountUser: updatedAt,
	}, nil
}

func (u *JwtService) VerifyToken(ctx context.Context, tokenString string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.config.AccessSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*User); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
