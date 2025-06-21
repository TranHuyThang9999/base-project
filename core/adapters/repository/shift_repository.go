package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

type shiftRepository struct {
	db *adapters.Pgsql
}

func NewShiftRepository(db *adapters.Pgsql) domain.ShiftRepository {
	return &shiftRepository{db: db}
}

func (r *shiftRepository) Create(ctx context.Context, shift *domain.Shifts) error {
	shift.ModifiedDate = time.Now()
	return r.db.DB().WithContext(ctx).Create(shift).Error
}

func (r *shiftRepository) Update(ctx context.Context, shift *domain.Shifts) error {
	shift.ModifiedDate = time.Now()
	result := r.db.DB().WithContext(ctx).Updates(shift)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("shift not found")
	}
	return nil
}

func (r *shiftRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.DB().WithContext(ctx).Delete(&domain.Shifts{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("shift not found")
	}
	return nil
}

func (r *shiftRepository) GetByID(ctx context.Context, id int64) (*domain.Shifts, error) {
	var shift domain.Shifts
	if err := r.db.DB().WithContext(ctx).First(&shift, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &shift, nil
}

func (r *shiftRepository) GetByBranch(ctx context.Context, branchId int, tenantId int) ([]*domain.Shifts, error) {
	var shifts []*domain.Shifts
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND TenantId = ?", branchId, tenantId).
		Find(&shifts).Error; err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *shiftRepository) GetByTimeRange(ctx context.Context, branchId int, from int64, to int64) ([]*domain.Shifts, error) {
	var shifts []*domain.Shifts
	if err := r.db.DB().WithContext(ctx).
		Where("BranchId = ? AND (`From` >= ? AND `To` <= ?)", branchId, from, to).
		Find(&shifts).Error; err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *shiftRepository) ListAll(ctx context.Context, tenantId int) ([]*domain.Shifts, error) {
	var shifts []*domain.Shifts
	if err := r.db.DB().WithContext(ctx).
		Where("TenantId = ?", tenantId).
		Find(&shifts).Error; err != nil {
		return nil, err
	}
	return shifts, nil
}

func (r *shiftRepository) UpdateModifiedInfo(ctx context.Context, id int64, modifiedBy int64) error {
	result := r.db.DB().WithContext(ctx).
		Model(&domain.Shifts{}).
		Where("Id = ?", id).
		Updates(map[string]interface{}{
			"ModifiedBy":   modifiedBy,
			"ModifiedDate": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("shift not found")
	}
	return nil
}
