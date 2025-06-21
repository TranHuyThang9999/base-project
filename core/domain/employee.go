package domain

import (
	"context"
	"time"
)

type Employees struct {
	Id                  int64      `gorm:"column:Id;primaryKey"`
	NickName            string     `gorm:"column:NickName"`
	Name                string     `gorm:"column:Name"`
	DOB                 *time.Time `gorm:"column:DOB"`
	IsActive            bool       `gorm:"column:IsActive"`
	IdentityNumber      string     `gorm:"column:IdentityNumber"`
	MobilePhone         string     `gorm:"column:MobilePhone"`
	UserId              *int64     `gorm:"column:UserId"`
	DepartmentId        *int64     `gorm:"column:DepartmentId"`
	JobTitleId          *int64     `gorm:"column:JobTitleId"`
	IdentityKeyClocking string     `gorm:"column:IdentityKeyClocking"`
	AccountSecretKey    string     `gorm:"column:AccountSecretKey"`
	TenantId            int        `gorm:"column:TenantId"`
	BranchId            int        `gorm:"column:BranchId"`
	CreatedBy           int64      `gorm:"column:CreatedBy"`
	ModifiedBy          *int64     `gorm:"column:ModifiedBy"`
	StartWorkingDate    *time.Time `gorm:"column:StartWorkingDate"`
	ModelDeviceClocking string     `gorm:"column:ModelDeviceClocking"`
	IsRecoverable       bool       `gorm:"column:IsRecoverable"`
	ReturnToWorkDate    *time.Time `gorm:"column:ReturnToWorkDate"`
}
type EmployeeRepository interface {
	Create(ctx context.Context, employee *Employees) error
	Update(ctx context.Context, employee *Employees) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*Employees, error)
	GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*Employees, error)
	GetActive(ctx context.Context, branchId int, tenantId int) ([]*Employees, error)
}
