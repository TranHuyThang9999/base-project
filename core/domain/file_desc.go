package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type FileDescriptors struct {
	Id        int64          `json:"id,omitempty" gorm:"column:id;type:bigint;primaryKey"`
	ObjectID  int64          `json:"object_id,omitempty"`
	Url       string         `json:"url,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (FileDescriptors) TableName() string {
	return "file_descriptors"
}

type RepositoryFileDescriptors interface {
	Add(ctx context.Context, file *FileDescriptors) error
	ListByObjectID(ctx context.Context, objectID int64) ([]*FileDescriptors, error)
	DeleteFileByID(ctx context.Context, fileID int64) error
	DeleteFileByObjectID(ctx context.Context, objectID int64) error
	AddWithTransaction(ctx context.Context, db *gorm.DB, file *FileDescriptors) error
	AddListFileWithTransaction(ctx context.Context, db *gorm.DB, files []*FileDescriptors) error
}
