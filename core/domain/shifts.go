package domain

import (
	"context"
	"time"
)

type Shifts struct {
	Id            int64     `gorm:"column:Id;primaryKey"`
	Name          string    `gorm:"column:Name"`
	From          int64     `gorm:"column:From"`
	To            int64     `gorm:"column:To"`
	BranchId      int       `gorm:"column:BranchId"`
	CheckInBefore int64     `gorm:"column:CheckInBefore"`
	CheckOutAfter int64     `gorm:"column:CheckOutAfter"`
	TenantId      int       `gorm:"column:TenantId"`
	CreatedBy     int64     `gorm:"column:CreatedBy"`
	ModifiedBy    int64     `gorm:"column:ModifiedBy"`
	ModifiedDate  time.Time `gorm:"column:ModifiedDate"`
}

type ShiftRepository interface {
	Create(ctx context.Context, shift *Shifts) error
	Update(ctx context.Context, shift *Shifts) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*Shifts, error)
	GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*Shifts, error)
	GetByTimeRange(ctx context.Context, branchId int, from int64, to int64) ([]*Shifts, error)
	ListAll(ctx context.Context, tenantId int) ([]*Shifts, error)
	UpdateModifiedInfo(ctx context.Context, id int64, modifiedBy int64) error
}
