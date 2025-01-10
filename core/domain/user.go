package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id           int64          `json:"id,omitempty" gorm:"column:id;type:bigint;primaryKey"`
	UserName     string         `json:"user_name,omitempty" gorm:"column:user_name;unique"`
	Password     string         `json:"password,omitempty"`
	PhoneNumber  string         `json:"phone_number,omitempty"`
	GoogleUserId string         `json:"google_user_id,omitempty"`
	Email        string         `json:"email,omitempty"`
	NickName     string         `json:"nick_name,omitempty"`
	Avatar       string         `json:"avatar,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type RepositoryUser interface {
	Create(ctx context.Context, db *gorm.DB, user *Users) error
	Update(ctx context.Context, user *Users) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*Users, error)
	FindByUsername(ctx context.Context, username string) (*Users, error)
	FindByEmail(ctx context.Context, email string) (*Users, error)
	UpdatePassword(ctx context.Context, id int64, newPassword string) error
}
