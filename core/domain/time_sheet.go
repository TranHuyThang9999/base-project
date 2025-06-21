package domain

import (
	"context"
	"time"
)

type TimeSheets struct {
	Id                         int64     `gorm:"column:Id;primaryKey"`
	EmployeeId                 int64     `gorm:"column:EmployeeId"`
	StartDate                  time.Time `gorm:"column:StartDate"`
	EndDate                    time.Time `gorm:"column:EndDate"`
	IsRepeat                   bool      `gorm:"column:IsRepeat"`
	RepeatType                 int8      `gorm:"column:RepeatType"`
	RepeatEachDay              int8      `gorm:"column:RepeatEachDay"`
	BranchId                   int       `gorm:"column:BranchId"`
	TenantId                   int       `gorm:"column:TenantId"`
	CreatedBy                  int64     `gorm:"column:CreatedBy"`
	TimeSheetStatus            int8      `gorm:"column:TimeSheetStatus"`
	SaveOnDaysOffOfBranch      bool      `gorm:"column:SaveOnDaysOffOfBranch"`
	SaveOnHoliday              bool      `gorm:"column:SaveOnHoliday"`
	Note                       string    `gorm:"column:Note"`
	AutoGenerateClockingStatus int8      `gorm:"column:AutoGenerateClockingStatus"`
	IsAppliedForNextWeeks      bool      `gorm:"column:IsAppliedForNextWeeks"`
	ParentTimeSheetId          int64     `gorm:"column:ParentTimeSheetId"`
	TimeSheetTrackingType      int8      `gorm:"column:TimeSheetTrackingType"`
}

type TimeSheetRepository interface {
	Create(ctx context.Context, timeSheet *TimeSheets) error
	Update(ctx context.Context, timeSheet *TimeSheets) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*TimeSheets, error)
	GetByEmployee(ctx context.Context, employeeId int64) ([]*TimeSheets, error)
	GetByDateRange(ctx context.Context, branchId int, startDate, endDate time.Time) ([]*TimeSheets, error)
	GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*TimeSheets, error)
	GetActiveTimeSheets(ctx context.Context, branchId int, tenantId int) ([]*TimeSheets, error)
	GetByParentID(ctx context.Context, parentId int64) ([]*TimeSheets, error)
	UpdateTimeSheetStatus(ctx context.Context, id int64, status int8) error
	GetConflictingTimeSheets(ctx context.Context, employeeId int64, startDate, endDate time.Time) ([]*TimeSheets, error)
}
