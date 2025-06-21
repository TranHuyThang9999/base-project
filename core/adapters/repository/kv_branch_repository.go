package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

type kvBranchRepository struct {
	db *adapters.Pgsql
}

func NewKvBranchRepository(db *adapters.Pgsql) domain.KvBranchRepository {
	return &kvBranchRepository{db: db}
}

func (r *kvBranchRepository) Create(ctx context.Context, branch *domain.KvBranchs) error {
	branch.CreatedDate = time.Now()
	branch.ModifiedDate = time.Now()
	return r.db.DB().WithContext(ctx).Create(branch).Error
}

func (r *kvBranchRepository) Update(ctx context.Context, branch *domain.KvBranchs) error {
	branch.ModifiedDate = time.Now()
	result := r.db.DB().WithContext(ctx).Updates(branch)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("branch not found")
	}
	return nil
}

func (r *kvBranchRepository) Delete(ctx context.Context, appId int) error {
	result := r.db.DB().WithContext(ctx).Delete(&domain.KvBranchs{}, appId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("branch not found")
	}
	return nil
}

func (r *kvBranchRepository) GetByID(ctx context.Context, appId int) (*domain.KvBranchs, error) {
	var branch domain.KvBranchs
	if err := r.db.DB().WithContext(ctx).First(&branch, appId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &branch, nil
}

func (r *kvBranchRepository) GetByTenantAndBranch(ctx context.Context, tenantId int, branchId int) (*domain.KvBranchs, error) {
	var branch domain.KvBranchs
	if err := r.db.DB().WithContext(ctx).
		Where("TenantId = ? AND BranchId = ?", tenantId, branchId).
		First(&branch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &branch, nil
}

func (r *kvBranchRepository) GetActiveBranches(ctx context.Context, tenantId int) ([]*domain.KvBranchs, error) {
	var branches []*domain.KvBranchs
	if err := r.db.DB().WithContext(ctx).
		Where("TenantId = ? AND IsActive = ?", tenantId, true).
		Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *kvBranchRepository) GetByTimeZone(ctx context.Context, timeZoneId string) ([]*domain.KvBranchs, error) {
	var branches []*domain.KvBranchs
	if err := r.db.DB().WithContext(ctx).
		Where("TimeZoneId = ?", timeZoneId).
		Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *kvBranchRepository) UpdateStatus(ctx context.Context, appId int, isActive bool) error {
	result := r.db.DB().WithContext(ctx).
		Model(&domain.KvBranchs{}).
		Where("AppId = ?", appId).
		Updates(map[string]interface{}{
			"IsActive":     isActive,
			"ModifiedDate": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("branch not found")
	}
	return nil
}
