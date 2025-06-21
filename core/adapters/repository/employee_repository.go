package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"
	"errors"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *adapters.Pgsql
}

func NewEmployeeRepository(db *adapters.Pgsql) domain.EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(ctx context.Context, employee *domain.Employees) error {
	return r.db.DB().WithContext(ctx).Create(employee).Error
}

func (r *employeeRepository) Update(ctx context.Context, employee *domain.Employees) error {
	result := r.db.DB().WithContext(ctx).Updates(employee)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	return nil
}

func (r *employeeRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.DB().WithContext(ctx).Delete(&domain.Employees{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	return nil
}

func (r *employeeRepository) GetByID(ctx context.Context, id int64) (*domain.Employees, error) {
	var employee domain.Employees
	if err := r.db.DB().WithContext(ctx).First(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*domain.Employees, error) {
	var employees []*domain.Employees
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND TenantId = ?", branchId, tenantId).
		Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) GetActive(ctx context.Context, branchId int, tenantId int) ([]*domain.Employees, error) {
	var employees []*domain.Employees
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND TenantId = ? AND IsActive = ?", branchId, tenantId, true).
		Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
