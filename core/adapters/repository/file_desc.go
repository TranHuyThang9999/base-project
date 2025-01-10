package repository

import (
	"context"
	"errors"
	"rices/core/adapters"
	"rices/core/domain"

	"gorm.io/gorm"
)

type fileDescRepository struct {
	db *adapters.Pgsql
}

func NewRepositoryFileDesc(db *adapters.Pgsql) domain.RepositoryFileDescriptors {
	return &fileDescRepository{
		db: db,
	}
}

// AddListFileWithTransaction implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) AddListFileWithTransaction(ctx context.Context, db *gorm.DB, files []*domain.FileDescriptors) error {
	result := db.WithContext(ctx).Create(files)
	return result.Error
}

// AddWithTransaction implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) AddWithTransaction(ctx context.Context, db *gorm.DB, file *domain.FileDescriptors) error {
	if err := db.WithContext(ctx).Create(file).Error; err != nil {
		return err
	}
	return nil
}

// Add implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) Add(ctx context.Context, file *domain.FileDescriptors) error {
	if err := f.db.DB().WithContext(ctx).Create(file).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFileByID implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) DeleteFileByID(ctx context.Context, fileID int64) error {
	var file domain.FileDescriptors
	if err := f.db.DB().WithContext(ctx).Where("id = ?", fileID).Delete(&file).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFileByObjectID implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) DeleteFileByObjectID(ctx context.Context, objectID int64) error {
	var file domain.FileDescriptors
	if err := f.db.DB().WithContext(ctx).Where("object_id = ?", objectID).Delete(&file).Error; err != nil {
		return err
	}
	return nil
}

// ListByObjectID implements domain.RepositoryFileDescriptors.
func (f *fileDescRepository) ListByObjectID(ctx context.Context, objectID int64) ([]*domain.FileDescriptors, error) {
	var files []*domain.FileDescriptors
	if err := f.db.DB().WithContext(ctx).Where("object_id = ?", objectID).Find(&files).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return files, nil
}
