package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"
	"errors"

	"gorm.io/gorm"
)

type timeSheetShiftRepository struct {
	db *adapters.Pgsql
}

func NewTimeSheetShiftRepository(db *adapters.Pgsql) domain.TimeSheetShiftRepository {
	return &timeSheetShiftRepository{db: db}
}

func (r *timeSheetShiftRepository) Create(ctx context.Context, timeSheetShift *domain.TimeSheetShifts) error {
	return r.db.DB().WithContext(ctx).Create(timeSheetShift).Error
}

func (r *timeSheetShiftRepository) Update(ctx context.Context, timeSheetShift *domain.TimeSheetShifts) error {
	result := r.db.DB().WithContext(ctx).Updates(timeSheetShift)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("timesheet shift not found")
	}
	return nil
}

func (r *timeSheetShiftRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.DB().WithContext(ctx).Delete(&domain.TimeSheetShifts{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("timesheet shift not found")
	}
	return nil
}

func (r *timeSheetShiftRepository) GetByID(ctx context.Context, id int64) (*domain.TimeSheetShifts, error) {
	var timeSheetShift domain.TimeSheetShifts
	if err := r.db.DB().WithContext(ctx).First(&timeSheetShift, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &timeSheetShift, nil
}

func (r *timeSheetShiftRepository) GetByTimeSheetID(ctx context.Context, timeSheetId int64) ([]*domain.TimeSheetShifts, error) {
	var timeSheetShifts []*domain.TimeSheetShifts
	if err := r.db.DB().WithContext(ctx).
		Where("TimeSheetId = ?", timeSheetId).
		Find(&timeSheetShifts).Error; err != nil {
		return nil, err
	}
	return timeSheetShifts, nil
}

func (r *timeSheetShiftRepository) GetByShiftID(ctx context.Context, shiftId int64) ([]*domain.TimeSheetShifts, error) {
	var timeSheetShifts []*domain.TimeSheetShifts
	if err := r.db.DB().WithContext(ctx).
		Where("ShiftIds LIKE ?", "%"+string(shiftId)+"%").
		Find(&timeSheetShifts).Error; err != nil {
		return nil, err
	}
	return timeSheetShifts, nil
}

func (r *timeSheetShiftRepository) GetByDayOfWeek(ctx context.Context, dayOfWeek string) ([]*domain.TimeSheetShifts, error) {
	var timeSheetShifts []*domain.TimeSheetShifts
	if err := r.db.DB().WithContext(ctx).
		Where("DayOfWeek = ?", dayOfWeek).
		Find(&timeSheetShifts).Error; err != nil {
		return nil, err
	}
	return timeSheetShifts, nil
}

func (r *timeSheetShiftRepository) BatchCreate(ctx context.Context, timeSheetShifts []*domain.TimeSheetShifts) error {
	return r.db.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, shift := range timeSheetShifts {
			if err := tx.Create(shift).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *timeSheetShiftRepository) DeleteByTimeSheetID(ctx context.Context, timeSheetId int64) error {
	result := r.db.DB().WithContext(ctx).
		Where("TimeSheetId = ?", timeSheetId).
		Delete(&domain.TimeSheetShifts{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no timesheet shifts found")
	}
	return nil
}
