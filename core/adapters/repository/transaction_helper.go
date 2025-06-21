package repository

import (
	"context"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/domain"

	"gorm.io/gorm"
)

// // Example usage:
//
//	err := txHelper.Transaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
//	    if err := tx.Create(&user).Error; err != nil {
//	        return err
//	    }
//	    return tx.Create(&profile).Error
//	})
type TransactionHelperRepository struct {
	db *adapters.Pgsql
}

func NewRepositoryTransaction(db *adapters.Pgsql) domain.RepositoryTransactionHelper {
	return &TransactionHelperRepository{
		db: db,
	}
}

func (t *TransactionHelperRepository) Transaction(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) error) error {
	return t.db.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
