package domain

import "context"

type TimeSheetShifts struct {
	Id               int64  `gorm:"column:Id;primaryKey"`
	TimeSheetId      int64  `gorm:"column:TimeSheetId"`
	ShiftIds         string `gorm:"column:ShiftIds"`
	RepeatDaysOfWeek string `gorm:"column:RepeatDaysOfWeek"`
	DayOfWeek        string `gorm:"column:DayOfWeek"`
}
type TimeSheetShiftRepository interface {
	Create(ctx context.Context, timeSheetShift *TimeSheetShifts) error
	Update(ctx context.Context, timeSheetShift *TimeSheetShifts) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*TimeSheetShifts, error)
	GetByTimeSheetID(ctx context.Context, timeSheetId int64) ([]*TimeSheetShifts, error)
	GetByShiftID(ctx context.Context, shiftId int64) ([]*TimeSheetShifts, error)
	GetByDayOfWeek(ctx context.Context, dayOfWeek string) ([]*TimeSheetShifts, error)
	BatchCreate(ctx context.Context, timeSheetShifts []*TimeSheetShifts) error
	DeleteByTimeSheetID(ctx context.Context, timeSheetId int64) error
}
