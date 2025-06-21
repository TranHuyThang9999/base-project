package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

type timeSheetRepository struct {
	db *adapters.Pgsql
}

func NewTimeSheetRepository(db *adapters.Pgsql) domain.TimeSheetRepository {
	return &timeSheetRepository{db: db}
}

func (r *timeSheetRepository) Create(ctx context.Context, timeSheet *domain.TimeSheets) error {
	return r.db.DB().WithContext(ctx).Create(timeSheet).Error
}

func (r *timeSheetRepository) Update(ctx context.Context, timeSheet *domain.TimeSheets) error {
	result := r.db.DB().WithContext(ctx).Updates(timeSheet)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("timesheet not found")
	}
	return nil
}

func (r *timeSheetRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.DB().WithContext(ctx).Delete(&domain.TimeSheets{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("timesheet not found")
	}
	return nil
}

func (r *timeSheetRepository) GetByID(ctx context.Context, id int64) (*domain.TimeSheets, error) {
	var timeSheet domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).First(&timeSheet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &timeSheet, nil
}

func (r *timeSheetRepository) GetByEmployee(ctx context.Context, employeeId int64) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("EmployeeId = ?", employeeId).
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}

func (r *timeSheetRepository) GetByDateRange(ctx context.Context, branchId int, startDate, endDate time.Time) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND (StartDate <= ? AND EndDate >= ?)",
			branchId, endDate, startDate).
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}

func (r *timeSheetRepository) GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND TenantId = ?", branchId, tenantId).
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}

func (r *timeSheetRepository) GetActiveTimeSheets(ctx context.Context, branchId int, tenantId int) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND TenantId = ? AND TimeSheetStatus = ?",
			branchId, tenantId, 1). // Assuming 1 is active status
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}

func (r *timeSheetRepository) GetByParentID(ctx context.Context, parentId int64) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("ParentTimeSheetId = ?", parentId).
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}

func (r *timeSheetRepository) UpdateTimeSheetStatus(ctx context.Context, id int64, status int8) error {
	result := r.db.DB().WithContext(ctx).
		Model(&domain.TimeSheets{}).
		Where("Id = ?", id).
		Update("TimeSheetStatus", status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("timesheet not found")
	}
	return nil
}

func (r *timeSheetRepository) GetConflictingTimeSheets(ctx context.Context, employeeId int64, startDate, endDate time.Time) ([]*domain.TimeSheets, error) {
	var timeSheets []*domain.TimeSheets
	if err := r.db.DB().WithContext(ctx).
		Where("EmployeeId = ? AND TimeSheetStatus = ? AND "+
			"((StartDate BETWEEN ? AND ?) OR (EndDate BETWEEN ? AND ?) OR "+
			"(StartDate <= ? AND EndDate >= ?))",
			employeeId, 1, startDate, endDate, startDate, endDate, startDate, endDate).
		Find(&timeSheets).Error; err != nil {
		return nil, err
	}
	return timeSheets, nil
}
