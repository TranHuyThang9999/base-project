package domain

import (
	"context"
	"time"
)

type KvBranchs struct {
	AppId               int       `gorm:"column:AppId;primaryKey"`
	TenantId            int       `gorm:"column:TenantId"`
	BranchId            int       `gorm:"column:BranchId"`
	Name                string    `gorm:"column:Name"`
	BaseUtcOffsetMinute int       `gorm:"column:BaseUtcOffsetMinute"`
	IsActive            bool      `gorm:"column:IsActive"`
	ModifiedDate        time.Time `gorm:"column:ModifiedDate"`
	CreatedDate         time.Time `gorm:"column:CreatedDate"`
	TimeZoneId          string    `gorm:"column:TimeZoneId"`
}

type KvBranchRepository interface {
	Create(ctx context.Context, branch *KvBranchs) error
	Update(ctx context.Context, branch *KvBranchs) error
	Delete(ctx context.Context, appId int) error
	GetByID(ctx context.Context, appId int) (*KvBranchs, error)
	GetByTenantAndBranch(ctx context.Context, tenantId int, branchId int) (*KvBranchs, error)
	GetActiveBranches(ctx context.Context, tenantId int) ([]*KvBranchs, error)
	GetByTimeZone(ctx context.Context, timeZoneId string) ([]*KvBranchs, error)
	UpdateStatus(ctx context.Context, appId int, isActive bool) error
}
