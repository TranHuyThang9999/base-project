package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id          int64          `json:"id,omitempty" gorm:"column:id;type:bigint;primaryKey"`
	UserName    string         `json:"user_name,omitempty" gorm:"column:user_name;unique"`
	Password    string         `json:"password,omitempty"`
	Email       string         `json:"email,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type RepositoryUser interface {
}
